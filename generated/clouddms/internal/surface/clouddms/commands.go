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

package clouddms

import (
	clouddms "cloud.google.com/go/clouddms/apiv1"
	"cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the datamigration command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "datamigration",
		Usage: "manage Database Migration API resources",
		Commands: []*cli.Command{
			{
				Name:  "connection-profiles",
				Usage: "Manage connection-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters connection profiles listed in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order results according to.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of connection profiles to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListConnectionProfiles` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ListConnectionProfilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListConnectionProfiles(ctx, req)
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
						Usage: "describe connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.GetConnectionProfileRequest{
								Name: name,
							}

							resp, err := client.GetConnectionProfile(ctx, req)
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
						Name:  "create",
						Usage: "create connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile-id", Usage: "The connection profile identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.BoolFlag{Name: "skip-validation", Usage: "Create the connection profile without validating it.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the connection profile, but don't create any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.CreateConnectionProfileRequest{
								Parent:              parent,
								ConnectionProfileId: cmd.String("connection-profile-id"),
								RequestId:           cmd.String("request-id"),
								ValidateOnly:        cmd.Bool("validate-only"),
								SkipValidation:      cmd.Bool("skip-validation"),
							}

							op, err := client.CreateConnectionProfile(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.BoolFlag{Name: "skip-validation", Usage: "Update the connection profile without validating it.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the connection profile, but don't update any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "connection_profile.name" not yet supported.
							connection_profile_name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							fmt.Printf("Executing update on %s\n", connection_profile_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "In case of force delete, the CloudSQL replica database is also deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConnectionProfile %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.DeleteConnectionProfileRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteConnectionProfile(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "conversion-workspaces",
				Usage: "Manage conversion-workspaces resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.GetConversionWorkspaceRequest{
								Name: name,
							}

							resp, err := client.GetConversionWorkspace(ctx, req)
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
						Name:  "list",
						Usage: "list conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters conversion workspaces listed in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of conversion workspaces to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The nextPageToken value received in the previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ListConversionWorkspacesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListConversionWorkspaces(ctx, req)
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
						Usage: "create conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace-id", Usage: "The ID of the conversion workspace to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.CreateConversionWorkspaceRequest{
								Parent:                parent,
								ConversionWorkspaceId: cmd.String("conversion-workspace-id"),
								RequestId:             cmd.String("request-id"),
							}

							op, err := client.CreateConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "conversion_workspace.name" not yet supported.
							conversion_workspace_name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							fmt.Printf("Executing update on %s\n", conversion_workspace_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Force delete the conversion workspace, even if there's a running migration.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConversionWorkspace %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.DeleteConversionWorkspaceRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "seed",
						Usage: "seed conversion-workspaces",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-commit", Usage: "Should the conversion workspace be committed automatically after the.", Required: false},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.SeedConversionWorkspaceRequest{
								Name:       name,
								AutoCommit: cmd.Bool("auto-commit"),
							}

							op, err := client.SeedConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "convert",
						Usage: "convert conversion-workspaces",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-commit", Usage: "Specifies whether the conversion workspace is to be committed.", Required: false},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.BoolFlag{Name: "convert-full-path", Usage: "Automatically convert the full entity path for each entity.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Filter the entities to convert.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ConvertConversionWorkspaceRequest{
								Name:            name,
								AutoCommit:      cmd.Bool("auto-commit"),
								Filter:          cmd.String("filter"),
								ConvertFullPath: cmd.Bool("convert-full-path"),
							}

							op, err := client.ConvertConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "commit",
						Usage: "commit conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commit-name", Usage: "Optional name of the commit.", Required: false},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.CommitConversionWorkspaceRequest{
								Name:       name,
								CommitName: cmd.String("commit-name"),
							}

							op, err := client.CommitConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "rollback",
						Usage: "rollback conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.RollbackConversionWorkspaceRequest{
								Name: name,
							}

							op, err := client.RollbackConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "apply",
						Usage: "apply conversion-workspaces",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-commit", Usage: "Specifies whether the conversion workspace is to be committed.", Required: false},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.BoolFlag{Name: "dry-run", Usage: "Only validates the apply process, but doesn't change the.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Filter which entities to apply.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ApplyConversionWorkspaceRequest{
								Name:       name,
								Filter:     cmd.String("filter"),
								DryRun:     cmd.Bool("dry-run"),
								AutoCommit: cmd.Bool("auto-commit"),
							}

							op, err := client.ApplyConversionWorkspace(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "describe-database-entities",
						Usage: "describe-database-entities conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commit-id", Usage: "Request a specific commit ID.", Required: false},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter the returned entities based on AIP-160 standard.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of entities to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The nextPageToken value received in the previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tree", Usage: "The tree to fetch.", Required: true},
							&cli.BoolFlag{Name: "uncommitted", Usage: "Whether to retrieve the latest committed version of the entities.", Required: false},
							&cli.StringFlag{Name: "view", Usage: "Results view based on AIP-157.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversion_workspace := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							fmt.Printf("Executing describe-database-entities on %s\n", conversion_workspace)
							return nil
						},
					},

					{
						Name:  "search-background-jobs",
						Usage: "search-background-jobs conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-size", Usage: "The maximum number of jobs to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-most-recent-per-job-type", Usage: "Whether or not to return just the most recent job per job type,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversion_workspace := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							fmt.Printf("Executing search-background-jobs on %s\n", conversion_workspace)
							return nil
						},
					},

					{
						Name:  "describe-conversion-workspace-revisions",
						Usage: "describe-conversion-workspace-revisions conversion-workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commit-id", Usage: "Optional filter to request a specific commit ID.", Required: false},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversion_workspace := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							fmt.Printf("Executing describe-conversion-workspace-revisions on %s\n", conversion_workspace)
							return nil
						},
					},
				},
			},
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch-static-ips",
						Usage: "fetch-static-ips locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of IPs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `FetchStaticIps` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.FetchStaticIpsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.FetchStaticIps(ctx, req)
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
				Name:  "mapping-rules",
				Usage: "Manage mapping-rules resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create mapping-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mapping-rule-id", Usage: "The ID of the rule to create.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.CreateMappingRuleRequest{
								Parent:        parent,
								MappingRuleId: cmd.String("mapping-rule-id"),
								RequestId:     cmd.String("request-id"),
							}

							resp, err := client.CreateMappingRule(ctx, req)
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
						Name:  "delete",
						Usage: "delete mapping-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mapping-rule", Usage: "The ID of the mapping rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s/mappingRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"), cmd.String("mapping-rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteMappingRule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.DeleteMappingRuleRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							if err := client.DeleteMappingRule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list mapping-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of rules to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The nextPageToken value received in the previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ListMappingRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMappingRules(ctx, req)
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
						Usage: "describe mapping-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mapping-rule", Usage: "The ID of the mapping rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s/mappingRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"), cmd.String("mapping-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.GetMappingRuleRequest{
								Name: name,
							}

							resp, err := client.GetMappingRule(ctx, req)
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
						Name:  "import",
						Usage: "import mapping-rules",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-commit", Usage: "Should the conversion workspace be committed automatically after.", Required: true},
							&cli.StringFlag{Name: "conversion-workspace", Usage: "The ID of the conversion workspace.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rules-format", Usage: "The format of the rules content file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion-workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ImportMappingRulesRequest{
								Parent:      parent,
								RulesFormat: clouddmspb.ImportRulesFileFormat(clouddmspb.ImportRulesFileFormat_value[cmd.String("rules-format")]),
								AutoCommit:  cmd.Bool("auto-commit"),
							}

							op, err := client.ImportMappingRules(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "migration-jobs",
				Usage: "Manage migration-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters migration jobs listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results based on the migration job name.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of migration jobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The nextPageToken value received in the previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ListMigrationJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMigrationJobs(ctx, req)
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
						Usage: "describe migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.GetMigrationJobRequest{
								Name: name,
							}

							resp, err := client.GetMigrationJob(ctx, req)
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
						Name:  "create",
						Usage: "create migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job-id", Usage: "The ID of the instance to create.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.CreateMigrationJobRequest{
								Parent:         parent,
								MigrationJobId: cmd.String("migration-job-id"),
								RequestId:      cmd.String("request-id"),
							}

							op, err := client.CreateMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "migration_job.name" not yet supported.
							migration_job_name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							fmt.Printf("Executing update on %s\n", migration_job_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete migration-jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "The destination CloudSQL connection profile is always deleted with the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMigrationJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.DeleteMigrationJobRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "start",
						Usage: "start migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-validation", Usage: "Start the migration job without running prior configuration.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.StartMigrationJobRequest{
								Name:           name,
								SkipValidation: cmd.Bool("skip-validation"),
							}

							op, err := client.StartMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "stop",
						Usage: "stop migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.StopMigrationJobRequest{
								Name: name,
							}

							op, err := client.StopMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "resume",
						Usage: "resume migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ResumeMigrationJobRequest{
								Name: name,
							}

							op, err := client.ResumeMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "promote",
						Usage: "promote migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.PromoteMigrationJobRequest{
								Name: name,
							}

							op, err := client.PromoteMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "verify",
						Usage: "verify migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.VerifyMigrationJobRequest{
								Name: name,
							}

							op, err := client.VerifyMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "restart",
						Usage: "restart migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-validation", Usage: "Restart the migration job without running prior configuration.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.RestartMigrationJobRequest{
								Name:           name,
								SkipValidation: cmd.Bool("skip-validation"),
							}

							op, err := client.RestartMigrationJob(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "generate-ssh-script",
						Usage: "generate-ssh-script migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "vm", Usage: "Bastion VM Instance name to use or to create.", Required: true},
							&cli.IntFlag{Name: "vm-port", Usage: "The port that will be open on the bastion host.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migration_job := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							fmt.Printf("Executing generate-ssh-script on %s\n", migration_job)
							return nil
						},
					},

					{
						Name:  "generate-tcp-proxy-script",
						Usage: "generate-tcp-proxy-script migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "vm-machine-type", Usage: "The type of the Compute instance that will host the proxy.", Required: true},
							&cli.StringFlag{Name: "vm-name", Usage: "The name of the Compute instance that will host the proxy.", Required: true},
							&cli.StringFlag{Name: "vm-subnet", Usage: "The name of the subnet the Compute instance will use for private.", Required: true},
							&cli.StringFlag{Name: "vm-zone", Usage: "The Google Cloud Platform zone to create the VM in.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migration_job := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							fmt.Printf("Executing generate-tcp-proxy-script on %s\n", migration_job)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migration-job", Usage: "The ID of the migration job.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration-job"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "private-connections",
				Usage: "Manage private-connections resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection-id", Usage: "The private connection identifier.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.BoolFlag{Name: "skip-validation", Usage: "If set to true, will skip validations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.CreatePrivateConnectionRequest{
								Parent:              parent,
								PrivateConnectionId: cmd.String("private-connection-id"),
								RequestId:           cmd.String("request-id"),
								SkipValidation:      cmd.Bool("skip-validation"),
							}

							op, err := client.CreatePrivateConnection(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "describe private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.GetPrivateConnectionRequest{
								Name: name,
							}

							resp, err := client.GetPrivateConnection(ctx, req)
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
						Name:  "list",
						Usage: "list private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters private connections listed in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of private connections to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListPrivateConnections` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.ListPrivateConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPrivateConnections(ctx, req)
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
						Name:  "delete",
						Usage: "delete private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePrivateConnection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := clouddms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &clouddmspb.DeletePrivateConnectionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeletePrivateConnection(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
		},
	}
}
