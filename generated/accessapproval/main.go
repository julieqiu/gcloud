package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	accessapproval "cloud.google.com/go/accessapproval/apiv1"
	"cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
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
				Name:  "accessapproval",
				Usage: "manage Access Approval API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "access-approval-settings",
						Usage: "Manage access-approval-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe access-approval-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/accessApprovalSettings", cmd.String("project"))
									client, err := accessapproval.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &accessapprovalpb.GetAccessApprovalSettingsMessage{Name: name}
									resp, err := client.GetAccessApprovalSettings(ctx, req)
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
								Usage: "update access-approval-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "active-key-version", Usage: "The active key version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/accessApprovalSettings", cmd.String("project"))
									client, err := accessapproval.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &accessapprovalpb.UpdateAccessApprovalSettingsMessage{}
									req.Settings = &accessapprovalpb.AccessApprovalSettings{
										Name:             name,
										Name:             cmd.String("name"),
										ActiveKeyVersion: cmd.String("active-key-version"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("active-key-version") {
										paths = append(paths, "active_key_version")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateAccessApprovalSettings(ctx, req)
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
								Usage: "delete access-approval-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/accessApprovalSettings", cmd.String("project"))
									client, err := accessapproval.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &accessapprovalpb.DeleteAccessApprovalSettingsMessage{Name: name}
									if err := client.DeleteAccessApprovalSettings(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "approval-requests",
						Usage: "Manage approval-requests resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list approval-requests",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := accessapproval.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &accessapprovalpb.ListApprovalRequestsMessage{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListApprovalRequests(ctx, req)
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
								Usage: "describe approval-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "approval_request", Usage: "The approval_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval_request"))
									client, err := accessapproval.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &accessapprovalpb.GetApprovalRequestMessage{Name: name}
									resp, err := client.GetApprovalRequest(ctx, req)
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
								Name:  "approve",
								Usage: "approve approval-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "approval_request", Usage: "The approval_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval_request"))
									fmt.Printf("Executing approve on %s\n", name)
									return nil
								},
							},
							{
								Name:  "dismiss",
								Usage: "dismiss approval-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "approval_request", Usage: "The approval_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval_request"))
									fmt.Printf("Executing dismiss on %s\n", name)
									return nil
								},
							},
							{
								Name:  "invalidate",
								Usage: "invalidate approval-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "approval_request", Usage: "The approval_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval_request"))
									fmt.Printf("Executing invalidate on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "service-account",
						Usage: "Manage service-account resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe service-account",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
