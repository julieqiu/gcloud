package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	privilegedaccessmanager "cloud.google.com/go/privilegedaccessmanager/apiv1"
	"cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "privilegedaccessmanager",
				Usage: "manage Privileged Access Manager API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "entitlements",
						Usage: "Manage entitlements resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list entitlements",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &privilegedaccessmanagerpb.ListEntitlementsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListEntitlements(ctx, req)
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
								Name:  "search",
								Usage: "search entitlements",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing search on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe entitlements",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.GetEntitlementRequest{Name: name}
									resp, err := client.GetEntitlement(ctx, req)
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
								Name:  "create",
								Usage: "create entitlements",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement-id", Usage: "The entitlement id.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.CreateEntitlementRequest{Parent: parent}
									req.EntitlementId = cmd.String("entitlement-id")
									req.Entitlement = &privilegedaccessmanagerpb.Entitlement{
										Etag: cmd.String("etag"),
									}
									op, err := client.CreateEntitlement(ctx, req)
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
								Usage: "delete entitlements",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.DeleteEntitlementRequest{Name: name}
									op, err := client.DeleteEntitlement(ctx, req)
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
								Name:  "update",
								Usage: "update entitlements",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.UpdateEntitlementRequest{}
									req.Entitlement = &privilegedaccessmanagerpb.Entitlement{
										Name: name,
										Etag: cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateEntitlement(ctx, req)
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
						},
					},
					{
						Name:  "grants",
						Usage: "Manage grants resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &privilegedaccessmanagerpb.ListGrantsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGrants(ctx, req)
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
								Name:  "search",
								Usage: "search grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									fmt.Printf("Executing search on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
									&cli.StringFlag{Name: "grant", Usage: "The grant.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.GetGrantRequest{Name: name}
									resp, err := client.GetGrant(ctx, req)
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
								Name:  "create",
								Usage: "create grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.CreateGrantRequest{Parent: parent}
									resp, err := client.CreateGrant(ctx, req)
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
								Usage: "approve grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
									&cli.StringFlag{Name: "grant", Usage: "The grant.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
									fmt.Printf("Executing approve on %s\n", name)
									return nil
								},
							},
							{
								Name:  "deny",
								Usage: "deny grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
									&cli.StringFlag{Name: "grant", Usage: "The grant.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
									fmt.Printf("Executing deny on %s\n", name)
									return nil
								},
							},
							{
								Name:  "revoke",
								Usage: "revoke grants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
									&cli.StringFlag{Name: "grant", Usage: "The grant.", Required: true},
									&cli.StringFlag{Name: "reason", Usage: "The reason.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privilegedaccessmanagerpb.RevokeGrantRequest{Name: name}
									req.Reason = cmd.String("reason")
									op, err := client.RevokeGrant(ctx, req)
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
						},
					},
					{
						Name:  "locations",
						Usage: "Manage locations resources",
						Commands: []*cli.Command{
							{
								Name:  "check-onboarding-status",
								Usage: "check-onboarding-status locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "entitlement", Usage: "The entitlement.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
									fmt.Printf("Executing check-onboarding-status on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list locations",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
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
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
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
									client, err := privilegedaccessmanager.NewClient(ctx)
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
									client, err := privilegedaccessmanager.NewClient(ctx)
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
								Name:  "delete",
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := privilegedaccessmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
