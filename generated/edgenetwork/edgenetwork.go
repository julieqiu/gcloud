package edgenetwork

import (
	"context"
	"fmt"
	"strings"

	edgenetwork "cloud.google.com/go/edgenetwork/apiv1"
	"cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud edgenetwork command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &edgenetworkpb.ListInterconnectAttachmentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListInterconnectAttachments(ctx, req)
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
						Usage: "describe interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "interconnect_attachment", Usage: "The interconnect_attachment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnectAttachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect_attachment"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.GetInterconnectAttachmentRequest{Name: name}
							resp, err := client.GetInterconnectAttachment(ctx, req)
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
						Usage: "create interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "interconnect-attachment-id", Usage: "The interconnect attachment id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "interconnect", Usage: "The interconnect.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.IntFlag{Name: "vlan-id", Usage: "The vlan id.", Required: true},
							&cli.IntFlag{Name: "mtu", Usage: "The mtu.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.CreateInterconnectAttachmentRequest{Parent: parent}
							req.InterconnectAttachmentId = cmd.String("interconnect-attachment-id")
							req.InterconnectAttachment = &edgenetworkpb.InterconnectAttachment{
								Name:         cmd.String("name"),
								Description:  cmd.String("description"),
								Interconnect: cmd.String("interconnect"),
								Network:      cmd.String("network"),
								VlanId:       int32(cmd.Int("vlan-id")),
								Mtu:          int32(cmd.Int("mtu")),
							}
							op, err := client.CreateInterconnectAttachment(ctx, req)
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
						Name:  "delete",
						Usage: "delete interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "interconnect_attachment", Usage: "The interconnect_attachment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnectAttachments/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect_attachment"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.DeleteInterconnectAttachmentRequest{Name: name}
							op, err := client.DeleteInterconnectAttachment(ctx, req)
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
				Name:  "interconnects",
				Usage: "Manage interconnects resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &edgenetworkpb.ListInterconnectsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListInterconnects(ctx, req)
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
						Usage: "describe interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "interconnect", Usage: "The interconnect.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnects/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.GetInterconnectRequest{Name: name}
							resp, err := client.GetInterconnect(ctx, req)
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
						Name:  "diagnose",
						Usage: "diagnose interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "interconnect", Usage: "The interconnect.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/interconnects/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("interconnect"))
							fmt.Printf("Executing diagnose on %s\n", name)
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
							client, err := edgenetwork.NewClient(ctx)
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
							client, err := edgenetwork.NewClient(ctx)
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
				Name:  "networks",
				Usage: "Manage networks resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &edgenetworkpb.ListNetworksRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListNetworks(ctx, req)
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
						Usage: "describe networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("network"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.GetNetworkRequest{Name: name}
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
						Name:  "diagnose",
						Usage: "diagnose networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("network"))
							fmt.Printf("Executing diagnose on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "network-id", Usage: "The network id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.IntFlag{Name: "mtu", Usage: "The mtu.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.CreateNetworkRequest{Parent: parent}
							req.NetworkId = cmd.String("network-id")
							req.Network = &edgenetworkpb.Network{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Mtu:         int32(cmd.Int("mtu")),
							}
							op, err := client.CreateNetwork(ctx, req)
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
						Name:  "delete",
						Usage: "delete networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/networks/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("network"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.DeleteNetworkRequest{Name: name}
							op, err := client.DeleteNetwork(ctx, req)
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
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &longrunningpb.ListOperationsRequest{Name: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListOperations(ctx, req)
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
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &longrunningpb.GetOperationRequest{Name: name}
							resp, err := client.GetOperation(ctx, req)
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
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &longrunningpb.DeleteOperationRequest{Name: name}
							if err := client.DeleteOperation(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &longrunningpb.CancelOperationRequest{Name: name}
							if err := client.CancelOperation(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Cancelled %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &edgenetworkpb.ListRoutersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListRouters(ctx, req)
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
						Usage: "describe routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.GetRouterRequest{Name: name}
							resp, err := client.GetRouter(ctx, req)
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
						Name:  "diagnose",
						Usage: "diagnose routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							fmt.Printf("Executing diagnose on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "router-id", Usage: "The router id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.CreateRouterRequest{Parent: parent}
							req.RouterId = cmd.String("router-id")
							req.Router = &edgenetworkpb.Router{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Network:     cmd.String("network"),
							}
							op, err := client.CreateRouter(ctx, req)
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
						Name:  "update",
						Usage: "update routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The router.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.UpdateRouterRequest{}
							req.Router = &edgenetworkpb.Router{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Network:     cmd.String("network"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("network") {
								paths = append(paths, "network")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateRouter(ctx, req)
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
						Name:  "delete",
						Usage: "delete routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/routers/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("router"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.DeleteRouterRequest{Name: name}
							op, err := client.DeleteRouter(ctx, req)
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
				Name:  "subnets",
				Usage: "Manage subnets resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &edgenetworkpb.ListSubnetsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListSubnets(ctx, req)
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
						Usage: "describe subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("subnet"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.GetSubnetRequest{Name: name}
							resp, err := client.GetSubnet(ctx, req)
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
						Usage: "create subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "subnet-id", Usage: "The subnet id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
							&cli.IntFlag{Name: "vlan-id", Usage: "The vlan id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.CreateSubnetRequest{Parent: parent}
							req.SubnetId = cmd.String("subnet-id")
							req.Subnet = &edgenetworkpb.Subnet{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Network:     cmd.String("network"),
								VlanId:      int32(cmd.Int("vlan-id")),
							}
							op, err := client.CreateSubnet(ctx, req)
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
						Name:  "update",
						Usage: "update subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The subnet.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.IntFlag{Name: "vlan-id", Usage: "The vlan id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("subnet"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.UpdateSubnetRequest{}
							req.Subnet = &edgenetworkpb.Subnet{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Network:     cmd.String("network"),
								VlanId:      int32(cmd.Int("vlan-id")),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("network") {
								paths = append(paths, "network")
							}
							if cmd.IsSet("vlan-id") {
								paths = append(paths, "vlan_id")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateSubnet(ctx, req)
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
						Name:  "delete",
						Usage: "delete subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"), cmd.String("subnet"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.DeleteSubnetRequest{Name: name}
							op, err := client.DeleteSubnet(ctx, req)
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
				Name:  "zones",
				Usage: "Manage zones resources",
				Commands: []*cli.Command{
					{
						Name:  "initialize",
						Usage: "initialize zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							fmt.Printf("Executing initialize on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &edgenetworkpb.ListZonesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListZones(ctx, req)
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
						Usage: "describe zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("zone"))
							client, err := edgenetwork.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &edgenetworkpb.GetZoneRequest{Name: name}
							resp, err := client.GetZone(ctx, req)
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
		},
	}
}
