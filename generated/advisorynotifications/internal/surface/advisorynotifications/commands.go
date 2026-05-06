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

package advisorynotifications

import (
	advisorynotifications "cloud.google.com/go/advisorynotifications/apiv1"
	"cloud.google.com/go/advisorynotifications/apiv1/advisorynotificationspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the advisorynotifications command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "advisorynotifications",
		Usage: "manage Advisory Notifications API resources",
		Commands: []*cli.Command{
			{
				Name:  "notifications",
				Usage: "Manage notifications resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list notifications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "language-code", Usage: "ISO code for requested localization language.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of notifications to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "view", Usage: "Specifies which parts of the notification resource should be returned.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := advisorynotifications.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &advisorynotificationspb.ListNotificationsRequest{
								Parent:       parent,
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								View:         advisorynotificationspb.NotificationView(advisorynotificationspb.NotificationView_value[cmd.String("view")]),
								LanguageCode: cmd.String("language-code"),
							}

							limit := cmd.Int("limit")
							it := client.ListNotifications(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe notifications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "language-code", Usage: "ISO code for requested localization language.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notification", Usage: "The ID of the notification.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/notifications/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := advisorynotifications.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &advisorynotificationspb.GetNotificationRequest{
								Name:         name,
								LanguageCode: cmd.String("language-code"),
							}

							resp, err := client.GetNotification(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "settings",
				Usage: "Manage settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/settings", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := advisorynotifications.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &advisorynotificationspb.GetSettingsRequest{
								Name: name,
							}

							resp, err := client.GetSettings(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "settings.name" not yet supported.
							settings_name := fmt.Sprintf("organizations/%s/locations/%s/settings", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", settings_name)
							return nil
						},
					},
				},
			},
		},
	}
}
