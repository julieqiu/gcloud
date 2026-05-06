package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	dataproc "cloud.google.com/go/dataproc/apiv1"
	"cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "dataproc",
				Usage: "manage Cloud Dataproc API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "autoscaling-policies",
						Usage: "Manage autoscaling-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create autoscaling-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.CreateAutoscalingPolicyRequest{Parent: parent}
									req.Policy = &dataprocpb.AutoscalingPolicy{
										Id: cmd.String("id"),
									}
									resp, err := client.CreateAutoscalingPolicy(ctx, req)
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
								Usage: "update autoscaling-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "autoscaling_policy", Usage: "The autoscaling_policy.", Required: true},
									&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("autoscaling_policy"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.UpdateAutoscalingPolicyRequest{}
									req.Policy = &dataprocpb.AutoscalingPolicy{
										Name: name,
										Id:   cmd.String("id"),
									}
									var paths []string
									if cmd.IsSet("id") {
										paths = append(paths, "id")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateAutoscalingPolicy(ctx, req)
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
								Usage: "describe autoscaling-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "autoscaling_policy", Usage: "The autoscaling_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("autoscaling_policy"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.GetAutoscalingPolicyRequest{Name: name}
									resp, err := client.GetAutoscalingPolicy(ctx, req)
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
								Usage: "list autoscaling-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataprocpb.ListAutoscalingPoliciesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAutoscalingPolicies(ctx, req)
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
								Usage: "delete autoscaling-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "autoscaling_policy", Usage: "The autoscaling_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("autoscaling_policy"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.DeleteAutoscalingPolicyRequest{Name: name}
									if err := client.DeleteAutoscalingPolicy(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "batches",
						Usage: "Manage batches resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create batches",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "batch-id", Usage: "The batch id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewBatchControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.CreateBatchRequest{Parent: parent}
									req.BatchId = cmd.String("batch-id")
									op, err := client.CreateBatch(ctx, req)
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
								Usage: "describe batches",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "batch", Usage: "The batch.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/batches/%s", cmd.String("project"), cmd.String("location"), cmd.String("batch"))
									client, err := dataproc.NewBatchControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.GetBatchRequest{Name: name}
									resp, err := client.GetBatch(ctx, req)
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
								Usage: "list batches",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewBatchControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataprocpb.ListBatchesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListBatches(ctx, req)
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
								Usage: "delete batches",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "batch", Usage: "The batch.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/batches/%s", cmd.String("project"), cmd.String("location"), cmd.String("batch"))
									client, err := dataproc.NewBatchControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.DeleteBatchRequest{Name: name}
									if err := client.DeleteBatch(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "clusters",
						Usage: "Manage clusters resources",
						Commands: []*cli.Command{
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "stop",
								Usage: "stop clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop...")
									return nil
								},
							},
							{
								Name:  "start",
								Usage: "start clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "diagnose",
								Usage: "diagnose clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing diagnose...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "jobs",
						Usage: "Manage jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "submit",
								Usage: "submit jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing submit...")
									return nil
								},
							},
							{
								Name:  "submit-as-operation",
								Usage: "submit-as-operation jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing submit-as-operation...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "node-groups",
						Usage: "Manage node-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create node-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "node-group-id", Usage: "The node group id.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
									client, err := dataproc.NewNodeGroupControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.CreateNodeGroupRequest{Parent: parent}
									req.NodeGroupId = cmd.String("node-group-id")
									req.NodeGroup = &dataprocpb.NodeGroup{
										Name: cmd.String("name"),
									}
									op, err := client.CreateNodeGroup(ctx, req)
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
								Name:  "resize",
								Usage: "resize node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe node-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "node_group", Usage: "The node_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/clusters/%s/nodeGroups/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"), cmd.String("node_group"))
									client, err := dataproc.NewNodeGroupControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.GetNodeGroupRequest{Name: name}
									resp, err := client.GetNodeGroup(ctx, req)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewAutoscalingPolicyClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewBatchControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewBatchControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewBatchControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewBatchControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewClusterControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewClusterControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewClusterControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewClusterControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewJobControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewJobControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewJobControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewJobControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewNodeGroupControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewNodeGroupControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewNodeGroupControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewNodeGroupControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewSessionControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewSessionControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewSessionControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewSessionControllerClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
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
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
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
						Name:  "session-templates",
						Usage: "Manage session-templates resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create session-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.CreateSessionTemplateRequest{Parent: parent}
									req.SessionTemplate = &dataprocpb.SessionTemplate{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreateSessionTemplate(ctx, req)
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
								Usage: "update session-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "template", Usage: "The template.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("template"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.UpdateSessionTemplateRequest{}
									req.SessionTemplate = &dataprocpb.SessionTemplate{
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
									resp, err := client.UpdateSessionTemplate(ctx, req)
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
								Usage: "describe session-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "template", Usage: "The template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("template"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.GetSessionTemplateRequest{Name: name}
									resp, err := client.GetSessionTemplate(ctx, req)
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
								Usage: "list session-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataprocpb.ListSessionTemplatesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSessionTemplates(ctx, req)
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
								Usage: "delete session-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "template", Usage: "The template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("template"))
									client, err := dataproc.NewSessionTemplateControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.DeleteSessionTemplateRequest{Name: name}
									if err := client.DeleteSessionTemplate(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "sessions",
						Usage: "Manage sessions resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "session-id", Usage: "The session id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: false},
									&cli.StringFlag{Name: "session-template", Usage: "The session template.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewSessionControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.CreateSessionRequest{Parent: parent}
									req.SessionId = cmd.String("session-id")
									req.Session = &dataprocpb.Session{
										Name:            cmd.String("name"),
										User:            cmd.String("user"),
										SessionTemplate: cmd.String("session-template"),
									}
									op, err := client.CreateSession(ctx, req)
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
								Usage: "describe sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("session"))
									client, err := dataproc.NewSessionControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.GetSessionRequest{Name: name}
									resp, err := client.GetSession(ctx, req)
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
								Usage: "list sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataproc.NewSessionControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataprocpb.ListSessionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSessions(ctx, req)
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
								Name:  "terminate",
								Usage: "terminate sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("session"))
									client, err := dataproc.NewSessionControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.TerminateSessionRequest{Name: name}
									op, err := client.TerminateSession(ctx, req)
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
								Usage: "delete sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("session"))
									client, err := dataproc.NewSessionControllerClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.DeleteSessionRequest{Name: name}
									op, err := client.DeleteSession(ctx, req)
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
						Name:  "workflow-templates",
						Usage: "Manage workflow-templates resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "id", Usage: "The id.", Required: true},
									&cli.IntFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.CreateWorkflowTemplateRequest{Parent: parent}
									req.Template = &dataprocpb.WorkflowTemplate{
										Id:      cmd.String("id"),
										Version: int32(cmd.Int("version")),
									}
									resp, err := client.CreateWorkflowTemplate(ctx, req)
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
								Usage: "describe workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "workflow_template", Usage: "The workflow_template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("region"), cmd.String("workflow_template"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.GetWorkflowTemplateRequest{Name: name}
									resp, err := client.GetWorkflowTemplate(ctx, req)
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
								Name:  "instantiate",
								Usage: "instantiate workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "workflow_template", Usage: "The workflow_template.", Required: true},
									&cli.IntFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("region"), cmd.String("workflow_template"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.InstantiateWorkflowTemplateRequest{Name: name}
									req.Version = int32(cmd.Int("version"))
									op, err := client.InstantiateWorkflowTemplate(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("InstantiateWorkflowTemplate completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "instantiate-inline",
								Usage: "instantiate-inline workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s", cmd.String("project"), cmd.String("region"))
									fmt.Printf("Executing instantiate-inline on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "workflow_template", Usage: "The workflow_template.", Required: true},
									&cli.StringFlag{Name: "id", Usage: "The id.", Required: false},
									&cli.IntFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("region"), cmd.String("workflow_template"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.UpdateWorkflowTemplateRequest{}
									req.Template = &dataprocpb.WorkflowTemplate{
										Name:    name,
										Id:      cmd.String("id"),
										Version: int32(cmd.Int("version")),
									}
									var paths []string
									if cmd.IsSet("id") {
										paths = append(paths, "id")
									}
									if cmd.IsSet("version") {
										paths = append(paths, "version")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateWorkflowTemplate(ctx, req)
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
								Usage: "list workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/regions/%s", cmd.String("project"), cmd.String("region"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataprocpb.ListWorkflowTemplatesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWorkflowTemplates(ctx, req)
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
								Usage: "delete workflow-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "region", Usage: "The region.", Required: true},
									&cli.StringFlag{Name: "workflow_template", Usage: "The workflow_template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/regions/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("region"), cmd.String("workflow_template"))
									client, err := dataproc.NewWorkflowTemplateClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataprocpb.DeleteWorkflowTemplateRequest{Name: name}
									if err := client.DeleteWorkflowTemplate(ctx, req); err != nil {
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
