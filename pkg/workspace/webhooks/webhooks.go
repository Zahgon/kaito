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

package webhooks

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	knativeinjection "knative.dev/pkg/injection"
	"knative.dev/pkg/webhook/resourcesemantics"

	kaitov1alpha1 "github.com/kaito-project/kaito/api/v1alpha1"
	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

func NewControllerWebhooks() []knativeinjection.ControllerConstructor {
	_ = "STUB: not implemented"
	return nil
}

func NewWorkspaceCRDValidationWebhook(ctx context.Context, _ configmap.Watcher) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceSetCRDValidationWebhook(ctx context.Context, _ configmap.Watcher) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

func NewMultiRoleInferenceCRDValidationWebhook(ctx context.Context, _ configmap.Watcher) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

var WorkspaceResources = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
	kaitov1alpha1.GroupVersion.WithKind("Workspace"): &kaitov1alpha1.Workspace{},
	kaitov1beta1.GroupVersion.WithKind("Workspace"):  &kaitov1beta1.Workspace{},
}

var InferenceSetResources = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
	kaitov1alpha1.GroupVersion.WithKind("InferenceSet"): &kaitov1alpha1.InferenceSet{},
}

var MultiRoleInferenceResources = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
	kaitov1alpha1.GroupVersion.WithKind("MultiRoleInference"): &kaitov1alpha1.MultiRoleInference{},
}
