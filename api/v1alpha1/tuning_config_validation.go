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

package v1alpha1

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
)

type Config struct {
	TrainingConfig TrainingConfig `yaml:"training_config"`
}

type TrainingConfig struct {
	ModelConfig        map[string]runtime.RawExtension `yaml:"ModelConfig"`
	QuantizationConfig map[string]runtime.RawExtension `yaml:"QuantizationConfig"`
	LoraConfig         map[string]runtime.RawExtension `yaml:"LoraConfig"`
	TrainingArguments  map[string]runtime.RawExtension `yaml:"TrainingArguments"`
	DatasetConfig      map[string]runtime.RawExtension `yaml:"DatasetConfig"`
	DataCollator       map[string]runtime.RawExtension `yaml:"DataCollator"`
}

func validateNilOrBool(value interface{}) error { _ = "STUB: not implemented"; return nil }

// nil is acceptable

// Correct type

// UnmarshalYAML custom method
func (t *TrainingConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

// This function converts a map[string]interface{} to a map[string]runtime.RawExtension.
// It does this by setting the raw marshalled data of the unmarshalled YAML to
// be the raw data of the runtime.RawExtension object.

func UnmarshalTrainingConfig(cm *corev1.ConfigMap) (*Config, *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateTrainingArgsViaConfigMap(cm *corev1.ConfigMap) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// If specified, ensure output dir is of type string

// Ensure the user-specified directory is under baseDir

// TODO: Here we perform the tuning GPU Memory Checks!

func validateMethodViaConfigMap(cm *corev1.ConfigMap, methodLowerCase string) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// Validate QuantizationConfig if it exists

// Dynamic field search for quantization settings within ModelConfig

// Validate both loadIn4bit and loadIn8bit

// Validation Logic

// getStructInstances dynamically generates instances of all sections in any config struct.
func getStructInstances(s any) map[string]any { _ = "STUB: not implemented"; return nil }

// Dereference pointer to get the struct type

// Create a new instance of the type pointed to by the field

func validateTuningConfigMapSchema(cm *corev1.ConfigMap) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// Extract the actual training configuration map

// Check if valid sections

func (r *TuningSpec) validateConfigMap(ctx context.Context, namespace string, methodLowerCase string, configMapName string) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}
