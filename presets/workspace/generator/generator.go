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

package generator

import (
	"regexp"

	"github.com/kaito-project/kaito/pkg/model"
)

const (
	SystemFileDiskSizeGiB  = 80
	DefaultModelTokenLimit = 2048
	HuggingFaceWebsite     = "https://huggingface.co"
)

// Please update the following model-specific configurations when adding new models to model catalog
var (
	safetensorRegex = regexp.MustCompile(`.*\.safetensors`)
	binRegex        = regexp.MustCompile(`.*\.bin`)
	mistralRegex    = regexp.MustCompile(`consolidated.*\.safetensors`)
	// source: https://github.com/vllm-project/vllm/blob/main/vllm/reasoning/__init__.py
	reasoningParserModeNamePrefixMap = map[string]string{
		"deepseek-r1":  "deepseek_r1",
		"deepseek-v3":  "deepseek_v3",
		"ernie-4.5":    "ernie45",
		"gemma-4":      "gemma4",
		"glm-4.5":      "glm45",
		"granite-3.2":  "granite",
		"holo2":        "holo2",
		"hunyuan-a13b": "hunyuan_a13b",
		"kimi-k2":      "kimi_k2",
		"minimax-m2":   "minimax_m2_append_think",
		"mistral":      "mistral",
		"olmo-3":       "olmo3",
		"qwen3":        "qwen3",
		"qwq-32b":      "deepseek_r1",
		"step3":        "step3",
	}
	reasoningParserArchMap = map[string]string{
		"DeepseekV3ForCausalLM":                  "deepseek_v3",
		"Ernie4_5_VLMoeForConditionalGeneration": "ernie45",
		"Ernie4_5_MoeForCausalLM":                "ernie45",
		"Gemma4ForConditionalGeneration":         "gemma4",
		"Glm4MoeForCausalLM":                     "glm45",
		"HunYuanMoEV1ForCausalLM":                "hunyuan_a13b",
		"GraniteForCausalLM":                     "granite",
		"KimiK2ForCausalLM":                      "kimi_k2",
		"KimiK25ForConditionalGeneration":        "kimi_k2",
		"MiniMaxM2ForCausalLM":                   "minimax_m2_append_think",
		"Mistral3ForConditionalGeneration":       "mistral",
		"MistralForCausalLM":                     "mistral",
		"NemotronForCausalLM":                    "nemotron_v3",
		"NemotronHForCausalLM":                   "nemotron_v3",
		"NemotronH_Nano_VL_V2":                   "nemotron_v3",
		"OlmoForCausalLM":                        "olmo3",
		"Qwen3ForCausalLM":                       "qwen3",
		"Qwen3MoeForCausalLM":                    "qwen3",
		"Qwen3_5ForConditionalGeneration":        "qwen3",
		"Qwen3_5MoeForConditionalGeneration":     "qwen3",
		"GptOssForCausalLM":                      "openai_gptoss",
		"Step3TextForCausalLM":                   "step3",
		"Step3VLForConditionalGeneration":        "step3",
	}

	// source: https://github.com/vllm-project/vllm/blob/main/vllm/tool_parsers/__init__.py
	// key is model name prefix, value is ToolCallParser mode name
	toolCallParserModeNamePrefixMap = map[string]string{
		"hermes-2":      "hermes",
		"hermes-3":      "hermes",
		"mistral":       "mistral",
		"meta-llama-3":  "llama3_json",
		"meta-llama-4":  "llama4_pythonic",
		"granite-3":     "granite",
		"granite-4":     "hermes",
		"internlm":      "internlm",
		"ai21-jamba":    "jamba",
		"llama-xlama":   "xlam",
		"xlam":          "xlam",
		"qwq-32b":       "hermes",
		"qwen2.5":       "hermes",
		"minimax":       "minimax",
		"deepseek-r1":   "deepseek_v3",
		"deepseek-v3":   "deepseek_v3",
		"deepseek-v3.1": "deepseek_v31",
		"deepseek-v3.2": "deepseek_v32",
		"kimi_k2":       "kimi_k2",
		"hunyuan-a13b":  "hunyuan_a13b",
		"longcat":       "longcat",
		"glm-4":         "glm45",
		"glm-4.7":       "glm47",
		"qwen3":         "hermes",
		"qwen3-coder":   "qwen3_xml",
		"qwen3.5":       "qwen3_coder",
		"qwen3.6":       "qwen3_coder",
		"olmo-3":        "olmo3",
		"gigachat3":     "gigachat3",
		"ernie-4.5":     "ernie45",
		"phi4-mini":     "phi4_mini_json",
		"step3p5":       "step3p5",
		"step3":         "step3",
		"seed-oss":      "seed_oss",
		"gemma-3":       "functiongemma",
		"gemma-4":       "gemma4",
	}

	// key is model architecture name, value is ToolCallParser mode name
	toolCallParserArchMap = map[string]string{
		"MistralForCausalLM":                     "mistral",
		"MistralLarge3ForCausalLM":               "mistral",
		"LlamaForCausalLM":                       "llama3_json",
		"Llama4ForConditionalGeneration":         "llama4_pythonic",
		"GraniteForCausalLM":                     "granite",
		"GraniteMoeForCausalLM":                  "granite",
		"GraniteMoeHybridForCausalLM":            "hermes",
		"GPTBigCodeForCausalLM":                  "granite-20b-fc",
		"InternLM2ForCausalLM":                   "internlm",
		"JambaForCausalLM":                       "jamba",
		"Qwen2ForCausalLM":                       "hermes",
		"Qwen3ForCausalLM":                       "hermes",
		"Qwen3MoeForCausalLM":                    "qwen3_xml",
		"Qwen3_5ForConditionalGeneration":        "qwen3_coder",
		"Qwen3_5MoeForConditionalGeneration":     "qwen3_coder",
		"MiniMaxM1ForCausalLM":                   "minimax",
		"MiniMaxM2ForCausalLM":                   "minimax_m2",
		"DeepseekV3ForCausalLM":                  "deepseek_v3",
		"DeepseekV32ForCausalLM":                 "deepseek_v32",
		"GptOssForCausalLM":                      "openai",
		"HunYuanMoEV1ForCausalLM":                "hunyuan_a13b",
		"LongcatFlashForCausalLM":                "longcat",
		"Glm4MoeForCausalLM":                     "glm45",
		"Glm47MoeForCausalLM":                    "glm47",
		"Gemma3ForCausalLM":                      "functiongemma",
		"Gemma4ForConditionalGeneration":         "gemma4",
		"Olmo3ForCausalLM":                       "olmo3",
		"SeedOssForCausalLM":                     "seed_oss",
		"Ernie4_5_VLMoeForConditionalGeneration": "ernie45",
		"Ernie4_5_MoeForCausalLM":                "ernie45",
		"Step3TextForCausalLM":                   "step3",
		"Step3p5TextForCausalLM":                 "step3p5",
		"NemotronHForCausalLM":                   "qwen3_coder",
		"NemotronH_Nano_VL_V2":                   "qwen3_coder",
		"Phi4MiniForCausalLM":                    "phi4_mini_json",
		"KimiK2ForCausalLM":                      "kimi_k2",
		"KimiK25ForConditionalGeneration":        "kimi_k2",
		"GigaChat3ForCausalLM":                   "gigachat3",
	}

	// chatTemplatePrefixMap maps model name prefixes to vllm-customized chat templates.
	// Templates are located in /workspace/chat_templates/ in the KAITO container image.
	// source: https://github.com/vllm-project/vllm/tree/main/examples
	chatTemplatePrefixMap = map[string]string{
		"deepseek-r1": "tool-chat-deepseekr1.jinja",
		"deepseek-v3": "tool-chat-deepseekv3.jinja",
		"llama-3":     "tool-chat-llama3.1-json.jinja",
		"phi-4-mini":  "tool-chat-phi4-mini.jinja",
		"qwen2.5":     "tool-chat-hermes.jinja",
	}

	// tokenizerModePrefixMap maps model name prefixes to their vLLM tokenizer mode.
	tokenizerModePrefixMap = map[string]string{
		// Use deepseek_v32 tokenizer mode for both DeepSeek R1 and V3 models to avoid special token decoding issues:
		// https://github.com/kaito-project/kaito/issues/1976
		"deepseek-r1": "deepseek_v32",
		"deepseek-v3": "deepseek_v32",
	}

	// vllmAttentionBackendPrefixMap maps model name prefixes to their vLLM attention backend.
	// source: https://docs.vllm.ai/en/latest/design/attention_backends/
	vllmAttentionBackendPrefixMap = map[string]string{
		// flashinfer attention backend is chosen by default for LLaMA 3 models, which requires the FlashInfer library to be installed lively.
		// Pin to triton backend as a workaround.
		"llama-3": "TRITON_ATTN",
	}

	// vllmMoeBackendOverride maps exact model names to their vLLM MoE backend.
	// source: https://docs.vllm.ai/en/latest/configuration/engine_args/#-moe-backend
	vllmMoeBackendOverride = map[string]string{
		// Mistral Small 4 FP8 defaults to FlashInfer CUTLASS MoE backend which requires
		// JIT compilation with CUDA dev headers (nvcc, cublasLt, nvrtc).
		// Pin to triton backend to avoid the JIT dependency for now.
		"mistral-small-4-119b-2603": "triton",
		// MiniMax-M2.7 FP8 MoE also defaults to FlashInfer CUTLASS which needs nvcc.
		"minimax-m2.7": "triton",
	}

	// vllmGdnPrefillBackendPrefixMap maps model name prefixes to their vLLM GDN prefill backend.
	// Qwen3.5/3.6 models use hybrid GDN (Gated DeltaNet) attention which defaults to
	// FlashInfer JIT compilation requiring nvcc. Pin to triton to avoid the dependency.
	// source: https://docs.vllm.ai/en/latest/configuration/engine_args/#-gdn-prefill-backend
	vllmGdnPrefillBackendPrefixMap = map[string]string{
		"qwen3.5": "triton",
		"qwen3.6": "triton",
	}

	// vllmExpertParallelEnabled maps model name prefixes to enable expert parallelism.
	// Expert parallelism distributes MoE experts across TP ranks, which can avoid
	// FP8 block quantization issues when expert weight dimensions are not divisible
	// by the quantization block size.
	// source: https://docs.vllm.ai/en/latest/configuration/engine_args/#-enable-expert-parallel
	vllmExpertParallelEnabled = map[string]bool{
		"minimax-m2": true,
	}

	// catalogOverrides provides hardcoded values for models whose HuggingFace
	// config.json omits fields that are required in model_catalog.yaml.
	// Keys are lowercased HuggingFace repo names.
	catalogOverrides = map[string]CatalogEntry{
		// source: https://github.com/huggingface/transformers/blob/main/src/transformers/models/gemma3/configuration_gemma3.py
		"google/gemma-3-4b-it": {
			ModelTokenLimit:   131072,
			NumAttentionHeads: 8,
			NumKeyValueHeads:  4,
			HeadDim:           256,
		},
		// Based on Gemma 3 model card, the 128K context window (131072 tokens) applies to all Gemma 3 4B/12B/27B sizes
		// source: https://huggingface.co/google/gemma-3-27b-it
		"google/gemma-3-27b-it": {
			ModelTokenLimit: 131072,
		},
		"mistralai/mistral-large-3-675b-instruct-2512": {
			// source: https://docs.vllm.ai/en/v0.17.1/api/vllm/model_executor/models/mistral_large_3/
			Architectures: []string{"MistralLarge3ForCausalLM"},
			PipelineTag:   "text-generation",
		},
	}
)

