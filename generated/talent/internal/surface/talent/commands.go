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

package talent

import (
	talent "cloud.google.com/go/talent/apiv4"
	"cloud.google.com/go/talent/apiv4/talentpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the jobs command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "jobs",
		Usage: "manage Cloud Talent Solution API resources",
		Commands: []*cli.Command{
			{
				Name:  "client-events",
				Usage: "Manage client-events resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create client-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.CreateClientEventRequest{
								Parent: parent,
							}

							resp, err := client.CreateClientEvent(ctx, req)
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
				Name:  "companies",
				Usage: "Manage companies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.CreateCompanyRequest{
								Parent: parent,
							}

							resp, err := client.CreateCompany(ctx, req)
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
						Usage: "describe companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "companie", Usage: "The ID of the companie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("companie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.GetCompanyRequest{
								Name: name,
							}

							resp, err := client.GetCompany(ctx, req)
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
						Usage: "update companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "companie", Usage: "The ID of the companie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "company.name" not yet supported.
							company_name := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("companie"))
							fmt.Printf("Executing update on %s\n", company_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "companie", Usage: "The ID of the companie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("companie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCompany on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.DeleteCompanyRequest{
								Name: name,
							}

							if err := client.DeleteCompany(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list companies",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of companies to be returned, at most 100.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The starting indicator from which to return results.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "require-open-jobs", Usage: "Set to true if the companies requested must have open jobs.", Required: false},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.ListCompaniesRequest{
								Parent:          parent,
								PageToken:       cmd.String("page-token"),
								PageSize:        int32(cmd.Int("page-size")),
								RequireOpenJobs: cmd.Bool("require-open-jobs"),
							}

							limit := cmd.Int("limit")
							it := client.ListCompanies(ctx, req)
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
				Name:  "jobs",
				Usage: "Manage jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.CreateJobRequest{
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
						Name:  "batch-create",
						Usage: "batch-create jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.BatchCreateJobsRequest{
								Parent: parent,
							}

							op, err := client.BatchCreateJobs(ctx, req)
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
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/jobs/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.GetJobRequest{
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
						Name:  "update",
						Usage: "update jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "job.name" not yet supported.
							job_name := fmt.Sprintf("projects/%s/tenants/%s/jobs/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("job"))
							fmt.Printf("Executing update on %s\n", job_name)
							return nil
						},
					},

					{
						Name:  "batch-update",
						Usage: "batch-update jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.BatchUpdateJobsRequest{
								Parent: parent,
							}

							op, err := client.BatchUpdateJobs(ctx, req)
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
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/jobs/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.DeleteJobRequest{
								Name: name,
							}

							if err := client.DeleteJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "batch-delete",
						Usage: "batch-delete jobs",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the jobs to delete.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.BatchDeleteJobsRequest{
								Parent: parent,
								Names:  cmd.StringSlice("names"),
							}

							op, err := client.BatchDeleteJobs(ctx, req)
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
						Name:  "list",
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter string specifies the jobs to be enumerated.", Required: true},
							&cli.StringFlag{Name: "job-view", Usage: "The desired job attributes returned for jobs in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of jobs to be returned per page of results.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The starting point of a query result.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.ListJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								JobView:   talentpb.JobView(talentpb.JobView_value[cmd.String("job-view")]),
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
						Name:  "search",
						Usage: "search jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "disable-keyword-match", Usage: "This field is deprecated.", Required: false},
							&cli.StringFlag{Name: "diversification-level", Usage: "Controls whether highly similar jobs are returned next to each other in.", Required: false},
							&cli.BoolFlag{Name: "enable-broadening", Usage: "Controls whether to broaden the search when it produces sparse results.", Required: false},
							&cli.StringFlag{Name: "job-view", Usage: "The desired job attributes returned for jobs in the search response.", Required: false},
							&cli.StringFlag{Name: "keyword-match-mode", Usage: "Controls what keyword match options to use.", Required: false},
							&cli.IntFlag{Name: "max-page-size", Usage: "A limit on the number of jobs returned in the search results.", Required: false},
							&cli.IntFlag{Name: "offset", Usage: "An integer that specifies the current offset (that is, starting result.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "The criteria determining how search results are sorted.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The token specifying the current offset within.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "relevance-threshold", Usage: "The relevance threshold of the search results.", Required: false},
							&cli.StringFlag{Name: "search-mode", Usage: "Mode of a search.", Required: false},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.SearchJobsRequest{
								Parent:               parent,
								SearchMode:           talentpb.SearchJobsRequest_SearchMode(talentpb.SearchJobsRequest_SearchMode_value[cmd.String("search-mode")]),
								EnableBroadening:     cmd.Bool("enable-broadening"),
								JobView:              talentpb.JobView(talentpb.JobView_value[cmd.String("job-view")]),
								Offset:               int32(cmd.Int("offset")),
								MaxPageSize:          int32(cmd.Int("max-page-size")),
								PageToken:            cmd.String("page-token"),
								OrderBy:              cmd.String("order-by"),
								DiversificationLevel: talentpb.SearchJobsRequest_DiversificationLevel(talentpb.SearchJobsRequest_DiversificationLevel_value[cmd.String("diversification-level")]),
								DisableKeywordMatch:  cmd.Bool("disable-keyword-match"),
								KeywordMatchMode:     talentpb.SearchJobsRequest_KeywordMatchMode(talentpb.SearchJobsRequest_KeywordMatchMode_value[cmd.String("keyword-match-mode")]),
								RelevanceThreshold:   talentpb.SearchJobsRequest_RelevanceThreshold(talentpb.SearchJobsRequest_RelevanceThreshold_value[cmd.String("relevance-threshold")]),
							}

							resp, err := client.SearchJobs(ctx, req)
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
						Name:  "search-for-alert",
						Usage: "search-for-alert jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "disable-keyword-match", Usage: "This field is deprecated.", Required: false},
							&cli.StringFlag{Name: "diversification-level", Usage: "Controls whether highly similar jobs are returned next to each other in.", Required: false},
							&cli.BoolFlag{Name: "enable-broadening", Usage: "Controls whether to broaden the search when it produces sparse results.", Required: false},
							&cli.StringFlag{Name: "job-view", Usage: "The desired job attributes returned for jobs in the search response.", Required: false},
							&cli.StringFlag{Name: "keyword-match-mode", Usage: "Controls what keyword match options to use.", Required: false},
							&cli.IntFlag{Name: "max-page-size", Usage: "A limit on the number of jobs returned in the search results.", Required: false},
							&cli.IntFlag{Name: "offset", Usage: "An integer that specifies the current offset (that is, starting result.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "The criteria determining how search results are sorted.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The token specifying the current offset within.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "relevance-threshold", Usage: "The relevance threshold of the search results.", Required: false},
							&cli.StringFlag{Name: "search-mode", Usage: "Mode of a search.", Required: false},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.SearchJobsRequest{
								Parent:               parent,
								SearchMode:           talentpb.SearchJobsRequest_SearchMode(talentpb.SearchJobsRequest_SearchMode_value[cmd.String("search-mode")]),
								EnableBroadening:     cmd.Bool("enable-broadening"),
								JobView:              talentpb.JobView(talentpb.JobView_value[cmd.String("job-view")]),
								Offset:               int32(cmd.Int("offset")),
								MaxPageSize:          int32(cmd.Int("max-page-size")),
								PageToken:            cmd.String("page-token"),
								OrderBy:              cmd.String("order-by"),
								DiversificationLevel: talentpb.SearchJobsRequest_DiversificationLevel(talentpb.SearchJobsRequest_DiversificationLevel_value[cmd.String("diversification-level")]),
								DisableKeywordMatch:  cmd.Bool("disable-keyword-match"),
								KeywordMatchMode:     talentpb.SearchJobsRequest_KeywordMatchMode(talentpb.SearchJobsRequest_KeywordMatchMode_value[cmd.String("keyword-match-mode")]),
								RelevanceThreshold:   talentpb.SearchJobsRequest_RelevanceThreshold(talentpb.SearchJobsRequest_RelevanceThreshold_value[cmd.String("relevance-threshold")]),
							}

							resp, err := client.SearchJobsForAlert(ctx, req)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "tenants",
				Usage: "Manage tenants resources",
				Commands: []*cli.Command{

					{
						Name:  "complete-query",
						Usage: "complete-query tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "company", Usage: "If provided, restricts completion to specified company.", Required: false},
							&cli.StringSliceFlag{Name: "language-codes", Usage: "The list of languages of the query.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Completion result count.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "The query used to generate suggestions.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope of the completion.", Required: false},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The completion topic.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tenant := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							fmt.Printf("Executing complete-query on %s\n", tenant)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.CreateTenantRequest{
								Parent: parent,
							}

							resp, err := client.CreateTenant(ctx, req)
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
						Usage: "describe tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.GetTenantRequest{
								Name: name,
							}

							resp, err := client.GetTenant(ctx, req)
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
						Usage: "update tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tenant.name" not yet supported.
							tenant_name := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							fmt.Printf("Executing update on %s\n", tenant_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tenant", Usage: "The ID of the tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTenant on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.DeleteTenantRequest{
								Name: name,
							}

							if err := client.DeleteTenant(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list tenants",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of tenants to be returned, at most 100.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The starting indicator from which to return results.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := talent.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &talentpb.ListTenantsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListTenants(ctx, req)
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
		},
	}
}
