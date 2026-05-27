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

package model

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kaito-project/kaito/pkg/sku"
)

type Model interface {
	// GetInferenceParameters returns the preset inference parameters for the model.
	GetInferenceParameters() *PresetParam

	// GetTuningParameters returns the preset tuning parameters for the model.
	GetTuningParameters() *PresetParam

	// SupportDistributedInference checks if the model supports distributed inference.
	SupportDistributedInference() bool

	// SupportTuning checks if the model supports tuning.
	SupportTuning() bool
}

// RuntimeName is LLM runtime name.
type RuntimeName string

const (
	RuntimeNameHuggingfaceTransformers RuntimeName = "transformers"
	RuntimeNameVLLM                    RuntimeName = "vllm"

	DefaultTuningMainFile = "/workspace/tfs/fine_tuning.py"
	ConfigfileNameVLLM    = "inference_config.yaml"

	// PortRayCluster is the default port for communication between the head and worker nodes in a Ray cluster.
	PortRayCluster = 6379
)

// Metadata defines the metadata for a model.
type Metadata struct {
	// Name is the name of the model, which serves as a unique identifier.
	// It is used to register the model information and retrieve it later.
	Name string `yaml:"name"`

	// ModelType is the type of the model, which indicates the kind of model
	// it is. Currently, the only supported types are "text-generation" and
	// "llama2-completion" (deprecated).
	ModelType string `yaml:"type"`

	// Version is the version of the model. It is a URL that points to the
	// model's huggingface page, which contains the model's repository ID
	// and revision ID, e.g. https://huggingface.co/mistralai/Mistral-7B-v0.3/commit/d8cadc02ac76bd617a919d50b092e59d2d110aff.
	Version string `yaml:"version"`

	// Runtime is the runtime environment in which the model operates.
	// Currently, the only supported runtime is "tfs".
	Runtime string `yaml:"runtime"`

	// DownloadAtRuntime indicates whether the model should be downloaded
	// at runtime. If set to true, the model will be downloaded when the
	// model deployment is created, and the container image will always be
	// the KAITO base image. If set to false, a container image whose name
	// contains the model name will be used, in which the model weights are baked.
	// +optional
	DownloadAtRuntime bool `yaml:"downloadAtRuntime,omitempty"`

	// DownloadAuthRequired indicates whether the model requires authentication to download.
	// +optional
	DownloadAuthRequired bool `yaml:"downloadAuthRequired,omitempty"`

	// Tag is the tag of the container image used to run the model.
	// If the model uses the KAITO base image, the tag field can be ignored
	// +optional
	Tag string `yaml:"tag,omitempty"`

	// Registry is the container registry where the model is stored.
	// If this is empty, os.Getenv("PRESET_REGISTRY_NAME") will be used.
	// +optional
	Registry string `yaml:"registry,omitempty"`

	// Deprecated indicates if the model is deprecated.
	// +optional
	Deprecated bool `yaml:"deprecated,omitempty"`

	// Architectures specifies the supported architectures for the model
	// This field is only for best effort supported vLLM models.
	// +optional
	Architectures []string `yaml:"architectures,omitempty"`

	// DType specifies the data type used by the model (e.g., "bfloat16", "float16", "float32").
	// This field is only for best effort supported vLLM models.
	// +optional
	DType string `yaml:"dtype,omitempty"`

	// ModelFileSize is the size of the model file, example: 14Gi.
	// This field is only for best effort supported vLLM models.
	// +optional
	ModelFileSize string `yaml:"modelFileSize,omitempty"`

	// DiskStorageRequirement is the disk storage requirement for the model, example: 90Gi.
	// This field is only for best effort supported vLLM models.
	// +optional
	DiskStorageRequirement string `yaml:"diskStorageRequirement,omitempty"`

	// BytesPerToken is the number of bytes used to represent each token in the model.
	// This field is only for best effort supported vLLM models.
	// +optional
	BytesPerToken int `yaml:"bytesPerToken,omitempty"`

	// ModelTokenLimit is the maximum number of tokens (context window) supported by the model.
	// This field is only for best effort supported vLLM models.
	// +optional
	ModelTokenLimit int `yaml:"modelTokenLimit,omitempty"`

	// ToolCallParser specifies the parser used for tool calls within the model.
	// This field is only for best effort supported vLLM models.
	// +optional
	ToolCallParser string `yaml:"toolCallParser,omitempty"`

	// ReasoningParser specifies the parser used for reasoning within the model.
	// This field is only for best effort supported vLLM models.
	// +optional
	ReasoningParser string `yaml:"reasoningParser,omitempty"`

	// ChatTemplate is the chat template file name used for chat models.
	// This field is only for best effort supported vLLM models.
	// +optional
	ChatTemplate string `yaml:"chatTemplate,omitempty"`

	// AllowRemoteFiles indicates whether the model allows loading remote files.
	// This field is only for best effort supported vLLM models.
	// +optional
	AllowRemoteFiles bool `yaml:"allowRemoteFiles,omitempty"`

	// QuantMethod specifies the weight quantization method used by the model
	// (e.g., "awq", "gptq"). Maps to quantization_config.quant_method in the
	// model's HuggingFace config.json.
	// +optional
	QuantMethod string `yaml:"quantMethod,omitempty"`

	// QuantBits specifies the number of bits used for weight quantization
	// (e.g., 4 for 4-bit AWQ). Maps to quantization_config.bits in the
	// model's HuggingFace config.json.
	// +optional
	QuantBits int `yaml:"quantBits,omitempty"`
}

