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

package procurement

import (
	procurement "cloud.google.com/go/procurement/apiv1"
	"cloud.google.com/go/procurement/apiv1/procurementpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudcommerceconsumerprocurement command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudcommerceconsumerprocurement",
		Usage: "manage Cloud Commerce Consumer Procurement API resources",
		Commands: []*cli.Command{
			{
				Name:  "license-pool",
				Usage: "Manage license-pool resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe license-pool",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s/orders/%s/licensePool", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.GetLicensePoolRequest{
								Name: name,
							}

							resp, err := client.GetLicensePool(ctx, req)
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
						Usage: "update license-pool",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "license_pool.name" not yet supported.
							license_pool_name := fmt.Sprintf("billingAccounts/%s/orders/%s/licensePool", cmd.String("billing-account"), cmd.String("order"))
							fmt.Printf("Executing update on %s\n", license_pool_name)
							return nil
						},
					},

					{
						Name:  "assign",
						Usage: "assign license-pool",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
							&cli.StringSliceFlag{Name: "usernames", Usage: "Username.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("billingAccounts/%s/orders/%s/licensePool", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.AssignRequest{
								Parent:    parent,
								Usernames: cmd.StringSlice("usernames"),
							}

							resp, err := client.Assign(ctx, req)
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
						Name:  "unassign",
						Usage: "unassign license-pool",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
							&cli.StringSliceFlag{Name: "usernames", Usage: "Username.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("billingAccounts/%s/orders/%s/licensePool", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.UnassignRequest{
								Parent:    parent,
								Usernames: cmd.StringSlice("usernames"),
							}

							resp, err := client.Unassign(ctx, req)
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
						Name:  "enumerate-licensed-users",
						Usage: "enumerate-licensed-users license-pool",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of users to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `EnumerateLicensedUsers`.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("billingAccounts/%s/orders/%s/licensePool", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.EnumerateLicensedUsersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.EnumerateLicensedUsers(ctx, req)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s/orders/%s/operations/%s", cmd.String("billing-account"), cmd.String("order"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s/orders/%s/operations/%s", cmd.String("billing-account"), cmd.String("order"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "orders",
				Usage: "Manage orders resources",
				Commands: []*cli.Command{

					{
						Name:  "place",
						Usage: "place orders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The user-specified name of the order being placed.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.PlaceOrderRequest{
								Parent:      parent,
								DisplayName: cmd.String("display-name"),
								RequestId:   cmd.String("request-id"),
							}

							op, err := client.PlaceOrder(ctx, req)
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
						Usage: "describe orders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s/orders/%s", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.GetOrderRequest{
								Name: name,
							}

							resp, err := client.GetOrder(ctx, req)
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
						Usage: "list orders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter that you can use to limit the list request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of entries requested.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The token for fetching the next page.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.ListOrdersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListOrders(ctx, req)
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
						Name:  "modify",
						Usage: "modify orders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "Updated display name of the order, leave as empty if you do not.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The weak etag, which can be optionally populated, of the order.", Required: false},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s/orders/%s", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.ModifyOrderRequest{
								Name:        name,
								DisplayName: cmd.String("display-name"),
								Etag:        cmd.String("etag"),
							}

							op, err := client.ModifyOrder(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel orders",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "cancellation-policy", Usage: "Cancellation policy of this request.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The weak etag, which can be optionally populated, of the order.", Required: false},
							&cli.StringFlag{Name: "order", Usage: "The ID of the order.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s/orders/%s", cmd.String("billing-account"), cmd.String("order"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := procurement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &procurementpb.CancelOrderRequest{
								Name:               name,
								Etag:               cmd.String("etag"),
								CancellationPolicy: procurementpb.CancelOrderRequest_CancellationPolicy(procurementpb.CancelOrderRequest_CancellationPolicy_value[cmd.String("cancellation-policy")]),
							}

							op, err := client.CancelOrder(ctx, req)
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
		},
	}
}
