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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

const (
	GuardrailsPolicyVolumeName = "guardrails-policy"
	GuardrailsPolicyMountPath  = "/etc/ragengine/guardrails"
	GuardrailsPolicyFileName   = kaitov1beta1.GuardrailsPolicyFileName
	GuardrailsPolicyFilePath   = GuardrailsPolicyMountPath + "/" + GuardrailsPolicyFileName
)

func GenerateRAGDeploymentManifest(ragEngineObj *kaitov1beta1.RAGEngine, revisionNum string, imageName string,
	imagePullSecretRefs []corev1.LocalObjectReference, commands []string, containerPorts []corev1.ContainerPort,
	livenessProbe, readinessProbe *corev1.Probe, resourceRequirements corev1.ResourceRequirements,
	tolerations []corev1.Toleration, volumes []corev1.Volume, volumeMount []corev1.VolumeMount) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

// we only set node affinity if there are node requirements specified. If there are no requirements, we don't set affinity at all to allow scheduling on any node.

// RAGEngine requires exactly 1 replica

// Configuration for rolling updates: allows no extra pods during the update and permits at most one unavailable pod at a time。

func RAGSetEnv(ragEngineObj *kaitov1beta1.RAGEngine) []corev1.EnvVar {
	_ = "STUB: not implemented"
	return nil

	// Add Pod metadata as environment variables for lifecycle hooks
}

// TODO: Model ID Env

// Determine vector DB type from CRD spec or default to "faiss"

// Inject vector DB connection info if configured

// Set the vector database persist directory based on storage configuration
// default in-memory/ephemeral storage

// Append RAGEngine name to ensure unique directory per instance

// Only add LLM_INFERENCE_URL if URL is not empty (URL is optional)

func GenerateRAGServiceManifest(ragObj *kaitov1beta1.RAGEngine, serviceName string, serviceType corev1.ServiceType) *corev1.Service {
	_ = "STUB: not implemented"
	return nil
}
