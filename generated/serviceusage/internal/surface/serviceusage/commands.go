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

package serviceusage

import (
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	"cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the serviceusage command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "serviceusage",
		Usage: "manage Service Usage API resources",
		Commands: []*cli.Command{
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
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{

					{
						Name:  "enable",
						Usage: "enable services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := serviceusage.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &serviceusagepb.EnableServiceRequest{
								Name: name,
							}

							op, err := client.EnableService(ctx, req)
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
						Usage: "disable services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "check-if-service-has-usage", Usage: "Defines the behavior for checking service usage when disabling a service.", Required: false},
							&cli.BoolFlag{Name: "disable-dependent-services", Usage: "Indicates if services that are enabled and which depend on this service.", Required: false},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := serviceusage.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &serviceusagepb.DisableServiceRequest{
								Name:                     name,
								DisableDependentServices: cmd.Bool("disable-dependent-services"),
								CheckIfServiceHasUsage:   serviceusagepb.DisableServiceRequest_CheckIfServiceHasUsage(serviceusagepb.DisableServiceRequest_CheckIfServiceHasUsage_value[cmd.String("check-if-service-has-usage")]),
							}

							op, err := client.DisableService(ctx, req)
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
						Usage: "describe services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := serviceusage.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &serviceusagepb.GetServiceRequest{
								Name: name,
							}

							resp, err := client.GetService(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "list services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only list services that conform to the given filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested size of the next page of data.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token identifying which result to start with, which is returned by a.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := serviceusage.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &serviceusagepb.ListServicesRequest{
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListServices(ctx, req)
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
						Name:  "batch-enable",
						Usage: "batch-enable services",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "service-ids", Usage: "The identifiers of the services to enable on the project.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := serviceusage.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &serviceusagepb.BatchEnableServicesRequest{
								ServiceIds: cmd.StringSlice("service-ids"),
							}

							op, err := client.BatchEnableServices(ctx, req)
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
						Name:  "batch-get",
						Usage: "batch-get services",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "names", Usage: "Names of the services to retrieve.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := serviceusage.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &serviceusagepb.BatchGetServicesRequest{
								Names: cmd.StringSlice("names"),
							}

							resp, err := client.BatchGetServices(ctx, req)
							if err != nil {
								return err
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
