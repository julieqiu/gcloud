// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package runtime provides shared output and prompt helpers for generated
// gcloud surface packages.
package runtime

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

// FormatResponse marshals msg according to format ("json" or "yaml") and
// returns the result as a string. Empty format defaults to "json".
func FormatResponse(format string, msg proto.Message) (string, error) {
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		return "", err
	}
	switch format {
	case "", "json":
		return string(jsonBytes), nil
	case "yaml":
		yamlBytes, err := yaml.JSONToYAML(jsonBytes)
		if err != nil {
			return "", err
		}
		return string(yamlBytes), nil
	}
	return "", fmt.Errorf("unsupported format %q (want json or yaml)", format)
}

// Ptr returns a pointer to v. Used by generated request struct literals to
// assign values to proto3 optional fields, which are encoded as pointers in
// the protobuf-generated Go types.
func Ptr[T any](v T) *T {
	return &v
}

// Confirm reads a y/N answer from stdin. It returns nil on yes, an error on
// no or EOF. Skip the call when --quiet is set.
func Confirm(prompt string) error {
	fmt.Fprintf(os.Stderr, "%s [y/N] ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return errors.New("aborted")
	}
	answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
	if answer != "y" && answer != "yes" {
		return errors.New("aborted")
	}
	return nil
}
