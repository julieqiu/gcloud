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
				Name:  "storagetransfer",
				Usage: "manage Storage Transfer API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "agent-pools",
						Usage: "Manage agent-pools resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create agent-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_id", Usage: "The project_id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project_id"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update agent-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_id", Usage: "The project_id.", Required: true},
									&cli.StringFlag{Name: "agent_pool_id", Usage: "The agent_pool_id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/agentPools/%s", cmd.String("project_id"), cmd.String("agent_pool_id"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe agent-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list agent-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete agent-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "google-service-accounts",
						Usage: "Manage google-service-accounts resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe google-service-accounts",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "transfer-jobs",
						Usage: "Manage transfer-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create transfer-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update transfer-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe transfer-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list transfer-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "run",
								Usage: "run transfer-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing run...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete transfer-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "transfer-operations",
						Usage: "Manage transfer-operations resources",
						Commands: []*cli.Command{
							{
								Name:  "pause",
								Usage: "pause transfer-operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing pause...")
									return nil
								},
							},
							{
								Name:  "resume",
								Usage: "resume transfer-operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resume...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list transfer-operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe transfer-operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel transfer-operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
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
