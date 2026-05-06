package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	optimization "cloud.google.com/go/optimization/apiv1"
	"cloud.google.com/go/optimization/apiv1/optimizationpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "cloudoptimization",
				Usage: "manage Cloud Optimization API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "locations",
						Usage: "Manage locations resources",
						Commands: []*cli.Command{
							{
								Name:  "optimize-tours",
								Usage: "optimize-tours locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing optimize-tours...")
									return nil
								},
							},
							{
								Name:  "batch-optimize-tours",
								Usage: "batch-optimize-tours locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing batch-optimize-tours...")
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := optimization.NewFleetRoutingClient(ctx)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
