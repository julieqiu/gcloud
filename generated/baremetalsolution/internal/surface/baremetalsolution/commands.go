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

package baremetalsolution

import (
	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	"cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the baremetalsolution command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "baremetalsolution",
		Usage: "manage Bare Metal Solution API resources",
		Commands: []*cli.Command{
			{
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListInstancesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListInstances(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetInstanceRequest{
								Name: name,
							}

							resp, err := client.GetInstance(ctx, req)
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
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "instance.name" not yet supported.
							instance_name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", instance_name)
							return nil
						},
					},

					{
						Name:  "rename",
						Usage: "rename instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-instance-id", Usage: "The new `id` of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.RenameInstanceRequest{
								Name:          name,
								NewInstanceId: cmd.String("new-instance-id"),
							}

							resp, err := client.RenameInstance(ctx, req)
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
						Name:  "reset",
						Usage: "reset instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ResetInstanceRequest{
								Name: name,
							}

							op, err := client.ResetInstance(ctx, req)
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
						Name:  "start",
						Usage: "start instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.StartInstanceRequest{
								Name: name,
							}

							op, err := client.StartInstance(ctx, req)
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
						Name:  "stop",
						Usage: "stop instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.StopInstanceRequest{
								Name: name,
							}

							op, err := client.StopInstance(ctx, req)
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
						Name:  "enable-interactive-serial-console",
						Usage: "enable-interactive-serial-console instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.EnableInteractiveSerialConsoleRequest{
								Name: name,
							}

							op, err := client.EnableInteractiveSerialConsole(ctx, req)
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
						Name:  "disable-interactive-serial-console",
						Usage: "disable-interactive-serial-console instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.DisableInteractiveSerialConsoleRequest{
								Name: name,
							}

							op, err := client.DisableInteractiveSerialConsole(ctx, req)
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
						Name:  "detach-lun",
						Usage: "detach-lun instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "lun", Usage: "Name of the Lun to detach.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-reboot", Usage: "If true, performs lun unmapping without instance reboot.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							instance := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							fmt.Printf("Executing detach-lun on %s\n", instance)
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
				Name:  "luns",
				Usage: "Manage luns resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe luns",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "lun", Usage: "The ID of the lun.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/luns/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("lun"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetLunRequest{
								Name: name,
							}

							resp, err := client.GetLun(ctx, req)
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
						Usage: "list luns",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListLunsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListLuns(ctx, req)
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
						Name:  "evict",
						Usage: "evict luns",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "lun", Usage: "The ID of the lun.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/luns/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("lun"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("EvictLun %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.EvictLunRequest{
								Name: name,
							}

							op, err := client.EvictLun(ctx, req)
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
				Name:  "networks",
				Usage: "Manage networks resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListNetworksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
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
						Name:  "list",
						Usage: "list networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", location)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetNetworkRequest{
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
						Name:  "update",
						Usage: "update networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "network.name" not yet supported.
							network_name := fmt.Sprintf("projects/%s/locations/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("network"))
							fmt.Printf("Executing update on %s\n", network_name)
							return nil
						},
					},

					{
						Name:  "rename",
						Usage: "rename networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "new-network-id", Usage: "The new `id` of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.RenameNetworkRequest{
								Name:         name,
								NewNetworkId: cmd.String("new-network-id"),
							}

							resp, err := client.RenameNetwork(ctx, req)
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
				Name:  "nfs-shares",
				Usage: "Manage nfs-shares resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe nfs-shares",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nfs-share", Usage: "The ID of the nfs share.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs-share"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetNfsShareRequest{
								Name: name,
							}

							resp, err := client.GetNfsShare(ctx, req)
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
						Usage: "list nfs-shares",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListNfsSharesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListNfsShares(ctx, req)
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
						Usage: "update nfs-shares",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nfs-share", Usage: "The ID of the nfs share.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "nfs_share.name" not yet supported.
							nfs_share_name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs-share"))
							fmt.Printf("Executing update on %s\n", nfs_share_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create nfs-shares",
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
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.CreateNfsShareRequest{
								Parent: parent,
							}

							op, err := client.CreateNfsShare(ctx, req)
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
						Name:  "rename",
						Usage: "rename nfs-shares",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-nfsshare-id", Usage: "The new `id` of the nfsshare.", Required: true},
							&cli.StringFlag{Name: "nfs-share", Usage: "The ID of the nfs share.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs-share"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.RenameNfsShareRequest{
								Name:          name,
								NewNfsshareId: cmd.String("new-nfsshare-id"),
							}

							resp, err := client.RenameNfsShare(ctx, req)
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
						Usage: "delete nfs-shares",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "nfs-share", Usage: "The ID of the nfs share.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs-share"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNfsShare %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.DeleteNfsShareRequest{
								Name: name,
							}

							op, err := client.DeleteNfsShare(ctx, req)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "os-images",
				Usage: "Manage os-images resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list os-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListOSImagesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListOSImages(ctx, req)
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
				Name:  "provisioning-configs",
				Usage: "Manage provisioning-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "submit",
						Usage: "submit provisioning-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "email", Usage: "Email provided to send a confirmation with provisioning config.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.SubmitProvisioningConfigRequest{
								Parent: parent,
								Email:  cmd.String("email"),
							}

							resp, err := client.SubmitProvisioningConfig(ctx, req)
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
						Usage: "describe provisioning-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "provisioning-config", Usage: "The ID of the provisioning config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/provisioningConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("provisioning-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetProvisioningConfigRequest{
								Name: name,
							}

							resp, err := client.GetProvisioningConfig(ctx, req)
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
						Usage: "create provisioning-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "email", Usage: "Email provided to send a confirmation with provisioning config.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.CreateProvisioningConfigRequest{
								Parent: parent,
								Email:  cmd.String("email"),
							}

							resp, err := client.CreateProvisioningConfig(ctx, req)
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
						Usage: "update provisioning-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "email", Usage: "Email provided to send a confirmation with provisioning config.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "provisioning-config", Usage: "The ID of the provisioning config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "provisioning_config.name" not yet supported.
							provisioning_config_name := fmt.Sprintf("projects/%s/locations/%s/provisioningConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("provisioning-config"))
							fmt.Printf("Executing update on %s\n", provisioning_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "provisioning-quotas",
				Usage: "Manage provisioning-quotas resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list provisioning-quotas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListProvisioningQuotasRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListProvisioningQuotas(ctx, req)
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
				Name:  "snapshots",
				Usage: "Manage snapshots resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.CreateVolumeSnapshotRequest{
								Parent: parent,
							}

							resp, err := client.CreateVolumeSnapshot(ctx, req)
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
						Name:  "restore-volume-snapshot",
						Usage: "restore-volume-snapshot snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
							&cli.StringFlag{Name: "volume-snapshot", Usage: "Name of the snapshot which will be used to restore its parent.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							volume_snapshot := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
							fmt.Printf("Executing restore-volume-snapshot on %s\n", volume_snapshot)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteVolumeSnapshot on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.DeleteVolumeSnapshotRequest{
								Name: name,
							}

							if err := client.DeleteVolumeSnapshot(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetVolumeSnapshotRequest{
								Name: name,
							}

							resp, err := client.GetVolumeSnapshot(ctx, req)
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
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListVolumeSnapshotsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListVolumeSnapshots(ctx, req)
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
				Name:  "ssh-keys",
				Usage: "Manage ssh-keys resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list ssh-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListSSHKeysRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSSHKeys(ctx, req)
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
						Usage: "create ssh-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "ssh-key-id", Usage: "The ID to use for the key, which will become the final component.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.CreateSSHKeyRequest{
								Parent:   parent,
								SshKeyId: cmd.String("ssh-key-id"),
							}

							resp, err := client.CreateSSHKey(ctx, req)
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
						Usage: "delete ssh-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "ssh-key", Usage: "The ID of the ssh key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sshKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("ssh-key"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSSHKey on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.DeleteSSHKeyRequest{
								Name: name,
							}

							if err := client.DeleteSSHKey(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "volumes",
				Usage: "Manage volumes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results from the server.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.ListVolumesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListVolumes(ctx, req)
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
						Usage: "describe volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.GetVolumeRequest{
								Name: name,
							}

							resp, err := client.GetVolume(ctx, req)
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
						Usage: "update volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "volume.name" not yet supported.
							volume_name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							fmt.Printf("Executing update on %s\n", volume_name)
							return nil
						},
					},

					{
						Name:  "rename",
						Usage: "rename volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-volume-id", Usage: "The new `id` of the volume.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.RenameVolumeRequest{
								Name:        name,
								NewVolumeId: cmd.String("new-volume-id"),
							}

							resp, err := client.RenameVolume(ctx, req)
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
						Name:  "evict",
						Usage: "evict volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("EvictVolume %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := baremetalsolution.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &baremetalsolutionpb.EvictVolumeRequest{
								Name: name,
							}

							op, err := client.EvictVolume(ctx, req)
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
						Name:  "resize",
						Usage: "resize volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "size-gib", Usage: "New Volume size, in GiB.", Required: false},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							volume := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							fmt.Printf("Executing resize on %s\n", volume)
							return nil
						},
					},
				},
			},
		},
	}
}
