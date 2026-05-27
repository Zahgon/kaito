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
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	kaitov1alpha1 "github.com/kaito-project/kaito/api/v1alpha1"
	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
)

var (
	// PollInterval defines the interval time for a poll operation.
	PollInterval = 2 * time.Second
	// PollTimeout defines the time after which the poll operation times out.
	PollTimeout = 120 * time.Second
)

func GetEnv(envVar string) string { _ = "STUB: not implemented"; return "" }

// GenerateRandomString generates a random number between 0 and 1000 and returns it as a string.
func GenerateRandomString() string { _ = "STUB: not implemented"; return "" }

// Generate a random number between 0 and 1000

func GetModelConfigInfo(configFilePath string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetPodNameForWorkspace(coreClient *kubernetes.Clientset, namespace, workspaceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetK8sConfig() (*rest.Config, error) { _ = "STUB: not implemented"; return nil, nil }

// Use kubeconfig file for local development

func GetK8sClientset() (*kubernetes.Clientset, error) { _ = "STUB: not implemented"; return nil, nil }

func GetPodLogs(coreClient *kubernetes.Clientset, namespace, podName, containerName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ExecSync(ctx context.Context, config *rest.Config, coreClient *kubernetes.Clientset, namespace, podName string, options corev1.PodExecOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func PrintPodLogsOnFailure(namespace, labelSelector string) { _ = "STUB: not implemented"; return }

func CopySecret(original *corev1.Secret, targetNamespace string) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func ExtractModelVersion(configs map[string]interface{}) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Using 'tag' as the version

func GenerateInferenceWorkspaceManifest(name, namespace, imageName string, resourceCount int, instanceType string,
	labelSelector *metav1.LabelSelector, preferredNodes []string, presetName kaitov1beta1.ModelName, imagePullSecret []string,
	podTemplate *corev1.PodTemplateSpec, adapters []kaitov1beta1.AdapterSpec, modelAccessSecret, customConfigMapName string) *kaitov1beta1.Workspace {
	_ = "STUB: not implemented"
	return nil
}

// If presetName is not nil, we are using a preset,
// otherwise we are using a custom template

func GenerateInferenceWorkspaceManifestWithVLLM(name, namespace, imageName string, resourceCount int, instanceType string,
	labelSelector *metav1.LabelSelector, preferredNodes []string, presetName kaitov1beta1.ModelName, imagePullSecret []string,
	podTemplate *corev1.PodTemplateSpec, adapters []kaitov1beta1.AdapterSpec, modelAccessSecret, customConfigMapName string) *kaitov1beta1.Workspace {
	_ = "STUB: not implemented"
	return nil
}

func GenerateInferenceSetManifestWithVLLM(name, namespace, imageName string, replicas int, instanceType string,
	labelSelector *metav1.LabelSelector, presetName kaitov1beta1.ModelName, imagePullSecret []string,
	adapters []kaitov1beta1.AdapterSpec, modelAccessSecret string) *kaitov1alpha1.InferenceSet {
	_ = "STUB: not implemented"
	return nil
}

func GenerateInferenceSetManifest(name, namespace, imageName string, replicas int, instanceType string,
	labelSelector *metav1.LabelSelector, presetName kaitov1beta1.ModelName, imagePullSecret []string,
	adapters []kaitov1beta1.AdapterSpec, modelAccessSecret string) *kaitov1alpha1.InferenceSet {
	_ = "STUB: not implemented"
	return nil
}

func GenerateTuningWorkspaceManifest(name, namespace, imageName string, resourceCount int, instanceType string,
	labelSelector *metav1.LabelSelector, preferredNodes []string, input *kaitov1beta1.DataSource,
	output *kaitov1beta1.DataDestination, preset *kaitov1beta1.PresetSpec, method kaitov1beta1.TuningMethod) *kaitov1beta1.Workspace {
	_ = "STUB: not implemented"
	return nil
}

func GenerateE2ETuningWorkspaceManifest(name, namespace, imageName, datasetImageName, outputRegistry string,
	resourceCount int, instanceType string, labelSelector *metav1.LabelSelector,
	preferredNodes []string, presetName kaitov1beta1.ModelName, imagePullSecret []string,
	customConfigMapName string, datasetVolume *corev1.Volume, outputVolume *corev1.Volume) *kaitov1beta1.Workspace {
	_ = "STUB: not implemented"
	return nil
}

// If presetName is not nil, we are using a preset,
// otherwise we are using a custom template

// GenerateE2ETuningConfigMapManifest generates a ConfigMap manifest for E2E tuning.
func GenerateE2ETuningConfigMapManifest(namespace string) *corev1.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}

// Same as workspace namespace

// GenerateE2EInferenceConfigMapManifest generates a ConfigMap manifest for E2E inference.
func GenerateE2EInferenceConfigMapManifest(name, namespace string) *corev1.ConfigMap {
	_ = "STUB: not implemented"
	return nil
}

// Same as workspace namespace

func GeneratePodTemplate(name, namespace, image string, labels map[string]string) *corev1.PodTemplateSpec {
	_ = "STUB: not implemented"
	return nil
}

func CompareSecrets(refs []corev1.LocalObjectReference, secrets []string) bool {
	_ = "STUB: not implemented"
	return false
}
