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

package gkebackup

import (
	gkebackup "cloud.google.com/go/gkebackup/apiv1"
	"cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the gkebackup command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkebackup",
		Usage: "manage Backup for GKE API resources",
		Commands: []*cli.Command{
			{
				Name:  "backup-channels",
				Usage: "Manage backup-channels resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-channel-id", Usage: "The client-provided short name for the BackupChannel resource.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.CreateBackupChannelRequest{
								Parent:          parent,
								BackupChannelId: cmd.String("backup-channel-id"),
							}

							op, err := client.CreateBackupChannel(ctx, req)
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
						Usage: "list backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListBackupChannelsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupChannels(ctx, req)
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
						Usage: "describe backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-channel", Usage: "The ID of the backup channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetBackupChannelRequest{
								Name: name,
							}

							resp, err := client.GetBackupChannel(ctx, req)
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
						Usage: "update backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-channel", Usage: "The ID of the backup channel.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup_channel.name" not yet supported.
							backup_channel_name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-channel"))
							fmt.Printf("Executing update on %s\n", backup_channel_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-channel", Usage: "The ID of the backup channel.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If provided, this value must match the current value of the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any BackupPlanAssociations below this.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-channel"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackupChannel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.DeleteBackupChannelRequest{
								Name:  name,
								Etag:  cmd.String("etag"),
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteBackupChannel(ctx, req)
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
				Name:  "backup-plan-bindings",
				Usage: "Manage backup-plan-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list backup-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-channel", Usage: "The ID of the backup channel.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListBackupPlanBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupPlanBindings(ctx, req)
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
						Usage: "describe backup-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-channel", Usage: "The ID of the backup channel.", Required: true},
							&cli.StringFlag{Name: "backup-plan-binding", Usage: "The ID of the backup plan binding.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s/backupPlanBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-channel"), cmd.String("backup-plan-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetBackupPlanBindingRequest{
								Name: name,
							}

							resp, err := client.GetBackupPlanBinding(ctx, req)
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
				Name:  "backup-plans",
				Usage: "Manage backup-plans resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-id", Usage: "The client-provided short name for the BackupPlan resource.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.CreateBackupPlanRequest{
								Parent:       parent,
								BackupPlanId: cmd.String("backup-plan-id"),
							}

							op, err := client.CreateBackupPlan(ctx, req)
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
						Usage: "list backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListBackupPlansRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupPlans(ctx, req)
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
						Usage: "describe backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetBackupPlanRequest{
								Name: name,
							}

							resp, err := client.GetBackupPlan(ctx, req)
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
						Usage: "update backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup_plan.name" not yet supported.
							backup_plan_name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							fmt.Printf("Executing update on %s\n", backup_plan_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If provided, this value must match the current value of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackupPlan %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.DeleteBackupPlanRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteBackupPlan(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "backups",
				Usage: "Manage backups resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-id", Usage: "The client-provided short name for the Backup resource.", Required: false},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.CreateBackupRequest{
								Parent:   parent,
								BackupId: cmd.String("backup-id"),
							}

							op, err := client.CreateBackup(ctx, req)
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
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If set to true, the response will return partial results when.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListBackupsRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								Filter:               cmd.String("filter"),
								OrderBy:              cmd.String("order-by"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackups(ctx, req)
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
						Usage: "describe backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetBackupRequest{
								Name: name,
							}

							resp, err := client.GetBackup(ctx, req)
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
						Usage: "update backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup.name" not yet supported.
							backup_name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("backup"))
							fmt.Printf("Executing update on %s\n", backup_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If provided, this value must match the current value of the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any VolumeBackups below this Backup will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("backup"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.DeleteBackupRequest{
								Name:  name,
								Etag:  cmd.String("etag"),
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteBackup(ctx, req)
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
						Name:  "describe",
						Usage: "describe backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							backup := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("backup"))
							fmt.Printf("Executing describe on %s\n", backup)
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
				Name:  "restore-channels",
				Usage: "Manage restore-channels resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-channel-id", Usage: "The client-provided short name for the RestoreChannel resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.CreateRestoreChannelRequest{
								Parent:           parent,
								RestoreChannelId: cmd.String("restore-channel-id"),
							}

							op, err := client.CreateRestoreChannel(ctx, req)
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
						Usage: "list restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListRestoreChannelsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRestoreChannels(ctx, req)
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
						Usage: "describe restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-channel", Usage: "The ID of the restore channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetRestoreChannelRequest{
								Name: name,
							}

							resp, err := client.GetRestoreChannel(ctx, req)
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
						Usage: "update restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-channel", Usage: "The ID of the restore channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "restore_channel.name" not yet supported.
							restore_channel_name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-channel"))
							fmt.Printf("Executing update on %s\n", restore_channel_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If provided, this value must match the current value of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-channel", Usage: "The ID of the restore channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-channel"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRestoreChannel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.DeleteRestoreChannelRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteRestoreChannel(ctx, req)
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
				Name:  "restore-plan-bindings",
				Usage: "Manage restore-plan-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list restore-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-channel", Usage: "The ID of the restore channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListRestorePlanBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRestorePlanBindings(ctx, req)
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
						Usage: "describe restore-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-channel", Usage: "The ID of the restore channel.", Required: true},
							&cli.StringFlag{Name: "restore-plan-binding", Usage: "The ID of the restore plan binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s/restorePlanBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-channel"), cmd.String("restore-plan-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetRestorePlanBindingRequest{
								Name: name,
							}

							resp, err := client.GetRestorePlanBinding(ctx, req)
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
				Name:  "restore-plans",
				Usage: "Manage restore-plans resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-plan-id", Usage: "The client-provided short name for the RestorePlan resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.CreateRestorePlanRequest{
								Parent:        parent,
								RestorePlanId: cmd.String("restore-plan-id"),
							}

							op, err := client.CreateRestorePlan(ctx, req)
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
						Usage: "list restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListRestorePlansRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRestorePlans(ctx, req)
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
						Usage: "describe restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetRestorePlanRequest{
								Name: name,
							}

							resp, err := client.GetRestorePlan(ctx, req)
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
						Usage: "update restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "restore_plan.name" not yet supported.
							restore_plan_name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"))
							fmt.Printf("Executing update on %s\n", restore_plan_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If provided, this value must match the current value of the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any Restores below this RestorePlan will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRestorePlan %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.DeleteRestorePlanRequest{
								Name:  name,
								Etag:  cmd.String("etag"),
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteRestorePlan(ctx, req)
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
				Name:  "restores",
				Usage: "Manage restores resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-id", Usage: "The client-provided short name for the Restore resource.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.CreateRestoreRequest{
								Parent:    parent,
								RestoreId: cmd.String("restore-id"),
							}

							op, err := client.CreateRestore(ctx, req)
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
						Usage: "list restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListRestoresRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRestores(ctx, req)
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
						Usage: "describe restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The ID of the restore.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"), cmd.String("restore"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetRestoreRequest{
								Name: name,
							}

							resp, err := client.GetRestore(ctx, req)
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
						Usage: "update restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The ID of the restore.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "restore.name" not yet supported.
							restore_name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"), cmd.String("restore"))
							fmt.Printf("Executing update on %s\n", restore_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If provided, this value must match the current value of the.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any VolumeRestores below this restore will also.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The ID of the restore.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"), cmd.String("restore"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRestore %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.DeleteRestoreRequest{
								Name:  name,
								Etag:  cmd.String("etag"),
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteRestore(ctx, req)
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
				Name:  "volume-backups",
				Usage: "Manage volume-backups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list volume-backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListVolumeBackupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListVolumeBackups(ctx, req)
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
						Usage: "describe volume-backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume-backup", Usage: "The ID of the volume backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s/volumeBackups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("backup"), cmd.String("volume-backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetVolumeBackupRequest{
								Name: name,
							}

							resp, err := client.GetVolumeBackup(ctx, req)
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
				Name:  "volume-restores",
				Usage: "Manage volume-restores resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list volume-restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The ID of the restore.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"), cmd.String("restore"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.ListVolumeRestoresRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListVolumeRestores(ctx, req)
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
						Usage: "describe volume-restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The ID of the restore.", Required: true},
							&cli.StringFlag{Name: "restore-plan", Usage: "The ID of the restore plan.", Required: true},
							&cli.StringFlag{Name: "volume-restore", Usage: "The ID of the volume restore.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s/volumeRestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore-plan"), cmd.String("restore"), cmd.String("volume-restore"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkebackup.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkebackuppb.GetVolumeRestoreRequest{
								Name: name,
							}

							resp, err := client.GetVolumeRestore(ctx, req)
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
