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

package karpenter

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/nodeprovision"
	"github.com/kaito-project/kaito/pkg/workspace/resource"
)

// NodeClassConfig holds cloud-specific NodeClass reference info.
// Group, Kind, Version, and ResourceName are injected via CLI flags.
// DefaultName is derived by Start() from ConfigMap labels.
type NodeClassConfig struct {
	Group        string // e.g. "karpenter.azure.com"
	Kind         string // e.g. "AKSNodeClass"
	Version      string // e.g. "v1beta1"
	ResourceName string // plural resource name (e.g. "aksnodeclasses"); combined with Group for CRD lookup
	DefaultName  string // populated by Start(): name of entry with karpenter.kaito.sh/default=true
}

// KarpenterProvisioner implements NodeProvisioner using the cloud-agnostic
// Karpenter API (NodePool / NodeClaim). Cloud-specific details (NodeClass
// group, kind, name mapping) are provided via NodeClassConfig.
type KarpenterProvisioner struct {
	client              client.Client
	nodeClassConfig     NodeClassConfig
	nodeResourceManager *resource.NodeManager
}

var _ nodeprovision.NodeProvisioner = (*KarpenterProvisioner)(nil)

// NewKarpenterProvisioner creates a new KarpenterProvisioner.
func NewKarpenterProvisioner(c client.Client, cfg NodeClassConfig) *KarpenterProvisioner {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the provisioner name.
func (p *KarpenterProvisioner) Name() string { _ = "STUB: not implemented"; return "" }

const nodeClassConfigMapName = "kaito-nodeclasses"

// Start verifies that the Karpenter CRDs are installed, creates
// NodeClass resources from the ConfigMap, and derives DefaultName from labels.
// Returns an error if Karpenter is not installed.
func (p *KarpenterProvisioner) Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Check if the core Karpenter CRDs exist.
	return nil
}

// Check if the provider-specific NodeClass CRD exists.

// Read the ConfigMap containing NodeClass manifests.

// Create each NodeClass and derive DefaultName from labels.

// Track the default NodeClass entry.

// Wait for the default NodeClass to be ready.

