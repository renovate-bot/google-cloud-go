// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pstest provides a fake Cloud PubSub service for testing. It implements a
// simplified form of the service, suitable for unit tests. It may behave
// differently from the actual service in ways in which the service is
// non-deterministic or unspecified: timing, delivery order, etc.
//
// This package is EXPERIMENTAL and is subject to change without notice.
//
// See the example for usage.
package pstest

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"path"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/internal/testutil"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"go.einride.tech/aip/filtering"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	durpb "google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReactorOptions is a map that Server uses to look up reactors.
// Key is the function name, value is array of reactor for the function.
type ReactorOptions map[string][]Reactor

// Reactor is an interface to allow reaction function to a certain call.
type Reactor interface {
	// React handles the message types and returns results.  If "handled" is false,
	// then the test server will ignore the results and continue to the next reactor
	// or the original handler.
	React(_ interface{}) (handled bool, ret interface{}, err error)
}

// ServerReactorOption is options passed to the server for reactor creation.
type ServerReactorOption struct {
	FuncName string
	Reactor  Reactor
}

type publishResponse struct {
	resp *pb.PublishResponse
	err  error
}

// Server is a fake Pub/Sub server.
type Server struct {
	srv     *testutil.Server
	Addr    string  // The address that the server is listening on.
	GServer GServer // Not intended to be used directly.
}

// GServer is the underlying service implementor. It is not intended to be used
// directly.
type GServer struct {
	pb.UnimplementedPublisherServer
	pb.UnimplementedSubscriberServer
	pb.UnimplementedSchemaServiceServer

	timeNowFunc atomic.Value

	mu             sync.Mutex
	topics         map[string]*topic
	subs           map[string]*subscription
	msgs           []*Message // all messages ever published
	msgsByID       map[string]*Message
	wg             sync.WaitGroup
	nextID         int
	streamTimeout  time.Duration
	reactorOptions ReactorOptions
	// schemas is a map of schemaIDs to a slice of schema revisions.
	// the last element in the slice is the most recent schema.
	schemas map[string][]*pb.Schema

	// PublishResponses is a channel of responses to use for Publish.
	publishResponses chan *publishResponse
	// autoPublishResponse enables the server to automatically generate
	// PublishResponse when publish is called. Otherwise, responses
	// are generated from the publishResponses channel.
	autoPublishResponse bool
}

// NewServer creates a new fake server running in the current process.
func NewServer(opts ...ServerReactorOption) *Server {
	return NewServerWithPort(0, opts...)
}

// NewServerWithPort creates a new fake server running in the current process at
// the specified port.
func NewServerWithPort(port int, opts ...ServerReactorOption) *Server {
	return NewServerWithAddress(fmt.Sprintf("localhost:%d", port), opts...)
}

// NewServerWithAddress creates a new fake server running in the current process
// at the specified address (host and port).
func NewServerWithAddress(address string, opts ...ServerReactorOption) *Server {
	return initNewServer(address, func(*grpc.Server) { /* empty */ }, opts...)
}

// NewServerWithCallback creates new fake server running in the current process
// at the specified port. Before starting the server, the provided callback is
// called to allow caller to register additional fakes into grpc server.
func NewServerWithCallback(port int, callback func(*grpc.Server), opts ...ServerReactorOption) *Server {
	return initNewServer(fmt.Sprintf("localhost:%d", port), callback, opts...)
}

// NewServerByAddressWithCallback creates new fake server running in the current
// process at with the provided address (host and port).
// Before starting the server, the provided callback is called to allow caller
// to register additional fakes into grpc server.
func initNewServer(address string, callback func(*grpc.Server), opts ...ServerReactorOption) *Server {
	srv, err := testutil.NewServerWithAddress(address)
	if err != nil {
		panic(fmt.Sprintf("pstest.initNewServer: %v", err))
	}
	reactorOptions := ReactorOptions{}
	for _, opt := range opts {
		reactorOptions[opt.FuncName] = append(reactorOptions[opt.FuncName], opt.Reactor)
	}
	s := &Server{
		srv:  srv,
		Addr: srv.Addr,
		GServer: GServer{
			topics:              map[string]*topic{},
			subs:                map[string]*subscription{},
			msgsByID:            map[string]*Message{},
			reactorOptions:      reactorOptions,
			publishResponses:    make(chan *publishResponse, 100),
			autoPublishResponse: true,
			schemas:             map[string][]*pb.Schema{},
		},
	}
	s.GServer.timeNowFunc.Store(time.Now)
	pb.RegisterPublisherServer(srv.Gsrv, &s.GServer)
	pb.RegisterSubscriberServer(srv.Gsrv, &s.GServer)
	pb.RegisterSchemaServiceServer(srv.Gsrv, &s.GServer)

	callback(srv.Gsrv)

	srv.Start()
	return s
}

// SetTimeNowFunc registers f as a function to
// be used instead of time.Now for this server.
func (s *Server) SetTimeNowFunc(f func() time.Time) {
	s.GServer.timeNowFunc.Store(f)
}

func (s *GServer) now() time.Time {
	return s.timeNowFunc.Load().(func() time.Time)()
}

// Publish behaves as if the Publish RPC was called with a message with the given
// data and attrs. It returns the ID of the message.
// The topic will be created if it doesn't exist.
//
// Publish panics if there is an error, which is appropriate for testing.
func (s *Server) Publish(topic string, data []byte, attrs map[string]string) string {
	return s.PublishOrdered(topic, data, attrs, "")
}

// PublishOrdered behaves as if the Publish RPC was called with a message with the given
// data, attrs and ordering key. It returns the ID of the message.
// The topic will be created if it doesn't exist.
//
// PublishOrdered panics if there is an error, which is appropriate for testing.
func (s *Server) PublishOrdered(topic string, data []byte, attrs map[string]string, orderingKey string) string {
	const topicPattern = "projects/*/topics/*"
	ok, err := path.Match(topicPattern, topic)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(fmt.Sprintf("topic name must be of the form %q", topicPattern))
	}
	s.GServer.CreateTopic(context.TODO(), &pb.Topic{Name: topic})
	req := &pb.PublishRequest{
		Topic:    topic,
		Messages: []*pb.PubsubMessage{{Data: data, Attributes: attrs, OrderingKey: orderingKey}},
	}
	res, err := s.GServer.Publish(context.TODO(), req)
	if err != nil {
		panic(fmt.Sprintf("pstest.Server.Publish: %v", err))
	}
	return res.MessageIds[0]
}

// AddPublishResponse adds a new publish response to the channel used for
// responding to publish requests.
func (s *Server) AddPublishResponse(pbr *pb.PublishResponse, err error) {
	pr := &publishResponse{}
	if err != nil {
		pr.err = err
	} else {
		pr.resp = pbr
	}
	s.GServer.publishResponses <- pr
}

