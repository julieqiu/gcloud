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

package artifactregistry

import (
	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	"cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the artifactregistry command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "artifactregistry",
		Usage: "manage Artifact Registry API resources",
		Commands: []*cli.Command{
			{
				Name:  "apt-artifacts",
				Usage: "Manage apt-artifacts resources",
				Commands: []*cli.Command{

					{
						Name:  "import",
						Usage: "import apt-artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ImportAptArtifactsRequest{
								Parent: parent,
							}

							op, err := client.ImportAptArtifacts(ctx, req)
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
				Name:  "attachments",
				Usage: "Manage attachments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of attachments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListAttachmentsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAttachments(ctx, req)
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
						Usage: "describe attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attachment", Usage: "The ID of the attachment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/attachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("attachment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetAttachmentRequest{
								Name: name,
							}

							resp, err := client.GetAttachment(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "create attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attachment-id", Usage: "The attachment id to use for this attachment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.CreateAttachmentRequest{
								Parent:       parent,
								AttachmentId: cmd.String("attachment-id"),
							}

							op, err := client.CreateAttachment(ctx, req)
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
						Usage: "delete attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attachment", Usage: "The ID of the attachment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/attachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("attachment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAttachment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeleteAttachmentRequest{
								Name: name,
							}

							op, err := client.DeleteAttachment(ctx, req)
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
				Name:  "docker-images",
				Usage: "Manage docker-images resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list docker-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to order the results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of artifacts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListDockerImagesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDockerImages(ctx, req)
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
						Usage: "describe docker-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "docker-image", Usage: "The ID of the docker image.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/dockerImages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("docker-image"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetDockerImageRequest{
								Name: name,
							}

							resp, err := client.GetDockerImage(ctx, req)
							if err != nil {
								return err
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
				Name:  "files",
				Usage: "Manage files resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to order the results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of files to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListFilesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFiles(ctx, req)
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
						Usage: "describe files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "file", Usage: "The ID of the file.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/files/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("file"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetFileRequest{
								Name: name,
							}

							resp, err := client.GetFile(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "file", Usage: "The ID of the file.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/files/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("file"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFile %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeleteFileRequest{
								Name: name,
							}

							op, err := client.DeleteFile(ctx, req)
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
						Usage: "update files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "file", Usage: "The ID of the file.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "file.name" not yet supported.
							file_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/files/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("file"))
							fmt.Printf("Executing update on %s\n", file_name)
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
				},
			},
			{
				Name:  "maven-artifacts",
				Usage: "Manage maven-artifacts resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list maven-artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of artifacts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListMavenArtifactsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMavenArtifacts(ctx, req)
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
						Usage: "describe maven-artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "maven-artifact", Usage: "The ID of the maven artifact.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/mavenArtifacts/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("maven-artifact"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetMavenArtifactRequest{
								Name: name,
							}

							resp, err := client.GetMavenArtifact(ctx, req)
							if err != nil {
								return err
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
				Name:  "npm-packages",
				Usage: "Manage npm-packages resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list npm-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of artifacts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListNpmPackagesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListNpmPackages(ctx, req)
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
						Usage: "describe npm-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "npm-package", Usage: "The ID of the npm package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/npmPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("npm-package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetNpmPackageRequest{
								Name: name,
							}

							resp, err := client.GetNpmPackage(ctx, req)
							if err != nil {
								return err
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
				Name:  "packages",
				Usage: "Manage packages resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to order the results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of packages to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListPackagesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPackages(ctx, req)
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
						Usage: "describe packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetPackageRequest{
								Name: name,
							}

							resp, err := client.GetPackage(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "delete packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePackage %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeletePackageRequest{
								Name: name,
							}

							op, err := client.DeletePackage(ctx, req)
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
						Usage: "update packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "package.name" not yet supported.
							package_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							fmt.Printf("Executing update on %s\n", package_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "project-settings",
				Usage: "Manage project-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe project-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/projectSettings", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetProjectSettingsRequest{
								Name: name,
							}

							resp, err := client.GetProjectSettings(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update project-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "project_settings.name" not yet supported.
							project_settings_name := fmt.Sprintf("projects/%s/projectSettings", cmd.String("project"))
							fmt.Printf("Executing update on %s\n", project_settings_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "python-packages",
				Usage: "Manage python-packages resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list python-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of artifacts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListPythonPackagesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPythonPackages(ctx, req)
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
						Usage: "describe python-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "python-package", Usage: "The ID of the python package.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pythonPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("python-package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetPythonPackageRequest{
								Name: name,
							}

							resp, err := client.GetPythonPackage(ctx, req)
							if err != nil {
								return err
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
				Name:  "repositories",
				Usage: "Manage repositories resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to order the results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of repositories to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListRepositoriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRepositories(ctx, req)
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
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetRepositoryRequest{
								Name: name,
							}

							resp, err := client.GetRepository(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "create repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repository-id", Usage: "The repository id to use for this repository.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.CreateRepositoryRequest{
								Parent:       parent,
								RepositoryId: cmd.String("repository-id"),
							}

							op, err := client.CreateRepository(ctx, req)
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
						Usage: "update repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "repository.name" not yet supported.
							repository_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing update on %s\n", repository_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRepository %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeleteRepositoryRequest{
								Name: name,
							}

							op, err := client.DeleteRepository(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "export-artifact",
						Usage: "export-artifact repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The repository of the artifact to export.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							repository := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing export-artifact on %s\n", repository)
							return nil
						},
					},
				},
			},
			{
				Name:  "rules",
				Usage: "Manage rules resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "rule-id", Usage: "The rule id to use for this repository.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.CreateRuleRequest{
								Parent: parent,
								RuleId: cmd.String("rule-id"),
							}

							resp, err := client.CreateRule(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of rules to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRules(ctx, req)
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
						Usage: "describe rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetRuleRequest{
								Name: name,
							}

							resp, err := client.GetRule(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "rule.name" not yet supported.
							rule_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("rule"))
							fmt.Printf("Executing update on %s\n", rule_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteRule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeleteRuleRequest{
								Name: name,
							}

							if err := client.DeleteRule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "tags",
				Usage: "Manage tags resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of tags to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListTagsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTags(ctx, req)
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
						Usage: "describe tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The ID of the tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"), cmd.String("tag"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetTagRequest{
								Name: name,
							}

							resp, err := client.GetTag(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "create tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "tag-id", Usage: "The tag id to use for this repository.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.CreateTagRequest{
								Parent: parent,
								TagId:  cmd.String("tag-id"),
							}

							resp, err := client.CreateTag(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The ID of the tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tag.name" not yet supported.
							tag_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"), cmd.String("tag"))
							fmt.Printf("Executing update on %s\n", tag_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The ID of the tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"), cmd.String("tag"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTag on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeleteTagRequest{
								Name: name,
							}

							if err := client.DeleteTag(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to order the results by.", Required: false},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of versions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous list request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view that should be returned in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ListVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								View:      artifactregistrypb.VersionView(artifactregistrypb.VersionView_value[cmd.String("view")]),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view that should be returned in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetVersionRequest{
								Name: name,
								View: artifactregistrypb.VersionView(artifactregistrypb.VersionView_value[cmd.String("view")]),
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
						Name:  "delete",
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "By default, a version that is tagged may not be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.DeleteVersionRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteVersion(ctx, req)
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
						Usage: "batch-delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the versions to delete.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request is performed without deleting data, following AIP-163.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("BatchDeleteVersions %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.BatchDeleteVersionsRequest{
								Parent:       parent,
								Names:        cmd.StringSlice("names"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.BatchDeleteVersions(ctx, req)
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
						Usage: "update versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "package", Usage: "The ID of the package.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "version.name" not yet supported.
							version_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/packages/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("package"), cmd.String("version"))
							fmt.Printf("Executing update on %s\n", version_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "vpcsc-config",
				Usage: "Manage vpcsc-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe vpcsc-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcscConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.GetVPCSCConfigRequest{
								Name: name,
							}

							resp, err := client.GetVPCSCConfig(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "update vpcsc-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "vpcsc_config.name" not yet supported.
							vpcsc_config_name := fmt.Sprintf("projects/%s/locations/%s/vpcscConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", vpcsc_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "yum-artifacts",
				Usage: "Manage yum-artifacts resources",
				Commands: []*cli.Command{

					{
						Name:  "import",
						Usage: "import yum-artifacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := artifactregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &artifactregistrypb.ImportYumArtifactsRequest{
								Parent: parent,
							}

							op, err := client.ImportYumArtifacts(ctx, req)
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
