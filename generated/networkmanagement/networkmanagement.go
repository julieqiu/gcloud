package networkmanagement

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	networkmanagement "cloud.google.com/go/networkmanagement/apiv1"
	"cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud networkmanagement command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "networkmanagement",
		Usage: "manage Network Management API resources",
		Commands: []*cli.Command{
			{
				Name:  "connectivity-tests",
				Usage: "Manage connectivity-tests resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list connectivity-tests",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "test", Usage: "The test.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("test"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.GetConnectivityTestRequest{Name: name}
							resp, err := client.GetConnectivityTest(ctx, req)
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
						Usage: "create connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "test-id", Usage: "The test id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "protocol", Usage: "The protocol.", Required: false},
							&cli.BoolFlag{Name: "round-trip", Usage: "The round trip.", Required: false},
							&cli.BoolFlag{Name: "bypass-firewall-checks", Usage: "The bypass firewall checks.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.CreateConnectivityTestRequest{Parent: parent}
							req.TestId = cmd.String("test-id")
							req.Resource = &networkmanagementpb.ConnectivityTest{
								Description:          cmd.String("description"),
								Protocol:             cmd.String("protocol"),
								RoundTrip:            cmd.Bool("round-trip"),
								BypassFirewallChecks: cmd.Bool("bypass-firewall-checks"),
							}
							op, err := client.CreateConnectivityTest(ctx, req)
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
						Usage: "update connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "test", Usage: "The test.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "protocol", Usage: "The protocol.", Required: false},
							&cli.BoolFlag{Name: "round-trip", Usage: "The round trip.", Required: false},
							&cli.BoolFlag{Name: "bypass-firewall-checks", Usage: "The bypass firewall checks.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("test"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.UpdateConnectivityTestRequest{}
							req.Resource = &networkmanagementpb.ConnectivityTest{
								Name:                 name,
								Description:          cmd.String("description"),
								Protocol:             cmd.String("protocol"),
								RoundTrip:            cmd.Bool("round-trip"),
								BypassFirewallChecks: cmd.Bool("bypass-firewall-checks"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("protocol") {
								paths = append(paths, "protocol")
							}
							if cmd.IsSet("round-trip") {
								paths = append(paths, "round_trip")
							}
							if cmd.IsSet("bypass-firewall-checks") {
								paths = append(paths, "bypass_firewall_checks")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateConnectivityTest(ctx, req)
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
						Name:  "rerun",
						Usage: "rerun connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "test", Usage: "The test.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("test"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.RerunConnectivityTestRequest{Name: name}
							op, err := client.RerunConnectivityTest(ctx, req)
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
						Usage: "delete connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "test", Usage: "The test.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("test"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.DeleteConnectivityTestRequest{Name: name}
							op, err := client.DeleteConnectivityTest(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
						Usage: "test-iam-permissions connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
						Usage: "set-iam-policy connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
						Usage: "test-iam-permissions connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
						Usage: "set-iam-policy connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
						Usage: "test-iam-permissions connectivity-tests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connectivityTest", Usage: "The connectivityTest.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", cmd.String("project"), cmd.String("connectivityTest"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewReachabilityClient(ctx)
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
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
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
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
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
				Name:  "vpc-flow-logs-configs",
				Usage: "Manage vpc-flow-logs-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkmanagementpb.ListVpcFlowLogsConfigsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListVpcFlowLogsConfigs(ctx, req)
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
						Usage: "describe vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc_flow_logs_config", Usage: "The vpc_flow_logs_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpc_flow_logs_config"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.GetVpcFlowLogsConfigRequest{Name: name}
							resp, err := client.GetVpcFlowLogsConfig(ctx, req)
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
						Usage: "create vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc-flow-logs-config-id", Usage: "The vpc flow logs config id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "filter-expr", Usage: "The filter expr.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.CreateVpcFlowLogsConfigRequest{Parent: parent}
							req.VpcFlowLogsConfigId = cmd.String("vpc-flow-logs-config-id")
							req.VpcFlowLogsConfig = &networkmanagementpb.VpcFlowLogsConfig{
								Description: cmd.String("description"),
								FilterExpr:  cmd.String("filter-expr"),
							}
							op, err := client.CreateVpcFlowLogsConfig(ctx, req)
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
						Usage: "update vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc_flow_logs_config", Usage: "The vpc_flow_logs_config.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "filter-expr", Usage: "The filter expr.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpc_flow_logs_config"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.UpdateVpcFlowLogsConfigRequest{}
							req.VpcFlowLogsConfig = &networkmanagementpb.VpcFlowLogsConfig{
								Name:        name,
								Description: cmd.String("description"),
								FilterExpr:  cmd.String("filter-expr"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("filter-expr") {
								paths = append(paths, "filter_expr")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateVpcFlowLogsConfig(ctx, req)
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
						Usage: "delete vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc_flow_logs_config", Usage: "The vpc_flow_logs_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpc_flow_logs_config"))
							client, err := networkmanagement.NewVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.DeleteVpcFlowLogsConfigRequest{Name: name}
							op, err := client.DeleteVpcFlowLogsConfig(ctx, req)
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
						Name:  "query-org-vpc-flow-logs-configs",
						Usage: "query-org-vpc-flow-logs-configs vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing query-org-vpc-flow-logs-configs on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "show-effective-flow-logs-configs",
						Usage: "show-effective-flow-logs-configs vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing show-effective-flow-logs-configs on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &networkmanagementpb.ListVpcFlowLogsConfigsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListVpcFlowLogsConfigs(ctx, req)
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
						Usage: "describe vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc_flow_logs_config", Usage: "The vpc_flow_logs_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpc_flow_logs_config"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.GetVpcFlowLogsConfigRequest{Name: name}
							resp, err := client.GetVpcFlowLogsConfig(ctx, req)
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
						Usage: "create vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc-flow-logs-config-id", Usage: "The vpc flow logs config id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "filter-expr", Usage: "The filter expr.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.CreateVpcFlowLogsConfigRequest{Parent: parent}
							req.VpcFlowLogsConfigId = cmd.String("vpc-flow-logs-config-id")
							req.VpcFlowLogsConfig = &networkmanagementpb.VpcFlowLogsConfig{
								Description: cmd.String("description"),
								FilterExpr:  cmd.String("filter-expr"),
							}
							op, err := client.CreateVpcFlowLogsConfig(ctx, req)
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
						Usage: "update vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc_flow_logs_config", Usage: "The vpc_flow_logs_config.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "filter-expr", Usage: "The filter expr.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpc_flow_logs_config"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.UpdateVpcFlowLogsConfigRequest{}
							req.VpcFlowLogsConfig = &networkmanagementpb.VpcFlowLogsConfig{
								Name:        name,
								Description: cmd.String("description"),
								FilterExpr:  cmd.String("filter-expr"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("filter-expr") {
								paths = append(paths, "filter_expr")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateVpcFlowLogsConfig(ctx, req)
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
						Usage: "delete vpc-flow-logs-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "vpc_flow_logs_config", Usage: "The vpc_flow_logs_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpcFlowLogsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpc_flow_logs_config"))
							client, err := networkmanagement.NewOrganizationVpcFlowLogsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &networkmanagementpb.DeleteVpcFlowLogsConfigRequest{Name: name}
							op, err := client.DeleteVpcFlowLogsConfig(ctx, req)
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
