package firestore

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud firestore command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "firestore",
		Usage: "manage Cloud Firestore API resources",
		Commands: []*cli.Command{
			{
				Name:  "documents",
				Usage: "Manage documents resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
					{
						Name:  "batch-get",
						Usage: "batch-get documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-get...")
							return nil
						},
					},
					{
						Name:  "begin-transaction",
						Usage: "begin-transaction documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing begin-transaction...")
							return nil
						},
					},
					{
						Name:  "commit",
						Usage: "commit documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing commit...")
							return nil
						},
					},
					{
						Name:  "rollback",
						Usage: "rollback documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing rollback...")
							return nil
						},
					},
					{
						Name:  "run-query",
						Usage: "run-query documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing run-query...")
							return nil
						},
					},
					{
						Name:  "execute-pipeline",
						Usage: "execute-pipeline documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing execute-pipeline...")
							return nil
						},
					},
					{
						Name:  "run-aggregation-query",
						Usage: "run-aggregation-query documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing run-aggregation-query...")
							return nil
						},
					},
					{
						Name:  "partition-query",
						Usage: "partition-query documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing partition-query...")
							return nil
						},
					},
					{
						Name:  "write",
						Usage: "write documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing write...")
							return nil
						},
					},
					{
						Name:  "listen",
						Usage: "listen documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing listen...")
							return nil
						},
					},
					{
						Name:  "list-collection-ids",
						Usage: "list-collection-ids documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list-collection-ids...")
							return nil
						},
					},
					{
						Name:  "batch-write",
						Usage: "batch-write documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-write...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/operations/%s", cmd.String("project"), cmd.String("database"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/operations/%s", cmd.String("project"), cmd.String("database"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/operations/%s", cmd.String("project"), cmd.String("database"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
