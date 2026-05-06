package policytroubleshooter

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud policytroubleshooter command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "policytroubleshooter",
		Usage: "manage Policy Troubleshooter API resources",
		Commands: []*cli.Command{
			{
				Name:  "iam",
				Usage: "Manage iam resources",
				Commands: []*cli.Command{
					{
						Name:  "troubleshoot",
						Usage: "troubleshoot iam",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing troubleshoot...")
							return nil
						},
					},
				},
			},
		},
	}
}
