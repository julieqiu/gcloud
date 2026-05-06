package vmwareengine

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	vmwareengine "cloud.google.com/go/vmwareengine/apiv1"
	"cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud vmwareengine command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "vmwareengine",
		Usage: "manage VMware Engine API resources",
		Commands: []*cli.Command{
			{
				Name:  "clusters",
				Usage: "Manage clusters resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListClusters(ctx, req)
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
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("cluster"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetClusterRequest{Name: name}
							resp, err := client.GetCluster(ctx, req)
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
						Usage: "create clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "cluster-id", Usage: "The cluster id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateClusterRequest{Parent: parent}
							req.ClusterId = cmd.String("cluster-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							op, err := client.CreateCluster(ctx, req)
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
						Usage: "update clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("cluster"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateClusterRequest{}
							req.Cluster = &vmwareenginepb.Cluster{
								Name: name,
							}
							op, err := client.UpdateCluster(ctx, req)
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
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("cluster"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteClusterRequest{Name: name}
							op, err := client.DeleteCluster(ctx, req)
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
				Name:  "dns-bind-permission",
				Usage: "Manage dns-bind-permission resources",
				Commands: []*cli.Command{
					{
						Name:  "grant",
						Usage: "grant dns-bind-permission",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsBindPermission", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GrantDnsBindPermissionRequest{Name: name}
							op, err := client.GrantDnsBindPermission(ctx, req)
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
						Name:  "describe",
						Usage: "describe dns-bind-permission",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsBindPermission", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetDnsBindPermissionRequest{Name: name}
							resp, err := client.GetDnsBindPermission(ctx, req)
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
						Name:  "revoke",
						Usage: "revoke dns-bind-permission",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsBindPermission", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.RevokeDnsBindPermissionRequest{Name: name}
							op, err := client.RevokeDnsBindPermission(ctx, req)
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
				Name:  "dns-forwarding",
				Usage: "Manage dns-forwarding resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe dns-forwarding",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/dnsForwarding", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetDnsForwardingRequest{Name: name}
							resp, err := client.GetDnsForwarding(ctx, req)
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
						Usage: "update dns-forwarding",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/dnsForwarding", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateDnsForwardingRequest{}
							req.DnsForwarding = &vmwareenginepb.DnsForwarding{
								Name: name,
							}
							op, err := client.UpdateDnsForwarding(ctx, req)
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
				Name:  "external-access-rules",
				Usage: "Manage external-access-rules resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListExternalAccessRulesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListExternalAccessRules(ctx, req)
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
						Usage: "describe external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
							&cli.StringFlag{Name: "external_access_rule", Usage: "The external_access_rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"), cmd.String("external_access_rule"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetExternalAccessRuleRequest{Name: name}
							resp, err := client.GetExternalAccessRule(ctx, req)
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
						Usage: "create external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
							&cli.StringFlag{Name: "external-access-rule-id", Usage: "The external access rule id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: false},
							&cli.StringFlag{Name: "ip-protocol", Usage: "The ip protocol.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateExternalAccessRuleRequest{Parent: parent}
							req.ExternalAccessRuleId = cmd.String("external-access-rule-id")
							req.ExternalAccessRule = &vmwareenginepb.ExternalAccessRule{
								Description: cmd.String("description"),
								Priority:    int32(cmd.Int("priority")),
								IpProtocol:  cmd.String("ip-protocol"),
							}
							op, err := client.CreateExternalAccessRule(ctx, req)
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
						Usage: "update external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
							&cli.StringFlag{Name: "external_access_rule", Usage: "The external_access_rule.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: false},
							&cli.StringFlag{Name: "ip-protocol", Usage: "The ip protocol.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"), cmd.String("external_access_rule"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateExternalAccessRuleRequest{}
							req.ExternalAccessRule = &vmwareenginepb.ExternalAccessRule{
								Name:        name,
								Description: cmd.String("description"),
								Priority:    int32(cmd.Int("priority")),
								IpProtocol:  cmd.String("ip-protocol"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("priority") {
								paths = append(paths, "priority")
							}
							if cmd.IsSet("ip-protocol") {
								paths = append(paths, "ip_protocol")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateExternalAccessRule(ctx, req)
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
						Usage: "delete external-access-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
							&cli.StringFlag{Name: "external_access_rule", Usage: "The external_access_rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s/externalAccessRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"), cmd.String("external_access_rule"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteExternalAccessRuleRequest{Name: name}
							op, err := client.DeleteExternalAccessRule(ctx, req)
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
				Name:  "external-addresses",
				Usage: "Manage external-addresses resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListExternalAddressesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListExternalAddresses(ctx, req)
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
						Usage: "describe external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "external_address", Usage: "The external_address.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("external_address"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetExternalAddressRequest{Name: name}
							resp, err := client.GetExternalAddress(ctx, req)
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
						Usage: "create external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "external-address-id", Usage: "The external address id.", Required: true},
							&cli.StringFlag{Name: "internal-ip", Usage: "The internal ip.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateExternalAddressRequest{Parent: parent}
							req.ExternalAddressId = cmd.String("external-address-id")
							req.ExternalAddress = &vmwareenginepb.ExternalAddress{
								InternalIp:  cmd.String("internal-ip"),
								Description: cmd.String("description"),
							}
							op, err := client.CreateExternalAddress(ctx, req)
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
						Usage: "update external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "external_address", Usage: "The external_address.", Required: true},
							&cli.StringFlag{Name: "internal-ip", Usage: "The internal ip.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("external_address"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateExternalAddressRequest{}
							req.ExternalAddress = &vmwareenginepb.ExternalAddress{
								Name:        name,
								InternalIp:  cmd.String("internal-ip"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("internal-ip") {
								paths = append(paths, "internal_ip")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateExternalAddress(ctx, req)
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
						Usage: "delete external-addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "external_address", Usage: "The external_address.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/externalAddresses/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("external_address"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteExternalAddressRequest{Name: name}
							op, err := client.DeleteExternalAddress(ctx, req)
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
				Name:  "hcx-activation-keys",
				Usage: "Manage hcx-activation-keys resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create hcx-activation-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "hcx-activation-key-id", Usage: "The hcx activation key id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateHcxActivationKeyRequest{Parent: parent}
							req.HcxActivationKeyId = cmd.String("hcx-activation-key-id")
							op, err := client.CreateHcxActivationKey(ctx, req)
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
						Name:  "list",
						Usage: "list hcx-activation-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListHcxActivationKeysRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListHcxActivationKeys(ctx, req)
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
						Usage: "describe hcx-activation-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "hcx_activation_key", Usage: "The hcx_activation_key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/hcxActivationKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("hcx_activation_key"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetHcxActivationKeyRequest{Name: name}
							resp, err := client.GetHcxActivationKey(ctx, req)
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
							client, err := vmwareengine.NewClient(ctx)
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
							client, err := vmwareengine.NewClient(ctx)
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
				Name:  "logging-servers",
				Usage: "Manage logging-servers resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListLoggingServersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListLoggingServers(ctx, req)
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
						Usage: "describe logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "logging_server", Usage: "The logging_server.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/loggingServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("logging_server"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetLoggingServerRequest{Name: name}
							resp, err := client.GetLoggingServer(ctx, req)
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
						Usage: "create logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "logging-server-id", Usage: "The logging server id.", Required: true},
							&cli.StringFlag{Name: "hostname", Usage: "The hostname.", Required: true},
							&cli.IntFlag{Name: "port", Usage: "The port.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateLoggingServerRequest{Parent: parent}
							req.LoggingServerId = cmd.String("logging-server-id")
							req.LoggingServer = &vmwareenginepb.LoggingServer{
								Hostname: cmd.String("hostname"),
								Port:     int32(cmd.Int("port")),
							}
							op, err := client.CreateLoggingServer(ctx, req)
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
						Usage: "update logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "logging_server", Usage: "The logging_server.", Required: true},
							&cli.StringFlag{Name: "hostname", Usage: "The hostname.", Required: false},
							&cli.IntFlag{Name: "port", Usage: "The port.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/loggingServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("logging_server"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateLoggingServerRequest{}
							req.LoggingServer = &vmwareenginepb.LoggingServer{
								Name:     name,
								Hostname: cmd.String("hostname"),
								Port:     int32(cmd.Int("port")),
							}
							var paths []string
							if cmd.IsSet("hostname") {
								paths = append(paths, "hostname")
							}
							if cmd.IsSet("port") {
								paths = append(paths, "port")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateLoggingServer(ctx, req)
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
						Usage: "delete logging-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "logging_server", Usage: "The logging_server.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/loggingServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("logging_server"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteLoggingServerRequest{Name: name}
							op, err := client.DeleteLoggingServer(ctx, req)
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
				Name:  "management-dns-zone-bindings",
				Usage: "Manage management-dns-zone-bindings resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListManagementDnsZoneBindingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListManagementDnsZoneBindings(ctx, req)
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
						Usage: "describe management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "management_dns_zone_binding", Usage: "The management_dns_zone_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("management_dns_zone_binding"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetManagementDnsZoneBindingRequest{Name: name}
							resp, err := client.GetManagementDnsZoneBinding(ctx, req)
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
						Usage: "create management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "management-dns-zone-binding-id", Usage: "The management dns zone binding id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateManagementDnsZoneBindingRequest{Parent: parent}
							req.ManagementDnsZoneBindingId = cmd.String("management-dns-zone-binding-id")
							req.ManagementDnsZoneBinding = &vmwareenginepb.ManagementDnsZoneBinding{
								Description: cmd.String("description"),
							}
							op, err := client.CreateManagementDnsZoneBinding(ctx, req)
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
						Usage: "update management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "management_dns_zone_binding", Usage: "The management_dns_zone_binding.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("management_dns_zone_binding"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateManagementDnsZoneBindingRequest{}
							req.ManagementDnsZoneBinding = &vmwareenginepb.ManagementDnsZoneBinding{
								Name:        name,
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateManagementDnsZoneBinding(ctx, req)
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
						Usage: "delete management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "management_dns_zone_binding", Usage: "The management_dns_zone_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("management_dns_zone_binding"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteManagementDnsZoneBindingRequest{Name: name}
							op, err := client.DeleteManagementDnsZoneBinding(ctx, req)
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
						Name:  "repair",
						Usage: "repair management-dns-zone-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "management_dns_zone_binding", Usage: "The management_dns_zone_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/managementDnsZoneBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("management_dns_zone_binding"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.RepairManagementDnsZoneBindingRequest{Name: name}
							op, err := client.RepairManagementDnsZoneBinding(ctx, req)
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
				Name:  "network-peerings",
				Usage: "Manage network-peerings resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_peering", Usage: "The network_peering.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_peering"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetNetworkPeeringRequest{Name: name}
							resp, err := client.GetNetworkPeering(ctx, req)
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
						Usage: "list network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListNetworkPeeringsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListNetworkPeerings(ctx, req)
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
						Name:  "create",
						Usage: "create network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network-peering-id", Usage: "The network peering id.", Required: true},
							&cli.StringFlag{Name: "peer-network", Usage: "The peer network.", Required: true},
							&cli.BoolFlag{Name: "export-custom-routes", Usage: "The export custom routes.", Required: false},
							&cli.BoolFlag{Name: "import-custom-routes", Usage: "The import custom routes.", Required: false},
							&cli.BoolFlag{Name: "exchange-subnet-routes", Usage: "The exchange subnet routes.", Required: false},
							&cli.BoolFlag{Name: "export-custom-routes-with-public-ip", Usage: "The export custom routes with public ip.", Required: false},
							&cli.BoolFlag{Name: "import-custom-routes-with-public-ip", Usage: "The import custom routes with public ip.", Required: false},
							&cli.IntFlag{Name: "peer-mtu", Usage: "The peer mtu.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The vmware engine network.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateNetworkPeeringRequest{Parent: parent}
							req.NetworkPeeringId = cmd.String("network-peering-id")
							req.NetworkPeering = &vmwareenginepb.NetworkPeering{
								PeerNetwork:                    cmd.String("peer-network"),
								ExportCustomRoutes:             cmd.Bool("export-custom-routes"),
								ImportCustomRoutes:             cmd.Bool("import-custom-routes"),
								ExchangeSubnetRoutes:           cmd.Bool("exchange-subnet-routes"),
								ExportCustomRoutesWithPublicIp: cmd.Bool("export-custom-routes-with-public-ip"),
								ImportCustomRoutesWithPublicIp: cmd.Bool("import-custom-routes-with-public-ip"),
								PeerMtu:                        int32(cmd.Int("peer-mtu")),
								VmwareEngineNetwork:            cmd.String("vmware-engine-network"),
								Description:                    cmd.String("description"),
							}
							op, err := client.CreateNetworkPeering(ctx, req)
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
						Usage: "delete network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_peering", Usage: "The network_peering.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_peering"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteNetworkPeeringRequest{Name: name}
							op, err := client.DeleteNetworkPeering(ctx, req)
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
						Name:  "update",
						Usage: "update network-peerings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_peering", Usage: "The network_peering.", Required: true},
							&cli.StringFlag{Name: "peer-network", Usage: "The peer network.", Required: false},
							&cli.BoolFlag{Name: "export-custom-routes", Usage: "The export custom routes.", Required: false},
							&cli.BoolFlag{Name: "import-custom-routes", Usage: "The import custom routes.", Required: false},
							&cli.BoolFlag{Name: "exchange-subnet-routes", Usage: "The exchange subnet routes.", Required: false},
							&cli.BoolFlag{Name: "export-custom-routes-with-public-ip", Usage: "The export custom routes with public ip.", Required: false},
							&cli.BoolFlag{Name: "import-custom-routes-with-public-ip", Usage: "The import custom routes with public ip.", Required: false},
							&cli.IntFlag{Name: "peer-mtu", Usage: "The peer mtu.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The vmware engine network.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPeerings/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_peering"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateNetworkPeeringRequest{}
							req.NetworkPeering = &vmwareenginepb.NetworkPeering{
								Name:                           name,
								PeerNetwork:                    cmd.String("peer-network"),
								ExportCustomRoutes:             cmd.Bool("export-custom-routes"),
								ImportCustomRoutes:             cmd.Bool("import-custom-routes"),
								ExchangeSubnetRoutes:           cmd.Bool("exchange-subnet-routes"),
								ExportCustomRoutesWithPublicIp: cmd.Bool("export-custom-routes-with-public-ip"),
								ImportCustomRoutesWithPublicIp: cmd.Bool("import-custom-routes-with-public-ip"),
								PeerMtu:                        int32(cmd.Int("peer-mtu")),
								VmwareEngineNetwork:            cmd.String("vmware-engine-network"),
								Description:                    cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("peer-network") {
								paths = append(paths, "peer_network")
							}
							if cmd.IsSet("export-custom-routes") {
								paths = append(paths, "export_custom_routes")
							}
							if cmd.IsSet("import-custom-routes") {
								paths = append(paths, "import_custom_routes")
							}
							if cmd.IsSet("exchange-subnet-routes") {
								paths = append(paths, "exchange_subnet_routes")
							}
							if cmd.IsSet("export-custom-routes-with-public-ip") {
								paths = append(paths, "export_custom_routes_with_public_ip")
							}
							if cmd.IsSet("import-custom-routes-with-public-ip") {
								paths = append(paths, "import_custom_routes_with_public_ip")
							}
							if cmd.IsSet("peer-mtu") {
								paths = append(paths, "peer_mtu")
							}
							if cmd.IsSet("vmware-engine-network") {
								paths = append(paths, "vmware_engine_network")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateNetworkPeering(ctx, req)
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
				Name:  "network-policies",
				Usage: "Manage network-policies resources",
				Commands: []*cli.Command{
					{
						Name:  "fetch-external-addresses",
						Usage: "fetch-external-addresses network-policies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch-external-addresses...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetNetworkPolicyRequest{Name: name}
							resp, err := client.GetNetworkPolicy(ctx, req)
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
						Usage: "list network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListNetworkPoliciesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListNetworkPolicies(ctx, req)
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
						Name:  "create",
						Usage: "create network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network-policy-id", Usage: "The network policy id.", Required: true},
							&cli.StringFlag{Name: "edge-services-cidr", Usage: "The edge services cidr.", Required: true},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The vmware engine network.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateNetworkPolicyRequest{Parent: parent}
							req.NetworkPolicyId = cmd.String("network-policy-id")
							req.NetworkPolicy = &vmwareenginepb.NetworkPolicy{
								EdgeServicesCidr:    cmd.String("edge-services-cidr"),
								VmwareEngineNetwork: cmd.String("vmware-engine-network"),
								Description:         cmd.String("description"),
							}
							op, err := client.CreateNetworkPolicy(ctx, req)
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
						Usage: "update network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
							&cli.StringFlag{Name: "edge-services-cidr", Usage: "The edge services cidr.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The vmware engine network.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateNetworkPolicyRequest{}
							req.NetworkPolicy = &vmwareenginepb.NetworkPolicy{
								Name:                name,
								EdgeServicesCidr:    cmd.String("edge-services-cidr"),
								VmwareEngineNetwork: cmd.String("vmware-engine-network"),
								Description:         cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("edge-services-cidr") {
								paths = append(paths, "edge_services_cidr")
							}
							if cmd.IsSet("vmware-engine-network") {
								paths = append(paths, "vmware_engine_network")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateNetworkPolicy(ctx, req)
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
						Usage: "delete network-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "network_policy", Usage: "The network_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/networkPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("network_policy"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteNetworkPolicyRequest{Name: name}
							op, err := client.DeleteNetworkPolicy(ctx, req)
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
				Name:  "node-types",
				Usage: "Manage node-types resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list node-types",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe node-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node_type", Usage: "The node_type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodeTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node_type"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetNodeTypeRequest{Name: name}
							resp, err := client.GetNodeType(ctx, req)
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
				Name:  "nodes",
				Usage: "Manage nodes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListNodesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListNodes(ctx, req)
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
						Usage: "describe nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/clusters/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("cluster"), cmd.String("node"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetNodeRequest{Name: name}
							resp, err := client.GetNode(ctx, req)
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
							client, err := vmwareengine.NewClient(ctx)
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
							client, err := vmwareengine.NewClient(ctx)
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
							client, err := vmwareengine.NewClient(ctx)
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
				},
			},
			{
				Name:  "peering-routes",
				Usage: "Manage peering-routes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list peering-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListPeeringRoutesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPeeringRoutes(ctx, req)
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
						Name:  "list",
						Usage: "list peering-routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListPrivateConnectionPeeringRoutesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPrivateConnectionPeeringRoutes(ctx, req)
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
				Name:  "private-clouds",
				Usage: "Manage private-clouds resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list private-clouds",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetPrivateCloudRequest{Name: name}
							resp, err := client.GetPrivateCloud(ctx, req)
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
						Usage: "create private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private-cloud-id", Usage: "The private cloud id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreatePrivateCloudRequest{Parent: parent}
							req.PrivateCloudId = cmd.String("private-cloud-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.PrivateCloud = &vmwareenginepb.PrivateCloud{
								Description: cmd.String("description"),
							}
							op, err := client.CreatePrivateCloud(ctx, req)
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
						Usage: "update private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdatePrivateCloudRequest{}
							req.PrivateCloud = &vmwareenginepb.PrivateCloud{
								Name:        name,
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdatePrivateCloud(ctx, req)
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
						Usage: "delete private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeletePrivateCloudRequest{Name: name}
							op, err := client.DeletePrivateCloud(ctx, req)
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
						Name:  "undelete",
						Usage: "undelete private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UndeletePrivateCloudRequest{Name: name}
							op, err := client.UndeletePrivateCloud(ctx, req)
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
						Name:  "show-nsx-credentials",
						Usage: "show-nsx-credentials private-clouds",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing show-nsx-credentials...")
							return nil
						},
					},
					{
						Name:  "show-vcenter-credentials",
						Usage: "show-vcenter-credentials private-clouds",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing show-vcenter-credentials...")
							return nil
						},
					},
					{
						Name:  "reset-nsx-credentials",
						Usage: "reset-nsx-credentials private-clouds",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing reset-nsx-credentials...")
							return nil
						},
					},
					{
						Name:  "reset-vcenter-credentials",
						Usage: "reset-vcenter-credentials private-clouds",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing reset-vcenter-credentials...")
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "privateCloud", Usage: "The privateCloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("privateCloud"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "privateCloud", Usage: "The privateCloud.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("privateCloud"))
							client, err := vmwareengine.NewClient(ctx)
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
						Usage: "test-iam-permissions private-clouds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "privateCloud", Usage: "The privateCloud.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s", cmd.String("project"), cmd.String("location"), cmd.String("privateCloud"))
							client, err := vmwareengine.NewClient(ctx)
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
				Name:  "private-connections",
				Usage: "Manage private-connections resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private-connection-id", Usage: "The private connection id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The vmware engine network.", Required: true},
							&cli.StringFlag{Name: "service-network", Usage: "The service network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreatePrivateConnectionRequest{Parent: parent}
							req.PrivateConnectionId = cmd.String("private-connection-id")
							req.PrivateConnection = &vmwareenginepb.PrivateConnection{
								Description:         cmd.String("description"),
								VmwareEngineNetwork: cmd.String("vmware-engine-network"),
								ServiceNetwork:      cmd.String("service-network"),
							}
							op, err := client.CreatePrivateConnection(ctx, req)
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
						Name:  "describe",
						Usage: "describe private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_connection", Usage: "The private_connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_connection"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetPrivateConnectionRequest{Name: name}
							resp, err := client.GetPrivateConnection(ctx, req)
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
						Usage: "list private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListPrivateConnectionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPrivateConnections(ctx, req)
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
						Usage: "update private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_connection", Usage: "The private_connection.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "vmware-engine-network", Usage: "The vmware engine network.", Required: false},
							&cli.StringFlag{Name: "service-network", Usage: "The service network.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_connection"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdatePrivateConnectionRequest{}
							req.PrivateConnection = &vmwareenginepb.PrivateConnection{
								Name:                name,
								Description:         cmd.String("description"),
								VmwareEngineNetwork: cmd.String("vmware-engine-network"),
								ServiceNetwork:      cmd.String("service-network"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("vmware-engine-network") {
								paths = append(paths, "vmware_engine_network")
							}
							if cmd.IsSet("service-network") {
								paths = append(paths, "service_network")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdatePrivateConnection(ctx, req)
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
						Usage: "delete private-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_connection", Usage: "The private_connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_connection"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeletePrivateConnectionRequest{Name: name}
							op, err := client.DeletePrivateConnection(ctx, req)
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
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &vmwareenginepb.ListSubnetsRequest{Parent: parent}
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
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("subnet"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetSubnetRequest{Name: name}
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
						Name:  "update",
						Usage: "update subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "private_cloud", Usage: "The private_cloud.", Required: true},
							&cli.StringFlag{Name: "subnet", Usage: "The subnet.", Required: true},
							&cli.StringFlag{Name: "ip-cidr-range", Usage: "The ip cidr range.", Required: false},
							&cli.StringFlag{Name: "gateway-ip", Usage: "The gateway ip.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/privateClouds/%s/subnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_cloud"), cmd.String("subnet"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateSubnetRequest{}
							req.Subnet = &vmwareenginepb.Subnet{
								Name:        name,
								IpCidrRange: cmd.String("ip-cidr-range"),
								GatewayIp:   cmd.String("gateway-ip"),
							}
							var paths []string
							if cmd.IsSet("ip-cidr-range") {
								paths = append(paths, "ip_cidr_range")
							}
							if cmd.IsSet("gateway-ip") {
								paths = append(paths, "gateway_ip")
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
				},
			},
			{
				Name:  "vmware-engine-networks",
				Usage: "Manage vmware-engine-networks resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vmware-engine-network-id", Usage: "The vmware engine network id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.CreateVmwareEngineNetworkRequest{Parent: parent}
							req.VmwareEngineNetworkId = cmd.String("vmware-engine-network-id")
							req.VmwareEngineNetwork = &vmwareenginepb.VmwareEngineNetwork{
								Description: cmd.String("description"),
								Etag:        cmd.String("etag"),
							}
							op, err := client.CreateVmwareEngineNetwork(ctx, req)
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
						Usage: "update vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vmware_engine_network", Usage: "The vmware_engine_network.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("vmware_engine_network"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.UpdateVmwareEngineNetworkRequest{}
							req.VmwareEngineNetwork = &vmwareenginepb.VmwareEngineNetwork{
								Name:        name,
								Description: cmd.String("description"),
								Etag:        cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateVmwareEngineNetwork(ctx, req)
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
						Usage: "delete vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vmware_engine_network", Usage: "The vmware_engine_network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("vmware_engine_network"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.DeleteVmwareEngineNetworkRequest{Name: name}
							op, err := client.DeleteVmwareEngineNetwork(ctx, req)
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
						Name:  "describe",
						Usage: "describe vmware-engine-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vmware_engine_network", Usage: "The vmware_engine_network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vmwareEngineNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("vmware_engine_network"))
							client, err := vmwareengine.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &vmwareenginepb.GetVmwareEngineNetworkRequest{Name: name}
							resp, err := client.GetVmwareEngineNetwork(ctx, req)
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
						Usage: "list vmware-engine-networks",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
		},
	}
}
