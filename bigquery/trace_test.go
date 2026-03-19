// Copyright 2026 Google LLC
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

package bigquery

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/googleapis/gax-go/v2"
	"github.com/googleapis/gax-go/v2/callctx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"google.golang.org/api/option"
)

func TestSetDatasetTraceMetadata(t *testing.T) {
	// Enable tracing feature for the test
	os.Setenv("GOOGLE_SDK_GO_EXPERIMENTAL_TRACING", "true")
	defer os.Unsetenv("GOOGLE_SDK_GO_EXPERIMENTAL_TRACING")
	gax.TestOnlyResetIsFeatureEnabled()
	defer gax.TestOnlyResetIsFeatureEnabled()

	ctx := context.Background()
	projectID := "test-project"
	datasetID := "test-dataset"
	resourceName := "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset"
	urlTemplate := "/bigquery/v2/projects/{projectId}/datasets/{datasetId}"

	ctx = setDatasetTraceMetadata(ctx, projectID, datasetID)

	res, ok := callctx.TelemetryFromContext(ctx, "resource_name")
	if !ok || res != resourceName {
		t.Errorf("expected resource_name %q, got %q", resourceName, res)
	}

	urlTmpl, ok := callctx.TelemetryFromContext(ctx, "url_template")
	if !ok || urlTmpl != urlTemplate {
		t.Errorf("expected url_template %q, got %q", urlTemplate, urlTmpl)
	}
}

