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
				Name:  "complaint-transactions",
				Usage: "Manage complaint-transactions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list complaint-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export complaint-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
							return nil
						},
					},
				},
			},
			{
				Name:  "complaints",
				Usage: "Manage complaints resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create complaints",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "resolve",
						Usage: "resolve complaints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "complaint", Usage: "The complaint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/complaints/%s", cmd.String("project"), cmd.String("complaint"))
							fmt.Printf("Executing resolve on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "disputes",
				Usage: "Manage disputes resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create disputes",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "resolve",
						Usage: "resolve disputes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dispute", Usage: "The dispute.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/disputes/%s", cmd.String("project"), cmd.String("dispute"))
							fmt.Printf("Executing resolve on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "financial-transactions",
				Usage: "Manage financial-transactions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list financial-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export financial-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
							return nil
						},
					},
				},
			},
			{
				Name:  "issuer-participants",
				Usage: "Manage issuer-participants resources",
				Commands: []*cli.Command{
					{
						Name:  "fetch",
						Usage: "fetch issuer-participants",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update issuer-participants",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "activate",
						Usage: "activate issuer-participants",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing activate...")
							return nil
						},
					},
					{
						Name:  "deactivate",
						Usage: "deactivate issuer-participants",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing deactivate...")
							return nil
						},
					},
					{
						Name:  "mobile-number-changed",
						Usage: "mobile-number-changed issuer-participants",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing mobile-number-changed...")
							return nil
						},
					},
				},
			},
			{
				Name:  "mandate-transactions",
				Usage: "Manage mandate-transactions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list mandate-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export mandate-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
							return nil
						},
					},
				},
			},
			{
				Name:  "metadata",
				Usage: "Manage metadata resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list metadata",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s", cmd.String("project"), cmd.String("rule"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "metadata-transactions",
				Usage: "Manage metadata-transactions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list metadata-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export metadata-transactions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
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
				Name:  "rules",
				Usage: "Manage rules resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "values",
				Usage: "Manage values resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
							&cli.StringFlag{Name: "metadata", Usage: "The metadata.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "batch-create",
						Usage: "batch-create values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
							&cli.StringFlag{Name: "metadata", Usage: "The metadata.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
							fmt.Printf("Executing batch-create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "batch-delete",
						Usage: "batch-delete values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
							&cli.StringFlag{Name: "metadata", Usage: "The metadata.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
							fmt.Printf("Executing batch-delete on %s\n", parent)
							return nil
						},
					},
				},
			},
		},
	}
}
