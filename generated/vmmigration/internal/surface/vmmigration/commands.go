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

package vmmigration

import (
	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	"cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the vmmigration command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "vmmigration",
		Usage: "manage VM Migration API resources",
		Commands: []*cli.Command{
			{
				Name:  "clone-jobs",
				Usage: "Manage clone-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create clone-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "clone-job-id", Usage: "The clone job identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateCloneJobRequest{
								Parent:     parent,
								CloneJobId: cmd.String("clone-job-id"),
								RequestId:  cmd.String("request-id"),
							}

							op, err := client.CreateCloneJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel clone-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "clone-job", Usage: "The ID of the clone job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cloneJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"), cmd.String("clone-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CancelCloneJobRequest{
								Name: name,
							}

							op, err := client.CancelCloneJob(ctx, req)
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
						Usage: "list clone-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of clone jobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListCloneJobs` call.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListCloneJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCloneJobs(ctx, req)
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
						Usage: "describe clone-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "clone-job", Usage: "The ID of the clone job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cloneJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"), cmd.String("clone-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetCloneJobRequest{
								Name: name,
							}

							resp, err := client.GetCloneJob(ctx, req)
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
				Name:  "cutover-jobs",
				Usage: "Manage cutover-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create cutover-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cutover-job-id", Usage: "The cutover job identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateCutoverJobRequest{
								Parent:       parent,
								CutoverJobId: cmd.String("cutover-job-id"),
								RequestId:    cmd.String("request-id"),
							}

							op, err := client.CreateCutoverJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel cutover-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cutover-job", Usage: "The ID of the cutover job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cutoverJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"), cmd.String("cutover-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CancelCutoverJobRequest{
								Name: name,
							}

							op, err := client.CancelCutoverJob(ctx, req)
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
						Usage: "list cutover-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of cutover jobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListCutoverJobs` call.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListCutoverJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCutoverJobs(ctx, req)
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
						Usage: "describe cutover-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cutover-job", Usage: "The ID of the cutover job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/cutoverJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"), cmd.String("cutover-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetCutoverJobRequest{
								Name: name,
							}

							resp, err := client.GetCutoverJob(ctx, req)
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
				Name:  "datacenter-connectors",
				Usage: "Manage datacenter-connectors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list datacenter-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of connectors to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDatacenterConnectors`.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListDatacenterConnectorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDatacenterConnectors(ctx, req)
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
						Usage: "describe datacenter-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "datacenter-connector", Usage: "The ID of the datacenter connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/datacenterConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("datacenter-connector"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetDatacenterConnectorRequest{
								Name: name,
							}

							resp, err := client.GetDatacenterConnector(ctx, req)
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
						Usage: "create datacenter-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "datacenter-connector-id", Usage: "The datacenterConnector identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateDatacenterConnectorRequest{
								Parent:                parent,
								DatacenterConnectorId: cmd.String("datacenter-connector-id"),
								RequestId:             cmd.String("request-id"),
							}

							op, err := client.CreateDatacenterConnector(ctx, req)
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
						Usage: "delete datacenter-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "datacenter-connector", Usage: "The ID of the datacenter connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/datacenterConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("datacenter-connector"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDatacenterConnector %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteDatacenterConnectorRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteDatacenterConnector(ctx, req)
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
						Name:  "upgrade-appliance",
						Usage: "upgrade-appliance datacenter-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "datacenter-connector", Usage: "The ID of the datacenter connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							datacenter_connector := fmt.Sprintf("projects/%s/locations/%s/sources/%s/datacenterConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("datacenter-connector"))
							fmt.Printf("Executing upgrade-appliance on %s\n", datacenter_connector)
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
							&cli.StringFlag{Name: "disk-migration-job-id", Usage: "The DiskMigrationJob identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateDiskMigrationJobRequest{
								Parent:             parent,
								DiskMigrationJobId: cmd.String("disk-migration-job-id"),
								RequestId:          cmd.String("request-id"),
							}

							op, err := client.CreateDiskMigrationJob(ctx, req)
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
						Usage: "list disk-migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request (according to <a.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Ordering of the result list.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of disk migration jobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDiskMigrationJobs`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListDiskMigrationJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDiskMigrationJobs(ctx, req)
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
						Usage: "describe disk-migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk-migration-job", Usage: "The ID of the disk migration job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk-migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetDiskMigrationJobRequest{
								Name: name,
							}

							resp, err := client.GetDiskMigrationJob(ctx, req)
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
						Usage: "update disk-migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk-migration-job", Usage: "The ID of the disk migration job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "disk_migration_job.name" not yet supported.
							disk_migration_job_name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk-migration-job"))
							fmt.Printf("Executing update on %s\n", disk_migration_job_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete disk-migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk-migration-job", Usage: "The ID of the disk migration job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk-migration-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDiskMigrationJob %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteDiskMigrationJobRequest{
								Name: name,
							}

							op, err := client.DeleteDiskMigrationJob(ctx, req)
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
						Name:  "run",
						Usage: "run disk-migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk-migration-job", Usage: "The ID of the disk migration job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk-migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.RunDiskMigrationJobRequest{
								Name: name,
							}

							op, err := client.RunDiskMigrationJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel disk-migration-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk-migration-job", Usage: "The ID of the disk migration job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/diskMigrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("disk-migration-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CancelDiskMigrationJobRequest{
								Name: name,
							}

							op, err := client.CancelDiskMigrationJob(ctx, req)
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
				Name:  "groups",
				Usage: "Manage groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of groups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListGroups` call.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGroups(ctx, req)
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
						Usage: "describe groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetGroupRequest{
								Name: name,
							}

							resp, err := client.GetGroup(ctx, req)
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
						Usage: "create groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group-id", Usage: "The group identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateGroupRequest{
								Parent:    parent,
								GroupId:   cmd.String("group-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateGroup(ctx, req)
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
						Usage: "update groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "group.name" not yet supported.
							group_name := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
							fmt.Printf("Executing update on %s\n", group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteGroup(ctx, req)
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
						Name:  "add-group-migration",
						Usage: "add-group-migration groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The full path name of the MigratingVm to add.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							group := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
							fmt.Printf("Executing add-group-migration on %s\n", group)
							return nil
						},
					},

					{
						Name:  "remove-group-migration",
						Usage: "remove-group-migration groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The MigratingVm to remove.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							group := fmt.Sprintf("projects/%s/locations/%s/groups/%s", cmd.String("project"), cmd.String("location"), cmd.String("group"))
							fmt.Printf("Executing remove-group-migration on %s\n", group)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter request (according to <a.", Required: false},
							&cli.StringFlag{Name: "image-import", Usage: "The ID of the image import.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The order by fields for the result (according to <a.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of targets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListImageImportJobs`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("image-import"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListImageImportJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListImageImportJobs(ctx, req)
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
						Usage: "describe image-import-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image-import", Usage: "The ID of the image import.", Required: true},
							&cli.StringFlag{Name: "image-import-job", Usage: "The ID of the image import job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s/imageImportJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("image-import"), cmd.String("image-import-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetImageImportJobRequest{
								Name: name,
							}

							resp, err := client.GetImageImportJob(ctx, req)
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
						Usage: "cancel image-import-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image-import", Usage: "The ID of the image import.", Required: true},
							&cli.StringFlag{Name: "image-import-job", Usage: "The ID of the image import job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s/imageImportJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("image-import"), cmd.String("image-import-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CancelImageImportJobRequest{
								Name: name,
							}

							op, err := client.CancelImageImportJob(ctx, req)
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
				Name:  "image-imports",
				Usage: "Manage image-imports resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list image-imports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request (according to <a.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The order by fields for the result (according to <a.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of targets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListImageImports` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListImageImportsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListImageImports(ctx, req)
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
						Usage: "describe image-imports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image-import", Usage: "The ID of the image import.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("image-import"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetImageImportRequest{
								Name: name,
							}

							resp, err := client.GetImageImport(ctx, req)
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
						Usage: "create image-imports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image-import-id", Usage: "The image import identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateImageImportRequest{
								Parent:        parent,
								ImageImportId: cmd.String("image-import-id"),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.CreateImageImport(ctx, req)
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
						Usage: "delete image-imports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image-import", Usage: "The ID of the image import.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/imageImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("image-import"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteImageImport %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteImageImportRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteImageImport(ctx, req)
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
				Name:  "migrating-vms",
				Usage: "Manage migrating-vms resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm-id", Usage: "The migratingVm identifier.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateMigratingVmRequest{
								Parent:        parent,
								MigratingVmId: cmd.String("migrating-vm-id"),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.CreateMigratingVm(ctx, req)
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
						Usage: "list migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of migrating VMs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListMigratingVms` call.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The level of details of each migrating VM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListMigratingVmsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								View:      vmmigrationpb.MigratingVmView(vmmigrationpb.MigratingVmView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListMigratingVms(ctx, req)
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
						Usage: "describe migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The level of details of the migrating VM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetMigratingVmRequest{
								Name: name,
								View: vmmigrationpb.MigratingVmView(vmmigrationpb.MigratingVmView_value[cmd.String("view")]),
							}

							resp, err := client.GetMigratingVm(ctx, req)
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
						Usage: "update migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "migrating_vm.name" not yet supported.
							migrating_vm_name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							fmt.Printf("Executing update on %s\n", migrating_vm_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMigratingVm %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteMigratingVmRequest{
								Name: name,
							}

							op, err := client.DeleteMigratingVm(ctx, req)
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
						Name:  "start-migration",
						Usage: "start-migration migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migrating_vm := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							fmt.Printf("Executing start-migration on %s\n", migrating_vm)
							return nil
						},
					},

					{
						Name:  "resume-migration",
						Usage: "resume-migration migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migrating_vm := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							fmt.Printf("Executing resume-migration on %s\n", migrating_vm)
							return nil
						},
					},

					{
						Name:  "pause-migration",
						Usage: "pause-migration migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migrating_vm := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							fmt.Printf("Executing pause-migration on %s\n", migrating_vm)
							return nil
						},
					},

					{
						Name:  "finalize-migration",
						Usage: "finalize-migration migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migrating_vm := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							fmt.Printf("Executing finalize-migration on %s\n", migrating_vm)
							return nil
						},
					},

					{
						Name:  "extend-migration",
						Usage: "extend-migration migrating-vms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							migrating_vm := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							fmt.Printf("Executing extend-migration on %s\n", migrating_vm)
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
				Name:  "replication-cycles",
				Usage: "Manage replication-cycles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list replication-cycles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of replication cycles to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListReplicationCycles`.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListReplicationCyclesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListReplicationCycles(ctx, req)
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
						Usage: "describe replication-cycles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "migrating-vm", Usage: "The ID of the migrating vm.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication-cycle", Usage: "The ID of the replication cycle.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/migratingVms/%s/replicationCycles/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("migrating-vm"), cmd.String("replication-cycle"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetReplicationCycleRequest{
								Name: name,
							}

							resp, err := client.GetReplicationCycle(ctx, req)
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
				Name:  "sources",
				Usage: "Manage sources resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of sources to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSources` call.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListSourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListSources(ctx, req)
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
						Usage: "describe sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetSourceRequest{
								Name: name,
							}

							resp, err := client.GetSource(ctx, req)
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
						Usage: "create sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-id", Usage: "The source identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateSourceRequest{
								Parent:    parent,
								SourceId:  cmd.String("source-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateSource(ctx, req)
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
						Usage: "update sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "source.name" not yet supported.
							source_name := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							fmt.Printf("Executing update on %s\n", source_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSource %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteSourceRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteSource(ctx, req)
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
						Name:  "fetch-inventory",
						Usage: "fetch-inventory sources",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force-refresh", Usage: "If this flag is set to true, the source will be queried instead of using.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							source := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							fmt.Printf("Executing fetch-inventory on %s\n", source)
							return nil
						},
					},

					{
						Name:  "fetch-storage-inventory",
						Usage: "fetch-storage-inventory sources",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force-refresh", Usage: "If this flag is set to true, the source will be queried instead.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of VMs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `FetchStorageInventory`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The type of the storage inventory to fetch.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							source := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							fmt.Printf("Executing fetch-storage-inventory on %s\n", source)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of targets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListTargets` call.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListTargetProjectsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTargetProjects(ctx, req)
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
						Usage: "describe target-projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target-project", Usage: "The ID of the target project.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/targetProjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("target-project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetTargetProjectRequest{
								Name: name,
							}

							resp, err := client.GetTargetProject(ctx, req)
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
						Usage: "create target-projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target-project-id", Usage: "The target_project identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateTargetProjectRequest{
								Parent:          parent,
								TargetProjectId: cmd.String("target-project-id"),
								RequestId:       cmd.String("request-id"),
							}

							op, err := client.CreateTargetProject(ctx, req)
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
						Usage: "update target-projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target-project", Usage: "The ID of the target project.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "target_project.name" not yet supported.
							target_project_name := fmt.Sprintf("projects/%s/locations/%s/targetProjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("target-project"))
							fmt.Printf("Executing update on %s\n", target_project_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target-project", Usage: "The ID of the target project.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/targetProjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("target-project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTargetProject %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteTargetProjectRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteTargetProject(ctx, req)
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
				Name:  "utilization-reports",
				Usage: "Manage utilization-reports resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list utilization-reports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "the order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of reports to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListUtilizationReports`.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The level of details of each report.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.ListUtilizationReportsRequest{
								Parent:    parent,
								View:      vmmigrationpb.UtilizationReportView(vmmigrationpb.UtilizationReportView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListUtilizationReports(ctx, req)
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
						Usage: "describe utilization-reports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "utilization-report", Usage: "The ID of the utilization report.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The level of details of the report.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/utilizationReports/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("utilization-report"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.GetUtilizationReportRequest{
								Name: name,
								View: vmmigrationpb.UtilizationReportView(vmmigrationpb.UtilizationReportView_value[cmd.String("view")]),
							}

							resp, err := client.GetUtilizationReport(ctx, req)
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
						Usage: "create utilization-reports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "utilization-report-id", Usage: "The ID to use for the report, which will become the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/sources/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.CreateUtilizationReportRequest{
								Parent:              parent,
								UtilizationReportId: cmd.String("utilization-report-id"),
								RequestId:           cmd.String("request-id"),
							}

							op, err := client.CreateUtilizationReport(ctx, req)
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
						Usage: "delete utilization-reports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "utilization-report", Usage: "The ID of the utilization report.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sources/%s/utilizationReports/%s", cmd.String("project"), cmd.String("location"), cmd.String("source"), cmd.String("utilization-report"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteUtilizationReport %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := vmmigration.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &vmmigrationpb.DeleteUtilizationReportRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteUtilizationReport(ctx, req)
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
		},
	}
}
