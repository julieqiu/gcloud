package networkconnectivity

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	networkconnectivity "cloud.google.com/go/networkconnectivity/apiv1"
	"cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud networkconnectivity command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkconnectivitypb.ListDestinationsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDestinations(ctx, req)
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
						Usage: "describe destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
							&cli.StringFlag{Name: "destination", Usage: "The destination.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"), cmd.String("destination"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetDestinationRequest{Name: name}
							resp, err := client.GetDestination(ctx, req)
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
						Usage: "create destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
							&cli.StringFlag{Name: "destination-id", Usage: "The destination id.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "ip-prefix", Usage: "The ip prefix.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateDestinationRequest{Parent: parent}
							req.DestinationId = cmd.String("destination-id")
							req.Destination = &networkconnectivitypb.Destination{
								Etag:        cmd.String("etag"),
								Description: cmd.String("description"),
								IpPrefix:    cmd.String("ip-prefix"),
							}
							op, err := client.CreateDestination(ctx, req)
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
						Usage: "update destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
							&cli.StringFlag{Name: "destination", Usage: "The destination.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "ip-prefix", Usage: "The ip prefix.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"), cmd.String("destination"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateDestinationRequest{}
							req.Destination = &networkconnectivitypb.Destination{
								Name:        name,
								Etag:        cmd.String("etag"),
								Description: cmd.String("description"),
								IpPrefix:    cmd.String("ip-prefix"),
							}
							var paths []string
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("ip-prefix") {
								paths = append(paths, "ip_prefix")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateDestination(ctx, req)
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
						Usage: "delete destinations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
							&cli.StringFlag{Name: "destination", Usage: "The destination.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s/destinations/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"), cmd.String("destination"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteDestinationRequest{Name: name}
							op, err := client.DeleteDestination(ctx, req)
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
				Name:  "groups",
				Usage: "Manage groups resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/groups/%s", cmd.String("project"), cmd.String("hub"), cmd.String("group"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetGroupRequest{Name: name}
							resp, err := client.GetGroup(ctx, req)
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
						Usage: "list groups",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkconnectivitypb.ListGroupsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListGroups(ctx, req)
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
						Name:  "update",
						Usage: "update groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/groups/%s", cmd.String("project"), cmd.String("hub"), cmd.String("group"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateGroupRequest{}
							req.Group = &networkconnectivitypb.Group{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateGroup(ctx, req)
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
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
						Name:  "list",
						Usage: "list hubs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetHubRequest{Name: name}
							resp, err := client.GetHub(ctx, req)
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
						Usage: "create hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub-id", Usage: "The hub id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.BoolFlag{Name: "export-psc", Usage: "The export psc.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateHubRequest{Parent: parent}
							req.HubId = cmd.String("hub-id")
							req.Hub = &networkconnectivitypb.Hub{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								ExportPsc:   cmd.Bool("export-psc"),
							}
							op, err := client.CreateHub(ctx, req)
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
						Usage: "update hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.BoolFlag{Name: "export-psc", Usage: "The export psc.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateHubRequest{}
							req.Hub = &networkconnectivitypb.Hub{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								ExportPsc:   cmd.Bool("export-psc"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("export-psc") {
								paths = append(paths, "export_psc")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateHub(ctx, req)
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
						Usage: "delete hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteHubRequest{Name: name}
							op, err := client.DeleteHub(ctx, req)
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
					{
						Name:  "list",
						Usage: "list hubs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkconnectivitypb.ListHubSpokesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListHubSpokes(ctx, req)
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
						Name:  "query-status",
						Usage: "query-status hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing query-status on %s\n", name)
							return nil
						},
					},
					{
						Name:  "reject-spoke",
						Usage: "reject-spoke hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The spoke uri.", Required: true},
							&cli.StringFlag{Name: "details", Usage: "The details.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.RejectHubSpokeRequest{Name: name}
							req.SpokeUri = cmd.String("spoke-uri")
							req.Details = cmd.String("details")
							op, err := client.RejectHubSpoke(ctx, req)
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
						Name:  "accept-spoke",
						Usage: "accept-spoke hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The spoke uri.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.AcceptHubSpokeRequest{Name: name}
							req.SpokeUri = cmd.String("spoke-uri")
							op, err := client.AcceptHubSpoke(ctx, req)
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
						Name:  "accept-spoke-update",
						Usage: "accept-spoke-update hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The spoke uri.", Required: true},
							&cli.StringFlag{Name: "spoke-etag", Usage: "The spoke etag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.AcceptSpokeUpdateRequest{Name: name}
							req.SpokeUri = cmd.String("spoke-uri")
							req.SpokeEtag = cmd.String("spoke-etag")
							op, err := client.AcceptSpokeUpdate(ctx, req)
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
						Name:  "reject-spoke-update",
						Usage: "reject-spoke-update hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "spoke-uri", Usage: "The spoke uri.", Required: true},
							&cli.StringFlag{Name: "spoke-etag", Usage: "The spoke etag.", Required: true},
							&cli.StringFlag{Name: "details", Usage: "The details.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.RejectSpokeUpdateRequest{Name: name}
							req.SpokeUri = cmd.String("spoke-uri")
							req.SpokeEtag = cmd.String("spoke-etag")
							req.Details = cmd.String("details")
							op, err := client.RejectSpokeUpdate(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
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
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
						Usage: "test-iam-permissions hubs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
				},
			},
			{
				Name:  "internal-ranges",
				Usage: "Manage internal-ranges resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list internal-ranges",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "internal_range", Usage: "The internal_range.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/internalRanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("internal_range"))
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetInternalRangeRequest{Name: name}
							resp, err := client.GetInternalRange(ctx, req)
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
						Usage: "create internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "internal-range-id", Usage: "The internal range id.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "ip-cidr-range", Usage: "The ip cidr range.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.IntFlag{Name: "prefix-length", Usage: "The prefix length.", Required: false},
							&cli.BoolFlag{Name: "immutable", Usage: "The immutable.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateInternalRangeRequest{Parent: parent}
							req.InternalRangeId = cmd.String("internal-range-id")
							req.InternalRange = &networkconnectivitypb.InternalRange{
								Description:  cmd.String("description"),
								IpCidrRange:  cmd.String("ip-cidr-range"),
								Network:      cmd.String("network"),
								PrefixLength: int32(cmd.Int("prefix-length")),
								Immutable:    cmd.Bool("immutable"),
							}
							op, err := client.CreateInternalRange(ctx, req)
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
						Usage: "update internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "internal_range", Usage: "The internal_range.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "ip-cidr-range", Usage: "The ip cidr range.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.IntFlag{Name: "prefix-length", Usage: "The prefix length.", Required: false},
							&cli.BoolFlag{Name: "immutable", Usage: "The immutable.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/internalRanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("internal_range"))
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateInternalRangeRequest{}
							req.InternalRange = &networkconnectivitypb.InternalRange{
								Name:         name,
								Description:  cmd.String("description"),
								IpCidrRange:  cmd.String("ip-cidr-range"),
								Network:      cmd.String("network"),
								PrefixLength: int32(cmd.Int("prefix-length")),
								Immutable:    cmd.Bool("immutable"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("ip-cidr-range") {
								paths = append(paths, "ip_cidr_range")
							}
							if cmd.IsSet("network") {
								paths = append(paths, "network")
							}
							if cmd.IsSet("prefix-length") {
								paths = append(paths, "prefix_length")
							}
							if cmd.IsSet("immutable") {
								paths = append(paths, "immutable")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateInternalRange(ctx, req)
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
						Usage: "delete internal-ranges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "internal_range", Usage: "The internal_range.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/internalRanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("internal_range"))
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteInternalRangeRequest{Name: name}
							op, err := client.DeleteInternalRange(ctx, req)
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
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
							client, err := networkconnectivity.NewHubClient(ctx)
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
							client, err := networkconnectivity.NewHubClient(ctx)
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
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
				Name:  "multicloud-data-transfer-configs",
				Usage: "Manage multicloud-data-transfer-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list multicloud-data-transfer-configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetMulticloudDataTransferConfigRequest{Name: name}
							resp, err := client.GetMulticloudDataTransferConfig(ctx, req)
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
						Usage: "create multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud-data-transfer-config-id", Usage: "The multicloud data transfer config id.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateMulticloudDataTransferConfigRequest{Parent: parent}
							req.MulticloudDataTransferConfigId = cmd.String("multicloud-data-transfer-config-id")
							req.MulticloudDataTransferConfig = &networkconnectivitypb.MulticloudDataTransferConfig{
								Etag:        cmd.String("etag"),
								Description: cmd.String("description"),
							}
							op, err := client.CreateMulticloudDataTransferConfig(ctx, req)
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
						Usage: "update multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateMulticloudDataTransferConfigRequest{}
							req.MulticloudDataTransferConfig = &networkconnectivitypb.MulticloudDataTransferConfig{
								Name:        name,
								Etag:        cmd.String("etag"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateMulticloudDataTransferConfig(ctx, req)
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
						Usage: "delete multicloud-data-transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_config", Usage: "The multicloud_data_transfer_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_config"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteMulticloudDataTransferConfigRequest{Name: name}
							op, err := client.DeleteMulticloudDataTransferConfig(ctx, req)
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
				Name:  "multicloud-data-transfer-supported-services",
				Usage: "Manage multicloud-data-transfer-supported-services resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe multicloud-data-transfer-supported-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "multicloud_data_transfer_supported_service", Usage: "The multicloud_data_transfer_supported_service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferSupportedServices/%s", cmd.String("project"), cmd.String("location"), cmd.String("multicloud_data_transfer_supported_service"))
							client, err := networkconnectivity.NewDataTransferClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetMulticloudDataTransferSupportedServiceRequest{Name: name}
							resp, err := client.GetMulticloudDataTransferSupportedService(ctx, req)
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
						Usage: "list multicloud-data-transfer-supported-services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
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
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
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
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
							client, err := networkconnectivity.NewDataTransferClient(ctx)
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
							client, err := networkconnectivity.NewHubClient(ctx)
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
							client, err := networkconnectivity.NewHubClient(ctx)
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
							client, err := networkconnectivity.NewHubClient(ctx)
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
							client, err := networkconnectivity.NewHubClient(ctx)
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
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
							client, err := networkconnectivity.NewInternalRangeClient(ctx)
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
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
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
				Name:  "policy-based-routes",
				Usage: "Manage policy-based-routes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list policy-based-routes",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy_based_route", Usage: "The policy_based_route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/PolicyBasedRoutes/%s", cmd.String("project"), cmd.String("policy_based_route"))
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetPolicyBasedRouteRequest{Name: name}
							resp, err := client.GetPolicyBasedRoute(ctx, req)
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
						Usage: "create policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy-based-route-id", Usage: "The policy based route id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreatePolicyBasedRouteRequest{Parent: parent}
							req.PolicyBasedRouteId = cmd.String("policy-based-route-id")
							req.PolicyBasedRoute = &networkconnectivitypb.PolicyBasedRoute{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Network:     cmd.String("network"),
								Priority:    int32(cmd.Int("priority")),
							}
							op, err := client.CreatePolicyBasedRoute(ctx, req)
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
						Usage: "delete policy-based-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy_based_route", Usage: "The policy_based_route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/PolicyBasedRoutes/%s", cmd.String("project"), cmd.String("policy_based_route"))
							client, err := networkconnectivity.NewPolicyBasedRoutingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeletePolicyBasedRouteRequest{Name: name}
							op, err := client.DeletePolicyBasedRoute(ctx, req)
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
				Name:  "route-tables",
				Usage: "Manage route-tables resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe route-tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "route_table", Usage: "The route_table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/routeTables/%s", cmd.String("project"), cmd.String("hub"), cmd.String("route_table"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetRouteTableRequest{Name: name}
							resp, err := client.GetRouteTable(ctx, req)
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
						Usage: "list route-tables",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkconnectivitypb.ListRouteTablesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListRouteTables(ctx, req)
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
				Name:  "routes",
				Usage: "Manage routes resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.StringFlag{Name: "route_table", Usage: "The route_table.", Required: true},
							&cli.StringFlag{Name: "route", Usage: "The route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/hubs/%s/routeTables/%s/routes/%s", cmd.String("project"), cmd.String("hub"), cmd.String("route_table"), cmd.String("route"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetRouteRequest{Name: name}
							resp, err := client.GetRoute(ctx, req)
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
						Usage: "list routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global/hubs/%s", cmd.String("project"), cmd.String("hub"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkconnectivitypb.ListRoutesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListRoutes(ctx, req)
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
				Name:  "service-classes",
				Usage: "Manage service-classes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list service-classes",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_class", Usage: "The service_class.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_class"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetServiceClassRequest{Name: name}
							resp, err := client.GetServiceClass(ctx, req)
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
						Usage: "update service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_class", Usage: "The service_class.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_class"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateServiceClassRequest{}
							req.ServiceClass = &networkconnectivitypb.ServiceClass{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Etag:        cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateServiceClass(ctx, req)
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
						Usage: "delete service-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_class", Usage: "The service_class.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_class"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteServiceClassRequest{Name: name}
							op, err := client.DeleteServiceClass(ctx, req)
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
				Name:  "service-connection-maps",
				Usage: "Manage service-connection-maps resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list service-connection-maps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_map", Usage: "The service_connection_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_map"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetServiceConnectionMapRequest{Name: name}
							resp, err := client.GetServiceConnectionMap(ctx, req)
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
						Usage: "create service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service-connection-map-id", Usage: "The service connection map id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "service-class", Usage: "The service class.", Required: false},
							&cli.StringFlag{Name: "token", Usage: "The token.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateServiceConnectionMapRequest{Parent: parent}
							req.ServiceConnectionMapId = cmd.String("service-connection-map-id")
							req.ServiceConnectionMap = &networkconnectivitypb.ServiceConnectionMap{
								Name:         cmd.String("name"),
								Description:  cmd.String("description"),
								ServiceClass: cmd.String("service-class"),
								Token:        cmd.String("token"),
								Etag:         cmd.String("etag"),
							}
							op, err := client.CreateServiceConnectionMap(ctx, req)
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
						Usage: "update service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_map", Usage: "The service_connection_map.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "service-class", Usage: "The service class.", Required: false},
							&cli.StringFlag{Name: "token", Usage: "The token.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_map"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateServiceConnectionMapRequest{}
							req.ServiceConnectionMap = &networkconnectivitypb.ServiceConnectionMap{
								Name:         name,
								Name:         cmd.String("name"),
								Description:  cmd.String("description"),
								ServiceClass: cmd.String("service-class"),
								Token:        cmd.String("token"),
								Etag:         cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("service-class") {
								paths = append(paths, "service_class")
							}
							if cmd.IsSet("token") {
								paths = append(paths, "token")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateServiceConnectionMap(ctx, req)
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
						Usage: "delete service-connection-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_map", Usage: "The service_connection_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_map"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteServiceConnectionMapRequest{Name: name}
							op, err := client.DeleteServiceConnectionMap(ctx, req)
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
				Name:  "service-connection-policies",
				Usage: "Manage service-connection-policies resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list service-connection-policies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_policy", Usage: "The service_connection_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_policy"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetServiceConnectionPolicyRequest{Name: name}
							resp, err := client.GetServiceConnectionPolicy(ctx, req)
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
						Usage: "create service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service-connection-policy-id", Usage: "The service connection policy id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.StringFlag{Name: "service-class", Usage: "The service class.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateServiceConnectionPolicyRequest{Parent: parent}
							req.ServiceConnectionPolicyId = cmd.String("service-connection-policy-id")
							req.ServiceConnectionPolicy = &networkconnectivitypb.ServiceConnectionPolicy{
								Name:         cmd.String("name"),
								Description:  cmd.String("description"),
								Network:      cmd.String("network"),
								ServiceClass: cmd.String("service-class"),
								Etag:         cmd.String("etag"),
							}
							op, err := client.CreateServiceConnectionPolicy(ctx, req)
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
						Usage: "update service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_policy", Usage: "The service_connection_policy.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.StringFlag{Name: "service-class", Usage: "The service class.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_policy"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateServiceConnectionPolicyRequest{}
							req.ServiceConnectionPolicy = &networkconnectivitypb.ServiceConnectionPolicy{
								Name:         name,
								Name:         cmd.String("name"),
								Description:  cmd.String("description"),
								Network:      cmd.String("network"),
								ServiceClass: cmd.String("service-class"),
								Etag:         cmd.String("etag"),
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
							if cmd.IsSet("service-class") {
								paths = append(paths, "service_class")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateServiceConnectionPolicy(ctx, req)
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
						Usage: "delete service-connection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_policy", Usage: "The service_connection_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_policy"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteServiceConnectionPolicyRequest{Name: name}
							op, err := client.DeleteServiceConnectionPolicy(ctx, req)
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
				Name:  "service-connection-tokens",
				Usage: "Manage service-connection-tokens resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_token", Usage: "The service_connection_token.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionTokens/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_token"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetServiceConnectionTokenRequest{Name: name}
							resp, err := client.GetServiceConnectionToken(ctx, req)
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
						Usage: "list service-connection-tokens",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service-connection-token-id", Usage: "The service connection token id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateServiceConnectionTokenRequest{Parent: parent}
							req.ServiceConnectionTokenId = cmd.String("service-connection-token-id")
							req.ServiceConnectionToken = &networkconnectivitypb.ServiceConnectionToken{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Network:     cmd.String("network"),
								Etag:        cmd.String("etag"),
							}
							op, err := client.CreateServiceConnectionToken(ctx, req)
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
						Usage: "delete service-connection-tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "service_connection_token", Usage: "The service_connection_token.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConnectionTokens/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_connection_token"))
							client, err := networkconnectivity.NewCrossNetworkAutomationClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteServiceConnectionTokenRequest{Name: name}
							op, err := client.DeleteServiceConnectionToken(ctx, req)
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
				Name:  "spokes",
				Usage: "Manage spokes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list spokes",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "spoke", Usage: "The spoke.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/spokes/%s", cmd.String("project"), cmd.String("location"), cmd.String("spoke"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.GetSpokeRequest{Name: name}
							resp, err := client.GetSpoke(ctx, req)
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
						Usage: "create spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "spoke-id", Usage: "The spoke id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: false},
							&cli.StringFlag{Name: "group", Usage: "The group.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.CreateSpokeRequest{Parent: parent}
							req.SpokeId = cmd.String("spoke-id")
							req.Spoke = &networkconnectivitypb.Spoke{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Hub:         cmd.String("hub"),
								Group:       cmd.String("group"),
								Etag:        cmd.String("etag"),
							}
							op, err := client.CreateSpoke(ctx, req)
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
						Usage: "update spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "spoke", Usage: "The spoke.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "hub", Usage: "The hub.", Required: false},
							&cli.StringFlag{Name: "group", Usage: "The group.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/spokes/%s", cmd.String("project"), cmd.String("location"), cmd.String("spoke"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.UpdateSpokeRequest{}
							req.Spoke = &networkconnectivitypb.Spoke{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Hub:         cmd.String("hub"),
								Group:       cmd.String("group"),
								Etag:        cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("hub") {
								paths = append(paths, "hub")
							}
							if cmd.IsSet("group") {
								paths = append(paths, "group")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateSpoke(ctx, req)
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
						Usage: "delete spokes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "spoke", Usage: "The spoke.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/spokes/%s", cmd.String("project"), cmd.String("location"), cmd.String("spoke"))
							client, err := networkconnectivity.NewHubClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkconnectivitypb.DeleteSpokeRequest{Name: name}
							op, err := client.DeleteSpoke(ctx, req)
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
		},
	}
}
