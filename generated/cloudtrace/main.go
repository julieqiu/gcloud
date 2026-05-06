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
				Name:  "cloudtrace",
				Usage: "manage Stackdriver Trace API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "spans",
						Usage: "Manage spans resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create spans",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
						},
					},
					{
						Name:  "traces",
						Usage: "Manage traces resources",
						Commands: []*cli.Command{
							{
								Name:  "batch-write",
								Usage: "batch-write traces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "trace", Usage: "The trace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/traces/%s", cmd.String("project"), cmd.String("trace"))
									fmt.Printf("Executing batch-write on %s\n", parent)
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
