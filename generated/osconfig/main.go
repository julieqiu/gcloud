package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	osconfig "cloud.google.com/go/osconfig/apiv1"
	"cloud.google.com/go/osconfig/apiv1/osconfigpb"
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
				Name:  "osconfig",
				Usage: "manage OS Config API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "instance-details",
						Usage: "Manage instance-details resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list instance-details",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := osconfig.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &osconfigpb.ListPatchJobInstanceDetailsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPatchJobInstanceDetails(ctx, req)
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
						Name:  "inventories",
						Usage: "Manage inventories resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list inventories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/zones/%s", cmd.String("project"), cmd.String("zone"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &osconfigpb.ListInventoriesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInventories(ctx, req)
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
						Name:  "inventory",
						Usage: "Manage inventory resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe inventory",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/inventory", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.GetInventoryRequest{Name: name}
									resp, err := client.GetInventory(ctx, req)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "osPolicyAssignment", Usage: "The osPolicyAssignment.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("osPolicyAssignment"), cmd.String("operation"))
									client, err := osconfig.NewClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "osPolicyAssignment", Usage: "The osPolicyAssignment.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("osPolicyAssignment"), cmd.String("operation"))
									client, err := osconfig.NewClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "osPolicyAssignment", Usage: "The osPolicyAssignment.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("osPolicyAssignment"), cmd.String("operation"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "osPolicyAssignment", Usage: "The osPolicyAssignment.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("osPolicyAssignment"), cmd.String("operation"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
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
						Name:  "os-policy-assignments",
						Usage: "Manage os-policy-assignments resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create os-policy-assignments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "os-policy-assignment-id", Usage: "The os policy assignment id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.CreateOSPolicyAssignmentRequest{Parent: parent}
									req.OsPolicyAssignmentId = cmd.String("os-policy-assignment-id")
									req.OsPolicyAssignment = &osconfigpb.OSPolicyAssignment{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
									}
									op, err := client.CreateOSPolicyAssignment(ctx, req)
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
								Usage: "update os-policy-assignments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "os_policy_assignment", Usage: "The os_policy_assignment.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os_policy_assignment"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.UpdateOSPolicyAssignmentRequest{}
									req.OsPolicyAssignment = &osconfigpb.OSPolicyAssignment{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateOSPolicyAssignment(ctx, req)
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
								Usage: "describe os-policy-assignments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "os_policy_assignment", Usage: "The os_policy_assignment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os_policy_assignment"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.GetOSPolicyAssignmentRequest{Name: name}
									resp, err := client.GetOSPolicyAssignment(ctx, req)
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
								Usage: "list os-policy-assignments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list os-policy-assignments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &osconfigpb.ListOSPolicyAssignmentRevisionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOSPolicyAssignmentRevisions(ctx, req)
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
								Usage: "delete os-policy-assignments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "os_policy_assignment", Usage: "The os_policy_assignment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/osPolicyAssignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("os_policy_assignment"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.DeleteOSPolicyAssignmentRequest{Name: name}
									op, err := client.DeleteOSPolicyAssignment(ctx, req)
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
						Name:  "patch-deployments",
						Usage: "Manage patch-deployments resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create patch-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch-deployment-id", Usage: "The patch deployment id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := osconfig.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.CreatePatchDeploymentRequest{Parent: parent}
									req.PatchDeploymentId = cmd.String("patch-deployment-id")
									req.PatchDeployment = &osconfigpb.PatchDeployment{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreatePatchDeployment(ctx, req)
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
								Usage: "describe patch-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_deployment", Usage: "The patch_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch_deployment"))
									client, err := osconfig.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.GetPatchDeploymentRequest{Name: name}
									resp, err := client.GetPatchDeployment(ctx, req)
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
								Usage: "list patch-deployments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete patch-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_deployment", Usage: "The patch_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch_deployment"))
									client, err := osconfig.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.DeletePatchDeploymentRequest{Name: name}
									if err := client.DeletePatchDeployment(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update patch-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_deployment", Usage: "The patch_deployment.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch_deployment"))
									client, err := osconfig.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.UpdatePatchDeploymentRequest{}
									req.PatchDeployment = &osconfigpb.PatchDeployment{
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
									resp, err := client.UpdatePatchDeployment(ctx, req)
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
								Name:  "pause",
								Usage: "pause patch-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_deployment", Usage: "The patch_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch_deployment"))
									fmt.Printf("Executing pause on %s\n", name)
									return nil
								},
							},
							{
								Name:  "resume",
								Usage: "resume patch-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_deployment", Usage: "The patch_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchDeployments/%s", cmd.String("project"), cmd.String("patch_deployment"))
									fmt.Printf("Executing resume on %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing execute...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe patch-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_job", Usage: "The patch_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchJobs/%s", cmd.String("project"), cmd.String("patch_job"))
									client, err := osconfig.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.GetPatchJobRequest{Name: name}
									resp, err := client.GetPatchJob(ctx, req)
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
								Name:  "cancel",
								Usage: "cancel patch-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "patch_job", Usage: "The patch_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/patchJobs/%s", cmd.String("project"), cmd.String("patch_job"))
									fmt.Printf("Executing cancel on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list patch-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
									&cli.StringFlag{Name: "assignment", Usage: "The assignment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/osPolicyAssignments/%s/report", cmd.String("project"), cmd.String("location"), cmd.String("instance"), cmd.String("assignment"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.GetOSPolicyAssignmentReportRequest{Name: name}
									resp, err := client.GetOSPolicyAssignmentReport(ctx, req)
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
						Name:  "reports",
						Usage: "Manage reports resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list reports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &osconfigpb.ListOSPolicyAssignmentReportsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOSPolicyAssignmentReports(ctx, req)
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
						Name:  "vulnerability-report",
						Usage: "Manage vulnerability-report resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe vulnerability-report",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/vulnerabilityReport", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osconfigpb.GetVulnerabilityReportRequest{Name: name}
									resp, err := client.GetVulnerabilityReport(ctx, req)
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
						Name:  "vulnerability-reports",
						Usage: "Manage vulnerability-reports resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list vulnerability-reports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/zones/%s", cmd.String("project"), cmd.String("zone"))
									client, err := osconfig.NewOsConfigZonalClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &osconfigpb.ListVulnerabilityReportsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVulnerabilityReports(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
