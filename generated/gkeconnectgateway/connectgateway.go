package connectgateway

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud connectgateway command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "connectgateway",
		Usage: "manage Connect Gateway API resources",
		Commands: []*cli.Command{
			{
				Name:  "memberships",
				Usage: "Manage memberships resources",
				Commands: []*cli.Command{
					{
						Name:  "generate-credentials",
						Usage: "generate-credentials memberships",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-credentials...")
							return nil
						},
					},
				},
			},
		},
	}
}
