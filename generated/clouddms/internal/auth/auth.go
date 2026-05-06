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

// Package auth provides shared client-option helpers for generated gcloud
// surface packages.
package auth

import (
	"context"
	"fmt"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/auth/credentials"
	"cloud.google.com/go/auth/credentials/impersonate"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/option"
)

// ClientOptions returns ClientOption values for the persistent auth flags.
// It reads --impersonate-service-account, --credential-file-override, and
// --billing-project from cmd and produces a credentials chain via
// cloud.google.com/go/auth.
func ClientOptions(ctx context.Context, cmd *cli.Command) ([]option.ClientOption, error) {
	scopes := []string{"https://www.googleapis.com/auth/cloud-platform"}

	var creds *auth.Credentials
	if file := cmd.String("credential-file-override"); file != "" {
		c, err := credentials.DetectDefault(&credentials.DetectOptions{
			CredentialsFile: file,
			Scopes:          scopes,
		})
		if err != nil {
			return nil, fmt.Errorf("loading credentials from %q: %w", file, err)
		}
		creds = c
	} else {
		c, err := credentials.DetectDefault(&credentials.DetectOptions{Scopes: scopes})
		if err != nil {
			return nil, fmt.Errorf("loading default credentials: %w", err)
		}
		creds = c
	}

	if sa := cmd.String("impersonate-service-account"); sa != "" {
		c, err := impersonate.NewCredentials(&impersonate.CredentialsOptions{
			TargetPrincipal: sa,
			Scopes:          scopes,
			Credentials:     creds,
		})
		if err != nil {
			return nil, fmt.Errorf("impersonating %q: %w", sa, err)
		}
		creds = c
	}

	opts := []option.ClientOption{option.WithAuthCredentials(creds)}
	if bp := cmd.String("billing-project"); bp != "" {
		opts = append(opts, option.WithQuotaProject(bp))
	}
	return opts, nil
}
