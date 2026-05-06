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

package billing

import (
	billing "cloud.google.com/go/billing/apiv1"
	"cloud.google.com/go/billing/apiv1/billingpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudbilling command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudbilling",
		Usage: "manage Cloud Billing API resources",
		Commands: []*cli.Command{
			{
				Name:  "billing-accounts",
				Usage: "Manage billing-accounts resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.GetBillingAccountRequest{
								Name: name,
							}

							resp, err := client.GetBillingAccount(ctx, req)
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
						Usage: "list billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Options for how to filter the returned billing accounts.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.ListBillingAccountsRequest{
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListBillingAccounts(ctx, req)
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
						Usage: "update billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.UpdateBillingAccountRequest{
								Name: name,
							}

							resp, err := client.UpdateBillingAccount(ctx, req)
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
						Usage: "create billing-accounts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.CreateBillingAccountRequest{}

							resp, err := client.CreateBillingAccount(ctx, req)
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
						Usage: "describe billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.GetIamPolicyRequest{
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
						Usage: "set-iam-policy billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.SetIamPolicyRequest{
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
						Usage: "test-iam-permissions billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.TestIamPermissionsRequest{
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

					{
						Name:  "move",
						Usage: "move billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.StringFlag{Name: "destination-parent", Usage: "The resource name of the Organization to move.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.MoveBillingAccountRequest{
								Name:              name,
								DestinationParent: cmd.String("destination-parent"),
							}

							resp, err := client.MoveBillingAccount(ctx, req)
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
				Name:  "billing-info",
				Usage: "Manage billing-info resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe billing-info",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.GetProjectBillingInfoRequest{
								Name: name,
							}

							resp, err := client.GetProjectBillingInfo(ctx, req)
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
						Usage: "update billing-info",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.UpdateProjectBillingInfoRequest{
								Name: name,
							}

							resp, err := client.UpdateProjectBillingInfo(ctx, req)
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
				Name:  "projects",
				Usage: "Manage projects resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account", Usage: "The ID of the billing account.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to be returned.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing-account"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.ListProjectBillingInfoRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListProjectBillingInfo(ctx, req)
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
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list services",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.ListServicesRequest{
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
				},
			},
			{
				Name:  "skus",
				Usage: "Manage skus resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list skus",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "currency-code", Usage: "The ISO 4217 currency code for the pricing info in the response proto.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to return.", Required: false},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := billing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &billingpb.ListSkusRequest{
								Parent:       parent,
								CurrencyCode: cmd.String("currency-code"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSkus(ctx, req)
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
