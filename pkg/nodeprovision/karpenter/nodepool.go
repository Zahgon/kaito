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
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

const (
	maxNodePoolNameLen = 253
	hashSuffixLen      = 9
)

// truncatedName returns a deterministic, truncated string with a hash suffix
// for uniqueness when the input exceeds maxLen.
func truncatedName(workspaceNamespace, workspaceName string, maxLen int) string {
	_ = "STUB: not implemented"
	return ""
}

// 1 for dash separator

// NodePoolName returns a deterministic, DNS-safe name for the NodePool
// derived from the workspace namespace and name.
// If the result exceeds 253 characters, it is truncated and a 9-char
// SHA-256 hex suffix is appended for uniqueness.
func NodePoolName(workspaceNamespace, workspaceName string) string {
	_ = "STUB: not implemented"
	return ""
}

// resolveNodeClassName determines the NodeClass resource name for a Workspace.
// It checks for the node-class-name annotation on the workspace, then falls
// back to the configured default.
func resolveNodeClassName(ws *kaitov1beta1.Workspace, cfg NodeClassConfig) string {
	_ = "STUB: not implemented"
	return ""
}

// isInferenceSetWorkspace returns true if the Workspace was created by an InferenceSet.
func isInferenceSetWorkspace(ws *kaitov1beta1.Workspace) bool {
	_ = "STUB: not implemented"
	return false
}

// nodePoolRequirements builds the NodePool requirements list.
// The instance-type requirement is always included. Provider-specific
// requirements (e.g. Azure placement scope) are added based on the
// NodeClassConfig group.
func nodePoolRequirements(ws *kaitov1beta1.Workspace, cfg NodeClassConfig) []karpenterv1.NodeSelectorRequirementWithMinValues {
	_ = "STUB: not implemented"
	return nil
}

// Azure Karpenter requires regional placement scope.

// generateNodePool builds a karpenter NodePool manifest for the given Workspace.
func generateNodePool(ws *kaitov1beta1.Workspace, cfg NodeClassConfig) *karpenterv1.NodePool {
	_ = "STUB: not implemented"
	return nil
}

// Drift budget: InferenceSet workspaces start with "0" (blocked),
// standalone workspaces use "1" (karpenter handles autonomously).

// Template labels propagated to NodeClaims and Nodes.

// Include the user's matchLabels so that inference pods' nodeAffinity
// (built from matchLabels) is satisfied. KAITO-reserved keys are stripped
// to avoid clobbering controller-managed labels.

// InferenceSet workspaces get additional labels so the drift controller
// can map NodeClaim events back to the owning InferenceSet.

// NodePool-level labels for management and lookup.

// InferenceSet workspaces get labels on NodePool ObjectMeta so the drift
// controller can List NodePools by InferenceSet directly.
