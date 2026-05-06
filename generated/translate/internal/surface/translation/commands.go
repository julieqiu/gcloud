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

package translation

import (
	translation "cloud.google.com/go/translation/apiv3"
	"cloud.google.com/go/translation/apiv3/translationpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the translate command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "translate",
		Usage: "manage Cloud Translation API resources",
		Commands: []*cli.Command{
			{
				Name:  "adaptive-mt-datasets",
				Usage: "Manage adaptive-mt-datasets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create adaptive-mt-datasets",
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.CreateAdaptiveMtDatasetRequest{
								Parent: parent,
							}

							resp, err := client.CreateAdaptiveMtDataset(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAdaptiveMtDataset on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DeleteAdaptiveMtDatasetRequest{
								Name: name,
							}

							if err := client.DeleteAdaptiveMtDataset(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetAdaptiveMtDatasetRequest{
								Name: name,
							}

							resp, err := client.GetAdaptiveMtDataset(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListAdaptiveMtDatasetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListAdaptiveMtDatasets(ctx, req)
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
						Name:  "import-adaptive-mt-file",
						Usage: "import-adaptive-mt-file adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ImportAdaptiveMtFileRequest{
								Parent: parent,
							}

							resp, err := client.ImportAdaptiveMtFile(ctx, req)
							if err != nil {
								return err
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
				Name:  "adaptive-mt-files",
				Usage: "Manage adaptive-mt-files resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe adaptive-mt-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "adaptive-mt-file", Usage: "The ID of the adaptive mt file.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s/adaptiveMtFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"), cmd.String("adaptive-mt-file"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetAdaptiveMtFileRequest{
								Name: name,
							}

							resp, err := client.GetAdaptiveMtFile(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete adaptive-mt-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "adaptive-mt-file", Usage: "The ID of the adaptive mt file.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s/adaptiveMtFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"), cmd.String("adaptive-mt-file"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAdaptiveMtFile on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DeleteAdaptiveMtFileRequest{
								Name: name,
							}

							if err := client.DeleteAdaptiveMtFile(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list adaptive-mt-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListAdaptiveMtFilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAdaptiveMtFiles(ctx, req)
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
				Name:  "adaptive-mt-sentences",
				Usage: "Manage adaptive-mt-sentences resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list adaptive-mt-sentences",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "adaptive-mt-dataset", Usage: "The ID of the adaptive mt dataset.", Required: true},
							&cli.StringFlag{Name: "adaptive-mt-file", Usage: "The ID of the adaptive mt file.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s/adaptiveMtFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("adaptive-mt-dataset"), cmd.String("adaptive-mt-file"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListAdaptiveMtSentencesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAdaptiveMtSentences(ctx, req)
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.CreateDatasetRequest{
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetDatasetRequest{
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
						Name:  "list",
						Usage: "list datasets",
						Flags: []cli.Flag{
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListDatasetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DeleteDatasetRequest{
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
						Name:  "import-data",
						Usage: "import-data datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							dataset := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing import-data on %s\n", dataset)
							return nil
						},
					},

					{
						Name:  "export-data",
						Usage: "export-data datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							dataset := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing export-data on %s\n", dataset)
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
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the examples that will be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results for the server to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListExamplesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
				},
			},
			{
				Name:  "glossaries",
				Usage: "Manage glossaries resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create glossaries",
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.CreateGlossaryRequest{
								Parent: parent,
							}

							op, err := client.CreateGlossary(ctx, req)
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
						Usage: "update glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "glossary.name" not yet supported.
							glossary_name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							fmt.Printf("Executing update on %s\n", glossary_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter specifying constraints of a list operation.", Required: false},
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListGlossariesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListGlossaries(ctx, req)
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
						Usage: "describe glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetGlossaryRequest{
								Name: name,
							}

							resp, err := client.GetGlossary(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DeleteGlossaryRequest{
								Name: name,
							}

							op, err := client.DeleteGlossary(ctx, req)
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
				Name:  "glossary-entries",
				Usage: "Manage glossary-entries resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "glossary-entrie", Usage: "The ID of the glossary entrie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/glossaryEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("glossary-entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetGlossaryEntryRequest{
								Name: name,
							}

							resp, err := client.GetGlossaryEntry(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListGlossaryEntriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGlossaryEntries(ctx, req)
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
						Usage: "create glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.CreateGlossaryEntryRequest{
								Parent: parent,
							}

							resp, err := client.CreateGlossaryEntry(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "glossary-entrie", Usage: "The ID of the glossary entrie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "glossary_entry.name" not yet supported.
							glossary_entry_name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/glossaryEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("glossary-entrie"))
							fmt.Printf("Executing update on %s\n", glossary_entry_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "glossary-entrie", Usage: "The ID of the glossary entrie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/glossaryEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("glossary-entrie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGlossaryEntry on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DeleteGlossaryEntryRequest{
								Name: name,
							}

							if err := client.DeleteGlossaryEntry(ctx, req); err != nil {
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
						Name:  "translate-text",
						Usage: "translate-text locations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "contents", Usage: "The content of the input in string format.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mime-type", Usage: "The format of the source text, for example, \"text/html\",.", Required: false},
							&cli.StringFlag{Name: "model", Usage: "The `model` type requested for this translation.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-language-code", Usage: "The ISO-639 language code of the input text if.", Required: false},
							&cli.StringFlag{Name: "target-language-code", Usage: "The ISO-639 language code to use for translation of the input.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.TranslateTextRequest{
								Parent:             parent,
								Contents:           cmd.StringSlice("contents"),
								MimeType:           cmd.String("mime-type"),
								SourceLanguageCode: cmd.String("source-language-code"),
								TargetLanguageCode: cmd.String("target-language-code"),
								Model:              cmd.String("model"),
							}

							resp, err := client.TranslateText(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "romanize-text",
						Usage: "romanize-text locations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "contents", Usage: "The content of the input in string format.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-language-code", Usage: "The ISO-639 language code of the input text if.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.RomanizeTextRequest{
								Parent:             parent,
								Contents:           cmd.StringSlice("contents"),
								SourceLanguageCode: cmd.String("source-language-code"),
							}

							resp, err := client.RomanizeText(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "detect-language",
						Usage: "detect-language locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mime-type", Usage: "The format of the source text, for example, \"text/html\",.", Required: false},
							&cli.StringFlag{Name: "model", Usage: "The language detection model to be used.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DetectLanguageRequest{
								Parent:   parent,
								Model:    cmd.String("model"),
								MimeType: cmd.String("mime-type"),
							}

							resp, err := client.DetectLanguage(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "translate-document",
						Usage: "translate-document locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customized-attribution", Usage: "This flag is to support user customized attribution.", Required: false},
							&cli.BoolFlag{Name: "enable-rotation-correction", Usage: "If true, enable auto rotation correction in DVS.", Required: false},
							&cli.BoolFlag{Name: "enable-shadow-removal-native-pdf", Usage: "If true, use the text removal server to remove the shadow text on.", Required: false},
							&cli.BoolFlag{Name: "is-translate-native-pdf-only", Usage: "is_translate_native_pdf_only field for external customers.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The `model` type requested for this translation.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-language-code", Usage: "The ISO-639 language code of the input document if known, for.", Required: false},
							&cli.StringFlag{Name: "target-language-code", Usage: "The ISO-639 language code to use for translation of the input.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.TranslateDocumentRequest{
								Parent:                       parent,
								SourceLanguageCode:           cmd.String("source-language-code"),
								TargetLanguageCode:           cmd.String("target-language-code"),
								Model:                        cmd.String("model"),
								CustomizedAttribution:        cmd.String("customized-attribution"),
								IsTranslateNativePdfOnly:     cmd.Bool("is-translate-native-pdf-only"),
								EnableShadowRemovalNativePdf: cmd.Bool("enable-shadow-removal-native-pdf"),
								EnableRotationCorrection:     cmd.Bool("enable-rotation-correction"),
							}

							resp, err := client.TranslateDocument(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "batch-translate-text",
						Usage: "batch-translate-text locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-language-code", Usage: "Source language code.", Required: true},
							&cli.StringSliceFlag{Name: "target-language-codes", Usage: "Specify up to 10 language codes here.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.BatchTranslateTextRequest{
								Parent:              parent,
								SourceLanguageCode:  cmd.String("source-language-code"),
								TargetLanguageCodes: cmd.StringSlice("target-language-codes"),
							}

							op, err := client.BatchTranslateText(ctx, req)
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
						Name:  "batch-translate-document",
						Usage: "batch-translate-document locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customized-attribution", Usage: "This flag is to support user customized attribution.", Required: false},
							&cli.BoolFlag{Name: "enable-rotation-correction", Usage: "If true, enable auto rotation correction in DVS.", Required: false},
							&cli.BoolFlag{Name: "enable-shadow-removal-native-pdf", Usage: "If true, use the text removal server to remove the shadow text on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-language-code", Usage: "The ISO-639 language code of the input document if known, for.", Required: true},
							&cli.StringSliceFlag{Name: "target-language-codes", Usage: "The ISO-639 language code to use for translation of the input.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.BatchTranslateDocumentRequest{
								Parent:                       parent,
								SourceLanguageCode:           cmd.String("source-language-code"),
								TargetLanguageCodes:          cmd.StringSlice("target-language-codes"),
								CustomizedAttribution:        cmd.String("customized-attribution"),
								EnableShadowRemovalNativePdf: cmd.Bool("enable-shadow-removal-native-pdf"),
								EnableRotationCorrection:     cmd.Bool("enable-rotation-correction"),
							}

							op, err := client.BatchTranslateDocument(ctx, req)
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
						Name:  "adaptive-mt-translate",
						Usage: "adaptive-mt-translate locations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "content", Usage: "The content of the input in string format.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The resource name for the dataset to use for adaptive MT.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mime-type", Usage: "The format of the source text.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.AdaptiveMtTranslateRequest{
								Parent:   parent,
								Dataset:  cmd.String("dataset"),
								Content:  cmd.StringSlice("content"),
								MimeType: cmd.String("mime-type"),
							}

							resp, err := client.AdaptiveMtTranslate(ctx, req)
							if err != nil {
								return err
							}
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
				Name:  "models",
				Usage: "Manage models resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create models",
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.CreateModelRequest{
								Parent: parent,
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
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the models that will be returned.", Required: false},
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.ListModelsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetModelRequest{
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
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.DeleteModelRequest{
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
				},
			},
			{
				Name:  "supported-languages",
				Usage: "Manage supported-languages resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe supported-languages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "display-language-code", Usage: "The language to use to return localized, human readable names.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "Get supported languages of this model.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := translation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &translationpb.GetSupportedLanguagesRequest{
								Parent:              parent,
								DisplayLanguageCode: cmd.String("display-language-code"),
								Model:               cmd.String("model"),
							}

							resp, err := client.GetSupportedLanguages(ctx, req)
							if err != nil {
								return err
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
