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

package multiroleinference

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1alpha1 "github.com/kaito-project/kaito/api/v1alpha1"
)

const (
	// MultiRoleInferenceFinalizer is the finalizer for MultiRoleInference objects.
	MultiRoleInferenceFinalizer = "multiroleinference.kaito.sh/finalizer"

	// ConditionTypeDeleting indicates the MRI is being deleted.
	ConditionTypeDeleting = "Deleting"
)

// MultiRoleInferenceReconciler reconciles a MultiRoleInference object.
type MultiRoleInferenceReconciler struct {
	client.Client
	Log                          logr.Logger
	Scheme                       *runtime.Scheme
	Recorder                     record.EventRecorder
	EnableGatewayAPIInferenceExt bool
}

// NewMultiRoleInferenceReconciler creates a new reconciler.
func NewMultiRoleInferenceReconciler(client client.Client, scheme *runtime.Scheme, log logr.Logger, recorder record.EventRecorder, enableGWIE bool) *MultiRoleInferenceReconciler {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups=kaito.sh,resources=multiroleinferences,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kaito.sh,resources=multiroleinferences/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kaito.sh,resources=multiroleinferences/finalizers,verbs=update
// +kubebuilder:rbac:groups=kaito.sh,resources=inferencesets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=source.toolkit.fluxcd.io,resources=ocirepositories,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=helm.toolkit.fluxcd.io,resources=helmreleases,verbs=get;list;watch;create;update;patch;delete

func (r *MultiRoleInferenceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Fetch the MultiRoleInference instance.

// Handle deletion vs normal reconciliation.

// Ensure finalizer is present.

// MRI is being deleted — run garbage collection.

// ensureFinalizer adds the finalizer to the MRI if not already present.
func (r *MultiRoleInferenceReconciler) ensureFinalizer(ctx context.Context, mri *kaitov1alpha1.MultiRoleInference) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteMultiRoleInference handles MRI deletion: sets Deleting condition, GCs children, removes finalizer.
func (r *MultiRoleInferenceReconciler) deleteMultiRoleInference(ctx context.Context, log logr.Logger, mri *kaitov1alpha1.MultiRoleInference) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Set Deleting condition.

// garbageCollectMultiRoleInference deletes all child InferenceSets and removes
// the finalizer. OCIRepository and HelmRelease are GC'd via ownerReferences.
func (r *MultiRoleInferenceReconciler) garbageCollectMultiRoleInference(ctx context.Context, log logr.Logger, mri *kaitov1alpha1.MultiRoleInference) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	// List all child InferenceSets owned by this MRI.
	return *new(ctrl.Result), nil
}

// Delete each child InferenceSet that hasn't been deleted yet.

// Wait until all child InferenceSets are fully removed before removing the finalizer.

// Remove the finalizer.

// addOrUpdateMultiRoleInference handles normal reconciliation: create/update child InferenceSets.
func (r *MultiRoleInferenceReconciler) addOrUpdateMultiRoleInference(ctx context.Context, log logr.Logger, mri *kaitov1alpha1.MultiRoleInference) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Create or update child InferenceSets for each role.

// Clean up stale InferenceSets (roles removed from spec).

// Reconcile InferencePool via Flux OCIRepository + HelmRelease — only when GWIE is enabled.
// EPP plugins config is passed inline through Helm values (pluginsCustomConfig),
// so no separate ConfigMap reconciliation is needed.

// Aggregate status from child InferenceSets.

// aggregateStatus reads child InferenceSet conditions and updates MRI status accordingly.
func (r *MultiRoleInferenceReconciler) aggregateStatus(ctx context.Context, log logr.Logger, mri *kaitov1alpha1.MultiRoleInference) error {
	_ = "STUB: not implemented"
	// List all child InferenceSets.
	return nil
}

// Build a map from role → InferenceSet.

// Check individual role readiness.

// Set prefill condition.

// Check decode InferenceSet readiness.

// Check InferencePool readiness.

// Set overall Ready condition.

// isInferenceSetReady checks if an InferenceSet has the InferenceSetReady condition set to True.
func (r *MultiRoleInferenceReconciler) isInferenceSetReady(is *kaitov1alpha1.InferenceSet) bool {
	_ = "STUB: not implemented"
	return false
}

// cleanupStaleInferenceSets deletes InferenceSets whose role has been removed from the MRI spec.
func (r *MultiRoleInferenceReconciler) cleanupStaleInferenceSets(ctx context.Context, mri *kaitov1alpha1.MultiRoleInference) error {
	_ = "STUB: not implemented"
	// Build set of expected InferenceSet names.
	return nil
}

// List all child InferenceSets.

// reconcileInferenceSet creates or updates a child InferenceSet for the given role.
func (r *MultiRoleInferenceReconciler) reconcileInferenceSet(
	ctx context.Context,
	mri *kaitov1alpha1.MultiRoleInference,
	role kaitov1alpha1.MultiRoleInferenceRoleSpec,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Build the desired InferenceSet.

// Set owner reference so the InferenceSet is garbage-collected with the MRI.

// Labels on the InferenceSet metadata.

// Spec — only reconcile replicas when explicitly set (non-nil).
// When nil, autoscaling is assumed and the controller skips replica reconciliation.

// LabelSelector — start from the MRI's labelSelector and inject role info.
// The InferenceSet controller propagates Spec.Selector to workspace.Resource.LabelSelector,
// so role-specific labels must be in the selector to ensure correct node selection.

// Template metadata labels: propagate selector matchLabels (includes role labels).

// Resource.

// Inference — preset with shared model config.

// Role-specific runtime config.

const (
	// eppPluginsConfigKey is the filename key for EPP plugins config,
	// used both as the ConfigMap data key and the Helm pluginsConfigFile name.
	eppPluginsConfigKey = "config.yaml"
)

// defaultPDPluginsConfigTemplate is the default EPP plugins YAML template for P/D disaggregated serving.
// Uses the llm-d EndpointPickerConfig format with schedulingProfiles for prefill and decode.
const defaultPDPluginsConfigTemplate = `apiVersion: inference.networking.x-k8s.io/v1alpha1
kind: EndpointPickerConfig
plugins:
  - type: disagg-headers-handler
  - type: prefix-based-pd-decider
    parameters:
      nonCachedTokens: 4
  - type: disagg-profile-handler
    parameters:
      deciders:
        prefill: prefix-based-pd-decider
  - type: by-label-selector
    name: prefill-filter
    parameters:
      matchLabels:
        kaito.sh/inference-role: prefill
  - type: by-label-selector
    name: decode-filter
    parameters:
      matchLabels:
        kaito.sh/inference-role: decode
  - type: load-aware-scorer
    parameters:
      threshold: 10
  - type: max-score-picker
schedulingProfiles:
  - name: prefill
    plugins:
      - pluginRef: prefill-filter
      - pluginRef: load-aware-scorer
        weight: 10
      - pluginRef: max-score-picker
  - name: decode
    plugins:
      - pluginRef: decode-filter
      - pluginRef: load-aware-scorer
        weight: 10
      - pluginRef: max-score-picker
`

// defaultPDPluginsConfig returns the default P/D plugins config.
// Note: precise-prefix-cache-scorer is omitted because it requires a tokenizer
// sidecar (UDS socket) that is not yet deployed by KAITO. Once tokenizer sidecar
// support is added, this config should be updated to include it.
func defaultPDPluginsConfig() string { _ = "STUB: not implemented"; return "" }

// inferencePoolName returns the name of the InferencePool resources for the MRI.
// Delegates to the shared utils.InferencePoolName helper.
func inferencePoolName(mriName string) string { _ = "STUB: not implemented"; return "" }

// reconcileInferencePool creates or updates the Flux OCIRepository and HelmRelease
// that render an InferencePool CR + EPP deployment, owned by the MRI.
func (r *MultiRoleInferenceReconciler) reconcileInferencePool(
	ctx context.Context,
	mri *kaitov1alpha1.MultiRoleInference,
) error {
	_ = "STUB: not implemented"
	return nil
}

// --- OCIRepository ---

// --- HelmRelease ---
// InferencePool selects ALL MRI pods (prefill + decode). EPP's internal
// prefill-filter / decode-filter plugins handle role-based selection.
// targetPort=5001 (sidecar) is only used by Envoy for user-facing traffic;
// EPP routes user requests exclusively to decode pods via decode-filter.
// Prefill communication is initiated by the decode sidecar directly to
// prefill pod:5000, bypassing InferencePool/Envoy entirely.

// Only leader pod (ordinal 0) serves inference traffic

// Build EPP extension values with llm-d image and P/D plugins config.

// Load plugins config: either from user-provided ConfigMap or auto-generated default.

// Disable EPP secure-serving (self-signed TLS) — MRI pools run plaintext
// behind the mesh, matching standalone InferenceSet behavior.

// routing sidecar port; GWIE CRD allows only one targetPort (maxItems: 1)

// isInferencePoolReady checks if the InferencePool HelmRelease is ready.
// When Gateway API Inference Extension is disabled, no pool is reconciled so we return true.
func (r *MultiRoleInferenceReconciler) isInferencePoolReady(ctx context.Context, mri *kaitov1alpha1.MultiRoleInference) bool {
	_ = "STUB: not implemented"
	return false
}

// SetupWithManager sets up the controller with the Manager.
func (r *MultiRoleInferenceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Only watch Flux resources when Gateway API Inference Extension is enabled,
// because the Flux CRDs are only installed under that feature gate.

// Verify prerequisite CRDs exist before configuring watches.
