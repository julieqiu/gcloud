package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	edgecontainer "cloud.google.com/go/edgecontainer/apiv1"
	"cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
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
				Name:  "edgecontainer",
				Usage: "manage Distributed Cloud Edge Container API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
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
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &edgecontainerpb.ListClustersRequest{Parent: parent}
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
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.GetClusterRequest{Name: name}
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
									&cli.StringFlag{Name: "cluster-id", Usage: "The cluster id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.IntFlag{Name: "default-max-pods-per-node", Usage: "The default max pods per node.", Required: false},
									&cli.StringFlag{Name: "target-version", Usage: "The target version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.CreateClusterRequest{Parent: parent}
									req.ClusterId = cmd.String("cluster-id")
									req.Cluster = &edgecontainerpb.Cluster{
										Name:                  cmd.String("name"),
										DefaultMaxPodsPerNode: int32(cmd.Int("default-max-pods-per-node")),
										TargetVersion:         cmd.String("target-version"),
									}
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
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.IntFlag{Name: "default-max-pods-per-node", Usage: "The default max pods per node.", Required: false},
									&cli.StringFlag{Name: "target-version", Usage: "The target version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.UpdateClusterRequest{}
									req.Cluster = &edgecontainerpb.Cluster{
										Name:                  name,
										Name:                  cmd.String("name"),
										DefaultMaxPodsPerNode: int32(cmd.Int("default-max-pods-per-node")),
										TargetVersion:         cmd.String("target-version"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("default-max-pods-per-node") {
										paths = append(paths, "default_max_pods_per_node")
									}
									if cmd.IsSet("target-version") {
										paths = append(paths, "target_version")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
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
								Name:  "upgrade",
								Usage: "upgrade clusters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "target-version", Usage: "The target version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.UpgradeClusterRequest{Name: name}
									req.TargetVersion = cmd.String("target-version")
									op, err := client.UpgradeCluster(ctx, req)
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
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.DeleteClusterRequest{Name: name}
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
							{
								Name:  "generate-access-token",
								Usage: "generate-access-token clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-access-token...")
									return nil
								},
							},
							{
								Name:  "generate-offline-credential",
								Usage: "generate-offline-credential clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-offline-credential...")
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
									client, err := edgecontainer.NewClient(ctx)
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
									client, err := edgecontainer.NewClient(ctx)
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
						Name:  "machines",
						Usage: "Manage machines resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list machines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &edgecontainerpb.ListMachinesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMachines(ctx, req)
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
								Usage: "describe machines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "machine", Usage: "The machine.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/machines/%s", cmd.String("project"), cmd.String("location"), cmd.String("machine"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.GetMachineRequest{Name: name}
									resp, err := client.GetMachine(ctx, req)
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
						Name:  "node-pools",
						Usage: "Manage node-pools resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list node-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &edgecontainerpb.ListNodePoolsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListNodePools(ctx, req)
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
								Usage: "describe node-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "node_pool", Usage: "The node_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("node_pool"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.GetNodePoolRequest{Name: name}
									resp, err := client.GetNodePool(ctx, req)
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
								Usage: "create node-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "node-pool-id", Usage: "The node pool id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "node-location", Usage: "The node location.", Required: false},
									&cli.IntFlag{Name: "node-count", Usage: "The node count.", Required: true},
									&cli.StringFlag{Name: "machine-filter", Usage: "The machine filter.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.CreateNodePoolRequest{Parent: parent}
									req.NodePoolId = cmd.String("node-pool-id")
									req.NodePool = &edgecontainerpb.NodePool{
										Name:          cmd.String("name"),
										NodeLocation:  cmd.String("node-location"),
										NodeCount:     int32(cmd.Int("node-count")),
										MachineFilter: cmd.String("machine-filter"),
									}
									op, err := client.CreateNodePool(ctx, req)
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
								Usage: "update node-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "node_pool", Usage: "The node_pool.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "node-location", Usage: "The node location.", Required: false},
									&cli.IntFlag{Name: "node-count", Usage: "The node count.", Required: false},
									&cli.StringFlag{Name: "machine-filter", Usage: "The machine filter.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("node_pool"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.UpdateNodePoolRequest{}
									req.NodePool = &edgecontainerpb.NodePool{
										Name:          name,
										Name:          cmd.String("name"),
										NodeLocation:  cmd.String("node-location"),
										NodeCount:     int32(cmd.Int("node-count")),
										MachineFilter: cmd.String("machine-filter"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("node-location") {
										paths = append(paths, "node_location")
									}
									if cmd.IsSet("node-count") {
										paths = append(paths, "node_count")
									}
									if cmd.IsSet("machine-filter") {
										paths = append(paths, "machine_filter")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateNodePool(ctx, req)
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
								Usage: "delete node-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "node_pool", Usage: "The node_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("node_pool"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.DeleteNodePoolRequest{Name: name}
									op, err := client.DeleteNodePool(ctx, req)
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
									client, err := edgecontainer.NewClient(ctx)
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
									client, err := edgecontainer.NewClient(ctx)
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
									client, err := edgecontainer.NewClient(ctx)
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
									client, err := edgecontainer.NewClient(ctx)
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
						Name:  "server-config",
						Usage: "Manage server-config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe server-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "vpn-connections",
						Usage: "Manage vpn-connections resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list vpn-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &edgecontainerpb.ListVpnConnectionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVpnConnections(ctx, req)
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
								Usage: "describe vpn-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vpn_connection", Usage: "The vpn_connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vpnConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpn_connection"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.GetVpnConnectionRequest{Name: name}
									resp, err := client.GetVpnConnection(ctx, req)
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
								Usage: "create vpn-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vpn-connection-id", Usage: "The vpn connection id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "nat-gateway-ip", Usage: "The nat gateway ip.", Required: false},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: false},
									&cli.StringFlag{Name: "vpc", Usage: "The vpc.", Required: false},
									&cli.BoolFlag{Name: "enable-high-availability", Usage: "The enable high availability.", Required: false},
									&cli.StringFlag{Name: "router", Usage: "The router.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.CreateVpnConnectionRequest{Parent: parent}
									req.VpnConnectionId = cmd.String("vpn-connection-id")
									req.VpnConnection = &edgecontainerpb.VpnConnection{
										Name:                   cmd.String("name"),
										NatGatewayIp:           cmd.String("nat-gateway-ip"),
										Cluster:                cmd.String("cluster"),
										Vpc:                    cmd.String("vpc"),
										EnableHighAvailability: cmd.Bool("enable-high-availability"),
										Router:                 cmd.String("router"),
									}
									op, err := client.CreateVpnConnection(ctx, req)
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
								Usage: "delete vpn-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vpn_connection", Usage: "The vpn_connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vpnConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpn_connection"))
									client, err := edgecontainer.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &edgecontainerpb.DeleteVpnConnectionRequest{Name: name}
									op, err := client.DeleteVpnConnection(ctx, req)
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
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
