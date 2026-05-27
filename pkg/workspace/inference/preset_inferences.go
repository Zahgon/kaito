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

package inference

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kaito-project/kaito/api/v1beta1"
	pkgmodel "github.com/kaito-project/kaito/pkg/model"
	"github.com/kaito-project/kaito/pkg/sku"
	"github.com/kaito-project/kaito/pkg/utils/consts"
	"github.com/kaito-project/kaito/pkg/utils/generator"
)

const (
	ProbePath = "/health"

	// defaultStartupProbeTimeout is the startup probe timeout for models that do not
	// specify ReadinessTimeout. 30 minutes covers all current models.
	defaultStartupProbeTimeout = 30 * time.Minute
)

var (
	containerPorts = []corev1.ContainerPort{{
		ContainerPort: int32(consts.PortInferenceServer),
	}}

	// defaultLivenessProbe has no initial delay because the startup probe ensures
	// the model is up before liveness evaluation begins.
	defaultLivenessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Port: intstr.FromInt32(consts.PortInferenceServer),
				Path: ProbePath,
			},
		},
		InitialDelaySeconds: 0,
		PeriodSeconds:       10,
		FailureThreshold:    3,
	}

	defaultReadinessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Port: intstr.FromInt32(consts.PortInferenceServer),
				Path: ProbePath,
			},
		},
		InitialDelaySeconds: 30,
		PeriodSeconds:       10,
	}
)

func defaultTolerations(ws *v1beta1.Workspace) []corev1.Toleration {
	_ = "STUB: not implemented"
	return nil
}

func GetInferenceImageInfo(ctx context.Context, workspaceObj *v1beta1.Workspace) []corev1.LocalObjectReference {
	_ = "STUB: not implemented"
	return nil
}

// Check if the workspace preset's access mode is private

// GenerateModelFileCacheVolume generates a volume for caching model files.
// These files would be stored in the local pv and its lifetime is tied to the pod.
// Use NVMe for storage acceleration if it's available.
//
// notes: no capacity check here because NVMe is typically a TiB level storage,
// which is sufficient for almost all models. check it if this assumption is not true.
func GenerateModelWeightsCacheVolume(ctx context.Context, workspaceObj *v1beta1.Workspace, model pkgmodel.Model) corev1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return *new(corev1.PersistentVolumeClaim)
}

// place model files in this volume

func GeneratePresetInference(ctx context.Context, workspaceObj *v1beta1.Workspace, revisionNum string,
	model pkgmodel.Model, kubeClient client.Client) (client.Object, error) {
	_ = "STUB: not implemented"
	return *new(client.Object), nil
}

// Set the target node count for the inference workload

// Use StatefulSet for all use cases to ensure consistent pod identity and storage management
// For multi-node distributed inference with vLLM, we need StatefulSet to ensure pods are
// created with individual identities (their ordinal indexes) -
// https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#pod-identity

