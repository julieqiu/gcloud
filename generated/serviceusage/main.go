package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "serviceusage",
				Usage: "manage Service Usage API resources",
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("operations/%s", cmd.String("operation"))
									fmt.Printf("Executing describe on %s\n", name)
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
								Name:  "enable",
								Usage: "enable services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing enable...")
									return nil
								},
							},
							{
								Name:  "disable",
								Usage: "disable services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing disable...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "batch-enable",
								Usage: "batch-enable services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing batch-enable...")
									return nil
								},
							},
							{
								Name:  "batch-get",
								Usage: "batch-get services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing batch-get...")
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
