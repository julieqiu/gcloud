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

package visionai

import (
	visionai "cloud.google.com/go/visionai/apiv1"
	"cloud.google.com/go/visionai/apiv1/visionaipb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the visionai command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "visionai",
		Usage: "manage Vision AI API resources",
		Commands: []*cli.Command{
			{
				Name:  "analyses",
				Usage: "Manage analyses resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListAnalysesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
						Name:  "describe",
						Usage: "describe analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analyse", Usage: "The ID of the analyse.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("analyse"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetAnalysisRequest{
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
						Name:  "create",
						Usage: "create analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analysis-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateAnalysisRequest{
								Parent:     parent,
								AnalysisId: cmd.String("analysis-id"),
								RequestId:  cmd.String("request-id"),
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
						Name:  "update",
						Usage: "update analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analyse", Usage: "The ID of the analyse.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "analysis.name" not yet supported.
							analysis_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("analyse"))
							fmt.Printf("Executing update on %s\n", analysis_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "analyse", Usage: "The ID of the analyse.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("analyse"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAnalysis %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteAnalysisRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteAnalysis(ctx, req)
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
				Name:  "annotations",
				Usage: "Manage annotations resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create annotations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "annotation-id", Usage: "The ID to use for the annotation, which will become the final.", Required: false},
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateAnnotationRequest{
								Parent:       parent,
								AnnotationId: runtime.Ptr(cmd.String("annotation-id")),
							}

							resp, err := client.CreateAnnotation(ctx, req)
							if err != nil {
								return err
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
						Usage: "describe annotations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "annotation", Usage: "The ID of the annotation.", Required: true},
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s/annotations/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"), cmd.String("annotation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetAnnotationRequest{
								Name: name,
							}

							resp, err := client.GetAnnotation(ctx, req)
							if err != nil {
								return err
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
						Usage: "list annotations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter applied to the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of annotations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListAnnotations` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListAnnotationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListAnnotations(ctx, req)
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
						Usage: "update annotations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "annotation", Usage: "The ID of the annotation.", Required: true},
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "annotation.name" not yet supported.
							annotation_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s/annotations/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"), cmd.String("annotation"))
							fmt.Printf("Executing update on %s\n", annotation_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete annotations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "annotation", Usage: "The ID of the annotation.", Required: true},
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s/annotations/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"), cmd.String("annotation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAnnotation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteAnnotationRequest{
								Name: name,
							}

							if err := client.DeleteAnnotation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "applications",
				Usage: "Manage applications resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListApplicationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListApplications(ctx, req)
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
						Usage: "describe applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetApplicationRequest{
								Name: name,
							}

							resp, err := client.GetApplication(ctx, req)
							if err != nil {
								return err
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
						Usage: "create applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateApplicationRequest{
								Parent:        parent,
								ApplicationId: cmd.String("application-id"),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.CreateApplication(ctx, req)
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
						Usage: "update applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "application.name" not yet supported.
							application_name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							fmt.Printf("Executing update on %s\n", application_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any instances and drafts from this application.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteApplication %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteApplicationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteApplication(ctx, req)
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
						Usage: "deploy applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.BoolFlag{Name: "enable-monitoring", Usage: "Whether or not to enable monitoring for the application on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the application graph, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeployApplicationRequest{
								Name:             name,
								ValidateOnly:     cmd.Bool("validate-only"),
								RequestId:        cmd.String("request-id"),
								EnableMonitoring: cmd.Bool("enable-monitoring"),
							}

							op, err := client.DeployApplication(ctx, req)
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
						Usage: "undeploy applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.UndeployApplicationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.UndeployApplication(ctx, req)
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
						Name:  "add-stream-input",
						Usage: "add-stream-input applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.AddApplicationStreamInputRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.AddApplicationStreamInput(ctx, req)
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
						Name:  "remove-stream-input",
						Usage: "remove-stream-input applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.RemoveApplicationStreamInputRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.RemoveApplicationStreamInput(ctx, req)
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
						Name:  "update-stream-input",
						Usage: "update-stream-input applications",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If true, UpdateApplicationStreamInput will insert stream input to.", Required: false},
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.UpdateApplicationStreamInputRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
							}

							op, err := client.UpdateApplicationStreamInput(ctx, req)
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
						Usage: "create applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateApplicationInstancesRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateApplicationInstances(ctx, req)
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
						Usage: "delete applications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringSliceFlag{Name: "instance-ids", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteApplicationInstancesRequest{
								Name:        name,
								InstanceIds: cmd.StringSlice("instance-ids"),
								RequestId:   cmd.String("request-id"),
							}

							op, err := client.DeleteApplicationInstances(ctx, req)
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
						Name:  "update-application-instances",
						Usage: "update-application-instances applications",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If true, Update Request will create one resource if the target resource.", Required: false},
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.UpdateApplicationInstancesRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
							}

							op, err := client.UpdateApplicationInstances(ctx, req)
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
				Name:  "assets",
				Usage: "Manage assets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset-id", Usage: "The ID to use for the asset, which will become the final.", Required: false},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateAssetRequest{
								Parent:  parent,
								AssetId: runtime.Ptr(cmd.String("asset-id")),
							}

							resp, err := client.CreateAsset(ctx, req)
							if err != nil {
								return err
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
						Usage: "update assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "asset.name" not yet supported.
							asset_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							fmt.Printf("Executing update on %s\n", asset_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetAssetRequest{
								Name: name,
							}

							resp, err := client.GetAsset(ctx, req)
							if err != nil {
								return err
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
						Usage: "list assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter applied to the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of assets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListAssets` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListAssetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListAssets(ctx, req)
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
						Usage: "delete assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAsset %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteAssetRequest{
								Name: name,
							}

							op, err := client.DeleteAsset(ctx, req)
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
						Name:  "upload",
						Usage: "upload assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.UploadAssetRequest{
								Name: name,
							}

							op, err := client.UploadAsset(ctx, req)
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
						Name:  "generate-retrieval-url",
						Usage: "generate-retrieval-url assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GenerateRetrievalUrlRequest{
								Name: name,
							}

							resp, err := client.GenerateRetrievalUrl(ctx, req)
							if err != nil {
								return err
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
						Name:  "analyze",
						Usage: "analyze assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.AnalyzeAssetRequest{
								Name: name,
							}

							op, err := client.AnalyzeAsset(ctx, req)
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
						Name:  "index",
						Usage: "index assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "index", Usage: "The name of the index.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.IndexAssetRequest{
								Name:  name,
								Index: cmd.String("index"),
							}

							op, err := client.IndexAsset(ctx, req)
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
						Name:  "remove-index",
						Usage: "remove-index assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "index", Usage: "The name of the index.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.RemoveIndexAssetRequest{
								Name:  name,
								Index: cmd.String("index"),
							}

							op, err := client.RemoveIndexAsset(ctx, req)
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
						Name:  "clip",
						Usage: "clip assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ClipAssetRequest{
								Name: name,
							}

							resp, err := client.ClipAsset(ctx, req)
							if err != nil {
								return err
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
						Name:  "generate-hls-uri",
						Usage: "generate-hls-uri assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.BoolFlag{Name: "live-view-enabled", Usage: "Option to exclusively show a livestream of the asset with up to 3 minutes.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GenerateHlsUriRequest{
								Name:            name,
								LiveViewEnabled: cmd.Bool("live-view-enabled"),
							}

							resp, err := client.GenerateHlsUri(ctx, req)
							if err != nil {
								return err
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
						Usage: "import assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ImportAssetsRequest{
								Parent: parent,
							}

							op, err := client.ImportAssets(ctx, req)
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
				Name:  "channels",
				Usage: "Manage channels resources",
				Commands: []*cli.Command{

					{
						Name:  "materialize-channel",
						Usage: "materialize-channel channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel-id", Usage: "Id of the channel.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.MaterializeChannelRequest{
								Parent:    parent,
								ChannelId: cmd.String("channel-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.MaterializeChannel(ctx, req)
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
				Name:  "clusters",
				Usage: "Manage clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "health-check",
						Usage: "health-check clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							cluster := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							fmt.Printf("Executing health-check on %s\n", cluster)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListClusters(ctx, req)
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
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetClusterRequest{
								Name: name,
							}

							resp, err := client.GetCluster(ctx, req)
							if err != nil {
								return err
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
						Usage: "create clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateClusterRequest{
								Parent:    parent,
								ClusterId: cmd.String("cluster-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateCluster(ctx, req)
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
						Usage: "update clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "cluster.name" not yet supported.
							cluster_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							fmt.Printf("Executing update on %s\n", cluster_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteCluster(ctx, req)
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
				Name:  "collections",
				Usage: "Manage collections resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection-id", Usage: "The ID to use for the collection, which will become the final.", Required: false},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateCollectionRequest{
								Parent:       parent,
								CollectionId: runtime.Ptr(cmd.String("collection-id")),
							}

							op, err := client.CreateCollection(ctx, req)
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
						Usage: "delete collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("collection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCollection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteCollectionRequest{
								Name: name,
							}

							op, err := client.DeleteCollection(ctx, req)
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
						Name:  "describe",
						Usage: "describe collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("collection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetCollectionRequest{
								Name: name,
							}

							resp, err := client.GetCollection(ctx, req)
							if err != nil {
								return err
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
						Usage: "update collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "collection.name" not yet supported.
							collection_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("collection"))
							fmt.Printf("Executing update on %s\n", collection_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of collections to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListCollectionsRequest` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListCollectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCollections(ctx, req)
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
						Name:  "add-collection-item",
						Usage: "add-collection-item collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "item.collection" not yet supported.
							item_collection := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("collection"))
							fmt.Printf("Executing add-collection-item on %s\n", item_collection)
							return nil
						},
					},

					{
						Name:  "remove-collection-item",
						Usage: "remove-collection-item collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "item.collection" not yet supported.
							item_collection := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("collection"))
							fmt.Printf("Executing remove-collection-item on %s\n", item_collection)
							return nil
						},
					},

					{
						Name:  "view-collection-items",
						Usage: "view-collection-items collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection", Usage: "The ID of the collection.", Required: true},
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of collections to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ViewCollectionItemsRequest` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							collection := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("collection"))
							fmt.Printf("Executing view-collection-items on %s\n", collection)
							return nil
						},
					},
				},
			},
			{
				Name:  "corpora",
				Usage: "Manage corpora resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create corpora",
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
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateCorpusRequest{
								Parent: parent,
							}

							op, err := client.CreateCorpus(ctx, req)
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
						Usage: "describe corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetCorpusRequest{
								Name: name,
							}

							resp, err := client.GetCorpus(ctx, req)
							if err != nil {
								return err
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
						Usage: "update corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "corpus.name" not yet supported.
							corpus_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							fmt.Printf("Executing update on %s\n", corpus_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter applied to the returned corpora list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results for the server to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListCorporaRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListCorpora(ctx, req)
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
						Usage: "delete corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCorpus on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteCorpusRequest{
								Name: name,
							}

							if err := client.DeleteCorpus(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "analyze",
						Usage: "analyze corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.AnalyzeCorpusRequest{
								Name: name,
							}

							op, err := client.AnalyzeCorpus(ctx, req)
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
						Name:  "search-assets",
						Usage: "search-assets corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "corpus", Usage: "The parent corpus to search.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The number of results to be returned in this page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The continuation token to fetch the next page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "result-annotation-keys", Usage: "A list of annotation keys to specify the annotations to be retrieved and.", Required: false},
							&cli.StringFlag{Name: "search-query", Usage: "Global search query.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							corpus := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							fmt.Printf("Executing search-assets on %s\n", corpus)
							return nil
						},
					},
				},
			},
			{
				Name:  "data-schemas",
				Usage: "Manage data-schemas resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateDataSchemaRequest{
								Parent: parent,
							}

							resp, err := client.CreateDataSchema(ctx, req)
							if err != nil {
								return err
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
						Usage: "update data-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "data-schema", Usage: "The ID of the data schema.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_schema.name" not yet supported.
							data_schema_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/dataSchemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("data-schema"))
							fmt.Printf("Executing update on %s\n", data_schema_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe data-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "data-schema", Usage: "The ID of the data schema.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/dataSchemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("data-schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetDataSchemaRequest{
								Name: name,
							}

							resp, err := client.GetDataSchema(ctx, req)
							if err != nil {
								return err
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
						Usage: "delete data-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "data-schema", Usage: "The ID of the data schema.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/dataSchemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("data-schema"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDataSchema on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteDataSchemaRequest{
								Name: name,
							}

							if err := client.DeleteDataSchema(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list data-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of data schemas to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDataSchemas` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListDataSchemasRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataSchemas(ctx, req)
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
				Name:  "drafts",
				Usage: "Manage drafts resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list drafts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListDraftsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDrafts(ctx, req)
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
						Usage: "describe drafts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "draft", Usage: "The ID of the draft.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/drafts/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("draft"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetDraftRequest{
								Name: name,
							}

							resp, err := client.GetDraft(ctx, req)
							if err != nil {
								return err
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
						Usage: "create drafts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "draft-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateDraftRequest{
								Parent:    parent,
								DraftId:   cmd.String("draft-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateDraft(ctx, req)
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
						Usage: "update drafts",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If true, UpdateDraftRequest will create one resource if the target resource.", Required: false},
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "draft", Usage: "The ID of the draft.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "draft.name" not yet supported.
							draft_name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/drafts/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("draft"))
							fmt.Printf("Executing update on %s\n", draft_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete drafts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "draft", Usage: "The ID of the draft.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/drafts/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("draft"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDraft %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteDraftRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteDraft(ctx, req)
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
				Name:  "events",
				Usage: "Manage events resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListEventsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEvents(ctx, req)
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
						Usage: "describe events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "event", Usage: "The ID of the event.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/events/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("event"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetEventRequest{
								Name: name,
							}

							resp, err := client.GetEvent(ctx, req)
							if err != nil {
								return err
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
						Usage: "create events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "event-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateEventRequest{
								Parent:    parent,
								EventId:   cmd.String("event-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateEvent(ctx, req)
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
						Usage: "update events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "event", Usage: "The ID of the event.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "event.name" not yet supported.
							event_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/events/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("event"))
							fmt.Printf("Executing update on %s\n", event_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "event", Usage: "The ID of the event.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/events/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("event"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEvent %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteEventRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteEvent(ctx, req)
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
				Name:  "index-endpoints",
				Usage: "Manage index-endpoints resources",
				Commands: []*cli.Command{

					{
						Name:  "search-index-endpoint",
						Usage: "search-index-endpoint index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The continuation token to fetch the next page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index_endpoint := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing search-index-endpoint on %s\n", index_endpoint)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint-id", Usage: "The ID to use for the IndexEndpoint, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateIndexEndpointRequest{
								Parent:          parent,
								IndexEndpointId: cmd.String("index-endpoint-id"),
							}

							op, err := client.CreateIndexEndpoint(ctx, req)
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
						Usage: "describe index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetIndexEndpointRequest{
								Name: name,
							}

							resp, err := client.GetIndexEndpoint(ctx, req)
							if err != nil {
								return err
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
						Usage: "list index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter applied to the returned list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListIndexEndpointsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListIndexEndpoints(ctx, req)
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
						Usage: "update index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "index_endpoint.name" not yet supported.
							index_endpoint_name := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing update on %s\n", index_endpoint_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIndexEndpoint %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteIndexEndpointRequest{
								Name: name,
							}

							op, err := client.DeleteIndexEndpoint(ctx, req)
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
						Name:  "deploy-index",
						Usage: "deploy-index index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index_endpoint := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing deploy-index on %s\n", index_endpoint)
							return nil
						},
					},

					{
						Name:  "undeploy-index",
						Usage: "undeploy-index index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index_endpoint := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing undeploy-index on %s\n", index_endpoint)
							return nil
						},
					},
				},
			},
			{
				Name:  "indexes",
				Usage: "Manage indexes resources",
				Commands: []*cli.Command{

					{
						Name:  "view-assets",
						Usage: "view-assets indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter applied to the returned list.", Required: false},
							&cli.StringFlag{Name: "index", Usage: "The index that owns this collection of assets.", Required: true},
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of assets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ViewIndexedAssets` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("indexe"))
							fmt.Printf("Executing view-assets on %s\n", index)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "index-id", Usage: "The ID for the index.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateIndexRequest{
								Parent:  parent,
								IndexId: cmd.String("index-id"),
							}

							op, err := client.CreateIndex(ctx, req)
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
						Usage: "update indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "index.name" not yet supported.
							index_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("indexe"))
							fmt.Printf("Executing update on %s\n", index_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("indexe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetIndexRequest{
								Name: name,
							}

							resp, err := client.GetIndex(ctx, req)
							if err != nil {
								return err
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
						Usage: "list indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of indexes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListIndexes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListIndexesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListIndexes(ctx, req)
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
						Usage: "delete indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("indexe"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIndex %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteIndexRequest{
								Name: name,
							}

							op, err := client.DeleteIndex(ctx, req)
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
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListInstancesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInstances(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "application", Usage: "The ID of the application.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetInstanceRequest{
								Name: name,
							}

							resp, err := client.GetInstance(ctx, req)
							if err != nil {
								return err
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
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListPublicOperatorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPublicOperators(ctx, req)
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
						Name:  "resolve-operator-info",
						Usage: "resolve-operator-info locations",
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
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ResolveOperatorInfoRequest{
								Parent: parent,
							}

							resp, err := client.ResolveOperatorInfo(ctx, req)
							if err != nil {
								return err
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
				Name:  "operators",
				Usage: "Manage operators resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListOperatorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListOperators(ctx, req)
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
						Usage: "describe operators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operator", Usage: "The ID of the operator.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operators/%s", cmd.String("project"), cmd.String("location"), cmd.String("operator"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetOperatorRequest{
								Name: name,
							}

							resp, err := client.GetOperator(ctx, req)
							if err != nil {
								return err
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
						Usage: "create operators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operator-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateOperatorRequest{
								Parent:     parent,
								OperatorId: cmd.String("operator-id"),
								RequestId:  cmd.String("request-id"),
							}

							op, err := client.CreateOperator(ctx, req)
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
						Usage: "update operators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operator", Usage: "The ID of the operator.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "operator.name" not yet supported.
							operator_name := fmt.Sprintf("projects/%s/locations/%s/operators/%s", cmd.String("project"), cmd.String("location"), cmd.String("operator"))
							fmt.Printf("Executing update on %s\n", operator_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operators",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operator", Usage: "The ID of the operator.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operators/%s", cmd.String("project"), cmd.String("location"), cmd.String("operator"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteOperator %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteOperatorRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteOperator(ctx, req)
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
				Name:  "processes",
				Usage: "Manage processes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list processes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListProcessesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListProcesses(ctx, req)
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
						Usage: "describe processes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processe", Usage: "The ID of the processe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/processes/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("processe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetProcessRequest{
								Name: name,
							}

							resp, err := client.GetProcess(ctx, req)
							if err != nil {
								return err
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
						Usage: "create processes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "process-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateProcessRequest{
								Parent:    parent,
								ProcessId: cmd.String("process-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateProcess(ctx, req)
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
						Usage: "update processes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processe", Usage: "The ID of the processe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "process.name" not yet supported.
							process_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/processes/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("processe"))
							fmt.Printf("Executing update on %s\n", process_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete processes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processe", Usage: "The ID of the processe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/processes/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("processe"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteProcess %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteProcessRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteProcess(ctx, req)
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
						Name:  "batch-run",
						Usage: "batch-run processes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "batch-id", Usage: "The batch ID.", Required: false},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.BatchRunProcessRequest{
								Parent:  parent,
								BatchId: cmd.String("batch-id"),
							}

							op, err := client.BatchRunProcess(ctx, req)
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
				Name:  "processors",
				Usage: "Manage processors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListProcessorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
						Name:  "prebuilt",
						Usage: "prebuilt processors",
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
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListPrebuiltProcessorsRequest{
								Parent: parent,
							}

							resp, err := client.ListPrebuiltProcessors(ctx, req)
							if err != nil {
								return err
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
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetProcessorRequest{
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
							&cli.StringFlag{Name: "processor-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateProcessorRequest{
								Parent:      parent,
								ProcessorId: cmd.String("processor-id"),
								RequestId:   cmd.String("request-id"),
							}

							op, err := client.CreateProcessor(ctx, req)
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
						Usage: "update processors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processor", Usage: "The ID of the processor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "processor.name" not yet supported.
							processor_name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
							fmt.Printf("Executing update on %s\n", processor_name)
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
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
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
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteProcessorRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
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
				},
			},
			{
				Name:  "search-configs",
				Usage: "Manage search-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create search-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-config-id", Usage: "ID to use for the new search config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateSearchConfigRequest{
								Parent:         parent,
								SearchConfigId: cmd.String("search-config-id"),
							}

							resp, err := client.CreateSearchConfig(ctx, req)
							if err != nil {
								return err
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
						Usage: "update search-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-config", Usage: "The ID of the search config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "search_config.name" not yet supported.
							search_config_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("search-config"))
							fmt.Printf("Executing update on %s\n", search_config_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe search-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-config", Usage: "The ID of the search config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("search-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetSearchConfigRequest{
								Name: name,
							}

							resp, err := client.GetSearchConfig(ctx, req)
							if err != nil {
								return err
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
						Usage: "delete search-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-config", Usage: "The ID of the search config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("search-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSearchConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteSearchConfigRequest{
								Name: name,
							}

							if err := client.DeleteSearchConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list search-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of search configurations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSearchConfigs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListSearchConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSearchConfigs(ctx, req)
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
				Name:  "search-hypernyms",
				Usage: "Manage search-hypernyms resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create search-hypernyms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-hypernym-id", Usage: "The search hypernym id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateSearchHypernymRequest{
								Parent:           parent,
								SearchHypernymId: runtime.Ptr(cmd.String("search-hypernym-id")),
							}

							resp, err := client.CreateSearchHypernym(ctx, req)
							if err != nil {
								return err
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
						Usage: "update search-hypernyms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-hypernym", Usage: "The ID of the search hypernym.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "search_hypernym.name" not yet supported.
							search_hypernym_name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchHypernyms/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("search-hypernym"))
							fmt.Printf("Executing update on %s\n", search_hypernym_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe search-hypernyms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-hypernym", Usage: "The ID of the search hypernym.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchHypernyms/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("search-hypernym"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetSearchHypernymRequest{
								Name: name,
							}

							resp, err := client.GetSearchHypernym(ctx, req)
							if err != nil {
								return err
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
						Usage: "delete search-hypernyms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-hypernym", Usage: "The ID of the search hypernym.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchHypernyms/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"), cmd.String("search-hypernym"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSearchHypernym on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteSearchHypernymRequest{
								Name: name,
							}

							if err := client.DeleteSearchHypernym(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list search-hypernyms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "corpora", Usage: "The ID of the corpora.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of SearchHypernyms returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `SearchHypernym` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListSearchHypernymsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSearchHypernyms(ctx, req)
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
				Name:  "series",
				Usage: "Manage series resources",
				Commands: []*cli.Command{

					{
						Name:  "acquire-lease",
						Usage: "acquire-lease series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "lease-type", Usage: "The lease type.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "owner", Usage: "The owner name.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serie", Usage: "The ID of the serie.", Required: true},
							&cli.StringFlag{Name: "series", Usage: "The series name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							series := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("serie"))
							fmt.Printf("Executing acquire-lease on %s\n", series)
							return nil
						},
					},

					{
						Name:  "renew-lease",
						Usage: "renew-lease series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "id", Usage: "Lease id.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "owner", Usage: "Lease owner.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serie", Usage: "The ID of the serie.", Required: true},
							&cli.StringFlag{Name: "series", Usage: "Series name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							series := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("serie"))
							fmt.Printf("Executing renew-lease on %s\n", series)
							return nil
						},
					},

					{
						Name:  "release-lease",
						Usage: "release-lease series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "id", Usage: "Lease id.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "owner", Usage: "Lease owner.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serie", Usage: "The ID of the serie.", Required: true},
							&cli.StringFlag{Name: "series", Usage: "Series name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							series := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("serie"))
							fmt.Printf("Executing release-lease on %s\n", series)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListSeriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSeries(ctx, req)
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
						Usage: "describe series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serie", Usage: "The ID of the serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("serie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetSeriesRequest{
								Name: name,
							}

							resp, err := client.GetSeries(ctx, req)
							if err != nil {
								return err
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
						Usage: "create series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "series-id", Usage: "Id of the requesting object.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateSeriesRequest{
								Parent:    parent,
								SeriesId:  cmd.String("series-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateSeries(ctx, req)
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
						Usage: "update series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "serie", Usage: "The ID of the serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "series.name" not yet supported.
							series_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("serie"))
							fmt.Printf("Executing update on %s\n", series_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "serie", Usage: "The ID of the serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("serie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSeries %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteSeriesRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteSeries(ctx, req)
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
				Name:  "streams",
				Usage: "Manage streams resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.ListStreamsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListStreams(ctx, req)
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
						Usage: "describe streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.GetStreamRequest{
								Name: name,
							}

							resp, err := client.GetStream(ctx, req)
							if err != nil {
								return err
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
						Usage: "create streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "stream-id", Usage: "Id of the requesting object.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.CreateStreamRequest{
								Parent:    parent,
								StreamId:  cmd.String("stream-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateStream(ctx, req)
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
						Usage: "update streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "stream.name" not yet supported.
							stream_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
							fmt.Printf("Executing update on %s\n", stream_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteStream %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := visionai.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &visionaipb.DeleteStreamRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteStream(ctx, req)
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
						Name:  "get-thumbnail",
						Usage: "get-thumbnail streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "event", Usage: "The name of the event.", Required: false},
							&cli.StringFlag{Name: "gcs-object-name", Usage: "The name of the GCS object to store the thumbnail image.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify the requests.", Required: false},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							stream := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
							fmt.Printf("Executing get-thumbnail on %s\n", stream)
							return nil
						},
					},

					{
						Name:  "generate-stream-hls-token",
						Usage: "generate-stream-hls-token streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							stream := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
							fmt.Printf("Executing generate-stream-hls-token on %s\n", stream)
							return nil
						},
					},
				},
			},
		},
	}
}