// SetAutoPublishResponse controls whether to automatically respond
// to messages published or to use user-added responses from the
// publishResponses channel.
func (s *Server) SetAutoPublishResponse(autoPublishResponse bool) {
	s.GServer.mu.Lock()
	defer s.GServer.mu.Unlock()
	s.GServer.autoPublishResponse = autoPublishResponse
}

// ResetPublishResponses resets the buffered publishResponses channel
// with a new buffered channel with the given size.
func (s *Server) ResetPublishResponses(size int) {
	s.GServer.mu.Lock()
	defer s.GServer.mu.Unlock()
	s.GServer.publishResponses = make(chan *publishResponse, size)
}

// SetStreamTimeout sets the amount of time a stream will be active before it shuts
// itself down. This mimics the real service's behavior of closing streams after 30
// minutes. If SetStreamTimeout is never called or is passed zero, streams never shut
// down.
func (s *Server) SetStreamTimeout(d time.Duration) {
	s.GServer.mu.Lock()
	defer s.GServer.mu.Unlock()
	s.GServer.streamTimeout = d
}

// A Message is a message that was published to the server.
type Message struct {
	ID          string
	Data        []byte
	Attributes  map[string]string
	PublishTime time.Time
	Deliveries  int      // number of times delivery of the message was attempted
	Acks        int      // number of acks received from clients
	Modacks     []Modack // modacks received by server for this message
	OrderingKey string
	Topic       string

	// protected by server mutex
	deliveries int
	acks       int
	modacks    []Modack
}

// Modack represents a modack sent to the server.
type Modack struct {
	AckID       string
	AckDeadline int32
	ReceivedAt  time.Time
}

// Messages returns information about all messages ever published.
func (s *Server) Messages() []*Message {
	s.GServer.mu.Lock()
	defer s.GServer.mu.Unlock()

	var msgs []*Message
	for _, m := range s.GServer.msgs {
		m.Deliveries = m.deliveries
		m.Acks = m.acks
		m.Modacks = append([]Modack(nil), m.modacks...)
		msgs = append(msgs, m)
	}
	return msgs
}

// Message returns the message with the given ID, or nil if no message
// with that ID was published.
func (s *Server) Message(id string) *Message {
	s.GServer.mu.Lock()
	defer s.GServer.mu.Unlock()

	m := s.GServer.msgsByID[id]
	if m != nil {
		m.Deliveries = m.deliveries
		m.Acks = m.acks
		m.Modacks = append([]Modack(nil), m.modacks...)
	}
	return m
}

// Wait blocks until all server activity has completed.
func (s *Server) Wait() {
	s.GServer.wg.Wait()
}

// ClearMessages removes all published messages
// from internal containers.
func (s *Server) ClearMessages() {
	s.GServer.mu.Lock()
	s.GServer.msgs = nil
	s.GServer.msgsByID = make(map[string]*Message)
	for _, sub := range s.GServer.subs {
		sub.msgs = map[string]*message{}
	}
	s.GServer.mu.Unlock()
}

// Close shuts down the server and releases all resources.
func (s *Server) Close() error {
	s.srv.Close()
	s.GServer.mu.Lock()
	defer s.GServer.mu.Unlock()
	for _, sub := range s.GServer.subs {
		sub.stop()
	}
	return nil
}

// CreateTopic creates a topic.
func (s *GServer) CreateTopic(_ context.Context, t *pb.Topic) (*pb.Topic, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(t, "CreateTopic", &pb.Topic{}); handled || err != nil {
		return ret.(*pb.Topic), err
	}

	if s.topics[t.Name] != nil {
		return nil, status.Errorf(codes.AlreadyExists, "topic %q", t.Name)
	}
	if err := checkTopicMessageRetention(t.MessageRetentionDuration); err != nil {
		return nil, err
	}
	// Take any ingestion setting to mean the topic is active.
	if t.IngestionDataSourceSettings != nil {
		t.State = pb.Topic_ACTIVE
	}
	top := newTopic(t)
	s.topics[t.Name] = top
	return top.proto, nil
}

// GetTopic gets a Pub/Sub topic.
func (s *GServer) GetTopic(_ context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "GetTopic", &pb.Topic{}); handled || err != nil {
		return ret.(*pb.Topic), err
	}

	if t := s.topics[req.Topic]; t != nil {
		return t.proto, nil
	}
	return nil, status.Errorf(codes.NotFound, "topic %q", req.Topic)
}

// UpdateTopic updates the Pub/Sub topic.
func (s *GServer) UpdateTopic(_ context.Context, req *pb.UpdateTopicRequest) (*pb.Topic, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "UpdateTopic", &pb.Topic{}); handled || err != nil {
		return ret.(*pb.Topic), err
	}

	t := s.topics[req.Topic.Name]
	if t == nil {
		return nil, status.Errorf(codes.NotFound, "topic %q", req.Topic.Name)
	}
	for _, path := range req.UpdateMask.Paths {
		switch path {
		case "labels":
			t.proto.Labels = req.Topic.Labels
		case "message_storage_policy":
			t.proto.MessageStoragePolicy = req.Topic.MessageStoragePolicy
		case "message_retention_duration":
			if err := checkTopicMessageRetention(req.Topic.MessageRetentionDuration); err != nil {
				return nil, err
			}
			t.proto.MessageRetentionDuration = req.Topic.MessageRetentionDuration
		case "schema_settings":
			t.proto.SchemaSettings = req.Topic.SchemaSettings
		case "schema_settings.schema":
			if t.proto.SchemaSettings == nil {
				t.proto.SchemaSettings = &pb.SchemaSettings{}
			}
			t.proto.SchemaSettings.Schema = req.Topic.SchemaSettings.Schema
		case "schema_settings.encoding":
			if t.proto.SchemaSettings == nil {
				t.proto.SchemaSettings = &pb.SchemaSettings{}
			}
			t.proto.SchemaSettings.Encoding = req.Topic.SchemaSettings.Encoding
		case "schema_settings.first_revision_id":
			if t.proto.SchemaSettings == nil {
				t.proto.SchemaSettings = &pb.SchemaSettings{}
			}
			t.proto.SchemaSettings.FirstRevisionId = req.Topic.SchemaSettings.FirstRevisionId
		case "schema_settings.last_revision_id":
			if t.proto.SchemaSettings == nil {
				t.proto.SchemaSettings = &pb.SchemaSettings{}
			}
			t.proto.SchemaSettings.LastRevisionId = req.Topic.SchemaSettings.LastRevisionId
		case "ingestion_data_source_settings":
			if t.proto.IngestionDataSourceSettings == nil {
				t.proto.IngestionDataSourceSettings = &pb.IngestionDataSourceSettings{}
			}
			t.proto.IngestionDataSourceSettings = req.Topic.IngestionDataSourceSettings
			// Take any ingestion setting to mean the topic is active.
			if t.proto.IngestionDataSourceSettings != nil {
				t.proto.State = pb.Topic_ACTIVE
			}
		case "message_transforms":
			t.proto.MessageTransforms = req.GetTopic().GetMessageTransforms()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unknown field name %q", path)
		}
	}
	return t.proto, nil
}