func TestTracingTelemetryAttributes(t *testing.T) {
	os.Setenv("GOOGLE_SDK_GO_EXPERIMENTAL_TRACING", "true")
	defer os.Unsetenv("GOOGLE_SDK_GO_EXPERIMENTAL_TRACING")
	gax.TestOnlyResetIsFeatureEnabled()
	defer gax.TestOnlyResetIsFeatureEnabled()

	tests := []struct {
		name             string
		callFunc         func(ctx context.Context, client *Client)
		mockResponse     string
		mockStatusCodes  []int
		wantResourceName string
		wantURLTemplate  string
		wantAttempts     int
	}{
		{
			name: "Dataset_Metadata",
			callFunc: func(ctx context.Context, client *Client) {
				_, _ = client.Dataset("test-dataset").Metadata(ctx)
			},
			mockResponse:     `{"id": "test-dataset", "datasetReference": {"projectId": "test-project", "datasetId": "test-dataset"}}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}",
			wantAttempts:     1,
		},
		{
			name: "Dataset_Create",
			callFunc: func(ctx context.Context, client *Client) {
				_ = client.Dataset("test-dataset").Create(ctx, &DatasetMetadata{})
			},
			mockResponse:     `{"id": "test-dataset", "datasetReference": {"projectId": "test-project", "datasetId": "test-dataset"}}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets",
			wantAttempts:     1,
		},
		{
			name: "Dataset_Update",
			callFunc: func(ctx context.Context, client *Client) {
				_, _ = client.Dataset("test-dataset").Update(ctx, DatasetMetadataToUpdate{}, "")
			},
			mockResponse:     `{"id": "test-dataset", "datasetReference": {"projectId": "test-project", "datasetId": "test-dataset"}}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}",
			wantAttempts:     1,
		},
		{
			name: "Dataset_Delete",
			callFunc: func(ctx context.Context, client *Client) {
				_ = client.Dataset("test-dataset").Delete(ctx)
			},
			mockResponse:     `{}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}",
			wantAttempts:     1,
		},
		{
			name: "Client_Datasets",
			callFunc: func(ctx context.Context, client *Client) {
				it := client.Datasets(ctx)
				_, _ = it.Next()
			},
			mockResponse:     `{"datasets": [{"datasetReference": {"projectId": "test-project", "datasetId": "test-dataset"}}]}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets",
			wantAttempts:     1,
		},
		{
			name: "Dataset_Models",
			callFunc: func(ctx context.Context, client *Client) {
				it := client.Dataset("test-dataset").Models(ctx)
				_, _ = it.Next()
			},
			mockResponse:     `{"models": [{"modelReference": {"projectId": "test-project", "datasetId": "test-dataset", "modelId": "test-model"}}]}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}/models",
			wantAttempts:     1,
		},
		{
			name: "Model_Metadata",
			callFunc: func(ctx context.Context, client *Client) {
				_, _ = client.Dataset("test-dataset").Model("test-model").Metadata(ctx)
			},
			mockResponse:     `{"modelReference": {"projectId": "test-project", "datasetId": "test-dataset", "modelId": "test-model"}}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset/models/test-model",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}/models/{modelId}",
			wantAttempts:     1,
		},
		{
			name: "Model_Update",
			callFunc: func(ctx context.Context, client *Client) {
				_, _ = client.Dataset("test-dataset").Model("test-model").Update(ctx, ModelMetadataToUpdate{}, "")
			},
			mockResponse:     `{"modelReference": {"projectId": "test-project", "datasetId": "test-dataset", "modelId": "test-model"}}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset/models/test-model",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}/models/{modelId}",
			wantAttempts:     1,
		},
		{
			name: "Model_Delete",
			callFunc: func(ctx context.Context, client *Client) {
				_ = client.Dataset("test-dataset").Model("test-model").Delete(ctx)
			},
			mockResponse:     `{}`,
			mockStatusCodes:  []int{http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset/models/test-model",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}/models/{modelId}",
			wantAttempts:     1,
		},
		{
			name: "Retry_Dataset_Metadata",
			callFunc: func(ctx context.Context, client *Client) {
				_, _ = client.Dataset("test-dataset").Metadata(ctx)
			},
			mockResponse:     `{"id": "test-dataset", "datasetReference": {"projectId": "test-project", "datasetId": "test-dataset"}}`,
			mockStatusCodes:  []int{http.StatusServiceUnavailable, http.StatusOK},
			wantResourceName: "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset",
			wantURLTemplate:  "/bigquery/v2/projects/{projectId}/datasets/{datasetId}",
			wantAttempts:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp := tracetest.NewInMemoryExporter()
			tp := sdktrace.NewTracerProvider(
				sdktrace.WithSyncer(exp),
			)
			otel.SetTracerProvider(tp)

			attempts := 0
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				status := http.StatusOK
				if attempts < len(tt.mockStatusCodes) {
					status = tt.mockStatusCodes[attempts]
				}
				attempts++
				w.WriteHeader(status)
				if status == http.StatusOK {
					w.Write([]byte(tt.mockResponse))
				}
			}))
			defer ts.Close()

			ctx := context.Background()
			client, err := NewClient(ctx, "test-project", option.WithEndpoint(ts.URL), option.WithoutAuthentication())
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}
			defer client.Close()

			tt.callFunc(ctx, client)

			spans := exp.GetSpans()
			if len(spans) == 0 {
				t.Fatalf("expected spans to be recorded, got 0")
			}

			if attempts != tt.wantAttempts {
				t.Errorf("expected %d attempts, got %d", tt.wantAttempts, attempts)
			}

			networkSpans := 0
			for _, span := range spans {
				// The otelAttributeTransport renames the span to "{METHOD} {url.template}" but might have a duplicated method name
				if strings.Contains(span.Name, tt.wantURLTemplate) {
					networkSpans++

					foundRes := false
					foundURL := false
					foundArtifact := false
					foundLanguage := false
					foundDomain := false

					for _, attr := range span.Attributes {
						if attr.Key == attribute.Key("gcp.resource.destination.id") && attr.Value.AsString() == tt.wantResourceName {
							foundRes = true
						}
						if attr.Key == attribute.Key("url.template") && attr.Value.AsString() == tt.wantURLTemplate {
							foundURL = true
						}
						if attr.Key == attribute.Key("gcp.client.artifact") && attr.Value.AsString() == "cloud.google.com/go/bigquery" {
							foundArtifact = true
						}
						if attr.Key == attribute.Key("gcp.client.language") && attr.Value.AsString() == "go" {
							foundLanguage = true
						}
						if attr.Key == attribute.Key("url.domain") && attr.Value.AsString() == "bigquery.googleapis.com" {
							foundDomain = true
						}
					}

					if !foundRes {
						t.Errorf("missing gcp.resource.destination.id attribute on network span attempt")
					}
					if !foundURL {
						t.Errorf("missing url.template attribute on network span attempt")
					}
					if !foundArtifact {
						t.Errorf("missing gcp.client.artifact attribute on network span attempt")
					}
					if !foundLanguage {
						t.Errorf("missing gcp.client.language attribute on network span attempt")
					}
					if !foundDomain {
						t.Errorf("missing url.domain attribute on network span attempt")
					}
				}
			}

			if networkSpans != tt.wantAttempts {
				var names []string
				for _, s := range spans {
					names = append(names, s.Name)
				}
				t.Errorf("expected %d network spans, got %d. Found span names: %v", tt.wantAttempts, networkSpans, names)
			}
		})
	}
}
