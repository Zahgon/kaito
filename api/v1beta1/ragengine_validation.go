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

package v1beta1

import (
	"context"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
)

func (w *RAGEngine) SupportedVerbs() []admissionregistrationv1.OperationType {
	_ = "STUB: not implemented"
	return nil
}

func (w *RAGEngine) Validate(ctx context.Context) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (w *RAGEngine) validateCreate() (errs *apis.FieldError) { _ = "STUB: not implemented"; return nil }

func (w *RAGEngine) validateGuardrails(ctx context.Context) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func validateGuardrailsPolicyConfigMap(cm *corev1.ConfigMap) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (w *RAGEngine) validateUpdate(old *RAGEngine) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResourceSpec) validateRAGCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Check for other instance types pattern matches if cloud provider is Azure

// Validate labelSelector

func (e *LocalEmbeddingSpec) validateCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Executes if image is of correct format

func (e *RemoteEmbeddingSpec) validateCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (e *InferenceServiceSpec) validateCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Only validate URL if it's provided
	return nil
}
