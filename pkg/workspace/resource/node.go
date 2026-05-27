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

package resource

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

type NodeManager struct {
	client.Client
}

func NewNodeManager(c client.Client) *NodeManager { _ = "STUB: not implemented"; return nil }

// CheckIfNodePluginsReady is used for ensuring node label(accelerator:nvidia) and GPU capacity on all auto-provisioned nodes for the workspace.
func (c *NodeManager) CheckIfNodePluginsReady(ctx context.Context, wObj *kaitov1beta1.Workspace, existingNodeClaims []*karpenterv1.NodeClaim) (bool, error) {
	_ = "STUB: not implemented"
	// ensure Nvidia device plugins are ready for the workspace when instance type is known.
	return false, nil
}

// checkNodePlugin ensures that NVIDIA device plugins are ready on all nodes for the workspace
func (c *NodeManager) checkNodePlugin(ctx context.Context, wObj *kaitov1beta1.Workspace, existingNodeClaims []*karpenterv1.NodeClaim) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check each node for NVIDIA accelerator label and GPU capacity

// getReadyNodesFromNodeClaims retrieves all ready nodes that are associated with NodeClaims for the workspace.
// This function excludes preferred nodes and only returns nodes that were provisioned through NodeClaims.
// It's primarily used for device plugin management where we need to ensure GPU nodes created by
// NodeClaims have the proper NVIDIA device plugins installed.
func (c *NodeManager) getReadyNodesFromNodeClaims(ctx context.Context, wObj *kaitov1beta1.Workspace, existingNodeClaims []*karpenterv1.NodeClaim) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
