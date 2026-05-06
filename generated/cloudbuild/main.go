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
				Name:  "cloudbuild",
				Usage: "manage Cloud Build API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "connections",
						Usage: "Manage connections resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create connections",
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
								Name:  "describe",
								Usage: "describe connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list connections",
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
								Name:  "update",
								Usage: "update connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-linkable-repositories",
								Usage: "fetch-linkable-repositories connections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-linkable-repositories...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy connections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy connections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions connections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									fmt.Printf("Executing cancel on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "repositories",
						Usage: "Manage repositories resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "batch-create",
								Usage: "batch-create repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing batch-create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repository"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("repository"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "access-read-write-token",
								Usage: "access-read-write-token repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing access-read-write-token...")
									return nil
								},
							},
							{
								Name:  "access-read-token",
								Usage: "access-read-token repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing access-read-token...")
									return nil
								},
							},
							{
								Name:  "fetch-git-refs",
								Usage: "fetch-git-refs repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-git-refs...")
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
