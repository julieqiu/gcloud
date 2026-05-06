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

package contactcenterinsights

import (
	contactcenterinsights "cloud.google.com/go/contactcenterinsights/apiv1"
	"cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the contactcenterinsights command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "contactcenterinsights",
		Usage: "manage Contact Center AI Insights API resources",
		Commands: []*cli.Command{
			{
				Name:  "analyses",
				Usage: "Manage analyses resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateAnalysisRequest{
								Parent: parent,
							}

							op, err := client.CreateAnalysis(ctx, req)
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
						Usage: "describe analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analyse", Usage: "The ID of the analyse.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("analyse"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetAnalysisRequest{
								Name: name,
							}

							resp, err := client.GetAnalysis(ctx, req)
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
						Usage: "list analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of analyses to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListAnalysesResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListAnalysesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListAnalyses(ctx, req)
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
						Usage: "delete analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analyse", Usage: "The ID of the analyse.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("analyse"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAnalysis on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteAnalysisRequest{
								Name: name,
							}

							if err := client.DeleteAnalysis(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "analysis-rules",
				Usage: "Manage analysis-rules resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create analysis-rules",
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
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateAnalysisRuleRequest{
								Parent: parent,
							}

							resp, err := client.CreateAnalysisRule(ctx, req)
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
						Usage: "describe analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analysis-rule", Usage: "The ID of the analysis rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("analysis-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetAnalysisRuleRequest{
								Name: name,
							}

							resp, err := client.GetAnalysisRule(ctx, req)
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
						Usage: "list analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of analysis rule to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListAnalysisRulesResponse`;.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListAnalysisRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAnalysisRules(ctx, req)
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
						Name:  "update",
						Usage: "update analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analysis-rule", Usage: "The ID of the analysis rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "analysis_rule.name" not yet supported.
							analysis_rule_name := fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("analysis-rule"))
							fmt.Printf("Executing update on %s\n", analysis_rule_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analysis-rule", Usage: "The ID of the analysis rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("analysis-rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAnalysisRule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteAnalysisRuleRequest{
								Name: name,
							}

							if err := client.DeleteAnalysisRule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "conversations",
				Usage: "Manage conversations resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-id", Usage: "A unique ID for the new conversation.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateConversationRequest{
								Parent:         parent,
								ConversationId: cmd.String("conversation-id"),
							}

							resp, err := client.CreateConversation(ctx, req)
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
						Name:  "upload",
						Usage: "upload conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-id", Usage: "A unique ID for the new conversation.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.UploadConversationRequest{
								Parent:         parent,
								ConversationId: cmd.String("conversation-id"),
							}

							op, err := client.UploadConversation(ctx, req)
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
						Usage: "update conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "conversation.name" not yet supported.
							conversation_name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							fmt.Printf("Executing update on %s\n", conversation_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The level of details of the conversation.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetConversationRequest{
								Name: name,
								View: contactcenterinsightspb.ConversationView(contactcenterinsightspb.ConversationView_value[cmd.String("view")]),
							}

							resp, err := client.GetConversation(ctx, req)
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
						Usage: "list conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The attribute by which to order conversations in the response.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of conversations to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListConversationsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The level of details of the conversation.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListConversationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								View:      contactcenterinsightspb.ConversationView(contactcenterinsightspb.ConversationView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListConversations(ctx, req)
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
						Usage: "delete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all of this conversation's analyses will also be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteConversation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteConversationRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteConversation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "bulk-analyze",
						Usage: "bulk-analyze conversations",
						Flags: []cli.Flag{
							&cli.FloatFlag{Name: "analysis-percentage", Usage: "Percentage of selected conversation to analyze, between.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter used to select the subset of conversations to analyze.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.BulkAnalyzeConversationsRequest{
								Parent:             parent,
								Filter:             cmd.String("filter"),
								AnalysisPercentage: float32(cmd.Float("analysis-percentage")),
							}

							op, err := client.BulkAnalyzeConversations(ctx, req)
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
						Name:  "bulk-delete",
						Usage: "bulk-delete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter used to select the subset of conversations to delete.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all of this conversation's analyses will also be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-delete-count", Usage: "Maximum number of conversations to delete.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.BulkDeleteConversationsRequest{
								Parent:         parent,
								Filter:         cmd.String("filter"),
								MaxDeleteCount: int32(cmd.Int("max-delete-count")),
								Force:          cmd.Bool("force"),
							}

							op, err := client.BulkDeleteConversations(ctx, req)
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
						Name:  "ingest",
						Usage: "ingest conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "sample-size", Usage: "If set, this fields indicates the number of objects to ingest.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.IngestConversationsRequest{
								Parent:     parent,
								SampleSize: runtime.Ptr(int32(cmd.Int("sample-size"))),
							}

							op, err := client.IngestConversations(ctx, req)
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
						Name:  "calculate-stats",
						Usage: "calculate-stats conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing calculate-stats on %s\n", location)
							return nil
						},
					},
				},
			},
			{
				Name:  "encryption-spec",
				Usage: "Manage encryption-spec resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe encryption-spec",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/encryptionSpec", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetEncryptionSpecRequest{
								Name: name,
							}

							resp, err := client.GetEncryptionSpec(ctx, req)
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
						Name:  "initialize",
						Usage: "initialize encryption-spec",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "encryption_spec.name" not yet supported.
							encryption_spec_name := fmt.Sprintf("projects/%s/locations/%s/encryptionSpec", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing initialize on %s\n", encryption_spec_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "feedback-labels",
				Usage: "Manage feedback-labels resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "feedback-label-id", Usage: "The ID of the feedback label to create.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateFeedbackLabelRequest{
								Parent:          parent,
								FeedbackLabelId: cmd.String("feedback-label-id"),
							}

							resp, err := client.CreateFeedbackLabel(ctx, req)
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
						Usage: "list feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of feedback labels to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListFeedbackLabelsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListFeedbackLabelsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeedbackLabels(ctx, req)
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
						Usage: "describe feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "feedback-label", Usage: "The ID of the feedback label.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/feedbackLabels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("feedback-label"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetFeedbackLabelRequest{
								Name: name,
							}

							resp, err := client.GetFeedbackLabel(ctx, req)
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
						Usage: "update feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "feedback-label", Usage: "The ID of the feedback label.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feedback_label.name" not yet supported.
							feedback_label_name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/feedbackLabels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("feedback-label"))
							fmt.Printf("Executing update on %s\n", feedback_label_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "feedback-label", Usage: "The ID of the feedback label.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/feedbackLabels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("feedback-label"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteFeedbackLabel on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteFeedbackLabelRequest{
								Name: name,
							}

							if err := client.DeleteFeedbackLabel(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "insightsdata",
				Usage: "Manage insightsdata resources",
				Commands: []*cli.Command{

					{
						Name:  "export",
						Usage: "export insightsdata",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "kms-key", Usage: "A fully qualified KMS key name for BigQuery tables protected by CMEK.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "write-disposition", Usage: "Options for what to do if the destination table already exists.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ExportInsightsDataRequest{
								Parent:           parent,
								Filter:           cmd.String("filter"),
								KmsKey:           cmd.String("kms-key"),
								WriteDisposition: contactcenterinsightspb.ExportInsightsDataRequest_WriteDisposition(contactcenterinsightspb.ExportInsightsDataRequest_WriteDisposition_value[cmd.String("write-disposition")]),
							}

							op, err := client.ExportInsightsData(ctx, req)
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
				Name:  "issue-models",
				Usage: "Manage issue-models resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create issue-models",
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
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateIssueModelRequest{
								Parent: parent,
							}

							op, err := client.CreateIssueModel(ctx, req)
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
						Usage: "update issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "issue_model.name" not yet supported.
							issue_model_name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							fmt.Printf("Executing update on %s\n", issue_model_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetIssueModelRequest{
								Name: name,
							}

							resp, err := client.GetIssueModel(ctx, req)
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
						Usage: "list issue-models",
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
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListIssueModelsRequest{
								Parent: parent,
							}

							resp, err := client.ListIssueModels(ctx, req)
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
						Usage: "delete issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIssueModel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteIssueModelRequest{
								Name: name,
							}

							op, err := client.DeleteIssueModel(ctx, req)
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
						Name:  "deploy",
						Usage: "deploy issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeployIssueModelRequest{
								Name: name,
							}

							op, err := client.DeployIssueModel(ctx, req)
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
						Name:  "undeploy",
						Usage: "undeploy issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.UndeployIssueModelRequest{
								Name: name,
							}

							op, err := client.UndeployIssueModel(ctx, req)
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
						Name:  "export",
						Usage: "export issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ExportIssueModelRequest{
								Name: name,
							}

							op, err := client.ExportIssueModel(ctx, req)
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
						Name:  "import",
						Usage: "import issue-models",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "create-new-model", Usage: "If set to true, will create an issue model from the imported file.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ImportIssueModelRequest{
								Parent:         parent,
								CreateNewModel: cmd.Bool("create-new-model"),
							}

							op, err := client.ImportIssueModel(ctx, req)
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
						Name:  "calculate-issue-model-stats",
						Usage: "calculate-issue-model-stats issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							issue_model := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							fmt.Printf("Executing calculate-issue-model-stats on %s\n", issue_model)
							return nil
						},
					},
				},
			},
			{
				Name:  "issues",
				Usage: "Manage issues resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"), cmd.String("issue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetIssueRequest{
								Name: name,
							}

							resp, err := client.GetIssue(ctx, req)
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
						Usage: "list issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListIssuesRequest{
								Parent: parent,
							}

							resp, err := client.ListIssues(ctx, req)
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
						Usage: "update issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "issue.name" not yet supported.
							issue_name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"), cmd.String("issue"))
							fmt.Printf("Executing update on %s\n", issue_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "issue-model", Usage: "The ID of the issue model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue-model"), cmd.String("issue"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIssue on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteIssueRequest{
								Name: name,
							}

							if err := client.DeleteIssue(ctx, req); err != nil {
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
						Name:  "query-metrics",
						Usage: "query-metrics locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter to select a subset of conversations to compute the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "time-granularity", Usage: "The time granularity of each data point in the time series.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing query-metrics on %s\n", location)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset in the entire.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of feedback labels to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListAllFeedbackLabelsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListAllFeedbackLabelsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListAllFeedbackLabels(ctx, req)
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
						Name:  "bulk-upload-feedback-labels",
						Usage: "bulk-upload-feedback-labels locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, upload will not happen and the labels will be validated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.BulkUploadFeedbackLabelsRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.BulkUploadFeedbackLabels(ctx, req)
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
						Name:  "bulk-download-feedback-labels",
						Usage: "bulk-download-feedback-labels locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-filter", Usage: "Filter parent conversations to download feedback labels for.", Required: false},
							&cli.StringFlag{Name: "feedback-label-type", Usage: "The type of feedback labels that will be downloaded.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-download-count", Usage: "Limits the maximum number of feedback labels that will be.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "template-qa-scorecard-id", Usage: "If set, a template for labeling conversations and scorecard.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.BulkDownloadFeedbackLabelsRequest{
								Parent:                parent,
								Filter:                cmd.String("filter"),
								MaxDownloadCount:      int32(cmd.Int("max-download-count")),
								FeedbackLabelType:     contactcenterinsightspb.BulkDownloadFeedbackLabelsRequest_FeedbackLabelType(contactcenterinsightspb.BulkDownloadFeedbackLabelsRequest_FeedbackLabelType_value[cmd.String("feedback-label-type")]),
								ConversationFilter:    cmd.String("conversation-filter"),
								TemplateQaScorecardId: cmd.StringSlice("template-qa-scorecard-id"),
							}

							op, err := client.BulkDownloadFeedbackLabels(ctx, req)
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
				Name:  "phrase-matchers",
				Usage: "Manage phrase-matchers resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create phrase-matchers",
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
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreatePhraseMatcherRequest{
								Parent: parent,
							}

							resp, err := client.CreatePhraseMatcher(ctx, req)
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
						Usage: "describe phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-matcher", Usage: "The ID of the phrase matcher.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseMatchers/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-matcher"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetPhraseMatcherRequest{
								Name: name,
							}

							resp, err := client.GetPhraseMatcher(ctx, req)
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
						Usage: "list phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of phrase matchers to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListPhraseMatchersResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListPhraseMatchersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPhraseMatchers(ctx, req)
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
						Usage: "delete phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-matcher", Usage: "The ID of the phrase matcher.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseMatchers/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-matcher"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePhraseMatcher on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeletePhraseMatcherRequest{
								Name: name,
							}

							if err := client.DeletePhraseMatcher(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-matcher", Usage: "The ID of the phrase matcher.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "phrase_matcher.name" not yet supported.
							phrase_matcher_name := fmt.Sprintf("projects/%s/locations/%s/phraseMatchers/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-matcher"))
							fmt.Printf("Executing update on %s\n", phrase_matcher_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "qa-questions",
				Usage: "Manage qa-questions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-question-id", Usage: "A unique ID for the new question.", Required: false},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateQaQuestionRequest{
								Parent:       parent,
								QaQuestionId: cmd.String("qa-question-id"),
							}

							resp, err := client.CreateQaQuestion(ctx, req)
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
						Usage: "describe qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-question", Usage: "The ID of the qa question.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"), cmd.String("qa-question"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetQaQuestionRequest{
								Name: name,
							}

							resp, err := client.GetQaQuestion(ctx, req)
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
						Usage: "update qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-question", Usage: "The ID of the qa question.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "qa_question.name" not yet supported.
							qa_question_name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"), cmd.String("qa-question"))
							fmt.Printf("Executing update on %s\n", qa_question_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-question", Usage: "The ID of the qa question.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"), cmd.String("qa-question"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteQaQuestion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteQaQuestionRequest{
								Name: name,
							}

							if err := client.DeleteQaQuestion(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of questions to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListQaQuestionsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListQaQuestionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListQaQuestions(ctx, req)
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
				Name:  "qa-scorecards",
				Usage: "Manage qa-scorecards resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard-id", Usage: "A unique ID for the new QaScorecard.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateQaScorecardRequest{
								Parent:        parent,
								QaScorecardId: cmd.String("qa-scorecard-id"),
							}

							resp, err := client.CreateQaScorecard(ctx, req)
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
						Usage: "describe qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetQaScorecardRequest{
								Name: name,
							}

							resp, err := client.GetQaScorecard(ctx, req)
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
						Usage: "update qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "qa_scorecard.name" not yet supported.
							qa_scorecard_name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"))
							fmt.Printf("Executing update on %s\n", qa_scorecard_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete qa-scorecards",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all of this QaScorecard's child resources will.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteQaScorecard on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteQaScorecardRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteQaScorecard(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of scorecards to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListQaScorecardsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListQaScorecardsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListQaScorecards(ctx, req)
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
				Name:  "revisions",
				Usage: "Manage revisions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard-revision-id", Usage: "A unique ID for the new QaScorecardRevision.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateQaScorecardRevisionRequest{
								Parent:                parent,
								QaScorecardRevisionId: cmd.String("qa-scorecard-revision-id"),
							}

							resp, err := client.CreateQaScorecardRevision(ctx, req)
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
						Usage: "describe revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetQaScorecardRevisionRequest{
								Name: name,
							}

							resp, err := client.GetQaScorecardRevision(ctx, req)
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
						Name:  "tune-qa-scorecard-revision",
						Usage: "tune-qa-scorecard-revision revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for selecting the feedback labels that needs to be.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Run in validate only mode, no fine tuning will actually run.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.TuneQaScorecardRevisionRequest{
								Parent:       parent,
								Filter:       cmd.String("filter"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.TuneQaScorecardRevision(ctx, req)
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
						Name:  "deploy",
						Usage: "deploy revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeployQaScorecardRevisionRequest{
								Name: name,
							}

							resp, err := client.DeployQaScorecardRevision(ctx, req)
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
						Name:  "undeploy",
						Usage: "undeploy revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.UndeployQaScorecardRevisionRequest{
								Name: name,
							}

							resp, err := client.UndeployQaScorecardRevision(ctx, req)
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
						Usage: "delete revisions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all of this QaScorecardRevision's child resources.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"), cmd.String("revision"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteQaScorecardRevision on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteQaScorecardRevisionRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteQaScorecardRevision(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter to reduce results to a specific subset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of scorecard revisions to return in the.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard", Usage: "The ID of the qa scorecard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa-scorecard"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListQaScorecardRevisionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListQaScorecardRevisions(ctx, req)
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
				Name:  "settings",
				Usage: "Manage settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/settings", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetSettingsRequest{
								Name: name,
							}

							resp, err := client.GetSettings(ctx, req)
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
						Usage: "update settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "settings.name" not yet supported.
							settings_name := fmt.Sprintf("projects/%s/locations/%s/settings", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", settings_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "views",
				Usage: "Manage views resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create views",
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
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.CreateViewRequest{
								Parent: parent,
							}

							resp, err := client.CreateView(ctx, req)
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
						Usage: "describe views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The ID of the view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.GetViewRequest{
								Name: name,
							}

							resp, err := client.GetView(ctx, req)
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
						Usage: "list views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of views to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListViewsResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.ListViewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListViews(ctx, req)
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
						Name:  "update",
						Usage: "update views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The ID of the view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "view.name" not yet supported.
							view_name := fmt.Sprintf("projects/%s/locations/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("view"))
							fmt.Printf("Executing update on %s\n", view_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The ID of the view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("view"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteView on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := contactcenterinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &contactcenterinsightspb.DeleteViewRequest{
								Name: name,
							}

							if err := client.DeleteView(ctx, req); err != nil {
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
