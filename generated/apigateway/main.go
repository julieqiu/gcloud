package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	apigateway "cloud.google.com/go/apigateway/apiv1"
	"cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "apigateway",
				Usage: "manage API Gateway API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "apis",
						Usage: "Manage apis resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list apis",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe apis",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/global/apis/%s", cmd.String("project"), cmd.String("api"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.GetApiRequest{Name: name}
									resp, err := client.GetApi(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create apis",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api-id", Usage: "The api id.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "managed-service", Usage: "The managed service.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.CreateApiRequest{Parent: parent}
									req.ApiId = cmd.String("api-id")
									req.Api = &apigatewaypb.Api{
										DisplayName:    cmd.String("display-name"),
										ManagedService: cmd.String("managed-service"),
									}
									op, err := client.CreateApi(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update apis",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "managed-service", Usage: "The managed service.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/global/apis/%s", cmd.String("project"), cmd.String("api"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.UpdateApiRequest{}
									req.Api = &apigatewaypb.Api{
										Name:           name,
										DisplayName:    cmd.String("display-name"),
										ManagedService: cmd.String("managed-service"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("managed-service") {
										paths = append(paths, "managed_service")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateApi(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete apis",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/global/apis/%s", cmd.String("project"), cmd.String("api"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.DeleteApiRequest{Name: name}
									op, err := client.DeleteApi(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &apigatewaypb.ListApiConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListApiConfigs(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
									&cli.StringFlag{Name: "api_config", Usage: "The api_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", cmd.String("project"), cmd.String("api"), cmd.String("api_config"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.GetApiConfigRequest{Name: name}
									resp, err := client.GetApiConfig(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
									&cli.StringFlag{Name: "api-config-id", Usage: "The api config id.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "gateway-service-account", Usage: "The gateway service account.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/global/apis/%s", cmd.String("project"), cmd.String("api"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.CreateApiConfigRequest{Parent: parent}
									req.ApiConfigId = cmd.String("api-config-id")
									req.ApiConfig = &apigatewaypb.ApiConfig{
										DisplayName:           cmd.String("display-name"),
										GatewayServiceAccount: cmd.String("gateway-service-account"),
									}
									op, err := client.CreateApiConfig(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
									&cli.StringFlag{Name: "api_config", Usage: "The api_config.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "gateway-service-account", Usage: "The gateway service account.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", cmd.String("project"), cmd.String("api"), cmd.String("api_config"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.UpdateApiConfigRequest{}
									req.ApiConfig = &apigatewaypb.ApiConfig{
										Name:                  name,
										DisplayName:           cmd.String("display-name"),
										GatewayServiceAccount: cmd.String("gateway-service-account"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("gateway-service-account") {
										paths = append(paths, "gateway_service_account")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateApiConfig(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "api", Usage: "The api.", Required: true},
									&cli.StringFlag{Name: "api_config", Usage: "The api_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/global/apis/%s/configs/%s", cmd.String("project"), cmd.String("api"), cmd.String("api_config"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.DeleteApiConfigRequest{Name: name}
									op, err := client.DeleteApiConfig(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.GetGatewayRequest{Name: name}
									resp, err := client.GetGateway(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway-id", Usage: "The gateway id.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "api-config", Usage: "The api config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.CreateGatewayRequest{Parent: parent}
									req.GatewayId = cmd.String("gateway-id")
									req.Gateway = &apigatewaypb.Gateway{
										DisplayName: cmd.String("display-name"),
										ApiConfig:   cmd.String("api-config"),
									}
									op, err := client.CreateGateway(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "api-config", Usage: "The api config.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.UpdateGatewayRequest{}
									req.Gateway = &apigatewaypb.Gateway{
										Name:        name,
										DisplayName: cmd.String("display-name"),
										ApiConfig:   cmd.String("api-config"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("api-config") {
										paths = append(paths, "api_config")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateGateway(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &apigatewaypb.DeleteGatewayRequest{Name: name}
									op, err := client.DeleteGateway(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := apigateway.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
