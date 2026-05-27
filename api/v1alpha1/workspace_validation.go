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

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	"knative.dev/pkg/apis"

	"github.com/kaito-project/kaito/pkg/model"
)

const (
	N_SERIES_PREFIX = "Standard_N"
	D_SERIES_PREFIX = "Standard_D"

	DefaultLoraConfigMapTemplate   = "lora-params-template"
	DefaultQloraConfigMapTemplate  = "qlora-params-template"
	DefaultInferenceConfigTemplate = "inference-params-template"
	MaxAdaptersNumber              = 10
)

func (w *Workspace) SupportedVerbs() []admissionregistrationv1.OperationType {
	_ = "STUB: not implemented"
	return nil
}

func (w *Workspace) Validate(ctx context.Context) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Check if the bypass resource checks annotation is set

// TODO: Add Adapter Spec Validation - Including DataSource Validation for Adapter

// TODO: Add validate resource based on Tuning Spec

func (w *Workspace) validateCreate() (errs *apis.FieldError) { _ = "STUB: not implemented"; return nil }

func (w *Workspace) validateNodeImageFamilyAnnotation() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Workspace) validateUpdate(old *Workspace) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (r *AdapterSpec) validateCreateorUpdate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Adapters support Image or Volume as source (not URLs)

func (r *TuningSpec) validateCreate(ctx context.Context, workspaceNamespace string) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Currently require a preset to specified, in future we can consider defining a template

func (r *TuningSpec) validateUpdate(old *TuningSpec) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Consider supporting config fields changing

func (r *DataSource) validateCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Ensure exactly one of URLs, Volume, or Image is specified

func (r *DataSource) validateUpdate(old *DataSource, isTuning bool) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataDestination) validateCreate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Cloud Provider requires credentials to push image

// Ensure exactly one of Volume or Image is specified

// TODO: Consider allowing both Volume and Image to be specified

func (r *DataDestination) validateUpdate() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResourceSpec) validateCreateWithTuning(tuning *TuningSpec) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResourceSpec) validateCreateWithInference(inference *InferenceSpec, bypassResourceChecks bool) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Since inference.Preset exists, we must validate preset name.

// If the preset is not valid, check if it is a deprecated model
// We use recover() to handle the panic from MustGet if the model is not found

// Return to skip the rest of checks, the Inference spec validation will return proper err msg.

// Check if instancetype exists in our SKUs map for the particular cloud provider

// InferenceSpec has been validated so the name is valid.

// Total GPU memory

// Check for other instance types pattern matches if cloud provider is Azure

// Validate labelSelector

func (r *ResourceSpec) validateUpdate(old *ResourceSpec) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// We disable changing node count for now.
	return nil
}

func (i *InferenceSpec) validateCreate(ctx context.Context, runtime model.RuntimeName) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Check if both Preset and Template are not set
	return nil
}

// Check if both Preset and Template are set at the same time

// Validate preset name

// Need to return here. Otherwise, a panic will be hit when doing following checks.

// Validate private preset has private image specified

// Additional validations for Preset

// Note: we don't enforce private access mode to have image secrets, in case anonymous pulling is enabled

// check if adapter names are duplicate

func (i *InferenceSpec) validateUpdate(old *InferenceSpec) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// inference.template can be changed, but cannot be set/unset.

// check if adapter names are duplicate

// check if adapter names are duplicate

func validateDuplicateName(adapters []AdapterSpec, nameMap map[string]bool) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}
