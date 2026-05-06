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

package backupdr

import (
	backupdr "cloud.google.com/go/backupdr/apiv1"
	"cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the backupdr command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "backupdr",
		Usage: "manage Backup and DR Service API resources",
		Commands: []*cli.Command{
			{
				Name:  "backup-plan-associations",
				Usage: "Manage backup-plan-associations resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-association-id", Usage: "The name of the backup plan association to create.", Required: true},
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.CreateBackupPlanAssociationRequest{
								Parent:                  parent,
								BackupPlanAssociationId: cmd.String("backup-plan-association-id"),
								RequestId:               cmd.String("request-id"),
							}

							op, err := client.CreateBackupPlanAssociation(ctx, req)
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
						Usage: "update backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-association", Usage: "The ID of the backup plan association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup_plan_association.name" not yet supported.
							backup_plan_association_name := fmt.Sprintf("projects/%s/locations/%s/backupPlanAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan-association"))
							fmt.Printf("Executing update on %s\n", backup_plan_association_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-association", Usage: "The ID of the backup plan association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlanAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan-association"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetBackupPlanAssociationRequest{
								Name: name,
							}

							resp, err := client.GetBackupPlanAssociation(ctx, req)
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
						Usage: "list backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListBackupPlanAssociationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupPlanAssociations(ctx, req)
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
						Name:  "fetch-for-resource-type",
						Usage: "fetch-for-resource-type backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results fetched in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of BackupPlanAssociations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous call of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-type", Usage: "The type of the GCP resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.FetchBackupPlanAssociationsForResourceTypeRequest{
								Parent:       parent,
								ResourceType: cmd.String("resource-type"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								Filter:       cmd.String("filter"),
								OrderBy:      cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.FetchBackupPlanAssociationsForResourceType(ctx, req)
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
						Usage: "delete backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-association", Usage: "The ID of the backup plan association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlanAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan-association"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackupPlanAssociation %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.DeleteBackupPlanAssociationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteBackupPlanAssociation(ctx, req)
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
						Name:  "trigger-backup",
						Usage: "trigger-backup backup-plan-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-association", Usage: "The ID of the backup plan association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "rule-id", Usage: "backup rule_id for which a backup needs to be triggered.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlanAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan-association"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.TriggerBackupRequest{
								Name:      name,
								RuleId:    cmd.String("rule-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.TriggerBackup(ctx, req)
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
				Name:  "backup-plans",
				Usage: "Manage backup-plans resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan-id", Usage: "The name of the `BackupPlan` to create.", Required: true},
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.CreateBackupPlanRequest{
								Parent:       parent,
								BackupPlanId: cmd.String("backup-plan-id"),
								RequestId:    cmd.String("request-id"),
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
						Name:  "update",
						Usage: "update backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup_plan.name" not yet supported.
							backup_plan_name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							fmt.Printf("Executing update on %s\n", backup_plan_name)
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetBackupPlanRequest{
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
						Name:  "list",
						Usage: "list backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Field match expression used to filter the results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field by which to sort the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `BackupPlans` to return in a single.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListBackupPlansRequest{
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
						Name:  "delete",
						Usage: "delete backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.DeleteBackupPlanRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
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
				},
			},
			{
				Name:  "backup-vaults",
				Usage: "Manage backup-vaults resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault-id", Usage: "ID of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.CreateBackupVaultRequest{
								Parent:        parent,
								BackupVaultId: cmd.String("backup-vault-id"),
								RequestId:     cmd.String("request-id"),
								ValidateOnly:  cmd.Bool("validate-only"),
							}

							op, err := client.CreateBackupVault(ctx, req)
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
						Usage: "list backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Reserved for future use to provide a BASIC & FULL view of Backup.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListBackupVaultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								View:      backupdrpb.BackupVaultView(backupdrpb.BackupVaultView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupVaults(ctx, req)
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
						Name:  "fetch-usable",
						Usage: "fetch-usable backup-vaults",
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.FetchUsableBackupVaultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.FetchUsableBackupVaults(ctx, req)
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
						Usage: "describe backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Reserved for future use to provide a BASIC & FULL view of Backup.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetBackupVaultRequest{
								Name: name,
								View: backupdrpb.BackupVaultView(backupdrpb.BackupVaultView_value[cmd.String("view")]),
							}

							resp, err := client.GetBackupVault(ctx, req)
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
						Usage: "update backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, will not check plan duration against backup vault.", Required: false},
							&cli.BoolFlag{Name: "force-update-access-restriction", Usage: "If set to true, we will force update access restriction even if.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup_vault.name" not yet supported.
							backup_vault_name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							fmt.Printf("Executing update on %s\n", backup_vault_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backup-vaults",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If true and the BackupVault is not found, the request will.", Required: false},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the backup vault.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any data source from this backup vault will also.", Required: false},
							&cli.BoolFlag{Name: "ignore-backup-plan-references", Usage: "If set to true, backupvault deletion will proceed even if there.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackupVault %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.DeleteBackupVaultRequest{
								Name:                       name,
								RequestId:                  cmd.String("request-id"),
								Force:                      cmd.Bool("force"),
								Etag:                       cmd.String("etag"),
								ValidateOnly:               cmd.Bool("validate-only"),
								AllowMissing:               cmd.Bool("allow-missing"),
								IgnoreBackupPlanReferences: cmd.Bool("ignore-backup-plan-references"),
							}

							op, err := client.DeleteBackupVault(ctx, req)
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
				Name:  "backups",
				Usage: "Manage backups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Reserved for future use to provide a BASIC & FULL view of Backup.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListBackupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								View:      backupdrpb.BackupView(backupdrpb.BackupView_value[cmd.String("view")]),
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
						Name:  "fetch-for-resource-type",
						Usage: "fetch-for-resource-type backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results fetched in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Backups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous call of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-type", Usage: "The type of the GCP resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "This parameter is used to specify the view of the backup.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.FetchBackupsForResourceTypeRequest{
								Parent:       parent,
								ResourceType: cmd.String("resource-type"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								Filter:       cmd.String("filter"),
								OrderBy:      cmd.String("order-by"),
								View:         backupdrpb.BackupView(backupdrpb.BackupView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.FetchBackupsForResourceType(ctx, req)
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
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Reserved for future use to provide a BASIC & FULL view of Backup.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetBackupRequest{
								Name: name,
								View: backupdrpb.BackupView(backupdrpb.BackupView_value[cmd.String("view")]),
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
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup.name" not yet supported.
							backup_name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"), cmd.String("backup"))
							fmt.Printf("Executing update on %s\n", backup_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.DeleteBackupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteBackup(ctx, req)
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
						Name:  "restore",
						Usage: "restore backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.RestoreBackupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.RestoreBackup(ctx, req)
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
				Name:  "data-source-references",
				Usage: "Manage data-source-references resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe data-source-references",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-source-reference", Usage: "The ID of the data source reference.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataSourceReferences/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-source-reference"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetDataSourceReferenceRequest{
								Name: name,
							}

							resp, err := client.GetDataSourceReference(ctx, req)
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
						Usage: "list data-source-references",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results listed in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of DataSourceReferences to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDataSourceReferences`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListDataSourceReferencesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataSourceReferences(ctx, req)
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
						Name:  "fetch-for-resource-type",
						Usage: "fetch-for-resource-type data-source-references",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters the results fetched in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields to order by, sorted in ascending.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of DataSourceReferences to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous call of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-type", Usage: "The type of the GCP resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.FetchDataSourceReferencesForResourceTypeRequest{
								Parent:       parent,
								ResourceType: cmd.String("resource-type"),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
								Filter:       cmd.String("filter"),
								OrderBy:      cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.FetchDataSourceReferencesForResourceType(ctx, req)
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
				Name:  "data-sources",
				Usage: "Manage data-sources resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListDataSourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataSources(ctx, req)
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
						Usage: "describe data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetDataSourceRequest{
								Name: name,
							}

							resp, err := client.GetDataSource(ctx, req)
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
						Usage: "update data-sources",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "Enable upsert.", Required: false},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_source.name" not yet supported.
							data_source_name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("data-source"))
							fmt.Printf("Executing update on %s\n", data_source_name)
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
				Name:  "management-servers",
				Usage: "Manage management-servers resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list management-servers",
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListManagementServersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    runtime.Ptr(cmd.String("filter")),
								OrderBy:   runtime.Ptr(cmd.String("order-by")),
							}

							limit := cmd.Int("limit")
							it := client.ListManagementServers(ctx, req)
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
						Usage: "describe management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetManagementServerRequest{
								Name: name,
							}

							resp, err := client.GetManagementServer(ctx, req)
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
						Usage: "create management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server-id", Usage: "The name of the management server to create.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.CreateManagementServerRequest{
								Parent:             parent,
								ManagementServerId: cmd.String("management-server-id"),
								RequestId:          cmd.String("request-id"),
							}

							op, err := client.CreateManagementServer(ctx, req)
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
						Usage: "delete management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteManagementServer %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.DeleteManagementServerRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteManagementServer(ctx, req)
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
						Usage: "set-iam-policy management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions management-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "management-server", Usage: "The ID of the management server.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/managementServers/%s", cmd.String("project"), cmd.String("location"), cmd.String("management-server"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
				Name:  "resource-backup-configs",
				Usage: "Manage resource-backup-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list resource-backup-configs",
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
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListResourceBackupConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListResourceBackupConfigs(ctx, req)
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
						Name:  "describe",
						Usage: "describe revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.GetBackupPlanRevisionRequest{
								Name: name,
							}

							resp, err := client.GetBackupPlanRevision(ctx, req)
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
							&cli.StringFlag{Name: "backup-plan", Usage: "The ID of the backup plan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `BackupPlans` to return in a single.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-plan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.ListBackupPlanRevisionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupPlanRevisions(ctx, req)
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
				Name:  "service-config",
				Usage: "Manage service-config resources",
				Commands: []*cli.Command{

					{
						Name:  "initialize",
						Usage: "initialize service-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-type", Usage: "The resource type to which the default service config will be.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serviceConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := backupdr.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &backupdrpb.InitializeServiceRequest{
								Name:         name,
								ResourceType: cmd.String("resource-type"),
								RequestId:    cmd.String("request-id"),
							}

							op, err := client.InitializeService(ctx, req)
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
		},
	}
}
