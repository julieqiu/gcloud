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

package servicedirectory

import (
	servicedirectory "cloud.google.com/go/servicedirectory/apiv1"
	"cloud.google.com/go/servicedirectory/apiv1/servicedirectorypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the servicedirectory command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "servicedirectory",
		Usage: "manage Service Directory API resources",
		Commands: []*cli.Command{
			{
				Name:  "endpoints",
				Usage: "Manage endpoints resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-id", Usage: "The Resource ID must be 1-63 characters long, and comply with.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.CreateEndpointRequest{
								Parent:     parent,
								EndpointId: cmd.String("endpoint-id"),
							}

							resp, err := client.CreateEndpoint(ctx, req)
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
						Usage: "list endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter to list results by.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The order to list results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.ListEndpointsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEndpoints(ctx, req)
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
						Usage: "describe endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"), cmd.String("endpoint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.GetEndpointRequest{
								Name: name,
							}

							resp, err := client.GetEndpoint(ctx, req)
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
						Usage: "update endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "endpoint.name" not yet supported.
							endpoint_name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"), cmd.String("endpoint"))
							fmt.Printf("Executing update on %s\n", endpoint_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"), cmd.String("endpoint"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEndpoint on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.DeleteEndpointRequest{
								Name: name,
							}

							if err := client.DeleteEndpoint(ctx, req); err != nil {
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
				Name:  "namespaces",
				Usage: "Manage namespaces resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace-id", Usage: "The Resource ID must be 1-63 characters long, and comply with.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.CreateNamespaceRequest{
								Parent:      parent,
								NamespaceId: cmd.String("namespace-id"),
							}

							resp, err := client.CreateNamespace(ctx, req)
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
						Usage: "list namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter to list results by.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The order to list results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.ListNamespacesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNamespaces(ctx, req)
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
						Usage: "describe namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.GetNamespaceRequest{
								Name: name,
							}

							resp, err := client.GetNamespace(ctx, req)
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
						Usage: "update namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "namespace.name" not yet supported.
							namespace_name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							fmt.Printf("Executing update on %s\n", namespace_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteNamespace on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.DeleteNamespaceRequest{
								Name: name,
							}

							if err := client.DeleteNamespace(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.GetIamPolicyRequest{
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.SetIamPolicyRequest{
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.TestIamPermissionsRequest{
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
				},
			},
			{
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{

					{
						Name:  "resolve",
						Usage: "resolve services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint-filter", Usage: "The filter applied to the endpoints of the resolved service.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-endpoints", Usage: "The maximum number of endpoints to return.", Required: false},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.ResolveServiceRequest{
								Name:           name,
								MaxEndpoints:   int32(cmd.Int("max-endpoints")),
								EndpointFilter: cmd.String("endpoint-filter"),
							}

							resp, err := client.ResolveService(ctx, req)
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
						Usage: "create services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-id", Usage: "The Resource ID must be 1-63 characters long, and comply with.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.CreateServiceRequest{
								Parent:    parent,
								ServiceId: cmd.String("service-id"),
							}

							resp, err := client.CreateService(ctx, req)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter to list results by.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The order to list results by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.ListServicesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.GetServiceRequest{
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
						Name:  "update",
						Usage: "update services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service.name" not yet supported.
							service_name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"))
							fmt.Printf("Executing update on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("namespace"), cmd.String("service"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteService on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := servicedirectory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicedirectorypb.DeleteServiceRequest{
								Name: name,
							}

							if err := client.DeleteService(ctx, req); err != nil {
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
