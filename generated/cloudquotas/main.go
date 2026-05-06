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
				Name:  "cloudquotas",
				Usage: "manage Cloud Quotas API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "quota-infos",
						Usage: "Manage quota-infos resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list quota-infos",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe quota-infos",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "quota_info", Usage: "The quota_info.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s/quotaInfos/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("quota_info"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "quota-preferences",
						Usage: "Manage quota-preferences resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list quota-preferences",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe quota-preferences",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "quota_preference", Usage: "The quota_preference.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/quotaPreferences/%s", cmd.String("project"), cmd.String("location"), cmd.String("quota_preference"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create quota-preferences",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update quota-preferences",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "quota_preference", Usage: "The quota_preference.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/quotaPreferences/%s", cmd.String("project"), cmd.String("location"), cmd.String("quota_preference"))
									fmt.Printf("Executing update on %s\n", name)
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
