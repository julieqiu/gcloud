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

package storagetransfer

import (
	storagetransfer "cloud.google.com/go/storagetransfer/apiv1"
	"cloud.google.com/go/storagetransfer/apiv1/storagetransferpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the storagetransfer command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "storagetransfer",
		Usage: "manage Storage Transfer API resources",
		Commands: []*cli.Command{
			{
				Name:  "agent-pools",
				Usage: "Manage agent-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create agent-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent-pool-id", Usage: "The ID of the agent pool to create.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing create on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update agent-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent-pool", Usage: "The ID of the agent pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "agent_pool.name" not yet supported.
							agent_pool_name := fmt.Sprintf("projects/%s/agentPools/%s", cmd.String("project"), cmd.String("agent-pool"))
							fmt.Printf("Executing update on %s\n", agent_pool_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe agent-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent-pool", Usage: "The ID of the agent pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agentPools/%s", cmd.String("project"), cmd.String("agent-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagetransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagetransferpb.GetAgentPoolRequest{
								Name: name,
							}

							resp, err := client.GetAgentPool(ctx, req)
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
						Usage: "list agent-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An optional list of query parameters specified as JSON text in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The list page token.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete agent-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent-pool", Usage: "The ID of the agent pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agentPools/%s", cmd.String("project"), cmd.String("agent-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAgentPool on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagetransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagetransferpb.DeleteAgentPoolRequest{
								Name: name,
							}

							if err := client.DeleteAgentPool(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "google-service-accounts",
				Usage: "Manage google-service-accounts resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe google-service-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "transfer-jobs",
				Usage: "Manage transfer-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create transfer-jobs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagetransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagetransferpb.CreateTransferJobRequest{}

							resp, err := client.CreateTransferJob(ctx, req)
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
						Usage: "update transfer-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-name", Usage: "The name of job to update.", Required: true},
							&cli.StringFlag{Name: "project-id", Usage: "The ID of the Google Cloud project that owns the.", Required: true},
							&cli.StringFlag{Name: "transfer-job", Usage: "The ID of the transfer job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							job_name := fmt.Sprintf("transferJobs/%s", cmd.String("transfer-job"))
							fmt.Printf("Executing update on %s\n", job_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe transfer-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-name", Usage: "The job to get.", Required: true},
							&cli.StringFlag{Name: "project-id", Usage: "The ID of the Google Cloud project that owns the.", Required: true},
							&cli.StringFlag{Name: "transfer-job", Usage: "The ID of the transfer job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							job_name := fmt.Sprintf("transferJobs/%s", cmd.String("transfer-job"))
							fmt.Printf("Executing describe on %s\n", job_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list transfer-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A list of query parameters specified as JSON text in the form of:.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The list page token.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagetransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagetransferpb.ListTransferJobsRequest{
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTransferJobs(ctx, req)
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
						Name:  "run",
						Usage: "run transfer-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-name", Usage: "The name of the transfer job.", Required: true},
							&cli.StringFlag{Name: "project-id", Usage: "The ID of the Google Cloud project that owns the transfer.", Required: true},
							&cli.StringFlag{Name: "transfer-job", Usage: "The ID of the transfer job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							job_name := fmt.Sprintf("transferJobs/%s", cmd.String("transfer-job"))
							fmt.Printf("Executing run on %s\n", job_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete transfer-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-name", Usage: "The job to delete.", Required: true},
							&cli.StringFlag{Name: "project-id", Usage: "The ID of the Google Cloud project that owns the.", Required: true},
							&cli.StringFlag{Name: "transfer-job", Usage: "The ID of the transfer job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							job_name := fmt.Sprintf("transferJobs/%s", cmd.String("transfer-job"))
							fmt.Printf("Executing delete on %s\n", job_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "transfer-operations",
				Usage: "Manage transfer-operations resources",
				Commands: []*cli.Command{

					{
						Name:  "pause",
						Usage: "pause transfer-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer-operation", Usage: "The ID of the transfer operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("transferOperations/%s", cmd.String("transfer-operation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute PauseTransferOperation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagetransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagetransferpb.PauseTransferOperationRequest{
								Name: name,
							}

							if err := client.PauseTransferOperation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "resume",
						Usage: "resume transfer-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer-operation", Usage: "The ID of the transfer operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("transferOperations/%s", cmd.String("transfer-operation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute ResumeTransferOperation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagetransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagetransferpb.ResumeTransferOperationRequest{
								Name: name,
							}

							if err := client.ResumeTransferOperation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list transfer-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe transfer-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer-operation", Usage: "The ID of the transfer operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("transferOperations/%s", cmd.String("transfer-operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel transfer-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer-operation", Usage: "The ID of the transfer operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("transferOperations/%s", cmd.String("transfer-operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
