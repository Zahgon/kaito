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

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

// garbageCollectRAGEngine remove finalizer associated with ragengine object.
func (c *RAGEngineReconciler) garbageCollectRAGEngine(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Only clean up NodeClaims when node auto-provisioning is enabled,
// since NodeClaim CRDs may not be installed when it's disabled.

// Check if there are any nodeClaims associated with this ragengine.

// We should delete all the nodeClaims that are created by this ragengine
