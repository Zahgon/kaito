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

package nodeclaim

import (
	"context"
	"reflect"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/utils/consts"
)

var (
	// nodeClaimStatusTimeoutInterval is the interval to check the nodeClaim status.
	nodeClaimStatusTimeoutInterval = 240 * time.Second

	WorkspaceSelector, _ = metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{Key: kaitov1beta1.LabelWorkspaceName, Operator: metav1.LabelSelectorOpExists},
		},
	})

	KarpenterWorkspaceSelector, _ = metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{Key: consts.KarpenterWorkspaceNameKey, Operator: metav1.LabelSelectorOpExists},
		},
	})

	RagEngineSelector, _ = metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{Key: kaitov1beta1.LabelRAGEngineName, Operator: metav1.LabelSelectorOpExists},
		},
	})

	NodeClaimPredicate = predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			nodeclaim, ok := e.Object.(*karpenterv1.NodeClaim)
			if !ok {
				return false
			}
			return isRelevantNodeClaim(nodeclaim.GetLabels())
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldNodeClaim, ok := e.ObjectOld.(*karpenterv1.NodeClaim)
			if !ok {
				return false
			}

			newNodeClaim, ok := e.ObjectNew.(*karpenterv1.NodeClaim)
			if !ok {
				return false
			}
			if !isRelevantNodeClaim(oldNodeClaim.GetLabels()) {
				return false
			}

			if !isRelevantNodeClaim(newNodeClaim.GetLabels()) {
				return false
			}

			oldNodeClaimCopy := oldNodeClaim.DeepCopy()
			newNodeClaimCopy := newNodeClaim.DeepCopy()

			// if only nodeclaim.Status.LastPodEventTime is changed, skip update event
			oldNodeClaimCopy.ResourceVersion = ""
			oldNodeClaimCopy.Status.LastPodEventTime = metav1.Time{}
			newNodeClaimCopy.ResourceVersion = ""
			newNodeClaimCopy.Status.LastPodEventTime = metav1.Time{}
			return !reflect.DeepEqual(oldNodeClaimCopy, newNodeClaimCopy)
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			nodeclaim, ok := e.Object.(*karpenterv1.NodeClaim)
			if !ok {
				return false
			}
			return isRelevantNodeClaim(nodeclaim.GetLabels())
		},
	}
)

// isRelevantNodeClaim returns true if the NodeClaim has labels indicating it
// belongs to a Workspace (legacy kaito.sh/* or karpenter.kaito.sh/*) or a RAGEngine.
func isRelevantNodeClaim(lbls map[string]string) bool { _ = "STUB: not implemented"; return false }

type ManifestOptions struct {
	DefaultNodeImageFamily string
}

// GenerateNodeClaimManifest generates a nodeClaim object from the given workspace or RAGEngine.
func GenerateNodeClaimManifest(storageRequirement string, obj client.Object) *karpenterv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

// GenerateNodeClaimManifestWithOptions generates a nodeClaim object from the given workspace or RAGEngine with explicit options.
func GenerateNodeClaimManifestWithOptions(storageRequirement string, obj client.Object, options ManifestOptions) *karpenterv1.NodeClaim {
	_ = "STUB: not implemented"
	return nil
}

// Determine the type of the input object and extract relevant fields

// Fake nodepool name to prevent Karpenter from scaling up.

// To prevent Karpenter from scaling down.

//azure

//aws

// GenerateNodeClaimName generates a nodeClaim name from the given workspace or RAGEngine.
func GenerateNodeClaimName(obj client.Object) string {
	_ = "STUB: not implemented"
	// Determine the type of the input object and extract relevant fields
	return ""
}

// We make sure the nodeClaim name is not fixed to the object

// CreateNodeClaim creates a nodeClaim object.
func CreateNodeClaim(ctx context.Context, nodeClaimObj *karpenterv1.NodeClaim, kubeClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPendingNodeClaims checks if there are any nodeClaims in provisioning condition. If so, wait until they are ready.
func WaitForPendingNodeClaims(ctx context.Context, obj client.Object, kubeClient client.Client) error {
	_ = "STUB: not implemented"

	// Determine the type of the input object and retrieve the InstanceType
	return nil
}

// check if the nodeClaim being created has the requested instance type

// wait until the nodeClaim is initialized

// ListNodeClaim lists all nodeClaim objects in the cluster that are created by the given workspace or RAGEngine.
func ListNodeClaim(ctx context.Context, obj client.Object, kubeClient client.Client) (*karpenterv1.NodeClaimList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build label selector based on the type of the input object

// CheckNodeClaimStatus checks the status of the nodeClaim. If the nodeClaim is not ready, then it will wait for the nodeClaim to be ready.
// If the nodeClaim is not ready after the timeout, then it will return an error.
// if the nodeClaim is ready, then it will return nil.
func CheckNodeClaimStatus(ctx context.Context, nodeClaimObj *karpenterv1.NodeClaim, kubeClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// if SKU is not available, then no need to retry.

// if nodeClaim is not ready, then continue.

// IsNodeClaimReadyNotDeleting checks if a NodeClaim is in ready state and not being deleted
func IsNodeClaimReadyNotDeleting(nodeClaim *karpenterv1.NodeClaim) bool {
	_ = "STUB: not implemented"
	return false
}
