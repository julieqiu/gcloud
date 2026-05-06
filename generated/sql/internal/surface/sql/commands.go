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

package sql

import (
	sql "cloud.google.com/go/sql/apiv1"
	"cloud.google.com/go/sql/apiv1/sqlpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the sqladmin command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "sqladmin",
		Usage: "manage Cloud SQL Admin API resources",
		Commands: []*cli.Command{
			{
				Name:  "acquire-ssrs-lease",
				Usage: "Manage acquire-ssrs-lease resources",
				Commands: []*cli.Command{

					{
						Name:  "acquire-ssrs-lease",
						Usage: "acquire-ssrs-lease acquire-ssrs-lease",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing acquire-ssrs-lease on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "add-entra-id-certificate",
				Usage: "Manage add-entra-id-certificate resources",
				Commands: []*cli.Command{

					{
						Name:  "add-entra-id-certificate",
						Usage: "add-entra-id-certificate add-entra-id-certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing add-entra-id-certificate on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "add-server-ca",
				Usage: "Manage add-server-ca resources",
				Commands: []*cli.Command{

					{
						Name:  "add-server-ca",
						Usage: "add-server-ca add-server-ca",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing add-server-ca on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "add-server-certificate",
				Usage: "Manage add-server-certificate resources",
				Commands: []*cli.Command{

					{
						Name:  "add-server-certificate",
						Usage: "add-server-certificate add-server-certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing add-server-certificate on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "backup-runs",
				Usage: "Manage backup-runs resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete backup-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "id", Usage: "The ID of the id.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("id"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backup-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "id", Usage: "The ID of the id.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("id"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert backup-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list backup-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "Maximum number of backup runs per response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A previously-returned page token representing part of the larger set of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := sql.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &sqlpb.CreateBackupRequest{
								Parent: parent,
							}

							resp, err := client.CreateBackup(ctx, req)
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/backups/%s", cmd.String("project"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := sql.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &sqlpb.GetBackupRequest{
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
							&cli.StringFlag{Name: "filter", Usage: "Multiple filter queries are separated by spaces.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of backups to return per response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListBackups` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := sql.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &sqlpb.ListBackupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
						Name:  "update",
						Usage: "update backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup.name" not yet supported.
							backup_name := fmt.Sprintf("projects/%s/backups/%s", cmd.String("project"), cmd.String("backup"))
							fmt.Printf("Executing update on %s\n", backup_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/backups/%s", cmd.String("project"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := sql.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &sqlpb.DeleteBackupRequest{
								Name: name,
							}

							resp, err := client.DeleteBackup(ctx, req)
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
				Name:  "cancel",
				Usage: "Manage cancel resources",
				Commands: []*cli.Command{

					{
						Name:  "cancel",
						Usage: "cancel cancel",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "clone",
				Usage: "Manage clone resources",
				Commands: []*cli.Command{

					{
						Name:  "clone",
						Usage: "clone clone",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing clone on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "connect-settings",
				Usage: "Manage connect-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe connect-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "create-ephemeral",
				Usage: "Manage create-ephemeral resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create create-ephemeral",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing create on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "databases",
				Usage: "Manage databases resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "demote",
				Usage: "Manage demote resources",
				Commands: []*cli.Command{

					{
						Name:  "demote",
						Usage: "demote demote",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing demote on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "demote-master",
				Usage: "Manage demote-master resources",
				Commands: []*cli.Command{

					{
						Name:  "demote-master",
						Usage: "demote-master demote-master",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing demote-master on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "execute-sql",
				Usage: "Manage execute-sql resources",
				Commands: []*cli.Command{

					{
						Name:  "execute-sql",
						Usage: "execute-sql execute-sql",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing execute-sql on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "export",
				Usage: "Manage export resources",
				Commands: []*cli.Command{

					{
						Name:  "export",
						Usage: "export export",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing export on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "failover",
				Usage: "Manage failover resources",
				Commands: []*cli.Command{

					{
						Name:  "failover",
						Usage: "failover failover",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing failover on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "flags",
				Usage: "Manage flags resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list flags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database-version", Usage: "Database type and version you want to retrieve flags for.", Required: false},
							&cli.StringFlag{Name: "flag-scope", Usage: "Specify the scope of flags to be returned by SqlFlagsListService.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := sql.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &sqlpb.SqlFlagsListRequest{
								DatabaseVersion: cmd.String("database-version"),
								FlagScope:       runtime.Ptr(sqlpb.SqlFlagScope(sqlpb.SqlFlagScope_value[cmd.String("flag-scope")])),
							}

							resp, err := client.List(ctx, req)
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
				Name:  "get-disk-shrink-config",
				Usage: "Manage get-disk-shrink-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe get-disk-shrink-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "get-latest-recovery-time",
				Usage: "Manage get-latest-recovery-time resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe get-latest-recovery-time",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "import",
				Usage: "Manage import resources",
				Commands: []*cli.Command{

					{
						Name:  "import",
						Usage: "import import",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing import on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "generate-ephemeral-cert",
						Usage: "generate-ephemeral-cert instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-token", Usage: "Access token to include in the signed certificate.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public-key", Usage: "PEM encoded public key to include in the signed certificate.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing generate-ephemeral-cert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "enable-final-backup", Usage: "Flag to opt-in for final backup.", Required: false},
							&cli.StringFlag{Name: "final-backup-description", Usage: "The description of the final backup.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of instances to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A previously-returned page token representing part of the larger set of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "list-entra-id-certificates",
				Usage: "Manage list-entra-id-certificates resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list list-entra-id-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "list-server-cas",
				Usage: "Manage list-server-cas resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list list-server-cas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "list-server-certificates",
				Usage: "Manage list-server-certificates resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list list-server-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
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
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "Cloud SQL instance ID.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "Maximum number of operations per response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A previously-returned page token representing part of the larger set of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "perform-disk-shrink",
				Usage: "Manage perform-disk-shrink resources",
				Commands: []*cli.Command{

					{
						Name:  "perform-disk-shrink",
						Usage: "perform-disk-shrink perform-disk-shrink",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing perform-disk-shrink on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "pre-check-major-version-upgrade",
				Usage: "Manage pre-check-major-version-upgrade resources",
				Commands: []*cli.Command{

					{
						Name:  "pre-check-major-version-upgrade",
						Usage: "pre-check-major-version-upgrade pre-check-major-version-upgrade",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing pre-check-major-version-upgrade on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "projects",
				Usage: "Manage projects resources",
				Commands: []*cli.Command{

					{
						Name:  "point-in-time-restore",
						Usage: "point-in-time-restore projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := sql.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &sqlpb.SqlInstancesPointInTimeRestoreRequest{
								Parent: parent,
							}

							resp, err := client.PointInTimeRestore(ctx, req)
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
				Name:  "promote-replica",
				Usage: "Manage promote-replica resources",
				Commands: []*cli.Command{

					{
						Name:  "promote-replica",
						Usage: "promote-replica promote-replica",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "failover", Usage: "Set to true to invoke a replica failover to the DR.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing promote-replica on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "reencrypt",
				Usage: "Manage reencrypt resources",
				Commands: []*cli.Command{

					{
						Name:  "reencrypt",
						Usage: "reencrypt reencrypt",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing reencrypt on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "release-ssrs-lease",
				Usage: "Manage release-ssrs-lease resources",
				Commands: []*cli.Command{

					{
						Name:  "release-ssrs-lease",
						Usage: "release-ssrs-lease release-ssrs-lease",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing release-ssrs-lease on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "reschedule-maintenance",
				Usage: "Manage reschedule-maintenance resources",
				Commands: []*cli.Command{

					{
						Name:  "reschedule-maintenance",
						Usage: "reschedule-maintenance reschedule-maintenance",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing reschedule-maintenance on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "reset-replica-size",
				Usage: "Manage reset-replica-size resources",
				Commands: []*cli.Command{

					{
						Name:  "reset-replica-size",
						Usage: "reset-replica-size reset-replica-size",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing reset-replica-size on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "reset-ssl-config",
				Usage: "Manage reset-ssl-config resources",
				Commands: []*cli.Command{

					{
						Name:  "reset-ssl-config",
						Usage: "reset-ssl-config reset-ssl-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "mode", Usage: "Reset SSL mode to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing reset-ssl-config on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "restart",
				Usage: "Manage restart resources",
				Commands: []*cli.Command{

					{
						Name:  "restart",
						Usage: "restart restart",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing restart on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "restore-backup",
				Usage: "Manage restore-backup resources",
				Commands: []*cli.Command{

					{
						Name:  "restore-backup",
						Usage: "restore-backup restore-backup",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing restore-backup on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "rotate-entra-id-certificate",
				Usage: "Manage rotate-entra-id-certificate resources",
				Commands: []*cli.Command{

					{
						Name:  "rotate-entra-id-certificate",
						Usage: "rotate-entra-id-certificate rotate-entra-id-certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing rotate-entra-id-certificate on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "rotate-server-ca",
				Usage: "Manage rotate-server-ca resources",
				Commands: []*cli.Command{

					{
						Name:  "rotate-server-ca",
						Usage: "rotate-server-ca rotate-server-ca",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing rotate-server-ca on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "rotate-server-certificate",
				Usage: "Manage rotate-server-certificate resources",
				Commands: []*cli.Command{

					{
						Name:  "rotate-server-certificate",
						Usage: "rotate-server-certificate rotate-server-certificate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing rotate-server-certificate on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "ssl-certs",
				Usage: "Manage ssl-certs resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete ssl-certs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sha1_fingerprint", Usage: "The ID of the sha1_fingerprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("sha1_fingerprint"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ssl-certs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sha1_fingerprint", Usage: "The ID of the sha1_fingerprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("sha1_fingerprint"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert ssl-certs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list ssl-certs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "start-external-sync",
				Usage: "Manage start-external-sync resources",
				Commands: []*cli.Command{

					{
						Name:  "start-external-sync",
						Usage: "start-external-sync start-external-sync",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "migration-type", Usage: "MigrationType configures the migration to use physical files or.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "replica-overwrite-enabled", Usage: "MySQL only.", Required: false},
							&cli.BoolFlag{Name: "skip-verification", Usage: "Whether to skip the verification step (VESS).", Required: false},
							&cli.StringFlag{Name: "sync-mode", Usage: "External sync mode.", Required: false},
							&cli.StringFlag{Name: "sync-parallel-level", Usage: "Parallel level for initial data sync.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing start-external-sync on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "start-replica",
				Usage: "Manage start-replica resources",
				Commands: []*cli.Command{

					{
						Name:  "start-replica",
						Usage: "start-replica start-replica",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing start-replica on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "stop-replica",
				Usage: "Manage stop-replica resources",
				Commands: []*cli.Command{

					{
						Name:  "stop-replica",
						Usage: "stop-replica stop-replica",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing stop-replica on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "switchover",
				Usage: "Manage switchover resources",
				Commands: []*cli.Command{

					{
						Name:  "switchover",
						Usage: "switchover switchover",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing switchover on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "tiers",
				Usage: "Manage tiers resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list tiers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "truncate-log",
				Usage: "Manage truncate-log resources",
				Commands: []*cli.Command{

					{
						Name:  "truncate-log",
						Usage: "truncate-log truncate-log",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing truncate-log on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "users",
				Usage: "Manage users resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host", Usage: "Host of the user in the instance.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "host", Usage: "Host of a user of the instance.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The ID of the name.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("instance"), cmd.String("name"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update users",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "database-roles", Usage: "List of database roles to grant to the user.", Required: false},
							&cli.StringFlag{Name: "host", Usage: "Host of the user in the instance.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "revoke-existing-roles", Usage: "Specifies whether to revoke existing roles that are not present.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "verify-external-sync-settings",
				Usage: "Manage verify-external-sync-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "verify-external-sync-settings",
						Usage: "verify-external-sync-settings verify-external-sync-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "migration-type", Usage: "MigrationType configures the migration to use physical files or.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "sync-mode", Usage: "External sync mode.", Required: false},
							&cli.StringFlag{Name: "sync-parallel-level", Usage: "Parallel level for initial data sync.", Required: false},
							&cli.BoolFlag{Name: "verify-connection-only", Usage: "Flag to enable verifying connection only.", Required: false},
							&cli.BoolFlag{Name: "verify-replication-only", Usage: "Flag to verify settings required by replication setup only.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing verify-external-sync-settings on %s\n", project)
							return nil
						},
					},
				},
			},
		},
	}
}
