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

package parallelstore

import (
	parallelstore "cloud.google.com/go/parallelstore/apiv1"
	"cloud.google.com/go/parallelstore/apiv1/parallelstorepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the parallelstore command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "parallelstore",
		Usage: "manage Parallelstore API resources",
		Commands: []*cli.Command{
			{
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list instances",
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
							client, err := parallelstore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &parallelstorepb.ListInstancesRequest{
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
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := parallelstore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &parallelstorepb.GetInstanceRequest{
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

					{
						Name:  "create",
						Usage: "create instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance-id", Usage: "The name of the Parallelstore instance.", Required: true},
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
							client, err := parallelstore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &parallelstorepb.CreateInstanceRequest{
								Parent:     parent,
								InstanceId: cmd.String("instance-id"),
								RequestId:  cmd.String("request-id"),
							}

							op, err := client.CreateInstance(ctx, req)
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
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "instance.name" not yet supported.
							instance_name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", instance_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInstance %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := parallelstore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &parallelstorepb.DeleteInstanceRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInstance(ctx, req)
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
						Usage: "import-data instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-account", Usage: "User-specified service account credentials to be used when.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := parallelstore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &parallelstorepb.ImportDataRequest{
								Name:           name,
								RequestId:      cmd.String("request-id"),
								ServiceAccount: cmd.String("service-account"),
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
						Name:  "export-data",
						Usage: "export-data instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-account", Usage: "User-specified Service Account (SA) credentials to be used when.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := parallelstore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &parallelstorepb.ExportDataRequest{
								Name:           name,
								RequestId:      cmd.String("request-id"),
								ServiceAccount: cmd.String("service-account"),
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
				},
			},
		},
	}
}
