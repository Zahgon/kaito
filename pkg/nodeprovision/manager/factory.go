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

package manager

import (
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kaito-project/kaito/pkg/nodeprovision"
)

// ProvisionerConfig holds all parameters needed to create a NodeProvisioner.
type ProvisionerConfig struct {
	KClient                client.Client
	DirectClient           client.Client
	Recorder               record.EventRecorder
	DefaultNodeImageFamily string
	ProvisionerType        string
	NodeClassGroup         string
	NodeClassKind          string
	NodeClassVersion       string
	NodeClassResourceName  string
}

// NewNodeProvisioner creates and returns a NodeProvisioner based on the provisionerType parameter.
//
//   - karpenter: KarpenterProvisioner (cloud-agnostic karpenter NodePool CRUD).
//   - byo: BYOProvisioner (all provisioning ops are no-ops).
//   - azure-gpu-provisioner (default): AzureGPUProvisioner (creates/deletes NodeClaims).
func NewNodeProvisioner(cfg ProvisionerConfig) nodeprovision.NodeProvisioner {
	_ = "STUB: not implemented"
	return *new(nodeprovision.NodeProvisioner)
}

// consts.NodeProvisionerAzureGPU
