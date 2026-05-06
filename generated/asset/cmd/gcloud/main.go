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

package main

import (
	"context"
	"fmt"
	"os"

	"gcloud/internal/surface/asset"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "async", Usage: "Return immediately without waiting for the operation to complete."},
			&cli.StringFlag{Name: "billing-project", Usage: "Project to charge for quota and billing. Used as the x-goog-user-project header."},
			&cli.StringFlag{Name: "credential-file-override", Usage: "Service-account JSON key file to use for authentication. Overrides Application Default Credentials."},
			&cli.StringFlag{Name: "format", Usage: "Output format: json or yaml.", Value: "json"},
			&cli.StringFlag{Name: "impersonate-service-account", Usage: "Service-account email to impersonate. Base credentials come from ADC or --credential-file-override."},
			&cli.IntFlag{Name: "limit", Usage: "Maximum number of results to return. 0 means no limit."},
			&cli.BoolFlag{Name: "quiet", Aliases: []string{"q"}, Usage: "Disable interactive prompts."},
		},
		Commands: []*cli.Command{
			asset.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
