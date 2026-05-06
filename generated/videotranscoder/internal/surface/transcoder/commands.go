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

package transcoder

import (
	transcoder "cloud.google.com/go/transcoder/apiv1"
	"cloud.google.com/go/transcoder/apiv1/transcoderpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the transcoder command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "transcoder",
		Usage: "manage Transcoder API resources",
		Commands: []*cli.Command{
			{
				Name:  "job-templates",
				Usage: "Manage job-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create job-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-template-id", Usage: "The ID to use for the job template, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.CreateJobTemplateRequest{
								Parent:        parent,
								JobTemplateId: cmd.String("job-template-id"),
							}

							resp, err := client.CreateJobTemplate(ctx, req)
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
						Usage: "list job-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression, following the syntax outlined in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "One or more fields to compare and use to sort the output.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `next_page_token` value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.ListJobTemplatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListJobTemplates(ctx, req)
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
						Usage: "describe job-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-template", Usage: "The ID of the job template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("job-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.GetJobTemplateRequest{
								Name: name,
							}

							resp, err := client.GetJobTemplate(ctx, req)
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
						Usage: "delete job-templates",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the job template is not found, the request will succeed.", Required: false},
							&cli.StringFlag{Name: "job-template", Usage: "The ID of the job template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("job-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteJobTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.DeleteJobTemplateRequest{
								Name:         name,
								AllowMissing: cmd.Bool("allow-missing"),
							}

							if err := client.DeleteJobTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "create",
						Usage: "create jobs",
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
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.CreateJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateJob(ctx, req)
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
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression, following the syntax outlined in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "One or more fields to compare and use to sort the output.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `next_page_token` value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.ListJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
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
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.GetJobRequest{
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
						Name:  "delete",
						Usage: "delete jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the job is not found, the request will succeed but no.", Required: false},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							client, err := transcoder.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &transcoderpb.DeleteJobRequest{
								Name:         name,
								AllowMissing: cmd.Bool("allow-missing"),
							}

							if err := client.DeleteJob(ctx, req); err != nil {
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