func getGPUConfig(ctx *generator.WorkspaceGeneratorContext) (*sku.GPUConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NAP is disabled (BYO scenario) - prefer to get GPU config from matching nodes with nvidia.com labels
// Only try to find matching nodes if we have a labelSelector and if WorkerNodes is not already populated

// NAP is enabled - try to get GPU config from known SKU

func shouldUseDistributedInference(ctx *generator.WorkspaceGeneratorContext, numNodes int) bool {
	_ = "STUB: not implemented"
	return false
}

type probeType string

const (
	probeTypeLiveness  probeType = "liveness"
	probeTypeReadiness probeType = "readiness"
)

func checkIfNVMeAvailable(ctx context.Context, gpuConfig *sku.GPUConfig, kubeClient client.Client) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if the required NVMe storage class exists

// getDistributedInferenceProbe returns a container probe configuration for the distributed inference workload.
func getDistributedInferenceProbe(probeType probeType, wObj *v1beta1.Workspace, initialDelaySeconds, periodSeconds, timeoutSeconds, failureThreshold int32) *corev1.Probe {
	_ = "STUB: not implemented"
	return nil
}

// for distributed inference, we cannot use the default http probe since only the leader pod
// exposes the health check endpoint. We need to use presets/workspace/inference/vllm/multi-node-health-check.py
// to check the health of both the leader and worker pods.

func buildStartupProbe(timeout time.Duration) *corev1.Probe { _ = "STUB: not implemented"; return nil }

// ceil(timeout / period) ensures the full timeout window is covered.

func buildDistributedStartupProbe(timeout time.Duration, wObj *v1beta1.Workspace) *corev1.Probe {
	_ = "STUB: not implemented"
	return nil
}

// buildBenchmarkStartupProbe returns an exec startup probe that runs
// benchmark_entrypoint.py on every kubelet tick.
//
// While vLLM is loading, the script exits 1 (/health not yet up), consuming the
// failureThreshold budget.  Once /health passes, the script runs the full benchmark
// and drain phase, then exits 0 — which marks the startup probe as passed and
// activates the readiness probe.
//
// When wObj is non-nil the probe is built for distributed inference: a shell
// conditional routes the leader (POD_INDEX=0) to benchmark_entrypoint.py and
// workers to the standard multi-node health check.
//
// timeoutSeconds is set to 600 to prevent kubelet killing the process mid-benchmark.
func buildBenchmarkStartupProbe(timeout time.Duration, wObj *v1beta1.Workspace, distributed bool) *corev1.Probe {
	_ = "STUB: not implemented"
	return nil
}

// covers benchmark duration + drain + buffer

func GetBaseImageName() string { _ = "STUB: not implemented"; return "" }

func GenerateInferencePodSpec(gpuConfig *sku.GPUConfig, numNodes int) func(*generator.WorkspaceGeneratorContext, *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// debug print of configVolume (requested)

// additional volume

// Add config volume mount

// add model weights volume mount

// add share memory for cross process communication

// node selector

// resource requirements

// inference command

// Calculate max-model-len for runtime context
// Default value

// First check if user provided explicit value in ConfigMap
//if v, ok2 := utils.ParseExplicitMaxModelLen(raw); ok2 {
//maxModelLen = v
//klog.Infof("[RuntimeContext] workspace=%s using user explicit max-model-len=%d", ctx.Workspace.Name, maxModelLen)
//} else {
// If no user value, compute planned value

//}

// When the routing sidecar is needed, it will be injected after the
// main container is created. vLLM keeps its default port (5000).

// Only set nodeAffinity when the user supplied selector labels.
// An empty MatchExpressions list is rejected by the Kubernetes API server.

// Use the model's ReadinessTimeout if specified; otherwise fall back to the
// default. containerStatuses[].started is reliable for downstream.

func SetModelDownloadInfo(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// add HF_TOKEN env var to the main inference container only

// additional initContainers

func SetAdapterPuller(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Find the main inference container by workspace name.

// Separate adapters by source type

// Handle image-based adapters (existing flow: EmptyDir + puller init containers)

// add container to pull adapters

// Handle volume-based adapters (mount volume directly, no puller needed)

// Propagate strength env vars for volume adapters

// SetBenchmarkConfig overrides the startup probe to run the benchmark entrypoint.
// It must be appended after GenerateInferencePodSpec (and SetDistributedInferenceProbe
// when distributed) so the container already exists.
func SetBenchmarkConfig(distributed bool) generator.TypedManifestModifier[generator.WorkspaceGeneratorContext, corev1.PodSpec] {
	_ = "STUB: not implemented"
	return nil
}

func SetDistributedInferenceProbe(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// 60 seconds initial delay for liveness probe to allow workers to join the cluster

func SetDefaultModelWeightsVolume(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// applyInferenceRoleEnv sets KAITO_INFERENCE_ROLE env var on the main inference
// container (identified by containerName) when the workspace has a valid
// inference-role label (prefill or decode). Only the main container needs this
// env var; sidecar containers do not use it.
func applyInferenceRoleEnv(labels map[string]string, containerName string, spec *corev1.PodSpec) {
	_ = "STUB: not implemented"
	return
}

// injectRoutingSidecar appends the llm-d routing sidecar container to the pod
// spec. The sidecar listens on PortRoutingSidecar (5001) and proxies to the
// main vLLM container which keeps its default PortInferenceServer (5000).
// No port or probe rewriting is needed on the main container.
func injectRoutingSidecar(spec *corev1.PodSpec) { _ = "STUB: not implemented"; return }

// needsRoutingSidecar returns true if the workspace requires the llm-d routing sidecar.
func needsRoutingSidecar(ws *v1beta1.Workspace) bool { _ = "STUB: not implemented"; return false }
