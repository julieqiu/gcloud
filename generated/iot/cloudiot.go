package cloudiot

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	iot "cloud.google.com/go/iot/apiv1"
	"cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud cloudiot command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &iotpb.ListDeviceConfigVersionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDeviceConfigVersions(ctx, req)
							limit := cmd.Int("limit")
							count := 0
							for {
								if limit > 0 && count >= limit {
									break
								}
								resp, err := it.Next()
								if err == iterator.Done {
									break
								}
								if err != nil {
									return err
								}
								out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
								if err != nil {
									return err
								}
								if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
									continue
								}
								if cmd.Bool("uri") {
									fmt.Println(resp.GetName())
								} else {
									fmt.Println(string(out))
								}
								count++
							}
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.BoolFlag{Name: "blocked", Usage: "The blocked.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.CreateDeviceRequest{Parent: parent}
							req.Device = &iotpb.Device{
								Id:      cmd.String("id"),
								Name:    cmd.String("name"),
								Blocked: cmd.Bool("blocked"),
							}
							resp, err := client.CreateDevice(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The device.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"), cmd.String("device"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.GetDeviceRequest{Name: name}
							resp, err := client.GetDevice(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The device.", Required: true},
							&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.BoolFlag{Name: "blocked", Usage: "The blocked.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"), cmd.String("device"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.UpdateDeviceRequest{}
							req.Device = &iotpb.Device{
								Name:    name,
								Id:      cmd.String("id"),
								Name:    cmd.String("name"),
								Blocked: cmd.Bool("blocked"),
							}
							var paths []string
							if cmd.IsSet("id") {
								paths = append(paths, "id")
							}
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("blocked") {
								paths = append(paths, "blocked")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateDevice(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The device.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"), cmd.String("device"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.DeleteDeviceRequest{Name: name}
							if err := client.DeleteDevice(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &iotpb.ListDevicesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDevices(ctx, req)
							limit := cmd.Int("limit")
							count := 0
							for {
								if limit > 0 && count >= limit {
									break
								}
								resp, err := it.Next()
								if err == iterator.Done {
									break
								}
								if err != nil {
									return err
								}
								out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
								if err != nil {
									return err
								}
								if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
									continue
								}
								if cmd.Bool("uri") {
									fmt.Println(resp.GetName())
								} else {
									fmt.Println(string(out))
								}
								count++
							}
							return nil
						},
					},
					{
						Name:  "modify-cloud-to-device-config",
						Usage: "modify-cloud-to-device-config devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The device.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"), cmd.String("device"))
							fmt.Printf("Executing modify-cloud-to-device-config on %s\n", name)
							return nil
						},
					},
					{
						Name:  "send-command-to-device",
						Usage: "send-command-to-device devices",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "device", Usage: "The device.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"), cmd.String("device"))
							fmt.Printf("Executing send-command-to-device on %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.CreateDeviceRegistryRequest{Parent: parent}
							req.DeviceRegistry = &iotpb.DeviceRegistry{
								Id:   cmd.String("id"),
								Name: cmd.String("name"),
							}
							resp, err := client.CreateDeviceRegistry(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.GetDeviceRegistryRequest{Name: name}
							resp, err := client.GetDeviceRegistry(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.UpdateDeviceRegistryRequest{}
							req.DeviceRegistry = &iotpb.DeviceRegistry{
								Name: name,
								Id:   cmd.String("id"),
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("id") {
								paths = append(paths, "id")
							}
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateDeviceRegistry(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iotpb.DeleteDeviceRegistryRequest{Name: name}
							if err := client.DeleteDeviceRegistry(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list registries",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registrie", Usage: "The registrie.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registrie"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "bind-device-to-gateway",
						Usage: "bind-device-to-gateway registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							fmt.Printf("Executing bind-device-to-gateway on %s\n", name)
							return nil
						},
					},
					{
						Name:  "unbind-device-from-gateway",
						Usage: "unbind-device-from-gateway registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							fmt.Printf("Executing unbind-device-from-gateway on %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registry", Usage: "The registry.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/registries/%s", cmd.String("project"), cmd.String("location"), cmd.String("registry"))
							client, err := iot.NewDeviceManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &iotpb.ListDeviceStatesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDeviceStates(ctx, req)
							limit := cmd.Int("limit")
							count := 0
							for {
								if limit > 0 && count >= limit {
									break
								}
								resp, err := it.Next()
								if err == iterator.Done {
									break
								}
								if err != nil {
									return err
								}
								out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
								if err != nil {
									return err
								}
								if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
									continue
								}
								if cmd.Bool("uri") {
									fmt.Println(resp.GetName())
								} else {
									fmt.Println(string(out))
								}
								count++
							}
							return nil
						},
					},
				},
			},
		},
	}
}
