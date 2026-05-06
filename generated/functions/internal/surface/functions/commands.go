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

package functions

import (
	functions "cloud.google.com/go/functions/apiv2"
	"cloud.google.com/go/functions/apiv2/functionspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudfunctions command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudfunctions",
		Usage: "manage Cloud Functions API resources",
		Commands: []*cli.Command{
			{
				Name:  "functions",
				Usage: "Manage functions resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The optional version of the 1st gen function whose details should.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.GetFunctionRequest{
								Name:     name,
								Revision: cmd.String("revision"),
							}

							resp, err := client.GetFunction(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter for Functions that match the filter expression,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of functions to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.ListFunctionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFunctions(ctx, req)
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
						Usage: "create functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function-id", Usage: "The ID to use for the function, which will become the final component of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.CreateFunctionRequest{
								Parent:     parent,
								FunctionId: cmd.String("function-id"),
							}

							op, err := client.CreateFunction(ctx, req)
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
						Usage: "update functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "function.name" not yet supported.
							function_name := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							fmt.Printf("Executing update on %s\n", function_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFunction %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.DeleteFunctionRequest{
								Name: name,
							}

							op, err := client.DeleteFunction(ctx, req)
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
						Name:  "generate-upload-url",
						Usage: "generate-upload-url functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The function environment the generated upload url will be used for.", Required: false},
							&cli.StringFlag{Name: "kms-key-name", Usage: "Resource name of a KMS crypto key (managed by the user) used to.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.GenerateUploadUrlRequest{
								Parent:      parent,
								KmsKeyName:  cmd.String("kms-key-name"),
								Environment: functionspb.Environment(functionspb.Environment_value[cmd.String("environment")]),
							}

							resp, err := client.GenerateUploadUrl(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "generate-download-url",
						Usage: "generate-download-url functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.GenerateDownloadUrlRequest{
								Name: name,
							}

							resp, err := client.GenerateDownloadUrl(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions functions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "function", Usage: "The ID of the function.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/functions/%s", cmd.String("project"), cmd.String("location"), cmd.String("function"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
				},
			},
			{
				Name:  "runtimes",
				Usage: "Manage runtimes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list runtimes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter for Runtimes that match the filter expression,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := functions.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &functionspb.ListRuntimesRequest{
								Parent: parent,
								Filter: cmd.String("filter"),
							}

							resp, err := client.ListRuntimes(ctx, req)
							if err != nil {
								return err
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
