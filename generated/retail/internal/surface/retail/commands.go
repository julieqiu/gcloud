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

package retail

import (
	retail "cloud.google.com/go/retail/apiv2"
	"cloud.google.com/go/retail/apiv2/retailpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the retail command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "retail",
		Usage: "manage Vertex AI Search for commerce API resources",
		Commands: []*cli.Command{
			{
				Name:  "attributes-config",
				Usage: "Manage attributes-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe attributes-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/attributesConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.GetAttributesConfigRequest{
								Name: name,
							}

							resp, err := client.GetAttributesConfig(ctx, req)
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
						Usage: "update attributes-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "attributes_config.name" not yet supported.
							attributes_config_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/attributesConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", attributes_config_name)
							return nil
						},
					},

					{
						Name:  "add-catalog-attribute",
						Usage: "add-catalog-attribute attributes-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attributes-config", Usage: "Full AttributesConfig resource name.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							attributes_config := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/attributesConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing add-catalog-attribute on %s\n", attributes_config)
							return nil
						},
					},

					{
						Name:  "remove-catalog-attribute",
						Usage: "remove-catalog-attribute attributes-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attributes-config", Usage: "Full AttributesConfig resource name.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "key", Usage: "The attribute name key of the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							attributes_config := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/attributesConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing remove-catalog-attribute on %s\n", attributes_config)
							return nil
						},
					},

					{
						Name:  "replace-catalog-attribute",
						Usage: "replace-catalog-attribute attributes-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attributes-config", Usage: "Full AttributesConfig resource name.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							attributes_config := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/attributesConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing replace-catalog-attribute on %s\n", attributes_config)
							return nil
						},
					},
				},
			},
			{
				Name:  "catalogs",
				Usage: "Manage catalogs resources",
				Commands: []*cli.Command{

					{
						Name:  "export-analytics-metrics",
						Usage: "export-analytics-metrics catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filtering expression to specify restrictions on returned metrics.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing export-analytics-metrics on %s\n", catalog)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [Catalog][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ListCatalogsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCatalogs(ctx, req)
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
						Usage: "update catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "catalog.name" not yet supported.
							catalog_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", catalog_name)
							return nil
						},
					},

					{
						Name:  "set-default-branch",
						Usage: "set-default-branch catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch-id", Usage: "The final component of the resource name of a branch.", Required: false},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, it permits switching to a branch with.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "note", Usage: "Some note on this request, this can be retrieved by.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing set-default-branch on %s\n", catalog)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing describe on %s\n", catalog)
							return nil
						},
					},

					{
						Name:  "complete-query",
						Usage: "complete-query catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "Determines which dataset to use for fetching completion.", Required: false},
							&cli.StringFlag{Name: "device-type", Usage: "The device type context for completion suggestions.", Required: false},
							&cli.BoolFlag{Name: "enable-attribute-suggestions", Usage: "If true, attribute suggestions are enabled and provided in the response.", Required: false},
							&cli.StringFlag{Name: "entity", Usage: "The entity for customers who run multiple entities, domains, sites, or.", Required: false},
							&cli.StringSliceFlag{Name: "language-codes", Usage: "Note that this field applies for `user-data` dataset only.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-suggestions", Usage: "Completion max suggestions.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "The query used to generate suggestions.", Required: true},
							&cli.StringFlag{Name: "visitor-id", Usage: "Recommended field.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing complete-query on %s\n", catalog)
							return nil
						},
					},
				},
			},
			{
				Name:  "completion-config",
				Usage: "Manage completion-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe completion-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/completionConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.GetCompletionConfigRequest{
								Name: name,
							}

							resp, err := client.GetCompletionConfig(ctx, req)
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
						Usage: "update completion-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "completion_config.name" not yet supported.
							completion_config_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/completionConfig", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", completion_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "completion-data",
				Usage: "Manage completion-data resources",
				Commands: []*cli.Command{

					{
						Name:  "import",
						Usage: "import completion-data",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notification-pubsub-topic", Usage: "Pub/Sub topic for receiving notification.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ImportCompletionDataRequest{
								Parent:                  parent,
								NotificationPubsubTopic: cmd.String("notification-pubsub-topic"),
							}

							op, err := client.ImportCompletionData(ctx, req)
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "control-id", Usage: "The ID to use for the Control, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.CreateControlRequest{
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "control", Usage: "The ID of the control.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/controls/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("control"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteControl on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.DeleteControlRequest{
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "control", Usage: "The ID of the control.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "control.name" not yet supported.
							control_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/controls/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("control"))
							fmt.Printf("Executing update on %s\n", control_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe controls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "control", Usage: "The ID of the control.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/controls/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("control"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.GetControlRequest{
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter to apply on the list results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListControls` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ListControlsRequest{
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
				Name:  "generative-question",
				Usage: "Manage generative-question resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update generative-question",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "generative_question_config.catalog" not yet supported.
							generative_question_config_catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", generative_question_config_catalog)
							return nil
						},
					},

					{
						Name:  "batch-update",
						Usage: "batch-update generative-question",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.BatchUpdateGenerativeQuestionConfigsRequest{
								Parent: parent,
							}

							resp, err := client.BatchUpdateGenerativeQuestionConfigs(ctx, req)
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
				Name:  "generative-question-feature",
				Usage: "Manage generative-question-feature resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update generative-question-feature",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "generative_questions_feature_config.catalog" not yet supported.
							generative_questions_feature_config_catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", generative_questions_feature_config_catalog)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe generative-question-feature",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							catalog := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							fmt.Printf("Executing describe on %s\n", catalog)
							return nil
						},
					},
				},
			},
			{
				Name:  "generative-questions",
				Usage: "Manage generative-questions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list generative-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ListGenerativeQuestionConfigsRequest{
								Parent: parent,
							}

							resp, err := client.ListGenerativeQuestionConfigs(ctx, req)
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
				Name:  "models",
				Usage: "Manage models resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.BoolFlag{Name: "dry-run", Usage: "Whether to run a dry run to validate the request (without.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.CreateModelRequest{
								Parent: parent,
								DryRun: cmd.Bool("dry-run"),
							}

							op, err := client.CreateModel(ctx, req)
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
						Usage: "describe models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.GetModelRequest{
								Name: name,
							}

							resp, err := client.GetModel(ctx, req)
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
						Name:  "pause",
						Usage: "pause models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.PauseModelRequest{
								Name: name,
							}

							resp, err := client.PauseModel(ctx, req)
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
						Name:  "resume",
						Usage: "resume models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ResumeModelRequest{
								Name: name,
							}

							resp, err := client.ResumeModel(ctx, req)
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
						Usage: "delete models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteModel on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.DeleteModelRequest{
								Name: name,
							}

							if err := client.DeleteModel(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListModels`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ListModelsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListModels(ctx, req)
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
						Usage: "update models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "model.name" not yet supported.
							model_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("model"))
							fmt.Printf("Executing update on %s\n", model_name)
							return nil
						},
					},

					{
						Name:  "tune",
						Usage: "tune models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.TuneModelRequest{
								Name: name,
							}

							op, err := client.TuneModel(ctx, req)
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
				},
			},
			{
				Name:  "placements",
				Usage: "Manage placements resources",
				Commands: []*cli.Command{

					{
						Name:  "conversational-search",
						Usage: "conversational-search placements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch", Usage: "The branch resource name, such as.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "conversation-id", Usage: "This field specifies the conversation id, which maintains the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "page-categories", Usage: "The categories associated with a category page.", Required: false},
							&cli.StringFlag{Name: "placement", Usage: "The ID of the placement.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Raw search query to be searched for.", Required: false},
							&cli.StringFlag{Name: "visitor-id", Usage: "A unique identifier for tracking visitors.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							placement := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/placements/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("placement"))
							fmt.Printf("Executing conversational-search on %s\n", placement)
							return nil
						},
					},

					{
						Name:  "predict",
						Usage: "predict placements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter for restricting prediction results with a length limit of 5,000.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "This field is not used; leave it unset.", Required: false},
							&cli.StringFlag{Name: "placement", Usage: "The ID of the placement.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Use validate only mode for this prediction query.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							placement := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/placements/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("placement"))
							fmt.Printf("Executing predict on %s\n", placement)
							return nil
						},
					},

					{
						Name:  "search",
						Usage: "search placements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch", Usage: "The branch resource name, such as.", Required: false},
							&cli.StringFlag{Name: "canonical-filter", Usage: "The default filter that is applied when a user performs a search without.", Required: false},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "entity", Usage: "The entity for customers that may run multiple different entities, domains,.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "The filter syntax consists of an expression language for constructing a.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The BCP-47 language code, such as \"en-US\" or \"sr-Latn\".", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "offset", Usage: "A 0-indexed integer that specifies the current offset (that is, starting.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "The order in which products are returned.", Required: false},
							&cli.StringSliceFlag{Name: "page-categories", Usage: "The categories associated with a category page.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [Product][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token.", Required: false},
							&cli.StringFlag{Name: "place-id", Usage: "An id corresponding to a place, such as a store id or region id.", Required: false},
							&cli.StringFlag{Name: "placement", Usage: "The ID of the placement.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Raw search query.", Required: false},
							&cli.StringFlag{Name: "region-code", Usage: "The Unicode country/region code (CLDR) of a location, such as.", Required: false},
							&cli.StringFlag{Name: "search-mode", Usage: "The search mode of the search request.", Required: false},
							&cli.StringSliceFlag{Name: "variant-rollup-keys", Usage: "The keys to fetch and rollup the matching.", Required: false},
							&cli.StringFlag{Name: "visitor-id", Usage: "A unique identifier for tracking visitors.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							placement := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/placements/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("placement"))
							fmt.Printf("Executing search on %s\n", placement)
							return nil
						},
					},
				},
			},
			{
				Name:  "products",
				Usage: "Manage products resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "product-id", Usage: "The ID to use for the [Product][google.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.CreateProductRequest{
								Parent:    parent,
								ProductId: cmd.String("product-id"),
							}

							resp, err := client.CreateProduct(ctx, req)
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
						Usage: "describe products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.GetProductRequest{
								Name: name,
							}

							resp, err := client.GetProduct(ctx, req)
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
						Usage: "list products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter to apply on the list results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of [Product][google.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ListProductsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListProducts(ctx, req)
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
						Usage: "update products",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Product][google.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "product.name" not yet supported.
							product_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							fmt.Printf("Executing update on %s\n", product_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteProduct on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.DeleteProductRequest{
								Name: name,
							}

							if err := client.DeleteProduct(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "purge",
						Usage: "purge products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter string to specify the products to be deleted with a.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Actually perform the purge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.PurgeProductsRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
								Force:  cmd.Bool("force"),
							}

							op, err := client.PurgeProducts(ctx, req)
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
						Usage: "import products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notification-pubsub-topic", Usage: "Full Pub/Sub topic name for receiving notification.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reconciliation-mode", Usage: "The mode of reconciliation between existing products and the products to be.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "Deprecated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ImportProductsRequest{
								Parent:                  parent,
								RequestId:               cmd.String("request-id"),
								ReconciliationMode:      retailpb.ImportProductsRequest_ReconciliationMode(retailpb.ImportProductsRequest_ReconciliationMode_value[cmd.String("reconciliation-mode")]),
								NotificationPubsubTopic: cmd.String("notification-pubsub-topic"),
							}

							op, err := client.ImportProducts(ctx, req)
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
						Name:  "set-inventory",
						Usage: "set-inventory products",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Product][google.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "inventory.name" not yet supported.
							inventory_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							fmt.Printf("Executing set-inventory on %s\n", inventory_name)
							return nil
						},
					},

					{
						Name:  "add-fulfillment-places",
						Usage: "add-fulfillment-places products",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Product][google.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "place-ids", Usage: "The IDs for this.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The fulfillment type, including commonly used types (such as.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							product := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							fmt.Printf("Executing add-fulfillment-places on %s\n", product)
							return nil
						},
					},

					{
						Name:  "remove-fulfillment-places",
						Usage: "remove-fulfillment-places products",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Product][google.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "place-ids", Usage: "The IDs for this.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The fulfillment type, including commonly used types (such as.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							product := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							fmt.Printf("Executing remove-fulfillment-places on %s\n", product)
							return nil
						},
					},

					{
						Name:  "add-local-inventories",
						Usage: "add-local-inventories products",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Product][google.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							product := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							fmt.Printf("Executing add-local-inventories on %s\n", product)
							return nil
						},
					},

					{
						Name:  "remove-local-inventories",
						Usage: "remove-local-inventories products",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Product][google.", Required: false},
							&cli.StringFlag{Name: "branche", Usage: "The ID of the branche.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "place-ids", Usage: "A list of place IDs to have their inventory deleted.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The ID of the product.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							product := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/branches/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("branche"), cmd.String("product"))
							fmt.Printf("Executing remove-local-inventories on %s\n", product)
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
						Name:  "create",
						Usage: "create serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config-id", Usage: "The ID to use for the ServingConfig, which will become the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.CreateServingConfigRequest{
								Parent:          parent,
								ServingConfigId: cmd.String("serving-config-id"),
							}

							resp, err := client.CreateServingConfig(ctx, req)
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
						Usage: "delete serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("serving-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteServingConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.DeleteServingConfigRequest{
								Name: name,
							}

							if err := client.DeleteServingConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "serving_config.name" not yet supported.
							serving_config_name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("serving-config"))
							fmt.Printf("Executing update on %s\n", serving_config_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("serving-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.GetServingConfigRequest{
								Name: name,
							}

							resp, err := client.GetServingConfig(ctx, req)
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
						Usage: "list serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListServingConfigs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ListServingConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListServingConfigs(ctx, req)
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
						Name:  "add-control",
						Usage: "add-control serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "control-id", Usage: "The id of the control to apply.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("serving-config"))
							fmt.Printf("Executing add-control on %s\n", serving_config)
							return nil
						},
					},

					{
						Name:  "remove-control",
						Usage: "remove-control serving-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "control-id", Usage: "The id of the control to apply.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serving-config", Usage: "The ID of the serving config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							serving_config := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/servingConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"), cmd.String("serving-config"))
							fmt.Printf("Executing remove-control on %s\n", serving_config)
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "write-async", Usage: "If set to true, the user event will be written asynchronously after.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.WriteUserEventRequest{
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.IntFlag{Name: "ets", Usage: "The event timestamp in milliseconds.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "raw-json", Usage: "An arbitrary serialized JSON string that contains necessary information.", Required: false},
							&cli.StringFlag{Name: "uri", Usage: "The URL including cgi-parameters but excluding the hash fragment with a.", Required: false},
							&cli.StringFlag{Name: "user-event", Usage: "URL encoded UserEvent proto with a length limit of 2,000,000.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.CollectUserEventRequest{
								Parent:    parent,
								UserEvent: cmd.String("user-event"),
								Uri:       cmd.String("uri"),
								Ets:       cmd.Int("ets"),
								RawJson:   cmd.String("raw-json"),
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter string to specify the events to be deleted with a.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Actually perform the purge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.PurgeUserEventsRequest{
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.ImportUserEventsRequest{
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

					{
						Name:  "rejoin",
						Usage: "rejoin user-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-event-rejoin-scope", Usage: "The type of the user event rejoin to define the scope and range of the user.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/catalogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := retail.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &retailpb.RejoinUserEventsRequest{
								Parent:               parent,
								UserEventRejoinScope: retailpb.RejoinUserEventsRequest_UserEventRejoinScope(retailpb.RejoinUserEventsRequest_UserEventRejoinScope_value[cmd.String("user-event-rejoin-scope")]),
							}

							op, err := client.RejoinUserEvents(ctx, req)
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
