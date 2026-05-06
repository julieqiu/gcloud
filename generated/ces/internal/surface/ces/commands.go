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

package ces

import (
	ces "cloud.google.com/go/ces/apiv1"
	"cloud.google.com/go/ces/apiv1/cespb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the ces command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "ces",
		Usage: "manage Gemini Enterprise for Customer Experience API resources",
		Commands: []*cli.Command{
			{
				Name:  "agents",
				Usage: "Manage agents resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the agents.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListAgentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("agent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetAgentRequest{
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
							&cli.StringFlag{Name: "agent-id", Usage: "The ID to use for the agent, which will become the final.", Required: false},
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateAgentRequest{
								Parent:  parent,
								AgentId: cmd.String("agent-id"),
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "agent.name" not yet supported.
							agent_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("agent"))
							fmt.Printf("Executing update on %s\n", agent_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete agents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent", Usage: "The ID of the agent.", Required: true},
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the agent.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "Indicates whether to forcefully delete the agent, even if it is.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("agent"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAgent on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteAgentRequest{
								Name:  name,
								Force: cmd.Bool("force"),
								Etag:  cmd.String("etag"),
							}

							if err := client.DeleteAgent(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "apps",
				Usage: "Manage apps resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the apps.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListAppsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListApps(ctx, req)
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
						Usage: "describe apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetAppRequest{
								Name: name,
							}

							resp, err := client.GetApp(ctx, req)
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
						Usage: "create apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app-id", Usage: "The ID to use for the app, which will become the final component.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateAppRequest{
								Parent: parent,
								AppId:  cmd.String("app-id"),
							}

							op, err := client.CreateApp(ctx, req)
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
						Usage: "update apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "app.name" not yet supported.
							app_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							fmt.Printf("Executing update on %s\n", app_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the app.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteApp %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteAppRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteApp(ctx, req)
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
						Name:  "export-app",
						Usage: "export-app apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "app-version", Usage: "The resource name of the app version to export.", Required: false},
							&cli.StringFlag{Name: "export-format", Usage: "The format to export the app in.", Required: true},
							&cli.StringFlag{Name: "gcs-uri", Usage: "The [Google Cloud.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ExportAppRequest{
								Name:         name,
								ExportFormat: cespb.ExportAppRequest_ExportFormat(cespb.ExportAppRequest_ExportFormat_value[cmd.String("export-format")]),
								GcsUri:       cmd.String("gcs-uri"),
								AppVersion:   cmd.String("app-version"),
							}

							op, err := client.ExportApp(ctx, req)
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
						Name:  "import-app",
						Usage: "import-app apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app-id", Usage: "The ID to use for the imported app.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name of the app to import.", Required: false},
							&cli.BoolFlag{Name: "ignore-app-lock", Usage: "Flag for overriding the app lock during import.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ImportAppRequest{
								Parent:        parent,
								DisplayName:   cmd.String("display-name"),
								AppId:         cmd.String("app-id"),
								IgnoreAppLock: cmd.Bool("ignore-app-lock"),
							}

							op, err := client.ImportApp(ctx, req)
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
						Name:  "execute-tool",
						Usage: "execute-tool apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ExecuteToolRequest{
								Parent: parent,
							}

							resp, err := client.ExecuteTool(ctx, req)
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
						Name:  "retrieve-tool-schema",
						Usage: "retrieve-tool-schema apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.RetrieveToolSchemaRequest{
								Parent: parent,
							}

							resp, err := client.RetrieveToolSchema(ctx, req)
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the changelogs.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListChangelogsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "changelog", Usage: "The ID of the changelog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/changelogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("changelog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetChangelogRequest{
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
				Name:  "conversations",
				Usage: "Manage conversations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the conversations.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "Indicate the source of the conversation.", Required: false},
							&cli.StringSliceFlag{Name: "sources", Usage: "Indicate the sources of the conversations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListConversationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								Source:    cespb.Conversation_Source(cespb.Conversation_Source_value[cmd.String("source")]),
								Sources:   cmd.StringSlice("sources"),
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
						Name:  "describe",
						Usage: "describe conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "Indicate the source of the conversation.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetConversationRequest{
								Name:   name,
								Source: cespb.Conversation_Source(cespb.Conversation_Source_value[cmd.String("source")]),
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
						Name:  "delete",
						Usage: "delete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "Indicate the source of the conversation.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("conversation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteConversation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteConversationRequest{
								Name:   name,
								Source: cespb.Conversation_Source(cespb.Conversation_Source_value[cmd.String("source")]),
							}

							if err := client.DeleteConversation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "batch-delete",
						Usage: "batch-delete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringSliceFlag{Name: "conversations", Usage: "The resource names of the conversations to delete.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.BatchDeleteConversationsRequest{
								Parent:        parent,
								Conversations: cmd.StringSlice("conversations"),
							}

							op, err := client.BatchDeleteConversations(ctx, req)
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
				Name:  "deployments",
				Usage: "Manage deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of deployments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDeployments` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetDeploymentRequest{
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

					{
						Name:  "create",
						Usage: "create deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "deployment-id", Usage: "The ID to use for the deployment, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateDeploymentRequest{
								Parent:       parent,
								DeploymentId: cmd.String("deployment-id"),
							}

							resp, err := client.CreateDeployment(ctx, req)
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
						Usage: "update deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deployment.name" not yet supported.
							deployment_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("deployment"))
							fmt.Printf("Executing update on %s\n", deployment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag of the deployment.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDeployment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteDeploymentRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteDeployment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "list",
						Usage: "list examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the examples.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListExamplesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "example", Usage: "The ID of the example.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("example"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetExampleRequest{
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
						Name:  "create",
						Usage: "create examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "example-id", Usage: "The ID to use for the example, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateExampleRequest{
								Parent:    parent,
								ExampleId: cmd.String("example-id"),
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
						Name:  "update",
						Usage: "update examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "example", Usage: "The ID of the example.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "example.name" not yet supported.
							example_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("example"))
							fmt.Printf("Executing update on %s\n", example_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the example.", Required: false},
							&cli.StringFlag{Name: "example", Usage: "The ID of the example.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("example"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteExample on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteExampleRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteExample(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "guardrails",
				Usage: "Manage guardrails resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list guardrails",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the guardrails.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListGuardrailsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGuardrails(ctx, req)
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
						Usage: "describe guardrails",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "guardrail", Usage: "The ID of the guardrail.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/guardrails/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("guardrail"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetGuardrailRequest{
								Name: name,
							}

							resp, err := client.GetGuardrail(ctx, req)
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
						Usage: "create guardrails",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "guardrail-id", Usage: "The ID to use for the guardrail, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateGuardrailRequest{
								Parent:      parent,
								GuardrailId: cmd.String("guardrail-id"),
							}

							resp, err := client.CreateGuardrail(ctx, req)
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
						Usage: "update guardrails",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "guardrail", Usage: "The ID of the guardrail.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "guardrail.name" not yet supported.
							guardrail_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/guardrails/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("guardrail"))
							fmt.Printf("Executing update on %s\n", guardrail_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete guardrails",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the guardrail.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "Indicates whether to forcefully delete the guardrail, even if it.", Required: false},
							&cli.StringFlag{Name: "guardrail", Usage: "The ID of the guardrail.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/guardrails/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("guardrail"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGuardrail on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteGuardrailRequest{
								Name:  name,
								Force: cmd.Bool("force"),
								Etag:  cmd.String("etag"),
							}

							if err := client.DeleteGuardrail(ctx, req); err != nil {
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
				Name:  "sessions",
				Usage: "Manage sessions resources",
				Commands: []*cli.Command{

					{
						Name:  "run-session",
						Usage: "run-session sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "config.session" not yet supported.
							config_session := fmt.Sprintf("projects/%s/locations/%s/apps/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("session"))
							fmt.Printf("Executing run-session on %s\n", config_session)
							return nil
						},
					},

					{
						Name:  "stream-run-session",
						Usage: "stream-run-session sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "config.session" not yet supported.
							config_session := fmt.Sprintf("projects/%s/locations/%s/apps/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("session"))
							fmt.Printf("Executing stream-run-session on %s\n", config_session)
							return nil
						},
					},

					{
						Name:  "generate-chat-token",
						Usage: "generate-chat-token sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment of the app to use for the session.", Required: true},
							&cli.BoolFlag{Name: "live-handoff-enabled", Usage: "Indicates if live handoff is enabled for the session.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recaptcha-token", Usage: "The reCAPTCHA token generated by the client-side chat widget.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GenerateChatTokenRequest{
								Name:               name,
								Deployment:         cmd.String("deployment"),
								RecaptchaToken:     cmd.String("recaptcha-token"),
								LiveHandoffEnabled: cmd.Bool("live-handoff-enabled"),
							}

							resp, err := client.GenerateChatToken(ctx, req)
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
						Name:  "list",
						Usage: "list tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the tools.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListToolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("tool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetToolRequest{
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
						Name:  "create",
						Usage: "create tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool-id", Usage: "The ID to use for the tool, which will become the final component.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateToolRequest{
								Parent: parent,
								ToolId: cmd.String("tool-id"),
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
						Name:  "update",
						Usage: "update tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tool.name" not yet supported.
							tool_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("tool"))
							fmt.Printf("Executing update on %s\n", tool_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the tool.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "Indicates whether to forcefully delete the tool, even if it is.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("tool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTool on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteToolRequest{
								Name:  name,
								Force: cmd.Bool("force"),
								Etag:  cmd.String("etag"),
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
				Name:  "toolsets",
				Usage: "Manage toolsets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list toolsets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the toolsets.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListToolsetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListToolsets(ctx, req)
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
						Usage: "describe toolsets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "toolset", Usage: "The ID of the toolset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/toolsets/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("toolset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetToolsetRequest{
								Name: name,
							}

							resp, err := client.GetToolset(ctx, req)
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
						Usage: "create toolsets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "toolset-id", Usage: "The ID to use for the toolset, which will become the final.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateToolsetRequest{
								Parent:    parent,
								ToolsetId: cmd.String("toolset-id"),
							}

							resp, err := client.CreateToolset(ctx, req)
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
						Usage: "update toolsets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "toolset", Usage: "The ID of the toolset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "toolset.name" not yet supported.
							toolset_name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/toolsets/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("toolset"))
							fmt.Printf("Executing update on %s\n", toolset_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete toolsets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the toolset.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "Indicates whether to forcefully delete the toolset, even if it is.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "toolset", Usage: "The ID of the toolset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/toolsets/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("toolset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteToolset on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteToolsetRequest{
								Name:  name,
								Force: cmd.Bool("force"),
								Etag:  cmd.String("etag"),
							}

							if err := client.DeleteToolset(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "retrieve-tools",
						Usage: "retrieve-tools toolsets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "tool-ids", Usage: "The identifiers of the tools to retrieve from the toolset.", Required: false},
							&cli.StringFlag{Name: "toolset", Usage: "The ID of the toolset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							toolset := fmt.Sprintf("projects/%s/locations/%s/apps/%s/toolsets/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("toolset"))
							fmt.Printf("Executing retrieve-tools on %s\n", toolset)
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
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied when listing the app versions.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.ListAppVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAppVersions(ctx, req)
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.GetAppVersionRequest{
								Name: name,
							}

							resp, err := client.GetAppVersion(ctx, req)
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "app-version-id", Usage: "The ID to use for the app version, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apps/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.CreateAppVersionRequest{
								Parent:       parent,
								AppVersionId: cmd.String("app-version-id"),
							}

							resp, err := client.CreateAppVersion(ctx, req)
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the app version.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAppVersion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.DeleteAppVersionRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteAppVersion(ctx, req); err != nil {
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apps/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("app"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := ces.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cespb.RestoreAppVersionRequest{
								Name: name,
							}

							op, err := client.RestoreAppVersion(ctx, req)
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
		},
	}
}
