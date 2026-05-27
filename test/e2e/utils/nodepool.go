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

package utils

import (
	"context"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

// ValidateWorkspaceTargetNodeCount verifies workspace.status.targetNodeCount
// matches the expected value.
func ValidateWorkspaceTargetNodeCount(ctx context.Context, workspaceObj *kaitov1beta1.Workspace, expectedCount int) {
	_ = "STUB: not implemented"
	return
}

// ValidateNodePoolShape verifies the NodePool created for a workspace has the
// expected structure: name, labels, nodeClassRef, requirements, taints,
// disruption config, and replicas.
func ValidateNodePoolShape(ctx context.Context, workspaceObj *kaitov1beta1.Workspace, expectedReplicas int) {
	_ = "STUB: not implemented"
	return
}

// --- Metadata ---

// --- Replicas ---

// --- Template Labels ---

// --- NodeClassRef ---

// --- Requirements ---

// --- Taints ---

// --- Disruption ---

// ValidateInferenceSetNodePoolShape verifies NodePool shape for an
// InferenceSet-managed workspace, including InferenceSet-specific labels
// and drift budget "0".
func ValidateInferenceSetNodePoolShape(ctx context.Context, workspaceObj *kaitov1beta1.Workspace,
	expectedReplicas int, inferenceSetName string) {
	_ = "STUB: not implemented"
	return
}

// --- Metadata labels ---

// InferenceSet labels on NodePool metadata

// --- Replicas ---

// --- Template labels ---

// --- NodeClassRef ---

// --- Requirements ---

// --- Taints ---

// --- Disruption — InferenceSet uses budget "0" ---

// ValidateNodeLabels verifies that karpenter-provisioned Nodes have the
// expected labels propagated from the NodePool template.
func ValidateNodeLabels(ctx context.Context, workspaceObj *kaitov1beta1.Workspace) {
	_ = "STUB: not implemented"
	return
}

// ValidateNodePoolNodeClassRef verifies that the NodePool for a workspace
// references the expected NodeClass name.
func ValidateNodePoolNodeClassRef(ctx context.Context, workspaceObj *kaitov1beta1.Workspace, expectedNodeClassName string) {
	_ = "STUB: not implemented"
	return
}

// ValidateNodePoolDeletion verifies that the NodePool for a workspace
// has been deleted (via workspace finalizer).
func ValidateNodePoolDeletion(ctx context.Context, workspaceObj *kaitov1beta1.Workspace) {
	_ = "STUB: not implemented"
	return
}

// still exists

// NotFound

// ValidateNodePoolIsolation verifies that multiple workspaces have distinct
// NodePools with no shared NodeClaims.
func ValidateNodePoolIsolation(ctx context.Context, workspaces []*kaitov1beta1.Workspace) {
	_ = "STUB: not implemented"
	return
}

// nodePoolName -> workspaceName
// nodeClaimName -> nodePoolName
