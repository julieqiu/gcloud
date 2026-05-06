package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	universalledger "cloud.google.com/go/universalledger/apiv1"
	"cloud.google.com/go/universalledger/apiv1/universalledgerpb"
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
				Name:  "universalledger",
				Usage: "manage Universal Ledger API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "endpoints",
						Usage: "Manage endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "submit-transaction",
								Usage: "submit-transaction endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing submit-transaction...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := universalledger.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &universalledgerpb.ListEndpointsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListEndpoints(ctx, req)
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
								Usage: "describe endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "endpoint", Usage: "The endpoint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
									client, err := universalledger.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &universalledgerpb.GetEndpointRequest{Name: name}
									resp, err := client.GetEndpoint(ctx, req)
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
								Name:  "submit-operational-transaction",
								Usage: "submit-operational-transaction endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing submit-operational-transaction...")
									return nil
								},
							},
							{
								Name:  "query-transaction-state",
								Usage: "query-transaction-state endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-transaction-state...")
									return nil
								},
							},
							{
								Name:  "query-account",
								Usage: "query-account endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-account...")
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
