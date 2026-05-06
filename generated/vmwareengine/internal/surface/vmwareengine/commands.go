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

package vmwareengine

import (
	vmwareengine "cloud.google.com/go/vmwareengine/apiv1"
	"cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the vmwareengine command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "vmwareengine",
		Usage: "manage VMware Engine API resources",
		Commands: []*cli.Command{
			{
				Name:  "clusters",
				Usage: "Manage clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of clusters to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListClusters` call.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListClusters(ctx, req)
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
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetClusterRequest{
								Name: name,
							}

							resp, err := client.GetCluster(ctx, req)
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
						Usage: "create clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-id", Usage: "The user-provided identifier of the new `Cluster`.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that zero.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "True if you want the request to be validated and not executed;.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateClusterRequest{
								Parent:       parent,
								ClusterId:    cmd.String("cluster-id"),
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateCluster(ctx, req)
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
						Usage: "update clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "True if you want the request to be validated and not executed;.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "cluster.name" not yet supported.
							cluster_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("cluster"))
							fmt.Printf("Executing update on %s\n", cluster_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that zero.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteCluster(ctx, req)
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
				Name:  "dns-bind-permission",
				Usage: "Manage dns-bind-permission resources",
				Commands: []*cli.Command{

					{
						Name:  "grant",
						Usage: "grant dns-bind-permission",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsBindPermission", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GrantDnsBindPermissionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.GrantDnsBindPermission(ctx, req)
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
						Usage: "describe dns-bind-permission",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsBindPermission", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetDnsBindPermissionRequest{
								Name: name,
							}

							resp, err := client.GetDnsBindPermission(ctx, req)
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
						Name:  "revoke",
						Usage: "revoke dns-bind-permission",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsBindPermission", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.RevokeDnsBindPermissionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.RevokeDnsBindPermission(ctx, req)
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
				Name:  "dns-forwarding",
				Usage: "Manage dns-forwarding resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe dns-forwarding",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/dnsForwarding", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetDnsForwardingRequest{
								Name: name,
							}

							resp, err := client.GetDnsForwarding(ctx, req)
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
						Usage: "update dns-forwarding",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "dns_forwarding.name" not yet supported.
							dns_forwarding_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/dnsForwarding", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing update on %s\n", dns_forwarding_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "external-access-rules",
				Usage: "Manage external-access-rules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of external access rules to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListExternalAccessRulesRequest`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListExternalAccessRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListExternalAccessRules(ctx, req)
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
						Usage: "describe external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-access-rule", Usage: "The ID of the external access rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"), cmd.String("external-access-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetExternalAccessRuleRequest{
								Name: name,
							}

							resp, err := client.GetExternalAccessRule(ctx, req)
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
						Usage: "create external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-access-rule-id", Usage: "The user-provided identifier of the `ExternalAccessRule` to be.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateExternalAccessRuleRequest{
								Parent:               parent,
								ExternalAccessRuleId: cmd.String("external-access-rule-id"),
								RequestId:            cmd.String("request-id"),
							}

							op, err := client.CreateExternalAccessRule(ctx, req)
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
						Usage: "update external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-access-rule", Usage: "The ID of the external access rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "external_access_rule.name" not yet supported.
							external_access_rule_name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"), cmd.String("external-access-rule"))
							fmt.Printf("Executing update on %s\n", external_access_rule_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-access-rule", Usage: "The ID of the external access rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"), cmd.String("external-access-rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteExternalAccessRule %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteExternalAccessRuleRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteExternalAccessRule(ctx, req)
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
				Name:  "external-addresses",
				Usage: "Manage external-addresses resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of external IP addresses to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListExternalAddresses` call.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListExternalAddressesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListExternalAddresses(ctx, req)
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
						Usage: "describe external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-addresse", Usage: "The ID of the external addresse.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("external-addresse"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetExternalAddressRequest{
								Name: name,
							}

							resp, err := client.GetExternalAddress(ctx, req)
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
						Usage: "create external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-address-id", Usage: "The user-provided identifier of the `ExternalAddress` to be.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateExternalAddressRequest{
								Parent:            parent,
								ExternalAddressId: cmd.String("external-address-id"),
								RequestId:         cmd.String("request-id"),
							}

							op, err := client.CreateExternalAddress(ctx, req)
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
						Usage: "update external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-addresse", Usage: "The ID of the external addresse.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "external_address.name" not yet supported.
							external_address_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("external-addresse"))
							fmt.Printf("Executing update on %s\n", external_address_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-addresse", Usage: "The ID of the external addresse.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("external-addresse"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteExternalAddress %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteExternalAddressRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteExternalAddress(ctx, req)
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
				Name:  "hcx-activation-keys",
				Usage: "Manage hcx-activation-keys resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create hcx-activation-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hcx-activation-key-id", Usage: "The user-provided identifier of the `HcxActivationKey` to be.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateHcxActivationKeyRequest{
								Parent:             parent,
								HcxActivationKeyId: cmd.String("hcx-activation-key-id"),
								RequestId:          cmd.String("request-id"),
							}

							op, err := client.CreateHcxActivationKey(ctx, req)
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
						Name:  "list",
						Usage: "list hcx-activation-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of HCX activation keys to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListHcxActivationKeys` call.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListHcxActivationKeysRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListHcxActivationKeys(ctx, req)
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
						Usage: "describe hcx-activation-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hcx-activation-key", Usage: "The ID of the hcx activation key.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/hcxActivationKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("hcx-activation-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetHcxActivationKeyRequest{
								Name: name,
							}

							resp, err := client.GetHcxActivationKey(ctx, req)
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
				Name:  "logging-servers",
				Usage: "Manage logging-servers resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of logging servers to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListLoggingServersRequest` call.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListLoggingServersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListLoggingServers(ctx, req)
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
						Usage: "describe logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "logging-server", Usage: "The ID of the logging server.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/loggingServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("logging-server"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetLoggingServerRequest{
								Name: name,
							}

							resp, err := client.GetLoggingServer(ctx, req)
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
						Usage: "create logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "logging-server-id", Usage: "The user-provided identifier of the `LoggingServer` to be.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateLoggingServerRequest{
								Parent:          parent,
								LoggingServerId: cmd.String("logging-server-id"),
								RequestId:       cmd.String("request-id"),
							}

							op, err := client.CreateLoggingServer(ctx, req)
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
						Usage: "update logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "logging-server", Usage: "The ID of the logging server.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "logging_server.name" not yet supported.
							logging_server_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/loggingServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("logging-server"))
							fmt.Printf("Executing update on %s\n", logging_server_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "logging-server", Usage: "The ID of the logging server.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/loggingServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("logging-server"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteLoggingServer %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteLoggingServerRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteLoggingServer(ctx, req)
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
				Name:  "management-dns-zone-bindings",
				Usage: "Manage management-dns-zone-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of management DNS zone bindings to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListManagementDnsZoneBindings`.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListManagementDnsZoneBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListManagementDnsZoneBindings(ctx, req)
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
						Usage: "describe management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-dns-zone-binding", Usage: "The ID of the management dns zone binding.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("management-dns-zone-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetManagementDnsZoneBindingRequest{
								Name: name,
							}

							resp, err := client.GetManagementDnsZoneBinding(ctx, req)
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
						Usage: "create management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-dns-zone-binding-id", Usage: "The user-provided identifier of the `ManagementDnsZoneBinding`.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateManagementDnsZoneBindingRequest{
								Parent:                     parent,
								ManagementDnsZoneBindingId: cmd.String("management-dns-zone-binding-id"),
								RequestId:                  cmd.String("request-id"),
							}

							op, err := client.CreateManagementDnsZoneBinding(ctx, req)
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
						Usage: "update management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-dns-zone-binding", Usage: "The ID of the management dns zone binding.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "management_dns_zone_binding.name" not yet supported.
							management_dns_zone_binding_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("management-dns-zone-binding"))
							fmt.Printf("Executing update on %s\n", management_dns_zone_binding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-dns-zone-binding", Usage: "The ID of the management dns zone binding.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("management-dns-zone-binding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteManagementDnsZoneBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteManagementDnsZoneBindingRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteManagementDnsZoneBinding(ctx, req)
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
						Name:  "repair",
						Usage: "repair management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-dns-zone-binding", Usage: "The ID of the management dns zone binding.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("management-dns-zone-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.RepairManagementDnsZoneBindingRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.RepairManagementDnsZoneBinding(ctx, req)
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
				Name:  "network-peerings",
				Usage: "Manage network-peerings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-peering", Usage: "The ID of the network peering.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-peering"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetNetworkPeeringRequest{
								Name: name,
							}

							resp, err := client.GetNetworkPeering(ctx, req)
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
						Usage: "list network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of network peerings to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListNetworkPeerings` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListNetworkPeeringsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNetworkPeerings(ctx, req)
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
						Usage: "create network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-peering-id", Usage: "The user-provided identifier of the new `NetworkPeering`.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateNetworkPeeringRequest{
								Parent:           parent,
								NetworkPeeringId: cmd.String("network-peering-id"),
								RequestId:        cmd.String("request-id"),
							}

							op, err := client.CreateNetworkPeering(ctx, req)
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
						Usage: "delete network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-peering", Usage: "The ID of the network peering.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-peering"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNetworkPeering %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteNetworkPeeringRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteNetworkPeering(ctx, req)
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
						Name:  "update",
						Usage: "update network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-peering", Usage: "The ID of the network peering.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "network_peering.name" not yet supported.
							network_peering_name := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-peering"))
							fmt.Printf("Executing update on %s\n", network_peering_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "network-policies",
				Usage: "Manage network-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch-external-addresses",
						Usage: "fetch-external-addresses network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "network-policy", Usage: "The resource name of the network policy to query for assigned.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of external IP addresses to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							network_policy := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"))
							fmt.Printf("Executing fetch-external-addresses on %s\n", network_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetNetworkPolicyRequest{
								Name: name,
							}

							resp, err := client.GetNetworkPolicy(ctx, req)
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
						Usage: "list network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of network policies to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListNetworkPolicies` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListNetworkPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNetworkPolicies(ctx, req)
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
						Usage: "create network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policy-id", Usage: "The user-provided identifier of the network policy to be created.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateNetworkPolicyRequest{
								Parent:          parent,
								NetworkPolicyId: cmd.String("network-policy-id"),
								RequestId:       cmd.String("request-id"),
							}

							op, err := client.CreateNetworkPolicy(ctx, req)
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
						Usage: "update network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "network_policy.name" not yet supported.
							network_policy_name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"))
							fmt.Printf("Executing update on %s\n", network_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-policie", Usage: "The ID of the network policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNetworkPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteNetworkPolicyRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteNetworkPolicy(ctx, req)
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
				Name:  "node-types",
				Usage: "Manage node-types resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list node-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of node types to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListNodeTypes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListNodeTypesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListNodeTypes(ctx, req)
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
						Usage: "describe node-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-type", Usage: "The ID of the node type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodeTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetNodeTypeRequest{
								Name: name,
							}

							resp, err := client.GetNodeType(ctx, req)
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
				Name:  "nodes",
				Usage: "Manage nodes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of nodes to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListNodes` call.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListNodesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListNodes(ctx, req)
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
						Usage: "describe nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The ID of the node.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("cluster"), cmd.String("node"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetNodeRequest{
								Name: name,
							}

							resp, err := client.GetNode(ctx, req)
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
				},
			},
			{
				Name:  "peering-routes",
				Usage: "Manage peering-routes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list peering-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-peering", Usage: "The ID of the network peering.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of peering routes to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListPeeringRoutes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network-peering"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListPeeringRoutesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPeeringRoutes(ctx, req)
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
						Usage: "list peering-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of peering routes to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListPrivateConnectionPeeringRoutes`.", Required: false},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListPrivateConnectionPeeringRoutesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPrivateConnectionPeeringRoutes(ctx, req)
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
				Name:  "private-clouds",
				Usage: "Manage private-clouds resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of private clouds to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListPrivateClouds` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListPrivateCloudsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPrivateClouds(ctx, req)
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
						Usage: "describe private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetPrivateCloudRequest{
								Name: name,
							}

							resp, err := client.GetPrivateCloud(ctx, req)
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
						Usage: "create private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud-id", Usage: "The user-provided identifier of the private cloud to be created.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that zero.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "True if you want the request to be validated and not executed;.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreatePrivateCloudRequest{
								Parent:         parent,
								PrivateCloudId: cmd.String("private-cloud-id"),
								RequestId:      cmd.String("request-id"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.CreatePrivateCloud(ctx, req)
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
						Usage: "update private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that zero.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "private_cloud.name" not yet supported.
							private_cloud_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing update on %s\n", private_cloud_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete private-clouds",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "delay-hours", Usage: "Time delay of the deletion specified in hours.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, cascade delete is enabled and all children of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that zero.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeletePrivateCloudRequest{
								Name:       name,
								RequestId:  cmd.String("request-id"),
								Force:      cmd.Bool("force"),
								DelayHours: runtime.Ptr(int32(cmd.Int("delay-hours"))),
							}

							op, err := client.DeletePrivateCloud(ctx, req)
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
						Name:  "undelete",
						Usage: "undelete private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "The request ID must be a valid UUID with the exception that zero.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.UndeletePrivateCloudRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.UndeletePrivateCloud(ctx, req)
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
						Name:  "show-nsx-credentials",
						Usage: "show-nsx-credentials private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							private_cloud := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing show-nsx-credentials on %s\n", private_cloud)
							return nil
						},
					},

					{
						Name:  "show-vcenter-credentials",
						Usage: "show-vcenter-credentials private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "username", Usage: "The username of the user to be queried for credentials.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							private_cloud := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing show-vcenter-credentials on %s\n", private_cloud)
							return nil
						},
					},

					{
						Name:  "reset-nsx-credentials",
						Usage: "reset-nsx-credentials private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							private_cloud := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing reset-nsx-credentials on %s\n", private_cloud)
							return nil
						},
					},

					{
						Name:  "reset-vcenter-credentials",
						Usage: "reset-vcenter-credentials private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "username", Usage: "The username of the user to be to reset the credentials.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							private_cloud := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing reset-vcenter-credentials on %s\n", private_cloud)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "private-connections",
				Usage: "Manage private-connections resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection-id", Usage: "The user-provided identifier of the new private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreatePrivateConnectionRequest{
								Parent:              parent,
								PrivateConnectionId: cmd.String("private-connection-id"),
								RequestId:           cmd.String("request-id"),
							}

							op, err := client.CreatePrivateConnection(ctx, req)
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
						Usage: "describe private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetPrivateConnectionRequest{
								Name: name,
							}

							resp, err := client.GetPrivateConnection(ctx, req)
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
						Usage: "list private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of private connections to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListPrivateConnections` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListPrivateConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPrivateConnections(ctx, req)
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
						Usage: "update private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "private_connection.name" not yet supported.
							private_connection_name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							fmt.Printf("Executing update on %s\n", private_connection_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePrivateConnection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeletePrivateConnectionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeletePrivateConnection(ctx, req)
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
				Name:  "subnets",
				Usage: "Manage subnets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of subnets to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSubnetsRequest` call.", Required: false},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListSubnetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSubnets(ctx, req)
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
						Usage: "describe subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The ID of the subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("subnet"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetSubnetRequest{
								Name: name,
							}

							resp, err := client.GetSubnet(ctx, req)
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
						Usage: "update subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-cloud", Usage: "The ID of the private cloud.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The ID of the subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "subnet.name" not yet supported.
							subnet_name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-cloud"), cmd.String("subnet"))
							fmt.Printf("Executing update on %s\n", subnet_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "vmware-engine-networks",
				Usage: "Manage vmware-engine-networks resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network-id", Usage: "The user-provided identifier of the new VMware Engine network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.CreateVmwareEngineNetworkRequest{
								Parent:                parent,
								VmwareEngineNetworkId: cmd.String("vmware-engine-network-id"),
								RequestId:             cmd.String("request-id"),
							}

							op, err := client.CreateVmwareEngineNetwork(ctx, req)
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
						Usage: "update vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The ID of the vmware engine network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "vmware_engine_network.name" not yet supported.
							vmware_engine_network_name := fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("vmware-engine-network"))
							fmt.Printf("Executing update on %s\n", vmware_engine_network_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Checksum used to ensure that the user-provided value is up to.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The ID of the vmware engine network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("vmware-engine-network"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteVmwareEngineNetwork %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.DeleteVmwareEngineNetworkRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      cmd.String("etag"),
							}

							op, err := client.DeleteVmwareEngineNetwork(ctx, req)
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
						Name:  "describe",
						Usage: "describe vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The ID of the vmware engine network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("vmware-engine-network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.GetVmwareEngineNetworkRequest{
								Name: name,
							}

							resp, err := client.GetVmwareEngineNetwork(ctx, req)
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
						Usage: "list vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that matches resources returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in one page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListVmwareEngineNetworks` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmwareengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmwareenginepb.ListVmwareEngineNetworksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListVmwareEngineNetworks(ctx, req)
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
		},
	}
}
