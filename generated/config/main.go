package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	config "cloud.google.com/go/config/apiv1"
	"cloud.google.com/go/config/apiv1/configpb"
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
				Name:  "config",
				Usage: "manage Infrastructure Manager API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "auto-migration-config",
						Usage: "Manage auto-migration-config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe auto-migration-config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/autoMigrationConfig", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetAutoMigrationConfigRequest{Name: name}
									resp, err := client.GetAutoMigrationConfig(ctx, req)
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
								Usage: "update auto-migration-config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.BoolFlag{Name: "auto-migration-enabled", Usage: "The auto migration enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/autoMigrationConfig", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.UpdateAutoMigrationConfigRequest{}
									req.AutoMigrationConfig = &configpb.AutoMigrationConfig{
										Name:                 name,
										AutoMigrationEnabled: cmd.Bool("auto-migration-enabled"),
									}
									var paths []string
									if cmd.IsSet("auto-migration-enabled") {
										paths = append(paths, "auto_migration_enabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAutoMigrationConfig(ctx, req)
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
						Name:  "deployment-groups",
						Usage: "Manage deployment-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment_group", Usage: "The deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment_group"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetDeploymentGroupRequest{Name: name}
									resp, err := client.GetDeploymentGroup(ctx, req)
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
								Usage: "create deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment-group-id", Usage: "The deployment group id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.CreateDeploymentGroupRequest{Parent: parent}
									req.DeploymentGroupId = cmd.String("deployment-group-id")
									op, err := client.CreateDeploymentGroup(ctx, req)
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
								Usage: "update deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment_group", Usage: "The deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment_group"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.UpdateDeploymentGroupRequest{}
									req.DeploymentGroup = &configpb.DeploymentGroup{
										Name: name,
									}
									op, err := client.UpdateDeploymentGroup(ctx, req)
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
								Usage: "delete deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment_group", Usage: "The deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment_group"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.DeleteDeploymentGroupRequest{Name: name}
									op, err := client.DeleteDeploymentGroup(ctx, req)
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
								Name:  "list",
								Usage: "list deployment-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "provision",
								Usage: "provision deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment_group", Usage: "The deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment_group"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.ProvisionDeploymentGroupRequest{Name: name}
									op, err := client.ProvisionDeploymentGroup(ctx, req)
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
								Name:  "deprovision",
								Usage: "deprovision deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment_group", Usage: "The deployment_group.", Required: true},
									&cli.BoolFlag{Name: "force", Usage: "The force.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment_group"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.DeprovisionDeploymentGroupRequest{Name: name}
									req.Force = cmd.Bool("force")
									op, err := client.DeprovisionDeploymentGroup(ctx, req)
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
						Name:  "deployments",
						Usage: "Manage deployments resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list deployments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetDeploymentRequest{Name: name}
									resp, err := client.GetDeployment(ctx, req)
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
								Usage: "create deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment-id", Usage: "The deployment id.", Required: true},
									&cli.StringFlag{Name: "artifacts-gcs-bucket", Usage: "The artifacts gcs bucket.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: true},
									&cli.BoolFlag{Name: "import-existing-resources", Usage: "The import existing resources.", Required: false},
									&cli.StringFlag{Name: "worker-pool", Usage: "The worker pool.", Required: false},
									&cli.StringFlag{Name: "tf-version-constraint", Usage: "The tf version constraint.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.CreateDeploymentRequest{Parent: parent}
									req.DeploymentId = cmd.String("deployment-id")
									req.Deployment = &configpb.Deployment{
										ArtifactsGcsBucket:      cmd.String("artifacts-gcs-bucket"),
										ServiceAccount:          cmd.String("service-account"),
										ImportExistingResources: cmd.Bool("import-existing-resources"),
										WorkerPool:              cmd.String("worker-pool"),
										TfVersionConstraint:     cmd.String("tf-version-constraint"),
									}
									op, err := client.CreateDeployment(ctx, req)
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
								Usage: "update deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
									&cli.StringFlag{Name: "artifacts-gcs-bucket", Usage: "The artifacts gcs bucket.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: false},
									&cli.BoolFlag{Name: "import-existing-resources", Usage: "The import existing resources.", Required: false},
									&cli.StringFlag{Name: "worker-pool", Usage: "The worker pool.", Required: false},
									&cli.StringFlag{Name: "tf-version-constraint", Usage: "The tf version constraint.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.UpdateDeploymentRequest{}
									req.Deployment = &configpb.Deployment{
										Name:                    name,
										ArtifactsGcsBucket:      cmd.String("artifacts-gcs-bucket"),
										ServiceAccount:          cmd.String("service-account"),
										ImportExistingResources: cmd.Bool("import-existing-resources"),
										WorkerPool:              cmd.String("worker-pool"),
										TfVersionConstraint:     cmd.String("tf-version-constraint"),
									}
									var paths []string
									if cmd.IsSet("artifacts-gcs-bucket") {
										paths = append(paths, "artifacts_gcs_bucket")
									}
									if cmd.IsSet("service-account") {
										paths = append(paths, "service_account")
									}
									if cmd.IsSet("import-existing-resources") {
										paths = append(paths, "import_existing_resources")
									}
									if cmd.IsSet("worker-pool") {
										paths = append(paths, "worker_pool")
									}
									if cmd.IsSet("tf-version-constraint") {
										paths = append(paths, "tf_version_constraint")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateDeployment(ctx, req)
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
								Usage: "delete deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.DeleteDeploymentRequest{Name: name}
									op, err := client.DeleteDeployment(ctx, req)
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
								Name:  "export-state",
								Usage: "export-state deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									fmt.Printf("Executing export-state on %s\n", name)
									return nil
								},
							},
							{
								Name:  "import-state",
								Usage: "import-state deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									fmt.Printf("Executing import-state on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.DeleteStatefileRequest{Name: name}
									if err := client.DeleteStatefile(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "lock",
								Usage: "lock deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.LockDeploymentRequest{Name: name}
									op, err := client.LockDeployment(ctx, req)
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
								Name:  "unlock",
								Usage: "unlock deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
									&cli.IntFlag{Name: "lock-id", Usage: "The lock id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.UnlockDeploymentRequest{Name: name}
									req.LockId = int64(cmd.Int("lock-id"))
									op, err := client.UnlockDeployment(ctx, req)
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
								Name:  "export-lock",
								Usage: "export-lock deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									fmt.Printf("Executing export-lock on %s\n", name)
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy deployments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy deployments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions deployments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
									client, err := config.NewClient(ctx)
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
									client, err := config.NewClient(ctx)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
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
									client, err := config.NewClient(ctx)
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
									client, err := config.NewClient(ctx)
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
									client, err := config.NewClient(ctx)
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
						Name:  "previews",
						Usage: "Manage previews resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create previews",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "preview-id", Usage: "The preview id.", Required: false},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: true},
									&cli.StringFlag{Name: "artifacts-gcs-bucket", Usage: "The artifacts gcs bucket.", Required: false},
									&cli.StringFlag{Name: "worker-pool", Usage: "The worker pool.", Required: false},
									&cli.StringFlag{Name: "tf-version-constraint", Usage: "The tf version constraint.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.CreatePreviewRequest{Parent: parent}
									req.PreviewId = cmd.String("preview-id")
									req.Preview = &configpb.Preview{
										Deployment:          cmd.String("deployment"),
										ServiceAccount:      cmd.String("service-account"),
										ArtifactsGcsBucket:  cmd.String("artifacts-gcs-bucket"),
										WorkerPool:          cmd.String("worker-pool"),
										TfVersionConstraint: cmd.String("tf-version-constraint"),
									}
									op, err := client.CreatePreview(ctx, req)
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
								Usage: "describe previews",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "preview", Usage: "The preview.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetPreviewRequest{Name: name}
									resp, err := client.GetPreview(ctx, req)
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
								Usage: "list previews",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete previews",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "preview", Usage: "The preview.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.DeletePreviewRequest{Name: name}
									op, err := client.DeletePreview(ctx, req)
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
								Name:  "export",
								Usage: "export previews",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "preview", Usage: "The preview.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/previews/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"))
									fmt.Printf("Executing export on %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configpb.ListResourceChangesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListResourceChanges(ctx, req)
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
								Usage: "describe resource-changes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "preview", Usage: "The preview.", Required: true},
									&cli.StringFlag{Name: "resource_change", Usage: "The resource_change.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/previews/%s/resourceChanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"), cmd.String("resource_change"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetResourceChangeRequest{Name: name}
									resp, err := client.GetResourceChange(ctx, req)
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
						Name:  "resource-drifts",
						Usage: "Manage resource-drifts resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list resource-drifts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configpb.ListResourceDriftsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListResourceDrifts(ctx, req)
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
								Usage: "describe resource-drifts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "preview", Usage: "The preview.", Required: true},
									&cli.StringFlag{Name: "resource_drift", Usage: "The resource_drift.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/previews/%s/resourceDrifts/%s", cmd.String("project"), cmd.String("location"), cmd.String("preview"), cmd.String("resource_drift"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetResourceDriftRequest{Name: name}
									resp, err := client.GetResourceDrift(ctx, req)
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
						Name:  "resources",
						Usage: "Manage resources resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe resources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
									&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
									&cli.StringFlag{Name: "resource", Usage: "The resource.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s/resources/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"), cmd.String("resource"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetResourceRequest{Name: name}
									resp, err := client.GetResource(ctx, req)
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
								Usage: "list resources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configpb.ListResourcesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListResources(ctx, req)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configpb.ListRevisionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRevisions(ctx, req)
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
								Usage: "describe revisions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
									&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetRevisionRequest{Name: name}
									resp, err := client.GetRevision(ctx, req)
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
								Name:  "export-state",
								Usage: "export-state revisions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
									&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployments/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment"), cmd.String("revision"))
									fmt.Printf("Executing export-state on %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe revisions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deployment_group", Usage: "The deployment_group.", Required: true},
									&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deploymentGroups/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("deployment_group"), cmd.String("revision"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetDeploymentGroupRevisionRequest{Name: name}
									resp, err := client.GetDeploymentGroupRevision(ctx, req)
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
								Usage: "list revisions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configpb.ListDeploymentGroupRevisionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDeploymentGroupRevisions(ctx, req)
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
						},
					},
					{
						Name:  "terraform-versions",
						Usage: "Manage terraform-versions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list terraform-versions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe terraform-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "terraform_version", Usage: "The terraform_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/terraformVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("terraform_version"))
									client, err := config.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configpb.GetTerraformVersionRequest{Name: name}
									resp, err := client.GetTerraformVersion(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
