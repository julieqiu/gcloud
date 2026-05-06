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

package iot

import (
	iot "cloud.google.com/go/iot/apiv1"
	"cloud.google.com/go/iot/apiv1/iotpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudiot command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudiot",
		Usage: "manage Cloud IoT API resources",
		Commands: []*cli.Command{
			{
				Name:  "config-versions",
				Usage: "Manage config-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list config-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "num-versions", Usage: "The number of versions to list.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.ListDeviceConfigVersionsRequest{
								Name:        name,
								NumVersions: int32(cmd.Int("num-versions")),
							}

							resp, err := client.ListDeviceConfigVersions(ctx, req)
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
				Name:  "devices",
				Usage: "Manage devices resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.CreateDeviceRequest{
								Parent: parent,
							}

							resp, err := client.CreateDevice(ctx, req)
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
						Usage: "describe devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.GetDeviceRequest{
								Name: name,
							}

							resp, err := client.GetDevice(ctx, req)
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
						Usage: "update devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "device.name" not yet supported.
							device_name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							fmt.Printf("Executing update on %s\n", device_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDevice on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.DeleteDeviceRequest{
								Name: name,
							}

							if err := client.DeleteDevice(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list devices",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "device-ids", Usage: "A list of device string IDs.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of devices to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListDevicesResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.ListDevicesRequest{
								Parent:    parent,
								DeviceIds: cmd.StringSlice("device-ids"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDevices(ctx, req)
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
						Name:  "modify-cloud-to-device-config",
						Usage: "modify-cloud-to-device-config devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "binary-data", Usage: "The configuration data for the device.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
							&cli.IntFlag{Name: "version-to-update", Usage: "The version number to update.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.ModifyCloudToDeviceConfigRequest{
								Name:            name,
								VersionToUpdate: cmd.Int("version-to-update"),
								BinaryData:      []byte(cmd.String("binary-data")),
							}

							resp, err := client.ModifyCloudToDeviceConfig(ctx, req)
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
						Name:  "send-command-to-device",
						Usage: "send-command-to-device devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "binary-data", Usage: "The command data to send to the device.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
							&cli.StringFlag{Name: "subfolder", Usage: "Optional subfolder for the command.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.SendCommandToDeviceRequest{
								Name:       name,
								BinaryData: []byte(cmd.String("binary-data")),
								Subfolder:  cmd.String("subfolder"),
							}

							resp, err := client.SendCommandToDevice(ctx, req)
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
				Name:  "registries",
				Usage: "Manage registries resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create registries",
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
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.CreateDeviceRegistryRequest{
								Parent: parent,
							}

							resp, err := client.CreateDeviceRegistry(ctx, req)
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
						Usage: "describe registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.GetDeviceRegistryRequest{
								Name: name,
							}

							resp, err := client.GetDeviceRegistry(ctx, req)
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
						Usage: "update registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "device_registry.name" not yet supported.
							device_registry_name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							fmt.Printf("Executing update on %s\n", device_registry_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDeviceRegistry on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.DeleteDeviceRegistryRequest{
								Name: name,
							}

							if err := client.DeleteDeviceRegistry(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of registries to return in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListDeviceRegistriesResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.ListDeviceRegistriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeviceRegistries(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
						Name:  "bind-device-to-gateway",
						Usage: "bind-device-to-gateway registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-id", Usage: "The device to associate with the specified gateway.", Required: true},
							&cli.StringFlag{Name: "gateway-id", Usage: "The value of `gateway_id` can be either the device numeric ID or the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.BindDeviceToGatewayRequest{
								Parent:    parent,
								GatewayId: cmd.String("gateway-id"),
								DeviceId:  cmd.String("device-id"),
							}

							resp, err := client.BindDeviceToGateway(ctx, req)
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
						Name:  "unbind-device-from-gateway",
						Usage: "unbind-device-from-gateway registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-id", Usage: "The device to disassociate from the specified gateway.", Required: true},
							&cli.StringFlag{Name: "gateway-id", Usage: "The value of `gateway_id` can be either the device numeric ID or the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.UnbindDeviceFromGatewayRequest{
								Parent:    parent,
								GatewayId: cmd.String("gateway-id"),
								DeviceId:  cmd.String("device-id"),
							}

							resp, err := client.UnbindDeviceFromGateway(ctx, req)
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
				Name:  "states",
				Usage: "Manage states resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list states",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device", Usage: "The ID of the device.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "num-states", Usage: "The number of states to list.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The ID of the registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"), cmd.String("device"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iot.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iotpb.ListDeviceStatesRequest{
								Name:      name,
								NumStates: int32(cmd.Int("num-states")),
							}

							resp, err := client.ListDeviceStates(ctx, req)
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
