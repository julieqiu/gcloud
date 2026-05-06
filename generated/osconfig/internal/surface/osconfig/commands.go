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

package osconfig

import (
	osconfig "cloud.google.com/go/osconfig/apiv1"
	"cloud.google.com/go/osconfig/apiv1/osconfigpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the osconfig command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "osconfig",
		Usage: "manage OS Config API resources",
		Commands: []*cli.Command{
			{
				Name:  "instance-details",
				Usage: "Manage instance-details resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list instance-details",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters results listed in the response.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of instance details records to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call.", Required: false},
							&cli.StringFlag{Name: "patch-job", Usage: "The ID of the patch job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/patchJobs/%s", cmd.String("project"), cmd.String("patch-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListPatchJobInstanceDetailsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPatchJobInstanceDetails(ctx, req)
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
				Name:  "inventories",
				Usage: "Manage inventories resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list inventories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If provided, this field specifies the criteria that must be met by a.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Inventory view indicating what information should be included in the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListInventoriesRequest{
								Parent:    parent,
								View:      osconfigpb.InventoryView(osconfigpb.InventoryView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListInventories(ctx, req)
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
				Name:  "inventory",
				Usage: "Manage inventory resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe inventory",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Inventory view indicating what information should be included in the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/inventory", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.GetInventoryRequest{
								Name: name,
								View: osconfigpb.InventoryView(osconfigpb.InventoryView_value[cmd.String("view")]),
							}

							resp, err := client.GetInventory(ctx, req)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "os-policy-assignments",
				Usage: "Manage os-policy-assignments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create os-policy-assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment-id", Usage: "The logical name of the OS policy assignment in the project.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.CreateOSPolicyAssignmentRequest{
								Parent:               parent,
								OsPolicyAssignmentId: cmd.String("os-policy-assignment-id"),
							}

							op, err := client.CreateOSPolicyAssignment(ctx, req)
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
						Usage: "update os-policy-assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "os_policy_assignment.name" not yet supported.
							os_policy_assignment_name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"))
							fmt.Printf("Executing update on %s\n", os_policy_assignment_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe os-policy-assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.GetOSPolicyAssignmentRequest{
								Name: name,
							}

							resp, err := client.GetOSPolicyAssignment(ctx, req)
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
						Usage: "list os-policy-assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of assignments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListOSPolicyAssignmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListOSPolicyAssignments(ctx, req)
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
						Usage: "list os-policy-assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of revisions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListOSPolicyAssignmentRevisionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListOSPolicyAssignmentRevisions(ctx, req)
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
						Usage: "delete os-policy-assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os-policy-assignment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteOSPolicyAssignment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.DeleteOSPolicyAssignmentRequest{
								Name: name,
							}

							op, err := client.DeleteOSPolicyAssignment(ctx, req)
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
				Name:  "patch-deployments",
				Usage: "Manage patch-deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create patch-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-deployment-id", Usage: "A name for the patch deployment in the project.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.CreatePatchDeploymentRequest{
								Parent:            parent,
								PatchDeploymentId: cmd.String("patch-deployment-id"),
							}

							resp, err := client.CreatePatchDeployment(ctx, req)
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
						Usage: "describe patch-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-deployment", Usage: "The ID of the patch deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.GetPatchDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetPatchDeployment(ctx, req)
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
						Usage: "list patch-deployments",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of patch deployments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListPatchDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPatchDeployments(ctx, req)
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
						Usage: "delete patch-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-deployment", Usage: "The ID of the patch deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch-deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePatchDeployment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.DeletePatchDeploymentRequest{
								Name: name,
							}

							if err := client.DeletePatchDeployment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update patch-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-deployment", Usage: "The ID of the patch deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "patch_deployment.name" not yet supported.
							patch_deployment_name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch-deployment"))
							fmt.Printf("Executing update on %s\n", patch_deployment_name)
							return nil
						},
					},

					{
						Name:  "pause",
						Usage: "pause patch-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-deployment", Usage: "The ID of the patch deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.PausePatchDeploymentRequest{
								Name: name,
							}

							resp, err := client.PausePatchDeployment(ctx, req)
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
						Name:  "resume",
						Usage: "resume patch-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-deployment", Usage: "The ID of the patch deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ResumePatchDeploymentRequest{
								Name: name,
							}

							resp, err := client.ResumePatchDeployment(ctx, req)
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
				Name:  "patch-jobs",
				Usage: "Manage patch-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "execute",
						Usage: "execute patch-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "description", Usage: "Description of the patch job.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "Display name for this patch job.", Required: false},
							&cli.BoolFlag{Name: "dry-run", Usage: "If this patch is a dry-run only, instances are contacted but.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ExecutePatchJobRequest{
								Parent:      parent,
								Description: cmd.String("description"),
								DryRun:      cmd.Bool("dry-run"),
								DisplayName: cmd.String("display-name"),
							}

							resp, err := client.ExecutePatchJob(ctx, req)
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
						Usage: "describe patch-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-job", Usage: "The ID of the patch job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/patchJobs/%s", cmd.String("project"), cmd.String("patch-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.GetPatchJobRequest{
								Name: name,
							}

							resp, err := client.GetPatchJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel patch-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "patch-job", Usage: "The ID of the patch job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/patchJobs/%s", cmd.String("project"), cmd.String("patch-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.CancelPatchJobRequest{
								Name: name,
							}

							resp, err := client.CancelPatchJob(ctx, req)
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
						Usage: "list patch-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If provided, this field specifies the criteria that must be met by patch.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of instance status to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListPatchJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPatchJobs(ctx, req)
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
				Name:  "report",
				Usage: "Manage report resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe report",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/osPolicyAssignments/%s/report", cmd.String("project"), cmd.String("location"), cmd.String("instance"), cmd.String("os-policy-assignment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.GetOSPolicyAssignmentReportRequest{
								Name: name,
							}

							resp, err := client.GetOSPolicyAssignmentReport(ctx, req)
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
				Name:  "reports",
				Usage: "Manage reports resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list reports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If provided, this field specifies the criteria that must be met by the.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "os-policy-assignment", Usage: "The ID of the os policy assignment.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call to the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"), cmd.String("os-policy-assignment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListOSPolicyAssignmentReportsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								Filter:    cmd.String("filter"),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListOSPolicyAssignmentReports(ctx, req)
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
				Name:  "vulnerability-report",
				Usage: "Manage vulnerability-report resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe vulnerability-report",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/vulnerabilityReport", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.GetVulnerabilityReportRequest{
								Name: name,
							}

							resp, err := client.GetVulnerabilityReport(ctx, req)
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
				Name:  "vulnerability-reports",
				Usage: "Manage vulnerability-reports resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list vulnerability-reports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "This field supports filtering by the severity level for the vulnerability.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := osconfig.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osconfigpb.ListVulnerabilityReportsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListVulnerabilityReports(ctx, req)
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
		},
	}
}
