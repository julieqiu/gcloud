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

package discoveryengine

import (
	discoveryengine "cloud.google.com/go/discoveryengine/apiv1"
	"cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the discoveryengine command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "discoveryengine",
		Usage: "manage Discovery Engine API resources",
		Commands: []*cli.Command{
			{
				Name:  "answers",
				Usage: "Manage answers resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe answers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "answer", Usage: "The ID of the answer.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s/answers/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"), cmd.String("answer"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetAnswerRequest{
								Name: name,
							}

							resp, err := client.GetAnswer(ctx, req)
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
				Name:  "assistants",
				Usage: "Manage assistants resources",
				Commands: []*cli.Command{

					{
						Name:  "stream-assist",
						Usage: "stream-assist assistants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "assistant", Usage: "The ID of the assistant.", Required: true},
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "engine", Usage: "The ID of the engine.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session to use for the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s/assistants/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("engine"), cmd.String("assistant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.StreamAssistRequest{
								Name:    name,
								Session: cmd.String("session"),
							}

							resp, err := client.StreamAssist(ctx, req)
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
				Name:  "batch-get-documents-metadata",
				Usage: "Manage batch-get-documents-metadata resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-get-documents-metadata",
						Usage: "batch-get-documents-metadata batch-get-documents-metadata",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.BatchGetDocumentsMetadataRequest{
								Parent: parent,
							}

							resp, err := client.BatchGetDocumentsMetadata(ctx, req)
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
				Name:  "cmek-config",
				Usage: "Manage cmek-config resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update cmek-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "set-default", Usage: "Set the following CmekConfig as the default to be used for child.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "config.name" not yet supported.
							config_name := fmt.Sprintf("projects/%s/locations/%s/cmekConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", config_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe cmek-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cmekConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetCmekConfigRequest{
								Name: name,
							}

							resp, err := client.GetCmekConfig(ctx, req)
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
				Name:  "cmek-configs",
				Usage: "Manage cmek-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list cmek-configs",
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
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListCmekConfigsRequest{
								Parent: parent,
							}

							resp, err := client.ListCmekConfigs(ctx, req)
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
						Usage: "delete cmek-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cmek-config", Usage: "The ID of the cmek config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cmekConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("cmek-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCmekConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteCmekConfigRequest{
								Name: name,
							}

							op, err := client.DeleteCmekConfig(ctx, req)
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
				},
			},
			{
				Name:  "completion-suggestions",
				Usage: "Manage completion-suggestions resources",
				Commands: []*cli.Command{

					{
						Name:  "import",
						Usage: "import completion-suggestions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ImportCompletionSuggestionsRequest{
								Parent: parent,
							}

							op, err := client.ImportCompletionSuggestions(ctx, req)
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
						Name:  "purge",
						Usage: "purge completion-suggestions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.PurgeCompletionSuggestionsRequest{
								Parent: parent,
							}

							op, err := client.PurgeCompletionSuggestions(ctx, req)
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
				Name:  "controls",
				Usage: "Manage controls resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create controls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "control-id", Usage: "The ID to use for the Control, which will become the final.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateControlRequest{
								Parent:    parent,
								ControlId: cmd.String("control-id"),
							}

							resp, err := client.CreateControl(ctx, req)
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
						Usage: "delete controls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "control", Usage: "The ID of the control.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/controls/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("control"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteControl on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteControlRequest{
								Name: name,
							}

							if err := client.DeleteControl(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update controls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "control", Usage: "The ID of the control.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "control.name" not yet supported.
							control_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/controls/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("control"))
							fmt.Printf("Executing update on %s\n", control_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe controls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "control", Usage: "The ID of the control.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/controls/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("control"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetControlRequest{
								Name: name,
							}

							resp, err := client.GetControl(ctx, req)
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
						Usage: "list controls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter to apply on the list results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListControls` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListControlsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListControls(ctx, req)
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
				Name:  "conversations",
				Usage: "Manage conversations resources",
				Commands: []*cli.Command{

					{
						Name:  "converse",
						Usage: "converse conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter syntax consists of an expression language for constructing a.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "safe-search", Usage: "Whether to turn on safe search.", Required: false},
							&cli.StringFlag{Name: "serving-config", Usage: "The resource name of the Serving Config to use.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ConverseConversationRequest{
								Name:          name,
								ServingConfig: cmd.String("serving-config"),
								SafeSearch:    cmd.Bool("safe-search"),
								Filter:        cmd.String("filter"),
							}

							resp, err := client.ConverseConversation(ctx, req)
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
						Usage: "create conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateConversationRequest{
								Parent: parent,
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
						Name:  "delete",
						Usage: "delete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("conversation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteConversation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteConversationRequest{
								Name: name,
							}

							if err := client.DeleteConversation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "conversation.name" not yet supported.
							conversation_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("conversation"))
							fmt.Printf("Executing update on %s\n", conversation_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetConversationRequest{
								Name: name,
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
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter to apply on the list results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListConversations` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListConversationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
				},
			},
			{
				Name:  "custom-models",
				Usage: "Manage custom-models resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list custom-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							data_store := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							fmt.Printf("Executing list on %s\n", data_store)
							return nil
						},
					},
				},
			},
			{
				Name:  "data-stores",
				Usage: "Manage data-stores resources",
				Commands: []*cli.Command{

					{
						Name:  "complete-query",
						Usage: "complete-query data-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.BoolFlag{Name: "include-tail-suggestions", Usage: "Indicates if tail suggestions should be returned if there are no.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "The typeahead input used to fetch suggestions.", Required: true},
							&cli.StringFlag{Name: "query-model", Usage: "Specifies the autocomplete data model.", Required: false},
							&cli.StringFlag{Name: "user-pseudo-id", Usage: "A unique identifier for tracking visitors.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							data_store := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							fmt.Printf("Executing complete-query on %s\n", data_store)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create data-stores",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "create-advanced-site-search", Usage: "A boolean flag indicating whether user want to directly create an advanced.", Required: false},
							&cli.StringFlag{Name: "data-store-id", Usage: "The ID to use for the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-default-schema-creation", Usage: "A boolean flag indicating whether to skip the default schema creation for.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateDataStoreRequest{
								Parent:                    parent,
								DataStoreId:               cmd.String("data-store-id"),
								CreateAdvancedSiteSearch:  cmd.Bool("create-advanced-site-search"),
								SkipDefaultSchemaCreation: cmd.Bool("skip-default-schema-creation"),
							}

							op, err := client.CreateDataStore(ctx, req)
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
						Usage: "describe data-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetDataStoreRequest{
								Name: name,
							}

							resp, err := client.GetDataStore(ctx, req)
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
						Usage: "list data-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter by solution type.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [DataStore][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListDataStoresRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataStores(ctx, req)
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
						Usage: "delete data-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataStore %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteDataStoreRequest{
								Name: name,
							}

							op, err := client.DeleteDataStore(ctx, req)
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
						Name:  "update",
						Usage: "update data-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_store.name" not yet supported.
							data_store_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							fmt.Printf("Executing update on %s\n", data_store_name)
							return nil
						},
					},

					{
						Name:  "train-custom-model",
						Usage: "train-custom-model data-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-id", Usage: "If not provided, a UUID will be generated.", Required: false},
							&cli.StringFlag{Name: "model-type", Usage: "Model to be trained.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							data_store := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							fmt.Printf("Executing train-custom-model on %s\n", data_store)
							return nil
						},
					},
				},
			},
			{
				Name:  "documents",
				Usage: "Manage documents resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s/documents/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"), cmd.String("document"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetDocumentRequest{
								Name: name,
							}

							resp, err := client.GetDocument(ctx, req)
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
						Usage: "list documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [Document][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListDocumentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDocuments(ctx, req)
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
						Usage: "create documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "document-id", Usage: "The ID to use for the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateDocumentRequest{
								Parent:     parent,
								DocumentId: cmd.String("document-id"),
							}

							resp, err := client.CreateDocument(ctx, req)
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
						Usage: "update documents",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to `true` and the.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "document.name" not yet supported.
							document_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s/documents/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"), cmd.String("document"))
							fmt.Printf("Executing update on %s\n", document_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s/documents/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"), cmd.String("document"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDocument on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteDocumentRequest{
								Name: name,
							}

							if err := client.DeleteDocument(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "import",
						Usage: "import documents",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-generate-ids", Usage: "Whether to automatically generate IDs for the documents if absent.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.BoolFlag{Name: "force-refresh-content", Usage: "Whether to force refresh the unstructured content of the.", Required: false},
							&cli.StringFlag{Name: "id-field", Usage: "The field indicates the ID field or column to be used as unique IDs of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reconciliation-mode", Usage: "The mode of reconciliation between existing documents and the documents to.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ImportDocumentsRequest{
								Parent:              parent,
								ReconciliationMode:  discoveryenginepb.ImportDocumentsRequest_ReconciliationMode(discoveryenginepb.ImportDocumentsRequest_ReconciliationMode_value[cmd.String("reconciliation-mode")]),
								AutoGenerateIds:     cmd.Bool("auto-generate-ids"),
								IdField:             cmd.String("id-field"),
								ForceRefreshContent: cmd.Bool("force-refresh-content"),
							}

							op, err := client.ImportDocuments(ctx, req)
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
						Name:  "purge",
						Usage: "purge documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter matching documents to purge.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Actually performs the purge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.PurgeDocumentsRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
								Force:  cmd.Bool("force"),
							}

							op, err := client.PurgeDocuments(ctx, req)
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
				Name:  "engines",
				Usage: "Manage engines resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "engine-id", Usage: "The ID to use for the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateEngineRequest{
								Parent:   parent,
								EngineId: cmd.String("engine-id"),
							}

							op, err := client.CreateEngine(ctx, req)
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
						Usage: "delete engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "engine", Usage: "The ID of the engine.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("engine"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEngine %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteEngineRequest{
								Name: name,
							}

							op, err := client.DeleteEngine(ctx, req)
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
						Name:  "update",
						Usage: "update engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "engine", Usage: "The ID of the engine.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "engine.name" not yet supported.
							engine_name := fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("engine"))
							fmt.Printf("Executing update on %s\n", engine_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "engine", Usage: "The ID of the engine.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetEngineRequest{
								Name: name,
							}

							resp, err := client.GetEngine(ctx, req)
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
						Usage: "list engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter by solution type.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Not supported.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Not supported.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListEnginesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListEngines(ctx, req)
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
				Name:  "grounding-configs",
				Usage: "Manage grounding-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "check",
						Usage: "check grounding-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "answer-candidate", Usage: "Answer candidate to check.", Required: false},
							&cli.StringFlag{Name: "grounding-config", Usage: "The ID of the grounding config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							grounding_config := fmt.Sprintf("projects/%s/locations/%s/groundingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("grounding-config"))
							fmt.Printf("Executing check on %s\n", grounding_config)
							return nil
						},
					},
				},
			},
			{
				Name:  "identity-mapping-stores",
				Usage: "Manage identity-mapping-stores resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "identity-mapping-store-id", Usage: "The ID of the Identity Mapping Store to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateIdentityMappingStoreRequest{
								Parent:                 parent,
								IdentityMappingStoreId: cmd.String("identity-mapping-store-id"),
							}

							resp, err := client.CreateIdentityMappingStore(ctx, req)
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
						Usage: "describe identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "identity-mapping-store", Usage: "The ID of the identity mapping store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/identityMappingStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("identity-mapping-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetIdentityMappingStoreRequest{
								Name: name,
							}

							resp, err := client.GetIdentityMappingStore(ctx, req)
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
						Usage: "delete identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "identity-mapping-store", Usage: "The ID of the identity mapping store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/identityMappingStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("identity-mapping-store"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIdentityMappingStore %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteIdentityMappingStoreRequest{
								Name: name,
							}

							op, err := client.DeleteIdentityMappingStore(ctx, req)
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
						Name:  "import-identity-mappings",
						Usage: "import-identity-mappings identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "identity-mapping-store", Usage: "The ID of the identity mapping store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							identity_mapping_store := fmt.Sprintf("projects/%s/locations/%s/identityMappingStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("identity-mapping-store"))
							fmt.Printf("Executing import-identity-mappings on %s\n", identity_mapping_store)
							return nil
						},
					},

					{
						Name:  "purge-identity-mappings",
						Usage: "purge-identity-mappings identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter matching identity mappings to purge.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "Actually performs the purge.", Required: false},
							&cli.StringFlag{Name: "identity-mapping-store", Usage: "The ID of the identity mapping store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							identity_mapping_store := fmt.Sprintf("projects/%s/locations/%s/identityMappingStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("identity-mapping-store"))
							fmt.Printf("Executing purge-identity-mappings on %s\n", identity_mapping_store)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "identity-mapping-store", Usage: "The ID of the identity mapping store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of IdentityMappings to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListIdentityMappings` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							identity_mapping_store := fmt.Sprintf("projects/%s/locations/%s/identityMappingStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("identity-mapping-store"))
							fmt.Printf("Executing list on %s\n", identity_mapping_store)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list identity-mapping-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of IdentityMappingStores to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListIdentityMappingStores` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListIdentityMappingStoresRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListIdentityMappingStores(ctx, req)
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
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "stream-generate-grounded-content",
						Usage: "stream-generate-grounded-content locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing stream-generate-grounded-content on %s\n", location)
							return nil
						},
					},

					{
						Name:  "generate-grounded-content",
						Usage: "generate-grounded-content locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing generate-grounded-content on %s\n", location)
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
				Name:  "projects",
				Usage: "Manage projects resources",
				Commands: []*cli.Command{

					{
						Name:  "provision",
						Usage: "provision projects",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "accept-data-use-terms", Usage: "Set to `true` to specify that caller has read and would like to.", Required: true},
							&cli.StringFlag{Name: "data-use-terms-version", Usage: "The version of the [Terms for data.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ProvisionProjectRequest{
								Name:                name,
								AcceptDataUseTerms:  cmd.Bool("accept-data-use-terms"),
								DataUseTermsVersion: cmd.String("data-use-terms-version"),
							}

							op, err := client.ProvisionProject(ctx, req)
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
				Name:  "ranking-configs",
				Usage: "Manage ranking-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "rank",
						Usage: "rank ranking-configs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "ignore-record-details-in-response", Usage: "If true, the response will contain only record ID and score.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The identifier of the model to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "The query to use.", Required: false},
							&cli.StringFlag{Name: "ranking-config", Usage: "The ID of the ranking config.", Required: true},
							&cli.IntFlag{Name: "top-n", Usage: "The number of results to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							ranking_config := fmt.Sprintf("projects/%s/locations/%s/rankingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("ranking-config"))
							fmt.Printf("Executing rank on %s\n", ranking_config)
							return nil
						},
					},
				},
			},
			{
				Name:  "schemas",
				Usage: "Manage schemas resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetSchemaRequest{
								Name: name,
							}

							resp, err := client.GetSchema(ctx, req)
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
						Usage: "list schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of [Schema][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListSchemasRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSchemas(ctx, req)
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
						Usage: "create schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-id", Usage: "The ID to use for the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateSchemaRequest{
								Parent:   parent,
								SchemaId: cmd.String("schema-id"),
							}

							op, err := client.CreateSchema(ctx, req)
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
						Usage: "update schemas",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Schema][google.", Required: false},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "schema.name" not yet supported.
							schema_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("schema"))
							fmt.Printf("Executing update on %s\n", schema_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("schema"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSchema %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteSchemaRequest{
								Name: name,
							}

							op, err := client.DeleteSchema(ctx, req)
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
				},
			},
			{
				Name:  "serving-configs",
				Usage: "Manage serving-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "answer",
						Usage: "answer serving-configs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "asynchronous-mode", Usage: "Deprecated: This field is deprecated.", Required: false},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session resource name.", Required: false},
							&cli.StringFlag{Name: "user-pseudo-id", Usage: "A unique identifier for tracking visitors.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("serving-config"))
							fmt.Printf("Executing answer on %s\n", serving_config)
							return nil
						},
					},

					{
						Name:  "stream-answer",
						Usage: "stream-answer serving-configs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "asynchronous-mode", Usage: "Deprecated: This field is deprecated.", Required: false},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session resource name.", Required: false},
							&cli.StringFlag{Name: "user-pseudo-id", Usage: "A unique identifier for tracking visitors.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("serving-config"))
							fmt.Printf("Executing stream-answer on %s\n", serving_config)
							return nil
						},
					},

					{
						Name:  "recommend",
						Usage: "recommend serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter for restricting recommendation results with a length limit of 5,000.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Use validate only mode for this recommendation query.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("serving-config"))
							fmt.Printf("Executing recommend on %s\n", serving_config)
							return nil
						},
					},

					{
						Name:  "search",
						Usage: "search serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch", Usage: "The branch resource name, such as.", Required: false},
							&cli.StringFlag{Name: "canonical-filter", Usage: "The default filter that is applied when a user performs a search without.", Required: false},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter syntax consists of an expression language for constructing a.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The BCP-47 language code, such as \"en-US\" or \"sr-Latn\".", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "offset", Usage: "A 0-indexed integer that specifies the current offset (that is, starting.", Required: false},
							&cli.IntFlag{Name: "one-box-page-size", Usage: "The maximum number of results to return for OneBox.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "The order in which documents are returned.", Required: false},
							&cli.StringSliceFlag{Name: "page-categories", Usage: "The categories associated with a category page.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [Document][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Raw search query.", Required: false},
							&cli.StringFlag{Name: "ranking-expression", Usage: "The ranking expression controls the customized ranking on.", Required: false},
							&cli.StringFlag{Name: "ranking-expression-backend", Usage: "The backend to use for the ranking expression evaluation.", Required: false},
							&cli.StringFlag{Name: "relevance-threshold", Usage: "The global relevance threshold of the search results.", Required: false},
							&cli.BoolFlag{Name: "safe-search", Usage: "Whether to turn on safe search.", Required: false},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session resource name.", Required: false},
							&cli.StringFlag{Name: "user-pseudo-id", Usage: "A unique identifier for tracking visitors.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("serving-config"))
							fmt.Printf("Executing search on %s\n", serving_config)
							return nil
						},
					},

					{
						Name:  "search-lite",
						Usage: "search-lite serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch", Usage: "The branch resource name, such as.", Required: false},
							&cli.StringFlag{Name: "canonical-filter", Usage: "The default filter that is applied when a user performs a search without.", Required: false},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter syntax consists of an expression language for constructing a.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The BCP-47 language code, such as \"en-US\" or \"sr-Latn\".", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "offset", Usage: "A 0-indexed integer that specifies the current offset (that is, starting.", Required: false},
							&cli.IntFlag{Name: "one-box-page-size", Usage: "The maximum number of results to return for OneBox.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "The order in which documents are returned.", Required: false},
							&cli.StringSliceFlag{Name: "page-categories", Usage: "The categories associated with a category page.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [Document][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Raw search query.", Required: false},
							&cli.StringFlag{Name: "ranking-expression", Usage: "The ranking expression controls the customized ranking on.", Required: false},
							&cli.StringFlag{Name: "ranking-expression-backend", Usage: "The backend to use for the ranking expression evaluation.", Required: false},
							&cli.StringFlag{Name: "relevance-threshold", Usage: "The global relevance threshold of the search results.", Required: false},
							&cli.BoolFlag{Name: "safe-search", Usage: "Whether to turn on safe search.", Required: false},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session resource name.", Required: false},
							&cli.StringFlag{Name: "user-pseudo-id", Usage: "A unique identifier for tracking visitors.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("serving-config"))
							fmt.Printf("Executing search-lite on %s\n", serving_config)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "serving_config.name" not yet supported.
							serving_config_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("serving-config"))
							fmt.Printf("Executing update on %s\n", serving_config_name)
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
						Name:  "create",
						Usage: "create sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateSessionRequest{
								Parent: parent,
							}

							resp, err := client.CreateSession(ctx, req)
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
						Usage: "delete sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSession on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteSessionRequest{
								Name: name,
							}

							if err := client.DeleteSession(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "session.name" not yet supported.
							session_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"))
							fmt.Printf("Executing update on %s\n", session_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.BoolFlag{Name: "include-answer-details", Usage: "If set to true, the full session including all answer details.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetSessionRequest{
								Name:                 name,
								IncludeAnswerDetails: cmd.Bool("include-answer-details"),
							}

							resp, err := client.GetSession(ctx, req)
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
						Usage: "list sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A comma-separated list of fields to filter by, in EBNF grammar.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSessions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListSessionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSessions(ctx, req)
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
						Usage: "create sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateSessionRequest{
								Parent: parent,
							}

							resp, err := client.CreateSession(ctx, req)
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
						Usage: "delete sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSession on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteSessionRequest{
								Name: name,
							}

							if err := client.DeleteSession(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "session.name" not yet supported.
							session_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"))
							fmt.Printf("Executing update on %s\n", session_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.BoolFlag{Name: "include-answer-details", Usage: "If set to true, the full session including all answer details.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetSessionRequest{
								Name:                 name,
								IncludeAnswerDetails: cmd.Bool("include-answer-details"),
							}

							resp, err := client.GetSession(ctx, req)
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
						Usage: "list sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A comma-separated list of fields to filter by, in EBNF grammar.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSessions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListSessionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSessions(ctx, req)
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
				Name:  "site-search-engine",
				Usage: "Manage site-search-engine resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe site-search-engine",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetSiteSearchEngineRequest{
								Name: name,
							}

							resp, err := client.GetSiteSearchEngine(ctx, req)
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
						Name:  "enable-advanced-site-search",
						Usage: "enable-advanced-site-search site-search-engine",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "site-search-engine", Usage: "Full resource name of the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							site_search_engine := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							fmt.Printf("Executing enable-advanced-site-search on %s\n", site_search_engine)
							return nil
						},
					},

					{
						Name:  "disable-advanced-site-search",
						Usage: "disable-advanced-site-search site-search-engine",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "site-search-engine", Usage: "Full resource name of the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							site_search_engine := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							fmt.Printf("Executing disable-advanced-site-search on %s\n", site_search_engine)
							return nil
						},
					},

					{
						Name:  "recrawl-uris",
						Usage: "recrawl-uris site-search-engine",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "site-credential", Usage: "Credential id to use for crawling.", Required: false},
							&cli.StringFlag{Name: "site-search-engine", Usage: "Full resource name of the.", Required: true},
							&cli.StringSliceFlag{Name: "uris", Usage: "List of URIs to crawl.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							site_search_engine := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							fmt.Printf("Executing recrawl-uris on %s\n", site_search_engine)
							return nil
						},
					},

					{
						Name:  "batch-verify-target-sites",
						Usage: "batch-verify-target-sites site-search-engine",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.BatchVerifyTargetSitesRequest{
								Parent: parent,
							}

							op, err := client.BatchVerifyTargetSites(ctx, req)
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
						Name:  "fetch-domain-verification-status",
						Usage: "fetch-domain-verification-status site-search-engine",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `FetchDomainVerificationStatus`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "site-search-engine", Usage: "The site search engine resource under which we fetch all the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							site_search_engine := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							fmt.Printf("Executing fetch-domain-verification-status on %s\n", site_search_engine)
							return nil
						},
					},
				},
			},
			{
				Name:  "sitemaps",
				Usage: "Manage sitemaps resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create sitemaps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateSitemapRequest{
								Parent: parent,
							}

							op, err := client.CreateSitemap(ctx, req)
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
						Usage: "delete sitemaps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sitemap", Usage: "The ID of the sitemap.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine/sitemaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("sitemap"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSitemap %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteSitemapRequest{
								Name: name,
							}

							op, err := client.DeleteSitemap(ctx, req)
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
						Name:  "fetch",
						Usage: "fetch sitemaps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.FetchSitemapsRequest{
								Parent: parent,
							}

							resp, err := client.FetchSitemaps(ctx, req)
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
				Name:  "suggestion-deny-list-entries",
				Usage: "Manage suggestion-deny-list-entries resources",
				Commands: []*cli.Command{

					{
						Name:  "import",
						Usage: "import suggestion-deny-list-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ImportSuggestionDenyListEntriesRequest{
								Parent: parent,
							}

							op, err := client.ImportSuggestionDenyListEntries(ctx, req)
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
						Name:  "purge",
						Usage: "purge suggestion-deny-list-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("collection"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.PurgeSuggestionDenyListEntriesRequest{
								Parent: parent,
							}

							op, err := client.PurgeSuggestionDenyListEntries(ctx, req)
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
				Name:  "target-sites",
				Usage: "Manage target-sites resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create target-sites",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CreateTargetSiteRequest{
								Parent: parent,
							}

							op, err := client.CreateTargetSite(ctx, req)
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
						Name:  "batch-create",
						Usage: "batch-create target-sites",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.BatchCreateTargetSitesRequest{
								Parent: parent,
							}

							op, err := client.BatchCreateTargetSites(ctx, req)
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
						Usage: "describe target-sites",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target-site", Usage: "The ID of the target site.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine/targetSites/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("target-site"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.GetTargetSiteRequest{
								Name: name,
							}

							resp, err := client.GetTargetSite(ctx, req)
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
						Usage: "update target-sites",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target-site", Usage: "The ID of the target site.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "target_site.name" not yet supported.
							target_site_name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine/targetSites/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("target-site"))
							fmt.Printf("Executing update on %s\n", target_site_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-sites",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target-site", Usage: "The ID of the target site.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine/targetSites/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"), cmd.String("target-site"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTargetSite %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.DeleteTargetSiteRequest{
								Name: name,
							}

							op, err := client.DeleteTargetSite(ctx, req)
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
						Name:  "list",
						Usage: "list target-sites",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListTargetSites` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s/siteSearchEngine", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListTargetSitesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTargetSites(ctx, req)
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
				Name:  "user-events",
				Usage: "Manage user-events resources",
				Commands: []*cli.Command{

					{
						Name:  "write",
						Usage: "write user-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "write-async", Usage: "If set to true, the user event is written asynchronously after.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.WriteUserEventRequest{
								Parent:     parent,
								WriteAsync: cmd.Bool("write-async"),
							}

							resp, err := client.WriteUserEvent(ctx, req)
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
						Name:  "collect",
						Usage: "collect user-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.IntFlag{Name: "ets", Usage: "The event timestamp in milliseconds.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "uri", Usage: "The URL including cgi-parameters but excluding the hash fragment with a.", Required: false},
							&cli.StringFlag{Name: "user-event", Usage: "URL encoded UserEvent proto with a length limit of 2,000,000.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.CollectUserEventRequest{
								Parent:    parent,
								UserEvent: cmd.String("user-event"),
								Uri:       runtime.Ptr(cmd.String("uri")),
								Ets:       runtime.Ptr(cmd.Int("ets")),
							}

							resp, err := client.CollectUserEvent(ctx, req)
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
						Name:  "purge",
						Usage: "purge user-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter string to specify the events to be deleted with a.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "The `force` field is currently not supported.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.PurgeUserEventsRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
								Force:  cmd.Bool("force"),
							}

							op, err := client.PurgeUserEvents(ctx, req)
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
						Usage: "import user-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-store", Usage: "The ID of the data store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ImportUserEventsRequest{
								Parent: parent,
							}

							op, err := client.ImportUserEvents(ctx, req)
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
				Name:  "user-licenses",
				Usage: "Manage user-licenses resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list user-licenses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for the list request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListUserLicenses` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-store", Usage: "The ID of the user store.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/userStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("user-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.ListUserLicensesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListUserLicenses(ctx, req)
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
				Name:  "user-stores",
				Usage: "Manage user-stores resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-update-user-licenses",
						Usage: "batch-update-user-licenses user-stores",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "delete-unassigned-user-licenses", Usage: "If true, if user licenses removed associated license config, the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-store", Usage: "The ID of the user store.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/userStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("user-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := discoveryengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &discoveryenginepb.BatchUpdateUserLicensesRequest{
								Parent:                       parent,
								DeleteUnassignedUserLicenses: cmd.Bool("delete-unassigned-user-licenses"),
							}

							op, err := client.BatchUpdateUserLicenses(ctx, req)
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
