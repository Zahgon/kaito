// Copyright (c) KAITO authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workspace

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	estimatorpkg "github.com/kaito-project/kaito/pkg/workspace/estimator"
)

// NodeEstimateRequestFromWorkspace builds a NodeEstimateRequest from a Workspace object.
// It fetches the HuggingFace access token from Kubernetes when ModelAccessSecret is set,
// so the returned request carries the resolved token value rather than a secret reference.
// This is a convenience helper for callers that already have a Workspace and a kube client.
func NodeEstimateRequestFromWorkspace(ctx context.Context, w *kaitov1beta1.Workspace, kubeClient client.Client) (estimatorpkg.NodeEstimateRequest, error) {
	_ = "STUB: not implemented"
	return *new(estimatorpkg.NodeEstimateRequest), nil
}

//nolint:staticcheck //SA1019: deprecate Resource.Count field

//nolint:staticcheck //SA1019: deprecate Resource.Count field

// Only HuggingFace models (names containing "/") require an access token for model lookup.
// Standard preset models use ModelAccessSecret only for runtime env-var injection,
// which is outside the estimator's concern.
