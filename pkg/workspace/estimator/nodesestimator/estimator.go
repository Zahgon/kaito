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

package nodesestimator

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	estimator "github.com/kaito-project/kaito/pkg/workspace/estimator"
)

// NodeEstimator estimates node count based on SKU memory and model memory requirement
type NodeEstimator struct {
	// no fields needed
}

func (c *NodeEstimator) Name() string { _ = "STUB: not implemented"; return "" }

func (c *NodeEstimator) EstimateNodeCount(ctx context.Context, req estimator.NodeEstimateRequest, cl client.Client) (int32, error) {
	_ = "STUB: not implemented"
	// If no preset is configured, default to the requested node count or 1.
	return 0, nil
}

// NAP is disabled (BYO scenario) — derive GPU config from existing ready nodes.

// NAP is enabled — instanceType is required and must be valid.

// Start with the user-requested node count (default is 1).

// maxModelLen: use the value resolved by the caller (RuntimeProfile.ContextSize), falling back to 2048.

// If GPU memory information is available, calculate the optimal node count

// vllm model size is about 102% of HuggingFace size

// utilization is set to default 0.84

// Overhead: fixed base (2.3GB) + KV cache for context length
