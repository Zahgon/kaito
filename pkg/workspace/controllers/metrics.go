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

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

var (
	workspacePhaseCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kaito_workspace_count",
			Help: "Number of Workspaces in a certain phase (succeeded, error, pending, deleting)",
		},
		[]string{"phase"},
	)

	workspacePresetCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kaito_workspace_preset_count",
			Help: "Number of Workspaces using each preset model, by preset name",
		},
		[]string{"preset_name"},
	)

	workspacePVCAllocatedBytes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kaito_workspace_pvc_allocated_bytes",
			Help: "Allocated (requested) PVC storage in bytes per PVC associated with a workspace",
		},
		[]string{"workspace_name", "workspace_namespace", "pvc_name"},
	)

	workspacePVCCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kaito_workspace_pvc_count",
			Help: "Number of PVCs associated with each workspace",
		},
		[]string{"workspace_name", "workspace_namespace"},
	)
)

func init() {
	metrics.Registry.MustRegister(workspacePhaseCount)
	metrics.Registry.MustRegister(workspacePresetCount)
	metrics.Registry.MustRegister(workspacePVCAllocatedBytes)
	metrics.Registry.MustRegister(workspacePVCCount)
}

func monitorWorkspaces(ctx context.Context, k8sClient client.Client) {
	_ = "STUB: not implemented"
	return
}

// Reset before re-setting so to remove stale keys

func collectPVCMetrics(ctx context.Context, k8sClient client.Client) {
	_ = "STUB: not implemented"
	return
}

// namespace -> workspace name -> count

// Prefer actual allocated capacity from status; fall back to spec request

func getWorkspacePresetName(ws *kaitov1beta1.Workspace) string {
	_ = "STUB: not implemented"
	return ""
}

func DetermineWorkspacePhase(ws *kaitov1beta1.Workspace) string {
	_ = "STUB: not implemented"
	return ""
}
