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
	//go:embed pusher.sh
	pusherSHTextData string

	pusherSHTemplate *template.Template
)

func init() {
	t, err := template.New("pusher.sh").Option("missingkey=zero").Parse(pusherSHTextData)
	if err != nil {
		panic(err)
	}

	pusherSHTemplate = t
}

func renderPusherSH(volDir string, imgRef string, annotationsData map[string]map[string]string, sentinelPath *string) string {
	_ = "STUB: not implemented"
	return ""
}

func NewPusherContainer(inputDirectory string, outputImage string, annotationsData map[string]map[string]string, sentinelPath *string) *corev1.Container {
	_ = "STUB: not implemented"
	return nil
}
