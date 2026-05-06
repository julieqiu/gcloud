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

package iap

import (
	iap "cloud.google.com/go/iap/apiv1"
	"cloud.google.com/go/iap/apiv1/iappb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the iap command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "iap",
		Usage: "manage Cloud Identity-Aware Proxy API resources",
		Commands: []*cli.Command{
			{
				Name:  "brands",
				Usage: "Manage brands resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list brands",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.ListBrandsRequest{
								Parent: parent,
							}

							resp, err := client.ListBrands(ctx, req)
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
						Usage: "create brands",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.CreateBrandRequest{
								Parent: parent,
							}

							resp, err := client.CreateBrand(ctx, req)
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
						Usage: "describe brands",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "brand", Usage: "The ID of the brand.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/brands/%s", cmd.String("project"), cmd.String("brand"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.GetBrandRequest{
								Name: name,
							}

							resp, err := client.GetBrand(ctx, req)
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
				Name:  "dest-groups",
				Usage: "Manage dest-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list dest-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of groups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListTunnelDestGroups`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.ListTunnelDestGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTunnelDestGroups(ctx, req)
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
						Usage: "create dest-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tunnel-dest-group-id", Usage: "The ID to use for the TunnelDestGroup, which becomes the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.CreateTunnelDestGroupRequest{
								Parent:            parent,
								TunnelDestGroupId: cmd.String("tunnel-dest-group-id"),
							}

							resp, err := client.CreateTunnelDestGroup(ctx, req)
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
						Usage: "describe dest-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dest-group", Usage: "The ID of the dest group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("dest-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.GetTunnelDestGroupRequest{
								Name: name,
							}

							resp, err := client.GetTunnelDestGroup(ctx, req)
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
						Usage: "delete dest-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dest-group", Usage: "The ID of the dest group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("dest-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTunnelDestGroup on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.DeleteTunnelDestGroupRequest{
								Name: name,
							}

							if err := client.DeleteTunnelDestGroup(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update dest-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dest-group", Usage: "The ID of the dest group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tunnel_dest_group.name" not yet supported.
							tunnel_dest_group_name := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("dest-group"))
							fmt.Printf("Executing update on %s\n", tunnel_dest_group_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "iap-settings",
				Usage: "Manage iap-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update iap-settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "iap_settings.name" not yet supported.
							fmt.Println("Executing update...")
							return nil
						},
					},
				},
			},
			{
				Name:  "identity-aware-proxy-clients",
				Usage: "Manage identity-aware-proxy-clients resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create identity-aware-proxy-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "brand", Usage: "The ID of the brand.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/brands/%s", cmd.String("project"), cmd.String("brand"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.CreateIdentityAwareProxyClientRequest{
								Parent: parent,
							}

							resp, err := client.CreateIdentityAwareProxyClient(ctx, req)
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
						Usage: "list identity-aware-proxy-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "brand", Usage: "The ID of the brand.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of clients to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListIdentityAwareProxyClients`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/brands/%s", cmd.String("project"), cmd.String("brand"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.ListIdentityAwareProxyClientsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListIdentityAwareProxyClients(ctx, req)
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
						Usage: "describe identity-aware-proxy-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "brand", Usage: "The ID of the brand.", Required: true},
							&cli.StringFlag{Name: "identity-aware-proxy-client", Usage: "The ID of the identity aware proxy client.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/brands/%s/identityAwareProxyClients/%s", cmd.String("project"), cmd.String("brand"), cmd.String("identity-aware-proxy-client"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.GetIdentityAwareProxyClientRequest{
								Name: name,
							}

							resp, err := client.GetIdentityAwareProxyClient(ctx, req)
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
						Name:  "reset-secret",
						Usage: "reset-secret identity-aware-proxy-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "brand", Usage: "The ID of the brand.", Required: true},
							&cli.StringFlag{Name: "identity-aware-proxy-client", Usage: "The ID of the identity aware proxy client.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/brands/%s/identityAwareProxyClients/%s", cmd.String("project"), cmd.String("brand"), cmd.String("identity-aware-proxy-client"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.ResetIdentityAwareProxyClientSecretRequest{
								Name: name,
							}

							resp, err := client.ResetIdentityAwareProxyClientSecret(ctx, req)
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
						Usage: "delete identity-aware-proxy-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "brand", Usage: "The ID of the brand.", Required: true},
							&cli.StringFlag{Name: "identity-aware-proxy-client", Usage: "The ID of the identity aware proxy client.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/brands/%s/identityAwareProxyClients/%s", cmd.String("project"), cmd.String("brand"), cmd.String("identity-aware-proxy-client"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIdentityAwareProxyClient on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.DeleteIdentityAwareProxyClientRequest{
								Name: name,
							}

							if err := client.DeleteIdentityAwareProxyClient(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "name",
				Usage: "Manage name resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe name",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Usage: "The ID of the name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("%s", cmd.String("name"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.GetIapSettingsRequest{
								Name: name,
							}

							resp, err := client.GetIapSettings(ctx, req)
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
						Name:  "validate-attribute-expression",
						Usage: "validate-attribute-expression name",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "expression", Usage: "User input string expression.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The ID of the name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("%s", cmd.String("name"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.ValidateIapAttributeExpressionRequest{
								Name:       name,
								Expression: cmd.String("expression"),
							}

							resp, err := client.ValidateIapAttributeExpression(ctx, req)
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
				Name:  "resource",
				Usage: "Manage resource resources",
				Commands: []*cli.Command{

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy resource",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("%s", cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.SetIamPolicyRequest{
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy resource",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("%s", cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.GetIamPolicyRequest{
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
						Usage: "test-iam-permissions resource",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("%s", cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iap.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iappb.TestIamPermissionsRequest{
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
		},
	}
}
