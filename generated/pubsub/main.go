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
				Name:  "pubsub",
				Usage: "manage Cloud Pub/Sub API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "schemas",
						Usage: "Manage schemas resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create schemas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list schemas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list schemas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "commit",
								Usage: "commit schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
									fmt.Printf("Executing commit on %s\n", name)
									return nil
								},
							},
							{
								Name:  "rollback",
								Usage: "rollback schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
									fmt.Printf("Executing rollback on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "validate",
								Usage: "validate schemas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing validate on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "validate-message",
								Usage: "validate-message schemas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing validate-message on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "snapshots",
						Usage: "Manage snapshots resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "create-snapshot",
								Usage: "create-snapshot snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/snapshots/%s", cmd.String("project"), cmd.String("snapshot"))
									fmt.Printf("Executing create-snapshot on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/snapshots/%s", cmd.String("project"), cmd.String("snapshot"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "subscriptions",
						Usage: "Manage subscriptions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "detach",
								Usage: "detach subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detach...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "create-subscription",
								Usage: "create-subscription subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create-subscription...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update subscriptions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "subscription", Usage: "The subscription.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "modify-ack-deadline",
								Usage: "modify-ack-deadline subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing modify-ack-deadline...")
									return nil
								},
							},
							{
								Name:  "acknowledge",
								Usage: "acknowledge subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing acknowledge...")
									return nil
								},
							},
							{
								Name:  "pull",
								Usage: "pull subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing pull...")
									return nil
								},
							},
							{
								Name:  "modify-push-config",
								Usage: "modify-push-config subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing modify-push-config...")
									return nil
								},
							},
							{
								Name:  "seek",
								Usage: "seek subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing seek...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions subscriptions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "topics",
						Usage: "Manage topics resources",
						Commands: []*cli.Command{
							{
								Name:  "create-topic",
								Usage: "create-topic topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create-topic...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update topics",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "topic", Usage: "The topic.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "publish",
								Usage: "publish topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing publish...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy topics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
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
