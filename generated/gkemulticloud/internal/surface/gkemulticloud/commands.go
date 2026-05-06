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

package gkemulticloud

import (
	gkemulticloud "cloud.google.com/go/gkemulticloud/apiv1"
	"cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the gkemulticloud command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkemulticloud",
		Usage: "manage GKE Multi-Cloud API resources",
		Commands: []*cli.Command{
			{
				Name:  "attached-clusters",
				Usage: "Manage attached-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attached-cluster-id", Usage: "A client provided ID the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually create the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.CreateAttachedClusterRequest{
								Parent:            parent,
								AttachedClusterId: cmd.String("attached-cluster-id"),
								ValidateOnly:      cmd.Bool("validate-only"),
							}

							op, err := client.CreateAttachedCluster(ctx, req)
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
						Usage: "update attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attached-cluster", Usage: "The ID of the attached cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually update the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "attached_cluster.name" not yet supported.
							attached_cluster_name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached-cluster"))
							fmt.Printf("Executing update on %s\n", attached_cluster_name)
							return nil
						},
					},

					{
						Name:  "import",
						Usage: "import attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "distribution", Usage: "The Kubernetes distribution of the underlying attached cluster.", Required: true},
							&cli.StringFlag{Name: "fleet-membership", Usage: "The name of the fleet membership resource to import.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "platform-version", Usage: "The platform version for the cluster (e.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually import the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ImportAttachedClusterRequest{
								Parent:          parent,
								ValidateOnly:    cmd.Bool("validate-only"),
								FleetMembership: cmd.String("fleet-membership"),
								PlatformVersion: cmd.String("platform-version"),
								Distribution:    cmd.String("distribution"),
							}

							op, err := client.ImportAttachedCluster(ctx, req)
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
						Name:  "describe",
						Usage: "describe attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attached-cluster", Usage: "The ID of the attached cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAttachedClusterRequest{
								Name: name,
							}

							resp, err := client.GetAttachedCluster(ctx, req)
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
						Usage: "list attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ListAttachedClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAttachedClusters(ctx, req)
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
						Name:  "delete",
						Usage: "delete attached-clusters",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the.", Required: false},
							&cli.StringFlag{Name: "attached-cluster", Usage: "The ID of the attached cluster.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the.", Required: false},
							&cli.BoolFlag{Name: "ignore-errors", Usage: "If set to true, the deletion of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually delete the resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached-cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAttachedCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.DeleteAttachedClusterRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								IgnoreErrors: cmd.Bool("ignore-errors"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteAttachedCluster(ctx, req)
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
						Name:  "generate-attached-cluster-agent-token",
						Usage: "generate-attached-cluster-agent-token attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attached-cluster", Usage: "The ID of the attached cluster.", Required: true},
							&cli.StringFlag{Name: "audience", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "grant-type", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "options", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "requested-token-type", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "subject-token", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "subject-token-type", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "Required.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							attached_cluster := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached-cluster"))
							fmt.Printf("Executing generate-attached-cluster-agent-token on %s\n", attached_cluster)
							return nil
						},
					},
				},
			},
			{
				Name:  "attached-server-config",
				Usage: "Manage attached-server-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe attached-server-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedServerConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAttachedServerConfigRequest{
								Name: name,
							}

							resp, err := client.GetAttachedServerConfig(ctx, req)
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
				Name:  "aws-clusters",
				Usage: "Manage aws-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster-id", Usage: "A client provided ID the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually create the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.CreateAwsClusterRequest{
								Parent:       parent,
								AwsClusterId: cmd.String("aws-cluster-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateAwsCluster(ctx, req)
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
						Usage: "update aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually update the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "aws_cluster.name" not yet supported.
							aws_cluster_name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							fmt.Printf("Executing update on %s\n", aws_cluster_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAwsClusterRequest{
								Name: name,
							}

							resp, err := client.GetAwsCluster(ctx, req)
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
						Usage: "list aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ListAwsClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAwsClusters(ctx, req)
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
						Name:  "delete",
						Usage: "delete aws-clusters",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the.", Required: false},
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the.", Required: false},
							&cli.BoolFlag{Name: "ignore-errors", Usage: "If set to true, the deletion of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually delete the resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAwsCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.DeleteAwsClusterRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								IgnoreErrors: cmd.Bool("ignore-errors"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteAwsCluster(ctx, req)
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
						Name:  "generate-aws-cluster-agent-token",
						Usage: "generate-aws-cluster-agent-token aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "audience", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "grant-type", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-pool-id", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "options", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "requested-token-type", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "subject-token", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "subject-token-type", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "Required.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							aws_cluster := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							fmt.Printf("Executing generate-aws-cluster-agent-token on %s\n", aws_cluster)
							return nil
						},
					},

					{
						Name:  "generate-aws-access-token",
						Usage: "generate-aws-access-token aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							aws_cluster := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							fmt.Printf("Executing generate-aws-access-token on %s\n", aws_cluster)
							return nil
						},
					},
				},
			},
			{
				Name:  "aws-node-pools",
				Usage: "Manage aws-node-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "aws-node-pool-id", Usage: "A client provided ID the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually create the node.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.CreateAwsNodePoolRequest{
								Parent:        parent,
								AwsNodePoolId: cmd.String("aws-node-pool-id"),
								ValidateOnly:  cmd.Bool("validate-only"),
							}

							op, err := client.CreateAwsNodePool(ctx, req)
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
						Usage: "update aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "aws-node-pool", Usage: "The ID of the aws node pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but don't actually update the node pool.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "aws_node_pool.name" not yet supported.
							aws_node_pool_name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"), cmd.String("aws-node-pool"))
							fmt.Printf("Executing update on %s\n", aws_node_pool_name)
							return nil
						},
					},

					{
						Name:  "rollback",
						Usage: "rollback aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "aws-node-pool", Usage: "The ID of the aws node pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "respect-pdb", Usage: "Option for rollback to ignore the PodDisruptionBudget when.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"), cmd.String("aws-node-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.RollbackAwsNodePoolUpdateRequest{
								Name:       name,
								RespectPdb: cmd.Bool("respect-pdb"),
							}

							op, err := client.RollbackAwsNodePoolUpdate(ctx, req)
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
						Name:  "describe",
						Usage: "describe aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "aws-node-pool", Usage: "The ID of the aws node pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"), cmd.String("aws-node-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAwsNodePoolRequest{
								Name: name,
							}

							resp, err := client.GetAwsNodePool(ctx, req)
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
						Usage: "list aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ListAwsNodePoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAwsNodePools(ctx, req)
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
						Name:  "delete",
						Usage: "delete aws-node-pools",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the.", Required: false},
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "aws-node-pool", Usage: "The ID of the aws node pool.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current ETag of the.", Required: false},
							&cli.BoolFlag{Name: "ignore-errors", Usage: "If set to true, the deletion of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually delete the node.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"), cmd.String("aws-node-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAwsNodePool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.DeleteAwsNodePoolRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								IgnoreErrors: cmd.Bool("ignore-errors"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteAwsNodePool(ctx, req)
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
				Name:  "aws-server-config",
				Usage: "Manage aws-server-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe aws-server-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsServerConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAwsServerConfigRequest{
								Name: name,
							}

							resp, err := client.GetAwsServerConfig(ctx, req)
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
				Name:  "azure-clients",
				Usage: "Manage azure-clients resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-client-id", Usage: "A client provided ID the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually create the client.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.CreateAzureClientRequest{
								Parent:        parent,
								AzureClientId: cmd.String("azure-client-id"),
								ValidateOnly:  cmd.Bool("validate-only"),
							}

							op, err := client.CreateAzureClient(ctx, req)
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
						Name:  "describe",
						Usage: "describe azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-client", Usage: "The ID of the azure client.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClients/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-client"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAzureClientRequest{
								Name: name,
							}

							resp, err := client.GetAzureClient(ctx, req)
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
						Usage: "list azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ListAzureClientsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAzureClients(ctx, req)
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
						Name:  "delete",
						Usage: "delete azure-clients",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the.", Required: false},
							&cli.StringFlag{Name: "azure-client", Usage: "The ID of the azure client.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually delete the resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClients/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-client"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAzureClient %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.DeleteAzureClientRequest{
								Name:         name,
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteAzureClient(ctx, req)
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
				Name:  "azure-clusters",
				Usage: "Manage azure-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster-id", Usage: "A client provided ID the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually create the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.CreateAzureClusterRequest{
								Parent:         parent,
								AzureClusterId: cmd.String("azure-cluster-id"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.CreateAzureCluster(ctx, req)
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
						Usage: "update azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually update the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "azure_cluster.name" not yet supported.
							azure_cluster_name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							fmt.Printf("Executing update on %s\n", azure_cluster_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAzureClusterRequest{
								Name: name,
							}

							resp, err := client.GetAzureCluster(ctx, req)
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
						Usage: "list azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ListAzureClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAzureClusters(ctx, req)
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
						Name:  "delete",
						Usage: "delete azure-clusters",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the.", Required: false},
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the.", Required: false},
							&cli.BoolFlag{Name: "ignore-errors", Usage: "If set to true, the deletion of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually delete the resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAzureCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.DeleteAzureClusterRequest{
								Name:         name,
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
								IgnoreErrors: cmd.Bool("ignore-errors"),
							}

							op, err := client.DeleteAzureCluster(ctx, req)
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
						Name:  "generate-azure-cluster-agent-token",
						Usage: "generate-azure-cluster-agent-token azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "audience", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "grant-type", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "node-pool-id", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "options", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "requested-token-type", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "Optional.", Required: false},
							&cli.StringFlag{Name: "subject-token", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "subject-token-type", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "Required.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							azure_cluster := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							fmt.Printf("Executing generate-azure-cluster-agent-token on %s\n", azure_cluster)
							return nil
						},
					},

					{
						Name:  "generate-azure-access-token",
						Usage: "generate-azure-access-token azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							azure_cluster := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							fmt.Printf("Executing generate-azure-access-token on %s\n", azure_cluster)
							return nil
						},
					},
				},
			},
			{
				Name:  "azure-node-pools",
				Usage: "Manage azure-node-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "azure-node-pool-id", Usage: "A client provided ID the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually create the node.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.CreateAzureNodePoolRequest{
								Parent:          parent,
								AzureNodePoolId: cmd.String("azure-node-pool-id"),
								ValidateOnly:    cmd.Bool("validate-only"),
							}

							op, err := client.CreateAzureNodePool(ctx, req)
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
						Usage: "update azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "azure-node-pool", Usage: "The ID of the azure node pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but don't actually update the node pool.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "azure_node_pool.name" not yet supported.
							azure_node_pool_name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"), cmd.String("azure-node-pool"))
							fmt.Printf("Executing update on %s\n", azure_node_pool_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "azure-node-pool", Usage: "The ID of the azure node pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"), cmd.String("azure-node-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAzureNodePoolRequest{
								Name: name,
							}

							resp, err := client.GetAzureNodePool(ctx, req)
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
						Usage: "list azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.ListAzureNodePoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAzureNodePools(ctx, req)
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
						Name:  "delete",
						Usage: "delete azure-node-pools",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the.", Required: false},
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "azure-node-pool", Usage: "The ID of the azure node pool.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current ETag of the.", Required: false},
							&cli.BoolFlag{Name: "ignore-errors", Usage: "If set to true, the deletion of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not actually delete the node.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"), cmd.String("azure-node-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAzureNodePool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.DeleteAzureNodePoolRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								Etag:         cmd.String("etag"),
								IgnoreErrors: cmd.Bool("ignore-errors"),
							}

							op, err := client.DeleteAzureNodePool(ctx, req)
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
				Name:  "azure-server-config",
				Usage: "Manage azure-server-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe azure-server-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureServerConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GetAzureServerConfigRequest{
								Name: name,
							}

							resp, err := client.GetAzureServerConfig(ctx, req)
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
				Name:  "jwks",
				Usage: "Manage jwks resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe jwks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							aws_cluster := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							fmt.Printf("Executing describe on %s\n", aws_cluster)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe jwks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							azure_cluster := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							fmt.Printf("Executing describe on %s\n", azure_cluster)
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
						Name:  "generate-attached-cluster-install-manifest",
						Usage: "generate-attached-cluster-install-manifest locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attached-cluster-id", Usage: "A client provided ID of the resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "platform-version", Usage: "The platform version for the cluster (e.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkemulticloud.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkemulticloudpb.GenerateAttachedClusterInstallManifestRequest{
								Parent:            parent,
								AttachedClusterId: cmd.String("attached-cluster-id"),
								PlatformVersion:   cmd.String("platform-version"),
							}

							resp, err := client.GenerateAttachedClusterInstallManifest(ctx, req)
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
				Name:  "openid-configuration",
				Usage: "Manage openid-configuration resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe openid-configuration",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aws-cluster", Usage: "The ID of the aws cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							aws_cluster := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws-cluster"))
							fmt.Printf("Executing describe on %s\n", aws_cluster)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe openid-configuration",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "azure-cluster", Usage: "The ID of the azure cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							azure_cluster := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure-cluster"))
							fmt.Printf("Executing describe on %s\n", azure_cluster)
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
		},
	}
}
