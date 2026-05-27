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

package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this RAGEngine (v1alpha1) to the Hub version (v1beta1).
func (src *RAGEngine) ConvertTo(dstRaw conversion.Hub) error { _ = "STUB: not implemented"; return nil }

// Copy TypeMeta

// Copy ObjectMeta

// Convert Spec

// Convert Compute

// Convert Storage: v1alpha1 flat -> v1beta1 nested

// Convert Embedding

// Convert InferenceService

// Note: QueryServiceName and IndexServiceName are v1alpha1-only fields, not converted

// Convert Status

// ConvertFrom converts from the Hub version (v1beta1) to this version (v1alpha1).
func (dst *RAGEngine) ConvertFrom(srcRaw conversion.Hub) error {
	_ = "STUB: not implemented"
	return nil
}

// Copy TypeMeta

// Copy ObjectMeta

// Convert Spec

// Convert Compute

// Convert Storage: v1beta1 nested -> v1alpha1 flat

// Convert Embedding

// Convert InferenceService

// QueryServiceName and IndexServiceName are v1alpha1-only fields, left empty

// Convert Status
