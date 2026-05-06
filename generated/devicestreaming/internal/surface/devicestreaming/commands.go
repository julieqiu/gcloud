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

package devicestreaming

import (
	devicestreaming "cloud.google.com/go/devicestreaming/apiv1"
	"cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the devicestreaming command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "devicestreaming",
		Usage: "manage Device Streaming API resources",
		Commands: []*cli.Command{
			{
				Name:  "device-sessions",
				Usage: "Manage device-sessions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-session-id", Usage: "The ID to use for the DeviceSession, which will become the final.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := devicestreaming.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &devicestreamingpb.CreateDeviceSessionRequest{
								Parent:          parent,
								DeviceSessionId: cmd.String("device-session-id"),
							}

							resp, err := client.CreateDeviceSession(ctx, req)
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
						Name:  "list",
						Usage: "list device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If specified, responses will be filtered by the given filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of DeviceSessions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A continuation token for paging.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := devicestreaming.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &devicestreamingpb.ListDeviceSessionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeviceSessions(ctx, req)
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
						Usage: "describe device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-session", Usage: "The ID of the device session.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deviceSessions/%s", cmd.String("project"), cmd.String("device-session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := devicestreaming.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &devicestreamingpb.GetDeviceSessionRequest{
								Name: name,
							}

							resp, err := client.GetDeviceSession(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-session", Usage: "The ID of the device session.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deviceSessions/%s", cmd.String("project"), cmd.String("device-session"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelDeviceSession on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := devicestreaming.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &devicestreamingpb.CancelDeviceSessionRequest{
								Name: name,
							}

							if err := client.CancelDeviceSession(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-session", Usage: "The ID of the device session.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "device_session.name" not yet supported.
							device_session_name := fmt.Sprintf("projects/%s/deviceSessions/%s", cmd.String("project"), cmd.String("device-session"))
							fmt.Printf("Executing update on %s\n", device_session_name)
							return nil
						},
					},
				},
			},
		},
	}
}