// ListTopics lists the topics in this server.
func (s *GServer) ListTopics(_ context.Context, req *pb.ListTopicsRequest) (*pb.ListTopicsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ListTopics", &pb.ListTopicsResponse{}); handled || err != nil {
		return ret.(*pb.ListTopicsResponse), err
	}

	var names []string
	for n := range s.topics {
		if strings.HasPrefix(n, req.Project) {
			names = append(names, n)
		}
	}
	sort.Strings(names)
	from, to, nextToken, err := testutil.PageBounds(int(req.PageSize), req.PageToken, len(names))
	if err != nil {
		return nil, err
	}
	res := &pb.ListTopicsResponse{NextPageToken: nextToken}
	for i := from; i < to; i++ {
		res.Topics = append(res.Topics, s.topics[names[i]].proto)
	}
	return res, nil
}

// ListTopicSubscriptions lists the subscriptions associated with a topic.
func (s *GServer) ListTopicSubscriptions(_ context.Context, req *pb.ListTopicSubscriptionsRequest) (*pb.ListTopicSubscriptionsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ListTopicSubscriptions", &pb.ListTopicSubscriptionsResponse{}); handled || err != nil {
		return ret.(*pb.ListTopicSubscriptionsResponse), err
	}

	var names []string
	for name, sub := range s.subs {
		if sub.topic.proto.Name == req.Topic {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	from, to, nextToken, err := testutil.PageBounds(int(req.PageSize), req.PageToken, len(names))
	if err != nil {
		return nil, err
	}
	return &pb.ListTopicSubscriptionsResponse{
		Subscriptions: names[from:to],
		NextPageToken: nextToken,
	}, nil
}

// DeleteTopic deletes the topic.
func (s *GServer) DeleteTopic(_ context.Context, req *pb.DeleteTopicRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "DeleteTopic", &emptypb.Empty{}); handled || err != nil {
		return ret.(*emptypb.Empty), err
	}

	t := s.topics[req.Topic]
	if t == nil {
		return nil, status.Errorf(codes.NotFound, "topic %q", req.Topic)
	}
	for _, sub := range s.subs {
		if sub.deadLetterTopic == nil {
			continue
		}
		if req.Topic == sub.deadLetterTopic.proto.Name {
			return nil, status.Errorf(codes.FailedPrecondition, "topic %q used as deadLetter for %s", req.Topic, sub.proto.Name)
		}
	}
	t.stop()
	delete(s.topics, req.Topic)
	return &emptypb.Empty{}, nil
}

// CreateSubscription creates a Pub/Sub subscription.
func (s *GServer) CreateSubscription(_ context.Context, ps *pb.Subscription) (*pb.Subscription, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(ps, "CreateSubscription", &pb.Subscription{}); handled || err != nil {
		return ret.(*pb.Subscription), err
	}

	if ps.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing name")
	}
	if s.subs[ps.Name] != nil {
		return nil, status.Errorf(codes.AlreadyExists, "subscription %q", ps.Name)
	}
	if ps.Topic == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing topic")
	}
	top := s.topics[ps.Topic]
	if top == nil {
		return nil, status.Errorf(codes.NotFound, "topic %q", ps.Topic)
	}
	if err := checkAckDeadline(ps.AckDeadlineSeconds); err != nil {
		return nil, err
	}
	if ps.MessageRetentionDuration == nil {
		ps.MessageRetentionDuration = defaultMessageRetentionDuration
	}
	if err := checkSubMessageRetention(ps.MessageRetentionDuration); err != nil {
		return nil, err
	}
	if ps.PushConfig == nil {
		ps.PushConfig = &pb.PushConfig{}
	} else if ps.PushConfig.Wrapper == nil {
		// Wrapper should default to PubsubWrapper.
		ps.PushConfig.Wrapper = &pb.PushConfig_PubsubWrapper_{
			PubsubWrapper: &pb.PushConfig_PubsubWrapper{},
		}
	}
	// Consider any table set to mean the config is active.
	// We don't convert nil config to empty like with PushConfig above
	// as this mimics the live service behavior.
	if ps.GetBigqueryConfig() != nil && ps.GetBigqueryConfig().GetTable() != "" {
		ps.BigqueryConfig.State = pb.BigQueryConfig_ACTIVE
	}
	if ps.CloudStorageConfig != nil && ps.CloudStorageConfig.Bucket != "" {
		ps.CloudStorageConfig.State = pb.CloudStorageConfig_ACTIVE
	}
	ps.TopicMessageRetentionDuration = top.proto.MessageRetentionDuration
	var deadLetterTopic *topic
	if ps.DeadLetterPolicy != nil {
		dlTopic, ok := s.topics[ps.DeadLetterPolicy.DeadLetterTopic]
		if !ok {
			return nil, status.Errorf(codes.NotFound, "deadLetter topic %q", ps.DeadLetterPolicy.DeadLetterTopic)
		}
		deadLetterTopic = dlTopic
	}

	sub := newSubscription(top, &s.mu, s.now, deadLetterTopic, ps)
	if ps.Filter != "" {
		filter, err := parseFilter(ps.Filter)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "bad filter: %v", err)
		}
		sub.filter = &filter
	}

	top.subs[ps.Name] = sub
	s.subs[ps.Name] = sub
	sub.start(&s.wg)
	return ps, nil
}

// Can be set for testing.
var minAckDeadlineSecs int32

// SetMinAckDeadline changes the minack deadline to n. Must be
// greater than or equal to 1 second. Remember to reset this value
// to the default after your test changes it. Example usage:
//
//	pstest.SetMinAckDeadlineSecs(1)
//	defer pstest.ResetMinAckDeadlineSecs()
func SetMinAckDeadline(n time.Duration) {
	if n < time.Second {
		panic("SetMinAckDeadline expects a value greater than 1 second")
	}

	minAckDeadlineSecs = int32(n / time.Second)
}

// ResetMinAckDeadline resets the minack deadline to the default.
func ResetMinAckDeadline() {
	minAckDeadlineSecs = 10
}

func checkAckDeadline(ads int32) error {
	if ads < minAckDeadlineSecs || ads > 600 {
		// PubSub service returns Unknown.
		return status.Errorf(codes.Unknown, "bad ack_deadline_seconds: %d", ads)
	}
	return nil
}

const (
	minTopicMessageRetentionDuration = 10 * time.Minute
	// 31 days is the maximum topic supported duration (https://cloud.google.com/pubsub/docs/replay-overview#configuring_message_retention)
	maxTopicMessageRetentionDuration = 31 * 24 * time.Hour
	minSubMessageRetentionDuration   = 10 * time.Minute
	// 7 days is the maximum subscription supported duration (https://cloud.google.com/pubsub/docs/replay-overview#configuring_message_retention)
	maxSubMessageRetentionDuration = 7 * 24 * time.Hour
)

