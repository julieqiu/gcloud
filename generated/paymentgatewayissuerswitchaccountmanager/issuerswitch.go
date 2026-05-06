package issuerswitch

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud issuerswitch command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "issuerswitch",
		Usage: "manage Issuer switch API resources",
		Commands: []*cli.Command{
			{
				Name:  "accounts",
				Usage: "Manage accounts resources",
				Commands: []*cli.Command{
					{
						Name:  "batch-reconcile-balance",
						Usage: "batch-reconcile-balance accounts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-reconcile-balance...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account_manager", Usage: "The account_manager.", Required: true},
							&cli.StringFlag{Name: "account", Usage: "The account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/accountManagers/%s/accounts/%s", cmd.String("project"), cmd.String("account_manager"), cmd.String("account"))
							fmt.Printf("Executing describe on %s\n", name)
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
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "transactions",
				Usage: "Manage transactions resources",
				Commands: []*cli.Command{
					{
						Name:  "export",
						Usage: "export transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account_manager", Usage: "The account_manager.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/accountManagers/%s", cmd.String("project"), cmd.String("account_manager"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "batch-reconcile",
						Usage: "batch-reconcile transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-reconcile...")
							return nil
						},
					},
				},
			},
		},
	}
}
