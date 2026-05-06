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

package apigateway

import (
	apigateway "cloud.google.com/go/apigateway/apiv1"
	"cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the apigateway command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "apigateway",
		Usage: "manage API Gateway API resources",
		Commands: []*cli.Command{
			{
				Name:  "apis",
				Usage: "Manage apis resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list apis",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by parameters.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.ListApisRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListApis(ctx, req)
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
						Usage: "describe apis",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apis/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.GetApiRequest{
								Name: name,
							}

							resp, err := client.GetApi(ctx, req)
							if err != nil {
								return err
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
						Usage: "create apis",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api-id", Usage: "Identifier to assign to the API.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.CreateApiRequest{
								Parent: parent,
								ApiId:  cmd.String("api-id"),
							}

							op, err := client.CreateApi(ctx, req)
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
						Usage: "update apis",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "api.name" not yet supported.
							api_name := fmt.Sprintf("projects/%s/locations/%s/apis/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"))
							fmt.Printf("Executing update on %s\n", api_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete apis",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apis/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteApi %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.DeleteApiRequest{
								Name: name,
							}

							op, err := client.DeleteApi(ctx, req)
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
				Name:  "configs",
				Usage: "Manage configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by parameters.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apis/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.ListApiConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListApiConfigs(ctx, req)
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
						Usage: "describe configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "config", Usage: "The ID of the config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specifies which fields of the API Config are returned in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apis/%s/configs/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"), cmd.String("config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.GetApiConfigRequest{
								Name: name,
								View: apigatewaypb.GetApiConfigRequest_ConfigView(apigatewaypb.GetApiConfigRequest_ConfigView_value[cmd.String("view")]),
							}

							resp, err := client.GetApiConfig(ctx, req)
							if err != nil {
								return err
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
						Usage: "create configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "api-config-id", Usage: "Identifier to assign to the API Config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/apis/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.CreateApiConfigRequest{
								Parent:      parent,
								ApiConfigId: cmd.String("api-config-id"),
							}

							op, err := client.CreateApiConfig(ctx, req)
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
						Usage: "update configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "config", Usage: "The ID of the config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "api_config.name" not yet supported.
							api_config_name := fmt.Sprintf("projects/%s/locations/%s/apis/%s/configs/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"), cmd.String("config"))
							fmt.Printf("Executing update on %s\n", api_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api", Usage: "The ID of the api.", Required: true},
							&cli.StringFlag{Name: "config", Usage: "The ID of the config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/apis/%s/configs/%s", cmd.String("project"), cmd.String("location"), cmd.String("api"), cmd.String("config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteApiConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.DeleteApiConfigRequest{
								Name: name,
							}

							op, err := client.DeleteApiConfig(ctx, req)
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
				Name:  "gateways",
				Usage: "Manage gateways resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by parameters.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.ListGatewaysRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGateways(ctx, req)
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
						Usage: "describe gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway", Usage: "The ID of the gateway.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.GetGatewayRequest{
								Name: name,
							}

							resp, err := client.GetGateway(ctx, req)
							if err != nil {
								return err
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
						Usage: "create gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-id", Usage: "Identifier to assign to the Gateway.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.CreateGatewayRequest{
								Parent:    parent,
								GatewayId: cmd.String("gateway-id"),
							}

							op, err := client.CreateGateway(ctx, req)
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
						Usage: "update gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway", Usage: "The ID of the gateway.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "gateway.name" not yet supported.
							gateway_name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
							fmt.Printf("Executing update on %s\n", gateway_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway", Usage: "The ID of the gateway.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGateway %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := apigateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &apigatewaypb.DeleteGatewayRequest{
								Name: name,
							}

							op, err := client.DeleteGateway(ctx, req)
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
				},
			},
		},
	}
}
