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

package migration

import (
	migration "cloud.google.com/go/migration/apiv2"
	"cloud.google.com/go/migration/apiv2/migrationpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the bigquerymigration command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigquerymigration",
		Usage: "manage BigQuery Migration API resources",
		Commands: []*cli.Command{
			{
				Name:  "subtasks",
				Usage: "Manage subtasks resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe subtasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subtask", Usage: "The ID of the subtask.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s/subtasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"), cmd.String("subtask"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.GetMigrationSubtaskRequest{
								Name: name,
							}

							resp, err := client.GetMigrationSubtask(ctx, req)
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
						Usage: "list subtasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter to apply.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of migration tasks to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from previous `ListMigrationSubtasks`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.ListMigrationSubtasksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMigrationSubtasks(ctx, req)
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
				Name:  "workflows",
				Usage: "Manage workflows resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create workflows",
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
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.CreateMigrationWorkflowRequest{
								Parent: parent,
							}

							resp, err := client.CreateMigrationWorkflow(ctx, req)
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
						Usage: "describe workflows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.GetMigrationWorkflowRequest{
								Name: name,
							}

							resp, err := client.GetMigrationWorkflow(ctx, req)
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
						Usage: "list workflows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of migration workflows to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from previous `ListMigrationWorkflows` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.ListMigrationWorkflowsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMigrationWorkflows(ctx, req)
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
						Usage: "delete workflows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteMigrationWorkflow on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.DeleteMigrationWorkflowRequest{
								Name: name,
							}

							if err := client.DeleteMigrationWorkflow(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "start",
						Usage: "start workflows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute StartMigrationWorkflow on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := migration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &migrationpb.StartMigrationWorkflowRequest{
								Name: name,
							}

							if err := client.StartMigrationWorkflow(ctx, req); err != nil {
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
