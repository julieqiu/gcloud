package gkemulticloud

import (
	"context"
	"fmt"
	"strings"

	gkemulticloud "cloud.google.com/go/gkemulticloud/apiv1"
	"cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud gkemulticloud command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "attached-cluster-id", Usage: "The attached cluster id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "platform-version", Usage: "The platform version.", Required: true},
							&cli.StringFlag{Name: "distribution", Usage: "The distribution.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.CreateAttachedClusterRequest{Parent: parent}
							req.AttachedClusterId = cmd.String("attached-cluster-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.AttachedCluster = &gkemulticloudpb.AttachedCluster{
								Name:            cmd.String("name"),
								Description:     cmd.String("description"),
								PlatformVersion: cmd.String("platform-version"),
								Distribution:    cmd.String("distribution"),
								Etag:            cmd.String("etag"),
							}
							op, err := client.CreateAttachedCluster(ctx, req)
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
						Usage: "update attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "attached_cluster", Usage: "The attached_cluster.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "platform-version", Usage: "The platform version.", Required: false},
							&cli.StringFlag{Name: "distribution", Usage: "The distribution.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached_cluster"))
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.UpdateAttachedClusterRequest{}
							req.AttachedCluster = &gkemulticloudpb.AttachedCluster{
								Name:            name,
								Name:            cmd.String("name"),
								Description:     cmd.String("description"),
								PlatformVersion: cmd.String("platform-version"),
								Distribution:    cmd.String("distribution"),
								Etag:            cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("platform-version") {
								paths = append(paths, "platform_version")
							}
							if cmd.IsSet("distribution") {
								paths = append(paths, "distribution")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateAttachedCluster(ctx, req)
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
						Name:  "import",
						Usage: "import attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing import on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "attached_cluster", Usage: "The attached_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached_cluster"))
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAttachedClusterRequest{Name: name}
							resp, err := client.GetAttachedCluster(ctx, req)
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
						Usage: "list attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkemulticloudpb.ListAttachedClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAttachedClusters(ctx, req)
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
						Name:  "delete",
						Usage: "delete attached-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "attached_cluster", Usage: "The attached_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached_cluster"))
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.DeleteAttachedClusterRequest{Name: name}
							op, err := client.DeleteAttachedCluster(ctx, req)
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
						Name:  "generate-attached-cluster-agent-token",
						Usage: "generate-attached-cluster-agent-token attached-clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-attached-cluster-agent-token...")
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedServerConfig", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAttachedServerConfigRequest{Name: name}
							resp, err := client.GetAttachedServerConfig(ctx, req)
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
				Name:  "aws-clusters",
				Usage: "Manage aws-clusters resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws-cluster-id", Usage: "The aws cluster id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "aws-region", Usage: "The aws region.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.CreateAwsClusterRequest{Parent: parent}
							req.AwsClusterId = cmd.String("aws-cluster-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.AwsCluster = &gkemulticloudpb.AwsCluster{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								AwsRegion:   cmd.String("aws-region"),
								Etag:        cmd.String("etag"),
							}
							op, err := client.CreateAwsCluster(ctx, req)
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
						Usage: "update aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "aws-region", Usage: "The aws region.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.UpdateAwsClusterRequest{}
							req.AwsCluster = &gkemulticloudpb.AwsCluster{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								AwsRegion:   cmd.String("aws-region"),
								Etag:        cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("aws-region") {
								paths = append(paths, "aws_region")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateAwsCluster(ctx, req)
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
						Usage: "describe aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAwsClusterRequest{Name: name}
							resp, err := client.GetAwsCluster(ctx, req)
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
						Usage: "list aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkemulticloudpb.ListAwsClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAwsClusters(ctx, req)
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
						Name:  "delete",
						Usage: "delete aws-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.DeleteAwsClusterRequest{Name: name}
							op, err := client.DeleteAwsCluster(ctx, req)
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
						Name:  "generate-aws-cluster-agent-token",
						Usage: "generate-aws-cluster-agent-token aws-clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-aws-cluster-agent-token...")
							return nil
						},
					},
					{
						Name:  "generate-aws-access-token",
						Usage: "generate-aws-access-token aws-clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-aws-access-token...")
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.StringFlag{Name: "aws-node-pool-id", Usage: "The aws node pool id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
							&cli.StringFlag{Name: "subnet-id", Usage: "The subnet id.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.CreateAwsNodePoolRequest{Parent: parent}
							req.AwsNodePoolId = cmd.String("aws-node-pool-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.AwsNodePool = &gkemulticloudpb.AwsNodePool{
								Name:     cmd.String("name"),
								Version:  cmd.String("version"),
								SubnetId: cmd.String("subnet-id"),
								Etag:     cmd.String("etag"),
							}
							op, err := client.CreateAwsNodePool(ctx, req)
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
						Usage: "update aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.StringFlag{Name: "aws_node_pool", Usage: "The aws_node_pool.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
							&cli.StringFlag{Name: "subnet-id", Usage: "The subnet id.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"), cmd.String("aws_node_pool"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.UpdateAwsNodePoolRequest{}
							req.AwsNodePool = &gkemulticloudpb.AwsNodePool{
								Name:     name,
								Name:     cmd.String("name"),
								Version:  cmd.String("version"),
								SubnetId: cmd.String("subnet-id"),
								Etag:     cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("version") {
								paths = append(paths, "version")
							}
							if cmd.IsSet("subnet-id") {
								paths = append(paths, "subnet_id")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateAwsNodePool(ctx, req)
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
						Name:  "rollback",
						Usage: "rollback aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.StringFlag{Name: "aws_node_pool", Usage: "The aws_node_pool.", Required: true},
							&cli.BoolFlag{Name: "respect-pdb", Usage: "The respect pdb.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"), cmd.String("aws_node_pool"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.RollbackAwsNodePoolUpdateRequest{Name: name}
							req.RespectPdb = cmd.Bool("respect-pdb")
							op, err := client.RollbackAwsNodePoolUpdate(ctx, req)
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
						Usage: "describe aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.StringFlag{Name: "aws_node_pool", Usage: "The aws_node_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"), cmd.String("aws_node_pool"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAwsNodePoolRequest{Name: name}
							resp, err := client.GetAwsNodePool(ctx, req)
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
						Usage: "list aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkemulticloudpb.ListAwsNodePoolsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAwsNodePools(ctx, req)
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
						Name:  "delete",
						Usage: "delete aws-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "aws_cluster", Usage: "The aws_cluster.", Required: true},
							&cli.StringFlag{Name: "aws_node_pool", Usage: "The aws_node_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsClusters/%s/awsNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("aws_cluster"), cmd.String("aws_node_pool"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.DeleteAwsNodePoolRequest{Name: name}
							op, err := client.DeleteAwsNodePool(ctx, req)
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
				Name:  "aws-server-config",
				Usage: "Manage aws-server-config resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe aws-server-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/awsServerConfig", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAwsServerConfigRequest{Name: name}
							resp, err := client.GetAwsServerConfig(ctx, req)
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
				Name:  "azure-clients",
				Usage: "Manage azure-clients resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure-client-id", Usage: "The azure client id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "tenant-id", Usage: "The tenant id.", Required: true},
							&cli.StringFlag{Name: "application-id", Usage: "The application id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.CreateAzureClientRequest{Parent: parent}
							req.AzureClientId = cmd.String("azure-client-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.AzureClient = &gkemulticloudpb.AzureClient{
								Name:          cmd.String("name"),
								TenantId:      cmd.String("tenant-id"),
								ApplicationId: cmd.String("application-id"),
							}
							op, err := client.CreateAzureClient(ctx, req)
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
						Usage: "describe azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_client", Usage: "The azure_client.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClients/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_client"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAzureClientRequest{Name: name}
							resp, err := client.GetAzureClient(ctx, req)
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
						Usage: "list azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkemulticloudpb.ListAzureClientsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAzureClients(ctx, req)
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
						Name:  "delete",
						Usage: "delete azure-clients",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_client", Usage: "The azure_client.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClients/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_client"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.DeleteAzureClientRequest{Name: name}
							op, err := client.DeleteAzureClient(ctx, req)
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
				Name:  "azure-clusters",
				Usage: "Manage azure-clusters resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure-cluster-id", Usage: "The azure cluster id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "azure-region", Usage: "The azure region.", Required: true},
							&cli.StringFlag{Name: "resource-group-id", Usage: "The resource group id.", Required: true},
							&cli.StringFlag{Name: "azure-client", Usage: "The azure client.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.CreateAzureClusterRequest{Parent: parent}
							req.AzureClusterId = cmd.String("azure-cluster-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.AzureCluster = &gkemulticloudpb.AzureCluster{
								Name:            cmd.String("name"),
								Description:     cmd.String("description"),
								AzureRegion:     cmd.String("azure-region"),
								ResourceGroupId: cmd.String("resource-group-id"),
								AzureClient:     cmd.String("azure-client"),
								Etag:            cmd.String("etag"),
							}
							op, err := client.CreateAzureCluster(ctx, req)
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
						Usage: "update azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "azure-region", Usage: "The azure region.", Required: false},
							&cli.StringFlag{Name: "resource-group-id", Usage: "The resource group id.", Required: false},
							&cli.StringFlag{Name: "azure-client", Usage: "The azure client.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.UpdateAzureClusterRequest{}
							req.AzureCluster = &gkemulticloudpb.AzureCluster{
								Name:            name,
								Name:            cmd.String("name"),
								Description:     cmd.String("description"),
								AzureRegion:     cmd.String("azure-region"),
								ResourceGroupId: cmd.String("resource-group-id"),
								AzureClient:     cmd.String("azure-client"),
								Etag:            cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("azure-region") {
								paths = append(paths, "azure_region")
							}
							if cmd.IsSet("resource-group-id") {
								paths = append(paths, "resource_group_id")
							}
							if cmd.IsSet("azure-client") {
								paths = append(paths, "azure_client")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateAzureCluster(ctx, req)
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
						Usage: "describe azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAzureClusterRequest{Name: name}
							resp, err := client.GetAzureCluster(ctx, req)
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
						Usage: "list azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkemulticloudpb.ListAzureClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAzureClusters(ctx, req)
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
						Name:  "delete",
						Usage: "delete azure-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.DeleteAzureClusterRequest{Name: name}
							op, err := client.DeleteAzureCluster(ctx, req)
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
						Name:  "generate-azure-cluster-agent-token",
						Usage: "generate-azure-cluster-agent-token azure-clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-azure-cluster-agent-token...")
							return nil
						},
					},
					{
						Name:  "generate-azure-access-token",
						Usage: "generate-azure-access-token azure-clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-azure-access-token...")
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
							&cli.StringFlag{Name: "azure-node-pool-id", Usage: "The azure node pool id.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
							&cli.StringFlag{Name: "subnet-id", Usage: "The subnet id.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "azure-availability-zone", Usage: "The azure availability zone.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.CreateAzureNodePoolRequest{Parent: parent}
							req.AzureNodePoolId = cmd.String("azure-node-pool-id")
							req.ValidateOnly = cmd.Bool("validate-only")
							req.AzureNodePool = &gkemulticloudpb.AzureNodePool{
								Name:                  cmd.String("name"),
								Version:               cmd.String("version"),
								SubnetId:              cmd.String("subnet-id"),
								Etag:                  cmd.String("etag"),
								AzureAvailabilityZone: cmd.String("azure-availability-zone"),
							}
							op, err := client.CreateAzureNodePool(ctx, req)
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
						Usage: "update azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
							&cli.StringFlag{Name: "azure_node_pool", Usage: "The azure_node_pool.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
							&cli.StringFlag{Name: "subnet-id", Usage: "The subnet id.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "azure-availability-zone", Usage: "The azure availability zone.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"), cmd.String("azure_node_pool"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.UpdateAzureNodePoolRequest{}
							req.AzureNodePool = &gkemulticloudpb.AzureNodePool{
								Name:                  name,
								Name:                  cmd.String("name"),
								Version:               cmd.String("version"),
								SubnetId:              cmd.String("subnet-id"),
								Etag:                  cmd.String("etag"),
								AzureAvailabilityZone: cmd.String("azure-availability-zone"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("version") {
								paths = append(paths, "version")
							}
							if cmd.IsSet("subnet-id") {
								paths = append(paths, "subnet_id")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							if cmd.IsSet("azure-availability-zone") {
								paths = append(paths, "azure_availability_zone")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateAzureNodePool(ctx, req)
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
						Usage: "describe azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
							&cli.StringFlag{Name: "azure_node_pool", Usage: "The azure_node_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"), cmd.String("azure_node_pool"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAzureNodePoolRequest{Name: name}
							resp, err := client.GetAzureNodePool(ctx, req)
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
						Usage: "list azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkemulticloudpb.ListAzureNodePoolsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAzureNodePools(ctx, req)
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
						Name:  "delete",
						Usage: "delete azure-node-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "azure_cluster", Usage: "The azure_cluster.", Required: true},
							&cli.StringFlag{Name: "azure_node_pool", Usage: "The azure_node_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureClusters/%s/azureNodePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("azure_cluster"), cmd.String("azure_node_pool"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.DeleteAzureNodePoolRequest{Name: name}
							op, err := client.DeleteAzureNodePool(ctx, req)
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
				Name:  "azure-server-config",
				Usage: "Manage azure-server-config resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe azure-server-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/azureServerConfig", cmd.String("project"), cmd.String("location"))
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkemulticloudpb.GetAzureServerConfigRequest{Name: name}
							resp, err := client.GetAzureServerConfig(ctx, req)
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
				Name:  "jwks",
				Usage: "Manage jwks resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe jwks",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe jwks",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "attached_cluster", Usage: "The attached_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/attachedClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("attached_cluster"))
							fmt.Printf("Executing generate-attached-cluster-install-manifest on %s\n", name)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe openid-configuration",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
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
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
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
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
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
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
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
							client, err := gkemulticloud.NewAttachedClustersClient(ctx)
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
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
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
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
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
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
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
							client, err := gkemulticloud.NewAwsClustersClient(ctx)
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
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
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
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
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
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
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
							client, err := gkemulticloud.NewAzureClustersClient(ctx)
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
		},
	}
}
