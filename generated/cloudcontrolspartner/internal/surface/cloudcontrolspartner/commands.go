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

package cloudcontrolspartner

import (
	cloudcontrolspartner "cloud.google.com/go/cloudcontrolspartner/apiv1"
	"cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudcontrolspartner command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudcontrolspartner",
		Usage: "manage Cloud Controls Partner API resources",
		Commands: []*cli.Command{
			{
				Name:  "access-approval-requests",
				Usage: "Manage access-approval-requests resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list access-approval-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of access requests to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.ListAccessApprovalRequestsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAccessApprovalRequests(ctx, req)
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
				Name:  "customers",
				Usage: "Manage customers resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe customers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.GetCustomerRequest{
								Name: name,
							}

							resp, err := client.GetCustomer(ctx, req)
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
						Usage: "list customers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Customers to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListCustomers` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.ListCustomersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCustomers(ctx, req)
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
						Usage: "create customers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer-id", Usage: "The customer id to use for the customer, which will become the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.CreateCustomerRequest{
								Parent:     parent,
								CustomerId: cmd.String("customer-id"),
							}

							resp, err := client.CreateCustomer(ctx, req)
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
						Usage: "update customers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "customer.name" not yet supported.
							customer_name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
							fmt.Printf("Executing update on %s\n", customer_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete customers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCustomer on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.DeleteCustomerRequest{
								Name: name,
							}

							if err := client.DeleteCustomer(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "ekm-connections",
				Usage: "Manage ekm-connections resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe ekm-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s/ekmConnections", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.GetEkmConnectionsRequest{
								Name: name,
							}

							resp, err := client.GetEkmConnections(ctx, req)
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
				Name:  "partner",
				Usage: "Manage partner resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe partner",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/partner", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.GetPartnerRequest{
								Name: name,
							}

							resp, err := client.GetPartner(ctx, req)
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
				Name:  "partner-permissions",
				Usage: "Manage partner-permissions resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe partner-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s/partnerPermissions", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.GetPartnerPermissionsRequest{
								Name: name,
							}

							resp, err := client.GetPartnerPermissions(ctx, req)
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
				Name:  "violations",
				Usage: "Manage violations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list violations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of customers row to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListViolations` call.", Required: false},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.ListViolationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListViolations(ctx, req)
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
						Usage: "describe violations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "violation", Usage: "The ID of the violation.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s/violations/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"), cmd.String("violation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.GetViolationRequest{
								Name: name,
							}

							resp, err := client.GetViolation(ctx, req)
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
				Name:  "workloads",
				Usage: "Manage workloads resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.GetWorkloadRequest{
								Name: name,
							}

							resp, err := client.GetWorkload(ctx, req)
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
						Usage: "list workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "customer", Usage: "The ID of the customer.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of workloads to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListWorkloads` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudcontrolspartner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudcontrolspartnerpb.ListWorkloadsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkloads(ctx, req)
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