// Validate checks if the Metadata is valid.
func (m *Metadata) Validate() error {
	_ = "STUB: not implemented"
	// Some models requiring authentication may not have a version URL, so we allow it to be empty until
	// we remove support for preset models requiring authentication.
	return nil
}

// PresetParam defines the preset inference parameters for a model.
type PresetParam struct {
	Metadata

	DiskStorageRequirement string // Disk storage requirements for the model.
	// DiskStorageRequirement is calculated as:
	// (TotalSafeTensorFileSize × 2.5 + 48) rounded up to the next multiple of 10.
	// This formula accounts for model weights, optimization files, and runtime overhead.
	// Example: For a 14Gi model, calculation is: 14 × 2.5 + 48 = 83, rounded up to 90Gi.

	ImageAccessMode               string         // Defines where the Image is Public or Private.
	GPUCountRequirement           string         // Number of GPUs required for the Preset. Used for inference.
	TotalSafeTensorFileSize       string         // Total SafeTensor file size for the Preset. Used for inference.
	TuningPerGPUMemoryRequirement map[string]int // Min GPU memory per tuning method (batch size 1). Used for tuning.
	BytesPerToken                 int            // Number of bytes per token for the model. It is calculated by 2 * hidden_layers * kv_heads * head_dim (hidden_size/num_attemtion_numbers) * dtype_size
	ModelTokenLimit               int            // Maximum number of tokens (context window) supported by the model. Maps to 'max_position_embeddings' in the model's Hugging Face config.json.

	// To determine TotalSafeTensorFileSize and BytesPerToken values for a new model,
	// run the presets/workspace/generator/preset_generator.py script
	// with the model's Hugging Face repository ID as an argument.

	// AttnType specifies the attention implementation (e.g., MHA, GQA, MLA).
	// Calculated by the preset generator based on model config.
	AttnType string `yaml:"attn_type,omitempty"`

	RuntimeParam

	// ReadinessTimeout defines the maximum duration for creating the workload.
	// This timeout accommodates the size of the image, ensuring pull completion
	// even under slower network conditions or unforeseen delays.
	ReadinessTimeout time.Duration
}

// RuntimeParam defines the llm runtime parameters.
type RuntimeParam struct {
	Transformers HuggingfaceTransformersParam
	VLLM         VLLMParam
	// Disable the tensor parallelism
	DisableTensorParallelism bool
}

type HuggingfaceTransformersParam struct {
	BaseCommand       string            // The initial command (e.g., 'accelerate launch') used in the command line.
	AccelerateParams  map[string]string // Parameters for configuring the accelerate command.
	InferenceMainFile string            // The main file for inference.
	ModelRunParams    map[string]string // Parameters for running the model training/inference.
	// The model name used in the OpenAI serving API.
	ModelName string
	// Tag is the ORAS image tag for pre-built model weights.
	Tag string
}

type VLLMParam struct {
	RayLeaderBaseCommand string
	RayLeaderParams      map[string]string
	RayWorkerBaseCommand string
	RayWorkerParams      map[string]string
	// BaseCommand is the command used to start the inference server.
	BaseCommand string
	// The model name used in the openai serving API.
	// see https://platform.openai.com/docs/api-reference/chat/create#chat-create-model.
	ModelName string
	// Parameters for running the model training/inference.
	ModelRunParams map[string]string
	// Indicates if vllm supports LoRA (Low-Rank Adaptation) for this model.
	// doc: https://docs.vllm.ai/en/latest/models/supported_models.html#text-generation-task-generate
	DisallowLoRA bool
}

func (p *PresetParam) DeepCopy() *PresetParam { _ = "STUB: not implemented"; return nil }

func (rp *RuntimeParam) DeepCopy() RuntimeParam {
	_ = "STUB: not implemented"
	return *new(RuntimeParam)
}

func (h *HuggingfaceTransformersParam) DeepCopy() HuggingfaceTransformersParam {
	_ = "STUB: not implemented"
	return *new(HuggingfaceTransformersParam)
}

func (v *VLLMParam) DeepCopy() VLLMParam { _ = "STUB: not implemented"; return *new(VLLMParam) }

// RuntimeContext defines the runtime context for a model.
type RuntimeContext struct {
	RuntimeName          RuntimeName
	GPUConfig            *sku.GPUConfig
	ConfigVolume         *corev1.VolumeMount
	SKUNumGPUs           int
	NumNodes             int
	WorkspaceMetadata    metav1.ObjectMeta
	DistributedInference bool
	MaxModelLen          int // max-model-len parameter for vLLM
	RuntimeContextExtraArguments
}

type RuntimeContextExtraArguments struct {
	AdaptersEnabled        bool
	AdapterStrengthEnabled bool
	PerformanceMode        string // vLLM --performance-mode; defaults to "balanced"
}

