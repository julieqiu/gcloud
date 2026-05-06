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

package eventarc

import (
	eventarc "cloud.google.com/go/eventarc/apiv1"
	"cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the eventarc command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "eventarc",
		Usage: "manage Eventarc API resources",
		Commands: []*cli.Command{
			{
				Name:  "channel-connections",
				Usage: "Manage channel-connections resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe channel-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel-connection", Usage: "The ID of the channel connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/channelConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetChannelConnectionRequest{
								Name: name,
							}

							resp, err := client.GetChannelConnection(ctx, req)
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
						Usage: "list channel-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of channel connections to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token` field in a.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListChannelConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListChannelConnections(ctx, req)
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
						Name:  "create",
						Usage: "create channel-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel-connection-id", Usage: "The user-provided ID to be assigned to the channel connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreateChannelConnectionRequest{
								Parent:              parent,
								ChannelConnectionId: cmd.String("channel-connection-id"),
							}

							op, err := client.CreateChannelConnection(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "delete",
						Usage: "delete channel-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel-connection", Usage: "The ID of the channel connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/channelConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeleteChannelConnectionRequest{
								Name: name,
							}

							op, err := client.DeleteChannelConnection(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "channels",
				Usage: "Manage channels resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel", Usage: "The ID of the channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/channels/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetChannelRequest{
								Name: name,
							}

							resp, err := client.GetChannel(ctx, req)
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
						Usage: "list channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of channels to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token` field in a.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListChannelsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListChannels(ctx, req)
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
						Name:  "create",
						Usage: "create channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel-id", Usage: "The user-provided ID to be assigned to the channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreateChannelRequest{
								Parent:       parent,
								ChannelId:    cmd.String("channel-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateChannel(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel", Usage: "The ID of the channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "channel.name" not yet supported.
							channel_name := fmt.Sprintf("projects/%s/locations/%s/channels/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel"))
							fmt.Printf("Executing update on %s\n", channel_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "channel", Usage: "The ID of the channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/channels/%s", cmd.String("project"), cmd.String("location"), cmd.String("channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeleteChannelRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteChannel(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "enrollments",
				Usage: "Manage enrollments resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe enrollments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "enrollment", Usage: "The ID of the enrollment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/enrollments/%s", cmd.String("project"), cmd.String("location"), cmd.String("enrollment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetEnrollmentRequest{
								Name: name,
							}

							resp, err := client.GetEnrollment(ctx, req)
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
						Usage: "list enrollments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter field that the list request will filter on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListEnrollmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListEnrollments(ctx, req)
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
						Name:  "create",
						Usage: "create enrollments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "enrollment-id", Usage: "The user-provided ID to be assigned to the Enrollment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreateEnrollmentRequest{
								Parent:       parent,
								EnrollmentId: cmd.String("enrollment-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateEnrollment(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update enrollments",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the Enrollment is not found, a new Enrollment.", Required: false},
							&cli.StringFlag{Name: "enrollment", Usage: "The ID of the enrollment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "enrollment.name" not yet supported.
							enrollment_name := fmt.Sprintf("projects/%s/locations/%s/enrollments/%s", cmd.String("project"), cmd.String("location"), cmd.String("enrollment"))
							fmt.Printf("Executing update on %s\n", enrollment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete enrollments",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the Enrollment is not found, the request will.", Required: false},
							&cli.StringFlag{Name: "enrollment", Usage: "The ID of the enrollment.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If provided, the Enrollment will only be deleted if the etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/enrollments/%s", cmd.String("project"), cmd.String("location"), cmd.String("enrollment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeleteEnrollmentRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteEnrollment(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "google-api-sources",
				Usage: "Manage google-api-sources resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe google-api-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "google-api-source", Usage: "The ID of the google api source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/googleApiSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("google-api-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetGoogleApiSourceRequest{
								Name: name,
							}

							resp, err := client.GetGoogleApiSource(ctx, req)
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
						Usage: "list google-api-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter field that the list request will filter on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListGoogleApiSourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListGoogleApiSources(ctx, req)
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
						Name:  "create",
						Usage: "create google-api-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "google-api-source-id", Usage: "The user-provided ID to be assigned to the GoogleApiSource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreateGoogleApiSourceRequest{
								Parent:            parent,
								GoogleApiSourceId: cmd.String("google-api-source-id"),
								ValidateOnly:      cmd.Bool("validate-only"),
							}

							op, err := client.CreateGoogleApiSource(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update google-api-sources",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the GoogleApiSource is not found, a new.", Required: false},
							&cli.StringFlag{Name: "google-api-source", Usage: "The ID of the google api source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "google_api_source.name" not yet supported.
							google_api_source_name := fmt.Sprintf("projects/%s/locations/%s/googleApiSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("google-api-source"))
							fmt.Printf("Executing update on %s\n", google_api_source_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete google-api-sources",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the MessageBus is not found, the request will.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "If provided, the MessageBus will only be deleted if the etag.", Required: false},
							&cli.StringFlag{Name: "google-api-source", Usage: "The ID of the google api source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/googleApiSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("google-api-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeleteGoogleApiSourceRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteGoogleApiSource(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "google-channel-config",
				Usage: "Manage google-channel-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe google-channel-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/googleChannelConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetGoogleChannelConfigRequest{
								Name: name,
							}

							resp, err := client.GetGoogleChannelConfig(ctx, req)
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
						Usage: "update google-channel-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "google_channel_config.name" not yet supported.
							google_channel_config_name := fmt.Sprintf("projects/%s/locations/%s/googleChannelConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", google_channel_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
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
						Name:  "describe",
						Usage: "describe message-buses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "message-buse", Usage: "The ID of the message buse.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", cmd.String("project"), cmd.String("location"), cmd.String("message-buse"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetMessageBusRequest{
								Name: name,
							}

							resp, err := client.GetMessageBus(ctx, req)
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
						Usage: "list message-buses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter field that the list request will filter on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListMessageBusesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMessageBuses(ctx, req)
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
						Name:  "list",
						Usage: "list message-buses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "message-buse", Usage: "The ID of the message buse.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", cmd.String("project"), cmd.String("location"), cmd.String("message-buse"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListMessageBusEnrollmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.ListMessageBusEnrollments(ctx, req)
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
						Name:  "create",
						Usage: "create message-buses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "message-bus-id", Usage: "The user-provided ID to be assigned to the MessageBus.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreateMessageBusRequest{
								Parent:       parent,
								MessageBusId: cmd.String("message-bus-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateMessageBus(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update message-buses",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the MessageBus is not found, a new MessageBus.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "message-buse", Usage: "The ID of the message buse.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "message_bus.name" not yet supported.
							message_bus_name := fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", cmd.String("project"), cmd.String("location"), cmd.String("message-buse"))
							fmt.Printf("Executing update on %s\n", message_bus_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete message-buses",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the MessageBus is not found, the request will.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "If provided, the MessageBus will only be deleted if the etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "message-buse", Usage: "The ID of the message buse.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", cmd.String("project"), cmd.String("location"), cmd.String("message-buse"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeleteMessageBusRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteMessageBus(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "pipelines",
				Usage: "Manage pipelines resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline", Usage: "The ID of the pipeline.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetPipelineRequest{
								Name: name,
							}

							resp, err := client.GetPipeline(ctx, req)
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
						Usage: "list pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter field that the list request will filter on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListPipelinesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPipelines(ctx, req)
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
						Name:  "create",
						Usage: "create pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline-id", Usage: "The user-provided ID to be assigned to the Pipeline.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreatePipelineRequest{
								Parent:       parent,
								PipelineId:   cmd.String("pipeline-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreatePipeline(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update pipelines",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the Pipeline is not found, a new Pipeline.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline", Usage: "The ID of the pipeline.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "pipeline.name" not yet supported.
							pipeline_name := fmt.Sprintf("projects/%s/locations/%s/pipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("pipeline"))
							fmt.Printf("Executing update on %s\n", pipeline_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete pipelines",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the Pipeline is not found, the request will.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "If provided, the Pipeline will only be deleted if the etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pipeline", Usage: "The ID of the pipeline.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeletePipelineRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeletePipeline(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "providers",
				Usage: "Manage providers resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe providers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "provider", Usage: "The ID of the provider.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/providers/%s", cmd.String("project"), cmd.String("location"), cmd.String("provider"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetProviderRequest{
								Name: name,
							}

							resp, err := client.GetProvider(ctx, req)
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
						Usage: "list providers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter field that the list request will filter on.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of providers to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token` field in a.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListProvidersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListProviders(ctx, req)
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
				},
			},
			{
				Name:  "triggers",
				Usage: "Manage triggers resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger", Usage: "The ID of the trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/triggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("trigger"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.GetTriggerRequest{
								Name: name,
							}

							resp, err := client.GetTrigger(ctx, req)
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
						Usage: "list triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter field.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sorting order of the resources returned.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of triggers to return on each page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token; provide the value from the `next_page_token` field in a.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.ListTriggersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListTriggers(ctx, req)
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
						Name:  "create",
						Usage: "create triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger-id", Usage: "The user-provided ID to be assigned to the trigger.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.CreateTriggerRequest{
								Parent:       parent,
								TriggerId:    cmd.String("trigger-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateTrigger(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update triggers",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the trigger is not found, a new trigger will be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger", Usage: "The ID of the trigger.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "trigger.name" not yet supported.
							trigger_name := fmt.Sprintf("projects/%s/locations/%s/triggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("trigger"))
							fmt.Printf("Executing update on %s\n", trigger_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete triggers",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the trigger is not found, the request will succeed.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "If provided, the trigger will only be deleted if the etag matches the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger", Usage: "The ID of the trigger.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/triggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("trigger"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := eventarc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &eventarcpb.DeleteTriggerRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteTrigger(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger", Usage: "The ID of the trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/triggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("trigger"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger", Usage: "The ID of the trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/triggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("trigger"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger", Usage: "The ID of the trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/triggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("trigger"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
		},
	}
}
