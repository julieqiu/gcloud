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

package executions

import (
	executions "cloud.google.com/go/executions/apiv1"
	"cloud.google.com/go/executions/apiv1/executionspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the workflowexecutions command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "workflowexecutions",
		Usage: "manage Workflow Executions API resources",
		Commands: []*cli.Command{
			{
				Name:  "executions",
				Usage: "Manage executions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filters applied to the [Executions.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The ordering applied to the [Executions.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of executions to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListExecutions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "A view defining which fields should be filled in the returned.", Required: false},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := executions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &executionspb.ListExecutionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								View:      executionspb.ExecutionView(executionspb.ExecutionView_value[cmd.String("view")]),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListExecutions(ctx, req)
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
						Usage: "create executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := executions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &executionspb.CreateExecutionRequest{
								Parent: parent,
							}

							resp, err := client.CreateExecution(ctx, req)
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
						Usage: "describe executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "A view defining which fields should be filled in the returned.", Required: false},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := executions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &executionspb.GetExecutionRequest{
								Name: name,
								View: executionspb.ExecutionView(executionspb.ExecutionView_value[cmd.String("view")]),
							}

							resp, err := client.GetExecution(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The ID of the workflow.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := executions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &executionspb.CancelExecutionRequest{
								Name: name,
							}

							resp, err := client.CancelExecution(ctx, req)
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
