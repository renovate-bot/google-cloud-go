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
	"fmt"
	"strings"

	"github.com/googleapis/gax-go/v2"
	"github.com/googleapis/gax-go/v2/callctx"
)

// setTraceMetadata sets the trace metadata into the context. It assumes the caller has already checked if tracing is enabled.
func setTraceMetadata(ctx context.Context, resourceName, urlTemplate string) context.Context {
	return callctx.WithTelemetryContext(ctx, "resource_name", resourceName, "url_template", urlTemplate)
}

// datasetResourceName constructs the standard resource name for a dataset.
// E.g., "//bigquery.googleapis.com/projects/{project}/datasets/{dataset}"
func datasetResourceName(projectID, datasetID string) string {
	return fmt.Sprintf("//bigquery.googleapis.com/projects/%s/datasets/%s", projectID, datasetID)
}

// modelResourceName constructs the standard resource name for a model.
// E.g., "//bigquery.googleapis.com/projects/{project}/datasets/{dataset}/models/{model}"
func modelResourceName(projectID, datasetID, modelID string) string {
	return fmt.Sprintf("//bigquery.googleapis.com/projects/%s/datasets/%s/models/%s", projectID, datasetID, modelID)
}

func fullyQualifiedDatasetResourceName(projectID, datasetID string) string {
	if strings.HasPrefix(datasetID, "projects/") {
		// Handle fully qualified names
		return fmt.Sprintf("//bigquery.googleapis.com/%s", datasetID)
	}
	return datasetResourceName(projectID, datasetID)
}

func setProjectItemTraceMetadata(ctx context.Context, projectID, childType string) context.Context {
	if !gax.IsFeatureEnabled("TRACING") {
		return ctx
	}
	return setTraceMetadata(ctx,
		fmt.Sprintf("//bigquery.googleapis.com/projects/%s", projectID),
		fmt.Sprintf("/bigquery/v2/projects/{projectId}/%s", childType))
}

func setDatasetTraceMetadata(ctx context.Context, projectID, datasetID string) context.Context {
	if !gax.IsFeatureEnabled("TRACING") {
		return ctx
	}
	return setTraceMetadata(ctx,
		fullyQualifiedDatasetResourceName(projectID, datasetID),
		"/bigquery/v2/projects/{projectId}/datasets/{datasetId}")
}

func setDatasetItemTraceMetadata(ctx context.Context, projectID, datasetID, childType string) context.Context {
	if !gax.IsFeatureEnabled("TRACING") {
		return ctx
	}
	return setTraceMetadata(ctx,
		fullyQualifiedDatasetResourceName(projectID, datasetID),
		fmt.Sprintf("/bigquery/v2/projects/{projectId}/datasets/{datasetId}/%s", childType))
}

func setModelTraceMetadata(ctx context.Context, projectID, datasetID, modelID string) context.Context {
	if !gax.IsFeatureEnabled("TRACING") {
		return ctx
	}
	return setTraceMetadata(ctx,
		modelResourceName(projectID, datasetID, modelID),
		"/bigquery/v2/projects/{projectId}/datasets/{datasetId}/models/{modelId}")
}
