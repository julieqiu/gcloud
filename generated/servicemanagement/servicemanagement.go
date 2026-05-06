package servicemanagement

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud servicemanagement command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "servicemanagement",
		Usage: "manage Service Management API resources",
		Commands: []*cli.Command{
			{
				Name:  "configs",
				Usage: "Manage configs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "submit",
						Usage: "submit configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing submit...")
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
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
				},
			},
			{
				Name:  "rollouts",
				Usage: "Manage rollouts resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list rollouts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe rollouts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create rollouts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
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
						Name:  "create",
						Usage: "create services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
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
					{
						Name:  "undelete",
						Usage: "undelete services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing undelete...")
							return nil
						},
					},
					{
						Name:  "generate-config-report",
						Usage: "generate-config-report services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-config-report...")
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
