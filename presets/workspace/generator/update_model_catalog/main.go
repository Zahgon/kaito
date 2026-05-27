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

// Command update_model_catalog updates model_catalog.yaml with latest
// values from HuggingFace.
//
// Usage:
//
//	go run ./presets/workspace/generator/update_model_catalog                       # Update all existing entries
//	go run ./presets/workspace/generator/update_model_catalog --repos org/m1,org/m2 # Add or update specific repos
//	go run ./presets/workspace/generator/update_model_catalog --dry-run             # Show changes without writing
//	go run ./presets/workspace/generator/update_model_catalog --token <HF_TOKEN>    # Use auth token
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaito-project/kaito/presets/workspace/generator"
)

func main() {
	repos := flag.String("repos", "", "Comma-separated list of repos to add or update (default: all existing)")
	dryRun := flag.Bool("dry-run", false, "Show changes without writing to file")
	token := flag.String("token", os.Getenv("HF_TOKEN"), "HuggingFace API token (default: $HF_TOKEN)")
	catalog := flag.String("catalog", filepath.Join("presets", "workspace", "models", "model_catalog.yaml"), "Path to model_catalog.yaml")
	flag.Parse()

	var repoList []string
	if *repos != "" {
		for _, r := range strings.Split(*repos, ",") {
			r = strings.TrimSpace(r)
			if r != "" {
				repoList = append(repoList, r)
			}
		}
	}

	if err := updateModelCatalog(repoList, *dryRun, *token, *catalog); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func formatDiff(old, new *generator.CatalogEntry) string { _ = "STUB: not implemented"; return "" }

func catalogFields(e *generator.CatalogEntry) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func catalogFieldLines(e *generator.CatalogEntry) []string { _ = "STUB: not implemented"; return nil }

func updateModelCatalog(repos []string, dryRun bool, token, catalogPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Infer missing architectures from base model on HuggingFace.
