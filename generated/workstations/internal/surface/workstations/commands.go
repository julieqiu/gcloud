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

package workstations

import (
	workstations "cloud.google.com/go/workstations/apiv1"
	"cloud.google.com/go/workstations/apiv1/workstationspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the workstations command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "workstations",
		Usage: "manage Cloud Workstations API resources",
		Commands: []*cli.Command{
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
				Name:  "workstation-clusters",
				Usage: "Manage workstation-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe workstation-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.GetWorkstationClusterRequest{
								Name: name,
							}

							resp, err := client.GetWorkstationCluster(ctx, req)
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
						Usage: "list workstation-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "next_page_token value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.ListWorkstationClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkstationClusters(ctx, req)
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
						Name:  "create",
						Usage: "create workstation-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster-id", Usage: "ID to use for the workstation cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.CreateWorkstationClusterRequest{
								Parent:               parent,
								WorkstationClusterId: cmd.String("workstation-cluster-id"),
								ValidateOnly:         cmd.Bool("validate-only"),
							}

							op, err := client.CreateWorkstationCluster(ctx, req)
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
						Usage: "update workstation-clusters",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set, and the workstation cluster is not found, a new.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "workstation_cluster.name" not yet supported.
							workstation_cluster_name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"))
							fmt.Printf("Executing update on %s\n", workstation_cluster_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete workstation-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If set, the request will be rejected if the latest version of the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set, any workstation configurations and workstations in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.DeleteWorkstationClusterRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
								Force:        cmd.Bool("force"),
							}

							op, err := client.DeleteWorkstationCluster(ctx, req)
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
				Name:  "workstation-configs",
				Usage: "Manage workstation-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.GetWorkstationConfigRequest{
								Name: name,
							}

							resp, err := client.GetWorkstationConfig(ctx, req)
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
						Usage: "list workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "next_page_token value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.ListWorkstationConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkstationConfigs(ctx, req)
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
						Name:  "list",
						Usage: "list workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "next_page_token value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.ListUsableWorkstationConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUsableWorkstationConfigs(ctx, req)
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
						Name:  "create",
						Usage: "create workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config-id", Usage: "ID to use for the workstation configuration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.CreateWorkstationConfigRequest{
								Parent:              parent,
								WorkstationConfigId: cmd.String("workstation-config-id"),
								ValidateOnly:        cmd.Bool("validate-only"),
							}

							op, err := client.CreateWorkstationConfig(ctx, req)
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
						Usage: "update workstation-configs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set and the workstation configuration is not found, a new.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "workstation_config.name" not yet supported.
							workstation_config_name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							fmt.Printf("Executing update on %s\n", workstation_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If set, the request is rejected if the latest version of the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set, any workstations in the workstation configuration are.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.DeleteWorkstationConfigRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
								Force:        cmd.Bool("force"),
							}

							op, err := client.DeleteWorkstationConfig(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions workstation-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "workstations",
				Usage: "Manage workstations resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation", Usage: "The ID of the workstation.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"), cmd.String("workstation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.GetWorkstationRequest{
								Name: name,
							}

							resp, err := client.GetWorkstation(ctx, req)
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
						Usage: "list workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "next_page_token value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.ListWorkstationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkstations(ctx, req)
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
						Name:  "list",
						Usage: "list workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "next_page_token value returned from a previous List request, if.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.ListUsableWorkstationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUsableWorkstations(ctx, req)
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
						Name:  "create",
						Usage: "create workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
							&cli.StringFlag{Name: "workstation-id", Usage: "ID to use for the workstation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.CreateWorkstationRequest{
								Parent:        parent,
								WorkstationId: cmd.String("workstation-id"),
								ValidateOnly:  cmd.Bool("validate-only"),
							}

							op, err := client.CreateWorkstation(ctx, req)
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
						Usage: "update workstations",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set and the workstation configuration is not found, a new.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation", Usage: "The ID of the workstation.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "workstation.name" not yet supported.
							workstation_name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"), cmd.String("workstation"))
							fmt.Printf("Executing update on %s\n", workstation_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If set, the request will be rejected if the latest version of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation", Usage: "The ID of the workstation.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"), cmd.String("workstation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.DeleteWorkstationRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteWorkstation(ctx, req)
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
						Name:  "start",
						Usage: "start workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If set, the request will be rejected if the latest version of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation", Usage: "The ID of the workstation.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"), cmd.String("workstation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.StartWorkstationRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.StartWorkstation(ctx, req)
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
						Name:  "stop",
						Usage: "stop workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If set, the request will be rejected if the latest version of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
							&cli.StringFlag{Name: "workstation", Usage: "The ID of the workstation.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"), cmd.String("workstation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := workstations.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &workstationspb.StopWorkstationRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.StopWorkstation(ctx, req)
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
						Name:  "generate-access-token",
						Usage: "generate-access-token workstations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "workstation", Usage: "The ID of the workstation.", Required: true},
							&cli.StringFlag{Name: "workstation-cluster", Usage: "The ID of the workstation cluster.", Required: true},
							&cli.StringFlag{Name: "workstation-config", Usage: "The ID of the workstation config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							workstation := fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s/workstations/%s", cmd.String("project"), cmd.String("location"), cmd.String("workstation-cluster"), cmd.String("workstation-config"), cmd.String("workstation"))
							fmt.Printf("Executing generate-access-token on %s\n", workstation)
							return nil
						},
					},
				},
			},
		},
	}
}
