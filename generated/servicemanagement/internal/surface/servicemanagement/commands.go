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

package servicemanagement

import (
	servicemanagement "cloud.google.com/go/servicemanagement/apiv1"
	"cloud.google.com/go/servicemanagement/apiv1/servicemanagementpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the servicemanagement command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "servicemanagement",
		Usage: "manage Service Management API resources",
		Commands: []*cli.Command{
			{
				Name:  "configs",
				Usage: "Manage configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list configs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The max number of items to include in the response list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The token of the page to retrieve.", Required: false},
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing list on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "config_id", Usage: "The ID of the config_id.", Required: true},
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specifies which parts of the Service Config should be returned in the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s/%s", cmd.String("service_name"), cmd.String("config_id"))
							fmt.Printf("Executing describe on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing create on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "submit",
						Usage: "submit configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, this will result in the generation of a.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing submit on %s\n", service_name)
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
				Name:  "rollouts",
				Usage: "Manage rollouts resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Use `filter` to return subset of rollouts.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The max number of items to include in the response list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The token of the page to retrieve.", Required: false},
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing list on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "rollout_id", Usage: "The ID of the rollout_id.", Required: true},
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s/%s", cmd.String("service_name"), cmd.String("rollout_id"))
							fmt.Printf("Executing describe on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing create on %s\n", service_name)
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
						Name:  "list",
						Usage: "list services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "consumer-id", Usage: "Include services consumed by the specified consumer.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The max number of items to include in the response list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token identifying which result to start with; returned by a previous list.", Required: false},
							&cli.StringFlag{Name: "producer-project-id", Usage: "Include services produced by the specified project.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicemanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicemanagementpb.ListServicesRequest{
								ProducerProjectId: cmd.String("producer-project-id"),
								PageSize:          int32(cmd.Int("page-size")),
								PageToken:         cmd.String("page-token"),
								ConsumerId:        cmd.String("consumer-id"),
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
						Name:  "describe",
						Usage: "describe services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing describe on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicemanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicemanagementpb.CreateServiceRequest{}

							op, err := client.CreateService(ctx, req)
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
						Usage: "delete services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing delete on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "undelete",
						Usage: "undelete services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service_name", Usage: "The ID of the service_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							service_name := fmt.Sprintf("%s", cmd.String("service_name"))
							fmt.Printf("Executing undelete on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "generate-config-report",
						Usage: "generate-config-report services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicemanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicemanagementpb.GenerateConfigReportRequest{}

							resp, err := client.GenerateConfigReport(ctx, req)
							if err != nil {
								return err
							}
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
						Usage: "set-iam-policy services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions services",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
		},
	}
}