var defaultMessageRetentionDuration = durpb.New(168 * time.Hour) // default is 7 days

func checkTopicMessageRetention(pmrd *durpb.Duration) error {
	if pmrd == nil {
		return nil
	}
	mrd := pmrd.AsDuration()
	if mrd < minTopicMessageRetentionDuration || mrd > maxTopicMessageRetentionDuration {
		return status.Errorf(codes.InvalidArgument, "bad message_retention_duration %+v", pmrd)
	}
	return nil
}

func checkSubMessageRetention(pmrd *durpb.Duration) error {
	if pmrd == nil {
		return nil
	}
	mrd := pmrd.AsDuration()
	if mrd < minSubMessageRetentionDuration || mrd > maxSubMessageRetentionDuration {
		return status.Errorf(codes.InvalidArgument, "bad message_retention_duration %+v", pmrd)
	}
	return nil
}

// GetSubscription fetches an existing Pub/Sub subscription details.
func (s *GServer) GetSubscription(_ context.Context, req *pb.GetSubscriptionRequest) (*pb.Subscription, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "GetSubscription", &pb.Subscription{}); handled || err != nil {
		return ret.(*pb.Subscription), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		return nil, err
	}
	return sub.proto, nil
}

// UpdateSubscription updates an existing Pub/Sub subscription.
func (s *GServer) UpdateSubscription(_ context.Context, req *pb.UpdateSubscriptionRequest) (*pb.Subscription, error) {
	if req.Subscription == nil {
		return nil, status.Errorf(codes.InvalidArgument, "missing subscription")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "UpdateSubscription", &pb.Subscription{}); handled || err != nil {
		return ret.(*pb.Subscription), err
	}

	sub, err := s.findSubscription(req.Subscription.Name)
	if err != nil {
		return nil, err
	}
	for _, path := range req.UpdateMask.Paths {
		switch path {
		case "push_config":
			sub.proto.PushConfig = req.Subscription.PushConfig

		case "bigquery_config":
			// If bq config is nil here, it will be cleared.
			// Otherwise, we'll consider the subscription active if any table is set.
			sub.proto.BigqueryConfig = req.GetSubscription().GetBigqueryConfig()
			if sub.proto.GetBigqueryConfig() != nil {
				if sub.proto.GetBigqueryConfig().GetTable() != "" {
					sub.proto.BigqueryConfig.State = pb.BigQueryConfig_ACTIVE
				} else {
					return nil, status.Errorf(codes.InvalidArgument, "table must be provided")
				}
			}

		case "cloud_storage_config":
			sub.proto.CloudStorageConfig = req.GetSubscription().GetCloudStorageConfig()
			// As long as the storage config is not nil, we assume it's valid
			// without additional checks.
			if sub.proto.GetCloudStorageConfig() != nil {
				sub.proto.CloudStorageConfig.State = pb.CloudStorageConfig_ACTIVE
			}

		case "ack_deadline_seconds":
			a := req.Subscription.AckDeadlineSeconds
			if err := checkAckDeadline(a); err != nil {
				return nil, err
			}
			sub.proto.AckDeadlineSeconds = a

		case "retain_acked_messages":
			sub.proto.RetainAckedMessages = req.Subscription.RetainAckedMessages

		case "message_retention_duration":
			if err := checkSubMessageRetention(req.Subscription.MessageRetentionDuration); err != nil {
				return nil, err
			}
			sub.proto.MessageRetentionDuration = req.Subscription.MessageRetentionDuration

		case "labels":
			sub.proto.Labels = req.Subscription.Labels

		case "expiration_policy":
			sub.proto.ExpirationPolicy = req.Subscription.ExpirationPolicy

		case "dead_letter_policy":
			sub.proto.DeadLetterPolicy = req.Subscription.DeadLetterPolicy
			if sub.proto.DeadLetterPolicy != nil {
				dlTopic, ok := s.topics[sub.proto.DeadLetterPolicy.DeadLetterTopic]
				if !ok {
					return nil, status.Errorf(codes.NotFound, "topic %q", sub.proto.DeadLetterPolicy.DeadLetterTopic)
				}
				sub.deadLetterTopic = dlTopic
			}

		case "retry_policy":
			sub.proto.RetryPolicy = req.Subscription.RetryPolicy

		case "filter":
			filter, err := parseFilter(req.Subscription.Filter)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "bad filter: %v", err)
			}
			sub.filter = &filter
			sub.proto.Filter = req.Subscription.Filter

		case "enable_exactly_once_delivery":
			sub.proto.EnableExactlyOnceDelivery = req.Subscription.EnableExactlyOnceDelivery
			for _, st := range sub.streams {
				st.enableExactlyOnceDelivery = req.Subscription.EnableExactlyOnceDelivery
			}
		case "message_transforms":
			sub.proto.MessageTransforms = req.GetSubscription().GetMessageTransforms()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unknown field name %q", path)
		}
	}
	return sub.proto, nil
}

// ListSubscriptions lists the Pub/Sub subscriptions in this server.
func (s *GServer) ListSubscriptions(_ context.Context, req *pb.ListSubscriptionsRequest) (*pb.ListSubscriptionsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ListSubscriptions", &pb.ListSubscriptionsResponse{}); handled || err != nil {
		return ret.(*pb.ListSubscriptionsResponse), err
	}

	var names []string
	for name := range s.subs {
		if strings.HasPrefix(name, req.Project) {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	from, to, nextToken, err := testutil.PageBounds(int(req.PageSize), req.PageToken, len(names))
	if err != nil {
		return nil, err
	}
	res := &pb.ListSubscriptionsResponse{NextPageToken: nextToken}
	for i := from; i < to; i++ {
		res.Subscriptions = append(res.Subscriptions, s.subs[names[i]].proto)
	}
	return res, nil
}

// DeleteSubscription deletes the Pub/Sub subscription.
func (s *GServer) DeleteSubscription(_ context.Context, req *pb.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "DeleteSubscription", &emptypb.Empty{}); handled || err != nil {
		return ret.(*emptypb.Empty), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		return nil, err
	}
	sub.stop()
	delete(s.subs, req.Subscription)
	sub.topic.deleteSub(sub)
	return &emptypb.Empty{}, nil
}

// DetachSubscription detaches the subscription from the topic.
func (s *GServer) DetachSubscription(_ context.Context, req *pb.DetachSubscriptionRequest) (*pb.DetachSubscriptionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "DetachSubscription", &pb.DetachSubscriptionResponse{}); handled || err != nil {
		return ret.(*pb.DetachSubscriptionResponse), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		return nil, err
	}
	sub.topic.deleteSub(sub)
	return &pb.DetachSubscriptionResponse{}, nil
}

