package eventarcpublishing

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud eventarcpublishing command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "eventarcpublishing",
		Usage: "manage Eventarc Publishing API resources",
		Commands: []*cli.Command{
			{
				Name:  "channel-connections",
				Usage: "Manage channel-connections resources",
				Commands: []*cli.Command{
					{
						Name:  "publish-events",
						Usage: "publish-events channel-connections",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing publish-events...")
							return nil
						},
					},
				},
			},
			{
				Name:  "channels",
				Usage: "Manage channels resources",
				Commands: []*cli.Command{
					{
						Name:  "publish-events",
						Usage: "publish-events channels",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing publish-events...")
							return nil
						},
					},
				},
			},
			{
				Name:  "message-buses",
				Usage: "Manage message-buses resources",
				Commands: []*cli.Command{
					{
						Name:  "publish",
						Usage: "publish message-buses",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing publish...")
							return nil
						},
					},
				},
			},
		},
	}
}
