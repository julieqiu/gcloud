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

package credentials

import (
	credentials "cloud.google.com/go/credentials/apiv1"
	"cloud.google.com/go/credentials/apiv1/credentialspb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the iamcredentials command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "iamcredentials",
		Usage: "manage IAM Service Account Credentials API resources",
		Commands: []*cli.Command{
			{
				Name:  "service-accounts",
				Usage: "Manage service-accounts resources",
				Commands: []*cli.Command{

					{
						Name:  "generate-access-token",
						Usage: "generate-access-token service-accounts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "delegates", Usage: "The sequence of service accounts in a delegation chain.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "scope", Usage: "Code to identify the scopes to be included in the OAuth 2.", Required: true},
							&cli.StringFlag{Name: "service-account", Usage: "The ID of the service account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := credentials.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &credentialspb.GenerateAccessTokenRequest{
								Name:      name,
								Delegates: cmd.StringSlice("delegates"),
								Scope:     cmd.StringSlice("scope"),
							}

							resp, err := client.GenerateAccessToken(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "generate-id-token",
						Usage: "generate-id-token service-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "audience", Usage: "The audience for the token, such as the API or account that this token.", Required: true},
							&cli.StringSliceFlag{Name: "delegates", Usage: "The sequence of service accounts in a delegation chain.", Required: false},
							&cli.BoolFlag{Name: "include-email", Usage: "Include the service account email in the token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-account", Usage: "The ID of the service account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := credentials.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &credentialspb.GenerateIdTokenRequest{
								Name:         name,
								Delegates:    cmd.StringSlice("delegates"),
								Audience:     cmd.String("audience"),
								IncludeEmail: cmd.Bool("include-email"),
							}

							resp, err := client.GenerateIdToken(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "sign-blob",
						Usage: "sign-blob service-accounts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "delegates", Usage: "The sequence of service accounts in a delegation chain.", Required: false},
							&cli.StringFlag{Name: "payload", Usage: "The bytes to sign.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-account", Usage: "The ID of the service account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := credentials.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &credentialspb.SignBlobRequest{
								Name:      name,
								Delegates: cmd.StringSlice("delegates"),
								Payload:   []byte(cmd.String("payload")),
							}

							resp, err := client.SignBlob(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "sign-jwt",
						Usage: "sign-jwt service-accounts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "delegates", Usage: "The sequence of service accounts in a delegation chain.", Required: false},
							&cli.StringFlag{Name: "payload", Usage: "The JWT payload to sign: a JSON object that contains a JWT Claims Set.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-account", Usage: "The ID of the service account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := credentials.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &credentialspb.SignJwtRequest{
								Name:      name,
								Delegates: cmd.StringSlice("delegates"),
								Payload:   cmd.String("payload"),
							}

							resp, err := client.SignJwt(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
		},
	}
}
