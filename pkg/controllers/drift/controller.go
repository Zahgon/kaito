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

package drift

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/nodeprovision"
)

const (
	// driftActiveRequeueInterval is used only when drift is actively in progress
	// (budget "1") and we're waiting for Karpenter to complete node replacement.
	driftActiveRequeueInterval = 30 * time.Second
)

// DriftReconciler orchestrates rolling drift upgrades for InferenceSet-managed Workspaces.
type DriftReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	Recorder    record.EventRecorder
	Provisioner nodeprovision.NodeProvisioner
}

// NewDriftReconciler creates a DriftReconciler.
func NewDriftReconciler(c client.Client, scheme *runtime.Scheme, recorder record.EventRecorder, provisioner nodeprovision.NodeProvisioner) *DriftReconciler {
	_ = "STUB: not implemented"
	return nil
}

// getDriftBudgetNodes extracts the Drifted budget Nodes value from an already-fetched NodePool.
// Returns an error if no budget entry with DisruptionReasonDrifted is found.
func getDriftBudgetNodes(np *karpenterv1.NodePool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// hasDriftedNodeClaimsInGroup checks whether any NodeClaim in the slice has the Drifted condition.
func hasDriftedNodeClaimsInGroup(nodeClaims []*karpenterv1.NodeClaim) bool {
	_ = "STUB: not implemented"
	return false
}

// Reconcile implements the drift upgrade state machine for a single InferenceSet.
//
//  1. Get InferenceSet
//  2. List NodePools by InferenceSet labels
//  3. List NodeClaims by InferenceSet labels, group by NodePool name
//  4. For each NodePool, check drift budget and apply state machine
func (r *DriftReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	// 1. Get InferenceSet.
	return *new(ctrl.Result), nil
}

// 2. List NodePools for this InferenceSet.

// 3. List all NodeClaims for this InferenceSet and group by NodePool name.

// Group NodeClaims by their NodePool name label.

// 4. State machine: find upgrading NodePool or next candidate.

// Check if this NodePool has drifted NodeClaims.

// Case A: One NodePool is upgrading (budget "1").

// No drifted NodeClaims — check workspace readiness.

// Workload ready — disable drift remediation.

// Requeue to check if more NodePools need upgrading (no event will fire for
// already-drifted NodeClaims sitting stable in other pools).

// Case B: No NodePool is upgrading. Find next candidate.

// Enable drift remediation on the next candidate.

// isWorkspaceReady returns true if the workspace has WorkspaceSucceeded=True.
func isWorkspaceReady(ws *kaitov1beta1.Workspace) bool { _ = "STUB: not implemented"; return false }

// inferenceSetNodeClaimPredicate filters to only NodeClaims with both the
// InferenceSet name and namespace labels.
func inferenceSetNodeClaimPredicate() predicate.Predicate {
	_ = "STUB: not implemented"
	return *new(predicate.Predicate)
}

// mapNodeClaimToInferenceSet extracts InferenceSet name/namespace from
// NodeClaim labels and returns a reconcile request.
func mapNodeClaimToInferenceSet(_ context.Context, o client.Object) []reconcile.Request {
	_ = "STUB: not implemented"
	return nil
}

// enqueueInferenceSetForNodeClaim maps NodeClaim events to the owning InferenceSet.
var enqueueInferenceSetForNodeClaim = handler.EnqueueRequestsFromMapFunc(mapNodeClaimToInferenceSet)

// inferenceSetWorkspacePredicate filters to only Workspaces created by an InferenceSet.
func inferenceSetWorkspacePredicate() predicate.Predicate {
	_ = "STUB: not implemented"
	return *new(predicate.Predicate)
}

// mapWorkspaceToInferenceSet extracts InferenceSet name from the workspace's
// created-by label and returns a reconcile request using the workspace's namespace.
func mapWorkspaceToInferenceSet(_ context.Context, o client.Object) []reconcile.Request {
	_ = "STUB: not implemented"
	return nil
}

// enqueueInferenceSetForWorkspace maps Workspace events to the owning InferenceSet.
var enqueueInferenceSetForWorkspace = handler.EnqueueRequestsFromMapFunc(mapWorkspaceToInferenceSet)

// SetupWithManager registers the controller with the manager.
func (r *DriftReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
