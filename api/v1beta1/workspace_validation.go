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

package v1beta1

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

func (w *Workspace) validateAnnotations() (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// valid

func (w *Workspace) validateCreate() (errs *apis.FieldError) { _ = "STUB: not implemented"; return nil }

// Check node auto-provisioning feature gate and validate instanceType accordingly
// This validation only applies to CREATE operations, not UPDATE (since instanceType is immutable)

// When NAP is disabled, instanceType must be empty (BYO scenario)

// When NAP is enabled, instanceType must be specified for node provisioning

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
	// If old is nil, this means Tuning is being toggled on, which should be caught by validateUpdate in Workspace
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

func (r *ResourceSpec) validateCreateWithInference(ctx context.Context, inference *InferenceSpec, bypassResourceChecks bool, runtime model.RuntimeName, wsNamespace string) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}

// Since inference.Preset exists, we must validate preset name.

// If the preset is not valid, check if it is a deprecated model
// We use recover() to handle the panic from MustGet if the model is not found

// Return to skip the rest of checks, the Inference spec validation will return proper err msg.

// Validate labelSelector

// Warn (don't reject) when the user-provided selector includes labels
// reserved for KAITO-managed resources. These keys are silently ignored
// at runtime to avoid cross-workspace/RAGEngine targeting.

// Reject when, after stripping KAITO-reserved keys, no usable matchLabels
// remain from a user-supplied set. An all-reserved selector would otherwise
// be silently dropped and end up matching every node in the cluster.

// If the user is using a custom pod template instead of a preset, we don't need to list the BYO nodes to get GPU info as we don't know the GPU requirements of a custom model.
// Note: for tests like aikit.yaml, it creates nodes with kind that do not have GPU labels, so we need to account for that case.

// List matching nodes (KAITO-reserved label keys are stripped to avoid
// matching nodes that belong to other Workspaces or RAGEngines).

// Try to get GPU configuration from nvidia.com labels first

// Verify uniformity

// NAP enabled
// Regardless of if preset is empty or not, we do want to make sure the instance type is valid for NAP and can't skip node validation like BYO.

// Check for other instance types pattern matches if cloud provider is Azure

// InferenceSpec has been validated so the name is valid.

// Total GPU memory

// GPU memory check and distributed inference runtime check: only run if TotalSafeTensorFileSize is specified

// If the model preset supports distributed inference, and a single machine has insufficient GPU memory to run the model,
// then we need to make sure the Workspace is not using the Huggingface Transformers runtime since it no longer supports
// multi-node distributed inference.

func (r *ResourceSpec) validateUpdate(old *ResourceSpec) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// We disable changing node count for now.
	return nil
}

// Check node auto-provisioning feature gate and validate instanceType accordingly

// When NAP is disabled, instanceType must be empty (BYO scenario)

// for backward compatibility, old.InstanceType is non-empty
// but update to empty is allowed.

func (i *InferenceSpec) validateCreate(ctx context.Context, runtime model.RuntimeName, wsNamespace string) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// Check if both Preset and Template are not set
	return nil
}

// Check if both Preset and Template are set at the same time

// Validate preset name

// Need to return here. Otherwise, a panic will be hit when doing following checks.

// For models that require downloading at runtime, we need to check if the modelAccessSecret is provided

// check if adapter names are duplicate

func (i *InferenceSpec) validateUpdate(old *InferenceSpec) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	// If old is nil, this means Inference is being toggled on, which should be caught by validateUpdate in Workspace
	return nil
}

// inference.template can be changed, but cannot be set/unset.

// check if adapter names are duplicate

// check if adapter names are duplicate

func validateDuplicateName(adapters []AdapterSpec, nameMap map[string]bool) (errs *apis.FieldError) {
	_ = "STUB: not implemented"
	return nil
}
