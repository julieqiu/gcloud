package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	cloudcontrolspartner "cloud.google.com/go/cloudcontrolspartner/apiv1"
	"cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
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
				Name:  "cloudcontrolspartner",
				Usage: "manage Cloud Controls Partner API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "access-approval-requests",
						Usage: "Manage access-approval-requests resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list access-approval-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cloudcontrolspartnerpb.ListAccessApprovalRequestsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAccessApprovalRequests(ctx, req)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.GetCustomerRequest{Name: name}
									resp, err := client.GetCustomer(ctx, req)
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
								Name:  "list",
								Usage: "list customers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cloudcontrolspartnerpb.ListCustomersRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCustomers(ctx, req)
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
								Name:  "create",
								Usage: "create customers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer-id", Usage: "The customer id.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.CreateCustomerRequest{Parent: parent}
									req.CustomerId = cmd.String("customer-id")
									req.Customer = &cloudcontrolspartnerpb.Customer{
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateCustomer(ctx, req)
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
								Usage: "update customers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.UpdateCustomerRequest{}
									req.Customer = &cloudcontrolspartnerpb.Customer{
										Name:        name,
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateCustomer(ctx, req)
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
								Usage: "delete customers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.DeleteCustomerRequest{Name: name}
									if err := client.DeleteCustomer(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s/ekmConnections", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.GetEkmConnectionsRequest{Name: name}
									resp, err := client.GetEkmConnections(ctx, req)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/partner", cmd.String("organization"), cmd.String("location"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.GetPartnerRequest{Name: name}
									resp, err := client.GetPartner(ctx, req)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s/partnerPermissions", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.GetPartnerPermissionsRequest{Name: name}
									resp, err := client.GetPartnerPermissions(ctx, req)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerMonitoringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cloudcontrolspartnerpb.ListViolationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListViolations(ctx, req)
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
								Usage: "describe violations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
									&cli.StringFlag{Name: "violation", Usage: "The violation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s/violations/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"), cmd.String("violation"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerMonitoringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.GetViolationRequest{Name: name}
									resp, err := client.GetViolation(ctx, req)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/customers/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"), cmd.String("workload"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cloudcontrolspartnerpb.GetWorkloadRequest{Name: name}
									resp, err := client.GetWorkload(ctx, req)
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
								Name:  "list",
								Usage: "list workloads",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "customer", Usage: "The customer.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s/customers/%s", cmd.String("organization"), cmd.String("location"), cmd.String("customer"))
									client, err := cloudcontrolspartner.NewCloudControlsPartnerCoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cloudcontrolspartnerpb.ListWorkloadsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWorkloads(ctx, req)
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
