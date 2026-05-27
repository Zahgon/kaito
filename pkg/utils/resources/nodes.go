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

package resources

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

const (
	LabelKeyNvidia    = "accelerator"
	LabelValueNvidia  = "nvidia"
	CapacityNvidiaGPU = "nvidia.com/gpu"
)

// GetNode get kubernetes node object with a provided name
func GetNode(ctx context.Context, nodeName string, kubeClient client.Client) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListNodes get list of kubernetes nodes
func ListNodes(ctx context.Context, kubeClient client.Client, labelSelector client.MatchingLabels) (*corev1.NodeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateNodeWithLabel update the node object with the label key/value
func UpdateNodeWithLabel(ctx context.Context, freshNode *corev1.Node, labelKey, labelValue string, kubeClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckNvidiaPlugin(ctx context.Context, nodeObj *corev1.Node) bool {
	_ = "STUB: not implemented"
	// check if label accelerator=nvidia exists in the node
	return false
}

// check Status.Capacity.nvidia.com/gpu has value

func ExtractObjFields(obj client.Object) (instanceType, namespace, name string, labelSelector *metav1.LabelSelector,
	nameLabel, namespaceLabel string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil, "", "", nil
}

// GetReadyNodes finds all ready nodes that match the workspace's label selector
func GetReadyNodes(ctx context.Context, c client.Client, wObj *kaitov1beta1.Workspace) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NodeIsReadyAndNotDeleting(node *corev1.Node) bool { _ = "STUB: not implemented"; return false }
