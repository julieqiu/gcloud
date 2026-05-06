package identitytoolkit

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud identitytoolkit command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "identitytoolkit",
		Usage: "manage Identity Toolkit API resources",
		Commands: []*cli.Command{
			{
				Name:  "mfa-enrollment",
				Usage: "Manage mfa-enrollment resources",
				Commands: []*cli.Command{
					{
						Name:  "finalize",
						Usage: "finalize mfa-enrollment",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing finalize...")
							return nil
						},
					},
					{
						Name:  "start",
						Usage: "start mfa-enrollment",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing start...")
							return nil
						},
					},
					{
						Name:  "withdraw",
						Usage: "withdraw mfa-enrollment",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing withdraw...")
							return nil
						},
					},
				},
			},
			{
				Name:  "mfa-sign-in",
				Usage: "Manage mfa-sign-in resources",
				Commands: []*cli.Command{
					{
						Name:  "finalize",
						Usage: "finalize mfa-sign-in",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing finalize...")
							return nil
						},
					},
					{
						Name:  "start",
						Usage: "start mfa-sign-in",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing start...")
							return nil
						},
					},
				},
			},
		},
	}
}
