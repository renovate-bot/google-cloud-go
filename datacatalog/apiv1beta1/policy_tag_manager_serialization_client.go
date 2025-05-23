// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package datacatalog

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	datacatalogpb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newPolicyTagManagerSerializationClientHook clientHook

// PolicyTagManagerSerializationCallOptions contains the retry settings for each method of PolicyTagManagerSerializationClient.
type PolicyTagManagerSerializationCallOptions struct {
	ImportTaxonomies []gax.CallOption
	ExportTaxonomies []gax.CallOption
}

func defaultPolicyTagManagerSerializationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("datacatalog.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("datacatalog.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("datacatalog.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://datacatalog.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPolicyTagManagerSerializationCallOptions() *PolicyTagManagerSerializationCallOptions {
	return &PolicyTagManagerSerializationCallOptions{
		ImportTaxonomies: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ExportTaxonomies: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultPolicyTagManagerSerializationRESTCallOptions() *PolicyTagManagerSerializationCallOptions {
	return &PolicyTagManagerSerializationCallOptions{
		ImportTaxonomies: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ExportTaxonomies: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalPolicyTagManagerSerializationClient is an interface that defines the methods available from Google Cloud Data Catalog API.
type internalPolicyTagManagerSerializationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ImportTaxonomies(context.Context, *datacatalogpb.ImportTaxonomiesRequest, ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error)
	ExportTaxonomies(context.Context, *datacatalogpb.ExportTaxonomiesRequest, ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error)
}

// PolicyTagManagerSerializationClient is a client for interacting with Google Cloud Data Catalog API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Policy tag manager serialization API service allows clients to manipulate
// their taxonomies and policy tags data with serialized format.
type PolicyTagManagerSerializationClient struct {
	// The internal transport-dependent client.
	internalClient internalPolicyTagManagerSerializationClient

	// The call options for this service.
	CallOptions *PolicyTagManagerSerializationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PolicyTagManagerSerializationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PolicyTagManagerSerializationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PolicyTagManagerSerializationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ImportTaxonomies imports all taxonomies and their policy tags to a project as new
// taxonomies.
//
// This method provides a bulk taxonomy / policy tag creation using nested
// proto structure.
func (c *PolicyTagManagerSerializationClient) ImportTaxonomies(ctx context.Context, req *datacatalogpb.ImportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error) {
	return c.internalClient.ImportTaxonomies(ctx, req, opts...)
}

// ExportTaxonomies exports all taxonomies and their policy tags in a project.
//
// This method generates SerializedTaxonomy protos with nested policy tags
// that can be used as an input for future ImportTaxonomies calls.
func (c *PolicyTagManagerSerializationClient) ExportTaxonomies(ctx context.Context, req *datacatalogpb.ExportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error) {
	return c.internalClient.ExportTaxonomies(ctx, req, opts...)
}

// policyTagManagerSerializationGRPCClient is a client for interacting with Google Cloud Data Catalog API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type policyTagManagerSerializationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PolicyTagManagerSerializationClient
	CallOptions **PolicyTagManagerSerializationCallOptions

	// The gRPC API client.
	policyTagManagerSerializationClient datacatalogpb.PolicyTagManagerSerializationClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewPolicyTagManagerSerializationClient creates a new policy tag manager serialization client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Policy tag manager serialization API service allows clients to manipulate
// their taxonomies and policy tags data with serialized format.
func NewPolicyTagManagerSerializationClient(ctx context.Context, opts ...option.ClientOption) (*PolicyTagManagerSerializationClient, error) {
	clientOpts := defaultPolicyTagManagerSerializationGRPCClientOptions()
	if newPolicyTagManagerSerializationClientHook != nil {
		hookOpts, err := newPolicyTagManagerSerializationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PolicyTagManagerSerializationClient{CallOptions: defaultPolicyTagManagerSerializationCallOptions()}

	c := &policyTagManagerSerializationGRPCClient{
		connPool:                            connPool,
		policyTagManagerSerializationClient: datacatalogpb.NewPolicyTagManagerSerializationClient(connPool),
		CallOptions:                         &client.CallOptions,
		logger:                              internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *policyTagManagerSerializationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *policyTagManagerSerializationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *policyTagManagerSerializationGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type policyTagManagerSerializationRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing PolicyTagManagerSerializationClient
	CallOptions **PolicyTagManagerSerializationCallOptions

	logger *slog.Logger
}

// NewPolicyTagManagerSerializationRESTClient creates a new policy tag manager serialization rest client.
//
// Policy tag manager serialization API service allows clients to manipulate
// their taxonomies and policy tags data with serialized format.
func NewPolicyTagManagerSerializationRESTClient(ctx context.Context, opts ...option.ClientOption) (*PolicyTagManagerSerializationClient, error) {
	clientOpts := append(defaultPolicyTagManagerSerializationRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPolicyTagManagerSerializationRESTCallOptions()
	c := &policyTagManagerSerializationRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &PolicyTagManagerSerializationClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPolicyTagManagerSerializationRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://datacatalog.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://datacatalog.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://datacatalog.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://datacatalog.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *policyTagManagerSerializationRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *policyTagManagerSerializationRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *policyTagManagerSerializationRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *policyTagManagerSerializationGRPCClient) ImportTaxonomies(ctx context.Context, req *datacatalogpb.ImportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ImportTaxonomies[0:len((*c.CallOptions).ImportTaxonomies):len((*c.CallOptions).ImportTaxonomies)], opts...)
	var resp *datacatalogpb.ImportTaxonomiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.policyTagManagerSerializationClient.ImportTaxonomies, req, settings.GRPC, c.logger, "ImportTaxonomies")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyTagManagerSerializationGRPCClient) ExportTaxonomies(ctx context.Context, req *datacatalogpb.ExportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ExportTaxonomies[0:len((*c.CallOptions).ExportTaxonomies):len((*c.CallOptions).ExportTaxonomies)], opts...)
	var resp *datacatalogpb.ExportTaxonomiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.policyTagManagerSerializationClient.ExportTaxonomies, req, settings.GRPC, c.logger, "ExportTaxonomies")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ImportTaxonomies imports all taxonomies and their policy tags to a project as new
// taxonomies.
//
// This method provides a bulk taxonomy / policy tag creation using nested
// proto structure.
func (c *policyTagManagerSerializationRESTClient) ImportTaxonomies(ctx context.Context, req *datacatalogpb.ImportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ImportTaxonomiesResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/taxonomies:import", req.GetParent())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ImportTaxonomies[0:len((*c.CallOptions).ImportTaxonomies):len((*c.CallOptions).ImportTaxonomies)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &datacatalogpb.ImportTaxonomiesResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "ImportTaxonomies")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ExportTaxonomies exports all taxonomies and their policy tags in a project.
//
// This method generates SerializedTaxonomy protos with nested policy tags
// that can be used as an input for future ImportTaxonomies calls.
func (c *policyTagManagerSerializationRESTClient) ExportTaxonomies(ctx context.Context, req *datacatalogpb.ExportTaxonomiesRequest, opts ...gax.CallOption) (*datacatalogpb.ExportTaxonomiesResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/taxonomies:export", req.GetParent())

	params := url.Values{}
	if req.GetSerializedTaxonomies() {
		params.Add("serializedTaxonomies", fmt.Sprintf("%v", req.GetSerializedTaxonomies()))
	}
	if items := req.GetTaxonomies(); len(items) > 0 {
		for _, item := range items {
			params.Add("taxonomies", fmt.Sprintf("%v", item))
		}
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ExportTaxonomies[0:len((*c.CallOptions).ExportTaxonomies):len((*c.CallOptions).ExportTaxonomies)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &datacatalogpb.ExportTaxonomiesResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ExportTaxonomies")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
