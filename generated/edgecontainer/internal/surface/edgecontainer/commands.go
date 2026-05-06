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

package edgecontainer

import (
	edgecontainer "cloud.google.com/go/edgecontainer/apiv1"
	"cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the edgecontainer command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "edgecontainer",
		Usage: "manage Distributed Cloud Edge Container API resources",
		Commands: []*cli.Command{
			{
				Name:  "clusters",
				Usage: "Manage clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only resources matching this filter will be listed.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the order in which resources will be listed.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of resources to list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.ListClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListClusters(ctx, req)
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
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.GetClusterRequest{
								Name: name,
							}

							resp, err := client.GetCluster(ctx, req)
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
						Usage: "create clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-id", Usage: "A client-specified unique identifier for the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.CreateClusterRequest{
								Parent:    parent,
								ClusterId: cmd.String("cluster-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateCluster(ctx, req)
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
						Usage: "update clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "cluster.name" not yet supported.
							cluster_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							fmt.Printf("Executing update on %s\n", cluster_name)
							return nil
						},
					},

					{
						Name:  "upgrade",
						Usage: "upgrade clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
							&cli.StringFlag{Name: "schedule", Usage: "The schedule for the upgrade.", Required: false},
							&cli.StringFlag{Name: "target-version", Usage: "The version the cluster is going to be upgraded to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.UpgradeClusterRequest{
								Name:          name,
								TargetVersion: cmd.String("target-version"),
								Schedule:      edgecontainerpb.UpgradeClusterRequest_Schedule(edgecontainerpb.UpgradeClusterRequest_Schedule_value[cmd.String("schedule")]),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.UpgradeCluster(ctx, req)
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
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.DeleteClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteCluster(ctx, req)
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
						Name:  "generate-access-token",
						Usage: "generate-access-token clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							cluster := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							fmt.Printf("Executing generate-access-token on %s\n", cluster)
							return nil
						},
					},

					{
						Name:  "generate-offline-credential",
						Usage: "generate-offline-credential clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							cluster := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							fmt.Printf("Executing generate-offline-credential on %s\n", cluster)
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
				Name:  "machines",
				Usage: "Manage machines resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list machines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only resources matching this filter will be listed.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the order in which resources will be listed.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of resources to list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.ListMachinesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMachines(ctx, req)
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
						Usage: "describe machines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "machine", Usage: "The ID of the machine.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/machines/%s", cmd.String("project"), cmd.String("location"), cmd.String("machine"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.GetMachineRequest{
								Name: name,
							}

							resp, err := client.GetMachine(ctx, req)
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
				Name:  "node-pools",
				Usage: "Manage node-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only resources matching this filter will be listed.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the order in which resources will be listed.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of resources to list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.ListNodePoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListNodePools(ctx, req)
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
						Usage: "describe node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-pool", Usage: "The ID of the node pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("node-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.GetNodePoolRequest{
								Name: name,
							}

							resp, err := client.GetNodePool(ctx, req)
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
						Usage: "create node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-pool-id", Usage: "A client-specified unique identifier for the node pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.CreateNodePoolRequest{
								Parent:     parent,
								NodePoolId: cmd.String("node-pool-id"),
								RequestId:  cmd.String("request-id"),
							}

							op, err := client.CreateNodePool(ctx, req)
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
						Usage: "update node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-pool", Usage: "The ID of the node pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "node_pool.name" not yet supported.
							node_pool_name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("node-pool"))
							fmt.Printf("Executing update on %s\n", node_pool_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-pool", Usage: "The ID of the node pool.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("node-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteNodePool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.DeleteNodePoolRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteNodePool(ctx, req)
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
				Name:  "server-config",
				Usage: "Manage server-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe server-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.GetServerConfigRequest{
								Name: name,
							}

							resp, err := client.GetServerConfig(ctx, req)
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
				Name:  "vpn-connections",
				Usage: "Manage vpn-connections resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list vpn-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only resources matching this filter will be listed.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the order in which resources will be listed.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of resources to list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from previous list request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.ListVpnConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListVpnConnections(ctx, req)
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
						Usage: "describe vpn-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "vpn-connection", Usage: "The ID of the vpn connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpnConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpn-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.GetVpnConnectionRequest{
								Name: name,
							}

							resp, err := client.GetVpnConnection(ctx, req)
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
						Usage: "create vpn-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
							&cli.StringFlag{Name: "vpn-connection-id", Usage: "The VPN connection identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.CreateVpnConnectionRequest{
								Parent:          parent,
								VpnConnectionId: cmd.String("vpn-connection-id"),
								RequestId:       cmd.String("request-id"),
							}

							op, err := client.CreateVpnConnection(ctx, req)
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
						Usage: "delete vpn-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
							&cli.StringFlag{Name: "vpn-connection", Usage: "The ID of the vpn connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/vpnConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("vpn-connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteVpnConnection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := edgecontainer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &edgecontainerpb.DeleteVpnConnectionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteVpnConnection(ctx, req)
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
		},
	}
}
