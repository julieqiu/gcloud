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

package secretmanager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the secretmanager command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "secretmanager",
		Usage: "manage Secret Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "secrets",
				Usage: "Manage secrets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter string, adhering to the rules in.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to be returned in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.ListSecretsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecrets(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret-id", Usage: "This must be unique within the project.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.CreateSecretRequest{
								Parent:   parent,
								SecretId: cmd.String("secret-id"),
							}

							resp, err := client.CreateSecret(ctx, req)
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
						Name:  "add-version",
						Usage: "add-version secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.AddSecretVersionRequest{
								Parent: parent,
							}

							resp, err := client.AddSecretVersion(ctx, req)
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
						Name:  "describe",
						Usage: "describe secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.GetSecretRequest{
								Name: name,
							}

							resp, err := client.GetSecret(ctx, req)
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
						Name:  "update",
						Usage: "update secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "secret.name" not yet supported.
							secret_name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							fmt.Printf("Executing update on %s\n", secret_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Etag of the [Secret][google.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSecret on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.DeleteSecretRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteSecret(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "describe",
						Usage: "describe secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions secrets",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
			{
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter string, adhering to the rules in.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to be returned in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.ListSecretVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecretVersions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.GetSecretVersionRequest{
								Name: name,
							}

							resp, err := client.GetSecretVersion(ctx, req)
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
						Name:  "access",
						Usage: "access versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.AccessSecretVersionRequest{
								Name: name,
							}

							resp, err := client.AccessSecretVersion(ctx, req)
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
						Name:  "disable",
						Usage: "disable versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Etag of the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.DisableSecretVersionRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.DisableSecretVersion(ctx, req)
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
						Name:  "enable",
						Usage: "enable versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Etag of the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.EnableSecretVersionRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.EnableSecretVersion(ctx, req)
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
						Name:  "destroy",
						Usage: "destroy versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Etag of the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "secret", Usage: "The ID of the secret.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := secretmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &secretmanagerpb.DestroySecretVersionRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.DestroySecretVersion(ctx, req)
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
