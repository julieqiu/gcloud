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
				Name:  "policytroubleshooter",
				Usage: "manage Policy Troubleshooter API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "iam",
						Usage: "Manage iam resources",
						Commands: []*cli.Command{
							{
								Name:  "troubleshoot",
								Usage: "troubleshoot iam",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing troubleshoot...")
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
