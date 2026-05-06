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

package cx

import (
	cx "cloud.google.com/go/cx/apiv3"
	"cloud.google.com/go/cx/apiv3/cxpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the dialogflow command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "dialogflow",
		Usage: "manage Dialogflow API resources",
		Commands: []*cli.Command{
			{
				Name:  "agents",
				Usage: "Manage agents resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListAgentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAgents(ctx, req)
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
						Usage: "describe agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetAgentRequest{
								Name: name,
							}

							resp, err := client.GetAgent(ctx, req)
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
						Usage: "create agents",
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
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateAgentRequest{
								Parent: parent,
							}

							resp, err := client.CreateAgent(ctx, req)
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
						Usage: "update agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "agent.name" not yet supported.
							agent_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							fmt.Printf("Executing update on %s\n", agent_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAgent on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteAgentRequest{
								Name: name,
							}

							if err := client.DeleteAgent(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "export",
						Usage: "export agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "agent-uri", Usage: "The [Google Cloud.", Required: false},
							&cli.StringFlag{Name: "data-format", Usage: "The data format of the exported agent.", Required: false},
							&cli.StringFlag{Name: "environment", Usage: "Environment name.", Required: false},
							&cli.BoolFlag{Name: "include-bigquery-export-settings", Usage: "Whether to include BigQuery Export setting.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ExportAgentRequest{
								Name:                          name,
								AgentUri:                      cmd.String("agent-uri"),
								DataFormat:                    cxpb.ExportAgentRequest_DataFormat(cxpb.ExportAgentRequest_DataFormat_value[cmd.String("data-format")]),
								Environment:                   cmd.String("environment"),
								IncludeBigqueryExportSettings: cmd.Bool("include-bigquery-export-settings"),
							}

							op, err := client.ExportAgent(ctx, req)
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
						Name:  "restore",
						Usage: "restore agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-option", Usage: "Agent restore mode.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("RestoreAgent %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.RestoreAgentRequest{
								Name:          name,
								RestoreOption: cxpb.RestoreAgentRequest_RestoreOption(cxpb.RestoreAgentRequest_RestoreOption_value[cmd.String("restore-option")]),
							}

							op, err := client.RestoreAgent(ctx, req)
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
						Name:  "validate",
						Usage: "validate agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "If not specified, the agent's default language is used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ValidateAgentRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.ValidateAgent(ctx, req)
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
				Name:  "changelogs",
				Usage: "Manage changelogs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list changelogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter string.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListChangelogsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListChangelogs(ctx, req)
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
						Usage: "describe changelogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "changelog", Usage: "The ID of the changelog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/changelogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("changelog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetChangelogRequest{
								Name: name,
							}

							resp, err := client.GetChangelog(ctx, req)
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
				Name:  "continuous-test-results",
				Usage: "Manage continuous-test-results resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list continuous-test-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListContinuousTestResultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListContinuousTestResults(ctx, req)
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
				Name:  "deployments",
				Usage: "Manage deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeployments(ctx, req)
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
						Usage: "describe deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetDeployment(ctx, req)
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
				Name:  "entity-types",
				Usage: "Manage entity-types resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to retrieve the entity type for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetEntityTypeRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetEntityType(ctx, req)
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
						Usage: "create entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `entity_type`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateEntityTypeRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.CreateEntityType(ctx, req)
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
						Usage: "update entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `entity_type`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entity_type.name" not yet supported.
							entity_type_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("entity-type"))
							fmt.Printf("Executing update on %s\n", entity_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for entity type not being used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEntityType on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteEntityTypeRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteEntityType(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list entity types for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListEntityTypesRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntityTypes(ctx, req)
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
						Name:  "export",
						Usage: "export entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "data-format", Usage: "The data format of the exported entity types.", Required: false},
							&cli.StringSliceFlag{Name: "entity-types", Usage: "The name of the entity types to export.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to retrieve the entity type for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ExportEntityTypesRequest{
								Parent:       parent,
								EntityTypes:  cmd.StringSlice("entity-types"),
								DataFormat:   cxpb.ExportEntityTypesRequest_DataFormat(cxpb.ExportEntityTypesRequest_DataFormat_value[cmd.String("data-format")]),
								LanguageCode: cmd.String("language-code"),
							}

							op, err := client.ExportEntityTypes(ctx, req)
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
						Usage: "import entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "merge-option", Usage: "Merge option for importing entity types.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target-entity-type", Usage: "The target entity type to import into.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ImportEntityTypesRequest{
								Parent:           parent,
								MergeOption:      cxpb.ImportEntityTypesRequest_MergeOption(cxpb.ImportEntityTypesRequest_MergeOption_value[cmd.String("merge-option")]),
								TargetEntityType: cmd.String("target-entity-type"),
							}

							op, err := client.ImportEntityTypes(ctx, req)
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
						Usage: "list entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListSessionEntityTypesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSessionEntityTypes(ctx, req)
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
						Usage: "describe entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetSessionEntityTypeRequest{
								Name: name,
							}

							resp, err := client.GetSessionEntityType(ctx, req)
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
						Usage: "create entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateSessionEntityTypeRequest{
								Parent: parent,
							}

							resp, err := client.CreateSessionEntityType(ctx, req)
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
						Usage: "update entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "session_entity_type.name" not yet supported.
							session_entity_type_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"), cmd.String("entity-type"))
							fmt.Printf("Executing update on %s\n", session_entity_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSessionEntityType on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteSessionEntityTypeRequest{
								Name: name,
							}

							if err := client.DeleteSessionEntityType(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "environments",
				Usage: "Manage environments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListEnvironmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEnvironments(ctx, req)
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
						Usage: "describe environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetEnvironmentRequest{
								Name: name,
							}

							resp, err := client.GetEnvironment(ctx, req)
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
						Usage: "create environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateEnvironmentRequest{
								Parent: parent,
							}

							op, err := client.CreateEnvironment(ctx, req)
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
						Usage: "update environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "environment.name" not yet supported.
							environment_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							fmt.Printf("Executing update on %s\n", environment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEnvironment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteEnvironmentRequest{
								Name: name,
							}

							if err := client.DeleteEnvironment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "lookup-environment-history",
						Usage: "lookup-environment-history environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.LookupEnvironmentHistoryRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.LookupEnvironmentHistory(ctx, req)
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
						Name:  "run-continuous-test",
						Usage: "run-continuous-test environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							fmt.Printf("Executing run-continuous-test on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "deploy-flow",
						Usage: "deploy-flow environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "flow-version", Usage: "The flow version to deploy.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							fmt.Printf("Executing deploy-flow on %s\n", environment)
							return nil
						},
					},
				},
			},
			{
				Name:  "examples",
				Usage: "Manage examples resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateExampleRequest{
								Parent: parent,
							}

							resp, err := client.CreateExample(ctx, req)
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
						Usage: "delete examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "example", Usage: "The ID of the example.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("example"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteExample on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteExampleRequest{
								Name: name,
							}

							if err := client.DeleteExample(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list examples for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The [next_page_token][ListExampleResponse.", Required: false},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListExamplesRequest{
								Parent:       parent,
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								LanguageCode: cmd.String("language-code"),
							}

							limit := cmd.Int("limit")
							it := client.ListExamples(ctx, req)
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
						Usage: "describe examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "example", Usage: "The ID of the example.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("example"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetExampleRequest{
								Name: name,
							}

							resp, err := client.GetExample(ctx, req)
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
						Usage: "update examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "example", Usage: "The ID of the example.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "example.name" not yet supported.
							example_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("example"))
							fmt.Printf("Executing update on %s\n", example_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "experiments",
				Usage: "Manage experiments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListExperimentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListExperiments(ctx, req)
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
						Usage: "describe experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetExperimentRequest{
								Name: name,
							}

							resp, err := client.GetExperiment(ctx, req)
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
						Usage: "create experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateExperimentRequest{
								Parent: parent,
							}

							resp, err := client.CreateExperiment(ctx, req)
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
						Usage: "update experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "experiment.name" not yet supported.
							experiment_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
							fmt.Printf("Executing update on %s\n", experiment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteExperiment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteExperimentRequest{
								Name: name,
							}

							if err := client.DeleteExperiment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "start",
						Usage: "start experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.StartExperimentRequest{
								Name: name,
							}

							resp, err := client.StartExperiment(ctx, req)
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
						Usage: "stop experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.StopExperimentRequest{
								Name: name,
							}

							resp, err := client.StopExperiment(ctx, req)
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
				Name:  "flows",
				Usage: "Manage flows resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `flow`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateFlowRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.CreateFlow(ctx, req)
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
						Usage: "delete flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for flows with no incoming transitions.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteFlow on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteFlowRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteFlow(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list flows for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListFlowsRequest{
								Parent:       parent,
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								LanguageCode: cmd.String("language-code"),
							}

							limit := cmd.Int("limit")
							it := client.ListFlows(ctx, req)
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
						Usage: "describe flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to retrieve the flow for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetFlowRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetFlow(ctx, req)
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
						Usage: "update flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `flow`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "flow.name" not yet supported.
							flow_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							fmt.Printf("Executing update on %s\n", flow_name)
							return nil
						},
					},

					{
						Name:  "train",
						Usage: "train flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("TrainFlow %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.TrainFlowRequest{
								Name: name,
							}

							op, err := client.TrainFlow(ctx, req)
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
						Name:  "validate",
						Usage: "validate flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "If not specified, the agent's default language is used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ValidateFlowRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.ValidateFlow(ctx, req)
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
						Usage: "import flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "import-option", Usage: "Flow import mode.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ImportFlowRequest{
								Parent:       parent,
								ImportOption: cxpb.ImportFlowRequest_ImportOption(cxpb.ImportFlowRequest_ImportOption_value[cmd.String("import-option")]),
							}

							op, err := client.ImportFlow(ctx, req)
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
						Usage: "export flows",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "flow-uri", Usage: "The [Google Cloud.", Required: false},
							&cli.BoolFlag{Name: "include-referenced-flows", Usage: "Whether to export flows referenced by the specified flow.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ExportFlowRequest{
								Name:                   name,
								FlowUri:                cmd.String("flow-uri"),
								IncludeReferencedFlows: cmd.Bool("include-referenced-flows"),
							}

							op, err := client.ExportFlow(ctx, req)
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
				Name:  "generative-settings",
				Usage: "Manage generative-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe generative-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "Language code of the generative settings.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generativeSettings", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetGenerativeSettingsRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetGenerativeSettings(ctx, req)
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
						Usage: "update generative-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "generative_settings.name" not yet supported.
							generative_settings_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generativeSettings", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							fmt.Printf("Executing update on %s\n", generative_settings_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "generators",
				Usage: "Manage generators resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list generators for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListGeneratorsRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGenerators(ctx, req)
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
						Usage: "describe generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list generators for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("generator"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetGeneratorRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetGenerator(ctx, req)
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
						Usage: "create generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to create generators for the following fields:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateGeneratorRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.CreateGenerator(ctx, req)
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
						Usage: "update generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list generators for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "generator.name" not yet supported.
							generator_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("generator"))
							fmt.Printf("Executing update on %s\n", generator_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for generators not being used.", Required: false},
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("generator"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGenerator on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteGeneratorRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteGenerator(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "intents",
				Usage: "Manage intents resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "intent-view", Usage: "The resource view to apply to the returned intent.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list intents for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListIntentsRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								IntentView:   cxpb.IntentView(cxpb.IntentView_value[cmd.String("intent-view")]),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListIntents(ctx, req)
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
						Usage: "describe intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "intent", Usage: "The ID of the intent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to retrieve the intent for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/intents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("intent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetIntentRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetIntent(ctx, req)
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
						Usage: "create intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `intent`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateIntentRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.CreateIntent(ctx, req)
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
						Usage: "update intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "intent", Usage: "The ID of the intent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `intent`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "intent.name" not yet supported.
							intent_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/intents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("intent"))
							fmt.Printf("Executing update on %s\n", intent_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "intent", Usage: "The ID of the intent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/intents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("intent"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIntent on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteIntentRequest{
								Name: name,
							}

							if err := client.DeleteIntent(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "import",
						Usage: "import intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "merge-option", Usage: "Merge option for importing intents.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ImportIntentsRequest{
								Parent:      parent,
								MergeOption: cxpb.ImportIntentsRequest_MergeOption(cxpb.ImportIntentsRequest_MergeOption_value[cmd.String("merge-option")]),
							}

							op, err := client.ImportIntents(ctx, req)
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
						Usage: "export intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "data-format", Usage: "The data format of the exported intents.", Required: false},
							&cli.StringSliceFlag{Name: "intents", Usage: "The name of the intents to export.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ExportIntentsRequest{
								Parent:     parent,
								Intents:    cmd.StringSlice("intents"),
								DataFormat: cxpb.ExportIntentsRequest_DataFormat(cxpb.ExportIntentsRequest_DataFormat_value[cmd.String("data-format")]),
							}

							op, err := client.ExportIntents(ctx, req)
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
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "pages",
				Usage: "Manage pages resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list pages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list pages for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListPagesRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPages(ctx, req)
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
						Usage: "describe pages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to retrieve the page for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "page", Usage: "The ID of the page.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("page"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetPageRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetPage(ctx, req)
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
						Usage: "create pages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `page`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreatePageRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.CreatePage(ctx, req)
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
						Usage: "update pages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `page`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "page", Usage: "The ID of the page.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "page.name" not yet supported.
							page_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("page"))
							fmt.Printf("Executing update on %s\n", page_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete pages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for pages with no incoming transitions.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "page", Usage: "The ID of the page.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("page"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePage on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeletePageRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeletePage(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "playbooks",
				Usage: "Manage playbooks resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreatePlaybookRequest{
								Parent: parent,
							}

							resp, err := client.CreatePlaybook(ctx, req)
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
						Usage: "delete playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePlaybook on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeletePlaybookRequest{
								Name: name,
							}

							if err := client.DeletePlaybook(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListPlaybooksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPlaybooks(ctx, req)
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
						Usage: "describe playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetPlaybookRequest{
								Name: name,
							}

							resp, err := client.GetPlaybook(ctx, req)
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
						Usage: "export playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "data-format", Usage: "The data format of the exported agent.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "playbook-uri", Usage: "The [Google Cloud.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ExportPlaybookRequest{
								Name:        name,
								PlaybookUri: cmd.String("playbook-uri"),
								DataFormat:  cxpb.ExportPlaybookRequest_DataFormat(cxpb.ExportPlaybookRequest_DataFormat_value[cmd.String("data-format")]),
							}

							op, err := client.ExportPlaybook(ctx, req)
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
						Usage: "import playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ImportPlaybookRequest{
								Parent: parent,
							}

							op, err := client.ImportPlaybook(ctx, req)
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
						Usage: "update playbooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "playbook.name" not yet supported.
							playbook_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							fmt.Printf("Executing update on %s\n", playbook_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "results",
				Usage: "Manage results resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter expression used to filter test case results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "test-case", Usage: "The ID of the test case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test-case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListTestCaseResultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListTestCaseResults(ctx, req)
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
						Usage: "describe results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "result", Usage: "The ID of the result.", Required: true},
							&cli.StringFlag{Name: "test-case", Usage: "The ID of the test case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s/results/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test-case"), cmd.String("result"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetTestCaseResultRequest{
								Name: name,
							}

							resp, err := client.GetTestCaseResult(ctx, req)
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
				Name:  "security-settings",
				Usage: "Manage security-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create security-settings",
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
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateSecuritySettingsRequest{
								Parent: parent,
							}

							resp, err := client.CreateSecuritySettings(ctx, req)
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
						Usage: "describe security-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-setting", Usage: "The ID of the security setting.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-setting"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetSecuritySettingsRequest{
								Name: name,
							}

							resp, err := client.GetSecuritySettings(ctx, req)
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
						Usage: "update security-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-setting", Usage: "The ID of the security setting.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "security_settings.name" not yet supported.
							security_settings_name := fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-setting"))
							fmt.Printf("Executing update on %s\n", security_settings_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list security-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListSecuritySettingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecuritySettings(ctx, req)
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
						Usage: "delete security-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-setting", Usage: "The ID of the security setting.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-setting"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSecuritySettings on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteSecuritySettingsRequest{
								Name: name,
							}

							if err := client.DeleteSecuritySettings(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "sessions",
				Usage: "Manage sessions resources",
				Commands: []*cli.Command{

					{
						Name:  "detect-intent",
						Usage: "detect-intent sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "response-view", Usage: "Specifies which fields in the.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							fmt.Printf("Executing detect-intent on %s\n", session)
							return nil
						},
					},

					{
						Name:  "server-streaming-detect-intent",
						Usage: "server-streaming-detect-intent sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "response-view", Usage: "Specifies which fields in the.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							fmt.Printf("Executing server-streaming-detect-intent on %s\n", session)
							return nil
						},
					},

					{
						Name:  "match-intent",
						Usage: "match-intent sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.BoolFlag{Name: "persist-parameter-changes", Usage: "Persist session parameter changes from `query_params`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							fmt.Printf("Executing match-intent on %s\n", session)
							return nil
						},
					},

					{
						Name:  "fulfill-intent",
						Usage: "fulfill-intent sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "match_intent_request.session" not yet supported.
							match_intent_request_session := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							fmt.Printf("Executing fulfill-intent on %s\n", match_intent_request_session)
							return nil
						},
					},

					{
						Name:  "submit-answer-feedback",
						Usage: "submit-answer-feedback sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "response-id", Usage: "ID of the response to update its feedback.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
							fmt.Printf("Executing submit-answer-feedback on %s\n", session)
							return nil
						},
					},
				},
			},
			{
				Name:  "test-cases",
				Usage: "Manage test-cases resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specifies whether response should include all fields or just the metadata.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListTestCasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								View:      cxpb.ListTestCasesRequest_TestCaseView(cxpb.ListTestCasesRequest_TestCaseView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListTestCases(ctx, req)
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
						Name:  "batch-delete",
						Usage: "batch-delete test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "Format of test case names:.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute BatchDeleteTestCases on %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.BatchDeleteTestCasesRequest{
								Parent: parent,
								Names:  cmd.StringSlice("names"),
							}

							if err := client.BatchDeleteTestCases(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "test-case", Usage: "The ID of the test case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test-case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetTestCaseRequest{
								Name: name,
							}

							resp, err := client.GetTestCase(ctx, req)
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
						Usage: "create test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateTestCaseRequest{
								Parent: parent,
							}

							resp, err := client.CreateTestCase(ctx, req)
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
						Usage: "update test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "test-case", Usage: "The ID of the test case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "test_case.name" not yet supported.
							test_case_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test-case"))
							fmt.Printf("Executing update on %s\n", test_case_name)
							return nil
						},
					},

					{
						Name:  "run",
						Usage: "run test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "Environment name.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "test-case", Usage: "The ID of the test case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test-case"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.RunTestCaseRequest{
								Name:        name,
								Environment: cmd.String("environment"),
							}

							op, err := client.RunTestCase(ctx, req)
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
						Name:  "batch-run",
						Usage: "batch-run test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "If not set, draft environment is assumed.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "test-cases", Usage: "Format:.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.BatchRunTestCasesRequest{
								Parent:      parent,
								Environment: cmd.String("environment"),
								TestCases:   cmd.StringSlice("test-cases"),
							}

							op, err := client.BatchRunTestCases(ctx, req)
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
						Name:  "calculate-coverage",
						Usage: "calculate-coverage test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The type of coverage requested.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							agent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							fmt.Printf("Executing calculate-coverage on %s\n", agent)
							return nil
						},
					},

					{
						Name:  "import",
						Usage: "import test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ImportTestCasesRequest{
								Parent: parent,
							}

							op, err := client.ImportTestCases(ctx, req)
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
						Usage: "export test-cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "data-format", Usage: "The data format of the exported test cases.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "The filter expression used to filter exported test cases, see.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ExportTestCasesRequest{
								Parent:     parent,
								DataFormat: cxpb.ExportTestCasesRequest_DataFormat(cxpb.ExportTestCasesRequest_DataFormat_value[cmd.String("data-format")]),
								Filter:     cmd.String("filter"),
							}

							op, err := client.ExportTestCases(ctx, req)
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
				Name:  "tools",
				Usage: "Manage tools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateToolRequest{
								Parent: parent,
							}

							resp, err := client.CreateTool(ctx, req)
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
						Usage: "list tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListToolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTools(ctx, req)
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
						Usage: "describe tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetToolRequest{
								Name: name,
							}

							resp, err := client.GetTool(ctx, req)
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
						Usage: "update tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tool.name" not yet supported.
							tool_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
							fmt.Printf("Executing update on %s\n", tool_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for Tools not being used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTool on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteToolRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteTool(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "transition-route-groups",
				Usage: "Manage transition-route-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list transition-route-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to list transition route groups for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListTransitionRouteGroupsRequest{
								Parent:       parent,
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								LanguageCode: cmd.String("language-code"),
							}

							limit := cmd.Int("limit")
							it := client.ListTransitionRouteGroups(ctx, req)
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
						Usage: "describe transition-route-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to retrieve the transition route group for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transition-route-group", Usage: "The ID of the transition route group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/transitionRouteGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("transition-route-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetTransitionRouteGroupRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetTransitionRouteGroup(ctx, req)
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
						Usage: "create transition-route-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `TransitionRouteGroup`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateTransitionRouteGroupRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.CreateTransitionRouteGroup(ctx, req)
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
						Usage: "update transition-route-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language of the following fields in `TransitionRouteGroup`:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transition-route-group", Usage: "The ID of the transition route group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "transition_route_group.name" not yet supported.
							transition_route_group_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/transitionRouteGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("transition-route-group"))
							fmt.Printf("Executing update on %s\n", transition_route_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete transition-route-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for transition route group that no page is using.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transition-route-group", Usage: "The ID of the transition route group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/transitionRouteGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("transition-route-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTransitionRouteGroup on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteTransitionRouteGroupRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteTransitionRouteGroup(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "validation-result",
				Usage: "Manage validation-result resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe validation-result",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "If not specified, the agent's default language is used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/validationResult", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetAgentValidationResultRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetAgentValidationResult(ctx, req)
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
						Usage: "describe validation-result",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "If not specified, the agent's default language is used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/validationResult", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetFlowValidationResultRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetFlowValidationResult(ctx, req)
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
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreatePlaybookVersionRequest{
								Parent: parent,
							}

							resp, err := client.CreatePlaybookVersion(ctx, req)
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
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetPlaybookVersionRequest{
								Name: name,
							}

							resp, err := client.GetPlaybookVersion(ctx, req)
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
						Name:  "restore",
						Usage: "restore versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.RestorePlaybookVersionRequest{
								Name: name,
							}

							resp, err := client.RestorePlaybookVersion(ctx, req)
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
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListPlaybookVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPlaybookVersions(ctx, req)
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
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "playbook", Usage: "The ID of the playbook.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePlaybookVersion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeletePlaybookVersionRequest{
								Name: name,
							}

							if err := client.DeletePlaybookVersion(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListToolVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListToolVersions(ctx, req)
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
						Usage: "create versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateToolVersionRequest{
								Parent: parent,
							}

							resp, err := client.CreateToolVersion(ctx, req)
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
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetToolVersionRequest{
								Name: name,
							}

							resp, err := client.GetToolVersion(ctx, req)
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
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for Tools not being used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteToolVersion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteToolVersionRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteToolVersion(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "restore",
						Usage: "restore versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.RestoreToolVersionRequest{
								Name: name,
							}

							resp, err := client.RestoreToolVersion(ctx, req)
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
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListVersions(ctx, req)
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
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetVersionRequest{
								Name: name,
							}

							resp, err := client.GetVersion(ctx, req)
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
						Usage: "create versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateVersionRequest{
								Parent: parent,
							}

							op, err := client.CreateVersion(ctx, req)
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
						Usage: "update versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "version.name" not yet supported.
							version_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
							fmt.Printf("Executing update on %s\n", version_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteVersion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteVersionRequest{
								Name: name,
							}

							if err := client.DeleteVersion(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "load",
						Usage: "load versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.BoolFlag{Name: "allow-override-agent-resources", Usage: "This field is used to prevent accidental overwrite of other agent.", Required: false},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("LoadVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.LoadVersionRequest{
								Name:                        name,
								AllowOverrideAgentResources: cmd.Bool("allow-override-agent-resources"),
							}

							op, err := client.LoadVersion(ctx, req)
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
						Name:  "compare-versions",
						Usage: "compare-versions versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "base-version", Usage: "Name of the base flow version to compare with the target version.", Required: true},
							&cli.StringFlag{Name: "flow", Usage: "The ID of the flow.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language to compare the flow versions for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target-version", Usage: "Name of the target flow version to compare with the.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							base_version := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
							fmt.Printf("Executing compare-versions on %s\n", base_version)
							return nil
						},
					},
				},
			},
			{
				Name:  "webhooks",
				Usage: "Manage webhooks resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list webhooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.ListWebhooksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWebhooks(ctx, req)
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
						Usage: "describe webhooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "webhook", Usage: "The ID of the webhook.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("webhook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.GetWebhookRequest{
								Name: name,
							}

							resp, err := client.GetWebhook(ctx, req)
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
						Usage: "create webhooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.CreateWebhookRequest{
								Parent: parent,
							}

							resp, err := client.CreateWebhook(ctx, req)
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
						Usage: "update webhooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "webhook", Usage: "The ID of the webhook.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "webhook.name" not yet supported.
							webhook_name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("webhook"))
							fmt.Printf("Executing update on %s\n", webhook_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete webhooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "This field has no effect for webhook not being used.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "webhook", Usage: "The ID of the webhook.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("webhook"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteWebhook on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cx.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cxpb.DeleteWebhookRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteWebhook(ctx, req); err != nil {
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
