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
	"context"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	"knative.dev/pkg/apis"
)

func (m *MultiRoleInference) SupportedVerbs() []admissionregistrationv1.OperationType {
	_ = "STUB: not implemented"
	return nil
}

func (m *MultiRoleInference) Validate(ctx context.Context) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Validate name is a valid DNS label.
	return nil
}

func (m *MultiRoleInference) validateCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Validate model name is not empty.
	return nil
}

// Validate labelSelector is not nil and not empty.

// Validate roles.

func (m *MultiRoleInference) validateUpdate(old *MultiRoleInference) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Model name is immutable.
	return nil
}

// Validate roles (same as create).

func (m *MultiRoleInference) validateRoles() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Validate exactly 2 roles.
	return nil
}

// Validate role type.

// Validate instanceType is not empty.

// Validate replicas >= 1 when specified (nil means autoscaling).
