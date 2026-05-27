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

package sku

func NewAzureSKUHandler() CloudSKUHandler { _ = "STUB: not implemented"; return *new(CloudSKUHandler) }

// https://learn.microsoft.com/en-us/azure/virtual-machines/sizes/gpu-accelerated/ncadsh100v5-series

// https://learn.microsoft.com/en-us/azure/virtual-machines/sizes/gpu-accelerated/ndh100v5-series

// https://learn.microsoft.com/en-us/azure/virtual-machines/sizes/gpu-accelerated/nccadsh100v5-series

// https://learn.microsoft.com/en-us/azure/virtual-machines/sizes/gpu-accelerated/nd-h200-v5-series

// Not supporting partial gpu skus for now
// {SKU: "Standard_NG8ads_V620_v1", GPUCount: 1.0 / 4.0, GPUMem: 8, GPUModel: "AMD Radeon PRO V620"},
// {SKU: "Standard_NG16ads_V620_v1", GPUCount: 1.0 / 2.0, GPUMem: 16, GPUModel: "AMD Radeon PRO V620"},
// {SKU: "Standard_NV4as_v4", GPUCount: 1.0 / 8.0, GPUMem: 2, GPUModel: "AMD Radeon Instinct MI25"},
// {SKU: "Standard_NV8as_v4", GPUCount: 1.0 / 4.0, GPUMem: 4, GPUModel: "AMD Radeon Instinct MI25"},
// {SKU: "Standard_NV16as_v4", GPUCount: 1.0 / 2.0, GPUMem: 8, GPUModel: "AMD Radeon Instinct MI25"},