// Publish sends a message to the topic.
func (s *GServer) Publish(_ context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "Publish", &pb.PublishResponse{}); handled || err != nil {
		return ret.(*pb.PublishResponse), err
	}

	if req.Topic == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing topic")
	}
	top := s.topics[req.Topic]
	if top == nil {
		return nil, status.Errorf(codes.NotFound, "topic %q", req.Topic)
	}

	if !s.autoPublishResponse {
		r := <-s.publishResponses
		if r.err != nil {
			return nil, r.err
		}
		return r.resp, nil
	}

	var ids []string
	for _, pm := range req.Messages {
		id := fmt.Sprintf("m%d", s.nextID)
		s.nextID++
		pm.MessageId = id
		pubTime := s.now()
		tsPubTime := timestamppb.New(pubTime)
		pm.PublishTime = tsPubTime
		m := &Message{
			ID:          id,
			Data:        pm.Data,
			Attributes:  pm.Attributes,
			PublishTime: pubTime,
			OrderingKey: pm.OrderingKey,
			Topic:       req.Topic,
		}
		top.publish(pm, m)
		ids = append(ids, id)
		s.msgs = append(s.msgs, m)
		s.msgsByID[id] = m
	}
	return &pb.PublishResponse{MessageIds: ids}, nil
}

type topic struct {
	proto *pb.Topic
	subs  map[string]*subscription
}

func newTopic(pt *pb.Topic) *topic {
	return &topic{
		proto: pt,
		subs:  map[string]*subscription{},
	}
}

func (t *topic) stop() {
	for _, sub := range t.subs {
		sub.proto.Topic = "_deleted-topic_"
	}
}

func (t *topic) deleteSub(sub *subscription) {
	delete(t.subs, sub.proto.Name)
}

func (t *topic) publish(pm *pb.PubsubMessage, m *Message) {
	for _, s := range t.subs {
		s.msgs[pm.MessageId] = &message{
			publishTime: m.PublishTime,
			proto: &pb.ReceivedMessage{
				AckId:   pm.MessageId,
				Message: pm,
			},
			deliveries:  &m.deliveries,
			acks:        &m.acks,
			streamIndex: -1,
		}
	}
}

type subscription struct {
	topic           *topic
	deadLetterTopic *topic
	mu              *sync.Mutex // the server mutex, here for convenience
	proto           *pb.Subscription
	ackTimeout      time.Duration
	msgs            map[string]*message // unacked messages by message ID
	streams         []*stream
	done            chan struct{}
	timeNowFunc     func() time.Time
	filter          *filtering.Filter
}

func newSubscription(t *topic, mu *sync.Mutex, timeNowFunc func() time.Time, deadLetterTopic *topic, ps *pb.Subscription) *subscription {
	at := time.Duration(ps.AckDeadlineSeconds) * time.Second
	if at == 0 {
		at = 10 * time.Second
	}
	ps.State = pb.Subscription_ACTIVE
	sub := &subscription{
		topic:           t,
		deadLetterTopic: deadLetterTopic,
		mu:              mu,
		proto:           ps,
		ackTimeout:      at,
		msgs:            map[string]*message{},
		done:            make(chan struct{}),
		timeNowFunc:     timeNowFunc,
	}
	return sub
}

func (s *subscription) start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-s.done:
				return
			case <-time.After(10 * time.Millisecond):
				s.deliver()
			}
		}
	}()
}

func (s *subscription) stop() {
	close(s.done)
}

// Acknowledge marks the message as acknowleged.
func (s *GServer) Acknowledge(_ context.Context, req *pb.AcknowledgeRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "Acknowledge", &emptypb.Empty{}); handled || err != nil {
		return ret.(*emptypb.Empty), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		return nil, err
	}
	for _, id := range req.AckIds {
		sub.ack(id)
	}
	return &emptypb.Empty{}, nil
}

// ModifyAckDeadline modifies the ack deadline of the message.
func (s *GServer) ModifyAckDeadline(_ context.Context, req *pb.ModifyAckDeadlineRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ModifyAckDeadline", &emptypb.Empty{}); handled || err != nil {
		return ret.(*emptypb.Empty), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	for _, id := range req.AckIds {
		s.msgsByID[id].modacks = append(s.msgsByID[id].modacks, Modack{AckID: id, AckDeadline: req.AckDeadlineSeconds, ReceivedAt: now})
	}
	dur := secsToDur(req.AckDeadlineSeconds)
	for _, id := range req.AckIds {
		sub.modifyAckDeadline(id, dur)
	}
	return &emptypb.Empty{}, nil
}

// Pull returns a list of unacknowledged messages from a subscription.
func (s *GServer) Pull(ctx context.Context, req *pb.PullRequest) (*pb.PullResponse, error) {
	s.mu.Lock()

	if handled, ret, err := s.runReactor(req, "Pull", &pb.PullResponse{}); handled || err != nil {
		s.mu.Unlock()
		return ret.(*pb.PullResponse), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		s.mu.Unlock()
		return nil, err
	}
	max := int(req.MaxMessages)
	if max < 0 {
		s.mu.Unlock()
		return nil, status.Error(codes.InvalidArgument, "MaxMessages cannot be negative")
	}
	if max == 0 { // MaxMessages not specified; use a default.
		max = 1000
	}
	msgs := sub.pull(max)
	s.mu.Unlock()
	// Implement the spec from the pubsub proto:
	// "If ReturnImmediately set to true, the system will respond immediately even if
	// it there are no messages available to return in the `Pull` response.
	// Otherwise, the system may wait (for a bounded amount of time) until at
	// least one message is available, rather than returning no messages."
	if len(msgs) == 0 && !req.ReturnImmediately {
		// Wait for a short amount of time for a message.
		// TODO: signal when a message arrives, so we don't wait the whole time.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(500 * time.Millisecond):
			s.mu.Lock()
			msgs = sub.pull(max)
			s.mu.Unlock()
		}
	}
	return &pb.PullResponse{ReceivedMessages: msgs}, nil
}

// StreamingPull return a stream to pull messages from a subscription.
func (s *GServer) StreamingPull(sps pb.Subscriber_StreamingPullServer) error {
	// Receive initial message configuring the pull.
	req, err := sps.Recv()
	if err != nil {
		return err
	}
	s.mu.Lock()
	sub, err := s.findSubscription(req.Subscription)
	s.mu.Unlock()
	if err != nil {
		return err
	}
	// Create a new stream to handle the pull.
	st := sub.newStream(sps, s.streamTimeout)
	st.ackTimeout = time.Duration(req.StreamAckDeadlineSeconds) * time.Second
	err = st.pull(&s.wg)
	sub.deleteStream(st)
	return err
}

