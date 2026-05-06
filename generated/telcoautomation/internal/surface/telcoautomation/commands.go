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

package telcoautomation

import (
	telcoautomation "cloud.google.com/go/telcoautomation/apiv1"
	"cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the telcoautomation command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "telcoautomation",
		Usage: "manage Telco Automation API resources",
		Commands: []*cli.Command{
			{
				Name:  "blueprints",
				Usage: "Manage blueprints resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint-id", Usage: "The name of the blueprint.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.CreateBlueprintRequest{
								Parent:      parent,
								BlueprintId: cmd.String("blueprint-id"),
							}

							resp, err := client.CreateBlueprint(ctx, req)
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
						Usage: "update blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "blueprint.name" not yet supported.
							blueprint_name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							fmt.Printf("Executing update on %s\n", blueprint_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Defines the type of view of the blueprint.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.GetBlueprintRequest{
								Name: name,
								View: telcoautomationpb.BlueprintView(telcoautomationpb.BlueprintView_value[cmd.String("view")]),
							}

							resp, err := client.GetBlueprint(ctx, req)
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
						Usage: "delete blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteBlueprint on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.DeleteBlueprintRequest{
								Name: name,
							}

							if err := client.DeleteBlueprint(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering only supports equality on blueprint state.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of blueprints to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous ListBlueprints call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListBlueprintsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBlueprints(ctx, req)
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
						Name:  "approve",
						Usage: "approve blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ApproveBlueprintRequest{
								Name: name,
							}

							resp, err := client.ApproveBlueprint(ctx, req)
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
						Name:  "propose",
						Usage: "propose blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ProposeBlueprintRequest{
								Name: name,
							}

							resp, err := client.ProposeBlueprint(ctx, req)
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
						Name:  "reject",
						Usage: "reject blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.RejectBlueprintRequest{
								Name: name,
							}

							resp, err := client.RejectBlueprint(ctx, req)
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
						Usage: "list blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of revisions to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous ListBlueprintRevisions call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListBlueprintRevisionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBlueprintRevisions(ctx, req)
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
						Name:  "search-revisions",
						Usage: "search-revisions blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of blueprints revisions to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous search call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Supported queries:.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.SearchBlueprintRevisionsRequest{
								Parent:    parent,
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchBlueprintRevisions(ctx, req)
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
						Name:  "discard",
						Usage: "discard blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "blueprint", Usage: "The ID of the blueprint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.DiscardBlueprintChangesRequest{
								Name: name,
							}

							resp, err := client.DiscardBlueprintChanges(ctx, req)
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
				Name:  "deployments",
				Usage: "Manage deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "search-revisions",
						Usage: "search-revisions deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of deployment revisions to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous search call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Supported queries:.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.SearchDeploymentRevisionsRequest{
								Parent:    parent,
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchDeploymentRevisions(ctx, req)
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
						Usage: "create deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-id", Usage: "The name of the deployment.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.CreateDeploymentRequest{
								Parent:       parent,
								DeploymentId: cmd.String("deployment-id"),
							}

							resp, err := client.CreateDeployment(ctx, req)
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
						Usage: "update deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deployment.name" not yet supported.
							deployment_name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							fmt.Printf("Executing update on %s\n", deployment_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Defines the type of view of the deployment.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.GetDeploymentRequest{
								Name: name,
								View: telcoautomationpb.DeploymentView(telcoautomationpb.DeploymentView_value[cmd.String("view")]),
							}

							resp, err := client.GetDeployment(ctx, req)
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
						Name:  "remove",
						Usage: "remove deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute RemoveDeployment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.RemoveDeploymentRequest{
								Name: name,
							}

							if err := client.RemoveDeployment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering only supports equality on deployment state.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of deployments to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous ListDeployments call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListDeploymentsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeployments(ctx, req)
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
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of revisions to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous ListDeploymentRevisions.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListDeploymentRevisionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeploymentRevisions(ctx, req)
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
						Name:  "discard",
						Usage: "discard deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.DiscardDeploymentChangesRequest{
								Name: name,
							}

							resp, err := client.DiscardDeploymentChanges(ctx, req)
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
						Name:  "apply",
						Usage: "apply deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ApplyDeploymentRequest{
								Name: name,
							}

							resp, err := client.ApplyDeployment(ctx, req)
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
						Name:  "compute-deployment-status",
						Usage: "compute-deployment-status deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ComputeDeploymentStatusRequest{
								Name: name,
							}

							resp, err := client.ComputeDeploymentStatus(ctx, req)
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
						Name:  "rollback",
						Usage: "rollback deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision-id", Usage: "The revision id of deployment to roll back to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.RollbackDeploymentRequest{
								Name:       name,
								RevisionId: cmd.String("revision-id"),
							}

							resp, err := client.RollbackDeployment(ctx, req)
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
				Name:  "edge-slms",
				Usage: "Manage edge-slms resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListEdgeSlmsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEdgeSlms(ctx, req)
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
						Usage: "describe edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-slm", Usage: "The ID of the edge slm.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/edgeSlms/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-slm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.GetEdgeSlmRequest{
								Name: name,
							}

							resp, err := client.GetEdgeSlm(ctx, req)
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
						Usage: "create edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-slm-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.CreateEdgeSlmRequest{
								Parent:    parent,
								EdgeSlmId: cmd.String("edge-slm-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateEdgeSlm(ctx, req)
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
						Usage: "delete edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "edge-slm", Usage: "The ID of the edge slm.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/edgeSlms/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge-slm"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEdgeSlm %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.DeleteEdgeSlmRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteEdgeSlm(ctx, req)
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
				Name:  "hydrated-deployments",
				Usage: "Manage hydrated-deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "hydrated-deployment", Usage: "The ID of the hydrated deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s/hydratedDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"), cmd.String("hydrated-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.GetHydratedDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetHydratedDeployment(ctx, req)
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
						Usage: "list hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of hydrated deployments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous ListHydratedDeployments.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListHydratedDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListHydratedDeployments(ctx, req)
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
						Name:  "update",
						Usage: "update hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "hydrated-deployment", Usage: "The ID of the hydrated deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "hydrated_deployment.name" not yet supported.
							hydrated_deployment_name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s/hydratedDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"), cmd.String("hydrated-deployment"))
							fmt.Printf("Executing update on %s\n", hydrated_deployment_name)
							return nil
						},
					},

					{
						Name:  "apply",
						Usage: "apply hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "hydrated-deployment", Usage: "The ID of the hydrated deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s/hydratedDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"), cmd.String("deployment"), cmd.String("hydrated-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ApplyHydratedDeploymentRequest{
								Name: name,
							}

							resp, err := client.ApplyHydratedDeployment(ctx, req)
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
							name := fmt.Sprintf("projects/%s/locations", cmd.String("project"))
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
							name := fmt.Sprintf("projects/%s/locations/%s/operations", cmd.String("project"), cmd.String("location"))
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
				Name:  "orchestration-clusters",
				Usage: "Manage orchestration-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListOrchestrationClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListOrchestrationClusters(ctx, req)
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
						Usage: "describe orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.GetOrchestrationClusterRequest{
								Name: name,
							}

							resp, err := client.GetOrchestrationCluster(ctx, req)
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
						Usage: "create orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.CreateOrchestrationClusterRequest{
								Parent:                 parent,
								OrchestrationClusterId: cmd.String("orchestration-cluster-id"),
								RequestId:              cmd.String("request-id"),
							}

							op, err := client.CreateOrchestrationCluster(ctx, req)
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
						Usage: "delete orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The ID of the orchestration cluster.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration-cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteOrchestrationCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.DeleteOrchestrationClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteOrchestrationCluster(ctx, req)
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
				Name:  "public-blueprints",
				Usage: "Manage public-blueprints resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list public-blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.ListPublicBlueprintsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPublicBlueprints(ctx, req)
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
						Usage: "describe public-blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public-blueprint", Usage: "The ID of the public blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/publicBlueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("public-blueprint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := telcoautomation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &telcoautomationpb.GetPublicBlueprintRequest{
								Name: name,
							}

							resp, err := client.GetPublicBlueprint(ctx, req)
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
		},
	}
}
