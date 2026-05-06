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

package dataproc

import (
	dataproc "cloud.google.com/go/dataproc/apiv1"
	"cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the dataproc command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "dataproc",
		Usage: "manage Cloud Dataproc API resources",
		Commands: []*cli.Command{
			{
				Name:  "autoscaling-policies",
				Usage: "Manage autoscaling-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create autoscaling-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.CreateAutoscalingPolicyRequest{
								Parent: parent,
							}

							resp, err := client.CreateAutoscalingPolicy(ctx, req)
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
						Usage: "update autoscaling-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaling-policie", Usage: "The ID of the autoscaling policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "policy.name" not yet supported.
							policy_name := fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("autoscaling-policie"))
							fmt.Printf("Executing update on %s\n", policy_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe autoscaling-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaling-policie", Usage: "The ID of the autoscaling policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("autoscaling-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.GetAutoscalingPolicyRequest{
								Name: name,
							}

							resp, err := client.GetAutoscalingPolicy(ctx, req)
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
						Usage: "list autoscaling-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in each response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, returned by a previous call, to request the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.ListAutoscalingPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutoscalingPolicies(ctx, req)
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
						Usage: "delete autoscaling-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaling-policie", Usage: "The ID of the autoscaling policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("autoscaling-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAutoscalingPolicy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.DeleteAutoscalingPolicyRequest{
								Name: name,
							}

							if err := client.DeleteAutoscalingPolicy(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "batch-id", Usage: "The ID to use for the batch, which will become the final.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.CreateBatchRequest{
								Parent:    parent,
								BatchId:   cmd.String("batch-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateBatch(ctx, req)
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
						Usage: "describe batches",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "batche", Usage: "The ID of the batche.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/batches/%s", cmd.String("project"), cmd.String("location"), cmd.String("batche"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.GetBatchRequest{
								Name: name,
							}

							resp, err := client.GetBatch(ctx, req)
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
						Usage: "list batches",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter for the batches to return in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field(s) on which to sort the list of batches.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of batches to return in each response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous `ListBatches` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.ListBatchesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListBatches(ctx, req)
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
						Usage: "delete batches",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "batche", Usage: "The ID of the batche.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/batches/%s", cmd.String("project"), cmd.String("location"), cmd.String("batche"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteBatch on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.DeleteBatchRequest{
								Name: name,
							}

							if err := client.DeleteBatch(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "action-on-failed-primary-workers", Usage: "Failure action when primary worker creation fails.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("region"))
							fmt.Printf("Executing create on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster_name", Usage: "The ID of the cluster_name.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("cluster_name"))
							fmt.Printf("Executing update on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "stop",
						Usage: "stop clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-uuid", Usage: "Specifying the `cluster_uuid` means the RPC will fail.", Required: false},
							&cli.StringFlag{Name: "cluster_name", Usage: "The ID of the cluster_name.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("cluster_name"))
							fmt.Printf("Executing stop on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "start",
						Usage: "start clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-uuid", Usage: "Specifying the `cluster_uuid` means the RPC will fail.", Required: false},
							&cli.StringFlag{Name: "cluster_name", Usage: "The ID of the cluster_name.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("cluster_name"))
							fmt.Printf("Executing start on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-uuid", Usage: "Specifying the `cluster_uuid` means the RPC should fail.", Required: false},
							&cli.StringFlag{Name: "cluster_name", Usage: "The ID of the cluster_name.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("cluster_name"))
							fmt.Printf("Executing delete on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster_name", Usage: "The ID of the cluster_name.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("cluster_name"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter constraining the clusters to list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard List page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard List page token.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "diagnose",
						Usage: "diagnose clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster_name", Usage: "The ID of the cluster_name.", Required: true},
							&cli.StringSliceFlag{Name: "jobs", Usage: "Specifies a list of jobs on which diagnosis is to be performed.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "tarball-access", Usage: "(Optional) The access type to the diagnostic tarball.", Required: false},
							&cli.StringFlag{Name: "tarball-gcs-dir", Usage: "(Optional) The output Cloud Storage directory for the diagnostic.", Required: false},
							&cli.StringSliceFlag{Name: "yarn-application-ids", Usage: "Specifies a list of yarn applications on which diagnosis is to be.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("cluster_name"))
							fmt.Printf("Executing diagnose on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing get-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique id used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("region"))
							fmt.Printf("Executing submit on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "submit-as-operation",
						Usage: "submit-as-operation jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique id used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("region"))
							fmt.Printf("Executing submit-as-operation on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("job_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-name", Usage: "If set, the returned jobs list includes only jobs that were.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "A filter constraining the jobs to list.", Required: false},
							&cli.StringFlag{Name: "job-state-matcher", Usage: "Specifies enumerated categories of jobs to list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The number of results to return in each response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, returned by a previous call, to request the.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("job_id"))
							fmt.Printf("Executing update on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("job_id"))
							fmt.Printf("Executing cancel on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("region"), cmd.String("job_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
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
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "node-group-id", Usage: "An optional node group ID.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/regions/%s/clusters/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.CreateNodeGroupRequest{
								Parent:      parent,
								NodeGroupId: cmd.String("node-group-id"),
								RequestId:   cmd.String("request-id"),
							}

							op, err := client.CreateNodeGroup(ctx, req)
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
						Name:  "resize",
						Usage: "resize node-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "node-group", Usage: "The ID of the node group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.IntFlag{Name: "size", Usage: "The number of running instances for the node group to maintain.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/clusters/%s/nodeGroups/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"), cmd.String("node-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.ResizeNodeGroupRequest{
								Name:      name,
								Size:      int32(cmd.Int("size")),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.ResizeNodeGroup(ctx, req)
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
						Usage: "describe node-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "node-group", Usage: "The ID of the node group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/clusters/%s/nodeGroups/%s", cmd.String("project"), cmd.String("region"), cmd.String("cluster"), cmd.String("node-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.GetNodeGroupRequest{
								Name: name,
							}

							resp, err := client.GetNodeGroup(ctx, req)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/regions/%s/operations/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.CreateSessionTemplateRequest{
								Parent: parent,
							}

							resp, err := client.CreateSessionTemplate(ctx, req)
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
						Usage: "update session-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session-template", Usage: "The ID of the session template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "session_template.name" not yet supported.
							session_template_name := fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("session-template"))
							fmt.Printf("Executing update on %s\n", session_template_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe session-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session-template", Usage: "The ID of the session template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("session-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.GetSessionTemplateRequest{
								Name: name,
							}

							resp, err := client.GetSessionTemplate(ctx, req)
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
						Usage: "list session-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter for the session templates to return in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of sessions to return in each response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous `ListSessions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.ListSessionTemplatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListSessionTemplates(ctx, req)
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
						Usage: "delete session-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session-template", Usage: "The ID of the session template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sessionTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("session-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSessionTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.DeleteSessionTemplateRequest{
								Name: name,
							}

							if err := client.DeleteSessionTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.StringFlag{Name: "session-id", Usage: "The ID to use for the session, which becomes the final component.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.CreateSessionRequest{
								Parent:    parent,
								SessionId: cmd.String("session-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateSession(ctx, req)
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
						Usage: "describe sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.GetSessionRequest{
								Name: name,
							}

							resp, err := client.GetSession(ctx, req)
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
						Usage: "list sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter for the sessions to return in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of sessions to return in each response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous `ListSessions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.ListSessionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListSessions(ctx, req)
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
						Name:  "terminate",
						Usage: "terminate sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.TerminateSessionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.TerminateSession(ctx, req)
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
						Usage: "delete sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique ID used to identify the request.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.DeleteSessionRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteSession(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.CreateWorkflowTemplateRequest{
								Parent: parent,
							}

							resp, err := client.CreateWorkflowTemplate(ctx, req)
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
						Usage: "describe workflow-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "version", Usage: "The version of workflow template to retrieve.", Required: false},
							&cli.StringFlag{Name: "workflow-template", Usage: "The ID of the workflow template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.GetWorkflowTemplateRequest{
								Name:    name,
								Version: int32(cmd.Int("version")),
							}

							resp, err := client.GetWorkflowTemplate(ctx, req)
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
						Name:  "instantiate",
						Usage: "instantiate workflow-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A tag that prevents multiple concurrent workflow.", Required: false},
							&cli.IntFlag{Name: "version", Usage: "The version of workflow template to instantiate.", Required: false},
							&cli.StringFlag{Name: "workflow-template", Usage: "The ID of the workflow template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("InstantiateWorkflowTemplate %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.InstantiateWorkflowTemplateRequest{
								Name:      name,
								Version:   int32(cmd.Int("version")),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.InstantiateWorkflowTemplate(ctx, req)
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
						Name:  "instantiate-inline",
						Usage: "instantiate-inline workflow-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A tag that prevents multiple concurrent workflow.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("InstantiateInlineWorkflowTemplate %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.InstantiateInlineWorkflowTemplateRequest{
								Parent:    parent,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.InstantiateInlineWorkflowTemplate(ctx, req)
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
						Name:  "update",
						Usage: "update workflow-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workflow-template", Usage: "The ID of the workflow template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "template.name" not yet supported.
							template_name := fmt.Sprintf("projects/%s/locations/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow-template"))
							fmt.Printf("Executing update on %s\n", template_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list workflow-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in each response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, returned by a previous call, to request the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.ListWorkflowTemplatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkflowTemplates(ctx, req)
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
						Usage: "delete workflow-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "version", Usage: "The version of workflow template to delete.", Required: false},
							&cli.StringFlag{Name: "workflow-template", Usage: "The ID of the workflow template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflowTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteWorkflowTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataproc.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataprocpb.DeleteWorkflowTemplateRequest{
								Name:    name,
								Version: int32(cmd.Int("version")),
							}

							if err := client.DeleteWorkflowTemplate(ctx, req); err != nil {
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