type Generator struct {
	ModelRepo      string
	Token          string
	Param          model.PresetParam
	CatalogData    []byte // Optional embedded catalog YAML
	IsMistralModel bool

	// Analyzed params
	LoadFormat    string
	ConfigFormat  string
	TokenizerMode string
	ModelConfig   map[string]interface{}
}

func NewGenerator(modelRepo, token string) *Generator { _ = "STUB: not implemented"; return nil }

// Initialize default PresetParam

func (g *Generator) getAuthHeader() string { _ = "STUB: not implemented"; return "" }

func (g *Generator) fetchURL(url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

func (g *Generator) FetchModelMetadata() error { _ = "STUB: not implemented"; return nil }

// listRepoFiles fetches the full file tree for the model repo from HuggingFace.
func (g *Generator) listRepoFiles() ([]FileInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// selectWeightFiles picks the model weight files to use and detects whether
// the model uses Mistral format. For Mistral-format models (those with
// consolidated*.safetensors), it sets g.IsMistralModel and returns only the
// consolidated files. For standard models, it prefers .safetensors over .bin
// when both are present.
func (g *Generator) selectWeightFiles(files []FileInfo) []FileInfo {
	_ = "STUB: not implemented"
	return nil
}

// Prefer safetensors over bin files when both exist.

func (g *Generator) setMistralMode() { _ = "STUB: not implemented"; return }

func calculateModelFileSize(files []FileInfo) string { _ = "STUB: not implemented"; return "" }

// fetchAndParseConfig downloads and parses the model's config.json. For
// Mistral-format models, it falls back to params.json if config.json is absent.
func (g *Generator) fetchAndParseConfig() error { _ = "STUB: not implemented"; return nil }

// config.json not available; fall back to params.json (Mistral native format).

// vLLM delegates the loading of config.json to HuggingFace transformer library
// (https://github.com/huggingface/transformers/blob/main/src/transformers/configuration_utils.py#L552).
// The library uses Python's json.loads which accepts non-standard JSON literals (e.g. Infinity, NaN).
// However, Go's standard library encoding/json only supports standard JSON values. sanitizeJSON replaces
// non-standard JSON literals (Infinity, -Infinity, NaN) with null so the data can be parsed by encoding/json.
func sanitizeJSON(data []byte) []byte {
	_ = "STUB: not implemented"
	// Replace standalone Infinity, -Infinity, NaN with null
	return nil
}

// Preserve the prefix (comma, bracket, colon, whitespace) before the value
//nolint:gocritic

func (g *Generator) fetchConfigFile(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mergeTextConfig promotes fields from a nested "text_config" or "llm_config"
// object into the top-level config. This is needed for multimodal models
// (e.g., Gemma-3, Ministral-3, Nemotron-VL) where architecture-specific
// parameters live under a nested config key.
func (g *Generator) mergeTextConfig() { _ = "STUB: not implemented"; return }

func getInt(config map[string]interface{}, keys []string, defaultVal int) int {
	_ = "STUB: not implemented"
	return 0
}

// getString looks up the first matching key in config that has a non-empty
// string value. Keys are tried in order; the first hit wins.
func getString(config map[string]interface{}, keys []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) ParseModelMetadata() { _ = "STUB: not implemented"; return }

// Override architectures for specific model families only when none were parsed

// set reasoning parser based on model name prefix

// set reasoning parser based on model architecture if not set by name prefix

// set ToolCallParser based on model name prefix
// sort the keys of toolCallParserModeNamePrefixMap in reverse alphabetical order and then iterate
// this is to ensure that longer (more specific) prefixes are matched first

// set ToolCallParser based on model architecture if not set by name prefix

// set ChatTemplate based on model name prefix

// Parse quantization config (e.g., AWQ, GPTQ) from HuggingFace config.json.

func (g *Generator) calculateStorageSize() string { _ = "STUB: not implemented"; return "" }

func (g *Generator) calculateKVCacheTokenSize() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// DeepSeek MLA

// Fallback KV heads

// TODO: honor kv-cache quantization instead of hardcoding fp16
// fp16

func (g *Generator) FinalizeParams() { _ = "STUB: not implemented"; return }

// VLLM Params

// Override tokenizer mode based on model name prefix

// Set attention backend based on model name prefix

// Set MoE backend based on exact model name match

// Set GDN prefill backend based on model name prefix

// Enable expert parallelism based on model name prefix

// loadFromCatalog checks whether the model repo exists in the embedded catalog.
// If found, it populates the generator's ModelConfig and Param fields from the
// catalog entry, avoiding any HuggingFace API calls.
func (g *Generator) loadFromCatalog() bool { _ = "STUB: not implemented"; return false }

// Populate ModelConfig from catalog entry so existing calculation
// functions (ParseModelMetadata, FinalizeParams) work unchanged.

// Restore quantization_config so ParseModelMetadata can pick it up.

// Set architectures in config for ParseModelMetadata to pick up

// Populate fields that FetchModelMetadata would have set

func (g *Generator) Generate() (*model.PresetParam, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GeneratePreset is the global function to generate preset param.
// If catalogData is provided, the generator will check for the model in the
// catalog before making any HuggingFace API calls.
func GeneratePreset(modelRepo, token string, catalogData ...[]byte) (*model.PresetParam, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
