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

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/nodeprovision"
	"github.com/kaito-project/kaito/pkg/utils"
	"github.com/kaito-project/kaito/pkg/workspace/estimator"
)

const (
	WorkspaceHashAnnotation = "workspace.kaito.io/hash"
	WorkspaceNameLabel      = "workspace.kaito.io/name"
	revisionHashSuffix      = 5
)

type WorkspaceReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder

	klogger         klog.Logger
	expectations    *utils.ControllerExpectations
	Estimator       estimator.NodesEstimator
	nodeProvisioner nodeprovision.NodeProvisioner
}

func NewWorkspaceReconciler(client client.Client, scheme *runtime.Scheme, log logr.Logger, Recorder record.EventRecorder, provisioner nodeprovision.NodeProvisioner) *WorkspaceReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (c *WorkspaceReconciler) Reconcile(ctx context.Context, req reconcile.Request) (result reconcile.Result, err error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Handle deleting workspace, garbage collect all the resources.

// update targetNodeCount for the workspace

func (c *WorkspaceReconciler) ensureFinalizer(ctx context.Context, workspaceObj *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *WorkspaceReconciler) reconcileNodes(ctx context.Context, wObj *kaitov1beta1.Workspace) (result *reconcile.Result, err error) {
	_ = "STUB: not implemented"
	// Provision nodes via the NodeProvisioner interface.
	// GpuProvisioner creates NodeClaims; BYOProvisioner (BYO mode) is a no-op.
	return nil, nil
}

// Check if nodes are ready.

func (c *WorkspaceReconciler) addOrUpdateWorkspace(ctx context.Context, wObj *kaitov1beta1.Workspace) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *WorkspaceReconciler) deleteWorkspace(ctx context.Context, wObj *kaitov1beta1.Workspace) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *WorkspaceReconciler) syncControllerRevision(ctx context.Context, wObj *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// nil checking.

func marshalSelectedFields(wObj *kaitov1beta1.Workspace) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ComputeHash(w *kaitov1beta1.Workspace) string { _ = "STUB: not implemented"; return "" }

func (c *WorkspaceReconciler) ensureService(ctx context.Context, wObj *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// headless service for worker pod to discover the leader pod

func (c *WorkspaceReconciler) applyTuning(ctx context.Context, wObj *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// applyInference applies inference spec.
func (c *WorkspaceReconciler) applyInference(ctx context.Context, wObj *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	// From v0.8.0 onwards, StatefulSet is the default workload for all workspaces.
	// This block purges existing Deployments and migrates them to StatefulSets later.
	// WARNING: This migration will cause a few minutes of service downtime.
	return nil
}

// TODO: handle update

// If the current workload revision matches the one in Workspace, we do not need to update it.

// Selectively update the pod spec fields that are relevant to inference,
// and leave the rest unchanged in case user has customized them.

// Update it with the latest one generated above.

func (c *WorkspaceReconciler) syncWorkspaceStatus(ctx context.Context, key types.NamespacedName, reconcileErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// Merge node conditions from provisioner: set returned conditions,
// remove any known node condition type that was not returned.

// Extract ResourceStatus condition status for downstream use.

type nodeStatusSnapshot struct {
	workerNodeNames []string
	conditions      []metav1.Condition
}

// nodeConditionTypes is the complete set of node-related condition types
// managed by NodeProvisioner. Conditions returned by CollectNodeStatusInfo
// are set; any type in this set not returned is removed from status.
var nodeConditionTypes = []string{
	string(kaitov1beta1.ConditionTypeNodeStatus),
	string(kaitov1beta1.ConditionTypeNodeClaimStatus),
	string(kaitov1beta1.ConditionTypeResourceStatus),
}

func (c *WorkspaceReconciler) collectNodeStatusSnapshot(ctx context.Context, wObj *kaitov1beta1.Workspace) (*nodeStatusSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect worker node names for status.

// Delegate status condition collection to the NodeProvisioner.

func (c *WorkspaceReconciler) collectInferenceReadyStatus(ctx context.Context, wObj *kaitov1beta1.Workspace) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type tuningStatusSnapshot struct {
	started   bool
	succeeded bool
	failed    bool
	active    int32
	ready     int32
}

func (c *WorkspaceReconciler) collectTuningStatusSnapshot(ctx context.Context, wObj *kaitov1beta1.Workspace) (*tuningStatusSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildReconcileErrMessageAppender(reconcileErr error) func(message string) string {
	_ = "STUB: not implemented"
	return nil
}

func setWorkspaceCondition(status *kaitov1beta1.WorkspaceStatus, generation int64, appendMessage func(string) string,
	conditionType kaitov1beta1.ConditionType, conditionStatus metav1.ConditionStatus, reason, message string) {
	_ = "STUB: not implemented"
	return
}

func applyTuningWorkspaceStatus(status *kaitov1beta1.WorkspaceStatus, generation int64, appendMessage func(string) string, snapshot *tuningStatusSnapshot) {
	_ = "STUB: not implemented"
	return
}

func applyInferenceWorkspaceStatus(ctx context.Context, status *kaitov1beta1.WorkspaceStatus, wObj *kaitov1beta1.Workspace, appendMessage func(string) string,
	inferenceReady bool, resourceConditionStatus metav1.ConditionStatus) {
	_ = "STUB: not implemented"
	return
}

// Clear benchmark state so applyBenchmarkStatus re-runs once inference recovers.
// This ensures a pod restart or rolling update doesn't leave stale results.

// applyBenchmarkStatus reads and parses the benchmark result from pod logs,
// then sets the BenchmarkCompleted condition on the workspace status.
// Returns nil on success (or when already recorded), non-nil on terminal failure.
// Skips the log read when BenchmarkCompleted is already True — the not-ready path
// clears the condition on any pod restart or rolling update.
func applyBenchmarkStatus(ctx context.Context, status *kaitov1beta1.WorkspaceStatus, wObj *kaitov1beta1.Workspace, generation int64, appendMessage func(string) string) error {
	_ = "STUB: not implemented"
	// Skip once the benchmark is done. Safe because the not-ready path clears BenchmarkCompleted
	// whenever inference goes down, so we won't get into a stale state.
	return nil
}

// These errors are terminal

func (c *WorkspaceReconciler) updateWorkspaceStatusIfChanged(ctx context.Context, key types.NamespacedName, modifyFn func(*kaitov1beta1.WorkspaceStatus) error) error {
	_ = "STUB: not implemented"
	return nil
}

func formatWorkspaceStatusChanges(oldStatus, newStatus kaitov1beta1.WorkspaceStatus) string {
	_ = "STUB: not implemented"
	return ""
}

// UpdateWorkspaceTargetNodeCount is used for updating the targetNodeCount in workspace status when it is 0.
func (c *WorkspaceReconciler) UpdateWorkspaceTargetNodeCount(ctx context.Context, wObj *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// Build the estimate request once, outside the status-update closure.

// Resolve the context window size from the workspace's inference ConfigMap (if any)
// and pass it through RuntimeProfile so the estimator does not need to do I/O.

// For non-vLLM runtime, use the Resource.Count directly
//nolint:staticcheck //SA1019: deprecate Resource.Count field

// Update the wObj to reflect the latest status change.

// SetupWithManager sets up the controller with the Manager.
func (c *WorkspaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Only watch NodeClaim resources if node auto-provisioning is enabled