// Seek updates a subscription to a specific point in time or snapshot.
func (s *GServer) Seek(ctx context.Context, req *pb.SeekRequest) (*pb.SeekResponse, error) {
	// Only handle time-based seeking for now.
	// This fake doesn't deal with snapshots.
	var target time.Time
	switch v := req.Target.(type) {
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "missing Seek target type")
	case *pb.SeekRequest_Time:
		target = v.Time.AsTime()
	default:
		return nil, status.Errorf(codes.Unimplemented, "unhandled Seek target type %T", v)
	}

	// The entire server must be locked while doing the work below,
	// because the messages don't have any other synchronization.
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "Seek", &pb.SeekResponse{}); handled || err != nil {
		return ret.(*pb.SeekResponse), err
	}

	sub, err := s.findSubscription(req.Subscription)
	if err != nil {
		return nil, err
	}
	// Drop all messages from sub that were published before the target time.
	for id, m := range sub.msgs {
		if m.publishTime.Before(target) {
			delete(sub.msgs, id)
			(*m.acks)++
		}
	}
	// Un-ack any already-acked messages after this time;
	// redelivering them to the subscription is the closest analogue here.
	for _, m := range s.msgs {
		if m.PublishTime.Before(target) {
			continue
		}
		sub.msgs[m.ID] = &message{
			publishTime: m.PublishTime,
			proto: &pb.ReceivedMessage{
				AckId: m.ID,
				// This was not preserved!
				//Message: pm,
			},
			deliveries:  &m.deliveries,
			acks:        &m.acks,
			streamIndex: -1,
		}
	}
	return &pb.SeekResponse{}, nil
}

// Gets a subscription that must exist.
// Must be called with the lock held.
func (s *GServer) findSubscription(name string) (*subscription, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing subscription")
	}
	sub := s.subs[name]
	if sub == nil {
		return nil, status.Errorf(codes.NotFound, "subscription %s", name)
	}
	return sub, nil
}

// Must be called with the lock held.
func (s *subscription) pull(max int) []*pb.ReceivedMessage {
	now := s.timeNowFunc()
	s.maintainMessages(now)
	var msgs []*pb.ReceivedMessage
	filterMsgs(s.msgs, s.filter)
	for id, m := range orderMsgs(s.msgs, s.proto.EnableMessageOrdering) {
		if m.outstanding() {
			continue
		}
		if s.deadLetterCandidate(m) {
			s.ack(id)
			s.publishToDeadLetter(m)
			continue
		}
		(*m.deliveries)++
		if s.proto.DeadLetterPolicy != nil {
			m.proto.DeliveryAttempt = int32(*m.deliveries)
		}
		m.ackDeadline = now.Add(s.ackTimeout)
		msgs = append(msgs, m.proto)
		if len(msgs) >= max {
			break
		}
	}
	return msgs
}

func orderMsgs(msgs map[string]*message, enableMessageOrdering bool) map[string]*message {
	if !enableMessageOrdering {
		return msgs
	}
	result := make(map[string]*message)

	type msg struct {
		id string
		m  *message
	}
	orderingKeyMap := make(map[string]msg)
	for id, m := range msgs {
		orderingKey := m.proto.Message.OrderingKey
		if orderingKey == "" {
			orderingKey = id
		}
		if val, ok := orderingKeyMap[orderingKey]; !ok || m.publishTime.Before(val.m.publishTime) {
			orderingKeyMap[orderingKey] = msg{m: m, id: id}
		}
	}
	for _, val := range orderingKeyMap {
		result[val.id] = val.m
	}
	return result
}

func filterMsgs(msgs map[string]*message, filter *filtering.Filter) {
	if filter == nil {
		return
	}

	filterByAttrs(msgs, filter, func(m *message) messageAttrs {
		return m.proto.Message.Attributes
	})
}

func (s *subscription) deliver() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := s.timeNowFunc()
	s.maintainMessages(now)
	// Try to deliver each remaining message.
	curIndex := 0
	filterMsgs(s.msgs, s.filter)
	for id, m := range orderMsgs(s.msgs, s.proto.EnableMessageOrdering) {
		if m.outstanding() {
			continue
		}
		if s.deadLetterCandidate(m) {
			s.ack(id)
			s.publishToDeadLetter(m)
			continue
		}
		// If the message was never delivered before, start with the stream at
		// curIndex. If it was delivered before, start with the stream after the one
		// that owned it.
		if m.streamIndex < 0 {
			delIndex, ok := s.tryDeliverMessage(m, curIndex, now)
			if !ok {
				break
			}
			curIndex = delIndex + 1
			m.streamIndex = curIndex
		} else {
			delIndex, ok := s.tryDeliverMessage(m, m.streamIndex, now)
			if !ok {
				break
			}
			m.streamIndex = delIndex
		}
	}
}

// tryDeliverMessage attempts to deliver m to the stream at index i. If it can't, it
// tries streams i+1, i+2, ..., wrapping around. Once it's tried all streams, it
// exits.
//
// It returns the index of the stream it delivered the message to, or 0, false if
// it didn't deliver the message.
//
// Must be called with the lock held.
func (s *subscription) tryDeliverMessage(m *message, start int, now time.Time) (int, bool) {
	// Optimistically increment DeliveryAttempt assuming we'll be able to deliver the message.  This is
	// safe since the lock is held for the duration of this function, and the channel receiver does not
	// modify the message.
	if s.proto.DeadLetterPolicy != nil {
		m.proto.DeliveryAttempt = int32(*m.deliveries) + 1
	}

	for i := 0; i < len(s.streams); i++ {
		idx := (i + start) % len(s.streams)

		st := s.streams[idx]
		select {
		case <-st.done:
			s.streams = deleteStreamAt(s.streams, idx)
			i--

		case st.msgc <- m.proto:
			(*m.deliveries)++
			m.ackDeadline = now.Add(st.ackTimeout)
			return idx, true

		default:
		}
	}
	// Restore the correct value of DeliveryAttempt if we were not able to deliver the message.
	if s.proto.DeadLetterPolicy != nil {
		m.proto.DeliveryAttempt = int32(*m.deliveries)
	}
	return 0, false
}

const retentionDuration = 10 * time.Minute

// Must be called with the lock held.
func (s *subscription) maintainMessages(now time.Time) {
	for id, m := range s.msgs {
		// Mark a message as re-deliverable if its ack deadline has expired.
		if m.outstanding() && now.After(m.ackDeadline) {
			m.makeAvailable()
		}
		pubTime := m.proto.Message.PublishTime.AsTime()
		// Remove messages that have been undelivered for a long time.
		if !m.outstanding() && now.Sub(pubTime) > retentionDuration {
			delete(s.msgs, id)
		}
	}
}

func (s *subscription) newStream(gs pb.Subscriber_StreamingPullServer, timeout time.Duration) *stream {
	st := &stream{
		sub:                       s,
		done:                      make(chan struct{}),
		msgc:                      make(chan *pb.ReceivedMessage),
		gstream:                   gs,
		ackTimeout:                s.ackTimeout,
		timeout:                   timeout,
		enableExactlyOnceDelivery: s.proto.EnableExactlyOnceDelivery,
		enableOrdering:            s.proto.EnableMessageOrdering,
	}
	s.mu.Lock()
	s.streams = append(s.streams, st)
	s.mu.Unlock()
	return st
}

