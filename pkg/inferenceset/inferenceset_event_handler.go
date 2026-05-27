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

package inferenceset

import (
	"context"

	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/utils"
)

type workspaceEventHandler struct {
	enqueueHandler handler.TypedEventHandler[client.Object, reconcile.Request]
	expectations   *utils.ControllerExpectations
	logger         klog.Logger
}

var _ handler.TypedEventHandler[client.Object, reconcile.Request] = (*workspaceEventHandler)(nil)

func getControllerKeyForWorkspace(ws *kaitov1beta1.Workspace) *client.ObjectKey {
	_ = "STUB: not implemented"
	return nil
}

func (n *workspaceEventHandler) Create(ctx context.Context, evt event.TypedCreateEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (n *workspaceEventHandler) Delete(ctx context.Context, evt event.TypedDeleteEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (n *workspaceEventHandler) Generic(ctx context.Context, evt event.TypedGenericEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (n *workspaceEventHandler) Update(ctx context.Context, evt event.TypedUpdateEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

var enqueueInferenceSetForWorkspace = handler.EnqueueRequestsFromMapFunc(
	func(ctx context.Context, o client.Object) []reconcile.Request {
		workspaceObj := o.(*kaitov1beta1.Workspace)
		key := getControllerKeyForWorkspace(workspaceObj)
		if key == nil {
			return nil
		}
		return []reconcile.Request{
			{
				NamespacedName: *key,
			},
		}
	})
