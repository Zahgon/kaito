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

package tuning

import (
	"context"
	_ "embed"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kaitov1beta1 "github.com/kaito-project/kaito/api/v1beta1"
	pkgmodel "github.com/kaito-project/kaito/pkg/model"
	"github.com/kaito-project/kaito/pkg/utils/consts"
	"github.com/kaito-project/kaito/pkg/utils/generator"
)

const (
	DefaultBaseDir          = "/mnt"
	DefaultOutputVolumePath = "/mnt/output"
)

var (
	//go:embed scripts/data-downloader.sh
	dataDownloaderScript string

	containerPorts = []corev1.ContainerPort{{
		ContainerPort: consts.PortInferenceServer,
	}}

	// Come up with valid liveness and readiness probes for fine-tuning
	// TODO: livenessProbe = &corev1.Probe{}
	// TODO: readinessProbe = &corev1.Probe{}
)

func defaultTolerations() []corev1.Toleration { _ = "STUB: not implemented"; return nil }

func GetTuningImageInfo() string { _ = "STUB: not implemented"; return "" }

// PrepareOutputDir ensures the output directory is within the base directory.
func PrepareOutputDir(outputDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetOutputDirFromTrainingArgs retrieves the output directory from training arguments if specified.
func GetOutputDirFromTrainingArgs(trainingArgs map[string]runtime.RawExtension) (string, *apis.FieldError) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetTrainingOutputDir retrieves and validates the output directory from the ConfigMap.
func GetTrainingOutputDir(ctx context.Context, configMap *corev1.ConfigMap) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CreatePresetTuning(ctx context.Context, workspaceObj *kaitov1beta1.Workspace, revisionNum string,
	model pkgmodel.Model, kubeClient client.Client) (client.Object, error) {
	_ = "STUB: not implemented"
	return *new(client.Object), nil
}

func GenerateBasicTuningPodSpec(skuNumGPUs int) func(*generator.WorkspaceGeneratorContext, *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// additional volume

// add share memory for cross process communication

// Add volume for model weights access

// resource requirements

// default env

// Append environment variable for default target modules if using Phi3 model

// Add Expandable Memory Feature to reduce Peak GPU Mem Usage

// tuning commands

// Add node affinity based on label selector from workspace resource

// Only set nodeAffinity when the user supplied selector labels.
// An empty MatchExpressions list is rejected by the Kubernetes API server.

func SetTrainingResultVolume(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Add shared volume for tuning parameters

// Add results volume for training output

// Now there are two options for data destination 1. Volume - 2. Image
// notes: this modifier requires the results volume to be set in the pod spec
func SetTrainingOutputImagePush(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// additional secret volume for image push

// additional sidecar container for uploading image

func SetTrainingInput(ctx *generator.WorkspaceGeneratorContext, spec *corev1.PodSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Now there are three options for DataSource: 1. URL - 2. Volume - 3. Image
func prepareDataSource(ctx context.Context, workspaceObj *kaitov1beta1.Workspace) (*corev1.Container, []corev1.Volume, []corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func handleURLDataSource(workspaceObj *kaitov1beta1.Workspace) (*corev1.Container, corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return nil, *new(corev1.Volume), *new(corev1.VolumeMount)
}