// checkNodeClassReady performs a single point-in-time check that the named
// NodeClass exists and has a Ready=True condition. Unlike waitForNodeClassReady,
// it does not poll — it returns an error immediately if the resource is missing
// or not ready.
func (p *KarpenterProvisioner) checkNodeClassReady(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// waitForNodeClassReady polls until the NodeClass has a Ready=True condition.
func (p *KarpenterProvisioner) waitForNodeClassReady(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// not created yet

// countCoveredNodes returns:
//   - coveredByNonKarpenter: nodes already handled outside karpenter. This includes:
//     (a) Ready BYO nodes (no NodeClaim, user-managed)
//     (b) Non-deleting legacy gpu-provisioner NodeClaims (whether their node is ready or not,
//     because karpenter will either self-heal them or delete them after a timeout,
//     at which point a reconcile is triggered via the NodeClaim watch)
//   - readyWithInstanceType: total ready nodes (BYO + legacy + karpenter) with correct instance type.
func countCoveredNodes(ctx context.Context, c client.Client, ws *kaitov1beta1.Workspace) (coveredByNonKarpenter int, readyWithInstanceType int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Count ready nodes (all types).

// BYO nodes: ready, correct instance type, no karpenter label, no legacy label.

// Count non-deleting legacy NodeClaims (whether ready or not).
// These are covered because karpenter engine will self-heal the node or
// delete the NodeClaim after a timeout (triggering a reconcile).

// ProvisionNodes creates or updates a NodePool for the Workspace.
// Computes delta-based replicas: desiredReplicas = max(0, targetNodeCount - coveredByNonKarpenterCount).
// If no NodePool exists and desiredReplicas is 0, no NodePool is created.
// If a NodePool exists, replicas are only increased (never decreased) to avoid
// disrupting running karpenter nodes when BYO nodes appear.
func (p *KarpenterProvisioner) ProvisionNodes(ctx context.Context, ws *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// Count non-karpenter ready nodes to compute delta.

// NodePool exists — only increase replicas, never decrease.
// This protects running karpenter nodes when BYO nodes appear after provisioning.

// DeleteNodes deletes the NodePool for the Workspace. Idempotent — NotFound is ignored.
// Karpenter cascades deletion: NodePool → NodeClaim → Node → VM.
func (p *KarpenterProvisioner) DeleteNodes(ctx context.Context, ws *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// nodeReadinessSnapshot holds pre-computed data about node and NodeClaim readiness
// for a workspace.
type nodeReadinessSnapshot struct {
	// readyWithInstanceTypeCount is the total number of ready nodes (BYO + legacy + karpenter)
	// that have the correct instance type for the workspace. Used to determine overall node readiness.
	readyWithInstanceTypeCount int
	// coveredByNonKarpenterCount is the number of nodes already handled outside karpenter:
	// ready BYO nodes + non-deleting legacy gpu-provisioner NodeClaims (regardless of node readiness).
	coveredByNonKarpenterCount int
	// targetNodeClaimCount is the number of karpenter NodeClaims needed:
	// max(0, ws.Status.TargetNodeCount - coveredByNonKarpenterCount).
	targetNodeClaimCount int
	// readyNodeClaims is the subset of karpenter-managed NodeClaims that are Ready and not deleting.
	// Used for GPU plugin readiness checks.
	readyNodeClaims []*karpenterv1.NodeClaim
}

// buildNodeReadinessSnapshot lists karpenter NodeClaims and all workspace nodes,
// returning counts needed for readiness decisions.
func (p *KarpenterProvisioner) buildNodeReadinessSnapshot(ctx context.Context, ws *kaitov1beta1.Workspace) (*nodeReadinessSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List karpenter NodeClaims.

// Count all ready nodes matching workspace labels (BYO + legacy + karpenter).

// EnsureNodesReady checks whether enough nodes are ready for the workspace.
// Counts all node types (BYO, legacy, karpenter) against targetNodeCount.
// Returns:
//   - ready: true when all expected nodes are present, Ready, and have GPU resources available.
//   - needRequeue: true when the caller should poll again because there is no watch/event
//     that will trigger reconciliation (e.g., Node registration or GPU plugin installation).
//   - err: non-nil on API errors.
func (p *KarpenterProvisioner) EnsureNodesReady(ctx context.Context, ws *kaitov1beta1.Workspace) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// Step 1: Check NodeClaim readiness.

// Step 2: Check that enough Nodes with the correct instance type are ready.

// Step 3: Check GPU device plugins on karpenter nodes.

// EnableDriftRemediation sets the Drifted budget to "1", allowing karpenter to replace drifted nodes.
func (p *KarpenterProvisioner) EnableDriftRemediation(ctx context.Context, workspaceNamespace, workspaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// DisableDriftRemediation sets the Drifted budget to "0", blocking karpenter from replacing drifted nodes.
func (p *KarpenterProvisioner) DisableDriftRemediation(ctx context.Context, workspaceNamespace, workspaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// setDriftBudget updates the Drifted budget entry in the NodePool.
// Uses RetryOnConflict with Get+Update inside the retry closure for optimistic concurrency.
func (p *KarpenterProvisioner) setDriftBudget(ctx context.Context, workspaceNamespace, workspaceName, nodes string) error {
	_ = "STUB: not implemented"
	return nil
}

// CollectNodeStatusInfo gathers status conditions for workspace status.
// Counts all ready nodes (BYO + legacy + karpenter) against targetNodeCount.
// GPU plugin readiness is checked only on karpenter NodeClaims.
func (p *KarpenterProvisioner) CollectNodeStatusInfo(ctx context.Context, ws *kaitov1beta1.Workspace) ([]metav1.Condition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeClaim condition: are enough karpenter NodeClaims ready?
// Non-karpenter nodes (BYO, legacy gpu-provisioner) reduce the target —
// they don't have karpenter NodeClaims, so we only need NodeClaims for the remainder.
// Legacy gpu-provisioner nodes also have NodeClaims, but those are already stable
// and not managed by this provisioner, so we treat them like BYO here.

// Node condition: are enough nodes ready with GPU resources?
// Uses total ready node count (all types). GPU plugin check only applies
// to karpenter NodeClaims — BYO/legacy nodes are assumed GPU-ready.

// No karpenter NodeClaims — all nodes are BYO/legacy, assume GPU ready.

// Derive resource condition.
