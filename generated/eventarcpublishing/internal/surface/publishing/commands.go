// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package publishing

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
)

// Command returns the eventarcpublishing command tree for inclusion under the gcloud root.
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel-connection", Usage: "The ID of the channel connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "text-events", Usage: "The text representation of events to publish.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							channel_connection := fmt.Sprintf("projects/%s/locations/%s/channelConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel-connection"))
							fmt.Printf("Executing publish-events on %s\n", channel_connection)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel", Usage: "The ID of the channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "text-events", Usage: "The text representation of events to publish.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							channel := fmt.Sprintf("projects/%s/locations/%s/channels/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel"))
							fmt.Printf("Executing publish-events on %s\n", channel)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "message-bus", Usage: "The full name of the message bus to publish events to.", Required: true},
							&cli.StringFlag{Name: "message-buse", Usage: "The ID of the message buse.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							message_bus := fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", cmd.String("project"), cmd.String("location"), cmd.String("message-buse"))
							fmt.Printf("Executing publish on %s\n", message_bus)
							return nil
						},
					},
				},
			},
		},
	}
}