func (s *subscription) deleteStream(st *stream) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var i int
	for i = 0; i < len(s.streams); i++ {
		if s.streams[i] == st {
			break
		}
	}
	if i < len(s.streams) {
		s.streams = deleteStreamAt(s.streams, i)
	}
}

func (s *subscription) deadLetterCandidate(m *message) bool {
	if s.proto.DeadLetterPolicy == nil {
		return false
	}
	if m.retriesDone(s.proto.DeadLetterPolicy.MaxDeliveryAttempts) {
		return true
	}
	return false
}

func (s *subscription) publishToDeadLetter(m *message) {
	acks := 0
	if m.acks != nil {
		acks = *m.acks
	}
	deliveries := 0
	if m.deliveries != nil {
		deliveries = *m.deliveries
	}
	s.deadLetterTopic.publish(m.proto.Message, &Message{
		PublishTime: m.publishTime,
		Acks:        acks,
		Deliveries:  deliveries,
	})
}

func deleteStreamAt(s []*stream, i int) []*stream {
	// Preserve order for round-robin delivery.
	return append(s[:i], s[i+1:]...)
}

type message struct {
	proto       *pb.ReceivedMessage
	publishTime time.Time
	ackDeadline time.Time
	deliveries  *int
	acks        *int
	streamIndex int // index of stream that currently owns msg, for round-robin delivery
}

// A message is outstanding if it is owned by some stream.
func (m *message) outstanding() bool {
	return !m.ackDeadline.IsZero()
}

// A message is outstanding if it is owned by some stream.
func (m *message) retriesDone(maxRetries int32) bool {
	return m.deliveries != nil && int32(*m.deliveries) >= maxRetries
}

func (m *message) makeAvailable() {
	m.ackDeadline = time.Time{}
}

type stream struct {
	sub                       *subscription
	done                      chan struct{} // closed when the stream is finished
	msgc                      chan *pb.ReceivedMessage
	gstream                   pb.Subscriber_StreamingPullServer
	ackTimeout                time.Duration
	timeout                   time.Duration
	enableExactlyOnceDelivery bool
	enableOrdering            bool
}

// pull manages the StreamingPull interaction for the life of the stream.
func (st *stream) pull(wg *sync.WaitGroup) error {
	errc := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		errc <- st.sendLoop()
	}()
	go func() {
		defer wg.Done()
		errc <- st.recvLoop()
	}()
	var tchan <-chan time.Time
	if st.timeout > 0 {
		tchan = time.After(st.timeout)
	}
	// Wait until one of the goroutines returns an error, or we time out.
	var err error
	select {
	case err = <-errc:
		if errors.Is(err, io.EOF) {
			err = nil
		}
	case <-tchan:
	}
	close(st.done) // stop the other goroutine
	return err
}

func (st *stream) sendLoop() error {
	for {
		select {
		case <-st.done:
			return nil
		case rm := <-st.msgc:
			res := &pb.StreamingPullResponse{
				ReceivedMessages: []*pb.ReceivedMessage{rm},
				SubscriptionProperties: &pb.StreamingPullResponse_SubscriptionProperties{
					ExactlyOnceDeliveryEnabled: st.enableExactlyOnceDelivery,
					MessageOrderingEnabled:     st.enableOrdering,
				},
			}
			if err := st.gstream.Send(res); err != nil {
				return err
			}
		}
	}
}

func (st *stream) recvLoop() error {
	for {
		req, err := st.gstream.Recv()
		if err != nil {
			return err
		}
		st.sub.handleStreamingPullRequest(st, req)
	}
}

func (s *subscription) handleStreamingPullRequest(st *stream, req *pb.StreamingPullRequest) {
	// Lock the entire server.
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ackID := range req.AckIds {
		s.ack(ackID)
	}
	for i, id := range req.ModifyDeadlineAckIds {
		s.modifyAckDeadline(id, secsToDur(req.ModifyDeadlineSeconds[i]))
	}
	if req.StreamAckDeadlineSeconds > 0 {
		st.ackTimeout = secsToDur(req.StreamAckDeadlineSeconds)
	}
}

// Must be called with the lock held.
func (s *subscription) ack(id string) {
	m := s.msgs[id]
	if m != nil {
		(*m.acks)++
		delete(s.msgs, id)
	}
}

// Must be called with the lock held.
func (s *subscription) modifyAckDeadline(id string, d time.Duration) {
	m := s.msgs[id]
	if m == nil { // already acked: ignore.
		return
	}
	if d == 0 { // nack
		m.makeAvailable()
	} else { // extend the deadline by d
		m.ackDeadline = s.timeNowFunc().Add(d)
	}
}

func secsToDur(secs int32) time.Duration {
	return time.Duration(secs) * time.Second
}

// runReactor looks up the reactors for a function, then launches them until handled=true
// or err is returned. If the reactor returns nil, the function returns defaultObj instead.
func (s *GServer) runReactor(req interface{}, funcName string, defaultObj interface{}) (bool, interface{}, error) {
	if val, ok := s.reactorOptions[funcName]; ok {
		for _, reactor := range val {
			handled, ret, err := reactor.React(req)
			// If handled=true, that means the reactor has successfully reacted to the request,
			// so use the output directly. If err occurs, that means the request is invalidated
			// by the reactor somehow.
			if handled || err != nil {
				if ret == nil {
					ret = defaultObj
				}
				return true, ret, err
			}
		}
	}
	return false, nil, nil
}

// errorInjectionReactor is a reactor to inject an error message with status code.
type errorInjectionReactor struct {
	code codes.Code
	msg  string
}

// React simply returns an error with defined error message and status code.
func (e *errorInjectionReactor) React(_ interface{}) (handled bool, ret interface{}, err error) {
	return true, nil, status.Errorf(e.code, e.msg)
}

// WithErrorInjection creates a ServerReactorOption that injects error with defined status code and
// message for a certain function.
func WithErrorInjection(funcName string, code codes.Code, msg string) ServerReactorOption {
	return ServerReactorOption{
		FuncName: funcName,
		Reactor:  &errorInjectionReactor{code: code, msg: msg},
	}
}

const letters = "abcdef1234567890"

