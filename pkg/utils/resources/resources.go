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

package resources

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateResource(ctx context.Context, resource client.Object, kubeClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the resource.

func GetResource(ctx context.Context, name, namespace string, kubeClient client.Client, resource client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckResourceStatus(obj client.Object, kubeClient client.Client, timeoutDuration time.Duration) error {
	_ = "STUB: not implemented"
	// Use Context for timeout
	return nil
}

// EnsureConfigOrCopyFromDefault handles two scenarios:
// 1. User provided config:
//   - Check if it exists in the target namespace
//   - If not found, return error as this is user-specified
//
// 2. No user config specified:
//   - Use the default config template
//   - Check if it exists in the target namespace
//   - If not, copy from release namespace to target namespace
func EnsureConfigOrCopyFromDefault(ctx context.Context, kubeClient client.Client,
	userProvided, systemDefault client.ObjectKey,
) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"

	// If user specified a config, use that
	return nil, nil
}

// Check if default configmap already exists in target namespace

// Copy default template from release namespace if not found

// Clear metadata not needed for creation
// Clear UID
