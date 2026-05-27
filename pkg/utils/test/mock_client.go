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

package test

import (
	"context"
	"reflect"

	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClient Client is a mock for the controller-runtime dynamic client interface.
type MockClient struct {
	mock.Mock

	ObjectMap  map[reflect.Type]map[k8sClient.ObjectKey]k8sClient.Object
	StatusMock *MockStatusClient
	UpdateCb   func(key types.NamespacedName)
}

var _ k8sClient.Client = &MockClient{}

func NewClient() *MockClient { _ = "STUB: not implemented"; return nil }

// Retrieves or creates a map associated with the type of obj
func (m *MockClient) ensureMapForType(t reflect.Type) map[k8sClient.ObjectKey]k8sClient.Object {
	_ = "STUB: not implemented"
	return nil
}

//create a new map with the object key if it doesn't exist

func (m *MockClient) CreateMapWithType(t interface{}) map[k8sClient.ObjectKey]k8sClient.Object {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) CreateOrUpdateObjectInMap(obj k8sClient.Object) {
	_ = "STUB: not implemented"
	return
}

func (m *MockClient) GetObjectFromMap(obj k8sClient.Object, key types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

// Get k8s Client interface
func (m *MockClient) Get(ctx context.Context, key types.NamespacedName, obj k8sClient.Object, opts ...k8sClient.GetOption) error {
	_ = "STUB: not implemented"
	//make any necessary changes to the object
	return nil
}

func (m *MockClient) List(ctx context.Context, list k8sClient.ObjectList, opts ...k8sClient.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) getObjectListFromMap(list k8sClient.ObjectList) k8sClient.ObjectList {
	_ = "STUB: not implemented"
	return *new(k8sClient.ObjectList)
}

//add additional object lists as needed

func (m *MockClient) Create(ctx context.Context, obj k8sClient.Object, opts ...k8sClient.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) Delete(ctx context.Context, obj k8sClient.Object, opts ...k8sClient.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) Update(ctx context.Context, obj k8sClient.Object, opts ...k8sClient.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) Patch(ctx context.Context, obj k8sClient.Object, patch k8sClient.Patch, opts ...k8sClient.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) DeleteAllOf(ctx context.Context, obj k8sClient.Object, opts ...k8sClient.DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

// SubResource implements client.Client
func (m *MockClient) SubResource(subResource string) k8sClient.SubResourceClient {
	_ = "STUB: not implemented"
	return *

	// GroupVersionKindFor implements client.Client
	new(k8sClient.SubResourceClient)
}

func (m *MockClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	_ = "STUB: not implemented"
	return *

	// IsObjectNamespaced implements client.Client
	new(schema.GroupVersionKind), nil
}

func (m *MockClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Apply implements client.Client
		nil
}

func (m *MockClient) Apply(ctx context.Context, obj runtime.ApplyConfiguration, opts ...k8sClient.ApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockClient) Scheme() *runtime.Scheme { _ = "STUB: not implemented"; return nil }

func (m *MockClient) RESTMapper() meta.RESTMapper {
	_ = "STUB: not implemented"
	return *new(meta.RESTMapper)
}

// StatusClient interface

func (m *MockClient) Status() k8sClient.StatusWriter {
	_ = "STUB: not implemented"
	return *new(k8sClient.StatusWriter)
}

type MockStatusClient struct {
	mock.Mock
}

func (m *MockStatusClient) Create(ctx context.Context, obj k8sClient.Object, subResource k8sClient.Object, opts ...k8sClient.SubResourceCreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStatusClient) Patch(ctx context.Context, obj k8sClient.Object, patch k8sClient.Patch, opts ...k8sClient.SubResourcePatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStatusClient) Update(ctx context.Context, obj k8sClient.Object, opts ...k8sClient.SubResourceUpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

var _ k8sClient.StatusWriter = &MockStatusClient{}
