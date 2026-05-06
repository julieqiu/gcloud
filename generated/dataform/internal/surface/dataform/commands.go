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

package dataform

import (
	dataform "cloud.google.com/go/dataform/apiv1"
	"cloud.google.com/go/dataform/apiv1/dataformpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the dataform command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "dataform",
		Usage: "manage Dataform API resources",
		Commands: []*cli.Command{
			{
				Name:  "compilation-results",
				Usage: "Manage compilation-results resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list compilation-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "This field only supports ordering by `name` and `create_time`.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of compilation results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListCompilationResults`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ListCompilationResultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListCompilationResults(ctx, req)
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
						Usage: "describe compilation-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "compilation-result", Usage: "The ID of the compilation result.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/compilationResults/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("compilation-result"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetCompilationResultRequest{
								Name: name,
							}

							resp, err := client.GetCompilationResult(ctx, req)
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
						Usage: "create compilation-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateCompilationResultRequest{
								Parent: parent,
							}

							resp, err := client.CreateCompilationResult(ctx, req)
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
						Name:  "query",
						Usage: "query compilation-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "compilation-result", Usage: "The ID of the compilation result.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Optional filter for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of compilation results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/compilationResults/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("compilation-result"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.QueryCompilationResultActionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.QueryCompilationResultActions(ctx, req)
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
				},
			},
			{
				Name:  "config",
				Usage: "Manage config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetConfigRequest{
								Name: name,
							}

							resp, err := client.GetConfig(ctx, req)
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
						Usage: "update config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "config.name" not yet supported.
							config_name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "folders",
				Usage: "Manage folders resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetFolderRequest{
								Name: name,
							}

							resp, err := client.GetFolder(ctx, req)
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
						Usage: "create folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateFolderRequest{
								Parent: parent,
							}

							resp, err := client.CreateFolder(ctx, req)
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
						Usage: "update folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "folder.name" not yet supported.
							folder_name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
							fmt.Printf("Executing update on %s\n", folder_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteFolder on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteFolderRequest{
								Name: name,
							}

							if err := client.DeleteFolder(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "delete-tree",
						Usage: "delete-tree folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If `false` (default): The operation will fail if any.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFolderTree %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteFolderTreeRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteFolderTree(ctx, req)
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
						Name:  "query-folder-contents",
						Usage: "query-folder-contents folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Optional filtering for the returned list.", Required: false},
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to additionally sort results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of paths to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `QueryFolderContents` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							folder := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
							fmt.Printf("Executing query-folder-contents on %s\n", folder)
							return nil
						},
					},

					{
						Name:  "move",
						Usage: "move folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination-containing-folder", Usage: "The name of the Folder, TeamFolder, or root location to move the.", Required: false},
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("MoveFolder %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.MoveFolderRequest{
								Name:                        name,
								DestinationContainingFolder: runtime.Ptr(cmd.String("destination-containing-folder")),
							}

							op, err := client.MoveFolder(ctx, req)
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
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "query-user-root-contents",
						Usage: "query-user-root-contents locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Optional filtering for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to additionally sort results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of paths to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `QueryUserRootContents` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing query-user-root-contents on %s\n", location)
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
				Name:  "release-configs",
				Usage: "Manage release-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list release-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of release configs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListReleaseConfigs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ListReleaseConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListReleaseConfigs(ctx, req)
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
						Usage: "describe release-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-config", Usage: "The ID of the release config.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("release-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetReleaseConfigRequest{
								Name: name,
							}

							resp, err := client.GetReleaseConfig(ctx, req)
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
						Usage: "create release-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-config-id", Usage: "The ID to use for the release config, which will become the final.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateReleaseConfigRequest{
								Parent:          parent,
								ReleaseConfigId: cmd.String("release-config-id"),
							}

							resp, err := client.CreateReleaseConfig(ctx, req)
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
						Usage: "update release-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-config", Usage: "The ID of the release config.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "release_config.name" not yet supported.
							release_config_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("release-config"))
							fmt.Printf("Executing update on %s\n", release_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete release-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-config", Usage: "The ID of the release config.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("release-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteReleaseConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteReleaseConfigRequest{
								Name: name,
							}

							if err := client.DeleteReleaseConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "repositories",
				Usage: "Manage repositories resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "This field only supports ordering by `name`.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of repositories to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListRepositories` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ListRepositoriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListRepositories(ctx, req)
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
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetRepositoryRequest{
								Name: name,
							}

							resp, err := client.GetRepository(ctx, req)
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
						Usage: "create repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repository-id", Usage: "The ID to use for the repository, which will become the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateRepositoryRequest{
								Parent:       parent,
								RepositoryId: cmd.String("repository-id"),
							}

							resp, err := client.CreateRepository(ctx, req)
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
						Usage: "update repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "repository.name" not yet supported.
							repository_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing update on %s\n", repository_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete repositories",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, child resources of this repository (compilation.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteRepository on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteRepositoryRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteRepository(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "move",
						Usage: "move repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination-containing-folder", Usage: "The name of the Folder, TeamFolder, or root location to move the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("MoveRepository %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.MoveRepositoryRequest{
								Name:                        name,
								DestinationContainingFolder: runtime.Ptr(cmd.String("destination-containing-folder")),
							}

							op, err := client.MoveRepository(ctx, req)
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
						Name:  "commit",
						Usage: "commit repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "required-head-commit-sha", Usage: "The commit SHA which must be the repository's current HEAD before.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CommitRepositoryChangesRequest{
								Name:                  name,
								RequiredHeadCommitSha: cmd.String("required-head-commit-sha"),
							}

							resp, err := client.CommitRepositoryChanges(ctx, req)
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
						Name:  "read-file",
						Usage: "read-file repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commit-sha", Usage: "The commit SHA for the commit to read from.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "Full file path to read including filename, from repository root.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ReadRepositoryFileRequest{
								Name:      name,
								CommitSha: cmd.String("commit-sha"),
								Path:      cmd.String("path"),
							}

							resp, err := client.ReadRepositoryFile(ctx, req)
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
						Name:  "query-directory-contents",
						Usage: "query-directory-contents repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commit-sha", Usage: "The Commit SHA for the commit to query from.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of paths to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "path", Usage: "The directory's full path including directory name, relative to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.QueryRepositoryDirectoryContentsRequest{
								Name:      name,
								CommitSha: cmd.String("commit-sha"),
								Path:      cmd.String("path"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.QueryRepositoryDirectoryContents(ctx, req)
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
						Name:  "fetch-history",
						Usage: "fetch-history repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of commits to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `FetchRepositoryHistory`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.FetchRepositoryHistoryRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.FetchRepositoryHistory(ctx, req)
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
						Name:  "compute-access-token-status",
						Usage: "compute-access-token-status repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ComputeRepositoryAccessTokenStatusRequest{
								Name: name,
							}

							resp, err := client.ComputeRepositoryAccessTokenStatus(ctx, req)
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
						Name:  "fetch-remote-branches",
						Usage: "fetch-remote-branches repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.FetchRemoteBranchesRequest{
								Name: name,
							}

							resp, err := client.FetchRemoteBranches(ctx, req)
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
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetIamPolicyRequest{
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.SetIamPolicyRequest{
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.TestIamPermissionsRequest{
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
				Name:  "team-folders",
				Usage: "Manage team-folders resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe team-folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "team-folder", Usage: "The ID of the team folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team-folder"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetTeamFolderRequest{
								Name: name,
							}

							resp, err := client.GetTeamFolder(ctx, req)
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
						Usage: "create team-folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateTeamFolderRequest{
								Parent: parent,
							}

							resp, err := client.CreateTeamFolder(ctx, req)
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
						Usage: "update team-folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "team-folder", Usage: "The ID of the team folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "team_folder.name" not yet supported.
							team_folder_name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team-folder"))
							fmt.Printf("Executing update on %s\n", team_folder_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete team-folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "team-folder", Usage: "The ID of the team folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team-folder"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTeamFolder on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteTeamFolderRequest{
								Name: name,
							}

							if err := client.DeleteTeamFolder(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "delete-tree",
						Usage: "delete-tree team-folders",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If `false` (default): The operation will fail if any.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "team-folder", Usage: "The ID of the team folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team-folder"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTeamFolderTree %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteTeamFolderTreeRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteTeamFolderTree(ctx, req)
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
						Name:  "query-contents",
						Usage: "query-contents team-folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Optional filtering for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to additionally sort results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of paths to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `QueryTeamFolderContents`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "team-folder", Usage: "The ID of the team folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							team_folder := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team-folder"))
							fmt.Printf("Executing query-contents on %s\n", team_folder)
							return nil
						},
					},

					{
						Name:  "search",
						Usage: "search team-folders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Optional filtering for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to additionally sort results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of TeamFolders to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `SearchTeamFolders` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing search on %s\n", location)
							return nil
						},
					},
				},
			},
			{
				Name:  "workflow-configs",
				Usage: "Manage workflow-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list workflow-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of workflow configs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListWorkflowConfigs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ListWorkflowConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkflowConfigs(ctx, req)
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
						Usage: "describe workflow-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-config", Usage: "The ID of the workflow config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetWorkflowConfigRequest{
								Name: name,
							}

							resp, err := client.GetWorkflowConfig(ctx, req)
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
						Usage: "create workflow-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-config-id", Usage: "The ID to use for the workflow config, which will become the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateWorkflowConfigRequest{
								Parent:           parent,
								WorkflowConfigId: cmd.String("workflow-config-id"),
							}

							resp, err := client.CreateWorkflowConfig(ctx, req)
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
						Usage: "update workflow-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-config", Usage: "The ID of the workflow config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "workflow_config.name" not yet supported.
							workflow_config_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-config"))
							fmt.Printf("Executing update on %s\n", workflow_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete workflow-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-config", Usage: "The ID of the workflow config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteWorkflowConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteWorkflowConfigRequest{
								Name: name,
							}

							if err := client.DeleteWorkflowConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "workflow-invocations",
				Usage: "Manage workflow-invocations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list workflow-invocations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "This field only supports ordering by `name`.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of workflow invocations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListWorkflowInvocations`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ListWorkflowInvocationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkflowInvocations(ctx, req)
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
						Usage: "describe workflow-invocations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-invocation", Usage: "The ID of the workflow invocation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-invocation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetWorkflowInvocationRequest{
								Name: name,
							}

							resp, err := client.GetWorkflowInvocation(ctx, req)
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
						Usage: "create workflow-invocations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateWorkflowInvocationRequest{
								Parent: parent,
							}

							resp, err := client.CreateWorkflowInvocation(ctx, req)
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
						Usage: "delete workflow-invocations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-invocation", Usage: "The ID of the workflow invocation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-invocation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteWorkflowInvocation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteWorkflowInvocationRequest{
								Name: name,
							}

							if err := client.DeleteWorkflowInvocation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel workflow-invocations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-invocation", Usage: "The ID of the workflow invocation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-invocation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CancelWorkflowInvocationRequest{
								Name: name,
							}

							resp, err := client.CancelWorkflowInvocation(ctx, req)
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
						Name:  "query",
						Usage: "query workflow-invocations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of workflow invocations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workflow-invocation", Usage: "The ID of the workflow invocation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workflow-invocation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.QueryWorkflowInvocationActionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.QueryWorkflowInvocationActions(ctx, req)
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
				},
			},
			{
				Name:  "workspaces",
				Usage: "Manage workspaces resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "This field only supports ordering by `name`.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of workspaces to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListWorkspaces` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ListWorkspacesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkspaces(ctx, req)
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
						Usage: "describe workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.GetWorkspaceRequest{
								Name: name,
							}

							resp, err := client.GetWorkspace(ctx, req)
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
						Usage: "create workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace-id", Usage: "The ID to use for the workspace, which will become the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CreateWorkspaceRequest{
								Parent:      parent,
								WorkspaceId: cmd.String("workspace-id"),
							}

							resp, err := client.CreateWorkspace(ctx, req)
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
						Usage: "delete workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteWorkspace on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.DeleteWorkspaceRequest{
								Name: name,
							}

							if err := client.DeleteWorkspace(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "install-npm-packages",
						Usage: "install-npm-packages workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing install-npm-packages on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "pull",
						Usage: "pull workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "remote-branch", Usage: "The name of the branch in the Git remote from which to pull.", Required: false},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.PullGitCommitsRequest{
								Name:         name,
								RemoteBranch: cmd.String("remote-branch"),
							}

							resp, err := client.PullGitCommits(ctx, req)
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
						Name:  "push",
						Usage: "push workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "remote-branch", Usage: "The name of the branch in the Git remote to which commits should.", Required: false},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.PushGitCommitsRequest{
								Name:         name,
								RemoteBranch: cmd.String("remote-branch"),
							}

							resp, err := client.PushGitCommits(ctx, req)
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
						Name:  "fetch-file-git-statuses",
						Usage: "fetch-file-git-statuses workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.FetchFileGitStatusesRequest{
								Name: name,
							}

							resp, err := client.FetchFileGitStatuses(ctx, req)
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
						Name:  "fetch-git-ahead-behind",
						Usage: "fetch-git-ahead-behind workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "remote-branch", Usage: "The name of the branch in the Git remote against which this.", Required: false},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.FetchGitAheadBehindRequest{
								Name:         name,
								RemoteBranch: cmd.String("remote-branch"),
							}

							resp, err := client.FetchGitAheadBehind(ctx, req)
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
						Usage: "commit workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commit-message", Usage: "The commit's message.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "paths", Usage: "Full file paths to commit including filename, rooted at workspace.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.CommitWorkspaceChangesRequest{
								Name:          name,
								CommitMessage: cmd.String("commit-message"),
								Paths:         cmd.StringSlice("paths"),
							}

							resp, err := client.CommitWorkspaceChanges(ctx, req)
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
						Name:  "reset",
						Usage: "reset workspaces",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "clean", Usage: "If set to true, untracked files will be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "paths", Usage: "Full file paths to reset back to their committed state including.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataformpb.ResetWorkspaceChangesRequest{
								Name:  name,
								Paths: cmd.StringSlice("paths"),
								Clean: cmd.Bool("clean"),
							}

							resp, err := client.ResetWorkspaceChanges(ctx, req)
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
						Name:  "fetch-file-diff",
						Usage: "fetch-file-diff workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The file's full path including filename, relative to the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing fetch-file-diff on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "query-directory-contents",
						Usage: "query-directory-contents workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of paths to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `QueryDirectoryContents`.", Required: false},
							&cli.StringFlag{Name: "path", Usage: "The directory's full path including directory name, relative to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specifies the metadata to return for each directory entry.", Required: false},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing query-directory-contents on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "search-files",
						Usage: "search-files workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Optional filter for the returned list in filtering format.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of search results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `SearchFilesRequest`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing search-files on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "make-directory",
						Usage: "make-directory workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The directory's full path including directory name, relative to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing make-directory on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "remove-directory",
						Usage: "remove-directory workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The directory's full path including directory name, relative to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing remove-directory on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "move-directory",
						Usage: "move-directory workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-path", Usage: "The new path for the directory including directory name, rooted.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The directory's full path including directory name, relative to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing move-directory on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "read-file",
						Usage: "read-file workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The file's full path including filename, relative to the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The Git revision of the file to return.", Required: false},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing read-file on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "remove-file",
						Usage: "remove-file workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The file's full path including filename, relative to the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing remove-file on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "move-file",
						Usage: "move-file workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-path", Usage: "The file's new path including filename, relative to the workspace.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The file's full path including filename, relative to the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing move-file on %s\n", workspace)
							return nil
						},
					},

					{
						Name:  "write-file",
						Usage: "write-file workspaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contents", Usage: "The file's contents.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "path", Usage: "The file.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "workspace", Usage: "The ID of the workspace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workspace := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("workspace"))
							fmt.Printf("Executing write-file on %s\n", workspace)
							return nil
						},
					},
				},
			},
		},
	}
}
