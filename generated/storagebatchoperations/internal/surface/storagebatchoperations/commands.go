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

package storagebatchoperations

import (
	storagebatchoperations "cloud.google.com/go/storagebatchoperations/apiv1"
	"cloud.google.com/go/storagebatchoperations/apiv1/storagebatchoperationspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the storagebatchoperations command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "storagebatchoperations",
		Usage: "manage Storage Batch Operations API resources",
		Commands: []*cli.Command{
			{
				Name:  "bucket-operations",
				Usage: "Manage bucket-operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list bucket-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filters results as defined by https://google.", Required: false},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.ListBucketOperationsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListBucketOperations(ctx, req)
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
						Usage: "describe bucket-operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket-operation", Usage: "The ID of the bucket operation.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s/bucketOperations/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("bucket-operation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.GetBucketOperationRequest{
								Name: name,
							}

							resp, err := client.GetBucketOperation(ctx, req)
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
				Name:  "jobs",
				Usage: "Manage jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filters results as defined by https://google.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.ListJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListJobs(ctx, req)
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
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.GetJobRequest{
								Name: name,
							}

							resp, err := client.GetJob(ctx, req)
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
						Usage: "create jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-id", Usage: "The optional `job_id` for this Job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.CreateJobRequest{
								Parent:    parent,
								JobId:     cmd.String("job-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateJob(ctx, req)
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
						Name:  "delete",
						Usage: "delete jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any child bucket operations of the job will also.", Required: false},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.DeleteJobRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							if err := client.DeleteJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := storagebatchoperations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &storagebatchoperationspb.CancelJobRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							resp, err := client.CancelJob(ctx, req)
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
		},
	}
}
