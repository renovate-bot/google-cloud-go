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

package dataflow

import (
	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	"google.golang.org/api/iterator"
)

// JobIterator manages a stream of *dataflowpb.Job.
type JobIterator struct {
	items    []*dataflowpb.Job
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*dataflowpb.Job, error) {
	var item *dataflowpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// JobMessageIterator manages a stream of *dataflowpb.JobMessage.
type JobMessageIterator struct {
	items    []*dataflowpb.JobMessage
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.JobMessage, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *JobMessageIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobMessageIterator) Next() (*dataflowpb.JobMessage, error) {
	var item *dataflowpb.JobMessage
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobMessageIterator) bufLen() int {
	return len(it.items)
}

func (it *JobMessageIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// StageSummaryIterator manages a stream of *dataflowpb.StageSummary.
type StageSummaryIterator struct {
	items    []*dataflowpb.StageSummary
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.StageSummary, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *StageSummaryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StageSummaryIterator) Next() (*dataflowpb.StageSummary, error) {
	var item *dataflowpb.StageSummary
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StageSummaryIterator) bufLen() int {
	return len(it.items)
}

func (it *StageSummaryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkerDetailsIterator manages a stream of *dataflowpb.WorkerDetails.
type WorkerDetailsIterator struct {
	items    []*dataflowpb.WorkerDetails
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.WorkerDetails, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *WorkerDetailsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkerDetailsIterator) Next() (*dataflowpb.WorkerDetails, error) {
	var item *dataflowpb.WorkerDetails
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkerDetailsIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkerDetailsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
