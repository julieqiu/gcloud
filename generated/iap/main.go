package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	iap "cloud.google.com/go/iap/apiv1"
	"cloud.google.com/go/iap/apiv1/iappb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "iap",
				Usage: "manage Cloud Identity-Aware Proxy API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "brands",
						Usage: "Manage brands resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list brands",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create brands",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe brands",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/iap_tunnel", cmd.String("project"))
									client, err := iap.NewIdentityAwareProxyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &iappb.ListTunnelDestGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTunnelDestGroups(ctx, req)
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
								Usage: "create dest-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tunnel-dest-group-id", Usage: "The tunnel dest group id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := iap.NewIdentityAwareProxyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &iappb.CreateTunnelDestGroupRequest{Parent: parent}
									req.TunnelDestGroupId = cmd.String("tunnel-dest-group-id")
									resp, err := client.CreateTunnelDestGroup(ctx, req)
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
								Usage: "describe dest-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dest_group", Usage: "The dest_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("dest_group"))
									client, err := iap.NewIdentityAwareProxyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &iappb.GetTunnelDestGroupRequest{Name: name}
									resp, err := client.GetTunnelDestGroup(ctx, req)
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
								Usage: "delete dest-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dest_group", Usage: "The dest_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("dest_group"))
									client, err := iap.NewIdentityAwareProxyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &iappb.DeleteTunnelDestGroupRequest{Name: name}
									if err := client.DeleteTunnelDestGroup(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update dest-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dest_group", Usage: "The dest_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/iap_tunnel/locations/%s/destGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("dest_group"))
									client, err := iap.NewIdentityAwareProxyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &iappb.UpdateTunnelDestGroupRequest{}
									req.TunnelDestGroup = &iappb.TunnelDestGroup{
										Name: name,
									}
									resp, err := client.UpdateTunnelDestGroup(ctx, req)
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
						Name:  "identity-aware-proxy-clients",
						Usage: "Manage identity-aware-proxy-clients resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create identity-aware-proxy-clients",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list identity-aware-proxy-clients",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe identity-aware-proxy-clients",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "reset-secret",
								Usage: "reset-secret identity-aware-proxy-clients",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reset-secret...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete identity-aware-proxy-clients",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
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
