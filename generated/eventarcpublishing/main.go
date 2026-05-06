package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "eventarcpublishing",
				Usage: "manage Eventarc Publishing API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
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
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
