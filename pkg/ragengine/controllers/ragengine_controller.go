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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

const (
	RAGEngineHashAnnotation = "ragengine.kaito.io/hash"
	RAGEngineNameLabel      = "ragengine.kaito.io/name"
	revisionHashSuffix      = 5
)

type RAGEngineReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

func NewRAGEngineReconciler(client client.Client, scheme *runtime.Scheme, log logr.Logger, Recorder record.EventRecorder) *RAGEngineReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (c *RAGEngineReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Handle deleting ragengine, garbage collect all the resources.

func (c *RAGEngineReconciler) ensureFinalizer(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RAGEngineReconciler) addRAGEngine(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// set resource status to true when no compute resource is needed.

func (c *RAGEngineReconciler) ensureService(ctx context.Context, ragObj *kaitov1beta1.RAGEngine) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure Service for index and query
// TODO: ServiceName currently does not accept customization for now

func (c *RAGEngineReconciler) applyRAG(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) error {
	_ = "STUB: not implemented"
	return nil
}

// Currently, all CRD changes are only passed through environment variables (env)

// Need to create a new workload

func (c *RAGEngineReconciler) deleteRAGEngine(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (c *RAGEngineReconciler) syncControllerRevision(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) error {
	_ = "STUB: not implemented"
	return nil
}

// nil checking.

func computeHash(ragEngineObj *kaitov1beta1.RAGEngine) string { _ = "STUB: not implemented"; return "" }

// applyRAGEngineResource applies RAGEngine resource spec.
func (c *RAGEngineReconciler) applyRAGEngineResource(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) error {
	_ = "STUB: not implemented"
	// Wait for pending nodeClaims if any before we decide whether to create new node or not.
	return nil
}

// Find all nodes that match the labelSelector and instanceType, they are not necessarily created by machines/nodeClaims.

// RAGEngine requires exactly 1 node

// No existing nodes, need to create one

// Select the best qualified node from existing nodes

// Ensure all gpu plugins are running successfully.

// If GetGPUConfigBySKU returns error, skip GPU plugin installation (e.g., CPU-only instances)

// Add the valid nodes names to the RAGEngineStatus.WorkerNodes.

// getAllQualifiedNodes returns all nodes that match the labelSelector and instanceType.
func (c *RAGEngineReconciler) getAllQualifiedNodes(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip nodes that are being deleted

// skip nodes that are not ready

// match the instanceType

// createAndValidateNode creates a new node and validates status.
func (c *RAGEngineReconciler) createAndValidateNode(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The default OS size is used

func (c *RAGEngineReconciler) CreateNodeClaim(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine, nodeOSDiskSize string) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check nodeClaim status until it is ready

// get the node object from the nodeClaim status nodeName.

// ensureNodePlugins ensures node plugins are installed.
func (c *RAGEngineReconciler) ensureNodePlugins(ctx context.Context, ragEngineObj *kaitov1beta1.RAGEngine, nodeObj *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// get fresh node object

//Nvidia Plugin

// isNodeClaimCRDAvailable checks if the Karpenter NodeClaim CRD is installed in the cluster
func isNodeClaimCRDAvailable(mgr ctrl.Manager) bool { _ = "STUB: not implemented"; return false }

// Check if karpenter.sh/v1 NodeClaim resource is available

// SetupWithManager sets up the controller with the Manager.
func (c *RAGEngineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Only watch NodeClaim resources if the CRD is actually installed

// watches for nodeClaim with labels indicating RAGEngine name.
func (c *RAGEngineReconciler) watchNodeClaims() handler.EventHandler {
	_ = "STUB: not implemented"
	return *new(handler.EventHandler)
}
