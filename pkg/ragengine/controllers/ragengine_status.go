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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

func (c *RAGEngineReconciler) updateRAGEngineStatus(ctx context.Context, name *client.ObjectKey, condition *metav1.Condition, workerNodes []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the latest version to avoid update conflict.

func (c *RAGEngineReconciler) updateStatusConditionIfNotMatch(ctx context.Context, ragObj *kaitov1beta1.RAGEngine, cType kaitov1beta1.ConditionType,
	cStatus metav1.ConditionStatus, cReason, cMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to change

func (c *RAGEngineReconciler) updateStatusNodeListIfNotMatch(ctx context.Context, ragObj *kaitov1beta1.RAGEngine, validNodeList []*corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}
