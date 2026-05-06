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

package aiplatform

import (
	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the aiplatform command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "aiplatform",
		Usage: "manage Vertex AI API resources",
		Commands: []*cli.Command{
			{
				Name:  "annotation-specs",
				Usage: "Manage annotation-specs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe annotation-specs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "annotation-spec", Usage: "The ID of the annotation spec.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/annotationSpecs/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("annotation-spec"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetAnnotationSpecRequest{
								Name: name,
							}

							resp, err := client.GetAnnotationSpec(ctx, req)
							if err != nil {
								return err
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
				Name:  "annotations",
				Usage: "Manage annotations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list annotations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-item", Usage: "The ID of the data item.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/dataItems/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("data-item"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListAnnotationsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
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
				},
			},
			{
				Name:  "artifacts",
				Usage: "Manage artifacts resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "artifact-id", Usage: "The {artifact} portion of the resource name with the format:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateArtifactRequest{
								Parent:     parent,
								ArtifactId: cmd.String("artifact-id"),
							}

							resp, err := client.CreateArtifact(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "artifact", Usage: "The ID of the artifact.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/artifacts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("artifact"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetArtifactRequest{
								Name: name,
							}

							resp, err := client.GetArtifact(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter specifying the boolean condition for the Artifacts to satisfy in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "How the list of messages is ordered.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Artifacts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListArtifactsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListArtifacts(ctx, req)
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
						Usage: "update artifacts",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Artifact][google.", Required: false},
							&cli.StringFlag{Name: "artifact", Usage: "The ID of the artifact.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "artifact.name" not yet supported.
							artifact_name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/artifacts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("artifact"))
							fmt.Printf("Executing update on %s\n", artifact_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "artifact", Usage: "The ID of the artifact.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag of the Artifact to delete.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/artifacts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("artifact"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteArtifact %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteArtifactRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteArtifact(ctx, req)
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
						Name:  "purge",
						Usage: "purge artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A required filter matching the Artifacts to be purged.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Flag to indicate to actually perform the purge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.PurgeArtifactsRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
								Force:  cmd.Bool("force"),
							}

							op, err := client.PurgeArtifacts(ctx, req)
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
						Name:  "query-artifact-lineage-subgraph",
						Usage: "query-artifact-lineage-subgraph artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "artifact", Usage: "The ID of the artifact.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter specifying the boolean condition for the Artifacts to satisfy in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-hops", Usage: "Specifies the size of the lineage graph in terms of number of hops from the.", Required: false},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							artifact := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/artifacts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("artifact"))
							fmt.Printf("Executing query-artifact-lineage-subgraph on %s\n", artifact)
							return nil
						},
					},
				},
			},
			{
				Name:  "batch-prediction-jobs",
				Usage: "Manage batch-prediction-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create batch-prediction-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateBatchPredictionJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateBatchPredictionJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe batch-prediction-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "batch-prediction-job", Usage: "The ID of the batch prediction job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/batchPredictionJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("batch-prediction-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetBatchPredictionJobRequest{
								Name: name,
							}

							resp, err := client.GetBatchPredictionJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list batch-prediction-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListBatchPredictionJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBatchPredictionJobs(ctx, req)
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
						Usage: "delete batch-prediction-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "batch-prediction-job", Usage: "The ID of the batch prediction job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/batchPredictionJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("batch-prediction-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBatchPredictionJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteBatchPredictionJobRequest{
								Name: name,
							}

							op, err := client.DeleteBatchPredictionJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel batch-prediction-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "batch-prediction-job", Usage: "The ID of the batch prediction job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/batchPredictionJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("batch-prediction-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelBatchPredictionJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelBatchPredictionJobRequest{
								Name: name,
							}

							if err := client.CancelBatchPredictionJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "cached-contents",
				Usage: "Manage cached-contents resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create cached-contents",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateCachedContentRequest{
								Parent: parent,
							}

							resp, err := client.CreateCachedContent(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe cached-contents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cached-content", Usage: "The ID of the cached content.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cachedContents/%s", cmd.String("project"), cmd.String("location"), cmd.String("cached-content"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetCachedContentRequest{
								Name: name,
							}

							resp, err := client.GetCachedContent(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update cached-contents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cached-content", Usage: "The ID of the cached content.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "cached_content.name" not yet supported.
							cached_content_name := fmt.Sprintf("projects/%s/locations/%s/cachedContents/%s", cmd.String("project"), cmd.String("location"), cmd.String("cached-content"))
							fmt.Printf("Executing update on %s\n", cached_content_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete cached-contents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cached-content", Usage: "The ID of the cached content.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cachedContents/%s", cmd.String("project"), cmd.String("location"), cmd.String("cached-content"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCachedContent on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteCachedContentRequest{
								Name: name,
							}

							if err := client.DeleteCachedContent(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list cached-contents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of cached contents to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListCachedContents` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListCachedContentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCachedContents(ctx, req)
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
				Name:  "contexts",
				Usage: "Manage contexts resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context-id", Usage: "The {context} portion of the resource name with the format:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateContextRequest{
								Parent:    parent,
								ContextId: cmd.String("context-id"),
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
						Name:  "describe",
						Usage: "describe contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetContextRequest{
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
						Name:  "list",
						Usage: "list contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter specifying the boolean condition for the Contexts to satisfy in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "How the list of messages is ordered.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Contexts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListContextsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
						Name:  "update",
						Usage: "update contexts",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Context][google.", Required: false},
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "context.name" not yet supported.
							context_name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							fmt.Printf("Executing update on %s\n", context_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag of the Context to delete.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "The force deletion semantics is still undefined.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteContext %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteContextRequest{
								Name:  name,
								Force: cmd.Bool("force"),
								Etag:  cmd.String("etag"),
							}

							op, err := client.DeleteContext(ctx, req)
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
						Name:  "purge",
						Usage: "purge contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A required filter matching the Contexts to be purged.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Flag to indicate to actually perform the purge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.PurgeContextsRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
								Force:  cmd.Bool("force"),
							}

							op, err := client.PurgeContexts(ctx, req)
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
						Name:  "add-context-artifacts-and-executions",
						Usage: "add-context-artifacts-and-executions contexts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "artifacts", Usage: "The resource names of the Artifacts to attribute to the Context.", Required: false},
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringSliceFlag{Name: "executions", Usage: "The resource names of the Executions to associate with the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							context := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							fmt.Printf("Executing add-context-artifacts-and-executions on %s\n", context)
							return nil
						},
					},

					{
						Name:  "add-context-children",
						Usage: "add-context-children contexts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "child-contexts", Usage: "The resource names of the child Contexts.", Required: false},
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							context := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							fmt.Printf("Executing add-context-children on %s\n", context)
							return nil
						},
					},

					{
						Name:  "remove-context-children",
						Usage: "remove-context-children contexts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "child-contexts", Usage: "The resource names of the child Contexts.", Required: false},
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							context := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							fmt.Printf("Executing remove-context-children on %s\n", context)
							return nil
						},
					},

					{
						Name:  "query-context-lineage-subgraph",
						Usage: "query-context-lineage-subgraph contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							context := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("context"))
							fmt.Printf("Executing query-context-lineage-subgraph on %s\n", context)
							return nil
						},
					},
				},
			},
			{
				Name:  "custom-jobs",
				Usage: "Manage custom-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create custom-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateCustomJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateCustomJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe custom-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-job", Usage: "The ID of the custom job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetCustomJobRequest{
								Name: name,
							}

							resp, err := client.GetCustomJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list custom-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListCustomJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCustomJobs(ctx, req)
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
						Usage: "delete custom-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-job", Usage: "The ID of the custom job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCustomJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteCustomJobRequest{
								Name: name,
							}

							op, err := client.DeleteCustomJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel custom-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-job", Usage: "The ID of the custom job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelCustomJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelCustomJobRequest{
								Name: name,
							}

							if err := client.CancelCustomJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "data-items",
				Usage: "Manage data-items resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list data-items",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListDataItemsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataItems(ctx, req)
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
				Name:  "data-labeling-jobs",
				Usage: "Manage data-labeling-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-labeling-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateDataLabelingJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateDataLabelingJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe data-labeling-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-labeling-job", Usage: "The ID of the data labeling job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataLabelingJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-labeling-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetDataLabelingJobRequest{
								Name: name,
							}

							resp, err := client.GetDataLabelingJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list data-labeling-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListDataLabelingJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataLabelingJobs(ctx, req)
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
						Usage: "delete data-labeling-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-labeling-job", Usage: "The ID of the data labeling job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataLabelingJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-labeling-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataLabelingJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteDataLabelingJobRequest{
								Name: name,
							}

							op, err := client.DeleteDataLabelingJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel data-labeling-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-labeling-job", Usage: "The ID of the data labeling job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataLabelingJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-labeling-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelDataLabelingJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelDataLabelingJobRequest{
								Name: name,
							}

							if err := client.CancelDataLabelingJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "dataset-versions",
				Usage: "Manage dataset-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create dataset-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateDatasetVersionRequest{
								Parent: parent,
							}

							op, err := client.CreateDatasetVersion(ctx, req)
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
						Usage: "update dataset-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "dataset-version", Usage: "The ID of the dataset version.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "dataset_version.name" not yet supported.
							dataset_version_name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/datasetVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("dataset-version"))
							fmt.Printf("Executing update on %s\n", dataset_version_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete dataset-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "dataset-version", Usage: "The ID of the dataset version.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/datasetVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("dataset-version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDatasetVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteDatasetVersionRequest{
								Name: name,
							}

							op, err := client.DeleteDatasetVersion(ctx, req)
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
						Usage: "describe dataset-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "dataset-version", Usage: "The ID of the dataset version.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/datasetVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("dataset-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetDatasetVersionRequest{
								Name: name,
							}

							resp, err := client.GetDatasetVersion(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list dataset-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListDatasetVersionsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDatasetVersions(ctx, req)
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
						Name:  "restore",
						Usage: "restore dataset-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "dataset-version", Usage: "The ID of the dataset version.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/datasetVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("dataset-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.RestoreDatasetVersionRequest{
								Name: name,
							}

							op, err := client.RestoreDatasetVersion(ctx, req)
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
				Name:  "datasets",
				Usage: "Manage datasets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create datasets",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateDatasetRequest{
								Parent: parent,
							}

							op, err := client.CreateDataset(ctx, req)
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
						Usage: "describe datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetDatasetRequest{
								Name: name,
							}

							resp, err := client.GetDataset(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "dataset.name" not yet supported.
							dataset_name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing update on %s\n", dataset_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListDatasetsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDatasets(ctx, req)
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
						Usage: "delete datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataset %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteDatasetRequest{
								Name: name,
							}

							op, err := client.DeleteDataset(ctx, req)
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
						Name:  "import",
						Usage: "import datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ImportDataRequest{
								Name: name,
							}

							op, err := client.ImportData(ctx, req)
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
						Usage: "export datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ExportDataRequest{
								Name: name,
							}

							op, err := client.ExportData(ctx, req)
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
						Name:  "search-data-items",
						Usage: "search-data-items datasets",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "annotation-filters", Usage: "An expression that specifies what Annotations will be returned per.", Required: false},
							&cli.StringFlag{Name: "annotations-filter", Usage: "An expression for filtering the Annotations that will be returned per.", Required: false},
							&cli.IntFlag{Name: "annotations-limit", Usage: "If set, only up to this many of Annotations will be returned per.", Required: false},
							&cli.StringFlag{Name: "data-item-filter", Usage: "An expression for filtering the DataItem that will be returned.", Required: false},
							&cli.StringFlag{Name: "data-labeling-job", Usage: "The resource name of a DataLabelingJob.", Required: false},
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results for the server to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "saved-query", Usage: "The resource name of a SavedQuery(annotation set in UI).", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							dataset := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing search-data-items on %s\n", dataset)
							return nil
						},
					},
				},
			},
			{
				Name:  "deployment-resource-pools",
				Usage: "Manage deployment-resource-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create deployment-resource-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-resource-pool-id", Usage: "The ID to use for the DeploymentResourcePool, which.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateDeploymentResourcePoolRequest{
								Parent:                   parent,
								DeploymentResourcePoolId: cmd.String("deployment-resource-pool-id"),
							}

							op, err := client.CreateDeploymentResourcePool(ctx, req)
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
						Usage: "describe deployment-resource-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-resource-pool", Usage: "The ID of the deployment resource pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-resource-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetDeploymentResourcePoolRequest{
								Name: name,
							}

							resp, err := client.GetDeploymentResourcePool(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list deployment-resource-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of DeploymentResourcePools to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDeploymentResourcePools` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListDeploymentResourcePoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeploymentResourcePools(ctx, req)
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
						Usage: "update deployment-resource-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-resource-pool", Usage: "The ID of the deployment resource pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deployment_resource_pool.name" not yet supported.
							deployment_resource_pool_name := fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-resource-pool"))
							fmt.Printf("Executing update on %s\n", deployment_resource_pool_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete deployment-resource-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-resource-pool", Usage: "The ID of the deployment resource pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-resource-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDeploymentResourcePool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteDeploymentResourcePoolRequest{
								Name: name,
							}

							op, err := client.DeleteDeploymentResourcePool(ctx, req)
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
						Name:  "query-deployed-models",
						Usage: "query-deployed-models deployment-resource-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-resource-pool", Usage: "The ID of the deployment resource pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of DeployedModels to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `QueryDeployedModels` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							deployment_resource_pool := fmt.Sprintf("projects/%s/locations/%s/deploymentResourcePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-resource-pool"))
							fmt.Printf("Executing query-deployed-models on %s\n", deployment_resource_pool)
							return nil
						},
					},
				},
			},
			{
				Name:  "endpoints",
				Usage: "Manage endpoints resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-id", Usage: "The ID to use for endpoint, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateEndpointRequest{
								Parent:     parent,
								EndpointId: cmd.String("endpoint-id"),
							}

							op, err := client.CreateEndpoint(ctx, req)
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
						Usage: "describe endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetEndpointRequest{
								Name: name,
							}

							resp, err := client.GetEndpoint(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListEndpointsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEndpoints(ctx, req)
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
						Usage: "update endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "endpoint.name" not yet supported.
							endpoint_name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing update on %s\n", endpoint_name)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "endpoint.name" not yet supported.
							endpoint_name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing update on %s\n", endpoint_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEndpoint %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteEndpointRequest{
								Name: name,
							}

							op, err := client.DeleteEndpoint(ctx, req)
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
						Name:  "deploy-model",
						Usage: "deploy-model endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing deploy-model on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "undeploy-model",
						Usage: "undeploy-model endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployed-model-id", Usage: "The ID of the DeployedModel to be undeployed from the Endpoint.", Required: true},
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing undeploy-model on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "mutate-deployed-model",
						Usage: "mutate-deployed-model endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing mutate-deployed-model on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "count-tokens",
						Usage: "count-tokens endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The name of the publisher model requested to serve the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing count-tokens on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "compute-tokens",
						Usage: "compute-tokens endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The name of the publisher model requested to serve the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing compute-tokens on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "predict",
						Usage: "predict endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing predict on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "raw-predict",
						Usage: "raw-predict endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing raw-predict on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "stream-raw-predict",
						Usage: "stream-raw-predict endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing stream-raw-predict on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "direct-predict",
						Usage: "direct-predict endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing direct-predict on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "direct-raw-predict",
						Usage: "direct-raw-predict endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "input", Usage: "The prediction input.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "method-name", Usage: "Fully qualified name of the API method being invoked to perform.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing direct-raw-predict on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "server-streaming-predict",
						Usage: "server-streaming-predict endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing server-streaming-predict on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "explain",
						Usage: "explain endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployed-model-id", Usage: "If specified, this ExplainRequest will be served by the chosen.", Required: false},
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing explain on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "generate-content",
						Usage: "generate-content endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cached-content", Usage: "The name of the cached content used as context to serve the.", Required: false},
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The fully qualified name of the publisher model or tuned model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							model := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing generate-content on %s\n", model)
							return nil
						},
					},

					{
						Name:  "stream-generate-content",
						Usage: "stream-generate-content endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cached-content", Usage: "The name of the cached content used as context to serve the.", Required: false},
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The fully qualified name of the publisher model or tuned model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							model := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing stream-generate-content on %s\n", model)
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
						Name:  "read-feature-values",
						Usage: "read-feature-values entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-id", Usage: "ID for a specific entity.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							entity_type := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing read-feature-values on %s\n", entity_type)
							return nil
						},
					},

					{
						Name:  "streaming-read-feature-values",
						Usage: "streaming-read-feature-values entity-types",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "entity-ids", Usage: "IDs of entities to read Feature values of.", Required: true},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							entity_type := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing streaming-read-feature-values on %s\n", entity_type)
							return nil
						},
					},

					{
						Name:  "write-feature-values",
						Usage: "write-feature-values entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							entity_type := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing write-feature-values on %s\n", entity_type)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type-id", Usage: "The ID to use for the EntityType, which will become the final.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateEntityTypeRequest{
								Parent:       parent,
								EntityTypeId: cmd.String("entity-type-id"),
							}

							op, err := client.CreateEntityType(ctx, req)
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
						Usage: "describe entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetEntityTypeRequest{
								Name: name,
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
						Name:  "list",
						Usage: "list entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the EntityTypes that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of EntityTypes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListEntityTypesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
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
						Name:  "update",
						Usage: "update entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entity_type.name" not yet supported.
							entity_type_name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing update on %s\n", entity_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any Features for this EntityType will also be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEntityType %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteEntityTypeRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteEntityType(ctx, req)
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
						Name:  "import-feature-values",
						Usage: "import-feature-values entity-types",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "disable-ingestion-analysis", Usage: "If true, API doesn't start ingestion analysis pipeline.", Required: false},
							&cli.BoolFlag{Name: "disable-online-serving", Usage: "If set, data will not be imported for online serving.", Required: false},
							&cli.StringFlag{Name: "entity-id-field", Usage: "Source column that holds entity IDs.", Required: false},
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "worker-count", Usage: "Specifies the number of workers that are used to write data to the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							entity_type := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing import-feature-values on %s\n", entity_type)
							return nil
						},
					},

					{
						Name:  "export-feature-values",
						Usage: "export-feature-values entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							entity_type := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing export-feature-values on %s\n", entity_type)
							return nil
						},
					},

					{
						Name:  "delete-feature-values",
						Usage: "delete-feature-values entity-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							entity_type := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							fmt.Printf("Executing delete-feature-values on %s\n", entity_type)
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
						Name:  "import",
						Usage: "import evaluations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ImportModelEvaluationRequest{
								Parent: parent,
							}

							resp, err := client.ImportModelEvaluation(ctx, req)
							if err != nil {
								return err
							}
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s/evaluations/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"), cmd.String("evaluation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetModelEvaluationRequest{
								Name: name,
							}

							resp, err := client.GetModelEvaluation(ctx, req)
							if err != nil {
								return err
							}
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
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListModelEvaluationsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListModelEvaluations(ctx, req)
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
				Name:  "events",
				Usage: "Manage events resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of events to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListEventsRequest{
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
				},
			},
			{
				Name:  "executions",
				Usage: "Manage executions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution-id", Usage: "The {execution} portion of the resource name with the format:.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateExecutionRequest{
								Parent:      parent,
								ExecutionId: cmd.String("execution-id"),
							}

							resp, err := client.CreateExecution(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetExecutionRequest{
								Name: name,
							}

							resp, err := client.GetExecution(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter specifying the boolean condition for the Executions to satisfy in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "How the list of messages is ordered.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Executions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListExecutionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListExecutions(ctx, req)
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
						Usage: "update executions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the [Execution][google.", Required: false},
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "execution.name" not yet supported.
							execution_name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("execution"))
							fmt.Printf("Executing update on %s\n", execution_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag of the Execution to delete.", Required: false},
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("execution"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteExecution %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteExecutionRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteExecution(ctx, req)
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
						Name:  "purge",
						Usage: "purge executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A required filter matching the Executions to be purged.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Flag to indicate to actually perform the purge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.PurgeExecutionsRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
								Force:  cmd.Bool("force"),
							}

							op, err := client.PurgeExecutions(ctx, req)
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
						Name:  "add-execution-events",
						Usage: "add-execution-events executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							execution := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("execution"))
							fmt.Printf("Executing add-execution-events on %s\n", execution)
							return nil
						},
					},

					{
						Name:  "query-execution-inputs-and-outputs",
						Usage: "query-execution-inputs-and-outputs executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							execution := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("execution"))
							fmt.Printf("Executing query-execution-inputs-and-outputs on %s\n", execution)
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
						Name:  "create",
						Usage: "create experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-experiment-id", Usage: "The ID to use for the Tensorboard experiment, which becomes the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTensorboardExperimentRequest{
								Parent:                  parent,
								TensorboardExperimentId: cmd.String("tensorboard-experiment-id"),
							}

							resp, err := client.CreateTensorboardExperiment(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTensorboardExperimentRequest{
								Name: name,
							}

							resp, err := client.GetTensorboardExperiment(ctx, req)
							if err != nil {
								return err
							}
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
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tensorboard_experiment.name" not yet supported.
							tensorboard_experiment_name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							fmt.Printf("Executing update on %s\n", tensorboard_experiment_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the TensorboardExperiments that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of TensorboardExperiments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTensorboardExperimentsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTensorboardExperiments(ctx, req)
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
						Usage: "delete experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTensorboardExperiment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteTensorboardExperimentRequest{
								Name: name,
							}

							op, err := client.DeleteTensorboardExperiment(ctx, req)
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
						Name:  "batch-create",
						Usage: "batch-create experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchCreateTensorboardTimeSeriesRequest{
								Parent: parent,
							}

							resp, err := client.BatchCreateTensorboardTimeSeries(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "write",
						Usage: "write experiments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-experiment", Usage: "The resource name of the TensorboardExperiment to write data to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard_experiment := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							fmt.Printf("Executing write on %s\n", tensorboard_experiment)
							return nil
						},
					},
				},
			},
			{
				Name:  "feature-groups",
				Usage: "Manage feature-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create feature-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group-id", Usage: "The ID to use for this FeatureGroup, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateFeatureGroupRequest{
								Parent:         parent,
								FeatureGroupId: cmd.String("feature-group-id"),
							}

							op, err := client.CreateFeatureGroup(ctx, req)
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
						Usage: "describe feature-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeatureGroupRequest{
								Name: name,
							}

							resp, err := client.GetFeatureGroup(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list feature-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the FeatureGroups that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of FeatureGroups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeatureGroupsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatureGroups(ctx, req)
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
						Usage: "update feature-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feature_group.name" not yet supported.
							feature_group_name := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"))
							fmt.Printf("Executing update on %s\n", feature_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete feature-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any Features under this FeatureGroup.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeatureGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteFeatureGroupRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteFeatureGroup(ctx, req)
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
				Name:  "feature-online-stores",
				Usage: "Manage feature-online-stores resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create feature-online-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store-id", Usage: "The ID to use for this FeatureOnlineStore, which will become the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateFeatureOnlineStoreRequest{
								Parent:               parent,
								FeatureOnlineStoreId: cmd.String("feature-online-store-id"),
							}

							op, err := client.CreateFeatureOnlineStore(ctx, req)
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
						Usage: "describe feature-online-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeatureOnlineStoreRequest{
								Name: name,
							}

							resp, err := client.GetFeatureOnlineStore(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list feature-online-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the FeatureOnlineStores that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of FeatureOnlineStores to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeatureOnlineStoresRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatureOnlineStores(ctx, req)
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
						Usage: "update feature-online-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feature_online_store.name" not yet supported.
							feature_online_store_name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"))
							fmt.Printf("Executing update on %s\n", feature_online_store_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete feature-online-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any FeatureViews and Features for this FeatureOnlineStore.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeatureOnlineStore %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteFeatureOnlineStoreRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteFeatureOnlineStore(ctx, req)
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
				Name:  "feature-view-syncs",
				Usage: "Manage feature-view-syncs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe feature-view-syncs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "feature-view-sync", Usage: "The ID of the feature view sync.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s/featureViewSyncs/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"), cmd.String("feature-view-sync"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeatureViewSyncRequest{
								Name: name,
							}

							resp, err := client.GetFeatureViewSync(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list feature-view-syncs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the FeatureViewSyncs that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of FeatureViewSyncs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeatureViewSyncsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatureViewSyncs(ctx, req)
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
				Name:  "feature-views",
				Usage: "Manage feature-views resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view-id", Usage: "The ID to use for the FeatureView, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "run-sync-immediately", Usage: "If set to true, one on demand sync will be run immediately,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateFeatureViewRequest{
								Parent:             parent,
								FeatureViewId:      cmd.String("feature-view-id"),
								RunSyncImmediately: cmd.Bool("run-sync-immediately"),
							}

							op, err := client.CreateFeatureView(ctx, req)
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
						Usage: "describe feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeatureViewRequest{
								Name: name,
							}

							resp, err := client.GetFeatureView(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the FeatureViews that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of FeatureViews to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeatureViewsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatureViews(ctx, req)
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
						Usage: "update feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feature_view.name" not yet supported.
							feature_view_name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							fmt.Printf("Executing update on %s\n", feature_view_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeatureView %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteFeatureViewRequest{
								Name: name,
							}

							op, err := client.DeleteFeatureView(ctx, req)
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
						Name:  "sync",
						Usage: "sync feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							feature_view := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							fmt.Printf("Executing sync on %s\n", feature_view)
							return nil
						},
					},

					{
						Name:  "fetch-feature-values",
						Usage: "fetch-feature-values feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-format", Usage: "Response data format.", Required: false},
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							feature_view := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							fmt.Printf("Executing fetch-feature-values on %s\n", feature_view)
							return nil
						},
					},

					{
						Name:  "search-nearest-entities",
						Usage: "search-nearest-entities feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-full-entity", Usage: "If set to true, the full entities (including all vector values.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							feature_view := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							fmt.Printf("Executing search-nearest-entities on %s\n", feature_view)
							return nil
						},
					},

					{
						Name:  "direct-write",
						Usage: "direct-write feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							feature_view := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							fmt.Printf("Executing direct-write on %s\n", feature_view)
							return nil
						},
					},

					{
						Name:  "generate-fetch-access-token",
						Usage: "generate-fetch-access-token feature-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-online-store", Usage: "The ID of the feature online store.", Required: true},
							&cli.StringFlag{Name: "feature-view", Usage: "The ID of the feature view.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							feature_view := fmt.Sprintf("projects/%s/locations/%s/featureOnlineStores/%s/featureViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-online-store"), cmd.String("feature-view"))
							fmt.Printf("Executing generate-fetch-access-token on %s\n", feature_view)
							return nil
						},
					},
				},
			},
			{
				Name:  "features",
				Usage: "Manage features resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "feature-id", Usage: "The ID to use for the Feature, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateFeatureRequest{
								Parent:    parent,
								FeatureId: cmd.String("feature-id"),
							}

							op, err := client.CreateFeature(ctx, req)
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
						Usage: "batch-create features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchCreateFeaturesRequest{
								Parent: parent,
							}

							op, err := client.BatchCreateFeatures(ctx, req)
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
						Usage: "describe features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"), cmd.String("feature"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeatureRequest{
								Name: name,
							}

							resp, err := client.GetFeature(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the Features that match the filter expression.", Required: false},
							&cli.IntFlag{Name: "latest-stats-count", Usage: "Only applicable for Vertex AI Feature Store (Legacy).", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Features to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeaturesRequest{
								Parent:           parent,
								Filter:           cmd.String("filter"),
								PageSize:         int32(cmd.Int("page-size")),
								PageToken:        cmd.String("page-token"),
								OrderBy:          cmd.String("order-by"),
								LatestStatsCount: int32(cmd.Int("latest-stats-count")),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatures(ctx, req)
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
						Usage: "update features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feature.name" not yet supported.
							feature_name := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"), cmd.String("feature"))
							fmt.Printf("Executing update on %s\n", feature_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "feature-group", Usage: "The ID of the feature group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featureGroups/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature-group"), cmd.String("feature"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeature %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteFeatureRequest{
								Name: name,
							}

							op, err := client.DeleteFeature(ctx, req)
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
						Name:  "create",
						Usage: "create features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "feature-id", Usage: "The ID to use for the Feature, which will become the final.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateFeatureRequest{
								Parent:    parent,
								FeatureId: cmd.String("feature-id"),
							}

							op, err := client.CreateFeature(ctx, req)
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
						Usage: "batch-create features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchCreateFeaturesRequest{
								Parent: parent,
							}

							op, err := client.BatchCreateFeatures(ctx, req)
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
						Usage: "describe features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"), cmd.String("feature"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeatureRequest{
								Name: name,
							}

							resp, err := client.GetFeature(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the Features that match the filter expression.", Required: false},
							&cli.IntFlag{Name: "latest-stats-count", Usage: "Only applicable for Vertex AI Feature Store (Legacy).", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Features to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeaturesRequest{
								Parent:           parent,
								Filter:           cmd.String("filter"),
								PageSize:         int32(cmd.Int("page-size")),
								PageToken:        cmd.String("page-token"),
								OrderBy:          cmd.String("order-by"),
								LatestStatsCount: int32(cmd.Int("latest-stats-count")),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatures(ctx, req)
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
						Usage: "update features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feature.name" not yet supported.
							feature_name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"), cmd.String("feature"))
							fmt.Printf("Executing update on %s\n", feature_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entity-type", Usage: "The ID of the entity type.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s/entityTypes/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"), cmd.String("entity-type"), cmd.String("feature"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeature %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteFeatureRequest{
								Name: name,
							}

							op, err := client.DeleteFeature(ctx, req)
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
				Name:  "featurestores",
				Usage: "Manage featurestores resources",
				Commands: []*cli.Command{

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore-id", Usage: "The ID to use for this Featurestore, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateFeaturestoreRequest{
								Parent:         parent,
								FeaturestoreId: cmd.String("featurestore-id"),
							}

							op, err := client.CreateFeaturestore(ctx, req)
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
						Usage: "describe featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetFeaturestoreRequest{
								Name: name,
							}

							resp, err := client.GetFeaturestore(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the featurestores that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Featurestores to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListFeaturestoresRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeaturestores(ctx, req)
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
						Usage: "update featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "featurestore.name" not yet supported.
							featurestore_name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing update on %s\n", featurestore_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any EntityTypes and Features for this Featurestore will.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeaturestore %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteFeaturestoreRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteFeaturestore(ctx, req)
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
						Name:  "batch-read-feature-values",
						Usage: "batch-read-feature-values featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							featurestore := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing batch-read-feature-values on %s\n", featurestore)
							return nil
						},
					},

					{
						Name:  "search-features",
						Usage: "search-features featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Features to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Query string that is a conjunction of field-restricted queries and/or.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing search-features on %s\n", location)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions featurestores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "featurestore", Usage: "The ID of the featurestore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/featurestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("featurestore"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "hyperparameter-tuning-jobs",
				Usage: "Manage hyperparameter-tuning-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create hyperparameter-tuning-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateHyperparameterTuningJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateHyperparameterTuningJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe hyperparameter-tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hyperparameter-tuning-job", Usage: "The ID of the hyperparameter tuning job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/hyperparameterTuningJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("hyperparameter-tuning-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetHyperparameterTuningJobRequest{
								Name: name,
							}

							resp, err := client.GetHyperparameterTuningJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list hyperparameter-tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListHyperparameterTuningJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListHyperparameterTuningJobs(ctx, req)
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
						Usage: "delete hyperparameter-tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hyperparameter-tuning-job", Usage: "The ID of the hyperparameter tuning job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/hyperparameterTuningJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("hyperparameter-tuning-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteHyperparameterTuningJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteHyperparameterTuningJobRequest{
								Name: name,
							}

							op, err := client.DeleteHyperparameterTuningJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel hyperparameter-tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hyperparameter-tuning-job", Usage: "The ID of the hyperparameter tuning job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/hyperparameterTuningJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("hyperparameter-tuning-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelHyperparameterTuningJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelHyperparameterTuningJobRequest{
								Name: name,
							}

							if err := client.CancelHyperparameterTuningJob(ctx, req); err != nil {
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
						Name:  "create",
						Usage: "create index-endpoints",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateIndexEndpointRequest{
								Parent: parent,
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetIndexEndpointRequest{
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
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListIndexEndpointsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteIndexEndpointRequest{
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
							&cli.StringFlag{Name: "deployed-index-id", Usage: "The ID of the DeployedIndex to be undeployed from the.", Required: true},
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

					{
						Name:  "mutate-deployed-index",
						Usage: "mutate-deployed-index index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index_endpoint := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing mutate-deployed-index on %s\n", index_endpoint)
							return nil
						},
					},

					{
						Name:  "find-neighbors",
						Usage: "find-neighbors index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployed-index-id", Usage: "The ID of the DeployedIndex that will serve the request.", Required: false},
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-full-datapoint", Usage: "If set to true, the full datapoints (including all vector values and.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index_endpoint := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing find-neighbors on %s\n", index_endpoint)
							return nil
						},
					},

					{
						Name:  "read-index-datapoints",
						Usage: "read-index-datapoints index-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployed-index-id", Usage: "The ID of the DeployedIndex that will serve the request.", Required: false},
							&cli.StringSliceFlag{Name: "ids", Usage: "IDs of the datapoints to be searched for.", Required: false},
							&cli.StringFlag{Name: "index-endpoint", Usage: "The ID of the index endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index_endpoint := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index-endpoint"))
							fmt.Printf("Executing read-index-datapoints on %s\n", index_endpoint)
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
						Name:  "create",
						Usage: "create indexes",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateIndexRequest{
								Parent: parent,
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
						Name:  "describe",
						Usage: "describe indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("indexe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetIndexRequest{
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
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListIndexesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
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
						Name:  "update",
						Usage: "update indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "index.name" not yet supported.
							index_name := fmt.Sprintf("projects/%s/locations/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("indexe"))
							fmt.Printf("Executing update on %s\n", index_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("indexe"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIndex %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteIndexRequest{
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

					{
						Name:  "upsert-datapoints",
						Usage: "upsert-datapoints indexes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "index", Usage: "The name of the Index resource to be updated.", Required: true},
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index := fmt.Sprintf("projects/%s/locations/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("indexe"))
							fmt.Printf("Executing upsert-datapoints on %s\n", index)
							return nil
						},
					},

					{
						Name:  "remove-datapoints",
						Usage: "remove-datapoints indexes",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "datapoint-ids", Usage: "A list of datapoint ids to be deleted.", Required: false},
							&cli.StringFlag{Name: "index", Usage: "The name of the Index resource to be updated.", Required: true},
							&cli.StringFlag{Name: "indexe", Usage: "The ID of the indexe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							index := fmt.Sprintf("projects/%s/locations/%s/indexes/%s", cmd.String("project"), cmd.String("location"), cmd.String("indexe"))
							fmt.Printf("Executing remove-datapoints on %s\n", index)
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
						Name:  "generate-synthetic-data",
						Usage: "generate-synthetic-data locations",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "count", Usage: "The number of synthetic examples to generate.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing generate-synthetic-data on %s\n", location)
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
						Name:  "evaluate-instances",
						Usage: "evaluate-instances locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing evaluate-instances on %s\n", location)
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
						Name:  "deploy",
						Usage: "deploy locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination", Usage: "The resource name of the Location to deploy the model in.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							destination := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing deploy on %s\n", destination)
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
						Name:  "retrieve-contexts",
						Usage: "retrieve-contexts locations",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.RetrieveContextsRequest{
								Parent: parent,
							}

							resp, err := client.RetrieveContexts(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "augment-prompt",
						Usage: "augment-prompt locations",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.AugmentPromptRequest{
								Parent: parent,
							}

							resp, err := client.AugmentPrompt(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "corroborate-content",
						Usage: "corroborate-content locations",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CorroborateContentRequest{
								Parent: parent,
							}

							resp, err := client.CorroborateContent(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "ask-contexts",
						Usage: "ask-contexts locations",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.AskContextsRequest{
								Parent: parent,
							}

							resp, err := client.AskContexts(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "async-retrieve-contexts",
						Usage: "async-retrieve-contexts locations",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.AsyncRetrieveContextsRequest{
								Parent: parent,
							}

							op, err := client.AsyncRetrieveContexts(ctx, req)
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
				Name:  "metadata-schemas",
				Usage: "Manage metadata-schemas resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create metadata-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-schema-id", Usage: "The {metadata_schema} portion of the resource name with the format:.", Required: false},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateMetadataSchemaRequest{
								Parent:           parent,
								MetadataSchemaId: cmd.String("metadata-schema-id"),
							}

							resp, err := client.CreateMetadataSchema(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe metadata-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-schema", Usage: "The ID of the metadata schema.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s/metadataSchemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"), cmd.String("metadata-schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetMetadataSchemaRequest{
								Name: name,
							}

							resp, err := client.GetMetadataSchema(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list metadata-schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A query to filter available MetadataSchemas for matching results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of MetadataSchemas to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListMetadataSchemasRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMetadataSchemas(ctx, req)
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
				Name:  "metadata-stores",
				Usage: "Manage metadata-stores resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create metadata-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store-id", Usage: "The {metadatastore} portion of the resource name with the format:.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateMetadataStoreRequest{
								Parent:          parent,
								MetadataStoreId: cmd.String("metadata-store-id"),
							}

							op, err := client.CreateMetadataStore(ctx, req)
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
						Usage: "describe metadata-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetMetadataStoreRequest{
								Name: name,
							}

							resp, err := client.GetMetadataStore(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list metadata-stores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Metadata Stores to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListMetadataStoresRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMetadataStores(ctx, req)
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
						Usage: "delete metadata-stores",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "Deprecated: Field is no longer supported.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-store", Usage: "The ID of the metadata store.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataStores/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-store"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMetadataStore %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteMetadataStoreRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteMetadataStore(ctx, req)
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
				Name:  "migratable-resources",
				Usage: "Manage migratable-resources resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search migratable-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter for your search.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.SearchMigratableResourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.SearchMigratableResources(ctx, req)
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
						Name:  "batch-migrate",
						Usage: "batch-migrate migratable-resources",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchMigrateResourcesRequest{
								Parent: parent,
							}

							op, err := client.BatchMigrateResources(ctx, req)
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
				Name:  "model-deployment-monitoring-jobs",
				Usage: "Manage model-deployment-monitoring-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create model-deployment-monitoring-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateModelDeploymentMonitoringJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateModelDeploymentMonitoringJob(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "search-model-deployment-monitoring-stats-anomalies",
						Usage: "search-model-deployment-monitoring-stats-anomalies model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployed-model-id", Usage: "The DeployedModel ID of the.", Required: true},
							&cli.StringFlag{Name: "feature-display-name", Usage: "The feature display name.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-deployment-monitoring-job", Usage: "The ID of the model deployment monitoring job.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							model_deployment_monitoring_job := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("model-deployment-monitoring-job"))
							fmt.Printf("Executing search-model-deployment-monitoring-stats-anomalies on %s\n", model_deployment_monitoring_job)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-deployment-monitoring-job", Usage: "The ID of the model deployment monitoring job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("model-deployment-monitoring-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetModelDeploymentMonitoringJobRequest{
								Name: name,
							}

							resp, err := client.GetModelDeploymentMonitoringJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListModelDeploymentMonitoringJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListModelDeploymentMonitoringJobs(ctx, req)
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
						Usage: "update model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-deployment-monitoring-job", Usage: "The ID of the model deployment monitoring job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "model_deployment_monitoring_job.name" not yet supported.
							model_deployment_monitoring_job_name := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("model-deployment-monitoring-job"))
							fmt.Printf("Executing update on %s\n", model_deployment_monitoring_job_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-deployment-monitoring-job", Usage: "The ID of the model deployment monitoring job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("model-deployment-monitoring-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteModelDeploymentMonitoringJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteModelDeploymentMonitoringJobRequest{
								Name: name,
							}

							op, err := client.DeleteModelDeploymentMonitoringJob(ctx, req)
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
						Name:  "pause",
						Usage: "pause model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-deployment-monitoring-job", Usage: "The ID of the model deployment monitoring job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("model-deployment-monitoring-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute PauseModelDeploymentMonitoringJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.PauseModelDeploymentMonitoringJobRequest{
								Name: name,
							}

							if err := client.PauseModelDeploymentMonitoringJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "resume",
						Usage: "resume model-deployment-monitoring-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-deployment-monitoring-job", Usage: "The ID of the model deployment monitoring job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/modelDeploymentMonitoringJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("model-deployment-monitoring-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute ResumeModelDeploymentMonitoringJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ResumeModelDeploymentMonitoringJobRequest{
								Name: name,
							}

							if err := client.ResumeModelDeploymentMonitoringJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "describe",
						Usage: "describe models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hugging-face-token", Usage: "Token used to access Hugging Face gated models.", Required: false},
							&cli.BoolFlag{Name: "is-hugging-face-model", Usage: "Boolean indicates whether the requested model is a Hugging Face.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The IETF BCP-47 language code representing the language in which.", Required: false},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "publisher", Usage: "The ID of the publisher.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "PublisherModel view specifying which fields to read.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("publishers/%s/models/%s", cmd.String("publisher"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetPublisherModelRequest{
								Name:               name,
								LanguageCode:       cmd.String("language-code"),
								View:               aiplatformpb.PublisherModelView(aiplatformpb.PublisherModelView_value[cmd.String("view")]),
								IsHuggingFaceModel: cmd.Bool("is-hugging-face-model"),
								HuggingFaceToken:   cmd.String("hugging-face-token"),
							}

							resp, err := client.GetPublisherModel(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "upload models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model-id", Usage: "The ID to use for the uploaded Model, which will become the final.", Required: false},
							&cli.StringFlag{Name: "parent-model", Usage: "The resource name of the model into which to upload the version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-account", Usage: "The user-provided custom service account to use to do the model.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.UploadModelRequest{
								Parent:         parent,
								ParentModel:    cmd.String("parent-model"),
								ModelId:        cmd.String("model-id"),
								ServiceAccount: cmd.String("service-account"),
							}

							op, err := client.UploadModel(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetModelRequest{
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
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListModelsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
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
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListModelVersionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListModelVersions(ctx, req)
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
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListModelVersionCheckpointsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListModelVersionCheckpoints(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "model.name" not yet supported.
							model_name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							fmt.Printf("Executing update on %s\n", model_name)
							return nil
						},
					},

					{
						Name:  "update-explanation-dataset",
						Usage: "update-explanation-dataset models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							model := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							fmt.Printf("Executing update-explanation-dataset on %s\n", model)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteModel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteModelRequest{
								Name: name,
							}

							op, err := client.DeleteModel(ctx, req)
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
						Name:  "delete",
						Usage: "delete models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteModelVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteModelVersionRequest{
								Name: name,
							}

							op, err := client.DeleteModelVersion(ctx, req)
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
						Name:  "merge-version-aliases",
						Usage: "merge-version-aliases models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "version-aliases", Usage: "The set of version aliases to merge.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.MergeVersionAliasesRequest{
								Name:           name,
								VersionAliases: cmd.StringSlice("version-aliases"),
							}

							resp, err := client.MergeVersionAliases(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "export models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ExportModelRequest{
								Name: name,
							}

							op, err := client.ExportModel(ctx, req)
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
						Name:  "copy",
						Usage: "copy models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-service-account", Usage: "The user-provided custom service account to use to do the copy.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-model", Usage: "The resource name of the Model to copy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CopyModelRequest{
								Parent:               parent,
								SourceModel:          cmd.String("source-model"),
								CustomServiceAccount: cmd.String("custom-service-account"),
							}

							op, err := client.CopyModel(ctx, req)
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
						Name:  "embed-content",
						Usage: "embed-content models",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-truncate", Usage: "Deprecated: Please use EmbedContentConfig.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.IntFlag{Name: "output-dimensionality", Usage: "Deprecated: Please use EmbedContentConfig.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "publisher", Usage: "The ID of the publisher.", Required: true},
							&cli.StringFlag{Name: "task-type", Usage: "Deprecated: Please use EmbedContentConfig.", Required: false},
							&cli.StringFlag{Name: "title", Usage: "Deprecated: Please use EmbedContentConfig.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							model := fmt.Sprintf("projects/%s/locations/%s/publishers/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("publisher"), cmd.String("model"))
							fmt.Printf("Executing embed-content on %s\n", model)
							return nil
						},
					},
				},
			},
			{
				Name:  "nas-jobs",
				Usage: "Manage nas-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create nas-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateNasJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateNasJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe nas-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nas-job", Usage: "The ID of the nas job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nasJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("nas-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetNasJobRequest{
								Name: name,
							}

							resp, err := client.GetNasJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list nas-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListNasJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListNasJobs(ctx, req)
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
						Usage: "delete nas-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nas-job", Usage: "The ID of the nas job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nasJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("nas-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNasJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteNasJobRequest{
								Name: name,
							}

							op, err := client.DeleteNasJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel nas-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nas-job", Usage: "The ID of the nas job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nasJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("nas-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelNasJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelNasJobRequest{
								Name: name,
							}

							if err := client.CancelNasJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "nas-trial-details",
				Usage: "Manage nas-trial-details resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe nas-trial-details",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nas-job", Usage: "The ID of the nas job.", Required: true},
							&cli.StringFlag{Name: "nas-trial-detail", Usage: "The ID of the nas trial detail.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nasJobs/%s/nasTrialDetails/%s", cmd.String("project"), cmd.String("location"), cmd.String("nas-job"), cmd.String("nas-trial-detail"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetNasTrialDetailRequest{
								Name: name,
							}

							resp, err := client.GetNasTrialDetail(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list nas-trial-details",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nas-job", Usage: "The ID of the nas job.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/nasJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("nas-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListNasTrialDetailsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListNasTrialDetails(ctx, req)
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
				Name:  "notebook-execution-jobs",
				Usage: "Manage notebook-execution-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create notebook-execution-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-execution-job-id", Usage: "User specified ID for the NotebookExecutionJob.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateNotebookExecutionJobRequest{
								Parent:                 parent,
								NotebookExecutionJobId: cmd.String("notebook-execution-job-id"),
							}

							op, err := client.CreateNotebookExecutionJob(ctx, req)
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
						Usage: "describe notebook-execution-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-execution-job", Usage: "The ID of the notebook execution job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The NotebookExecutionJob view.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookExecutionJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-execution-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetNotebookExecutionJobRequest{
								Name: name,
								View: aiplatformpb.NotebookExecutionJobView(aiplatformpb.NotebookExecutionJobView_value[cmd.String("view")]),
							}

							resp, err := client.GetNotebookExecutionJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list notebook-execution-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The NotebookExecutionJob view.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListNotebookExecutionJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								View:      aiplatformpb.NotebookExecutionJobView(aiplatformpb.NotebookExecutionJobView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListNotebookExecutionJobs(ctx, req)
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
						Usage: "delete notebook-execution-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-execution-job", Usage: "The ID of the notebook execution job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookExecutionJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-execution-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNotebookExecutionJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteNotebookExecutionJobRequest{
								Name: name,
							}

							op, err := client.DeleteNotebookExecutionJob(ctx, req)
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
				Name:  "notebook-runtime-templates",
				Usage: "Manage notebook-runtime-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create notebook-runtime-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime-template-id", Usage: "User specified ID for the notebook runtime template.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateNotebookRuntimeTemplateRequest{
								Parent:                    parent,
								NotebookRuntimeTemplateId: cmd.String("notebook-runtime-template-id"),
							}

							op, err := client.CreateNotebookRuntimeTemplate(ctx, req)
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
						Usage: "describe notebook-runtime-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime-template", Usage: "The ID of the notebook runtime template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetNotebookRuntimeTemplateRequest{
								Name: name,
							}

							resp, err := client.GetNotebookRuntimeTemplate(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list notebook-runtime-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListNotebookRuntimeTemplatesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNotebookRuntimeTemplates(ctx, req)
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
						Usage: "delete notebook-runtime-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime-template", Usage: "The ID of the notebook runtime template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNotebookRuntimeTemplate %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteNotebookRuntimeTemplateRequest{
								Name: name,
							}

							op, err := client.DeleteNotebookRuntimeTemplate(ctx, req)
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
						Usage: "update notebook-runtime-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime-template", Usage: "The ID of the notebook runtime template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "notebook_runtime_template.name" not yet supported.
							notebook_runtime_template_name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimeTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime-template"))
							fmt.Printf("Executing update on %s\n", notebook_runtime_template_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "notebook-runtimes",
				Usage: "Manage notebook-runtimes resources",
				Commands: []*cli.Command{

					{
						Name:  "assign",
						Usage: "assign notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime-id", Usage: "User specified ID for the notebook runtime.", Required: false},
							&cli.StringFlag{Name: "notebook-runtime-template", Usage: "The resource name of the NotebookRuntimeTemplate based on which a.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.AssignNotebookRuntimeRequest{
								Parent:                  parent,
								NotebookRuntimeTemplate: cmd.String("notebook-runtime-template"),
								NotebookRuntimeId:       cmd.String("notebook-runtime-id"),
							}

							op, err := client.AssignNotebookRuntime(ctx, req)
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
						Usage: "describe notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime", Usage: "The ID of the notebook runtime.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetNotebookRuntimeRequest{
								Name: name,
							}

							resp, err := client.GetNotebookRuntime(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListNotebookRuntimesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNotebookRuntimes(ctx, req)
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
						Usage: "delete notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime", Usage: "The ID of the notebook runtime.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNotebookRuntime %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteNotebookRuntimeRequest{
								Name: name,
							}

							op, err := client.DeleteNotebookRuntime(ctx, req)
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
						Name:  "upgrade",
						Usage: "upgrade notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime", Usage: "The ID of the notebook runtime.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.UpgradeNotebookRuntimeRequest{
								Name: name,
							}

							op, err := client.UpgradeNotebookRuntime(ctx, req)
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
						Name:  "start",
						Usage: "start notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime", Usage: "The ID of the notebook runtime.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.StartNotebookRuntimeRequest{
								Name: name,
							}

							op, err := client.StartNotebookRuntime(ctx, req)
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
						Name:  "stop",
						Usage: "stop notebook-runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notebook-runtime", Usage: "The ID of the notebook runtime.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/notebookRuntimes/%s", cmd.String("project"), cmd.String("location"), cmd.String("notebook-runtime"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.StopNotebookRuntimeRequest{
								Name: name,
							}

							op, err := client.StopNotebookRuntime(ctx, req)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "persistent-resources",
				Usage: "Manage persistent-resources resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create persistent-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "persistent-resource-id", Usage: "The ID to use for the PersistentResource, which become the final.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreatePersistentResourceRequest{
								Parent:               parent,
								PersistentResourceId: cmd.String("persistent-resource-id"),
							}

							op, err := client.CreatePersistentResource(ctx, req)
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
						Usage: "describe persistent-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "persistent-resource", Usage: "The ID of the persistent resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/persistentResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("persistent-resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetPersistentResourceRequest{
								Name: name,
							}

							resp, err := client.GetPersistentResource(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list persistent-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListPersistentResourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPersistentResources(ctx, req)
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
						Usage: "delete persistent-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "persistent-resource", Usage: "The ID of the persistent resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/persistentResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("persistent-resource"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePersistentResource %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeletePersistentResourceRequest{
								Name: name,
							}

							op, err := client.DeletePersistentResource(ctx, req)
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
						Usage: "update persistent-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "persistent-resource", Usage: "The ID of the persistent resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "persistent_resource.name" not yet supported.
							persistent_resource_name := fmt.Sprintf("projects/%s/locations/%s/persistentResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("persistent-resource"))
							fmt.Printf("Executing update on %s\n", persistent_resource_name)
							return nil
						},
					},

					{
						Name:  "reboot",
						Usage: "reboot persistent-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "persistent-resource", Usage: "The ID of the persistent resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/persistentResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("persistent-resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.RebootPersistentResourceRequest{
								Name: name,
							}

							op, err := client.RebootPersistentResource(ctx, req)
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
				Name:  "pipeline-jobs",
				Usage: "Manage pipeline-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline-job-id", Usage: "The ID to use for the PipelineJob, which will become the final component of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreatePipelineJobRequest{
								Parent:        parent,
								PipelineJobId: cmd.String("pipeline-job-id"),
							}

							resp, err := client.CreatePipelineJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline-job", Usage: "The ID of the pipeline job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pipelineJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("pipeline-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetPipelineJobRequest{
								Name: name,
							}

							resp, err := client.GetPipelineJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the PipelineJobs that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListPipelineJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPipelineJobs(ctx, req)
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
						Usage: "delete pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline-job", Usage: "The ID of the pipeline job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pipelineJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("pipeline-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePipelineJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeletePipelineJobRequest{
								Name: name,
							}

							op, err := client.DeletePipelineJob(ctx, req)
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
						Usage: "batch-delete pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the PipelineJobs to delete.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchDeletePipelineJobsRequest{
								Parent: parent,
								Names:  cmd.StringSlice("names"),
							}

							op, err := client.BatchDeletePipelineJobs(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline-job", Usage: "The ID of the pipeline job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pipelineJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("pipeline-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelPipelineJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelPipelineJobRequest{
								Name: name,
							}

							if err := client.CancelPipelineJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "batch-cancel",
						Usage: "batch-cancel pipeline-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the PipelineJobs to cancel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchCancelPipelineJobsRequest{
								Parent: parent,
								Names:  cmd.StringSlice("names"),
							}

							op, err := client.BatchCancelPipelineJobs(ctx, req)
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
				Name:  "rag-corpora",
				Usage: "Manage rag-corpora resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create rag-corpora",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateRagCorpusRequest{
								Parent: parent,
							}

							op, err := client.CreateRagCorpus(ctx, req)
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
						Usage: "update rag-corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "rag_corpus.name" not yet supported.
							rag_corpus_name := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"))
							fmt.Printf("Executing update on %s\n", rag_corpus_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe rag-corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetRagCorpusRequest{
								Name: name,
							}

							resp, err := client.GetRagCorpus(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list rag-corpora",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListRagCorporaRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRagCorpora(ctx, req)
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
						Usage: "delete rag-corpora",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any RagFiles in this RagCorpus will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRagCorpus %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteRagCorpusRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteRagCorpus(ctx, req)
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
				Name:  "rag-engine-config",
				Usage: "Manage rag-engine-config resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update rag-engine-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "rag_engine_config.name" not yet supported.
							rag_engine_config_name := fmt.Sprintf("projects/%s/locations/%s/ragEngineConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", rag_engine_config_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe rag-engine-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ragEngineConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetRagEngineConfigRequest{
								Name: name,
							}

							resp, err := client.GetRagEngineConfig(ctx, req)
							if err != nil {
								return err
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
				Name:  "rag-files",
				Usage: "Manage rag-files resources",
				Commands: []*cli.Command{

					{
						Name:  "upload",
						Usage: "upload rag-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.UploadRagFileRequest{
								Parent: parent,
							}

							resp, err := client.UploadRagFile(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "import rag-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ImportRagFilesRequest{
								Parent: parent,
							}

							op, err := client.ImportRagFiles(ctx, req)
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
						Usage: "describe rag-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
							&cli.StringFlag{Name: "rag-file", Usage: "The ID of the rag file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s/ragFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"), cmd.String("rag-file"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetRagFileRequest{
								Name: name,
							}

							resp, err := client.GetRagFile(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list rag-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListRagFilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRagFiles(ctx, req)
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
						Usage: "delete rag-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rag-corpora", Usage: "The ID of the rag corpora.", Required: true},
							&cli.StringFlag{Name: "rag-file", Usage: "The ID of the rag file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ragCorpora/%s/ragFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("rag-corpora"), cmd.String("rag-file"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRagFile %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteRagFileRequest{
								Name: name,
							}

							op, err := client.DeleteRagFile(ctx, req)
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
				Name:  "reasoning-engines",
				Usage: "Manage reasoning-engines resources",
				Commands: []*cli.Command{

					{
						Name:  "query",
						Usage: "query reasoning-engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "class-method", Usage: "Class method to be used for the query.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.QueryReasoningEngineRequest{
								Name:        name,
								ClassMethod: cmd.String("class-method"),
							}

							resp, err := client.QueryReasoningEngine(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "stream-query",
						Usage: "stream-query reasoning-engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "class-method", Usage: "Class method to be used for the stream query.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.StreamQueryReasoningEngineRequest{
								Name:        name,
								ClassMethod: cmd.String("class-method"),
							}

							resp, err := client.StreamQueryReasoningEngine(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "async-query",
						Usage: "async-query reasoning-engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "input-gcs-uri", Usage: "Input Cloud Storage URI for the Async query.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "output-gcs-uri", Usage: "Output Cloud Storage URI for the Async query.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.AsyncQueryReasoningEngineRequest{
								Name:         name,
								InputGcsUri:  cmd.String("input-gcs-uri"),
								OutputGcsUri: cmd.String("output-gcs-uri"),
							}

							op, err := client.AsyncQueryReasoningEngine(ctx, req)
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
						Usage: "create reasoning-engines",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateReasoningEngineRequest{
								Parent: parent,
							}

							op, err := client.CreateReasoningEngine(ctx, req)
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
						Usage: "describe reasoning-engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetReasoningEngineRequest{
								Name: name,
							}

							resp, err := client.GetReasoningEngine(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list reasoning-engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListReasoningEnginesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListReasoningEngines(ctx, req)
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
						Usage: "update reasoning-engines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "reasoning_engine.name" not yet supported.
							reasoning_engine_name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							fmt.Printf("Executing update on %s\n", reasoning_engine_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete reasoning-engines",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, child resources of this reasoning engine will.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteReasoningEngine %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteReasoningEngineRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteReasoningEngine(ctx, req)
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
				Name:  "runs",
				Usage: "Manage runs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-run-id", Usage: "The ID to use for the Tensorboard run, which becomes the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTensorboardRunRequest{
								Parent:           parent,
								TensorboardRunId: cmd.String("tensorboard-run-id"),
							}

							resp, err := client.CreateTensorboardRun(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "batch-create runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchCreateTensorboardRunsRequest{
								Parent: parent,
							}

							resp, err := client.BatchCreateTensorboardRuns(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTensorboardRunRequest{
								Name: name,
							}

							resp, err := client.GetTensorboardRun(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tensorboard_run.name" not yet supported.
							tensorboard_run_name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"))
							fmt.Printf("Executing update on %s\n", tensorboard_run_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the TensorboardRuns that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of TensorboardRuns to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTensorboardRunsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTensorboardRuns(ctx, req)
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
						Usage: "delete runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTensorboardRun %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteTensorboardRunRequest{
								Name: name,
							}

							op, err := client.DeleteTensorboardRun(ctx, req)
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
						Name:  "write",
						Usage: "write runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-run", Usage: "The resource name of the TensorboardRun to write data to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard_run := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"))
							fmt.Printf("Executing write on %s\n", tensorboard_run)
							return nil
						},
					},
				},
			},
			{
				Name:  "saved-queries",
				Usage: "Manage saved-queries resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListSavedQueriesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSavedQueries(ctx, req)
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
						Usage: "delete saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "saved-querie", Usage: "The ID of the saved querie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/savedQueries/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("saved-querie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSavedQuery %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteSavedQueryRequest{
								Name: name,
							}

							op, err := client.DeleteSavedQuery(ctx, req)
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
				Name:  "schedules",
				Usage: "Manage schedules resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create schedules",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateScheduleRequest{
								Parent: parent,
							}

							resp, err := client.CreateSchedule(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete schedules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schedule", Usage: "The ID of the schedule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schedules/%s", cmd.String("project"), cmd.String("location"), cmd.String("schedule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSchedule %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteScheduleRequest{
								Name: name,
							}

							op, err := client.DeleteSchedule(ctx, req)
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
						Usage: "describe schedules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schedule", Usage: "The ID of the schedule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schedules/%s", cmd.String("project"), cmd.String("location"), cmd.String("schedule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetScheduleRequest{
								Name: name,
							}

							resp, err := client.GetSchedule(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list schedules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the Schedules that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListSchedulesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSchedules(ctx, req)
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
						Name:  "pause",
						Usage: "pause schedules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schedule", Usage: "The ID of the schedule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schedules/%s", cmd.String("project"), cmd.String("location"), cmd.String("schedule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute PauseSchedule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.PauseScheduleRequest{
								Name: name,
							}

							if err := client.PauseSchedule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "resume",
						Usage: "resume schedules",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "catch-up", Usage: "Whether to backfill missed runs when the schedule is resumed from.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schedule", Usage: "The ID of the schedule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schedules/%s", cmd.String("project"), cmd.String("location"), cmd.String("schedule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute ResumeSchedule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ResumeScheduleRequest{
								Name:    name,
								CatchUp: cmd.Bool("catch-up"),
							}

							if err := client.ResumeSchedule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update schedules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schedule", Usage: "The ID of the schedule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "schedule.name" not yet supported.
							schedule_name := fmt.Sprintf("projects/%s/locations/%s/schedules/%s", cmd.String("project"), cmd.String("location"), cmd.String("schedule"))
							fmt.Printf("Executing update on %s\n", schedule_name)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
							&cli.StringFlag{Name: "session-id", Usage: "The user defined ID to use for session, which will become the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateSessionRequest{
								Parent:    parent,
								SessionId: cmd.String("session-id"),
							}

							op, err := client.CreateSession(ctx, req)
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
						Usage: "describe sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetSessionRequest{
								Name: name,
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
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of sessions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListSessionsRequest{
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
						Name:  "update",
						Usage: "update sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "session.name" not yet supported.
							session_name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"), cmd.String("session"))
							fmt.Printf("Executing update on %s\n", session_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"), cmd.String("session"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSession %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteSessionRequest{
								Name: name,
							}

							op, err := client.DeleteSession(ctx, req)
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
						Name:  "append-event",
						Usage: "append-event sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reasoning-engine", Usage: "The ID of the reasoning engine.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("reasoning-engine"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.AppendEventRequest{
								Name: name,
							}

							resp, err := client.AppendEvent(ctx, req)
							if err != nil {
								return err
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
				Name:  "slices",
				Usage: "Manage slices resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-import",
						Usage: "batch-import slices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/models/%s/evaluations/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"), cmd.String("evaluation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchImportModelEvaluationSlicesRequest{
								Parent: parent,
							}

							resp, err := client.BatchImportModelEvaluationSlices(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "batch-import",
						Usage: "batch-import slices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "slice", Usage: "The ID of the slice.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/models/%s/evaluations/%s/slices/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"), cmd.String("evaluation"), cmd.String("slice"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.BatchImportEvaluatedAnnotationsRequest{
								Parent: parent,
							}

							resp, err := client.BatchImportEvaluatedAnnotations(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe slices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "slice", Usage: "The ID of the slice.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s/evaluations/%s/slices/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"), cmd.String("evaluation"), cmd.String("slice"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetModelEvaluationSliceRequest{
								Name: name,
							}

							resp, err := client.GetModelEvaluationSlice(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list slices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "evaluation", Usage: "The ID of the evaluation.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The ID of the model.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/models/%s/evaluations/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"), cmd.String("evaluation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListModelEvaluationSlicesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListModelEvaluationSlices(ctx, req)
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
				Name:  "specialist-pools",
				Usage: "Manage specialist-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create specialist-pools",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateSpecialistPoolRequest{
								Parent: parent,
							}

							op, err := client.CreateSpecialistPool(ctx, req)
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
						Usage: "describe specialist-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "specialist-pool", Usage: "The ID of the specialist pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/specialistPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("specialist-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetSpecialistPoolRequest{
								Name: name,
							}

							resp, err := client.GetSpecialistPool(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list specialist-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListSpecialistPoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSpecialistPools(ctx, req)
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
						Usage: "delete specialist-pools",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any specialist managers in this SpecialistPool will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "specialist-pool", Usage: "The ID of the specialist pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/specialistPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("specialist-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSpecialistPool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteSpecialistPoolRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteSpecialistPool(ctx, req)
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
						Usage: "update specialist-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "specialist-pool", Usage: "The ID of the specialist pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "specialist_pool.name" not yet supported.
							specialist_pool_name := fmt.Sprintf("projects/%s/locations/%s/specialistPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("specialist-pool"))
							fmt.Printf("Executing update on %s\n", specialist_pool_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "studies",
				Usage: "Manage studies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create studies",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateStudyRequest{
								Parent: parent,
							}

							resp, err := client.CreateStudy(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe studies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/studies/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetStudyRequest{
								Name: name,
							}

							resp, err := client.GetStudy(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list studies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of studies to return per \"page\" of results.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token to request the next page of results.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListStudiesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListStudies(ctx, req)
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
						Usage: "delete studies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/studies/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteStudy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteStudyRequest{
								Name: name,
							}

							if err := client.DeleteStudy(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "lookup",
						Usage: "lookup studies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "display-name", Usage: "The user-defined display name of the Study.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.LookupStudyRequest{
								Parent:      parent,
								DisplayName: cmd.String("display-name"),
							}

							resp, err := client.LookupStudy(ctx, req)
							if err != nil {
								return err
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
				Name:  "tensorboards",
				Usage: "Manage tensorboards resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tensorboards",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTensorboardRequest{
								Parent: parent,
							}

							op, err := client.CreateTensorboard(ctx, req)
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
						Usage: "describe tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTensorboardRequest{
								Name: name,
							}

							resp, err := client.GetTensorboard(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tensorboard.name" not yet supported.
							tensorboard_name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							fmt.Printf("Executing update on %s\n", tensorboard_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the Tensorboards that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Tensorboards to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTensorboardsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTensorboards(ctx, req)
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
						Usage: "delete tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTensorboard %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteTensorboardRequest{
								Name: name,
							}

							op, err := client.DeleteTensorboard(ctx, req)
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
						Name:  "read-usage",
						Usage: "read-usage tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							fmt.Printf("Executing read-usage on %s\n", tensorboard)
							return nil
						},
					},

					{
						Name:  "read-size",
						Usage: "read-size tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							fmt.Printf("Executing read-size on %s\n", tensorboard)
							return nil
						},
					},

					{
						Name:  "batch-read",
						Usage: "batch-read tensorboards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringSliceFlag{Name: "time-series", Usage: "The resource names of the TensorboardTimeSeries to read data.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"))
							fmt.Printf("Executing batch-read on %s\n", tensorboard)
							return nil
						},
					},
				},
			},
			{
				Name:  "time-series",
				Usage: "Manage time-series resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-time-series-id", Usage: "The user specified unique ID to use for the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTensorboardTimeSeriesRequest{
								Parent:                  parent,
								TensorboardTimeSeriesId: cmd.String("tensorboard-time-series-id"),
							}

							resp, err := client.CreateTensorboardTimeSeries(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "time-serie", Usage: "The ID of the time serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s/timeSeries/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"), cmd.String("time-serie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTensorboardTimeSeriesRequest{
								Name: name,
							}

							resp, err := client.GetTensorboardTimeSeries(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "time-serie", Usage: "The ID of the time serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tensorboard_time_series.name" not yet supported.
							tensorboard_time_series_name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s/timeSeries/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"), cmd.String("time-serie"))
							fmt.Printf("Executing update on %s\n", tensorboard_time_series_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the TensorboardTimeSeries that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of TensorboardTimeSeries to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTensorboardTimeSeriesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTensorboardTimeSeries(ctx, req)
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
						Usage: "delete time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "time-serie", Usage: "The ID of the time serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s/timeSeries/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"), cmd.String("time-serie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTensorboardTimeSeries %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteTensorboardTimeSeriesRequest{
								Name: name,
							}

							op, err := client.DeleteTensorboardTimeSeries(ctx, req)
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
						Name:  "read",
						Usage: "read time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Reads the TensorboardTimeSeries' data that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-data-points", Usage: "The maximum number of TensorboardTimeSeries' data to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-time-series", Usage: "The resource name of the TensorboardTimeSeries to read data from.", Required: true},
							&cli.StringFlag{Name: "time-serie", Usage: "The ID of the time serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard_time_series := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s/timeSeries/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"), cmd.String("time-serie"))
							fmt.Printf("Executing read on %s\n", tensorboard_time_series)
							return nil
						},
					},

					{
						Name:  "read-blob-data",
						Usage: "read-blob-data time-series",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "blob-ids", Usage: "IDs of the blobs to read.", Required: false},
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "time-serie", Usage: "The ID of the time serie.", Required: true},
							&cli.StringFlag{Name: "time-series", Usage: "The resource name of the TensorboardTimeSeries to list Blobs.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							time_series := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s/timeSeries/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"), cmd.String("time-serie"))
							fmt.Printf("Executing read-blob-data on %s\n", time_series)
							return nil
						},
					},

					{
						Name:  "export-tensorboard-time-series",
						Usage: "export-tensorboard-time-series time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "experiment", Usage: "The ID of the experiment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Exports the TensorboardTimeSeries' data that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the TensorboardTimeSeries' data.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of data points to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "tensorboard", Usage: "The ID of the tensorboard.", Required: true},
							&cli.StringFlag{Name: "tensorboard-time-series", Usage: "The resource name of the TensorboardTimeSeries to export data.", Required: true},
							&cli.StringFlag{Name: "time-serie", Usage: "The ID of the time serie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							tensorboard_time_series := fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s/runs/%s/timeSeries/%s", cmd.String("project"), cmd.String("location"), cmd.String("tensorboard"), cmd.String("experiment"), cmd.String("run"), cmd.String("time-serie"))
							fmt.Printf("Executing export-tensorboard-time-series on %s\n", tensorboard_time_series)
							return nil
						},
					},
				},
			},
			{
				Name:  "training-pipelines",
				Usage: "Manage training-pipelines resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create training-pipelines",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTrainingPipelineRequest{
								Parent: parent,
							}

							resp, err := client.CreateTrainingPipeline(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe training-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "training-pipeline", Usage: "The ID of the training pipeline.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/trainingPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("training-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTrainingPipelineRequest{
								Name: name,
							}

							resp, err := client.GetTrainingPipeline(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list training-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTrainingPipelinesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTrainingPipelines(ctx, req)
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
						Usage: "delete training-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "training-pipeline", Usage: "The ID of the training pipeline.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/trainingPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("training-pipeline"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTrainingPipeline %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteTrainingPipelineRequest{
								Name: name,
							}

							op, err := client.DeleteTrainingPipeline(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel training-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "training-pipeline", Usage: "The ID of the training pipeline.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/trainingPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("training-pipeline"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelTrainingPipeline on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelTrainingPipelineRequest{
								Name: name,
							}

							if err := client.CancelTrainingPipeline(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "trials",
				Usage: "Manage trials resources",
				Commands: []*cli.Command{

					{
						Name:  "suggest",
						Usage: "suggest trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "client-id", Usage: "The identifier of the client that is requesting the suggestion.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.IntFlag{Name: "suggestion-count", Usage: "The number of suggestions requested.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/studies/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.SuggestTrialsRequest{
								Parent:          parent,
								SuggestionCount: int32(cmd.Int("suggestion-count")),
								ClientId:        cmd.String("client-id"),
							}

							op, err := client.SuggestTrials(ctx, req)
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
						Usage: "create trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/studies/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTrialRequest{
								Parent: parent,
							}

							resp, err := client.CreateTrial(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.StringFlag{Name: "trial", Usage: "The ID of the trial.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/studies/%s/trials/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"), cmd.String("trial"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTrialRequest{
								Name: name,
							}

							resp, err := client.GetTrial(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The number of Trials to retrieve per \"page\" of results.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token to request the next page of results.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/studies/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTrialsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListTrials(ctx, req)
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
						Name:  "add-trial-measurement",
						Usage: "add-trial-measurement trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.StringFlag{Name: "trial", Usage: "The ID of the trial.", Required: true},
							&cli.StringFlag{Name: "trial-name", Usage: "The name of the trial to add measurement.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							trial_name := fmt.Sprintf("projects/%s/locations/%s/studies/%s/trials/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"), cmd.String("trial"))
							fmt.Printf("Executing add-trial-measurement on %s\n", trial_name)
							return nil
						},
					},

					{
						Name:  "complete",
						Usage: "complete trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "infeasible-reason", Usage: "A human readable reason why the trial was infeasible.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.StringFlag{Name: "trial", Usage: "The ID of the trial.", Required: true},
							&cli.BoolFlag{Name: "trial-infeasible", Usage: "True if the Trial cannot be run with the given Parameter, and.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/studies/%s/trials/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"), cmd.String("trial"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CompleteTrialRequest{
								Name:             name,
								TrialInfeasible:  cmd.Bool("trial-infeasible"),
								InfeasibleReason: cmd.String("infeasible-reason"),
							}

							resp, err := client.CompleteTrial(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.StringFlag{Name: "trial", Usage: "The ID of the trial.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/studies/%s/trials/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"), cmd.String("trial"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTrial on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.DeleteTrialRequest{
								Name: name,
							}

							if err := client.DeleteTrial(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "check-trial-early-stopping-state",
						Usage: "check-trial-early-stopping-state trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.StringFlag{Name: "trial", Usage: "The ID of the trial.", Required: true},
							&cli.StringFlag{Name: "trial-name", Usage: "The Trial's name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							trial_name := fmt.Sprintf("projects/%s/locations/%s/studies/%s/trials/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"), cmd.String("trial"))
							fmt.Printf("Executing check-trial-early-stopping-state on %s\n", trial_name)
							return nil
						},
					},

					{
						Name:  "stop",
						Usage: "stop trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
							&cli.StringFlag{Name: "trial", Usage: "The ID of the trial.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/studies/%s/trials/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"), cmd.String("trial"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.StopTrialRequest{
								Name: name,
							}

							resp, err := client.StopTrial(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "list-optimal-trials",
						Usage: "list-optimal-trials trials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "studie", Usage: "The ID of the studie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/studies/%s", cmd.String("project"), cmd.String("location"), cmd.String("studie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListOptimalTrialsRequest{
								Parent: parent,
							}

							resp, err := client.ListOptimalTrials(ctx, req)
							if err != nil {
								return err
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
				Name:  "tuning-jobs",
				Usage: "Manage tuning-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tuning-jobs",
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
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CreateTuningJobRequest{
								Parent: parent,
							}

							resp, err := client.CreateTuningJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tuning-job", Usage: "The ID of the tuning job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tuningJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tuning-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.GetTuningJobRequest{
								Name: name,
							}

							resp, err := client.GetTuningJob(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.ListTuningJobsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTuningJobs(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel tuning-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tuning-job", Usage: "The ID of the tuning job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tuningJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("tuning-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelTuningJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.CancelTuningJobRequest{
								Name: name,
							}

							if err := client.CancelTuningJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "rebase-tuned-model",
						Usage: "rebase-tuned-model tuning-jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deploy-to-same-endpoint", Usage: "By default, bison to gemini migration will always create new.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := aiplatform.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &aiplatformpb.RebaseTunedModelRequest{
								Parent:               parent,
								DeployToSameEndpoint: cmd.Bool("deploy-to-same-endpoint"),
							}

							op, err := client.RebaseTunedModel(ctx, req)
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
