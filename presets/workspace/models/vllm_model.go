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

package models

import (
	"context"
	_ "embed"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kaito-project/kaito/pkg/model"
)

var (
	//go:embed model_catalog.yaml
	modelCatalogYAML []byte
)

// registerModel registers a HuggingFace model with the given ID and parameters
// into the model registry and returns the registered model. If param is nil,
// it returns nil and does not register a model.
func registerModel(hfModelCardID string, param *model.PresetParam) model.Model {
	_ = "STUB: not implemented"
	return *new(model.Model)
}

// GetModelByNameWithToken returns a vLLM-compatible model for the given modelName using
// a pre-resolved access token. Unlike GetModelByName, this function does not perform any
// Kubernetes Secret lookups; the caller is responsible for obtaining the token beforehand.
// Pass an empty string for token when working with public models that require no authentication.
func GetModelByNameWithToken(ctx context.Context, modelName, token string) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// Redirect legacy preset names (e.g. "phi-4") to their full HuggingFace
// model ID (e.g. "microsoft/phi-4").

// GetModelByName returns a vLLM-compatible model for the given modelName.
// If the modelName contains a "/", it fetches an access token from the
// Kubernetes Secret identified by secretName and secretNamespace,
// then generates a preset for the corresponding HuggingFace model.
// Prefer GetModelByNameWithToken when the token has already been resolved by the caller.
func GetModelByName(ctx context.Context, modelName, secretName, secretNamespace string, kubeClient client.Client) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// Redirect legacy preset names (e.g. "phi-4") to their full HuggingFace
// model ID (e.g. "microsoft/phi-4").

// only log the error here since token may not be required for public models

// generateHuggingFaceModel generates or retrieves a vLLM preset for modelName (which must
// contain a "/") using the provided token.
func generateHuggingFaceModel(modelName, token string) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// check whether the model is in the supported model architecture list

type vLLMCompatibleModel struct {
	model              model.Metadata
	generatedRunParams map[string]string // vLLM run params produced by the generator
}

func (m *vLLMCompatibleModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

// Apply defaults for keys the generator doesn't set.

// For quantized models, let vLLM auto-detect the optimal dtype
// TODO: test if we can always set dtype to "auto"

func (m *vLLMCompatibleModel) GetTuningParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*vLLMCompatibleModel) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	return false
}

func (m *vLLMCompatibleModel) SupportTuning() bool { _ = "STUB: not implemented"; return false }

// GetHFTokenFromSecret retrieves the HuggingFace token from a Kubernetes secret.
// If secretName is empty, it returns an empty string without error.
// If secretNamespace is empty, it defaults to "default".
// An error is returned if kubeClient is nil, the secret cannot be retrieved,
// or the HF_TOKEN key is not present in the secret data.
func GetHFTokenFromSecret(ctx context.Context, kubeClient client.Client, secretName, secretNamespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
