package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	"cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "baremetalsolution",
				Usage: "manage Bare Metal Solution API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "instances",
						Usage: "Manage instances resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetInstanceRequest{Name: name}
									resp, err := client.GetInstance(ctx, req)
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
								Usage: "update instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "machine-type", Usage: "The machine type.", Required: false},
									&cli.BoolFlag{Name: "hyperthreading-enabled", Usage: "The hyperthreading enabled.", Required: false},
									&cli.StringFlag{Name: "os-image", Usage: "The os image.", Required: false},
									&cli.StringFlag{Name: "pod", Usage: "The pod.", Required: false},
									&cli.StringFlag{Name: "network-template", Usage: "The network template.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.UpdateInstanceRequest{}
									req.Instance = &baremetalsolutionpb.Instance{
										Name:                  name,
										Name:                  cmd.String("name"),
										MachineType:           cmd.String("machine-type"),
										HyperthreadingEnabled: cmd.Bool("hyperthreading-enabled"),
										OsImage:               cmd.String("os-image"),
										Pod:                   cmd.String("pod"),
										NetworkTemplate:       cmd.String("network-template"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("machine-type") {
										paths = append(paths, "machine_type")
									}
									if cmd.IsSet("hyperthreading-enabled") {
										paths = append(paths, "hyperthreading_enabled")
									}
									if cmd.IsSet("os-image") {
										paths = append(paths, "os_image")
									}
									if cmd.IsSet("pod") {
										paths = append(paths, "pod")
									}
									if cmd.IsSet("network-template") {
										paths = append(paths, "network_template")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "rename",
								Usage: "rename instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									fmt.Printf("Executing rename on %s\n", name)
									return nil
								},
							},
							{
								Name:  "reset",
								Usage: "reset instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.ResetInstanceRequest{Name: name}
									op, err := client.ResetInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "start",
								Usage: "start instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.StartInstanceRequest{Name: name}
									op, err := client.StartInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "stop",
								Usage: "stop instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.StopInstanceRequest{Name: name}
									op, err := client.StopInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "enable-interactive-serial-console",
								Usage: "enable-interactive-serial-console instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.EnableInteractiveSerialConsoleRequest{Name: name}
									op, err := client.EnableInteractiveSerialConsole(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "disable-interactive-serial-console",
								Usage: "disable-interactive-serial-console instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.DisableInteractiveSerialConsoleRequest{Name: name}
									op, err := client.DisableInteractiveSerialConsole(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "detach-lun",
								Usage: "detach-lun instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detach-lun...")
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
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
								Name:  "describe",
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "lun", Usage: "The lun.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/luns/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("lun"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetLunRequest{Name: name}
									resp, err := client.GetLun(ctx, req)
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
								Name:  "list",
								Usage: "list luns",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &baremetalsolutionpb.ListLunsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLuns(ctx, req)
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
								Name:  "evict",
								Usage: "evict luns",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "lun", Usage: "The lun.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/luns/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("lun"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.EvictLunRequest{Name: name}
									op, err := client.EvictLun(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("EvictLun completed on %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe networks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("network"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetNetworkRequest{Name: name}
									resp, err := client.GetNetwork(ctx, req)
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
								Usage: "update networks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
									&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
									&cli.StringFlag{Name: "ip-address", Usage: "The ip address.", Required: false},
									&cli.StringFlag{Name: "vlan-id", Usage: "The vlan id.", Required: false},
									&cli.StringFlag{Name: "cidr", Usage: "The cidr.", Required: false},
									&cli.StringFlag{Name: "services-cidr", Usage: "The services cidr.", Required: false},
									&cli.BoolFlag{Name: "jumbo-frames-enabled", Usage: "The jumbo frames enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("network"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.UpdateNetworkRequest{}
									req.Network = &baremetalsolutionpb.Network{
										Name:               name,
										Id:                 cmd.String("id"),
										IpAddress:          cmd.String("ip-address"),
										VlanId:             cmd.String("vlan-id"),
										Cidr:               cmd.String("cidr"),
										ServicesCidr:       cmd.String("services-cidr"),
										JumboFramesEnabled: cmd.Bool("jumbo-frames-enabled"),
									}
									var paths []string
									if cmd.IsSet("id") {
										paths = append(paths, "id")
									}
									if cmd.IsSet("ip-address") {
										paths = append(paths, "ip_address")
									}
									if cmd.IsSet("vlan-id") {
										paths = append(paths, "vlan_id")
									}
									if cmd.IsSet("cidr") {
										paths = append(paths, "cidr")
									}
									if cmd.IsSet("services-cidr") {
										paths = append(paths, "services_cidr")
									}
									if cmd.IsSet("jumbo-frames-enabled") {
										paths = append(paths, "jumbo_frames_enabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateNetwork(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "rename",
								Usage: "rename networks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("network"))
									fmt.Printf("Executing rename on %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "nfs_share", Usage: "The nfs_share.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs_share"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetNfsShareRequest{Name: name}
									resp, err := client.GetNfsShare(ctx, req)
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
								Name:  "list",
								Usage: "list nfs-shares",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update nfs-shares",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "nfs_share", Usage: "The nfs_share.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.IntFlag{Name: "requested-size-gib", Usage: "The requested size gib.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs_share"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.UpdateNfsShareRequest{}
									req.NfsShare = &baremetalsolutionpb.NfsShare{
										Name:             name,
										Name:             cmd.String("name"),
										RequestedSizeGib: int64(cmd.Int("requested-size-gib")),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("requested-size-gib") {
										paths = append(paths, "requested_size_gib")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateNfsShare(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "create",
								Usage: "create nfs-shares",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.IntFlag{Name: "requested-size-gib", Usage: "The requested size gib.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.CreateNfsShareRequest{Parent: parent}
									req.NfsShare = &baremetalsolutionpb.NfsShare{
										Name:             cmd.String("name"),
										RequestedSizeGib: int64(cmd.Int("requested-size-gib")),
									}
									op, err := client.CreateNfsShare(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "rename",
								Usage: "rename nfs-shares",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "nfs_share", Usage: "The nfs_share.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs_share"))
									fmt.Printf("Executing rename on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete nfs-shares",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "nfs_share", Usage: "The nfs_share.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/nfsShares/%s", cmd.String("project"), cmd.String("location"), cmd.String("nfs_share"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.DeleteNfsShareRequest{Name: name}
									op, err := client.DeleteNfsShare(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing submit on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe provisioning-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "provisioning_config", Usage: "The provisioning_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/provisioningConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("provisioning_config"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetProvisioningConfigRequest{Name: name}
									resp, err := client.GetProvisioningConfig(ctx, req)
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
								Name:  "create",
								Usage: "create provisioning-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "email", Usage: "The email.", Required: false},
									&cli.StringFlag{Name: "ticket-id", Usage: "The ticket id.", Required: false},
									&cli.StringFlag{Name: "handover-service-account", Usage: "The handover service account.", Required: false},
									&cli.StringFlag{Name: "email", Usage: "The email.", Required: false},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: false},
									&cli.BoolFlag{Name: "vpc-sc-enabled", Usage: "The vpc sc enabled.", Required: false},
									&cli.StringFlag{Name: "status-message", Usage: "The status message.", Required: false},
									&cli.StringFlag{Name: "custom-id", Usage: "The custom id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.CreateProvisioningConfigRequest{Parent: parent}
									req.Email = cmd.String("email")
									req.ProvisioningConfig = &baremetalsolutionpb.ProvisioningConfig{
										TicketId:               cmd.String("ticket-id"),
										HandoverServiceAccount: cmd.String("handover-service-account"),
										Email:                  cmd.String("email"),
										Location:               cmd.String("location"),
										VpcScEnabled:           cmd.Bool("vpc-sc-enabled"),
										StatusMessage:          cmd.String("status-message"),
										CustomId:               cmd.String("custom-id"),
									}
									resp, err := client.CreateProvisioningConfig(ctx, req)
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
								Usage: "update provisioning-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "provisioning_config", Usage: "The provisioning_config.", Required: true},
									&cli.StringFlag{Name: "ticket-id", Usage: "The ticket id.", Required: false},
									&cli.StringFlag{Name: "handover-service-account", Usage: "The handover service account.", Required: false},
									&cli.StringFlag{Name: "email", Usage: "The email.", Required: false},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: false},
									&cli.BoolFlag{Name: "vpc-sc-enabled", Usage: "The vpc sc enabled.", Required: false},
									&cli.StringFlag{Name: "status-message", Usage: "The status message.", Required: false},
									&cli.StringFlag{Name: "custom-id", Usage: "The custom id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/provisioningConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("provisioning_config"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.UpdateProvisioningConfigRequest{}
									req.ProvisioningConfig = &baremetalsolutionpb.ProvisioningConfig{
										Name:                   name,
										TicketId:               cmd.String("ticket-id"),
										HandoverServiceAccount: cmd.String("handover-service-account"),
										Email:                  cmd.String("email"),
										Location:               cmd.String("location"),
										VpcScEnabled:           cmd.Bool("vpc-sc-enabled"),
										StatusMessage:          cmd.String("status-message"),
										CustomId:               cmd.String("custom-id"),
									}
									var paths []string
									if cmd.IsSet("ticket-id") {
										paths = append(paths, "ticket_id")
									}
									if cmd.IsSet("handover-service-account") {
										paths = append(paths, "handover_service_account")
									}
									if cmd.IsSet("email") {
										paths = append(paths, "email")
									}
									if cmd.IsSet("location") {
										paths = append(paths, "location")
									}
									if cmd.IsSet("vpc-sc-enabled") {
										paths = append(paths, "vpc_sc_enabled")
									}
									if cmd.IsSet("status-message") {
										paths = append(paths, "status_message")
									}
									if cmd.IsSet("custom-id") {
										paths = append(paths, "custom_id")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateProvisioningConfig(ctx, req)
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
						},
					},
					{
						Name:  "provisioning-quotas",
						Usage: "Manage provisioning-quotas resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list provisioning-quotas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.CreateVolumeSnapshotRequest{Parent: parent}
									req.VolumeSnapshot = &baremetalsolutionpb.VolumeSnapshot{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreateVolumeSnapshot(ctx, req)
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
								Name:  "restore-volume-snapshot",
								Usage: "restore-volume-snapshot snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing restore-volume-snapshot...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.DeleteVolumeSnapshotRequest{Name: name}
									if err := client.DeleteVolumeSnapshot(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetVolumeSnapshotRequest{Name: name}
									resp, err := client.GetVolumeSnapshot(ctx, req)
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
								Name:  "list",
								Usage: "list snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &baremetalsolutionpb.ListVolumeSnapshotsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVolumeSnapshots(ctx, req)
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
						Name:  "ssh-keys",
						Usage: "Manage ssh-keys resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list ssh-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create ssh-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ssh-key-id", Usage: "The ssh key id.", Required: true},
									&cli.StringFlag{Name: "public-key", Usage: "The public key.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.CreateSSHKeyRequest{Parent: parent}
									req.SshKeyId = cmd.String("ssh-key-id")
									req.SshKey = &baremetalsolutionpb.SSHKey{
										PublicKey: cmd.String("public-key"),
									}
									resp, err := client.CreateSSHKey(ctx, req)
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
								Usage: "delete ssh-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ssh_key", Usage: "The ssh_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sshKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("ssh_key"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.DeleteSSHKeyRequest{Name: name}
									if err := client.DeleteSSHKey(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.GetVolumeRequest{Name: name}
									resp, err := client.GetVolume(ctx, req)
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
								Usage: "update volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
									&cli.IntFlag{Name: "requested-size-gib", Usage: "The requested size gib.", Required: false},
									&cli.IntFlag{Name: "originally-requested-size-gib", Usage: "The originally requested size gib.", Required: false},
									&cli.IntFlag{Name: "current-size-gib", Usage: "The current size gib.", Required: false},
									&cli.IntFlag{Name: "emergency-size-gib", Usage: "The emergency size gib.", Required: false},
									&cli.IntFlag{Name: "max-size-gib", Usage: "The max size gib.", Required: false},
									&cli.IntFlag{Name: "auto-grown-size-gib", Usage: "The auto grown size gib.", Required: false},
									&cli.IntFlag{Name: "remaining-space-gib", Usage: "The remaining space gib.", Required: false},
									&cli.BoolFlag{Name: "snapshot-enabled", Usage: "The snapshot enabled.", Required: false},
									&cli.StringFlag{Name: "pod", Usage: "The pod.", Required: false},
									&cli.StringFlag{Name: "notes", Usage: "The notes.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.UpdateVolumeRequest{}
									req.Volume = &baremetalsolutionpb.Volume{
										Name:                       name,
										Id:                         cmd.String("id"),
										RequestedSizeGib:           int64(cmd.Int("requested-size-gib")),
										OriginallyRequestedSizeGib: int64(cmd.Int("originally-requested-size-gib")),
										CurrentSizeGib:             int64(cmd.Int("current-size-gib")),
										EmergencySizeGib:           int64(cmd.Int("emergency-size-gib")),
										MaxSizeGib:                 int64(cmd.Int("max-size-gib")),
										AutoGrownSizeGib:           int64(cmd.Int("auto-grown-size-gib")),
										RemainingSpaceGib:          int64(cmd.Int("remaining-space-gib")),
										SnapshotEnabled:            cmd.Bool("snapshot-enabled"),
										Pod:                        cmd.String("pod"),
										Notes:                      cmd.String("notes"),
									}
									var paths []string
									if cmd.IsSet("id") {
										paths = append(paths, "id")
									}
									if cmd.IsSet("requested-size-gib") {
										paths = append(paths, "requested_size_gib")
									}
									if cmd.IsSet("originally-requested-size-gib") {
										paths = append(paths, "originally_requested_size_gib")
									}
									if cmd.IsSet("current-size-gib") {
										paths = append(paths, "current_size_gib")
									}
									if cmd.IsSet("emergency-size-gib") {
										paths = append(paths, "emergency_size_gib")
									}
									if cmd.IsSet("max-size-gib") {
										paths = append(paths, "max_size_gib")
									}
									if cmd.IsSet("auto-grown-size-gib") {
										paths = append(paths, "auto_grown_size_gib")
									}
									if cmd.IsSet("remaining-space-gib") {
										paths = append(paths, "remaining_space_gib")
									}
									if cmd.IsSet("snapshot-enabled") {
										paths = append(paths, "snapshot_enabled")
									}
									if cmd.IsSet("pod") {
										paths = append(paths, "pod")
									}
									if cmd.IsSet("notes") {
										paths = append(paths, "notes")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateVolume(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "rename",
								Usage: "rename volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									fmt.Printf("Executing rename on %s\n", name)
									return nil
								},
							},
							{
								Name:  "evict",
								Usage: "evict volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := baremetalsolution.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &baremetalsolutionpb.EvictVolumeRequest{Name: name}
									op, err := client.EvictVolume(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("EvictVolume completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "resize",
								Usage: "resize volumes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
