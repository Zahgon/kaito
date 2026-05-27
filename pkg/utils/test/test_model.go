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

package test

import (
	"github.com/kaito-project/kaito/pkg/model"
)

type baseTestModel struct{}

var emptyParams = map[string]string{}

func (*baseTestModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*baseTestModel) GetTuningParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*baseTestModel) SupportDistributedInference() bool { _ = "STUB: not implemented"; return false }

func (*baseTestModel) SupportTuning() bool { _ = "STUB: not implemented"; return false }

type testModel struct {
	baseTestModel
}

func (*testModel) SupportDistributedInference() bool { _ = "STUB: not implemented"; return false }

type testDistributedModel struct {
	baseTestModel
}

func (*testDistributedModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*testDistributedModel) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	return false
}

type testNoTensorParallelModel struct {
	baseTestModel
}

func (*testNoTensorParallelModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*testNoTensorParallelModel) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	return false
}

type testNoLoraSupportModel struct {
	baseTestModel
}

type testModelDownload struct {
	baseTestModel
}

func (*testModelDownload) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	return false
}

func (*testModelDownload) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

type testModelDownloadA100 struct {
	baseTestModel
}

func (*testModelDownloadA100) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	return false
}

func (*testModelDownloadA100) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*testNoLoraSupportModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

func (*testNoLoraSupportModel) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	return false
}

type testFalcon7BModel struct {
	baseTestModel
}

func (*testFalcon7BModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

// falcon-7b has 71 attention heads (prime number)

func (*testFalcon7BModel) SupportDistributedInference() bool {
	_ = "STUB: not implemented"
	// Due to tensor parallelism being disabled
	return false
}

type testQwen25Coder32BModel struct {
	baseTestModel
}

func (*testQwen25Coder32BModel) GetInferenceParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	return nil
}

// Supports tensor parallelism

func (*testQwen25Coder32BModel) GetTuningParameters() *model.PresetParam {
	_ = "STUB: not implemented"
	// Not recommended for further fine-tuning instruct models
	return nil
}

func (*testQwen25Coder32BModel) SupportTuning() bool { _ = "STUB: not implemented"; return false }

func RegisterTestModel() { _ = "STUB: not implemented"; return }
