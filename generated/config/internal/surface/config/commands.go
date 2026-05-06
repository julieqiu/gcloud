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

package config

import (
	config "cloud.google.com/go/config/apiv1"
	"cloud.google.com/go/config/apiv1/configpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the config command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "manage Infrastructure Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "auto-migration-config",
				Usage: "Manage auto-migration-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe auto-migration-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autoMigrationConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetAutoMigrationConfigRequest{
								Name: name,
							}

							resp, err := client.GetAutoMigrationConfig(ctx, req)
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
						Usage: "update auto-migration-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "auto_migration_config.name" not yet supported.
							auto_migration_config_name := fmt.Sprintf("projects/%s/locations/%s/autoMigrationConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", auto_migration_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "deployment-groups",
				Usage: "Manage deployment-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetDeploymentGroupRequest{
								Name: name,
							}

							resp, err := client.GetDeploymentGroup(ctx, req)
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
						Usage: "create deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group-id", Usage: "The deployment group ID.", Required: true},
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
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.CreateDeploymentGroupRequest{
								Parent:            parent,
								DeploymentGroupId: cmd.String("deployment-group-id"),
								RequestId:         cmd.String("request-id"),
							}

							op, err := client.CreateDeploymentGroup(ctx, req)
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
						Usage: "update deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deployment_group.name" not yet supported.
							deployment_group_name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"))
							fmt.Printf("Executing update on %s\n", deployment_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.StringFlag{Name: "deployment-reference-policy", Usage: "Policy on how to handle referenced deployments when deleting the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any revisions for this deployment group will also.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.DeleteDeploymentGroupRequest{
								Name:                      name,
								RequestId:                 cmd.String("request-id"),
								Force:                     cmd.Bool("force"),
								DeploymentReferencePolicy: configpb.DeleteDeploymentGroupRequest_DeploymentReferencePolicy(configpb.DeleteDeploymentGroupRequest_DeploymentReferencePolicy_value[cmd.String("deployment-reference-policy")]),
							}

							op, err := client.DeleteDeploymentGroup(ctx, req)
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
						Name:  "list",
						Usage: "list deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the DeploymentGroups that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resources, 'page_size' specifies number.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListDeploymentGroups' which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListDeploymentGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeploymentGroups(ctx, req)
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
						Name:  "provision",
						Usage: "provision deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ProvisionDeploymentGroupRequest{
								Name: name,
							}

							op, err := client.ProvisionDeploymentGroup(ctx, req)
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
						Name:  "deprovision",
						Usage: "deprovision deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delete-policy", Usage: "Policy on how resources within each deployment should be handled.", Required: false},
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, this option is propagated to the deletion of each.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.DeprovisionDeploymentGroupRequest{
								Name:         name,
								Force:        cmd.Bool("force"),
								DeletePolicy: configpb.DeleteDeploymentRequest_DeletePolicy(configpb.DeleteDeploymentRequest_DeletePolicy_value[cmd.String("delete-policy")]),
							}

							op, err := client.DeprovisionDeploymentGroup(ctx, req)
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
				Name:  "deployments",
				Usage: "Manage deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the Deployments that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resources, 'page_size' specifies number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListDeployments' which specifies the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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
						Name:  "describe",
						Usage: "describe deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetDeploymentRequest{
								Name: name,
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
						Name:  "create",
						Usage: "create deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-id", Usage: "The Deployment ID.", Required: true},
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
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.CreateDeploymentRequest{
								Parent:       parent,
								DeploymentId: cmd.String("deployment-id"),
								RequestId:    cmd.String("request-id"),
							}

							op, err := client.CreateDeployment(ctx, req)
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
						Usage: "update deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deployment.name" not yet supported.
							deployment_name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							fmt.Printf("Executing update on %s\n", deployment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delete-policy", Usage: "Policy on how resources actuated by the deployment should be.", Required: false},
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any revisions for this deployment will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.DeleteDeploymentRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								Force:        cmd.Bool("force"),
								DeletePolicy: configpb.DeleteDeploymentRequest_DeletePolicy(configpb.DeleteDeploymentRequest_DeletePolicy_value[cmd.String("delete-policy")]),
							}

							op, err := client.DeleteDeployment(ctx, req)
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
						Name:  "export-state",
						Usage: "export-state deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.BoolFlag{Name: "draft", Usage: "If this flag is set to true, the exported deployment state file.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ExportDeploymentStatefileRequest{
								Parent: parent,
								Draft:  cmd.Bool("draft"),
							}

							resp, err := client.ExportDeploymentStatefile(ctx, req)
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
						Name:  "import-state",
						Usage: "import-state deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "lock-id", Usage: "Lock ID of the lock file to verify that the user who is importing.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-draft", Usage: "Optional.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ImportStatefileRequest{
								Parent:    parent,
								LockId:    cmd.Int("lock-id"),
								SkipDraft: cmd.Bool("skip-draft"),
							}

							resp, err := client.ImportStatefile(ctx, req)
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
						Usage: "delete deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "lock-id", Usage: "Lock ID of the lock file to verify that the user who is deleting.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteStatefile on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.DeleteStatefileRequest{
								Name:   name,
								LockId: cmd.Int("lock-id"),
							}

							if err := client.DeleteStatefile(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "lock",
						Usage: "lock deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.LockDeploymentRequest{
								Name: name,
							}

							op, err := client.LockDeployment(ctx, req)
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
						Name:  "unlock",
						Usage: "unlock deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "lock-id", Usage: "Lock ID of the lock file to be unlocked.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.UnlockDeploymentRequest{
								Name:   name,
								LockId: cmd.Int("lock-id"),
							}

							op, err := client.UnlockDeployment(ctx, req)
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
						Name:  "export-lock",
						Usage: "export-lock deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ExportLockInfoRequest{
								Name: name,
							}

							resp, err := client.ExportLockInfo(ctx, req)
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
						Usage: "set-iam-policy deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
				Name:  "previews",
				Usage: "Manage previews resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create previews",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "preview-id", Usage: "The preview ID.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.CreatePreviewRequest{
								Parent:    parent,
								PreviewId: cmd.String("preview-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreatePreview(ctx, req)
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
						Usage: "describe previews",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetPreviewRequest{
								Name: name,
							}

							resp, err := client.GetPreview(ctx, req)
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
						Usage: "list previews",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the Deployments that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resources, 'page_size' specifies number.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListDeployments' which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListPreviewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListPreviews(ctx, req)
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
						Usage: "delete previews",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.DeletePreviewRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeletePreview(ctx, req)
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
						Name:  "export",
						Usage: "export previews",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ExportPreviewResultRequest{
								Parent: parent,
							}

							resp, err := client.ExportPreviewResult(ctx, req)
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
				Name:  "resource-changes",
				Usage: "Manage resource-changes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list resource-changes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the resource changes that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resource changes, 'page_size' specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListResourceChanges' which.", Required: false},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListResourceChangesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListResourceChanges(ctx, req)
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
						Usage: "describe resource-changes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-change", Usage: "The ID of the resource change.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/previews/%s/resourceChanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"), cmd.String("resource-change"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetResourceChangeRequest{
								Name: name,
							}

							resp, err := client.GetResourceChange(ctx, req)
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
				Name:  "resource-drifts",
				Usage: "Manage resource-drifts resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list resource-drifts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the resource drifts that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resource drifts, 'page_size' specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListResourceDrifts' which.", Required: false},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListResourceDriftsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListResourceDrifts(ctx, req)
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
						Usage: "describe resource-drifts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "preview", Usage: "The ID of the preview.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-drift", Usage: "The ID of the resource drift.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/previews/%s/resourceDrifts/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"), cmd.String("resource-drift"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetResourceDriftRequest{
								Name: name,
							}

							resp, err := client.GetResourceDrift(ctx, req)
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
				Name:  "resources",
				Usage: "Manage resources resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s/resources/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"), cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetResourceRequest{
								Name: name,
							}

							resp, err := client.GetResource(ctx, req)
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
						Usage: "list resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the Resources that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resources, 'page_size' specifies number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListResources' which specifies the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListResourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListResources(ctx, req)
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
				},
			},
			{
				Name:  "revisions",
				Usage: "Manage revisions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Lists the Revisions that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resources, `page_size` specifies number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListRevisions' which specifies the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListRevisionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRevisions(ctx, req)
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
						Usage: "describe revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetRevisionRequest{
								Name: name,
							}

							resp, err := client.GetRevision(ctx, req)
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
						Name:  "export-state",
						Usage: "export-state revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ExportRevisionStatefileRequest{
								Parent: parent,
							}

							resp, err := client.ExportRevisionStatefile(ctx, req)
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
						Usage: "describe revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetDeploymentGroupRevisionRequest{
								Name: name,
							}

							resp, err := client.GetDeploymentGroupRevision(ctx, req)
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
						Usage: "list revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-group", Usage: "The ID of the deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of resources, 'page_size' specifies number.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListDeploymentGroupRevisions'.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListDeploymentGroupRevisionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeploymentGroupRevisions(ctx, req)
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
				},
			},
			{
				Name:  "terraform-versions",
				Usage: "Manage terraform-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list terraform-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists the TerraformVersions that match the filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to use to sort the list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a page of terraform versions, 'page_size'.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to 'ListTerraformVersions' which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.ListTerraformVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTerraformVersions(ctx, req)
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
						Usage: "describe terraform-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "terraform-version", Usage: "The ID of the terraform version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/terraformVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("terraform-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := config.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configpb.GetTerraformVersionRequest{
								Name: name,
							}

							resp, err := client.GetTerraformVersion(ctx, req)
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
