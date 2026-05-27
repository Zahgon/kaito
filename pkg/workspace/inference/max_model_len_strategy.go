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

package inference

import (
	pkgmodel "github.com/kaito-project/kaito/pkg/model"
	"github.com/kaito-project/kaito/pkg/sku"
)

// computeMaxModelLen calculates the optimal max model length for GPU memory efficiency.
func computeMaxModelLen(preset *pkgmodel.PresetParam, gpu *sku.GPUConfig, numRequiredNodes int) int {
	_ = "STUB: not implemented"
	// Validate input parameters
	return 0
}

// Parse model weight size using Kubernetes resource.Quantity

// Calculate available GPU memory and adjusted bytes per token

// Calculate raw token candidate

// Apply constraints and finalize result

// parseModelWeight parses the model weight string and returns the weight in GiB.
// It handles Kubernetes resource.Quantity format strings like "25.63Gi", etc.
func parseModelWeight(totalSafeTensorFileSize string) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Use Kubernetes resource.Quantity to parse the size string

// Convert bytes to GiB (1 GiB = 2^30 bytes)

// calculateMemoryParameters computes available GPU memory and adjusted bytes per token.
// Returns the available memory in bytes and the adjusted bytes per token for the calculation.
func calculateMemoryParameters(preset *pkgmodel.PresetParam, gpu *sku.GPUConfig, numRequiredNodes int, weightGiB float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Calculate available GPU memory using the formula:
// availableMemoryGiB = (gpuMemGB * 0.84) / gpuCount - (weightGiB * 1.02) / (nodes * gpuCount) - 2.3

// Available GPU memory per GPU with 84% utilization factor

// Model weight overhead per GPU with 2% safety margin

// Falcon models: don't divide by nodes and gpuCount

// Other models: distribute weight across nodes and GPUs

// Static overhead for activations and non-torch components
// Sum of max activations (1.7 GiB) and max non-torch overhead (0.6 GiB)

// Calculate adjusted bytes per token for distribution

// Falcon models: no distribution adjustment

// Other models: distribute across GPUs

// applyConstraintsAndAlignment applies token limit constraints and 256-token boundary alignment.
// This ensures the result stays within model limits and is optimally aligned for efficiency.
func applyConstraintsAndAlignment(candidate, modelTokenLimit int) int {
	_ = "STUB: not implemented"
	// Clamp to model's token limit if necessary
	return 0
}

// Align down to 256-token boundary for efficiency
// This helps with memory allocation patterns and performance optimization
