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

package gpuprovisioner

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/nodeprovision"
	"github.com/kaito-project/kaito/pkg/workspace/resource"
)

// AzureGPUProvisioner wraps the Azure gpu-provisioner
// (https://github.com/Azure/gpu-provisioner) logic behind the
// NodeProvisioner interface. It creates NodeClaims directly (the legacy
// path) and has no drift support.
type AzureGPUProvisioner struct {
	nodeClaimManager    *resource.NodeClaimManager
	nodeResourceManager *resource.NodeManager
}

var _ nodeprovision.NodeProvisioner = (*AzureGPUProvisioner)(nil)

// NewAzureGPUProvisioner creates an AzureGPUProvisioner that delegates to the existing
// NodeClaimManager and NodeManager.
func NewAzureGPUProvisioner(ncm *resource.NodeClaimManager, nm *resource.NodeManager) *AzureGPUProvisioner {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the provisioner name.
func (g *AzureGPUProvisioner) Name() string { _ = "STUB: not implemented"; return "" }

// Start is a no-op for AzureGPUProvisioner.
func (g *AzureGPUProvisioner) Start(ctx context.Context) error {
	_ = "STUB: not implemented"

	// ProvisionNodes creates NodeClaims via the Azure gpu-provisioner backend.
	return nil
}

func (g *AzureGPUProvisioner) ProvisionNodes(ctx context.Context, ws *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteNodes deletes all NodeClaims associated with the workspace.
func (g *AzureGPUProvisioner) DeleteNodes(ctx context.Context, ws *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

// EnableDriftRemediation is a no-op for Azure gpu-provisioner (no drift support).
func (g *AzureGPUProvisioner) EnableDriftRemediation(ctx context.Context, workspaceNamespace, workspaceName string) error {
	_ = "STUB: not implemented"

	// DisableDriftRemediation is a no-op for Azure gpu-provisioner (no drift support).
	return nil
}

func (g *AzureGPUProvisioner) DisableDriftRemediation(ctx context.Context, workspaceNamespace, workspaceName string) error {
	_ = "STUB: not implemented"

	// EnsureNodesReady checks that:
	//  1. All expected NodeClaims are in Ready state -> (false, false) if not.
	//  2. Enough Nodes with the correct instance type are ready -> (false, true) if not.
	//  3. GPU device plugins are installed on provisioned nodes -> (false, true) if not.
	return nil
}

func (g *AzureGPUProvisioner) EnsureNodesReady(ctx context.Context, ws *kaitov1beta1.Workspace) (bool, bool, error) {
	_ = "STUB: not implemented"
	// List nodes once and derive both readyNodes (for NodeClaim check) and
	// readyCount with correct instance type (for node readiness check).
	return false, false, nil
}

// Step 1: Check NodeClaims readiness.

// Step 2: Check that enough Nodes with the correct instance type are ready.

// Step 3: Check GPU device plugins on provisioned nodes.

// CollectNodeStatusInfo gathers status conditions for workspace status.
func (g *AzureGPUProvisioner) CollectNodeStatusInfo(ctx context.Context, ws *kaitov1beta1.Workspace) ([]metav1.Condition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeClaim readiness.

// Node readiness.

// Derive resource condition.
