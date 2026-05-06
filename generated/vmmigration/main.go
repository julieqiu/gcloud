package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	"cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
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
				Name:  "vmmigration",
				Usage: "manage VM Migration API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "clone-jobs",
						Usage: "Manage clone-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create clone-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "clone-job-id", Usage: "The clone job id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateCloneJobRequest{Parent: parent}
									req.CloneJobId = cmd.String("clone-job-id")
									op, err := client.CreateCloneJob(ctx, req)
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
								Name:  "cancel",
								Usage: "cancel clone-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "clone_job", Usage: "The clone_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cloneJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"), cmd.String("clone_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CancelCloneJobRequest{Name: name}
									op, err := client.CancelCloneJob(ctx, req)
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
								Usage: "list clone-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListCloneJobsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCloneJobs(ctx, req)
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
								Usage: "describe clone-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "clone_job", Usage: "The clone_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cloneJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"), cmd.String("clone_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetCloneJobRequest{Name: name}
									resp, err := client.GetCloneJob(ctx, req)
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
						Name:  "cutover-jobs",
						Usage: "Manage cutover-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create cutover-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "cutover-job-id", Usage: "The cutover job id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateCutoverJobRequest{Parent: parent}
									req.CutoverJobId = cmd.String("cutover-job-id")
									op, err := client.CreateCutoverJob(ctx, req)
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
								Name:  "cancel",
								Usage: "cancel cutover-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "cutover_job", Usage: "The cutover_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cutoverJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"), cmd.String("cutover_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CancelCutoverJobRequest{Name: name}
									op, err := client.CancelCutoverJob(ctx, req)
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
								Usage: "list cutover-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListCutoverJobsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCutoverJobs(ctx, req)
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
								Usage: "describe cutover-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "cutover_job", Usage: "The cutover_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cutoverJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"), cmd.String("cutover_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetCutoverJobRequest{Name: name}
									resp, err := client.GetCutoverJob(ctx, req)
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
						Name:  "datacenter-connectors",
						Usage: "Manage datacenter-connectors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list datacenter-connectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListDatacenterConnectorsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDatacenterConnectors(ctx, req)
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
								Usage: "describe datacenter-connectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "datacenter_connector", Usage: "The datacenter_connector.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/datacenterConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("datacenter_connector"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetDatacenterConnectorRequest{Name: name}
									resp, err := client.GetDatacenterConnector(ctx, req)
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
								Usage: "create datacenter-connectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "datacenter-connector-id", Usage: "The datacenter connector id.", Required: true},
									&cli.StringFlag{Name: "registration-id", Usage: "The registration id.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: false},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateDatacenterConnectorRequest{Parent: parent}
									req.DatacenterConnectorId = cmd.String("datacenter-connector-id")
									req.DatacenterConnector = &vmmigrationpb.DatacenterConnector{
										RegistrationId: cmd.String("registration-id"),
										ServiceAccount: cmd.String("service-account"),
										Version:        cmd.String("version"),
									}
									op, err := client.CreateDatacenterConnector(ctx, req)
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
								Usage: "delete datacenter-connectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "datacenter_connector", Usage: "The datacenter_connector.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/datacenterConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("datacenter_connector"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteDatacenterConnectorRequest{Name: name}
									op, err := client.DeleteDatacenterConnector(ctx, req)
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
								Name:  "upgrade-appliance",
								Usage: "upgrade-appliance datacenter-connectors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing upgrade-appliance...")
									return nil
								},
							},
						},
					},
					{
						Name:  "disk-migration-jobs",
						Usage: "Manage disk-migration-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "disk-migration-job-id", Usage: "The disk migration job id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateDiskMigrationJobRequest{Parent: parent}
									req.DiskMigrationJobId = cmd.String("disk-migration-job-id")
									op, err := client.CreateDiskMigrationJob(ctx, req)
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
								Usage: "list disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListDiskMigrationJobsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDiskMigrationJobs(ctx, req)
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
								Usage: "describe disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "disk_migration_job", Usage: "The disk_migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk_migration_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetDiskMigrationJobRequest{Name: name}
									resp, err := client.GetDiskMigrationJob(ctx, req)
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
								Usage: "update disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "disk_migration_job", Usage: "The disk_migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk_migration_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.UpdateDiskMigrationJobRequest{}
									req.DiskMigrationJob = &vmmigrationpb.DiskMigrationJob{
										Name: name,
									}
									op, err := client.UpdateDiskMigrationJob(ctx, req)
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
								Usage: "delete disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "disk_migration_job", Usage: "The disk_migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk_migration_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteDiskMigrationJobRequest{Name: name}
									op, err := client.DeleteDiskMigrationJob(ctx, req)
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
								Name:  "run",
								Usage: "run disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "disk_migration_job", Usage: "The disk_migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk_migration_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.RunDiskMigrationJobRequest{Name: name}
									op, err := client.RunDiskMigrationJob(ctx, req)
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
								Name:  "cancel",
								Usage: "cancel disk-migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "disk_migration_job", Usage: "The disk_migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk_migration_job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CancelDiskMigrationJobRequest{Name: name}
									op, err := client.CancelDiskMigrationJob(ctx, req)
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
						Name:  "groups",
						Usage: "Manage groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGroups(ctx, req)
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
								Usage: "describe groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetGroupRequest{Name: name}
									resp, err := client.GetGroup(ctx, req)
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
								Usage: "create groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "group-id", Usage: "The group id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateGroupRequest{Parent: parent}
									req.GroupId = cmd.String("group-id")
									req.Group = &vmmigrationpb.Group{
										Description: cmd.String("description"),
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreateGroup(ctx, req)
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
								Usage: "update groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.UpdateGroupRequest{}
									req.Group = &vmmigrationpb.Group{
										Name:        name,
										Description: cmd.String("description"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateGroup(ctx, req)
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
								Usage: "delete groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteGroupRequest{Name: name}
									op, err := client.DeleteGroup(ctx, req)
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
								Name:  "add-group-migration",
								Usage: "add-group-migration groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-group-migration...")
									return nil
								},
							},
							{
								Name:  "remove-group-migration",
								Usage: "remove-group-migration groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-group-migration...")
									return nil
								},
							},
						},
					},
					{
						Name:  "image-import-jobs",
						Usage: "Manage image-import-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list image-import-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListImageImportJobsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListImageImportJobs(ctx, req)
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
								Usage: "describe image-import-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
									&cli.StringFlag{Name: "result", Usage: "The result.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s/imageImportJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("result"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetImageImportJobRequest{Name: name}
									resp, err := client.GetImageImportJob(ctx, req)
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
								Usage: "cancel image-import-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
									&cli.StringFlag{Name: "result", Usage: "The result.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s/imageImportJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("result"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CancelImageImportJobRequest{Name: name}
									op, err := client.CancelImageImportJob(ctx, req)
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
						Name:  "image-imports",
						Usage: "Manage image-imports resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list image-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListImageImportsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListImageImports(ctx, req)
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
								Usage: "describe image-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetImageImportRequest{Name: name}
									resp, err := client.GetImageImport(ctx, req)
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
								Usage: "create image-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "image-import-id", Usage: "The image import id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateImageImportRequest{Parent: parent}
									req.ImageImportId = cmd.String("image-import-id")
									op, err := client.CreateImageImport(ctx, req)
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
								Usage: "delete image-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteImageImportRequest{Name: name}
									op, err := client.DeleteImageImport(ctx, req)
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
									client, err := vmmigration.NewClient(ctx)
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
									client, err := vmmigration.NewClient(ctx)
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
						Name:  "migrating-vms",
						Usage: "Manage migrating-vms resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create migrating-vms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating-vm-id", Usage: "The migrating vm id.", Required: true},
									&cli.StringFlag{Name: "source-vm-id", Usage: "The source vm id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateMigratingVmRequest{Parent: parent}
									req.MigratingVmId = cmd.String("migrating-vm-id")
									req.MigratingVm = &vmmigrationpb.MigratingVm{
										SourceVmId:  cmd.String("source-vm-id"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateMigratingVm(ctx, req)
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
								Usage: "list migrating-vms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListMigratingVmsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMigratingVms(ctx, req)
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
								Usage: "describe migrating-vms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetMigratingVmRequest{Name: name}
									resp, err := client.GetMigratingVm(ctx, req)
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
								Usage: "update migrating-vms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "source-vm-id", Usage: "The source vm id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.UpdateMigratingVmRequest{}
									req.MigratingVm = &vmmigrationpb.MigratingVm{
										Name:        name,
										SourceVmId:  cmd.String("source-vm-id"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("source-vm-id") {
										paths = append(paths, "source_vm_id")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMigratingVm(ctx, req)
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
								Usage: "delete migrating-vms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteMigratingVmRequest{Name: name}
									op, err := client.DeleteMigratingVm(ctx, req)
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
								Name:  "start-migration",
								Usage: "start-migration migrating-vms",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-migration...")
									return nil
								},
							},
							{
								Name:  "resume-migration",
								Usage: "resume-migration migrating-vms",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resume-migration...")
									return nil
								},
							},
							{
								Name:  "pause-migration",
								Usage: "pause-migration migrating-vms",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing pause-migration...")
									return nil
								},
							},
							{
								Name:  "finalize-migration",
								Usage: "finalize-migration migrating-vms",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing finalize-migration...")
									return nil
								},
							},
							{
								Name:  "extend-migration",
								Usage: "extend-migration migrating-vms",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing extend-migration...")
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
									client, err := vmmigration.NewClient(ctx)
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
									client, err := vmmigration.NewClient(ctx)
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
									client, err := vmmigration.NewClient(ctx)
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
									client, err := vmmigration.NewClient(ctx)
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
						Name:  "replication-cycles",
						Usage: "Manage replication-cycles resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list replication-cycles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListReplicationCyclesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListReplicationCycles(ctx, req)
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
								Usage: "describe replication-cycles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "migrating_vm", Usage: "The migrating_vm.", Required: true},
									&cli.StringFlag{Name: "replication_cycle", Usage: "The replication_cycle.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/replicationCycles/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating_vm"), cmd.String("replication_cycle"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetReplicationCycleRequest{Name: name}
									resp, err := client.GetReplicationCycle(ctx, req)
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
						Name:  "sources",
						Usage: "Manage sources resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListSourcesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSources(ctx, req)
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
								Usage: "describe sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetSourceRequest{Name: name}
									resp, err := client.GetSource(ctx, req)
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
								Usage: "create sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source-id", Usage: "The source id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateSourceRequest{Parent: parent}
									req.SourceId = cmd.String("source-id")
									req.Source = &vmmigrationpb.Source{
										Description: cmd.String("description"),
									}
									op, err := client.CreateSource(ctx, req)
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
								Usage: "update sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.UpdateSourceRequest{}
									req.Source = &vmmigrationpb.Source{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateSource(ctx, req)
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
								Usage: "delete sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteSourceRequest{Name: name}
									op, err := client.DeleteSource(ctx, req)
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
								Name:  "fetch-inventory",
								Usage: "fetch-inventory sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-inventory...")
									return nil
								},
							},
							{
								Name:  "fetch-storage-inventory",
								Usage: "fetch-storage-inventory sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-storage-inventory...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-projects",
						Usage: "Manage target-projects resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list target-projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListTargetProjectsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTargetProjects(ctx, req)
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
								Usage: "describe target-projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target_project", Usage: "The target_project.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/targetProjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("target_project"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetTargetProjectRequest{Name: name}
									resp, err := client.GetTargetProject(ctx, req)
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
								Usage: "create target-projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target-project-id", Usage: "The target project id.", Required: true},
									&cli.StringFlag{Name: "project", Usage: "The project.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateTargetProjectRequest{Parent: parent}
									req.TargetProjectId = cmd.String("target-project-id")
									req.TargetProject = &vmmigrationpb.TargetProject{
										Project:     cmd.String("project"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateTargetProject(ctx, req)
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
								Usage: "update target-projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target_project", Usage: "The target_project.", Required: true},
									&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/targetProjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("target_project"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.UpdateTargetProjectRequest{}
									req.TargetProject = &vmmigrationpb.TargetProject{
										Name:        name,
										Project:     cmd.String("project"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("project") {
										paths = append(paths, "project")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTargetProject(ctx, req)
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
								Usage: "delete target-projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target_project", Usage: "The target_project.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/targetProjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("target_project"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteTargetProjectRequest{Name: name}
									op, err := client.DeleteTargetProject(ctx, req)
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
						Name:  "utilization-reports",
						Usage: "Manage utilization-reports resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list utilization-reports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &vmmigrationpb.ListUtilizationReportsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListUtilizationReports(ctx, req)
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
								Usage: "describe utilization-reports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "utilization_report", Usage: "The utilization_report.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/utilizationReports/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("utilization_report"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.GetUtilizationReportRequest{Name: name}
									resp, err := client.GetUtilizationReport(ctx, req)
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
								Usage: "create utilization-reports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "utilization-report-id", Usage: "The utilization report id.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.CreateUtilizationReportRequest{Parent: parent}
									req.UtilizationReportId = cmd.String("utilization-report-id")
									req.UtilizationReport = &vmmigrationpb.UtilizationReport{
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreateUtilizationReport(ctx, req)
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
								Usage: "delete utilization-reports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "utilization_report", Usage: "The utilization_report.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/utilizationReports/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("utilization_report"))
									client, err := vmmigration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &vmmigrationpb.DeleteUtilizationReportRequest{Name: name}
									op, err := client.DeleteUtilizationReport(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
