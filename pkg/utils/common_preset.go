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
	corev1 "k8s.io/api/core/v1"
)

const (
	DefaultVolumeMountPath    = "/dev/shm"
	DefaultConfigMapMountPath = "/mnt/config"
	DefaultDataVolumePath     = "/mnt/data"
	DefaultAdapterVolumePath  = "/mnt/adapter"
	DefaultWeightsVolumePath  = "/workspace/weights"

	DefaultORASToolImage = "mcr.microsoft.com/oss/v2/oras-project/oras:v1.2.3"
)

var DefaultModelWeightsVolume = corev1.Volume{
	Name: "model-weights-volume",
	VolumeSource: corev1.VolumeSource{
		EmptyDir: &corev1.EmptyDirVolumeSource{},
	},
}
var DefaultModelWeightsVolumeMount = corev1.VolumeMount{
	Name:      "model-weights-volume",
	MountPath: DefaultWeightsVolumePath,
}

func ConfigResultsVolume(outputPath string, outputVolume *corev1.VolumeSource) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func FindResultsVolumeMount(spec *corev1.PodSpec) *corev1.VolumeMount {
	_ = "STUB: not implemented"
	return nil
}

func ConfigImagePullSecretVolume(nameSuffix string, imagePullSecrets []string) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func ConfigImagePushSecretVolume(imagePushSecret string) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func ConfigSHMVolume() (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func ConfigCMVolume(cmName string) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func ConfigDataVolume(inputVolumeSource *corev1.VolumeSource) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func ConfigAdapterVolume(inputVolumeSource *corev1.VolumeSource) (corev1.Volume, corev1.VolumeMount) {
	_ = "STUB: not implemented"
	return *new(corev1.Volume), *new(corev1.VolumeMount)
}

func GetPresetImageName(registry, name, tag string) string { _ = "STUB: not implemented"; return "" }
