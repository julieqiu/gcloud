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

package networkconnectivity

import (
	networkconnectivity "cloud.google.com/go/networkconnectivity/apiv1"
	"cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the networkconnectivity command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "networkconnectivity",
		Usage: "manage Network Connectivity API resources",
		Commands: []*cli.Command{
			{
				Name:  "destinations",
				Usage: "Manage destinations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sort order of the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results listed per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If `true`, allow partial responses for multi-regional aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListDestinationsRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								Filter:               cmd.String("filter"),
								OrderBy:              cmd.String("order-by"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListDestinations(ctx, req)
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
						Usage: "describe destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination", Usage: "The ID of the destination.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"), cmd.String("destination"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetDestinationRequest{
								Name: name,
							}

							resp, err := client.GetDestination(ctx, req)
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
						Usage: "create destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination-id", Usage: "The ID to use for the `Destination` resource, which becomes the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateDestinationRequest{
								Parent:        parent,
								DestinationId: cmd.String("destination-id"),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.CreateDestination(ctx, req)
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
						Usage: "update destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination", Usage: "The ID of the destination.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "destination.name" not yet supported.
							destination_name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"), cmd.String("destination"))
							fmt.Printf("Executing update on %s\n", destination_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination", Usage: "The ID of the destination.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag is computed by the server, and might be sent with update.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"), cmd.String("destination"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDestination %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteDestinationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      cmd.String("etag"),
							}

							op, err := client.DeleteDestination(ctx, req)
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
				Name:  "groups",
				Usage: "Manage groups resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/groups/%s", cmd.String("project"), cmd.String("hub"), cmd.String("group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetGroupRequest{
								Name: name,
							}

							resp, err := client.GetGroup(ctx, req)
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
						Usage: "list groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGroups(ctx, req)
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
						Name:  "update",
						Usage: "update groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "group.name" not yet supported.
							group_name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/groups/%s", cmd.String("project"), cmd.String("hub"), cmd.String("group"))
							fmt.Printf("Executing update on %s\n", group_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "hubs",
				Usage: "Manage hubs resources",
				Commands: []*cli.Command{

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListHubsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListHubs(ctx, req)
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
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetHubRequest{
								Name: name,
							}

							resp, err := client.GetHub(ctx, req)
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
						Usage: "create hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub-id", Usage: "A unique identifier for the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateHubRequest{
								Parent:    parent,
								HubId:     cmd.String("hub-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateHub(ctx, req)
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
						Usage: "update hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "hub.name" not yet supported.
							hub_name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing update on %s\n", hub_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteHub %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteHubRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteHub(ctx, req)
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
						Name:  "list",
						Usage: "list hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by name or create_time.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "spoke-locations", Usage: "A list of locations.", Required: false},
							&cli.StringFlag{Name: "view", Usage: "The view of the spoke to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListHubSpokesRequest{
								Name:           name,
								SpokeLocations: cmd.StringSlice("spoke-locations"),
								PageSize:       int32(cmd.Int("page-size")),
								PageToken:      cmd.String("page-token"),
								Filter:         cmd.String("filter"),
								OrderBy:        cmd.String("order-by"),
								View:           networkconnectivitypb.ListHubSpokesRequest_SpokeView(networkconnectivitypb.ListHubSpokesRequest_SpokeView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListHubSpokes(ctx, req)
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
						Name:  "query-status",
						Usage: "query-status hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "group-by", Usage: "Aggregate the results by the specified fields.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results in ascending order by the specified fields.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.QueryHubStatusRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								GroupBy:   cmd.String("group-by"),
							}

							limit := cmd.Int("limit")
							it := client.QueryHubStatus(ctx, req)
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
						Name:  "reject-spoke",
						Usage: "reject-spoke hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "details", Usage: "Additional information provided by the hub administrator.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The URI of the spoke to reject from the hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.RejectHubSpokeRequest{
								Name:      name,
								SpokeUri:  cmd.String("spoke-uri"),
								RequestId: cmd.String("request-id"),
								Details:   cmd.String("details"),
							}

							op, err := client.RejectHubSpoke(ctx, req)
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
						Name:  "accept-spoke",
						Usage: "accept-spoke hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The URI of the spoke to accept into the hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.AcceptHubSpokeRequest{
								Name:      name,
								SpokeUri:  cmd.String("spoke-uri"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.AcceptHubSpoke(ctx, req)
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
						Name:  "accept-spoke-update",
						Usage: "accept-spoke-update hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke-etag", Usage: "The etag of the spoke to accept update.", Required: true},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The URI of the spoke to accept update.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.AcceptSpokeUpdateRequest{
								Name:      name,
								SpokeUri:  cmd.String("spoke-uri"),
								SpokeEtag: cmd.String("spoke-etag"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.AcceptSpokeUpdate(ctx, req)
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
						Name:  "reject-spoke-update",
						Usage: "reject-spoke-update hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "details", Usage: "Additional information provided by the hub administrator.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke-etag", Usage: "The etag of the spoke to reject update.", Required: true},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The URI of the spoke to reject update.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.RejectSpokeUpdateRequest{
								Name:      name,
								SpokeUri:  cmd.String("spoke-uri"),
								SpokeEtag: cmd.String("spoke-etag"),
								Details:   cmd.String("details"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.RejectSpokeUpdate(ctx, req)
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
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "internal-ranges",
				Usage: "Manage internal-ranges resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListInternalRangesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInternalRanges(ctx, req)
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
						Usage: "describe internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "internal-range", Usage: "The ID of the internal range.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/internalRanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("internal-range"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetInternalRangeRequest{
								Name: name,
							}

							resp, err := client.GetInternalRange(ctx, req)
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
						Usage: "create internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "internal-range-id", Usage: "Resource ID.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateInternalRangeRequest{
								Parent:          parent,
								InternalRangeId: cmd.String("internal-range-id"),
								RequestId:       cmd.String("request-id"),
							}

							op, err := client.CreateInternalRange(ctx, req)
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
						Usage: "update internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "internal-range", Usage: "The ID of the internal range.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "internal_range.name" not yet supported.
							internal_range_name := fmt.Sprintf("projects/%s/locations/%s/internalRanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("internal-range"))
							fmt.Printf("Executing update on %s\n", internal_range_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "internal-range", Usage: "The ID of the internal range.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/internalRanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("internal-range"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInternalRange %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteInternalRangeRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInternalRange(ctx, req)
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
				Name:  "multicloud-data-transfer-configs",
				Usage: "Manage multicloud-data-transfer-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The sort order of the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results listed per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If `true`, allows partial responses for multi-regional aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListMulticloudDataTransferConfigsRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								Filter:               cmd.String("filter"),
								OrderBy:              cmd.String("order-by"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListMulticloudDataTransferConfigs(ctx, req)
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
						Usage: "describe multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetMulticloudDataTransferConfigRequest{
								Name: name,
							}

							resp, err := client.GetMulticloudDataTransferConfig(ctx, req)
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
						Usage: "create multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config-id", Usage: "The ID to use for the `MulticloudDataTransferConfig` resource,.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateMulticloudDataTransferConfigRequest{
								Parent:                         parent,
								MulticloudDataTransferConfigId: cmd.String("multicloud-data-transfer-config-id"),
								RequestId:                      cmd.String("request-id"),
							}

							op, err := client.CreateMulticloudDataTransferConfig(ctx, req)
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
						Usage: "update multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "multicloud_data_transfer_config.name" not yet supported.
							multicloud_data_transfer_config_name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"))
							fmt.Printf("Executing update on %s\n", multicloud_data_transfer_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag is computed by the server, and might be sent with update.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config", Usage: "The ID of the multicloud data transfer config.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMulticloudDataTransferConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteMulticloudDataTransferConfigRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      cmd.String("etag"),
							}

							op, err := client.DeleteMulticloudDataTransferConfig(ctx, req)
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
				Name:  "multicloud-data-transfer-supported-services",
				Usage: "Manage multicloud-data-transfer-supported-services resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe multicloud-data-transfer-supported-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-supported-service", Usage: "The ID of the multicloud data transfer supported service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferSupportedServices/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud-data-transfer-supported-service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetMulticloudDataTransferSupportedServiceRequest{
								Name: name,
							}

							resp, err := client.GetMulticloudDataTransferSupportedService(ctx, req)
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
						Usage: "list multicloud-data-transfer-supported-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results listed per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListMulticloudDataTransferSupportedServicesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMulticloudDataTransferSupportedServices(ctx, req)
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
				Name:  "policy-based-routes",
				Usage: "Manage policy-based-routes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListPolicyBasedRoutesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPolicyBasedRoutes(ctx, req)
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
						Usage: "describe policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy-based-route", Usage: "The ID of the policy based route.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/policyBasedRoutes/%s", cmd.String("project"), cmd.String("policy-based-route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetPolicyBasedRouteRequest{
								Name: name,
							}

							resp, err := client.GetPolicyBasedRoute(ctx, req)
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
						Usage: "create policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy-based-route-id", Usage: "Unique id for the policy-based route to create.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreatePolicyBasedRouteRequest{
								Parent:             parent,
								PolicyBasedRouteId: cmd.String("policy-based-route-id"),
								RequestId:          cmd.String("request-id"),
							}

							op, err := client.CreatePolicyBasedRoute(ctx, req)
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
						Usage: "delete policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy-based-route", Usage: "The ID of the policy based route.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/policyBasedRoutes/%s", cmd.String("project"), cmd.String("policy-based-route"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePolicyBasedRoute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeletePolicyBasedRouteRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeletePolicyBasedRoute(ctx, req)
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
				Name:  "route-tables",
				Usage: "Manage route-tables resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe route-tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route-table", Usage: "The ID of the route table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/routeTables/%s", cmd.String("project"), cmd.String("hub"), cmd.String("route-table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetRouteTableRequest{
								Name: name,
							}

							resp, err := client.GetRouteTable(ctx, req)
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
						Usage: "list route-tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListRouteTablesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRouteTables(ctx, req)
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
				Name:  "routes",
				Usage: "Manage routes resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route", Usage: "The ID of the route.", Required: true},
							&cli.StringFlag{Name: "route-table", Usage: "The ID of the route table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/routeTables/%s/routes/%s", cmd.String("project"), cmd.String("hub"), cmd.String("route-table"), cmd.String("route"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetRouteRequest{
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
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The ID of the hub.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route-table", Usage: "The ID of the route table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global/hubs/%s/routeTables/%s", cmd.String("project"), cmd.String("hub"), cmd.String("route-table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListRoutesRequest{
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
				},
			},
			{
				Name:  "service-classes",
				Usage: "Manage service-classes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListServiceClassesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceClasses(ctx, req)
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
						Usage: "describe service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-classe", Usage: "The ID of the service classe.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-classe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetServiceClassRequest{
								Name: name,
							}

							resp, err := client.GetServiceClass(ctx, req)
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
						Usage: "update service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-classe", Usage: "The ID of the service classe.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_class.name" not yet supported.
							service_class_name := fmt.Sprintf("projects/%s/locations/%s/serviceClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-classe"))
							fmt.Printf("Executing update on %s\n", service_class_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag is computed by the server, and may be sent on update and.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-classe", Usage: "The ID of the service classe.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-classe"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServiceClass %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteServiceClassRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      runtime.Ptr(cmd.String("etag")),
							}

							op, err := client.DeleteServiceClass(ctx, req)
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
				Name:  "service-connection-maps",
				Usage: "Manage service-connection-maps resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListServiceConnectionMapsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceConnectionMaps(ctx, req)
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
						Usage: "describe service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-connection-map", Usage: "The ID of the service connection map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-map"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetServiceConnectionMapRequest{
								Name: name,
							}

							resp, err := client.GetServiceConnectionMap(ctx, req)
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
						Usage: "create service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-map-id", Usage: "Resource ID.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateServiceConnectionMapRequest{
								Parent:                 parent,
								ServiceConnectionMapId: cmd.String("service-connection-map-id"),
								RequestId:              cmd.String("request-id"),
							}

							op, err := client.CreateServiceConnectionMap(ctx, req)
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
						Usage: "update service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-map", Usage: "The ID of the service connection map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_connection_map.name" not yet supported.
							service_connection_map_name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-map"))
							fmt.Printf("Executing update on %s\n", service_connection_map_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag is computed by the server, and may be sent on update and.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-map", Usage: "The ID of the service connection map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-map"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServiceConnectionMap %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteServiceConnectionMapRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      runtime.Ptr(cmd.String("etag")),
							}

							op, err := client.DeleteServiceConnectionMap(ctx, req)
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
				Name:  "service-connection-policies",
				Usage: "Manage service-connection-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListServiceConnectionPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceConnectionPolicies(ctx, req)
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
						Usage: "describe service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-connection-policie", Usage: "The ID of the service connection policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetServiceConnectionPolicyRequest{
								Name: name,
							}

							resp, err := client.GetServiceConnectionPolicy(ctx, req)
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
						Usage: "create service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-policy-id", Usage: "Resource ID.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateServiceConnectionPolicyRequest{
								Parent:                    parent,
								ServiceConnectionPolicyId: cmd.String("service-connection-policy-id"),
								RequestId:                 cmd.String("request-id"),
							}

							op, err := client.CreateServiceConnectionPolicy(ctx, req)
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
						Usage: "update service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-policie", Usage: "The ID of the service connection policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_connection_policy.name" not yet supported.
							service_connection_policy_name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-policie"))
							fmt.Printf("Executing update on %s\n", service_connection_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag is computed by the server, and may be sent on update and.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-policie", Usage: "The ID of the service connection policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServiceConnectionPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteServiceConnectionPolicyRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      runtime.Ptr(cmd.String("etag")),
							}

							op, err := client.DeleteServiceConnectionPolicy(ctx, req)
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
				Name:  "service-connection-tokens",
				Usage: "Manage service-connection-tokens resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-connection-token", Usage: "The ID of the service connection token.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionTokens/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-token"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetServiceConnectionTokenRequest{
								Name: name,
							}

							resp, err := client.GetServiceConnectionToken(ctx, req)
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
						Usage: "list service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListServiceConnectionTokensRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceConnectionTokens(ctx, req)
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
						Usage: "create service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-token-id", Usage: "Resource ID.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateServiceConnectionTokenRequest{
								Parent:                   parent,
								ServiceConnectionTokenId: cmd.String("service-connection-token-id"),
								RequestId:                cmd.String("request-id"),
							}

							op, err := client.CreateServiceConnectionToken(ctx, req)
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
						Usage: "delete service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag is computed by the server, and may be sent on update and.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service-connection-token", Usage: "The ID of the service connection token.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionTokens/%s", cmd.String("project"), cmd.String("location"), cmd.String("service-connection-token"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServiceConnectionToken %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteServiceConnectionTokenRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Etag:      runtime.Ptr(cmd.String("etag")),
							}

							op, err := client.DeleteServiceConnectionToken(ctx, req)
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
				Name:  "spokes",
				Usage: "Manage spokes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort the results by a certain order.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.ListSpokesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSpokes(ctx, req)
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
						Usage: "describe spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "spoke", Usage: "The ID of the spoke.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/spokes/%s", cmd.String("project"), cmd.String("location"), cmd.String("spoke"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.GetSpokeRequest{
								Name: name,
							}

							resp, err := client.GetSpoke(ctx, req)
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
						Usage: "create spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke-id", Usage: "Unique id for the spoke to create.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.CreateSpokeRequest{
								Parent:    parent,
								SpokeId:   cmd.String("spoke-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateSpoke(ctx, req)
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
						Usage: "update spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke", Usage: "The ID of the spoke.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "spoke.name" not yet supported.
							spoke_name := fmt.Sprintf("projects/%s/locations/%s/spokes/%s", cmd.String("project"), cmd.String("location"), cmd.String("spoke"))
							fmt.Printf("Executing update on %s\n", spoke_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "spoke", Usage: "The ID of the spoke.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/spokes/%s", cmd.String("project"), cmd.String("location"), cmd.String("spoke"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSpoke %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networkconnectivity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networkconnectivitypb.DeleteSpokeRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteSpoke(ctx, req)
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
		},
	}
}
