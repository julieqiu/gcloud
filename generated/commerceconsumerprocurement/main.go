package main

import (
	"context"
	"fmt"
	"os"

	procurement "cloud.google.com/go/commerce/consumer/procurement/apiv1"
	"cloud.google.com/go/commerce/consumer/procurement/apiv1/procurementpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "cloudcommerceconsumerprocurement",
				Usage: "manage Cloud Commerce Consumer Procurement API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "license-pool",
						Usage: "Manage license-pool resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe license-pool",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update license-pool",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "billing_account", Usage: "The billing_account.", Required: true},
									&cli.StringFlag{Name: "order", Usage: "The order.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("billingAccounts/%s/orders/%s/licensePool", cmd.String("billing_account"), cmd.String("order"))
									client, err := procurement.NewLicenseManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &procurementpb.UpdateLicensePoolRequest{}
									req.LicensePool = &procurementpb.LicensePool{
										Name: name,
									}
									resp, err := client.UpdateLicensePool(ctx, req)
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
								Name:  "assign",
								Usage: "assign license-pool",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing assign...")
									return nil
								},
							},
							{
								Name:  "unassign",
								Usage: "unassign license-pool",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing unassign...")
									return nil
								},
							},
							{
								Name:  "enumerate-licensed-users",
								Usage: "enumerate-licensed-users license-pool",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing enumerate-licensed-users...")
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
									&cli.StringFlag{Name: "billingAccount", Usage: "The billingAccount.", Required: true},
									&cli.StringFlag{Name: "order", Usage: "The order.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("billingAccounts/%s/orders/%s/operations/%s", cmd.String("billingAccount"), cmd.String("order"), cmd.String("operation"))
									client, err := procurement.NewLicenseManagementClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "billingAccount", Usage: "The billingAccount.", Required: true},
									&cli.StringFlag{Name: "order", Usage: "The order.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("billingAccounts/%s/orders/%s/operations/%s", cmd.String("billingAccount"), cmd.String("order"), cmd.String("operation"))
									client, err := procurement.NewConsumerProcurementClient(ctx)
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
						},
					},
					{
						Name:  "orders",
						Usage: "Manage orders resources",
						Commands: []*cli.Command{
							{
								Name:  "place",
								Usage: "place orders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing place...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe orders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list orders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "modify",
								Usage: "modify orders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing modify...")
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel orders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
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
