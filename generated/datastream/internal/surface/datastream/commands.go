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

package datastream

import (
	datastream "cloud.google.com/go/datastream/apiv1"
	"cloud.google.com/go/datastream/apiv1/datastreampb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the datastream command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "datastream",
		Usage: "manage Datastream API resources",
		Commands: []*cli.Command{
			{
				Name:  "connection-profiles",
				Usage: "Manage connection-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of connection profiles to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListConnectionProfiles` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.ListConnectionProfilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListConnectionProfiles(ctx, req)
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
						Usage: "describe connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.GetConnectionProfileRequest{
								Name: name,
							}

							resp, err := client.GetConnectionProfile(ctx, req)
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
						Usage: "create connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile-id", Usage: "The connection profile identifier.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Create the connection profile without validating it.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the connection profile, but don't create any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.CreateConnectionProfileRequest{
								Parent:              parent,
								ConnectionProfileId: cmd.String("connection-profile-id"),
								RequestId:           cmd.String("request-id"),
								ValidateOnly:        cmd.Bool("validate-only"),
								Force:               cmd.Bool("force"),
							}

							op, err := client.CreateConnectionProfile(ctx, req)
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
						Usage: "update connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Update the connection profile without validating it.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the connection profile, but don't update any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "connection_profile.name" not yet supported.
							connection_profile_name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							fmt.Printf("Executing update on %s\n", connection_profile_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-profile", Usage: "The ID of the connection profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConnectionProfile %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.DeleteConnectionProfileRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteConnectionProfile(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "discover",
						Usage: "discover connection-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.DiscoverConnectionProfileRequest{
								Parent: parent,
							}

							resp, err := client.DiscoverConnectionProfile(ctx, req)
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
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch-static-ips",
						Usage: "fetch-static-ips locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Ips to return, will likely not be specified.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListStaticIps` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.FetchStaticIpsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.FetchStaticIps(ctx, req)
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
				Name:  "objects",
				Usage: "Manage objects resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe objects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "object", Usage: "The ID of the object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/streams/%s/objects/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"), cmd.String("object"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.GetStreamObjectRequest{
								Name: name,
							}

							resp, err := client.GetStreamObject(ctx, req)
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
						Name:  "lookup",
						Usage: "lookup objects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.LookupStreamObjectRequest{
								Parent: parent,
							}

							resp, err := client.LookupStreamObject(ctx, req)
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
						Usage: "list objects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of objects to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListStreamObjectsRequest` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.ListStreamObjectsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListStreamObjects(ctx, req)
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
						Name:  "start-backfill-job",
						Usage: "start-backfill-job objects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "object", Usage: "The ID of the object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							object := fmt.Sprintf("projects/%s/locations/%s/streams/%s/objects/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"), cmd.String("object"))
							fmt.Printf("Executing start-backfill-job on %s\n", object)
							return nil
						},
					},

					{
						Name:  "stop-backfill-job",
						Usage: "stop-backfill-job objects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "object", Usage: "The ID of the object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							object := fmt.Sprintf("projects/%s/locations/%s/streams/%s/objects/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"), cmd.String("object"))
							fmt.Printf("Executing stop-backfill-job on %s\n", object)
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
				Name:  "private-connections",
				Usage: "Manage private-connections resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create private-connections",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, will skip validations.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection-id", Usage: "The private connectivity identifier.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "When supplied with PSC Interface config, will get/create the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.CreatePrivateConnectionRequest{
								Parent:              parent,
								PrivateConnectionId: cmd.String("private-connection-id"),
								RequestId:           cmd.String("request-id"),
								Force:               cmd.Bool("force"),
								ValidateOnly:        cmd.Bool("validate-only"),
							}

							op, err := client.CreatePrivateConnection(ctx, req)
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
						Name:  "describe",
						Usage: "describe private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.GetPrivateConnectionRequest{
								Name: name,
							}

							resp, err := client.GetPrivateConnection(ctx, req)
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
						Usage: "list private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of private connectivity configurations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListPrivateConnections` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.ListPrivateConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPrivateConnections(ctx, req)
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
						Name:  "delete",
						Usage: "delete private-connections",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any child routes that belong to this.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePrivateConnection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.DeletePrivateConnectionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeletePrivateConnection(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "routes",
				Usage: "Manage routes resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "route-id", Usage: "The Route identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.CreateRouteRequest{
								Parent:    parent,
								RouteId:   cmd.String("route-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateRoute(ctx, req)
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
						Name:  "describe",
						Usage: "describe routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route", Usage: "The ID of the route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s/routes/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"), cmd.String("route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.GetRouteRequest{
								Name: name,
							}

							resp, err := client.GetRoute(ctx, req)
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
						Usage: "list routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Routes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListRoutes` call.", Required: false},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.ListRoutesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRoutes(ctx, req)
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
						Name:  "delete",
						Usage: "delete routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "private-connection", Usage: "The ID of the private connection.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "route", Usage: "The ID of the route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s/routes/%s", cmd.String("project"), cmd.String("location"), cmd.String("private-connection"), cmd.String("route"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRoute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.DeleteRouteRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteRoute(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "streams",
				Usage: "Manage streams resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of streams to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListStreams` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.ListStreamsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListStreams(ctx, req)
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
						Usage: "describe streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.GetStreamRequest{
								Name: name,
							}

							resp, err := client.GetStream(ctx, req)
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
						Usage: "create streams",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "Create the stream without validating it.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "stream-id", Usage: "The stream identifier.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the stream, but don't create any resources.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.CreateStreamRequest{
								Parent:       parent,
								StreamId:     cmd.String("stream-id"),
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
								Force:        cmd.Bool("force"),
							}

							op, err := client.CreateStream(ctx, req)
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
						Usage: "update streams",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "Update the stream without validating it.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the stream with the changes, without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "stream.name" not yet supported.
							stream_name := fmt.Sprintf("projects/%s/locations/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"))
							fmt.Printf("Executing update on %s\n", stream_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete streams",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteStream %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.DeleteStreamRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteStream(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "run",
						Usage: "run streams",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "Update the stream without validating it.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream", Usage: "The ID of the stream.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("stream"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datastream.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datastreampb.RunStreamRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.RunStream(ctx, req)
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
		},
	}
}
