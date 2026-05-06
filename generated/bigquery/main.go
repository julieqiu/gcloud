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
				Name:  "bigquery",
				Usage: "manage BigQuery API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "cancel",
						Usage: "Manage cancel resources",
						Commands: []*cli.Command{
							{
								Name:  "cancel-job",
								Usage: "cancel-job cancel",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel-job...")
									return nil
								},
							},
						},
					},
					{
						Name:  "datasets",
						Usage: "Manage datasets resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert-dataset",
								Usage: "insert-dataset datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert-dataset...")
									return nil
								},
							},
							{
								Name:  "patch-dataset",
								Usage: "patch-dataset datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-dataset...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "undelete",
								Usage: "undelete datasets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing undelete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete",
						Usage: "Manage delete resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete delete",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "jobs",
						Usage: "Manage jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert-job",
								Usage: "insert-job jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert-job...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "models",
						Usage: "Manage models resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe models",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list models",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch-model",
								Usage: "patch-model models",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-model...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete models",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "queries",
						Usage: "Manage queries resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe queries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "query",
								Usage: "query queries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query...")
									return nil
								},
							},
						},
					},
					{
						Name:  "routines",
						Usage: "Manage routines resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe routines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert-routine",
								Usage: "insert-routine routines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert-routine...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update routines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete routines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list routines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "row-access-policies",
						Usage: "Manage row-access-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list row-access-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe row-access-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create row-access-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update row-access-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete row-access-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "batch-delete",
								Usage: "batch-delete row-access-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing batch-delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "service-account",
						Usage: "Manage service-account resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe service-account",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "tables",
						Usage: "Manage tables resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe tables",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert-table",
								Usage: "insert-table tables",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert-table...")
									return nil
								},
							},
							{
								Name:  "patch-table",
								Usage: "patch-table tables",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-table...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update tables",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete tables",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list tables",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
