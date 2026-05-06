package cloudtasks

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud cloudtasks command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudtasks",
		Usage: "manage Cloud Tasks API resources",
		Commands: []*cli.Command{
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "queues",
				Usage: "Manage queues resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list queues",
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
						Usage: "describe queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create queues",
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
						Usage: "update queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "purge",
						Usage: "purge queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing purge on %s\n", name)
							return nil
						},
					},
					{
						Name:  "pause",
						Usage: "pause queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing pause on %s\n", name)
							return nil
						},
					},
					{
						Name:  "resume",
						Usage: "resume queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing resume on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "tasks",
				Usage: "Manage tasks resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "run",
						Usage: "run tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							fmt.Printf("Executing run on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
