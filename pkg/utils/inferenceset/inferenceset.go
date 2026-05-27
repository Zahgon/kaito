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

package inferenceset

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1alpha1 "github.com/kaito-project/kaito/api/v1alpha1"
	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

// UpdateStatusConditionIfNotMatch updates the inferenceset status condition if it doesn't match the current values
func UpdateStatusConditionIfNotMatch(ctx context.Context, c client.Client, iObj *kaitov1alpha1.InferenceSet, cType kaitov1alpha1.ConditionType,
	cStatus metav1.ConditionStatus, cReason, cMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to change

// UpdateInferenceSetStatus updates the inferenceset status with the provided condition
func UpdateInferenceSetStatus(ctx context.Context, c client.Client, name *client.ObjectKey, modifyFn func(*kaitov1alpha1.InferenceSetStatus) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the latest version to avoid update conflict.

// UpdateInferenceSetWithRetry gets the latest inferenceset object, applies the modify function, and retries on conflict
func UpdateInferenceSetWithRetry(ctx context.Context, c client.Client, iObj *kaitov1alpha1.InferenceSet, modifyFn func(*kaitov1alpha1.InferenceSet) error) error {
	_ = "STUB: not implemented"
	return nil
}

// ListWorkspaces lists all workspace objects in the InferenceSet's namespace that are created by the given InferenceSet.
func ListWorkspaces(ctx context.Context, iObj *kaitov1alpha1.InferenceSet, kubeClient client.Client) (*kaitov1beta1.WorkspaceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all the workspaces in iObj.Namespace that are created by this inferenceset.
// We use label selector to find the workspaces.
// The label is "inferenceset.kaito.sh/created-by": <inferenceset-name>

func ComputeInferenceSetHash(iObj *kaitov1alpha1.InferenceSet) string {
	_ = "STUB: not implemented"
	return ""
}

func MarshalInferenceSetFields(iObj *kaitov1alpha1.InferenceSet) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
