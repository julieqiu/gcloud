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
				Name:  "iamcredentials",
				Usage: "manage IAM Service Account Credentials API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "service-accounts",
						Usage: "Manage service-accounts resources",
						Commands: []*cli.Command{
							{
								Name:  "generate-access-token",
								Usage: "generate-access-token service-accounts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service_account", Usage: "The service_account.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service_account"))
									fmt.Printf("Executing generate-access-token on %s\n", name)
									return nil
								},
							},
							{
								Name:  "generate-id-token",
								Usage: "generate-id-token service-accounts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service_account", Usage: "The service_account.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service_account"))
									fmt.Printf("Executing generate-id-token on %s\n", name)
									return nil
								},
							},
							{
								Name:  "sign-blob",
								Usage: "sign-blob service-accounts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service_account", Usage: "The service_account.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service_account"))
									fmt.Printf("Executing sign-blob on %s\n", name)
									return nil
								},
							},
							{
								Name:  "sign-jwt",
								Usage: "sign-jwt service-accounts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service_account", Usage: "The service_account.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/serviceAccounts/%s", cmd.String("project"), cmd.String("service_account"))
									fmt.Printf("Executing sign-jwt on %s\n", name)
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
