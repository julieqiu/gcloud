package spanner

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud spanner command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "spanner",
		Usage: "manage Cloud Spanner API resources",
		Commands: []*cli.Command{
			{
				Name:  "databases",
				Usage: "Manage databases resources",
				Commands: []*cli.Command{
					{
						Name:  "cache-update",
						Usage: "cache-update databases",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing cache-update...")
							return nil
						},
					},
				},
			},
			{
				Name:  "sessions",
				Usage: "Manage sessions resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "batch-create",
						Usage: "batch-create sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing batch-create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "execute-sql",
						Usage: "execute-sql sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing execute-sql...")
							return nil
						},
					},
					{
						Name:  "execute-streaming-sql",
						Usage: "execute-streaming-sql sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing execute-streaming-sql...")
							return nil
						},
					},
					{
						Name:  "execute-batch-dml",
						Usage: "execute-batch-dml sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing execute-batch-dml...")
							return nil
						},
					},
					{
						Name:  "read",
						Usage: "read sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing read...")
							return nil
						},
					},
					{
						Name:  "streaming-read",
						Usage: "streaming-read sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing streaming-read...")
							return nil
						},
					},
					{
						Name:  "begin-transaction",
						Usage: "begin-transaction sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing begin-transaction...")
							return nil
						},
					},
					{
						Name:  "commit",
						Usage: "commit sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing commit...")
							return nil
						},
					},
					{
						Name:  "rollback",
						Usage: "rollback sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing rollback...")
							return nil
						},
					},
					{
						Name:  "partition-query",
						Usage: "partition-query sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing partition-query...")
							return nil
						},
					},
					{
						Name:  "partition-read",
						Usage: "partition-read sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing partition-read...")
							return nil
						},
					},
					{
						Name:  "batch-write",
						Usage: "batch-write sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-write...")
							return nil
						},
					},
				},
			},
		},
	}
}
