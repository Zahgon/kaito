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

// CatalogEntry represents a pre-computed model entry in the catalog,
// storing the raw HuggingFace config values needed for preset generation
// without requiring runtime API calls.
type CatalogEntry struct {
	Name              string   `yaml:"name"`
	Description       string   `yaml:"description,omitempty"`
	License           string   `yaml:"license,omitempty"`
	PipelineTag       string   `yaml:"pipelineTag,omitempty"`
	BaseModel         []string `yaml:"baseModel,omitempty"`
	ModelFileSize     string   `yaml:"modelFileSize"`
	Architectures     []string `yaml:"architectures"`
	ModelTokenLimit   int      `yaml:"modelTokenLimit"`
	HiddenSize        int      `yaml:"hiddenSize"`
	NumHiddenLayers   int      `yaml:"numHiddenLayers"`
	NumAttentionHeads int      `yaml:"numAttentionHeads"`
	NumKeyValueHeads  int      `yaml:"numKeyValueHeads"`
	LoadFormat        string   `yaml:"loadFormat,omitempty"`
	ConfigFormat      string   `yaml:"configFormat,omitempty"`
	TokenizerMode     string   `yaml:"tokenizerMode,omitempty"`
	HeadDim           int      `yaml:"headDim,omitempty"`
	KVLoraRank        int      `yaml:"kvLoraRank,omitempty"`
	QKRopeHeadDim     int      `yaml:"qkRopeHeadDim,omitempty"`
	QuantMethod       string   `yaml:"quantMethod,omitempty"`
	QuantBits         int      `yaml:"quantBits,omitempty"`
}

// ModelCatalog holds the list of pre-computed model entries.
type ModelCatalog struct {
	Models []CatalogEntry `yaml:"models"`
}

// configKeyMap maps catalog field names to the ordered list of HuggingFace
// config keys to try, mirroring the getInt lookup order.
var configKeyMap = map[string][]string{
	"modelTokenLimit":   {"max_position_embeddings", "n_ctx", "seq_length", "max_seq_len", "max_sequence_length"},
	"hiddenSize":        {"hidden_size", "n_embd", "d_model", "dim"},
	"numHiddenLayers":   {"num_hidden_layers", "n_layer", "n_layers"},
	"numAttentionHeads": {"num_attention_heads", "n_head", "n_heads"},
	"numKeyValueHeads":  {"num_key_value_heads", "n_head_kv", "n_kv_heads"},
}

// optionalKeyMap holds catalog fields that are only stored when present.
var optionalKeyMap = map[string][]string{
	"headDim":       {"head_dim", "attention_head_dim"},
	"kvLoraRank":    {"kv_lora_rank"},
	"qkRopeHeadDim": {"qk_rope_head_dim"},
	"quantMethod":   {"format", "quant_algo", "quant_method"},
}

// fetchModelInfo fetches the model info from the HuggingFace API
// and returns license, pipeline_tag, and base_model.
func fetchModelInfo(g *Generator, repo string) (license, pipelineTag string, baseModel []string) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// pipeline_tag: prefer top-level, fall back to cardData

// cardData holds license and base_model

// When license is "other", HuggingFace stores the actual license
// identifier in the license_name field.

// base_model can be a string or a list of strings

// FetchCatalogEntry fetches a CatalogEntry for a model repo from HuggingFace.
func FetchCatalogEntry(repo, token string) (*CatalogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Fetch model info (license, pipeline, base_model)
}

// Extract quantization config (e.g., AWQ, GPTQ) from HuggingFace config.json.

// Copy format fields from generator (only when non-default)

// Apply hardcoded overrides — these always take precedence over HF values.

// LoadCatalog reads a model_catalog.yaml file and returns its entries.
func LoadCatalog(path string) ([]CatalogEntry, error) { _ = "STUB: not implemented"; return nil, nil }

// SaveCatalog writes catalog entries to a YAML file.
func SaveCatalog(path string, entries []CatalogEntry) error { _ = "STUB: not implemented"; return nil }
