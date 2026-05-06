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

package netapp

import (
	netapp "cloud.google.com/go/netapp/apiv1"
	"cloud.google.com/go/netapp/apiv1/netapppb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the netapp command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "netapp",
		Usage: "manage NetApp API resources",
		Commands: []*cli.Command{
			{
				Name:  "active-directories",
				Usage: "Manage active-directories resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list active-directories",
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
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListActiveDirectoriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListActiveDirectories(ctx, req)
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
						Usage: "describe active-directories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "active-directorie", Usage: "The ID of the active directorie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", cmd.String("project"), cmd.String("location"), cmd.String("active-directorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetActiveDirectoryRequest{
								Name: name,
							}

							resp, err := client.GetActiveDirectory(ctx, req)
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
						Usage: "create active-directories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "active-directory-id", Usage: "ID of the active directory to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateActiveDirectoryRequest{
								Parent:            parent,
								ActiveDirectoryId: cmd.String("active-directory-id"),
							}

							op, err := client.CreateActiveDirectory(ctx, req)
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
						Usage: "update active-directories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "active-directorie", Usage: "The ID of the active directorie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "active_directory.name" not yet supported.
							active_directory_name := fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", cmd.String("project"), cmd.String("location"), cmd.String("active-directorie"))
							fmt.Printf("Executing update on %s\n", active_directory_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete active-directories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "active-directorie", Usage: "The ID of the active directorie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", cmd.String("project"), cmd.String("location"), cmd.String("active-directorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteActiveDirectory %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteActiveDirectoryRequest{
								Name: name,
							}

							op, err := client.DeleteActiveDirectory(ctx, req)
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
				Name:  "backup-policies",
				Usage: "Manage backup-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create backup-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-policy-id", Usage: "The ID to use for the backup policy.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateBackupPolicyRequest{
								Parent:         parent,
								BackupPolicyId: cmd.String("backup-policy-id"),
							}

							op, err := client.CreateBackupPolicy(ctx, req)
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
						Usage: "describe backup-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-policie", Usage: "The ID of the backup policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetBackupPolicyRequest{
								Name: name,
							}

							resp, err := client.GetBackupPolicy(ctx, req)
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
						Usage: "list backup-policies",
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
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListBackupPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackupPolicies(ctx, req)
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
						Usage: "update backup-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-policie", Usage: "The ID of the backup policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup_policy.name" not yet supported.
							backup_policy_name := fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-policie"))
							fmt.Printf("Executing update on %s\n", backup_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backup-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-policie", Usage: "The ID of the backup policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackupPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteBackupPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteBackupPolicy(ctx, req)
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
							&cli.StringFlag{Name: "backup-vault-id", Usage: "The ID to use for the backupVault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateBackupVaultRequest{
								Parent:        parent,
								BackupVaultId: cmd.String("backup-vault-id"),
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
						Name:  "describe",
						Usage: "describe backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetBackupVaultRequest{
								Name: name,
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
						Name:  "list",
						Usage: "list backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value to use if there are additional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListBackupVaultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
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
						Name:  "update",
						Usage: "update backup-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteBackupVaultRequest{
								Name: name,
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
						Name:  "create",
						Usage: "create backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-id", Usage: "The ID to use for the backup.", Required: true},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateBackupRequest{
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
						Name:  "describe",
						Usage: "describe backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetBackupRequest{
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
						Name:  "list",
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The standard list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value to use if there are additional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListBackupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
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
						Name:  "delete",
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("backup"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteBackupRequest{
								Name: name,
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
						Name:  "update",
						Usage: "update backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "backup-vault", Usage: "The ID of the backup vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup.name" not yet supported.
							backup_name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup-vault"), cmd.String("backup"))
							fmt.Printf("Executing update on %s\n", backup_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "host-groups",
				Usage: "Manage host-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list host-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter to apply to the request.", Required: false},
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
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListHostGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListHostGroups(ctx, req)
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
						Usage: "describe host-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host-group", Usage: "The ID of the host group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("host-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetHostGroupRequest{
								Name: name,
							}

							resp, err := client.GetHostGroup(ctx, req)
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
						Usage: "create host-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host-group-id", Usage: "ID of the host group to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateHostGroupRequest{
								Parent:      parent,
								HostGroupId: cmd.String("host-group-id"),
							}

							op, err := client.CreateHostGroup(ctx, req)
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
						Usage: "update host-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host-group", Usage: "The ID of the host group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "host_group.name" not yet supported.
							host_group_name := fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("host-group"))
							fmt.Printf("Executing update on %s\n", host_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete host-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host-group", Usage: "The ID of the host group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("host-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteHostGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteHostGroupRequest{
								Name: name,
							}

							op, err := client.DeleteHostGroup(ctx, req)
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
				Name:  "kms-configs",
				Usage: "Manage kms-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value to use if there are additional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListKmsConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListKmsConfigs(ctx, req)
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
						Usage: "create kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-config-id", Usage: "Id of the requesting KmsConfig.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateKmsConfigRequest{
								Parent:      parent,
								KmsConfigId: cmd.String("kms-config-id"),
							}

							op, err := client.CreateKmsConfig(ctx, req)
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
						Usage: "describe kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-config", Usage: "The ID of the kms config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetKmsConfigRequest{
								Name: name,
							}

							resp, err := client.GetKmsConfig(ctx, req)
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
						Usage: "update kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-config", Usage: "The ID of the kms config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "kms_config.name" not yet supported.
							kms_config_name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms-config"))
							fmt.Printf("Executing update on %s\n", kms_config_name)
							return nil
						},
					},

					{
						Name:  "encrypt",
						Usage: "encrypt kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-config", Usage: "The ID of the kms config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.EncryptVolumesRequest{
								Name: name,
							}

							op, err := client.EncryptVolumes(ctx, req)
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
						Name:  "verify",
						Usage: "verify kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-config", Usage: "The ID of the kms config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.VerifyKmsConfigRequest{
								Name: name,
							}

							resp, err := client.VerifyKmsConfig(ctx, req)
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
						Usage: "delete kms-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-config", Usage: "The ID of the kms config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteKmsConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteKmsConfigRequest{
								Name: name,
							}

							op, err := client.DeleteKmsConfig(ctx, req)
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
				Name:  "ontap",
				Usage: "Manage ontap resources",
				Commands: []*cli.Command{

					{
						Name:  "execute-ontap-post",
						Usage: "execute-ontap-post ontap",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "ontap", Usage: "The ID of the ontap.", Required: true},
							&cli.StringFlag{Name: "ontap-path", Usage: "The resource path of the ONTAP resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							ontap_path := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s/ontap/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"), cmd.String("ontap"))
							fmt.Printf("Executing execute-ontap-post on %s\n", ontap_path)
							return nil
						},
					},

					{
						Name:  "execute-ontap-get",
						Usage: "execute-ontap-get ontap",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "ontap", Usage: "The ID of the ontap.", Required: true},
							&cli.StringFlag{Name: "ontap-path", Usage: "The resource path of the ONTAP resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							ontap_path := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s/ontap/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"), cmd.String("ontap"))
							fmt.Printf("Executing execute-ontap-get on %s\n", ontap_path)
							return nil
						},
					},

					{
						Name:  "execute-ontap-delete",
						Usage: "execute-ontap-delete ontap",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "ontap", Usage: "The ID of the ontap.", Required: true},
							&cli.StringFlag{Name: "ontap-path", Usage: "The resource path of the ONTAP resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							ontap_path := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s/ontap/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"), cmd.String("ontap"))
							fmt.Printf("Executing execute-ontap-delete on %s\n", ontap_path)
							return nil
						},
					},

					{
						Name:  "execute-ontap-patch",
						Usage: "execute-ontap-patch ontap",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "ontap", Usage: "The ID of the ontap.", Required: true},
							&cli.StringFlag{Name: "ontap-path", Usage: "The resource path of the ONTAP resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							ontap_path := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s/ontap/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"), cmd.String("ontap"))
							fmt.Printf("Executing execute-ontap-patch on %s\n", ontap_path)
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
				Name:  "quota-rules",
				Usage: "Manage quota-rules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list quota-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListQuotaRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListQuotaRules(ctx, req)
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
						Usage: "describe quota-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-rule", Usage: "The ID of the quota rule.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("quota-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetQuotaRuleRequest{
								Name: name,
							}

							resp, err := client.GetQuotaRule(ctx, req)
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
						Usage: "create quota-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-rule-id", Usage: "ID of the quota rule to create.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateQuotaRuleRequest{
								Parent:      parent,
								QuotaRuleId: cmd.String("quota-rule-id"),
							}

							op, err := client.CreateQuotaRule(ctx, req)
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
						Usage: "update quota-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-rule", Usage: "The ID of the quota rule.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "quota_rule.name" not yet supported.
							quota_rule_name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("quota-rule"))
							fmt.Printf("Executing update on %s\n", quota_rule_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete quota-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-rule", Usage: "The ID of the quota rule.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("quota-rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteQuotaRule %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteQuotaRuleRequest{
								Name: name,
							}

							op, err := client.DeleteQuotaRule(ctx, req)
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
				Name:  "replications",
				Usage: "Manage replications resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value to use if there are additional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListReplicationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListReplications(ctx, req)
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
						Usage: "describe replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetReplicationRequest{
								Name: name,
							}

							resp, err := client.GetReplication(ctx, req)
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
						Usage: "create replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication-id", Usage: "ID of the replication to create.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateReplicationRequest{
								Parent:        parent,
								ReplicationId: cmd.String("replication-id"),
							}

							op, err := client.CreateReplication(ctx, req)
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
						Usage: "delete replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteReplication %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteReplicationRequest{
								Name: name,
							}

							op, err := client.DeleteReplication(ctx, req)
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
						Name:  "update",
						Usage: "update replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "replication.name" not yet supported.
							replication_name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							fmt.Printf("Executing update on %s\n", replication_name)
							return nil
						},
					},

					{
						Name:  "stop",
						Usage: "stop replications",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "Indicates whether to stop replication forcefully while data transfer is in.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.StopReplicationRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.StopReplication(ctx, req)
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
						Name:  "resume",
						Usage: "resume replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ResumeReplicationRequest{
								Name: name,
							}

							op, err := client.ResumeReplication(ctx, req)
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
						Name:  "reverse-direction",
						Usage: "reverse-direction replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ReverseReplicationDirectionRequest{
								Name: name,
							}

							op, err := client.ReverseReplicationDirection(ctx, req)
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
						Name:  "establish-peering",
						Usage: "establish-peering replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "peer-cluster-name", Usage: "Name of the user's local source cluster to be peered with the.", Required: true},
							&cli.StringSliceFlag{Name: "peer-ip-addresses", Usage: "List of IPv4 ip addresses to be used for peering.", Required: false},
							&cli.StringFlag{Name: "peer-svm-name", Usage: "Name of the user's local source vserver svm to be peered with the.", Required: true},
							&cli.StringFlag{Name: "peer-volume-name", Usage: "Name of the user's local source volume to be peered with the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.EstablishPeeringRequest{
								Name:            name,
								PeerClusterName: cmd.String("peer-cluster-name"),
								PeerSvmName:     cmd.String("peer-svm-name"),
								PeerIpAddresses: cmd.StringSlice("peer-ip-addresses"),
								PeerVolumeName:  cmd.String("peer-volume-name"),
							}

							op, err := client.EstablishPeering(ctx, req)
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
						Name:  "sync",
						Usage: "sync replications",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "replication", Usage: "The ID of the replication.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.SyncReplicationRequest{
								Name: name,
							}

							op, err := client.SyncReplication(ctx, req)
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
				Name:  "snapshots",
				Usage: "Manage snapshots resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value to use if there are additional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListSnapshotsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListSnapshots(ctx, req)
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
						Usage: "describe snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetSnapshotRequest{
								Name: name,
							}

							resp, err := client.GetSnapshot(ctx, req)
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
						Usage: "create snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot-id", Usage: "ID of the snapshot to create.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateSnapshotRequest{
								Parent:     parent,
								SnapshotId: cmd.String("snapshot-id"),
							}

							op, err := client.CreateSnapshot(ctx, req)
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
						Usage: "delete snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSnapshot %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteSnapshotRequest{
								Name: name,
							}

							op, err := client.DeleteSnapshot(ctx, req)
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
						Name:  "update",
						Usage: "update snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "snapshot.name" not yet supported.
							snapshot_name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
							fmt.Printf("Executing update on %s\n", snapshot_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "storage-pools",
				Usage: "Manage storage-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "List filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value to use if there are additional.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListStoragePoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListStoragePools(ctx, req)
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
						Usage: "create storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool-id", Usage: "Id of the requesting storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateStoragePoolRequest{
								Parent:        parent,
								StoragePoolId: cmd.String("storage-pool-id"),
							}

							op, err := client.CreateStoragePool(ctx, req)
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
						Usage: "describe storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetStoragePoolRequest{
								Name: name,
							}

							resp, err := client.GetStoragePool(ctx, req)
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
						Usage: "update storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "storage_pool.name" not yet supported.
							storage_pool_name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"))
							fmt.Printf("Executing update on %s\n", storage_pool_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteStoragePool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteStoragePoolRequest{
								Name: name,
							}

							op, err := client.DeleteStoragePool(ctx, req)
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
						Name:  "validate-directory-service",
						Usage: "validate-directory-service storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "directory-service-type", Usage: "Type of directory service policy attached to the storage pool.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("ValidateDirectoryService %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ValidateDirectoryServiceRequest{
								Name:                 name,
								DirectoryServiceType: netapppb.DirectoryServiceType(netapppb.DirectoryServiceType_value[cmd.String("directory-service-type")]),
							}

							op, err := client.ValidateDirectoryService(ctx, req)
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
						Name:  "switch",
						Usage: "switch storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage-pool", Usage: "The ID of the storage pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.SwitchActiveReplicaZoneRequest{
								Name: name,
							}

							op, err := client.SwitchActiveReplicaZone(ctx, req)
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
				Name:  "volumes",
				Usage: "Manage volumes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list volumes",
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
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.ListVolumesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListVolumes(ctx, req)
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
						Usage: "describe volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.GetVolumeRequest{
								Name: name,
							}

							resp, err := client.GetVolume(ctx, req)
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
						Usage: "create volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume-id", Usage: "Id of the requesting volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.CreateVolumeRequest{
								Parent:   parent,
								VolumeId: cmd.String("volume-id"),
							}

							op, err := client.CreateVolume(ctx, req)
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
						Usage: "update volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "volume.name" not yet supported.
							volume_name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							fmt.Printf("Executing update on %s\n", volume_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete volumes",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If this field is set as true, CCFE will not block the volume resource.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteVolume %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.DeleteVolumeRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteVolume(ctx, req)
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
						Name:  "revert",
						Usage: "revert volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot-id", Usage: "The snapshot resource ID, in the format 'my-snapshot', where the.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.RevertVolumeRequest{
								Name:       name,
								SnapshotId: cmd.String("snapshot-id"),
							}

							op, err := client.RevertVolume(ctx, req)
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
						Name:  "establish-peering",
						Usage: "establish-peering volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "peer-cluster-name", Usage: "Name of the user's local source cluster to be peered with the.", Required: true},
							&cli.StringSliceFlag{Name: "peer-ip-addresses", Usage: "List of IPv4 ip addresses to be used for peering.", Required: false},
							&cli.StringFlag{Name: "peer-svm-name", Usage: "Name of the user's local source vserver svm to be peered with the.", Required: true},
							&cli.StringFlag{Name: "peer-volume-name", Usage: "Name of the user's local source volume to be peered with the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.EstablishVolumePeeringRequest{
								Name:            name,
								PeerClusterName: cmd.String("peer-cluster-name"),
								PeerSvmName:     cmd.String("peer-svm-name"),
								PeerIpAddresses: cmd.StringSlice("peer-ip-addresses"),
								PeerVolumeName:  cmd.String("peer-volume-name"),
							}

							op, err := client.EstablishVolumePeering(ctx, req)
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
						Usage: "restore volumes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The backup resource name, in the format.", Required: true},
							&cli.StringSliceFlag{Name: "file-list", Usage: "List of files to be restored, specified by their absolute path in.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "restore-destination-path", Usage: "Absolute directory path in the destination volume.", Required: false},
							&cli.StringFlag{Name: "volume", Usage: "The ID of the volume.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := netapp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &netapppb.RestoreBackupFilesRequest{
								Name:                   name,
								Backup:                 cmd.String("backup"),
								FileList:               cmd.StringSlice("file-list"),
								RestoreDestinationPath: cmd.String("restore-destination-path"),
							}

							op, err := client.RestoreBackupFiles(ctx, req)
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
