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
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kaito-project/kaito/api/v1beta1"
	"github.com/kaito-project/kaito/pkg/utils/consts"
	"github.com/kaito-project/kaito/pkg/utils/resources"
)

const (
	ProbePath           = "/health"
	PortInferenceServer = 5000
)

var (
	containerPorts = []corev1.ContainerPort{{
		ContainerPort: int32(PortInferenceServer),
	},
	}

	livenessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Port: intstr.FromInt(PortInferenceServer),
				Path: ProbePath,
			},
		},
		InitialDelaySeconds: 600, // 10 minutes
		PeriodSeconds:       10,
	}

	readinessProbe = &corev1.Probe{
		ProbeHandler: corev1.ProbeHandler{
			HTTPGet: &corev1.HTTPGetAction{
				Port: intstr.FromInt(PortInferenceServer),
				Path: ProbePath,
			},
		},
		InitialDelaySeconds: 30,
		PeriodSeconds:       10,
	}

	tolerations = []corev1.Toleration{
		{
			Effect:   corev1.TaintEffectNoSchedule,
			Operator: corev1.TolerationOpExists,
			Key:      resources.CapacityNvidiaGPU,
		},
		{
			Effect:   corev1.TaintEffectNoSchedule,
			Value:    consts.GPUString,
			Key:      consts.SKUString,
			Operator: corev1.TolerationOpEqual,
		},
	}
)

type ImageConfig struct {
	RegistryName string
	ImageName    string
	ImageTag     string
}

func (ic ImageConfig) GetImage() string { _ = "STUB: not implemented"; return "" }

func getImageConfig() ImageConfig { _ = "STUB: not implemented"; return *new(ImageConfig) }

func getEnv(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }

// configStorageVolume creates a volume and volume mount for vector database storage
func configStorageVolume(storageSpec *v1beta1.StorageSpec) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

// Use PVC for persistent storage

// Use emptyDir as fallback (data will not persist across pod restarts)

func configGuardrailsPolicyVolume(cmName string) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

// The ConfigMap data key and the mounted filename are intentionally the same.

func ensureGuardrailsPolicyConfigMap(ctx context.Context, ragEngineObj *v1beta1.RAGEngine, kubeClient client.Client) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// User-provided ConfigMaps belong to the user; never patch them.
// The copied default ConfigMap is shared namespace-wide, so it must stay
// unowned by any individual RAGEngine to avoid breaking other workloads.

func CreatePresetRAG(ctx context.Context, ragEngineObj *v1beta1.RAGEngine, revisionNum string, kubeClient client.Client) (client.Object, error) {
	_ = "STUB: not implemented"
	return *new(client.Object), nil
}

// Configure storage volume for FAISS vector database persistence

// If GetGPUConfigBySKU returns error, skip GPU resource allocation (e.g., CPU-only instances)

// If embedding is remote or compute instance type is not specified, do not allocate GPU resources by default
// and apply default CPU and memory requests to ensure the pod can be scheduled.
