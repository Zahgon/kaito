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
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/utils"
)

type NodeClaimManager struct {
	client.Client
	recorder               record.EventRecorder
	expectations           *utils.ControllerExpectations
	logger                 klog.Logger
	defaultNodeImageFamily string
}

func NewNodeClaimManager(c client.Client, recorder record.EventRecorder, expectations *utils.ControllerExpectations) *NodeClaimManager {
	_ = "STUB: not implemented"
	return nil
}

func (c *NodeClaimManager) SetDefaultNodeImageFamily(defaultNodeImageFamily string) {
	_ = "STUB: not implemented"
	return
}

// GetNumNodeClaimsNeeded calculates how many NodeClaims are needed to meet the target node count for the workspace.
func (c *NodeClaimManager) GetNumNodeClaimsNeeded(ctx context.Context, wObj *kaitov1beta1.Workspace, readyNodes []*corev1.Node) int {
	_ = "STUB: not implemented"
	return 0
}

// Count ready nodes that do NOT have a corresponding NodeClaim

// Calculate how many NodeClaims we need, including those already provisioned.

// CheckNodeClaims checks the current state of NodeClaims for the given workspace and determines how many additional NodeClaims need to be created to meet the target node count.
func (c *NodeClaimManager) CheckNodeClaims(ctx context.Context, wObj *kaitov1beta1.Workspace, readyNodes []*corev1.Node) (int, []*karpenterv1.NodeClaim, error) {
	_ = "STUB: not implemented"
	// We don't care in this case if the ready nodes come from NodeClaims, meaning ready nodes could come from BYO if the right size and properly labeled.
	return 0, nil, nil
}

// Calculate the total number of NodeClaims needed.

// Then, the number of NodeClaims to create is the difference between the total number needed and number of existing NodeClaims.

// CreateUpNodeClaims creates a specified number of NodeClaims as defined by nodesToCreate for the given workspace.
// this function will be invoked before creating workloads for workspace in order to ensure nodes.
func (c *NodeClaimManager) CreateUpNodeClaims(ctx context.Context, wObj *kaitov1beta1.Workspace, nodesToCreate int) error {
	_ = "STUB: not implemented"
	return nil
}

// Failed to create, decrement expectations

// should not return here or expectations will leak

// EnsureNodeClaimsReady is used for checking the number of ready nodeclaims(isNodeClaimReadyNotDeleting) meet the target NodeClaim count needed. Updates the
func (c *NodeClaimManager) EnsureNodeClaimsReady(ctx context.Context, wObj *kaitov1beta1.Workspace, readyNodes []*corev1.Node, existingNodeClaims []*karpenterv1.NodeClaim) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// determineNodeOSDiskSize returns the appropriate OS disk size for the workspace
func (c *NodeClaimManager) determineNodeOSDiskSize(ctx context.Context, wObj *kaitov1beta1.Workspace) string {
	_ = "STUB: not implemented"
	return ""
}

// The default OS size is used
