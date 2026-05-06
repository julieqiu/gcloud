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

package edgenetwork

import (
	edgenetwork "cloud.google.com/go/edgenetwork/apiv1"
	"cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the edgenetwork command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "edgenetwork",
		Usage: "manage Distributed Cloud Edge Network API resources",
		Commands: []*cli.Command{
			{
				Name:  "interconnect-attachments",
				Usage: "Manage interconnect-attachments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.ListInterconnectAttachmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInterconnectAttachments(ctx, req)
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
						Usage: "describe interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect-attachment", Usage: "The ID of the interconnect attachment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnectAttachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect-attachment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.GetInterconnectAttachmentRequest{
								Name: name,
							}

							resp, err := client.GetInterconnectAttachment(ctx, req)
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
						Usage: "create interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect-attachment-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.CreateInterconnectAttachmentRequest{
								Parent:                   parent,
								InterconnectAttachmentId: cmd.String("interconnect-attachment-id"),
								RequestId:                cmd.String("request-id"),
							}

							op, err := client.CreateInterconnectAttachment(ctx, req)
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
						Usage: "delete interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect-attachment", Usage: "The ID of the interconnect attachment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnectAttachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect-attachment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInterconnectAttachment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DeleteInterconnectAttachmentRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInterconnectAttachment(ctx, req)
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
				Name:  "interconnects",
				Usage: "Manage interconnects resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.ListInterconnectsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInterconnects(ctx, req)
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
						Usage: "describe interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnects/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.GetInterconnectRequest{
								Name: name,
							}

							resp, err := client.GetInterconnect(ctx, req)
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
						Name:  "diagnose",
						Usage: "diagnose interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnects/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DiagnoseInterconnectRequest{
								Name: name,
							}

							resp, err := client.DiagnoseInterconnect(ctx, req)
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
				Name:  "networks",
				Usage: "Manage networks resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.ListNetworksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNetworks(ctx, req)
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
						Usage: "describe networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.GetNetworkRequest{
								Name: name,
							}

							resp, err := client.GetNetwork(ctx, req)
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
						Name:  "diagnose",
						Usage: "diagnose networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DiagnoseNetworkRequest{
								Name: name,
							}

							resp, err := client.DiagnoseNetwork(ctx, req)
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
						Usage: "create networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.CreateNetworkRequest{
								Parent:    parent,
								NetworkId: cmd.String("network-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateNetwork(ctx, req)
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
						Usage: "delete networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("network"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNetwork %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DeleteNetworkRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteNetwork(ctx, req)
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
				Name:  "routers",
				Usage: "Manage routers resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.ListRoutersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRouters(ctx, req)
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
						Usage: "describe routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.GetRouterRequest{
								Name: name,
							}

							resp, err := client.GetRouter(ctx, req)
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
						Name:  "diagnose",
						Usage: "diagnose routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DiagnoseRouterRequest{
								Name: name,
							}

							resp, err := client.DiagnoseRouter(ctx, req)
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
						Usage: "create routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.CreateRouterRequest{
								Parent:    parent,
								RouterId:  cmd.String("router-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateRouter(ctx, req)
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
						Usage: "update routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "router.name" not yet supported.
							router_name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							fmt.Printf("Executing update on %s\n", router_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRouter %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DeleteRouterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteRouter(ctx, req)
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
				Name:  "subnets",
				Usage: "Manage subnets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.ListSubnetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSubnets(ctx, req)
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
						Usage: "describe subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The ID of the subnet.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("subnet"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.GetSubnetRequest{
								Name: name,
							}

							resp, err := client.GetSubnet(ctx, req)
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
						Usage: "create subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnet-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.CreateSubnetRequest{
								Parent:    parent,
								SubnetId:  cmd.String("subnet-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateSubnet(ctx, req)
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
						Usage: "update subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnet", Usage: "The ID of the subnet.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "subnet.name" not yet supported.
							subnet_name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("subnet"))
							fmt.Printf("Executing update on %s\n", subnet_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnet", Usage: "The ID of the subnet.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("subnet"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSubnet %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.DeleteSubnetRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteSubnet(ctx, req)
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
				Name:  "zones",
				Usage: "Manage zones resources",
				Commands: []*cli.Command{

					{
						Name:  "initialize",
						Usage: "initialize zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.InitializeZoneRequest{
								Name: name,
							}

							resp, err := client.InitializeZone(ctx, req)
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
						Usage: "list zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.ListZonesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListZones(ctx, req)
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
						Usage: "describe zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgenetwork.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgenetworkpb.GetZoneRequest{
								Name: name,
							}

							resp, err := client.GetZone(ctx, req)
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
