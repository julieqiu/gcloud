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

package documentai

import (
	documentai "cloud.google.com/go/documentai/apiv1"
	"cloud.google.com/go/documentai/apiv1/documentaipb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the documentai command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "documentai",
		Usage: "manage Cloud Document AI API resources",
		Commands: []*cli.Command{
			{
				Name:  "evaluations",
				Usage: "Manage evaluations resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s/evaluations/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"), cmd.String("evaluation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.GetEvaluationRequest{
								Name: name,
							}

							resp, err := client.GetEvaluation(ctx, req)
							if err != nil {
								return err
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListEvaluations` call.", Required: false},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.ListEvaluationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEvaluations(ctx, req)
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
				Name:  "human-review-config",
				Usage: "Manage human-review-config resources",
				Commands: []*cli.Command{

					{
						Name:  "review-document",
						Usage: "review-document human-review-config",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "enable-schema-validation", Usage: "Whether the validation should be performed on the ad-hoc review request.", Required: false},
							&cli.StringFlag{Name: "human-review-config", Usage: "The resource name of the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "priority", Usage: "The priority of the human review task.", Required: false},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							human_review_config := fmt.Sprintf("projects/%s/locations/%s/processors/%s/humanReviewConfig", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							fmt.Printf("Executing review-document on %s\n", human_review_config)
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
						Name:  "fetch-processor-types",
						Usage: "fetch-processor-types locations",
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
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.FetchProcessorTypesRequest{
								Parent: parent,
							}

							resp, err := client.FetchProcessorTypes(ctx, req)
							if err != nil {
								return err
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
							name := fmt.Sprintf("projects/%s/locations/%s/operations", cmd.String("project"), cmd.String("location"))
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
				Name:  "processor-types",
				Usage: "Manage processor-types resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list processor-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of processor types to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Used to retrieve the next page of results, empty if at the end of the list.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.ListProcessorTypesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListProcessorTypes(ctx, req)
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
						Usage: "describe processor-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor-type", Usage: "The ID of the processor type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processorTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.GetProcessorTypeRequest{
								Name: name,
							}

							resp, err := client.GetProcessorType(ctx, req)
							if err != nil {
								return err
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
				Name:  "processor-versions",
				Usage: "Manage processor-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "train",
						Usage: "train processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "base-processor-version", Usage: "The processor version to use as a base for training.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.TrainProcessorVersionRequest{
								Parent:               parent,
								BaseProcessorVersion: cmd.String("base-processor-version"),
							}

							op, err := client.TrainProcessorVersion(ctx, req)
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
						Usage: "describe processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.GetProcessorVersionRequest{
								Name: name,
							}

							resp, err := client.GetProcessorVersion(ctx, req)
							if err != nil {
								return err
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
						Usage: "list processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of processor versions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "We will return the processor versions sorted by creation time.", Required: false},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.ListProcessorVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListProcessorVersions(ctx, req)
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
						Usage: "delete processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteProcessorVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.DeleteProcessorVersionRequest{
								Name: name,
							}

							op, err := client.DeleteProcessorVersion(ctx, req)
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
						Usage: "deploy processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.DeployProcessorVersionRequest{
								Name: name,
							}

							op, err := client.DeployProcessorVersion(ctx, req)
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
						Usage: "undeploy processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.UndeployProcessorVersionRequest{
								Name: name,
							}

							op, err := client.UndeployProcessorVersion(ctx, req)
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
						Name:  "evaluate-processor-version",
						Usage: "evaluate-processor-version processor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "processor-version", Usage: "The ID of the processor version.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							processor_version := fmt.Sprintf("projects/%s/locations/%s/processors/%s/processorVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"), cmd.String("processor-version"))
							fmt.Printf("Executing evaluate-processor-version on %s\n", processor_version)
							return nil
						},
					},
				},
			},
			{
				Name:  "processors",
				Usage: "Manage processors resources",
				Commands: []*cli.Command{

					{
						Name:  "process",
						Usage: "process processors",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "imageless-mode", Usage: "Option to remove images from the document.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-human-review", Usage: "Whether human review should be skipped for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.ProcessRequest{
								Name:            name,
								SkipHumanReview: cmd.Bool("skip-human-review"),
								ImagelessMode:   cmd.Bool("imageless-mode"),
							}

							resp, err := client.ProcessDocument(ctx, req)
							if err != nil {
								return err
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
						Name:  "batch-process",
						Usage: "batch-process processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-human-review", Usage: "Whether human review should be skipped for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.BatchProcessRequest{
								Name:            name,
								SkipHumanReview: cmd.Bool("skip-human-review"),
							}

							op, err := client.BatchProcessDocuments(ctx, req)
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
						Usage: "list processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of processors to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "We will return the processors sorted by creation time.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.ListProcessorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListProcessors(ctx, req)
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
						Usage: "describe processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.GetProcessorRequest{
								Name: name,
							}

							resp, err := client.GetProcessor(ctx, req)
							if err != nil {
								return err
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
						Usage: "create processors",
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
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.CreateProcessorRequest{
								Parent: parent,
							}

							resp, err := client.CreateProcessor(ctx, req)
							if err != nil {
								return err
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
						Usage: "delete processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteProcessor %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.DeleteProcessorRequest{
								Name: name,
							}

							op, err := client.DeleteProcessor(ctx, req)
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
						Name:  "enable",
						Usage: "enable processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.EnableProcessorRequest{
								Name: name,
							}

							op, err := client.EnableProcessor(ctx, req)
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
						Name:  "disable",
						Usage: "disable processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := documentai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &documentaipb.DisableProcessorRequest{
								Name: name,
							}

							op, err := client.DisableProcessor(ctx, req)
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
						Name:  "set-default-processor-version",
						Usage: "set-default-processor-version processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "default-processor-version", Usage: "The resource name of child.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							processor := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							fmt.Printf("Executing set-default-processor-version on %s\n", processor)
							return nil
						},
					},
				},
			},
		},
	}
}