func (p *PresetParam) GetInferenceCommand(rc RuntimeContext) []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PresetParam) buildHuggingfaceInferenceCommand() []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PresetParam) buildVLLMInferenceCommand(rc RuntimeContext) []string {
	_ = "STUB: not implemented"
	// For InferenceSet-managed workspaces, determine the served-model-name:
	// - MRI workspaces (have multiroleinference.kaito.sh/created-by label): use VLLM.ModelName
	//   so all roles share a single model identifier for EPP routing.
	// - Standalone InferenceSet workspaces: use the InferenceSet name (label value)
	//   so EPP routes requests by InferenceSet identity.
	// - Fallback: use VLLM.ModelName if available.
	return nil
}

// Note: string literal used to avoid import cycle with api/v1alpha1 package.
// Matches v1alpha1.LabelMultiRoleInferenceParent.

// Dynamically determine dtype based on GPU compute capability.
// bfloat16 requires CUDA compute capability >= 8.0 (Ampere+).
// Fall back to float16 on older GPUs.

// Hybrid Mamba/Attention models (e.g., NemotronH) require the hybrid KV cache
// manager in vLLM, which is incompatible with LMCache KV cache CPU offloading.
// Disable offloading for these architectures to prevent startup crashes.

// Parallelism strategy follows a 3-tier hierarchy (see configureParallelism):
//  1. Data Parallelism (DP)   – model fits on a single GPU
//  2. Tensor Parallelism (TP) – model fits on a single node (multiple GPUs)
//  3. Pipeline Parallelism (PP) + TP – model requires multiple nodes

// Single-node path: no Ray cluster needed.

// Multi-node path: set up a Ray cluster for cross-node parallelism.

// configureParallelism sets the vLLM parallelism parameters according to a
// 3-tier strategy based on where the model can be placed:
//
//  1. Single-GPU (DP): If the model file size is less than 50% of a single GPU's
//     memory, each GPU can serve the model independently. We use data parallelism
//     to run replicas across GPUs for maximum throughput.
//
//  2. Single-node (TP): If the model is too large for one GPU but fits within the
//     combined memory of all GPUs on a single node, we shard the model across GPUs
//     using tensor parallelism.
//
//  3. Multi-node (PP + TP): If the model exceeds a single node's capacity, we use
//     pipeline parallelism across nodes, with tensor parallelism within each node.
func (p *PresetParam) configureParallelism(rc RuntimeContext) { _ = "STUB: not implemented"; return }

// Tier 1: Model fits on a single GPU → Data Parallelism.
// Use DP only on a single node; multi-node DP is not supported.

// In this branch, data-parallel-size is guaranteed to be > 1; disable kv cache CPU offloading
// due to conflicts between data parallelism and CPU offloading.

// Tier 2: Model fits on a single node → Tensor Parallelism.
// TP is set to the number of GPUs on the node.

// Tier 3: Model requires multiple nodes → Pipeline Parallelism + TP.

// Disable kv cache CPU offloading when pipeline parallelism is enabled.
// TODO: LMCache doesn't support cross-node PP in CPU offload mode.

// PP is set to the number of nodes.

// Since vllm 0.12.0, we need to set the distributed-executor-backend explicitly.

// buildMultiNodeRayCommand constructs the shell command for multi-node inference
// using a Ray cluster. Pod index 0 is the leader; all other pods are workers.
func (p *PresetParam) buildMultiNodeRayCommand(rc RuntimeContext) []string {
	_ = "STUB: not implemented"
	return nil
}

// leader if pod index is 0, otherwise worker
// leader: start ray head + model

// worker: join the cluster

// getModelFileSize returns the model file size as a resource.Quantity.
// It tries TotalSafeTensorFileSize first (preset models), then ModelFileSize (best-effort models).
func (p *PresetParam) getModelFileSize() *resource.Quantity { _ = "STUB: not implemented"; return nil }

// isVLLMHybridKVCacheManagerRequired returns true if the model uses a hybrid
// architecture (e.g., Mamba/Attention) that requires vLLM's hybrid KV cache manager
// (https://docs.vllm.ai/en/latest/design/hybrid_kv_cache_manager/)
func (p *PresetParam) isVLLMHybridKVCacheManagerRequired() bool {
	_ = "STUB: not implemented"
	return false
}

// modelFitsOnSingleGPU returns true when the model file size is smaller than
// 50% of a single GPU's memory, meaning the entire model can be loaded onto
// one GPU with headroom to spare.
func (p *PresetParam) modelFitsOnSingleGPU(rc RuntimeContext) bool {
	_ = "STUB: not implemented"
	return false
}

// Single GPU memory = total GPU memory / number of GPUs.
// Condition: modelSize < 0.5 * singleGPUMem
// Rearranged to avoid division: modelSize * numGPUs * 2 < totalGPUMem.

func (p *PresetParam) Validate(rc RuntimeContext) error { _ = "STUB: not implemented"; return nil }

// Only support Huggingface for now
func (p *PresetParam) GetTuningCommand(rc RuntimeContext) []string {
	_ = "STUB: not implemented"
	return nil
}
