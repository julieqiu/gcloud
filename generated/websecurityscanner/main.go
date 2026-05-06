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
				Name:  "websecurityscanner",
				Usage: "manage Web Security Scanner API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "crawled-urls",
						Usage: "Manage crawled-urls resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list crawled-urls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "finding-type-stats",
						Usage: "Manage finding-type-stats resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list finding-type-stats",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "findings",
						Usage: "Manage findings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe findings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list findings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "scan-configs",
						Usage: "Manage scan-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create scan-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete scan-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe scan-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list scan-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update scan-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "start",
								Usage: "start scan-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start...")
									return nil
								},
							},
						},
					},
					{
						Name:  "scan-runs",
						Usage: "Manage scan-runs resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe scan-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list scan-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "stop",
								Usage: "stop scan-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop...")
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
