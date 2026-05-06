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

package dialogflow

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
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
				Name:  "agent",
				Usage: "Manage agent resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetAgentRequest{
								Parent: parent,
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
						Name:  "set-agent",
						Usage: "set-agent agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "agent.parent" not yet supported.
							agent_parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing set-agent on %s\n", agent_parent)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAgent on %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteAgentRequest{
								Parent: parent,
							}

							if err := client.DeleteAgent(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "search",
						Usage: "search agent",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.SearchAgentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchAgents(ctx, req)
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
						Name:  "train",
						Usage: "train agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("TrainAgent %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.TrainAgentRequest{
								Parent: parent,
							}

							op, err := client.TrainAgent(ctx, req)
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
						Name:  "export",
						Usage: "export agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "agent-uri", Usage: "The [Google Cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ExportAgentRequest{
								Parent:   parent,
								AgentUri: cmd.String("agent-uri"),
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
						Name:  "import",
						Usage: "import agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("ImportAgent %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ImportAgentRequest{
								Parent: parent,
							}

							op, err := client.ImportAgent(ctx, req)
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
						Name:  "restore",
						Usage: "restore agent",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("RestoreAgent %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.RestoreAgentRequest{
								Parent: parent,
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
				},
			},
			{
				Name:  "answer-records",
				Usage: "Manage answer-records resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list answer-records",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filters to restrict results to specific answer records.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of records to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListAnswerRecordsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAnswerRecords(ctx, req)
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
						Usage: "update answer-records",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "answer-record", Usage: "The ID of the answer record.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "answer_record.name" not yet supported.
							answer_record_name := fmt.Sprintf("projects/%s/answerRecords/%s", cmd.String("project"), cmd.String("answer-record"))
							fmt.Printf("Executing update on %s\n", answer_record_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "contexts",
				Usage: "Manage contexts resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list contexts",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/sessions/%s", cmd.String("project"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListContextsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListContexts(ctx, req)
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
						Usage: "describe contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/sessions/%s/contexts/%s", cmd.String("project"), cmd.String("session"), cmd.String("context"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetContextRequest{
								Name: name,
							}

							resp, err := client.GetContext(ctx, req)
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
						Usage: "create contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/sessions/%s", cmd.String("project"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateContextRequest{
								Parent: parent,
							}

							resp, err := client.CreateContext(ctx, req)
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
						Usage: "update contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "context.name" not yet supported.
							context_name := fmt.Sprintf("projects/%s/agent/sessions/%s/contexts/%s", cmd.String("project"), cmd.String("session"), cmd.String("context"))
							fmt.Printf("Executing update on %s\n", context_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/sessions/%s/contexts/%s", cmd.String("project"), cmd.String("session"), cmd.String("context"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteContext on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteContextRequest{
								Name: name,
							}

							if err := client.DeleteContext(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/sessions/%s", cmd.String("project"), cmd.String("session"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAllContexts on %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteAllContextsRequest{
								Parent: parent,
							}

							if err := client.DeleteAllContexts(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "conversation-datasets",
				Usage: "Manage conversation-datasets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create conversation-datasets",
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
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateConversationDatasetRequest{
								Parent: parent,
							}

							op, err := client.CreateConversationDataset(ctx, req)
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
						Usage: "describe conversation-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-dataset", Usage: "The ID of the conversation dataset.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationDatasets/%s", cmd.String("project"), cmd.String("conversation-dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetConversationDatasetRequest{
								Name: name,
							}

							resp, err := client.GetConversationDataset(ctx, req)
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
						Usage: "list conversation-datasets",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of conversation datasets to return in a single.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListConversationDatasetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListConversationDatasets(ctx, req)
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
						Usage: "delete conversation-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-dataset", Usage: "The ID of the conversation dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversationDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation-dataset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConversationDataset %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteConversationDatasetRequest{
								Name: name,
							}

							op, err := client.DeleteConversationDataset(ctx, req)
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
						Name:  "import-conversation-data",
						Usage: "import-conversation-data conversation-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-dataset", Usage: "The ID of the conversation dataset.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationDatasets/%s", cmd.String("project"), cmd.String("conversation-dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ImportConversationDataRequest{
								Name: name,
							}

							op, err := client.ImportConversationData(ctx, req)
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
				Name:  "conversation-models",
				Usage: "Manage conversation-models resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create conversation-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateConversationModelRequest{
								Parent: parent,
							}

							op, err := client.CreateConversationModel(ctx, req)
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
						Usage: "describe conversation-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationModels/%s", cmd.String("project"), cmd.String("conversation-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetConversationModelRequest{
								Name: name,
							}

							resp, err := client.GetConversationModel(ctx, req)
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
						Usage: "list conversation-models",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of conversation models to return in a single.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListConversationModelsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListConversationModels(ctx, req)
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
						Usage: "delete conversation-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationModels/%s", cmd.String("project"), cmd.String("conversation-model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConversationModel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteConversationModelRequest{
								Name: name,
							}

							op, err := client.DeleteConversationModel(ctx, req)
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
						Usage: "deploy conversation-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationModels/%s", cmd.String("project"), cmd.String("conversation-model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeployConversationModel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeployConversationModelRequest{
								Name: name,
							}

							op, err := client.DeployConversationModel(ctx, req)
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
						Name:  "undeploy",
						Usage: "undeploy conversation-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationModels/%s", cmd.String("project"), cmd.String("conversation-model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("UndeployConversationModel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.UndeployConversationModelRequest{
								Name: name,
							}

							op, err := client.UndeployConversationModel(ctx, req)
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
				Name:  "conversation-profiles",
				Usage: "Manage conversation-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list conversation-profiles",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListConversationProfilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListConversationProfiles(ctx, req)
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
						Usage: "describe conversation-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-profile", Usage: "The ID of the conversation profile.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationProfiles/%s", cmd.String("project"), cmd.String("conversation-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetConversationProfileRequest{
								Name: name,
							}

							resp, err := client.GetConversationProfile(ctx, req)
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
						Usage: "create conversation-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateConversationProfileRequest{
								Parent: parent,
							}

							resp, err := client.CreateConversationProfile(ctx, req)
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
						Usage: "update conversation-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-profile", Usage: "The ID of the conversation profile.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "conversation_profile.name" not yet supported.
							conversation_profile_name := fmt.Sprintf("projects/%s/conversationProfiles/%s", cmd.String("project"), cmd.String("conversation-profile"))
							fmt.Printf("Executing update on %s\n", conversation_profile_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete conversation-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-profile", Usage: "The ID of the conversation profile.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationProfiles/%s", cmd.String("project"), cmd.String("conversation-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteConversationProfile on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteConversationProfileRequest{
								Name: name,
							}

							if err := client.DeleteConversationProfile(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "set-suggestion-feature-config",
						Usage: "set-suggestion-feature-config conversation-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-profile", Usage: "The ID of the conversation profile.", Required: true},
							&cli.StringFlag{Name: "participant-role", Usage: "The participant role to add or update the suggestion feature.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversation_profile := fmt.Sprintf("projects/%s/conversationProfiles/%s", cmd.String("project"), cmd.String("conversation-profile"))
							fmt.Printf("Executing set-suggestion-feature-config on %s\n", conversation_profile)
							return nil
						},
					},

					{
						Name:  "clear-suggestion-feature-config",
						Usage: "clear-suggestion-feature-config conversation-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-profile", Usage: "The ID of the conversation profile.", Required: true},
							&cli.StringFlag{Name: "participant-role", Usage: "The participant role to remove the suggestion feature.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "suggestion-feature-type", Usage: "The type of the suggestion feature to remove.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversation_profile := fmt.Sprintf("projects/%s/conversationProfiles/%s", cmd.String("project"), cmd.String("conversation-profile"))
							fmt.Printf("Executing clear-suggestion-feature-config on %s\n", conversation_profile)
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
							&cli.StringFlag{Name: "conversation-id", Usage: "Identifier of the conversation.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateConversationRequest{
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
						Name:  "list",
						Usage: "list conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters conversations listed in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListConversationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
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
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetConversationRequest{
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
						Name:  "complete",
						Usage: "complete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CompleteConversationRequest{
								Name: name,
							}

							resp, err := client.CompleteConversation(ctx, req)
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
						Name:  "ingest-context-references",
						Usage: "ingest-context-references conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversation := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							fmt.Printf("Executing ingest-context-references on %s\n", conversation)
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
						Name:  "list",
						Usage: "list documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression used to filter documents returned by the list method.", Required: false},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/knowledgeBases/%s", cmd.String("project"), cmd.String("knowledge-base"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListDocumentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
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
						Name:  "describe",
						Usage: "describe documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/knowledgeBases/%s/documents/%s", cmd.String("project"), cmd.String("knowledge-base"), cmd.String("document"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetDocumentRequest{
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
						Name:  "create",
						Usage: "create documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/knowledgeBases/%s", cmd.String("project"), cmd.String("knowledge-base"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateDocumentRequest{
								Parent: parent,
							}

							op, err := client.CreateDocument(ctx, req)
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
						Usage: "import documents",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "import-gcs-custom-metadata", Usage: "Whether to import custom metadata from Google Cloud Storage.", Required: false},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/knowledgeBases/%s", cmd.String("project"), cmd.String("knowledge-base"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ImportDocumentsRequest{
								Parent:                  parent,
								ImportGcsCustomMetadata: cmd.Bool("import-gcs-custom-metadata"),
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
						Name:  "delete",
						Usage: "delete documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/knowledgeBases/%s/documents/%s", cmd.String("project"), cmd.String("knowledge-base"), cmd.String("document"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDocument %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteDocumentRequest{
								Name: name,
							}

							op, err := client.DeleteDocument(ctx, req)
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
						Usage: "update documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "document.name" not yet supported.
							document_name := fmt.Sprintf("projects/%s/knowledgeBases/%s/documents/%s", cmd.String("project"), cmd.String("knowledge-base"), cmd.String("document"))
							fmt.Printf("Executing update on %s\n", document_name)
							return nil
						},
					},

					{
						Name:  "reload",
						Usage: "reload documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.BoolFlag{Name: "import-gcs-custom-metadata", Usage: "Whether to import custom metadata from Google Cloud Storage.", Required: false},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "smart-messaging-partial-update", Usage: "When enabled, the reload request is to apply partial update to.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/knowledgeBases/%s/documents/%s", cmd.String("project"), cmd.String("knowledge-base"), cmd.String("document"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ReloadDocumentRequest{
								Name:                        name,
								ImportGcsCustomMetadata:     cmd.Bool("import-gcs-custom-metadata"),
								SmartMessagingPartialUpdate: cmd.Bool("smart-messaging-partial-update"),
							}

							op, err := client.ReloadDocument(ctx, req)
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
						Usage: "export documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.BoolFlag{Name: "export-full-content", Usage: "When enabled, export the full content of the document including empirical.", Required: false},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "smart-messaging-partial-update", Usage: "When enabled, export the smart messaging allowlist document for partial.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/knowledgeBases/%s/documents/%s", cmd.String("project"), cmd.String("knowledge-base"), cmd.String("document"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ExportDocumentRequest{
								Name:                        name,
								ExportFullContent:           cmd.Bool("export-full-content"),
								SmartMessagingPartialUpdate: cmd.Bool("smart-messaging-partial-update"),
							}

							op, err := client.ExportDocument(ctx, req)
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
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetEncryptionSpecRequest{
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
				Name:  "entities",
				Usage: "Manage entities resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-create",
						Usage: "batch-create entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/entityTypes/%s", cmd.String("project"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("BatchCreateEntities %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchCreateEntitiesRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							op, err := client.BatchCreateEntities(ctx, req)
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
						Name:  "batch-update",
						Usage: "batch-update entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/entityTypes/%s", cmd.String("project"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("BatchUpdateEntities %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchUpdateEntitiesRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							op, err := client.BatchUpdateEntities(ctx, req)
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
						Name:  "batch-delete",
						Usage: "batch-delete entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringSliceFlag{Name: "entity-values", Usage: "The reference `values` of the entities to delete.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/entityTypes/%s", cmd.String("project"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("BatchDeleteEntities %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchDeleteEntitiesRequest{
								Parent:       parent,
								EntityValues: cmd.StringSlice("entity-values"),
								LanguageCode: cmd.String("language-code"),
							}

							op, err := client.BatchDeleteEntities(ctx, req)
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
				Name:  "entity-types",
				Usage: "Manage entity-types resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListEntityTypesRequest{
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
						Name:  "describe",
						Usage: "describe entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/entityTypes/%s", cmd.String("project"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetEntityTypeRequest{
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
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateEntityTypeRequest{
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
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entity_type.name" not yet supported.
							entity_type_name := fmt.Sprintf("projects/%s/agent/entityTypes/%s", cmd.String("project"), cmd.String("entity-type"))
							fmt.Printf("Executing update on %s\n", entity_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/entityTypes/%s", cmd.String("project"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEntityType on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteEntityTypeRequest{
								Name: name,
							}

							if err := client.DeleteEntityType(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "batch-update",
						Usage: "batch-update entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchUpdateEntityTypesRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							op, err := client.BatchUpdateEntityTypes(ctx, req)
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
						Name:  "batch-delete",
						Usage: "batch-delete entity-types",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "entity-type-names", Usage: "The names entity types to delete.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("BatchDeleteEntityTypes %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchDeleteEntityTypesRequest{
								Parent:          parent,
								EntityTypeNames: cmd.StringSlice("entity-type-names"),
							}

							op, err := client.BatchDeleteEntityTypes(ctx, req)
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
						Usage: "list entity-types",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/sessions/%s", cmd.String("project"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListSessionEntityTypesRequest{
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
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("session"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetSessionEntityTypeRequest{
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/sessions/%s", cmd.String("project"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateSessionEntityTypeRequest{
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
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "session_entity_type.name" not yet supported.
							session_entity_type_name := fmt.Sprintf("projects/%s/agent/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("session"), cmd.String("entity-type"))
							fmt.Printf("Executing update on %s\n", session_entity_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("session"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSessionEntityType on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteSessionEntityTypeRequest{
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
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListEnvironmentsRequest{
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
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/environments/%s", cmd.String("project"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetEnvironmentRequest{
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
							&cli.StringFlag{Name: "environment-id", Usage: "The unique id of the new environment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateEnvironmentRequest{
								Parent:        parent,
								EnvironmentId: cmd.String("environment-id"),
							}

							resp, err := client.CreateEnvironment(ctx, req)
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
							&cli.BoolFlag{Name: "allow-load-to-draft-and-discard-changes", Usage: "This field is used to prevent accidental overwrite of the default.", Required: false},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "environment.name" not yet supported.
							environment_name := fmt.Sprintf("projects/%s/agent/environments/%s", cmd.String("project"), cmd.String("environment"))
							fmt.Printf("Executing update on %s\n", environment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/environments/%s", cmd.String("project"), cmd.String("environment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEnvironment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteEnvironmentRequest{
								Name: name,
							}

							if err := client.DeleteEnvironment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "evaluations",
				Usage: "Manage evaluations resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversationModels/%s/evaluations/%s", cmd.String("project"), cmd.String("conversation-model"), cmd.String("evaluation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetConversationModelEvaluationRequest{
								Name: name,
							}

							resp, err := client.GetConversationModelEvaluation(ctx, req)
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
						Usage: "list evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of evaluations to return in a.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversationModels/%s", cmd.String("project"), cmd.String("conversation-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListConversationModelEvaluationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListConversationModelEvaluations(ctx, req)
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
						Usage: "create evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation-model", Usage: "The ID of the conversation model.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversationModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation-model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateConversationModelEvaluationRequest{
								Parent: parent,
							}

							op, err := client.CreateConversationModelEvaluation(ctx, req)
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
						Name:  "create",
						Usage: "create evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateGeneratorEvaluationRequest{
								Parent: parent,
							}

							op, err := client.CreateGeneratorEvaluation(ctx, req)
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
						Usage: "describe evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/generators/%s/evaluations/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"), cmd.String("evaluation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetGeneratorEvaluationRequest{
								Name: name,
							}

							resp, err := client.GetGeneratorEvaluation(ctx, req)
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
						Usage: "list evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of evaluations to return in a.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListGeneratorEvaluationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGeneratorEvaluations(ctx, req)
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
						Usage: "delete evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/generators/%s/evaluations/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"), cmd.String("evaluation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGeneratorEvaluation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteGeneratorEvaluationRequest{
								Name: name,
							}

							if err := client.DeleteGeneratorEvaluation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "fulfillment",
				Usage: "Manage fulfillment resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe fulfillment",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/fulfillment", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetFulfillmentRequest{
								Name: name,
							}

							resp, err := client.GetFulfillment(ctx, req)
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
						Usage: "update fulfillment",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "fulfillment.name" not yet supported.
							fulfillment_name := fmt.Sprintf("projects/%s/agent/fulfillment", cmd.String("project"))
							fmt.Printf("Executing update on %s\n", fulfillment_name)
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
						Name:  "create",
						Usage: "create generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "generator-id", Usage: "The ID to use for the generator, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateGeneratorRequest{
								Parent:      parent,
								GeneratorId: cmd.String("generator-id"),
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
						Name:  "describe",
						Usage: "describe generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetGeneratorRequest{
								Name: name,
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
						Name:  "list",
						Usage: "list generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of conversation models to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListGeneratorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
						Name:  "delete",
						Usage: "delete generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGenerator on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteGeneratorRequest{
								Name: name,
							}

							if err := client.DeleteGenerator(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update generators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "generator", Usage: "The ID of the generator.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "generator.name" not yet supported.
							generator_name := fmt.Sprintf("projects/%s/locations/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("generator"))
							fmt.Printf("Executing update on %s\n", generator_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "history",
				Usage: "Manage history resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe history",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent/environments/%s", cmd.String("project"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetEnvironmentHistoryRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.GetEnvironmentHistory(ctx, req)
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
				Name:  "intents",
				Usage: "Manage intents resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intent-view", Usage: "The resource view to apply to the returned intent.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListIntentsRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								IntentView:   dialogflowpb.IntentView(dialogflowpb.IntentView_value[cmd.String("intent-view")]),
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
							&cli.StringFlag{Name: "intent", Usage: "The ID of the intent.", Required: true},
							&cli.StringFlag{Name: "intent-view", Usage: "The resource view to apply to the returned intent.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/intents/%s", cmd.String("project"), cmd.String("intent"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetIntentRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
								IntentView:   dialogflowpb.IntentView(dialogflowpb.IntentView_value[cmd.String("intent-view")]),
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
							&cli.StringFlag{Name: "intent-view", Usage: "The resource view to apply to the returned intent.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateIntentRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								IntentView:   dialogflowpb.IntentView(dialogflowpb.IntentView_value[cmd.String("intent-view")]),
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
							&cli.StringFlag{Name: "intent", Usage: "The ID of the intent.", Required: true},
							&cli.StringFlag{Name: "intent-view", Usage: "The resource view to apply to the returned intent.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "intent.name" not yet supported.
							intent_name := fmt.Sprintf("projects/%s/agent/intents/%s", cmd.String("project"), cmd.String("intent"))
							fmt.Printf("Executing update on %s\n", intent_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intent", Usage: "The ID of the intent.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/intents/%s", cmd.String("project"), cmd.String("intent"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIntent on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteIntentRequest{
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
						Name:  "batch-update",
						Usage: "batch-update intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intent-view", Usage: "The resource view to apply to the returned intent.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language used to access language-specific data.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchUpdateIntentsRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
								IntentView:   dialogflowpb.IntentView(dialogflowpb.IntentView_value[cmd.String("intent-view")]),
							}

							op, err := client.BatchUpdateIntents(ctx, req)
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
						Name:  "batch-delete",
						Usage: "batch-delete intents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("BatchDeleteIntents %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.BatchDeleteIntentsRequest{
								Parent: parent,
							}

							op, err := client.BatchDeleteIntents(ctx, req)
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
				Name:  "knowledge-bases",
				Usage: "Manage knowledge-bases resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list knowledge-bases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression used to filter knowledge bases returned by the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListKnowledgeBasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListKnowledgeBases(ctx, req)
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
						Usage: "describe knowledge-bases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/knowledgeBases/%s", cmd.String("project"), cmd.String("knowledge-base"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetKnowledgeBaseRequest{
								Name: name,
							}

							resp, err := client.GetKnowledgeBase(ctx, req)
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
						Usage: "create knowledge-bases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateKnowledgeBaseRequest{
								Parent: parent,
							}

							resp, err := client.CreateKnowledgeBase(ctx, req)
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
						Usage: "delete knowledge-bases",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "Force deletes the knowledge base.", Required: false},
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/knowledgeBases/%s", cmd.String("project"), cmd.String("knowledge-base"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteKnowledgeBase on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteKnowledgeBaseRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteKnowledgeBase(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update knowledge-bases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "knowledge-base", Usage: "The ID of the knowledge base.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "knowledge_base.name" not yet supported.
							knowledge_base_name := fmt.Sprintf("projects/%s/knowledgeBases/%s", cmd.String("project"), cmd.String("knowledge-base"))
							fmt.Printf("Executing update on %s\n", knowledge_base_name)
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
				Name:  "messages",
				Usage: "Manage messages resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list messages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter on message fields.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListMessagesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMessages(ctx, req)
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
				Name:  "participants",
				Usage: "Manage participants resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateParticipantRequest{
								Parent: parent,
							}

							resp, err := client.CreateParticipant(ctx, req)
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
						Usage: "describe participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetParticipantRequest{
								Name: name,
							}

							resp, err := client.GetParticipant(ctx, req)
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
						Usage: "list participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListParticipantsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListParticipants(ctx, req)
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
						Usage: "update participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "participant.name" not yet supported.
							participant_name := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							fmt.Printf("Executing update on %s\n", participant_name)
							return nil
						},
					},

					{
						Name:  "analyze-content",
						Usage: "analyze-content participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							participant := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							fmt.Printf("Executing analyze-content on %s\n", participant)
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
							&cli.StringFlag{Name: "input-audio", Usage: "The natural language speech audio to be processed.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/agent/sessions/%s", cmd.String("project"), cmd.String("session"))
							fmt.Printf("Executing detect-intent on %s\n", session)
							return nil
						},
					},
				},
			},
			{
				Name:  "sip-trunks",
				Usage: "Manage sip-trunks resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create sip-trunks",
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
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateSipTrunkRequest{
								Parent: parent,
							}

							resp, err := client.CreateSipTrunk(ctx, req)
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
						Usage: "delete sip-trunks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sip-trunk", Usage: "The ID of the sip trunk.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sipTrunks/%s", cmd.String("project"), cmd.String("location"), cmd.String("sip-trunk"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSipTrunk on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteSipTrunkRequest{
								Name: name,
							}

							if err := client.DeleteSipTrunk(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list sip-trunks",
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
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListSipTrunksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSipTrunks(ctx, req)
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
						Usage: "describe sip-trunks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sip-trunk", Usage: "The ID of the sip trunk.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sipTrunks/%s", cmd.String("project"), cmd.String("location"), cmd.String("sip-trunk"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetSipTrunkRequest{
								Name: name,
							}

							resp, err := client.GetSipTrunk(ctx, req)
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
						Usage: "update sip-trunks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sip-trunk", Usage: "The ID of the sip trunk.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "sip_trunk.name" not yet supported.
							sip_trunk_name := fmt.Sprintf("projects/%s/locations/%s/sipTrunks/%s", cmd.String("project"), cmd.String("location"), cmd.String("sip-trunk"))
							fmt.Printf("Executing update on %s\n", sip_trunk_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "stateless-suggestion",
				Usage: "Manage stateless-suggestion resources",
				Commands: []*cli.Command{

					{
						Name:  "generate",
						Usage: "generate stateless-suggestion",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-settings", Usage: "Name of the CX SecuritySettings which is used to redact generated.", Required: false},
							&cli.StringSliceFlag{Name: "trigger-events", Usage: "A list of trigger events.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GenerateStatelessSuggestionRequest{
								Parent:           parent,
								TriggerEvents:    cmd.StringSlice("trigger-events"),
								SecuritySettings: cmd.String("security-settings"),
							}

							resp, err := client.GenerateStatelessSuggestion(ctx, req)
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
				Name:  "suggestions",
				Usage: "Manage suggestions resources",
				Commands: []*cli.Command{

					{
						Name:  "suggest-conversation-summary",
						Usage: "suggest-conversation-summary suggestions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "context-size", Usage: "Max number of messages prior to and including.", Required: false},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message used as context for.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversation := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							fmt.Printf("Executing suggest-conversation-summary on %s\n", conversation)
							return nil
						},
					},

					{
						Name:  "generate-stateless-summary",
						Usage: "generate-stateless-summary suggestions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message used as context for.", Required: false},
							&cli.IntFlag{Name: "max-context-size", Usage: "Max number of messages prior to and including.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "stateless_conversation.parent" not yet supported.
							stateless_conversation_parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing generate-stateless-summary on %s\n", stateless_conversation_parent)
							return nil
						},
					},

					{
						Name:  "search-knowledge",
						Usage: "search-knowledge suggestions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The conversation (between human agent and end user) where the.", Required: false},
							&cli.StringFlag{Name: "conversation-profile", Usage: "The conversation profile used to configure the search.", Required: true},
							&cli.BoolFlag{Name: "exact-search", Usage: "Whether to search the query exactly without query rewrite.", Required: false},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message when the request is.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-source", Usage: "The source of the query in the request.", Required: false},
							&cli.StringFlag{Name: "session-id", Usage: "The ID of the search session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.SearchKnowledgeRequest{
								Parent:              parent,
								ConversationProfile: cmd.String("conversation-profile"),
								SessionId:           cmd.String("session-id"),
								Conversation:        cmd.String("conversation"),
								LatestMessage:       cmd.String("latest-message"),
								QuerySource:         dialogflowpb.SearchKnowledgeRequest_QuerySource(dialogflowpb.SearchKnowledgeRequest_QuerySource_value[cmd.String("query-source")]),
								ExactSearch:         cmd.Bool("exact-search"),
							}

							resp, err := client.SearchKnowledge(ctx, req)
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
						Name:  "generate",
						Usage: "generate suggestions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message for which the request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "trigger-events", Usage: "A list of trigger events.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							conversation := fmt.Sprintf("projects/%s/conversations/%s", cmd.String("project"), cmd.String("conversation"))
							fmt.Printf("Executing generate on %s\n", conversation)
							return nil
						},
					},

					{
						Name:  "suggest-articles",
						Usage: "suggest-articles suggestions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "context-size", Usage: "Max number of messages prior to and including.", Required: false},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message to compile suggestion.", Required: false},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.SuggestArticlesRequest{
								Parent:        parent,
								LatestMessage: cmd.String("latest-message"),
								ContextSize:   int32(cmd.Int("context-size")),
							}

							resp, err := client.SuggestArticles(ctx, req)
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
						Name:  "suggest-faq-answers",
						Usage: "suggest-faq-answers suggestions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "context-size", Usage: "Max number of messages prior to and including.", Required: false},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message to compile suggestion.", Required: false},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.SuggestFaqAnswersRequest{
								Parent:        parent,
								LatestMessage: cmd.String("latest-message"),
								ContextSize:   int32(cmd.Int("context-size")),
							}

							resp, err := client.SuggestFaqAnswers(ctx, req)
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
						Name:  "suggest-smart-replies",
						Usage: "suggest-smart-replies suggestions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "context-size", Usage: "Max number of messages prior to and including.", Required: false},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message to compile suggestion.", Required: false},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.SuggestSmartRepliesRequest{
								Parent:        parent,
								LatestMessage: cmd.String("latest-message"),
								ContextSize:   int32(cmd.Int("context-size")),
							}

							resp, err := client.SuggestSmartReplies(ctx, req)
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
						Name:  "suggest-knowledge-assist",
						Usage: "suggest-knowledge-assist suggestions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "context-size", Usage: "Max number of messages prior to and including.", Required: false},
							&cli.StringFlag{Name: "conversation", Usage: "The ID of the conversation.", Required: true},
							&cli.StringFlag{Name: "latest-message", Usage: "The name of the latest conversation message to compile.", Required: false},
							&cli.StringFlag{Name: "participant", Usage: "The ID of the participant.", Required: true},
							&cli.StringFlag{Name: "previous-suggested-query", Usage: "The previously suggested query for the given conversation.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/conversations/%s/participants/%s", cmd.String("project"), cmd.String("conversation"), cmd.String("participant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.SuggestKnowledgeAssistRequest{
								Parent:                 parent,
								LatestMessage:          cmd.String("latest-message"),
								ContextSize:            int32(cmd.Int("context-size")),
								PreviousSuggestedQuery: cmd.String("previous-suggested-query"),
							}

							resp, err := client.SuggestKnowledgeAssist(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool-id", Usage: "The ID to use for the tool, which will become the final.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateToolRequest{
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
						Name:  "describe",
						Usage: "describe tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("tool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetToolRequest{
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
						Name:  "list",
						Usage: "list tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of conversation models to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListToolsRequest{
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
						Name:  "delete",
						Usage: "delete tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("tool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTool on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteToolRequest{
								Name: name,
							}

							if err := client.DeleteTool(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update tools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tool", Usage: "The ID of the tool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tool.name" not yet supported.
							tool_name := fmt.Sprintf("projects/%s/locations/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("tool"))
							fmt.Printf("Executing update on %s\n", tool_name)
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
							&cli.StringFlag{Name: "language-code", Usage: "The language for which you want a validation result.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetValidationResultRequest{
								Parent:       parent,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetValidationResult(ctx, req)
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
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return in a single page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.ListVersionsRequest{
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/versions/%s", cmd.String("project"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.GetVersionRequest{
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/agent", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.CreateVersionRequest{
								Parent: parent,
							}

							resp, err := client.CreateVersion(ctx, req)
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "version.name" not yet supported.
							version_name := fmt.Sprintf("projects/%s/agent/versions/%s", cmd.String("project"), cmd.String("version"))
							fmt.Printf("Executing update on %s\n", version_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/agent/versions/%s", cmd.String("project"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteVersion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dialogflow.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dialogflowpb.DeleteVersionRequest{
								Name: name,
							}

							if err := client.DeleteVersion(ctx, req); err != nil {
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
