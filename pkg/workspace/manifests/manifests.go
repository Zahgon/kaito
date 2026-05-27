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

package manifests

import (
	"context"

	helmv2 "github.com/fluxcd/helm-controller/api/v2"
	sourcev1 "github.com/fluxcd/source-controller/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"

	kaitov1alpha1 "github.com/kaito-project/kaito/api/v1alpha1"
	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	pkgmodel "github.com/kaito-project/kaito/pkg/model"
	"github.com/kaito-project/kaito/pkg/utils/generator"
)

func GenerateHeadlessServiceManifest(workspaceObj *kaitov1beta1.Workspace) *corev1.Service {
	_ = "STUB: not implemented"
	return nil
}

func GenerateServiceManifest(workspaceObj *kaitov1beta1.Workspace, serviceType corev1.ServiceType) *corev1.Service {
	_ = "STUB: not implemented"
	return nil
}

// select the pod with index 0 as the endpoint

// When the routing sidecar is present (decode role + vLLM), route
// external traffic through the sidecar port so that all requests
// pass through the routing layer. Kubelet container probes still
// hit vLLM directly on PortInferenceServer (via PodIP), while
// Service/Gateway traffic routes to the sidecar on PortRoutingSidecar.

// HTTP API Port

// Added this to allow pods to discover each other
// (DNS Resolution) During their initialization phase

func GenerateStatefulSetManifest(revisionNum string, replicas int) func(*generator.WorkspaceGeneratorContext, *appsv1.StatefulSet) error {
	_ = "STUB: not implemented"
	return nil
}

// if workspaceObj.Labels contains "inferenceset.kaito.sh/created-by", add it to selector for VPA/HPA purpose

// Propagate MRI parent and inference-role labels to pod templates for InferencePool endpoint selection.

func AddStatefulSetVolumeClaimTemplates(volumeClaimTemplates corev1.PersistentVolumeClaim) func(*generator.WorkspaceGeneratorContext, *appsv1.StatefulSet) error {
	_ = "STUB: not implemented"
	return nil
}

func SetStatefulSetPodSpec(podSpec *corev1.PodSpec) func(*generator.WorkspaceGeneratorContext, *appsv1.StatefulSet) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateTuningJobManifest(revisionNum string) func(*generator.WorkspaceGeneratorContext, *batchv1.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func SetJobPodSpec(podSpec *corev1.PodSpec) func(*generator.WorkspaceGeneratorContext, *batchv1.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func GeneratePullerContainers(wObj *kaitov1beta1.Workspace, adapters []kaitov1beta1.AdapterSpec, volumeMounts []corev1.VolumeMount) ([]corev1.Container, []corev1.EnvVar, []corev1.Volume) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GenerateManifestWithPodTemplate(workspaceObj *kaitov1beta1.Workspace, tolerations []corev1.Toleration) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

// if workspaceObj.Labels contains "inferenceset.kaito.sh/created-by", add it to selector for VPA/HPA purpose

// Propagate MRI parent and inference-role labels to pod templates for InferencePool endpoint selection.

// Overwrite affinity. Only set node affinity when there are user-defined
// node requirements; an empty MatchExpressions list is rejected by the
// Kubernetes API server.

// append tolerations

func GetModelImageName(presetObj *pkgmodel.PresetParam) string {
	_ = "STUB: not implemented"
	return ""
}

// GenerateModelPullerContainer creates an init container that pulls model images using ORAS
func GenerateModelPullerContainer(ctx context.Context, workspaceObj *kaitov1beta1.Workspace, presetObj *pkgmodel.PresetParam) []corev1.Container {
	_ = "STUB: not implemented"
	return nil
}

// If the preset is set to download at runtime, we don't need to pull the model weights.

// GenerateInferencePoolOCIRepository generates a Flux OCIRepository for the inference pool.
func GenerateInferencePoolOCIRepository(inferenceSetObj *kaitov1alpha1.InferenceSet) *sourcev1.OCIRepository {
	_ = "STUB: not implemented"
	return nil
}

// Chart source for Gateway API Inference Extension inference pool;
// keep in sync with consts.InferencePoolChartVersion when upgrading.

// inferencePoolTargetPort returns the target port for the InferencePool.
// For decode-role InferenceSets with vLLM runtime, traffic goes through the
// routing sidecar on PortRoutingSidecar. For all other cases, traffic goes
// directly to the inference server on PortInferenceServer.
// Runtime detection mirrors v1beta1.GetWorkspaceRuntimeName: when the vLLM
// feature gate is enabled (default), the runtime defaults to vLLM unless
// explicitly overridden by the kaito.sh/runtime annotation.
func inferencePoolTargetPort(inferenceSetObj *kaitov1alpha1.InferenceSet) int32 {
	_ = "STUB: not implemented"
	return 0
}

// Mirror GetWorkspaceRuntimeName logic: default to vLLM when feature gate is on.

// GenerateInferencePoolHelmRelease generates a Flux HelmRelease for the inference pool.
func GenerateInferencePoolHelmRelease(inferenceSetObj *kaitov1alpha1.InferenceSet) (*helmv2.HelmRelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The Endpoint Picker (EPP) from Gateway API Inference Extension picks an endpoint that can serve traffic.
// KAITO overrides the default GWIE EPP image with the llm-d inference scheduler, which provides
// advanced scheduling plugins (KV cache-aware routing, P/D disaggregation, pluggable filters/scorers).
// In a multi-node inference environment, this means we need to select the leader pod (with pod index 0)
// since only the leader pod is capable of serving traffic.

// Based on https://github.com/kubernetes-sigs/gateway-api-inference-extension/blob/v1.3.1/config/charts/inferencepool/values.yaml

// Referencing the OCIRepository created above
