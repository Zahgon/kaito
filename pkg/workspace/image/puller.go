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

package image

import (
	_ "embed"
	"text/template"

	corev1 "k8s.io/api/core/v1"
)

var (
	//go:embed puller.sh
	pullerSHTextData string

	pullerSHTemplate *template.Template
)

func init() {
	t, err := template.New("puller.sh").Option("missingkey=zero").Parse(pullerSHTextData)
	if err != nil {
		panic(err)
	}

	pullerSHTemplate = t
}

func renderPullerSH(imgRef string, volDir string) string { _ = "STUB: not implemented"; return "" }

func NewPullerContainer(inputImage string, outputDirectory string) *corev1.Container {
	_ = "STUB: not implemented"
	return nil
}
