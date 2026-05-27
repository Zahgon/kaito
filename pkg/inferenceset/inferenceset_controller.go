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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kaitov1alpha1 "github.com/kaito-project/kaito/api/v1alpha1"
	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/utils"
)

const (
	InferenceSetHashAnnotation = "inferenceset.kaito.io/hash"
	InferenceSetNameLabel      = "inferenceset.kaito.io/name"
	revisionHashSuffix         = 5
)

type InferenceSetReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder

	klogger      klog.Logger
	expectations *utils.ControllerExpectations
}

func NewInferenceSetReconciler(client client.Client, scheme *runtime.Scheme, log logr.Logger, Recorder record.EventRecorder) *InferenceSetReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (c *InferenceSetReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Handle deleting inferenceset, garbage collect all the resources.

func (c *InferenceSetReconciler) ensureFinalizer(ctx context.Context, iObj *kaitov1alpha1.InferenceSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *InferenceSetReconciler) deleteInferenceSet(ctx context.Context, iObj *kaitov1alpha1.InferenceSet) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// garbageCollectInferenceSet remove finalizer associated with inferenceset object.
func (c *InferenceSetReconciler) garbageCollectInferenceSet(ctx context.Context, iObj *kaitov1alpha1.InferenceSet) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Check if there are any workspaces associated with this inferenceset.

// We should delete all the workspaces that are created by this inferenceset

// aggregateBenchmarkResults scans workspaces and returns:
//   - totalTPM: sum of peakTokensPerMinute across all succeeded workspaces that have a valid result
//   - readyReplicas: count of succeeded workspaces
//   - benchmarkedReplicas: count of those workspaces
//   - hasBenchmarkTPMResult: true if at least one workspace contributed a TPM value
func aggregateBenchmarkResults(workspaces []kaitov1beta1.Workspace) (totalTPM float64, readyReplicas, benchmarkedReplicas int, hasBenchmarkTPMResult bool) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false
}

func (c *InferenceSetReconciler) addOrUpdateInferenceSet(ctx context.Context, iObj *kaitov1alpha1.InferenceSet) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Check if there are any existing workspaces associated with this inferenceset.

// first delete workspace that is not in ready state

// delete rest of extra workspaces

// check whether ws.Name is already in deletingWorkspaces

// After deleting the extra workspaces, we should requeue to wait for the deletion to complete

// Start with labels from the template metadata, then add controller labels.

// Also propagate select labels from the InferenceSet's own metadata,
// in case template.metadata.labels was pruned by the API server.

// Start with annotations from the template metadata.

// Propagate the disable-benchmark opt-out so each child workspace inherits it.
// Benchmark is on by default; only propagate when explicitly disabled.

// Reconcile labels on existing workspaces by additively propagating InferenceSet metadata labels.
// Note: this only adds/updates desired labels; it does not remove stale labels to avoid
// conflicting with labels managed by other controllers.
// This ensures label changes (e.g., adding kaito.sh/inference-role) propagate
// to workspaces that were created before the label was set.

// Propagate inference-role from InferenceSet metadata (reliable even if template labels are pruned).

// check whether all the workspaces are ready

// update the replicas in the status

// set selector for HPA/VPA

// No ready replica has a TPM result — clear the TPM key so the profile
// doesn't reflect a previous generation of workspaces.
// Other metric keys are left intact to be cleared by their own logic.

// Feature flag is off — clear any TPM value that may have been written
// when the flag was previously enabled (e.g. annotation removed).

// Surface benchmark progress when the annotation is set.

// ensureGatewayAPIInferenceExtension reconciles Gateway API Inference Extension components for a InferenceSet.
//
// How it works:
// 1) Dry-runs preset inference generation to determine if the target workload is a StatefulSet.
// 2) Renders a Flux OCIRepository and a HelmRelease for the InferencePool chart.
// 3) Creates the resources if absent; updates them if the desired spec differs.
// 4) Waits for resources to become ready using the model's inference readiness timeout.
// 5) Aggregates and returns any errors.
//
// Idempotent and safe to call on every reconcile; no-op if preconditions are not met.
func (c *InferenceSetReconciler) ensureGatewayAPIInferenceExtension(ctx context.Context, iObj *kaitov1alpha1.InferenceSet) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip GWIE for child InferenceSets managed by MultiRoleInference.
// The MRI controller creates a shared InferencePool + EPP for all child InferenceSets.
// Use OwnerReferences (controller-managed) instead of labels (easily user-modifiable)
// to prevent accidental GWIE bypass on standalone InferenceSets.

// Gateway API Inference Extension is specifically designed to work with vLLM and preset-based inference workloads.

// Create or update OCIRepository

// Check if HelmRelease exists

func (c *InferenceSetReconciler) syncControllerRevision(ctx context.Context, iObj *kaitov1alpha1.InferenceSet) error {
	_ = "STUB: not implemented"
	return nil
}

// nil checking.

// SetupWithManager sets up the controller with the Manager.
func (c *InferenceSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify that all prerequisite CRDs exist before configuring watches that depend on them.
// - FluxCD HelmRelease / OCIRepository: required for installing and reconciling the InferencePool Helm chart.
// - Gateway API Inference Extension InferencePool / InferenceModel: required runtime CRDs that the Workspace
//   controller indirectly relies on (Helm chart renders resources referencing them).
// Failing fast here provides a clear, actionable error instead of deferred reconcile failures later.

// We don't need to own InferencePool and InferenceModel because they are managed by Flux's HelmRelease
