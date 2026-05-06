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

package cloudbuild

import (
	cloudbuild "cloud.google.com/go/cloudbuild/apiv2"
	"cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudbuild command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudbuild",
		Usage: "manage Cloud Build API resources",
		Commands: []*cli.Command{
			{
				Name:  "connections",
				Usage: "Manage connections resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-id", Usage: "The ID to use for the Connection, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.CreateConnectionRequest{
								Parent:       parent,
								ConnectionId: cmd.String("connection-id"),
							}

							op, err := client.CreateConnection(ctx, req)
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
						Usage: "describe connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.GetConnectionRequest{
								Name: name,
							}

							resp, err := client.GetConnection(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results to return in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page start.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.ListConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListConnections(ctx, req)
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
						Usage: "update connections",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the connection is not found a new connection.", Required: false},
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the connection.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "connection.name" not yet supported.
							connection_name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing update on %s\n", connection_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the connection.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConnection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.DeleteConnectionRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteConnection(ctx, req)
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
						Name:  "fetch-linkable-repositories",
						Usage: "fetch-linkable-repositories connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results to return in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page start.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							connection := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing fetch-linkable-repositories on %s\n", connection)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
				Name:  "repositories",
				Usage: "Manage repositories resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repository-id", Usage: "The ID to use for the repository, which will become the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.CreateRepositoryRequest{
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
						Name:  "batch-create",
						Usage: "batch-create repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.BatchCreateRepositoriesRequest{
								Parent: parent,
							}

							op, err := client.BatchCreateRepositories(ctx, req)
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
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.GetRepositoryRequest{
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
						Name:  "list",
						Usage: "list repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results to return in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page start.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.ListRepositoriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
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
						Name:  "delete",
						Usage: "delete repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the repository.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repositorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRepository %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudbuild.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudbuildpb.DeleteRepositoryRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								ValidateOnly: cmd.Bool("validate-only"),
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
						Name:  "access-read-write-token",
						Usage: "access-read-write-token repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The resource name of the repository in the format.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							repository := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repositorie"))
							fmt.Printf("Executing access-read-write-token on %s\n", repository)
							return nil
						},
					},

					{
						Name:  "access-read-token",
						Usage: "access-read-token repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The resource name of the repository in the format.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							repository := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repositorie"))
							fmt.Printf("Executing access-read-token on %s\n", repository)
							return nil
						},
					},

					{
						Name:  "fetch-git-refs",
						Usage: "fetch-git-refs repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "ref-type", Usage: "Type of refs to fetch.", Required: false},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The resource name of the repository in the format.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							repository := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repositorie"))
							fmt.Printf("Executing fetch-git-refs on %s\n", repository)
							return nil
						},
					},
				},
			},
		},
	}
}