func genRevID() string {
	id := make([]byte, 8)
	for i := range id {
		id[i] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

// CreateSchema creates a new schema.
func (s *GServer) CreateSchema(_ context.Context, req *pb.CreateSchemaRequest) (*pb.Schema, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "CreateSchema", &pb.Schema{}); handled || err != nil {
		return ret.(*pb.Schema), err
	}

	name := fmt.Sprintf("%s/schemas/%s", req.Parent, req.SchemaId)
	sc := &pb.Schema{
		Name:               name,
		Type:               req.Schema.Type,
		Definition:         req.Schema.Definition,
		RevisionId:         genRevID(),
		RevisionCreateTime: timestamppb.Now(),
	}
	s.schemas[name] = append(s.schemas[name], sc)

	return sc, nil
}

// GetSchema gets an existing schema details.
func (s *GServer) GetSchema(_ context.Context, req *pb.GetSchemaRequest) (*pb.Schema, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "GetSchema", &pb.Schema{}); handled || err != nil {
		return ret.(*pb.Schema), err
	}

	ss := strings.Split(req.Name, "@")
	var schemaName, revisionID string
	if len := len(ss); len == 1 {
		schemaName = ss[0]
	} else if len == 2 {
		schemaName = ss[0]
		revisionID = ss[1]
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "schema(%q) name parse error", req.Name)
	}

	schemaRev, ok := s.schemas[schemaName]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "schema(%q) not found", req.Name)
	}

	if revisionID == "" {
		return schemaRev[len(schemaRev)-1], nil
	}

	for _, sc := range schemaRev {
		if sc.RevisionId == revisionID {
			return sc, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "schema %q not found", req.Name)
}

// ListSchemas lists the available schemas in this server.
func (s *GServer) ListSchemas(_ context.Context, req *pb.ListSchemasRequest) (*pb.ListSchemasResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ListSchemas", &pb.ListSchemasResponse{}); handled || err != nil {
		return ret.(*pb.ListSchemasResponse), err
	}
	ss := make([]*pb.Schema, 0)
	for _, sc := range s.schemas {
		ss = append(ss, sc[len(sc)-1])
	}
	return &pb.ListSchemasResponse{
		Schemas: ss,
	}, nil
}

// ListSchemaRevisions lists the schema revisions.
func (s *GServer) ListSchemaRevisions(_ context.Context, req *pb.ListSchemaRevisionsRequest) (*pb.ListSchemaRevisionsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ListSchemaRevisions", &pb.ListSchemasResponse{}); handled || err != nil {
		return ret.(*pb.ListSchemaRevisionsResponse), err
	}
	ss := make([]*pb.Schema, 0)
	ss = append(ss, s.schemas[req.Name]...)
	return &pb.ListSchemaRevisionsResponse{
		Schemas: ss,
	}, nil
}

// CommitSchema commits a new schema revision.
func (s *GServer) CommitSchema(_ context.Context, req *pb.CommitSchemaRequest) (*pb.Schema, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "CommitSchema", &pb.Schema{}); handled || err != nil {
		return ret.(*pb.Schema), err
	}

	sc := &pb.Schema{
		Name:       req.Name,
		Type:       req.Schema.Type,
		Definition: req.Schema.Definition,
	}
	sc.RevisionId = genRevID()
	sc.RevisionCreateTime = timestamppb.Now()

	s.schemas[req.Name] = append(s.schemas[req.Name], sc)

	return sc, nil
}

// RollbackSchema rolls back the current schema to a previous revision by copying and creating a new revision.
func (s *GServer) RollbackSchema(_ context.Context, req *pb.RollbackSchemaRequest) (*pb.Schema, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "RollbackSchema", &pb.Schema{}); handled || err != nil {
		return ret.(*pb.Schema), err
	}

	for _, sc := range s.schemas[req.Name] {
		if sc.RevisionId == req.RevisionId {
			cloned := proto.Clone(sc)
			newSchema := cloned.(*pb.Schema)
			newSchema.RevisionId = genRevID()
			newSchema.RevisionCreateTime = timestamppb.Now()
			s.schemas[req.Name] = append(s.schemas[req.Name], newSchema)
			return newSchema, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "schema %q@%q not found", req.Name, req.RevisionId)
}

// DeleteSchemaRevision deletes a schema revision.
func (s *GServer) DeleteSchemaRevision(_ context.Context, req *pb.DeleteSchemaRevisionRequest) (*pb.Schema, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "DeleteSchemaRevision", &pb.Schema{}); handled || err != nil {
		return ret.(*pb.Schema), err
	}

	schemaPath := strings.Split(req.Name, "@")
	if len(schemaPath) != 2 {
		return nil, status.Errorf(codes.InvalidArgument, "could not parse revision ID from schema name: %q", req.Name)
	}
	schemaName := schemaPath[0]
	revID := schemaPath[1]
	schemaRevisions, ok := s.schemas[schemaName]
	if ok {
		if len(schemaRevisions) == 1 {
			return nil, status.Errorf(codes.InvalidArgument, "cannot delete last revision for schema %q", req.Name)
		}
		for i, sc := range schemaRevisions {
			if sc.RevisionId == revID {
				s.schemas[schemaName] = append(schemaRevisions[:i], schemaRevisions[i+1:]...)
				return schemaRevisions[len(schemaRevisions)-1], nil
			}
		}
	}

	return nil, status.Errorf(codes.NotFound, "schema %q not found", req.Name)
}

// DeleteSchema deletes an existing schema.
func (s *GServer) DeleteSchema(_ context.Context, req *pb.DeleteSchemaRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "DeleteSchema", &emptypb.Empty{}); handled || err != nil {
		return ret.(*emptypb.Empty), err
	}

	schema := s.schemas[req.Name]
	if schema == nil {
		return nil, status.Errorf(codes.NotFound, "schema %q", req.Name)
	}

	delete(s.schemas, req.Name)
	return &emptypb.Empty{}, nil
}

// ValidateSchema mocks the ValidateSchema call but only checks that the schema definition is not empty.
func (s *GServer) ValidateSchema(_ context.Context, req *pb.ValidateSchemaRequest) (*pb.ValidateSchemaResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ValidateSchema", &pb.ValidateSchemaResponse{}); handled || err != nil {
		return ret.(*pb.ValidateSchemaResponse), err
	}

	if req.Schema.Definition == "" {
		return nil, status.Error(codes.InvalidArgument, "schema definition cannot be empty")
	}
	return &pb.ValidateSchemaResponse{}, nil
}

// ValidateMessage mocks the ValidateMessage call but only checks that the schema definition to validate the
// message against is not empty.
func (s *GServer) ValidateMessage(_ context.Context, req *pb.ValidateMessageRequest) (*pb.ValidateMessageResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if handled, ret, err := s.runReactor(req, "ValidateMessage", &pb.ValidateMessageResponse{}); handled || err != nil {
		return ret.(*pb.ValidateMessageResponse), err
	}

	spec := req.GetSchemaSpec()
	if valReq, ok := spec.(*pb.ValidateMessageRequest_Name); ok {
		sc, ok := s.schemas[valReq.Name]
		if !ok {
			return nil, status.Errorf(codes.NotFound, "schema(%q) not found", valReq.Name)
		}
		schema := sc[len(sc)-1]
		if schema.Definition == "" {
			return nil, status.Error(codes.InvalidArgument, "schema definition cannot be empty")
		}
	}
	if valReq, ok := spec.(*pb.ValidateMessageRequest_Schema); ok {
		if valReq.Schema.Definition == "" {
			return nil, status.Error(codes.InvalidArgument, "schema definition cannot be empty")
		}
	}

	return &pb.ValidateMessageResponse{}, nil
}
