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
				Name:  "language",
				Usage: "manage Cloud Natural Language API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "documents",
						Usage: "Manage documents resources",
						Commands: []*cli.Command{
							{
								Name:  "analyze-sentiment",
								Usage: "analyze-sentiment documents",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing analyze-sentiment...")
									return nil
								},
							},
							{
								Name:  "analyze-entities",
								Usage: "analyze-entities documents",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing analyze-entities...")
									return nil
								},
							},
							{
								Name:  "classify-text",
								Usage: "classify-text documents",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing classify-text...")
									return nil
								},
							},
							{
								Name:  "moderate-text",
								Usage: "moderate-text documents",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing moderate-text...")
									return nil
								},
							},
							{
								Name:  "annotate-text",
								Usage: "annotate-text documents",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing annotate-text...")
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
