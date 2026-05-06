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

package support

import (
	support "cloud.google.com/go/support/apiv2"
	"cloud.google.com/go/support/apiv2/supportpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudsupport command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudsupport",
		Usage: "manage Google Cloud Support API resources",
		Commands: []*cli.Command{
			{
				Name:  "attachments",
				Usage: "Manage attachments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of attachments fetched with each request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.ListAttachmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAttachments(ctx, req)
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
				Name:  "case-classifications",
				Usage: "Manage case-classifications resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search case-classifications",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of classifications fetched with each request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "query", Usage: "An expression used to filter case classifications.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.SearchCaseClassificationsRequest{
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchCaseClassifications(ctx, req)
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
				Name:  "cases",
				Usage: "Manage cases resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.GetCaseRequest{
								Name: name,
							}

							resp, err := client.GetCase(ctx, req)
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
						Usage: "list cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression used to filter cases.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of cases fetched with each request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.ListCasesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCases(ctx, req)
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
						Usage: "search cases",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of cases fetched with each request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "An expression used to filter cases.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.SearchCasesRequest{
								Parent:    parent,
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchCases(ctx, req)
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
						Usage: "create cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.CreateCaseRequest{
								Parent: parent,
							}

							resp, err := client.CreateCase(ctx, req)
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
						Usage: "update cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "case.name" not yet supported.
							case_name := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							fmt.Printf("Executing update on %s\n", case_name)
							return nil
						},
					},

					{
						Name:  "escalate",
						Usage: "escalate cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.EscalateCaseRequest{
								Name: name,
							}

							resp, err := client.EscalateCase(ctx, req)
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
						Name:  "close",
						Usage: "close cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.CloseCaseRequest{
								Name: name,
							}

							resp, err := client.CloseCase(ctx, req)
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
				Name:  "comments",
				Usage: "Manage comments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of comments to fetch.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.ListCommentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListComments(ctx, req)
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
						Usage: "create comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "case", Usage: "The ID of the case.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/cases/%s", cmd.String("project"), cmd.String("case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := support.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &supportpb.CreateCommentRequest{
								Parent: parent,
							}

							resp, err := client.CreateComment(ctx, req)
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
