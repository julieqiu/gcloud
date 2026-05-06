package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	domains "cloud.google.com/go/domains/apiv1"
	"cloud.google.com/go/domains/apiv1/domainspb"
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
				Name:  "domains",
				Usage: "manage Cloud Domains API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
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
									client, err := domains.NewClient(ctx)
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
									client, err := domains.NewClient(ctx)
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
						Name:  "registrations",
						Usage: "Manage registrations resources",
						Commands: []*cli.Command{
							{
								Name:  "search-domains",
								Usage: "search-domains registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search-domains...")
									return nil
								},
							},
							{
								Name:  "retrieve-register-parameters",
								Usage: "retrieve-register-parameters registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing retrieve-register-parameters...")
									return nil
								},
							},
							{
								Name:  "register",
								Usage: "register registrations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing register on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "retrieve-transfer-parameters",
								Usage: "retrieve-transfer-parameters registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing retrieve-transfer-parameters...")
									return nil
								},
							},
							{
								Name:  "transfer",
								Usage: "transfer registrations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing transfer on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe registrations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
									client, err := domains.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &domainspb.GetRegistrationRequest{Name: name}
									resp, err := client.GetRegistration(ctx, req)
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
								Usage: "update registrations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
									&cli.StringFlag{Name: "domain-name", Usage: "The domain name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
									client, err := domains.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &domainspb.UpdateRegistrationRequest{}
									req.Registration = &domainspb.Registration{
										Name:       name,
										DomainName: cmd.String("domain-name"),
									}
									var paths []string
									if cmd.IsSet("domain-name") {
										paths = append(paths, "domain_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateRegistration(ctx, req)
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
								Name:  "configure-management-settings",
								Usage: "configure-management-settings registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing configure-management-settings...")
									return nil
								},
							},
							{
								Name:  "configure-dns-settings",
								Usage: "configure-dns-settings registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing configure-dns-settings...")
									return nil
								},
							},
							{
								Name:  "configure-contact-settings",
								Usage: "configure-contact-settings registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing configure-contact-settings...")
									return nil
								},
							},
							{
								Name:  "export",
								Usage: "export registrations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
									client, err := domains.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &domainspb.ExportRegistrationRequest{Name: name}
									op, err := client.ExportRegistration(ctx, req)
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
								Usage: "delete registrations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
									client, err := domains.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &domainspb.DeleteRegistrationRequest{Name: name}
									op, err := client.DeleteRegistration(ctx, req)
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
							{
								Name:  "retrieve-authorization-code",
								Usage: "retrieve-authorization-code registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing retrieve-authorization-code...")
									return nil
								},
							},
							{
								Name:  "reset-authorization-code",
								Usage: "reset-authorization-code registrations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reset-authorization-code...")
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
