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

package networkservices

import (
	networkservices "cloud.google.com/go/networkservices/apiv1"
	"cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the networkservices command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "networkservices",
		Usage: "manage Network Services API resources",
		Commands: []*cli.Command{
			{
				Name:  "authz-extensions",
				Usage: "Manage authz-extensions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list authz-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint about how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results that the server returns.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListAuthzExtensionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAuthzExtensions(ctx, req)
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
						Usage: "describe authz-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-extension", Usage: "The ID of the authz extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz-extension"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetAuthzExtensionRequest{
								Name: name,
							}

							resp, err := client.GetAuthzExtension(ctx, req)
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
						Usage: "create authz-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-extension-id", Usage: "User-provided ID of the `AuthzExtension` resource to be.", Required: true},
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
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateAuthzExtensionRequest{
								Parent:           parent,
								AuthzExtensionId: cmd.String("authz-extension-id"),
								RequestId:        cmd.String("request-id"),
							}

							op, err := client.CreateAuthzExtension(ctx, req)
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
						Usage: "update authz-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-extension", Usage: "The ID of the authz extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "authz_extension.name" not yet supported.
							authz_extension_name := fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz-extension"))
							fmt.Printf("Executing update on %s\n", authz_extension_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete authz-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-extension", Usage: "The ID of the authz extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz-extension"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAuthzExtension %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteAuthzExtensionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteAuthzExtension(ctx, req)
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
				Name:  "edge-cache-keysets",
				Usage: "Manage edge-cache-keysets resources",
				Commands: []*cli.Command{

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy edge-cache-keysets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-cache-keyset", Usage: "The ID of the edge cache keyset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/edgeCacheKeysets/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-cache-keyset"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe edge-cache-keysets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-cache-keyset", Usage: "The ID of the edge cache keyset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/edgeCacheKeysets/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-cache-keyset"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions edge-cache-keysets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-cache-keyset", Usage: "The ID of the edge cache keyset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/edgeCacheKeysets/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-cache-keyset"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy edge-cache-keysets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-cache-keyset", Usage: "The ID of the edge cache keyset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/edgeCacheKeysets/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-cache-keyset"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe edge-cache-keysets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-cache-keyset", Usage: "The ID of the edge cache keyset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/edgeCacheKeysets/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-cache-keyset"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions edge-cache-keysets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-cache-keyset", Usage: "The ID of the edge cache keyset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/edgeCacheKeysets/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-cache-keyset"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "endpoint-policies",
				Usage: "Manage endpoint-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list endpoint-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of EndpointPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListEndpointPoliciesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListEndpointPoliciesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListEndpointPolicies(ctx, req)
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
						Usage: "describe endpoint-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-policie", Usage: "The ID of the endpoint policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/endpointPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetEndpointPolicyRequest{
								Name: name,
							}

							resp, err := client.GetEndpointPolicy(ctx, req)
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
						Usage: "create endpoint-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-policy-id", Usage: "Short name of the EndpointPolicy resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateEndpointPolicyRequest{
								Parent:           parent,
								EndpointPolicyId: cmd.String("endpoint-policy-id"),
							}

							op, err := client.CreateEndpointPolicy(ctx, req)
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
						Usage: "update endpoint-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-policie", Usage: "The ID of the endpoint policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "endpoint_policy.name" not yet supported.
							endpoint_policy_name := fmt.Sprintf("projects/%s/locations/%s/endpointPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint-policie"))
							fmt.Printf("Executing update on %s\n", endpoint_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete endpoint-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-policie", Usage: "The ID of the endpoint policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/endpointPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEndpointPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteEndpointPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteEndpointPolicy(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Gateways to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListGatewaysResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListGatewaysRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetGatewayRequest{
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
							&cli.StringFlag{Name: "gateway-id", Usage: "Short name of the Gateway resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateGatewayRequest{
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
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteGatewayRequest{
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
				Name:  "grpc-routes",
				Usage: "Manage grpc-routes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list grpc-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of GrpcRoutes to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListGrpcRoutesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListGrpcRoutesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListGrpcRoutes(ctx, req)
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
						Usage: "describe grpc-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "grpc-route", Usage: "The ID of the grpc route.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("grpc-route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetGrpcRouteRequest{
								Name: name,
							}

							resp, err := client.GetGrpcRoute(ctx, req)
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
						Usage: "create grpc-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "grpc-route-id", Usage: "Short name of the GrpcRoute resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateGrpcRouteRequest{
								Parent:      parent,
								GrpcRouteId: cmd.String("grpc-route-id"),
							}

							op, err := client.CreateGrpcRoute(ctx, req)
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
						Usage: "update grpc-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "grpc-route", Usage: "The ID of the grpc route.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "grpc_route.name" not yet supported.
							grpc_route_name := fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("grpc-route"))
							fmt.Printf("Executing update on %s\n", grpc_route_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete grpc-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "grpc-route", Usage: "The ID of the grpc route.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("grpc-route"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGrpcRoute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteGrpcRouteRequest{
								Name: name,
							}

							op, err := client.DeleteGrpcRoute(ctx, req)
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
				Name:  "http-routes",
				Usage: "Manage http-routes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list http-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of HttpRoutes to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListHttpRoutesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListHttpRoutesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListHttpRoutes(ctx, req)
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
						Usage: "describe http-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "http-route", Usage: "The ID of the http route.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("http-route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetHttpRouteRequest{
								Name: name,
							}

							resp, err := client.GetHttpRoute(ctx, req)
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
						Usage: "create http-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "http-route-id", Usage: "Short name of the HttpRoute resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateHttpRouteRequest{
								Parent:      parent,
								HttpRouteId: cmd.String("http-route-id"),
							}

							op, err := client.CreateHttpRoute(ctx, req)
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
						Usage: "update http-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "http-route", Usage: "The ID of the http route.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "http_route.name" not yet supported.
							http_route_name := fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("http-route"))
							fmt.Printf("Executing update on %s\n", http_route_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete http-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "http-route", Usage: "The ID of the http route.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("http-route"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteHttpRoute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteHttpRouteRequest{
								Name: name,
							}

							op, err := client.DeleteHttpRoute(ctx, req)
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
				Name:  "lb-edge-extensions",
				Usage: "Manage lb-edge-extensions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list lb-edge-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint about how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results that the server returns.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListLbEdgeExtensionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListLbEdgeExtensions(ctx, req)
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
						Usage: "describe lb-edge-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-edge-extension", Usage: "The ID of the lb edge extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-edge-extension"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetLbEdgeExtensionRequest{
								Name: name,
							}

							resp, err := client.GetLbEdgeExtension(ctx, req)
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
						Usage: "create lb-edge-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-edge-extension-id", Usage: "User-provided ID of the `LbEdgeExtension` resource to be created.", Required: true},
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
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateLbEdgeExtensionRequest{
								Parent:            parent,
								LbEdgeExtensionId: cmd.String("lb-edge-extension-id"),
								RequestId:         cmd.String("request-id"),
							}

							op, err := client.CreateLbEdgeExtension(ctx, req)
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
						Usage: "update lb-edge-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-edge-extension", Usage: "The ID of the lb edge extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "lb_edge_extension.name" not yet supported.
							lb_edge_extension_name := fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-edge-extension"))
							fmt.Printf("Executing update on %s\n", lb_edge_extension_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete lb-edge-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-edge-extension", Usage: "The ID of the lb edge extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-edge-extension"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteLbEdgeExtension %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteLbEdgeExtensionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteLbEdgeExtension(ctx, req)
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
				Name:  "lb-route-extensions",
				Usage: "Manage lb-route-extensions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list lb-route-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint about how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results that the server returns.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListLbRouteExtensionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListLbRouteExtensions(ctx, req)
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
						Usage: "describe lb-route-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-route-extension", Usage: "The ID of the lb route extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-route-extension"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetLbRouteExtensionRequest{
								Name: name,
							}

							resp, err := client.GetLbRouteExtension(ctx, req)
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
						Usage: "create lb-route-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-route-extension-id", Usage: "User-provided ID of the `LbRouteExtension` resource to be.", Required: true},
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
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateLbRouteExtensionRequest{
								Parent:             parent,
								LbRouteExtensionId: cmd.String("lb-route-extension-id"),
								RequestId:          cmd.String("request-id"),
							}

							op, err := client.CreateLbRouteExtension(ctx, req)
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
						Usage: "update lb-route-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-route-extension", Usage: "The ID of the lb route extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "lb_route_extension.name" not yet supported.
							lb_route_extension_name := fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-route-extension"))
							fmt.Printf("Executing update on %s\n", lb_route_extension_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete lb-route-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-route-extension", Usage: "The ID of the lb route extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-route-extension"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteLbRouteExtension %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteLbRouteExtensionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteLbRouteExtension(ctx, req)
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
				Name:  "lb-traffic-extensions",
				Usage: "Manage lb-traffic-extensions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list lb-traffic-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint about how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results that the server returns.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListLbTrafficExtensionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListLbTrafficExtensions(ctx, req)
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
						Usage: "describe lb-traffic-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-traffic-extension", Usage: "The ID of the lb traffic extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-traffic-extension"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetLbTrafficExtensionRequest{
								Name: name,
							}

							resp, err := client.GetLbTrafficExtension(ctx, req)
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
						Usage: "create lb-traffic-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-traffic-extension-id", Usage: "User-provided ID of the `LbTrafficExtension` resource to be.", Required: true},
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
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateLbTrafficExtensionRequest{
								Parent:               parent,
								LbTrafficExtensionId: cmd.String("lb-traffic-extension-id"),
								RequestId:            cmd.String("request-id"),
							}

							op, err := client.CreateLbTrafficExtension(ctx, req)
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
						Usage: "update lb-traffic-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-traffic-extension", Usage: "The ID of the lb traffic extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "lb_traffic_extension.name" not yet supported.
							lb_traffic_extension_name := fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-traffic-extension"))
							fmt.Printf("Executing update on %s\n", lb_traffic_extension_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete lb-traffic-extensions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lb-traffic-extension", Usage: "The ID of the lb traffic extension.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb-traffic-extension"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteLbTrafficExtension %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteLbTrafficExtensionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteLbTrafficExtension(ctx, req)
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
				Name:  "meshes",
				Usage: "Manage meshes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list meshes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Meshes to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListMeshesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListMeshesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListMeshes(ctx, req)
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
						Usage: "describe meshes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "meshe", Usage: "The ID of the meshe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("meshe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetMeshRequest{
								Name: name,
							}

							resp, err := client.GetMesh(ctx, req)
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
						Usage: "create meshes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mesh-id", Usage: "Short name of the Mesh resource to be created.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateMeshRequest{
								Parent: parent,
								MeshId: cmd.String("mesh-id"),
							}

							op, err := client.CreateMesh(ctx, req)
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
						Usage: "update meshes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "meshe", Usage: "The ID of the meshe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "mesh.name" not yet supported.
							mesh_name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("meshe"))
							fmt.Printf("Executing update on %s\n", mesh_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete meshes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "meshe", Usage: "The ID of the meshe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("meshe"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMesh %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteMeshRequest{
								Name: name,
							}

							op, err := client.DeleteMesh(ctx, req)
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
			{
				Name:  "route-views",
				Usage: "Manage route-views resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe route-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway", Usage: "The ID of the gateway.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route-view", Usage: "The ID of the route view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s/routeViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"), cmd.String("route-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetGatewayRouteViewRequest{
								Name: name,
							}

							resp, err := client.GetGatewayRouteView(ctx, req)
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
						Usage: "describe route-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "meshe", Usage: "The ID of the meshe.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route-view", Usage: "The ID of the route view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s/routeViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("meshe"), cmd.String("route-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetMeshRouteViewRequest{
								Name: name,
							}

							resp, err := client.GetMeshRouteView(ctx, req)
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
						Usage: "list route-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway", Usage: "The ID of the gateway.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of GatewayRouteViews to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListGatewayRouteViewsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListGatewayRouteViewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGatewayRouteViews(ctx, req)
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
						Usage: "list route-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "meshe", Usage: "The ID of the meshe.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of MeshRouteViews to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListMeshRouteViewsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("meshe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListMeshRouteViewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMeshRouteViews(ctx, req)
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
				Name:  "service-bindings",
				Usage: "Manage service-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list service-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of ServiceBindings to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListServiceBindingsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListServiceBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceBindings(ctx, req)
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
						Usage: "describe service-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-binding", Usage: "The ID of the service binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetServiceBindingRequest{
								Name: name,
							}

							resp, err := client.GetServiceBinding(ctx, req)
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
						Usage: "create service-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-binding-id", Usage: "Short name of the ServiceBinding resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateServiceBindingRequest{
								Parent:           parent,
								ServiceBindingId: cmd.String("service-binding-id"),
							}

							op, err := client.CreateServiceBinding(ctx, req)
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
						Usage: "update service-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-binding", Usage: "The ID of the service binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_binding.name" not yet supported.
							service_binding_name := fmt.Sprintf("projects/%s/locations/%s/serviceBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-binding"))
							fmt.Printf("Executing update on %s\n", service_binding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-binding", Usage: "The ID of the service binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-binding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServiceBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteServiceBindingRequest{
								Name: name,
							}

							op, err := client.DeleteServiceBinding(ctx, req)
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
				Name:  "service-lb-policies",
				Usage: "Manage service-lb-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list service-lb-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of ServiceLbPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListServiceLbPoliciesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListServiceLbPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceLbPolicies(ctx, req)
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
						Usage: "describe service-lb-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-lb-policie", Usage: "The ID of the service lb policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-lb-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetServiceLbPolicyRequest{
								Name: name,
							}

							resp, err := client.GetServiceLbPolicy(ctx, req)
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
						Usage: "create service-lb-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-lb-policy-id", Usage: "Short name of the ServiceLbPolicy resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateServiceLbPolicyRequest{
								Parent:            parent,
								ServiceLbPolicyId: cmd.String("service-lb-policy-id"),
							}

							op, err := client.CreateServiceLbPolicy(ctx, req)
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
						Usage: "update service-lb-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-lb-policie", Usage: "The ID of the service lb policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_lb_policy.name" not yet supported.
							service_lb_policy_name := fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-lb-policie"))
							fmt.Printf("Executing update on %s\n", service_lb_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-lb-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-lb-policie", Usage: "The ID of the service lb policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-lb-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServiceLbPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteServiceLbPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteServiceLbPolicy(ctx, req)
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
				Name:  "tcp-routes",
				Usage: "Manage tcp-routes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list tcp-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of TcpRoutes to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListTcpRoutesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListTcpRoutesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListTcpRoutes(ctx, req)
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
						Usage: "describe tcp-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tcp-route", Usage: "The ID of the tcp route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tcp-route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetTcpRouteRequest{
								Name: name,
							}

							resp, err := client.GetTcpRoute(ctx, req)
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
						Usage: "create tcp-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tcp-route-id", Usage: "Short name of the TcpRoute resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateTcpRouteRequest{
								Parent:     parent,
								TcpRouteId: cmd.String("tcp-route-id"),
							}

							op, err := client.CreateTcpRoute(ctx, req)
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
						Usage: "update tcp-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tcp-route", Usage: "The ID of the tcp route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tcp_route.name" not yet supported.
							tcp_route_name := fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tcp-route"))
							fmt.Printf("Executing update on %s\n", tcp_route_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tcp-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tcp-route", Usage: "The ID of the tcp route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tcp-route"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTcpRoute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteTcpRouteRequest{
								Name: name,
							}

							op, err := client.DeleteTcpRoute(ctx, req)
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
				Name:  "tls-routes",
				Usage: "Manage tls-routes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list tls-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of TlsRoutes to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListTlsRoutesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListTlsRoutesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListTlsRoutes(ctx, req)
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
						Usage: "describe tls-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-route", Usage: "The ID of the tls route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls-route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetTlsRouteRequest{
								Name: name,
							}

							resp, err := client.GetTlsRoute(ctx, req)
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
						Usage: "create tls-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-route-id", Usage: "Short name of the TlsRoute resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateTlsRouteRequest{
								Parent:     parent,
								TlsRouteId: cmd.String("tls-route-id"),
							}

							op, err := client.CreateTlsRoute(ctx, req)
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
						Usage: "update tls-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-route", Usage: "The ID of the tls route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tls_route.name" not yet supported.
							tls_route_name := fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls-route"))
							fmt.Printf("Executing update on %s\n", tls_route_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tls-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-route", Usage: "The ID of the tls route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls-route"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTlsRoute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteTlsRouteRequest{
								Name: name,
							}

							op, err := client.DeleteTlsRoute(ctx, req)
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
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of `WasmPluginVersion` resources to return per.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListWasmPluginVersionsResponse` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListWasmPluginVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWasmPluginVersions(ctx, req)
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetWasmPluginVersionRequest{
								Name: name,
							}

							resp, err := client.GetWasmPluginVersion(ctx, req)
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
						Usage: "create versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin-version-id", Usage: "User-provided ID of the `WasmPluginVersion` resource to be.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateWasmPluginVersionRequest{
								Parent:              parent,
								WasmPluginVersionId: cmd.String("wasm-plugin-version-id"),
							}

							op, err := client.CreateWasmPluginVersion(ctx, req)
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
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteWasmPluginVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteWasmPluginVersionRequest{
								Name: name,
							}

							op, err := client.DeleteWasmPluginVersion(ctx, req)
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
				Name:  "wasm-plugins",
				Usage: "Manage wasm-plugins resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list wasm-plugins",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of `WasmPlugin` resources to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListWasmPluginsResponse` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.ListWasmPluginsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWasmPlugins(ctx, req)
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
						Usage: "describe wasm-plugins",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Determines how much data must be returned in the response.", Required: false},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.GetWasmPluginRequest{
								Name: name,
								View: networkservicespb.WasmPluginView(networkservicespb.WasmPluginView_value[cmd.String("view")]),
							}

							resp, err := client.GetWasmPlugin(ctx, req)
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
						Usage: "create wasm-plugins",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin-id", Usage: "User-provided ID of the `WasmPlugin` resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.CreateWasmPluginRequest{
								Parent:       parent,
								WasmPluginId: cmd.String("wasm-plugin-id"),
							}

							op, err := client.CreateWasmPlugin(ctx, req)
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
						Usage: "update wasm-plugins",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "wasm_plugin.name" not yet supported.
							wasm_plugin_name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"))
							fmt.Printf("Executing update on %s\n", wasm_plugin_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete wasm-plugins",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wasm-plugin", Usage: "The ID of the wasm plugin.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm-plugin"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteWasmPlugin %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkservices.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkservicespb.DeleteWasmPluginRequest{
								Name: name,
							}

							op, err := client.DeleteWasmPlugin(ctx, req)
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
		},
	}
}
