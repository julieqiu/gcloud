package universalledger

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud universalledger command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "universalledger",
		Usage: "manage Universal Ledger API resources",
		Commands: []*cli.Command{
			{
				Name:  "endpoints",
				Usage: "Manage endpoints resources",
				Commands: []*cli.Command{
					{
						Name:  "submit-transaction",
						Usage: "submit-transaction endpoints",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing submit-transaction...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list endpoints",
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
						Usage: "describe endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "endpoint", Usage: "The endpoint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "submit-operational-transaction",
						Usage: "submit-operational-transaction endpoints",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing submit-operational-transaction...")
							return nil
						},
					},
					{
						Name:  "query-transaction-state",
						Usage: "query-transaction-state endpoints",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing query-transaction-state...")
							return nil
						},
					},
					{
						Name:  "query-account",
						Usage: "query-account endpoints",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing query-account...")
							return nil
						},
					},
				},
			},
		},
	}
}
