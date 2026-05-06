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
				Name:  "logging",
				Usage: "manage Cloud Logging API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "buckets",
						Usage: "Manage buckets resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list buckets",
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
								Usage: "describe buckets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create buckets",
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
								Name:  "update-async",
								Usage: "update-async buckets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing update-async on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create buckets",
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
								Usage: "update buckets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete buckets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "undelete",
								Usage: "undelete buckets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing undelete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "cmek-settings",
						Usage: "Manage cmek-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe cmek-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/cmekSettings", cmd.String("project"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update cmek-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/cmekSettings", cmd.String("project"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "entries",
						Usage: "Manage entries resources",
						Commands: []*cli.Command{
							{
								Name:  "write",
								Usage: "write entries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing write on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list entries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "tail",
								Usage: "tail entries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing tail...")
									return nil
								},
							},
							{
								Name:  "copy",
								Usage: "copy entries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing copy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "exclusions",
						Usage: "Manage exclusions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list exclusions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe exclusions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "exclusion", Usage: "The exclusion.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/exclusions/%s", cmd.String("project"), cmd.String("exclusion"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create exclusions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update exclusions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "exclusion", Usage: "The exclusion.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/exclusions/%s", cmd.String("project"), cmd.String("exclusion"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete exclusions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "exclusion", Usage: "The exclusion.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/exclusions/%s", cmd.String("project"), cmd.String("exclusion"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "links",
						Usage: "Manage links resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create links",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete links",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
									&cli.StringFlag{Name: "link", Usage: "The link.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s/links/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"), cmd.String("link"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list links",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe links",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
									&cli.StringFlag{Name: "link", Usage: "The link.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s/links/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"), cmd.String("link"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "logs",
						Usage: "Manage logs resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete logs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list logs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "metrics",
						Usage: "Manage metrics resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list metrics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe metrics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create metrics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update metrics",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "metric", Usage: "The metric.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/metrics/%s", cmd.String("project"), cmd.String("metric"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete metrics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "monitored-resource-descriptors",
						Usage: "Manage monitored-resource-descriptors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list monitored-resource-descriptors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
									return nil
								},
							},
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
									return nil
								},
							},
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
									return nil
								},
							},
						},
					},
					{
						Name:  "settings",
						Usage: "Manage settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/settings", cmd.String("project"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/settings", cmd.String("project"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "sinks",
						Usage: "Manage sinks resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list sinks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe sinks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create sinks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update sinks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "sink", Usage: "The sink.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/sinks/%s", cmd.String("project"), cmd.String("sink"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete sinks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "views",
						Usage: "Manage views resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list views",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
									&cli.StringFlag{Name: "view", Usage: "The view.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"), cmd.String("view"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/buckets/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
									&cli.StringFlag{Name: "view", Usage: "The view.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"), cmd.String("view"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "bucket", Usage: "The bucket.", Required: true},
									&cli.StringFlag{Name: "view", Usage: "The view.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/buckets/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("bucket"), cmd.String("view"))
									fmt.Printf("Executing delete on %s\n", name)
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
