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

package byoprovisioner

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/nodeprovision"
)

// BYOProvisioner is a no-op NodeProvisioner for BYO (Bring Your Own) node
// scenarios where node auto-provisioning is disabled. ProvisionNodes and
// DeleteNodes are no-ops. EnsureNodesReady only checks that enough
// matching Nodes are ready (no instance type validation, no GPU plugin checks).
type BYOProvisioner struct {
	client client.Client
}

var _ nodeprovision.NodeProvisioner = (*BYOProvisioner)(nil)

func NewBYOProvisioner(c client.Client) *BYOProvisioner { _ = "STUB: not implemented"; return nil }

// Name returns the provisioner name.
func (n *BYOProvisioner) Name() string { _ = "STUB: not implemented"; return "" }

// Start is a no-op for BYOProvisioner.
func (n *BYOProvisioner) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *BYOProvisioner) ProvisionNodes(ctx context.Context, ws *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *BYOProvisioner) DeleteNodes(ctx context.Context, ws *kaitov1beta1.Workspace) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *BYOProvisioner) EnableDriftRemediation(ctx context.Context, workspaceNamespace, workspaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *BYOProvisioner) DisableDriftRemediation(ctx context.Context, workspaceNamespace, workspaceName string) error {
	_ = "STUB: not implemented"

	// EnsureNodesReady checks that enough matching Nodes are ready for the
	// Workspace. In BYO mode there are no provisioning resources, so needRequeue
	// is always true when nodes are not ready.
	return nil
}

func (n *BYOProvisioner) EnsureNodesReady(ctx context.Context, ws *kaitov1beta1.Workspace) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// CollectNodeStatusInfo gathers status conditions for workspace status.
// In BYO mode, no NodeClaimStatus condition is returned.
func (n *BYOProvisioner) CollectNodeStatusInfo(ctx context.Context, ws *kaitov1beta1.Workspace) ([]metav1.Condition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BYO mode: no NodeClaimStatus condition.
