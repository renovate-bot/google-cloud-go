// Copyright 2024 Google LLC
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

package lfp_test

import (
	"context"

	lfp "cloud.google.com/go/shopping/merchant/lfp/apiv1beta"
	lfppb "cloud.google.com/go/shopping/merchant/lfp/apiv1beta/lfppb"
	"google.golang.org/api/iterator"
)

func ExampleNewLfpStoreClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := lfp.NewLfpStoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewLfpStoreRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := lfp.NewLfpStoreRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleLfpStoreClient_DeleteLfpStore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := lfp.NewLfpStoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &lfppb.DeleteLfpStoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/lfp/apiv1beta/lfppb#DeleteLfpStoreRequest.
	}
	err = c.DeleteLfpStore(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleLfpStoreClient_GetLfpStore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := lfp.NewLfpStoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &lfppb.GetLfpStoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/lfp/apiv1beta/lfppb#GetLfpStoreRequest.
	}
	resp, err := c.GetLfpStore(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleLfpStoreClient_InsertLfpStore() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := lfp.NewLfpStoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &lfppb.InsertLfpStoreRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/lfp/apiv1beta/lfppb#InsertLfpStoreRequest.
	}
	resp, err := c.InsertLfpStore(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleLfpStoreClient_ListLfpStores() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := lfp.NewLfpStoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &lfppb.ListLfpStoresRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/lfp/apiv1beta/lfppb#ListLfpStoresRequest.
	}
	it := c.ListLfpStores(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*lfppb.ListLfpStoresResponse)
	}
}