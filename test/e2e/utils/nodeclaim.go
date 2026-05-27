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

package utils

import (
	"context"

	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	"github.com/kaito-project/kaito/api/v1beta1"
)

// ValidateNodeClaimCreation Logic to validate the nodeClaim creation.
func ValidateNodeClaimCreation(ctx context.Context, workspaceObj *v1beta1.Workspace, expectedCount int) {
	_ = "STUB: not implemented"
	return
}

// GetAllValidNodeClaims get all valid nodeClaims.
func GetAllValidNodeClaims(ctx context.Context, workspaceObj *v1beta1.Workspace) (*karpenterv1.NodeClaimList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
