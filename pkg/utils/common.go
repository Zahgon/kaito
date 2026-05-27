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
	awsapis "github.com/aws/karpenter-provider-aws/pkg/apis"
	awsv1 "github.com/aws/karpenter-provider-aws/pkg/apis/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	karpenterapis "sigs.k8s.io/karpenter/pkg/apis"
	karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

	"github.com/kaito-project/kaito/pkg/sku"
)

const (
	errInvalidModelVersionURL = "invalid model version URL: %s. Expected format: https://huggingface.co/<org>/<model>/commit/<revision>"
)

var (
	karpenterSchemeGroupVersion = schema.GroupVersion{Group: karpenterapis.Group, Version: "v1"}
	awsSchemeGroupVersion       = schema.GroupVersion{Group: awsapis.Group, Version: "v1"}

	KarpenterSchemeBuilder = runtime.NewSchemeBuilder(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(karpenterSchemeGroupVersion,
			&karpenterv1.NodePool{},
			&karpenterv1.NodePoolList{},
			&karpenterv1.NodeClaim{},
			&karpenterv1.NodeClaimList{},
		)
		metav1.AddToGroupVersion(scheme, karpenterSchemeGroupVersion)
		return nil
	})
	AwsSchemeBuilder = runtime.NewSchemeBuilder(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(awsSchemeGroupVersion,
			&awsv1.EC2NodeClass{},
			&awsv1.EC2NodeClassList{},
		)
		metav1.AddToGroupVersion(scheme, awsSchemeGroupVersion)
		return nil
	})
)

func Contains(s []string, e string) bool { _ = "STUB: not implemented"; return false }

// SearchRawExtension performs a search for a key within a runtime.RawExtension.
func SearchRawExtension(raw runtime.RawExtension, key string) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func BuildCmdStr(baseCommand string, runParams ...map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func BuildIfElseCmdStr(condition string, trueCmd string, trueCmdParams map[string]string, falseCmd string, falseCmdParams map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func ShellCmd(command string) []string { _ = "STUB: not implemented"; return nil }

func GetReleaseNamespace() (string, error) {
	_ = "STUB: not implemented"
	// Path to the namespace file inside a Kubernetes pod
	return "", nil
}

// Attempt to read the namespace from the file

// Fallback: Read the namespace from an environment variable

func GetSKUHandler() (sku.CloudSKUHandler, error) {
	_ = "STUB: not implemented"
	// Get the cloud provider from the environment
	return *new(sku.CloudSKUHandler), nil
}

// Select the correct SKU handler based on the cloud provider

func IsAzureCloudProvider() bool { _ = "STUB: not implemented"; return false }

func GetGPUConfigBySKU(instanceType string) (*sku.GPUConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGPUConfigFromNodeLabels extracts GPU configuration from nvidia.com labels on a node
func GetGPUConfigFromNodeLabels(node *corev1.Node) (*sku.GPUConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if all required nvidia.com labels are present

// Parse GPU count

// Parse GPU memory (nvidia.com/gpu.memory is per-GPU memory in MiB).

// Parse CUDA compute capability from nvidia.com/cuda.compute.major and nvidia.com/cuda.compute.minor labels.
// These are set by the NVIDIA GPU Feature Discovery (GFD) DaemonSet.

// SKU is not available from node labels

func ExtractAndValidateRepoName(image string) error {
	_ = "STUB: not implemented"
	// Extract repository name (part after the last / and before the colon :)
	// For example given image: modelsregistry.azurecr.io/ADAPTER_HERE:0.0.1
	return nil
}

// Extracts "ADAPTER_HERE:0.0.1"
// Extracts "ADAPTER_HERE"

// Check if repository name is lowercase

func SelectNodes(qualified []*corev1.Node, preferred []string, previous []string, count int) []*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// either all are preferred, or none is preferred

// either all are previous, or none is previous

// Choose node created by gpu-provisioner and karpenter since it is more likely to be empty to use.

// ParseHuggingFaceModelVersion parses the model version in the format of https://huggingface.co/<org>/<model>/commit/<revision>
// and returns the repoId and revision. If the commit is not specified, it returns an empty string for revision,
// and the main branch HEAD commit is used.
//
// Example 1:
//
//	Version: "https://huggingface.co/tiiuae/falcon-7b/commit/ec89142b67d748a1865ea4451372db8313ada0d8"
//	RepoId: "tiiuae/falcon-7b"
//	Revision: "ec89142b67d748a1865ea4451372db8313ada0d8"
//
// Example 2:
//
//	Version: https://huggingface.co/tiiuae/falcon-7b
//	RepoId: "tiiuae/falcon-7b"
//	Revision: "" (main branch HEAD commit is used)
func ParseHuggingFaceModelVersion(version string) (repoId string, revision string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Expected path: "<org>/<model>"

// Expected path: "<org>/<model>/commit/<revision>"

// getRayLeaderHost constructs the leader host for the Ray cluster.
func GetRayLeaderHost(meta metav1.ObjectMeta) string { _ = "STUB: not implemented"; return "" }

// InferencePoolName returns the name of the inference pool for the given workspace.
func InferencePoolName(workspaceName string) string { _ = "STUB: not implemented"; return "" }

// ClientObjectSpecEqual compares the spec field of two client.Objects for equality.
// For example:
//
//	a:   {"apiVersion": "apps/v1", "kind": "Deployment", "spec": {"replicas": 2}}
//	b:   {"apiVersion": "apps/v1", "kind": "Deployment", "spec": {"replicas": 2}}
//	result: true
//
//	c:   {"apiVersion": "apps/v1", "kind": "Deployment", "spec": {"replicas": 3}}
//	d:   {"apiVersion": "apps/v1", "kind": "Deployment", "spec": {"replicas": 2}}
//	result: false
func ClientObjectSpecEqual(a, b client.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
