package appengine

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud appengine command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "appengine",
		Usage: "manage App Engine Admin API resources",
		Commands: []*cli.Command{
			{
				Name:  "apps",
				Usage: "Manage apps resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe apps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create apps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update apps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "repair",
						Usage: "repair apps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing repair...")
							return nil
						},
					},
				},
			},
			{
				Name:  "authorized-certificates",
				Usage: "Manage authorized-certificates resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list authorized-certificates",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe authorized-certificates",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create authorized-certificates",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update authorized-certificates",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete authorized-certificates",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
				},
			},
			{
				Name:  "authorized-domains",
				Usage: "Manage authorized-domains resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list authorized-domains",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "domain-mappings",
				Usage: "Manage domain-mappings resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list domain-mappings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe domain-mappings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create domain-mappings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update domain-mappings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete domain-mappings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
				},
			},
			{
				Name:  "ingress-rules",
				Usage: "Manage ingress-rules resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list ingress-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "batch-update",
						Usage: "batch-update ingress-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-update...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create ingress-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe ingress-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update ingress-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete ingress-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
				},
			},
			{
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
					{
						Name:  "debug",
						Usage: "debug instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing debug...")
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
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
				},
			},
			{
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
				},
			},
		},
	}
}
