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
				Name:  "identitytoolkit",
				Usage: "manage Identity Toolkit API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "mfa-enrollment",
						Usage: "Manage mfa-enrollment resources",
						Commands: []*cli.Command{
							{
								Name:  "finalize",
								Usage: "finalize mfa-enrollment",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing finalize...")
									return nil
								},
							},
							{
								Name:  "start",
								Usage: "start mfa-enrollment",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start...")
									return nil
								},
							},
							{
								Name:  "withdraw",
								Usage: "withdraw mfa-enrollment",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing withdraw...")
									return nil
								},
							},
						},
					},
					{
						Name:  "mfa-sign-in",
						Usage: "Manage mfa-sign-in resources",
						Commands: []*cli.Command{
							{
								Name:  "finalize",
								Usage: "finalize mfa-sign-in",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing finalize...")
									return nil
								},
							},
							{
								Name:  "start",
								Usage: "start mfa-sign-in",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start...")
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
