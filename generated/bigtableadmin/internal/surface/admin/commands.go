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

package admin

import (
	admin "cloud.google.com/go/admin/apiv2"
	"cloud.google.com/go/admin/apiv2/adminpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the bigtableadmin command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigtableadmin",
		Usage: "manage Cloud Bigtable Admin API resources",
		Commands: []*cli.Command{
			{
				Name:  "app-profiles",
				Usage: "Manage app-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app-profile-id", Usage: "The ID to be used when referring to the new app profile within.", Required: true},
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If true, ignore safety checks when creating the app profile.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateAppProfileRequest{
								Parent:         parent,
								AppProfileId:   cmd.String("app-profile-id"),
								IgnoreWarnings: cmd.Bool("ignore-warnings"),
							}

							resp, err := client.CreateAppProfile(ctx, req)
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
						Usage: "describe app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app-profile", Usage: "The ID of the app profile.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("app-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetAppProfileRequest{
								Name: name,
							}

							resp, err := client.GetAppProfile(ctx, req)
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
						Usage: "list app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of `next_page_token` returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListAppProfilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAppProfiles(ctx, req)
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
						Usage: "update app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app-profile", Usage: "The ID of the app profile.", Required: true},
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If true, ignore safety checks when updating the app profile.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "app_profile.name" not yet supported.
							app_profile_name := fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("app-profile"))
							fmt.Printf("Executing update on %s\n", app_profile_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app-profile", Usage: "The ID of the app profile.", Required: true},
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If true, ignore safety checks when deleting the app profile.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("app-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAppProfile on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteAppProfileRequest{
								Name:           name,
								IgnoreWarnings: cmd.Bool("ignore-warnings"),
							}

							if err := client.DeleteAppProfile(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "authorized-views",
				Usage: "Manage authorized-views resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorized-view-id", Usage: "The id of the AuthorizedView to create.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateAuthorizedViewRequest{
								Parent:           parent,
								AuthorizedViewId: cmd.String("authorized-view-id"),
							}

							op, err := client.CreateAuthorizedView(ctx, req)
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
						Usage: "list authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of `next_page_token` returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The resource_view to be applied to the returned AuthorizedViews'.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListAuthorizedViewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								View:      adminpb.AuthorizedView_ResponseView(adminpb.AuthorizedView_ResponseView_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListAuthorizedViews(ctx, req)
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
						Usage: "describe authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorized-view", Usage: "The ID of the authorized view.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The resource_view to be applied to the returned AuthorizedView's.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("authorized-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetAuthorizedViewRequest{
								Name: name,
								View: adminpb.AuthorizedView_ResponseView(adminpb.AuthorizedView_ResponseView_value[cmd.String("view")]),
							}

							resp, err := client.GetAuthorizedView(ctx, req)
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
						Usage: "update authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorized-view", Usage: "The ID of the authorized view.", Required: true},
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If true, ignore the safety checks when updating the.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "authorized_view.name" not yet supported.
							authorized_view_name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("authorized-view"))
							fmt.Printf("Executing update on %s\n", authorized_view_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorized-view", Usage: "The ID of the authorized view.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the AuthorizedView.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("authorized-view"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAuthorizedView on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteAuthorizedViewRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteAuthorizedView(ctx, req); err != nil {
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
							&cli.StringFlag{Name: "backup-id", Usage: "The id of the backup to be created.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateBackupRequest{
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
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("backup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetBackupRequest{
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
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backup.name" not yet supported.
							backup_name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("backup"))
							fmt.Printf("Executing update on %s\n", backup_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup", Usage: "The ID of the backup.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("backup"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteBackup on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteBackupRequest{
								Name: name,
							}

							if err := client.DeleteBackup(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters backups listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for specifying the sort order of the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Number of backups to be returned in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If non-empty, `page_token` should contain a.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListBackupsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
						Name:  "copy",
						Usage: "copy backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backup-id", Usage: "The id of the new backup.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-backup", Usage: "The source backup to be copied from.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CopyBackupRequest{
								Parent:       parent,
								BackupId:     cmd.String("backup-id"),
								SourceBackup: cmd.String("source-backup"),
							}

							op, err := client.CopyBackup(ctx, req)
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
				Name:  "clusters",
				Usage: "Manage clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster-id", Usage: "The ID to be used when referring to the new cluster within its.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateClusterRequest{
								Parent:    parent,
								ClusterId: cmd.String("cluster-id"),
							}

							op, err := client.CreateCluster(ctx, req)
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
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetClusterRequest{
								Name: name,
							}

							resp, err := client.GetCluster(ctx, req)
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
						Usage: "list clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "page-token", Usage: "DEPRECATED: This field is unused and ignored.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListClustersRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.ListClusters(ctx, req)
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
						Usage: "update clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "default-storage-type", Usage: "The type of storage used by this cluster to serve its.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location where this cluster's nodes and storage reside.", Required: false},
							&cli.StringFlag{Name: "node-scaling-factor", Usage: "The node scaling factor of this cluster.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "serve-nodes", Usage: "The number of nodes in the cluster.", Required: false},
							&cli.StringFlag{Name: "state", Usage: "The current state of the cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.Cluster{
								Name:               name,
								Location:           cmd.String("location"),
								State:              adminpb.Cluster_State(adminpb.Cluster_State_value[cmd.String("state")]),
								ServeNodes:         int32(cmd.Int("serve-nodes")),
								NodeScalingFactor:  adminpb.Cluster_NodeScalingFactor(adminpb.Cluster_NodeScalingFactor_value[cmd.String("node-scaling-factor")]),
								DefaultStorageType: adminpb.StorageType(adminpb.StorageType_value[cmd.String("default-storage-type")]),
							}

							op, err := client.UpdateCluster(ctx, req)
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
						Name:  "partial-update-cluster",
						Usage: "partial-update-cluster clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "cluster.name" not yet supported.
							cluster_name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							fmt.Printf("Executing partial-update-cluster on %s\n", cluster_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCluster on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteClusterRequest{
								Name: name,
							}

							if err := client.DeleteCluster(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "hot-tablets",
				Usage: "Manage hot-tablets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list hot-tablets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of `next_page_token` returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListHotTabletsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListHotTablets(ctx, req)
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
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance-id", Usage: "The ID to be used when referring to the new instance within its.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateInstanceRequest{
								Parent:     parent,
								InstanceId: cmd.String("instance-id"),
							}

							op, err := client.CreateInstance(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetInstanceRequest{
								Name: name,
							}

							resp, err := client.GetInstance(ctx, req)
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
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "page-token", Usage: "DEPRECATED: This field is unused and ignored.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListInstancesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.ListInstances(ctx, req)
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
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "display-name", Usage: "The descriptive name for this instance as it appears in UIs.", Required: true},
							&cli.StringFlag{Name: "edition", Usage: "The edition of the instance.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "satisfies-pzi", Usage: "Reserved for future use.", Required: false},
							&cli.BoolFlag{Name: "satisfies-pzs", Usage: "Reserved for future use.", Required: false},
							&cli.StringFlag{Name: "state", Usage: "The current state of the instance.", Required: false},
							&cli.StringFlag{Name: "type", Usage: "The type of the instance.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.Instance{
								Name:         name,
								DisplayName:  cmd.String("display-name"),
								State:        adminpb.Instance_State(adminpb.Instance_State_value[cmd.String("state")]),
								Type:         adminpb.Instance_Type(adminpb.Instance_Type_value[cmd.String("type")]),
								Edition:      adminpb.Instance_Edition(adminpb.Instance_Edition_value[cmd.String("edition")]),
								SatisfiesPzs: runtime.Ptr(cmd.Bool("satisfies-pzs")),
								SatisfiesPzi: runtime.Ptr(cmd.Bool("satisfies-pzi")),
							}

							resp, err := client.UpdateInstance(ctx, req)
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
						Name:  "partial-update-instance",
						Usage: "partial-update-instance instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "instance.name" not yet supported.
							instance_name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing partial-update-instance on %s\n", instance_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteInstance on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteInstanceRequest{
								Name: name,
							}

							if err := client.DeleteInstance(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Usage: "set-iam-policy instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
				Name:  "logical-views",
				Usage: "Manage logical-views resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "logical-view-id", Usage: "The ID to use for the logical view, which will become the final.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateLogicalViewRequest{
								Parent:        parent,
								LogicalViewId: cmd.String("logical-view-id"),
							}

							op, err := client.CreateLogicalView(ctx, req)
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
						Usage: "describe logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "logical-view", Usage: "The ID of the logical view.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("logical-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetLogicalViewRequest{
								Name: name,
							}

							resp, err := client.GetLogicalView(ctx, req)
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
						Usage: "list logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of logical views to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListLogicalViews` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListLogicalViewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListLogicalViews(ctx, req)
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
						Usage: "update logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "logical-view", Usage: "The ID of the logical view.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "logical_view.name" not yet supported.
							logical_view_name := fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("logical-view"))
							fmt.Printf("Executing update on %s\n", logical_view_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the logical view.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "logical-view", Usage: "The ID of the logical view.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("logical-view"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteLogicalView on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteLogicalViewRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteLogicalView(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "materialized-views",
				Usage: "Manage materialized-views resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "materialized-view-id", Usage: "The ID to use for the materialized view, which will become the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateMaterializedViewRequest{
								Parent:             parent,
								MaterializedViewId: cmd.String("materialized-view-id"),
							}

							op, err := client.CreateMaterializedView(ctx, req)
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
						Usage: "describe materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "materialized-view", Usage: "The ID of the materialized view.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("materialized-view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetMaterializedViewRequest{
								Name: name,
							}

							resp, err := client.GetMaterializedView(ctx, req)
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
						Usage: "list materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of materialized views to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListMaterializedViews`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListMaterializedViewsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMaterializedViews(ctx, req)
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
						Usage: "update materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "materialized-view", Usage: "The ID of the materialized view.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "materialized_view.name" not yet supported.
							materialized_view_name := fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("materialized-view"))
							fmt.Printf("Executing update on %s\n", materialized_view_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the materialized view.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "materialized-view", Usage: "The ID of the materialized view.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("materialized-view"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteMaterializedView on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteMaterializedViewRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteMaterializedView(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "schema-bundles",
				Usage: "Manage schema-bundles resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-bundle-id", Usage: "The unique ID to use for the schema bundle, which will become the.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateSchemaBundleRequest{
								Parent:         parent,
								SchemaBundleId: cmd.String("schema-bundle-id"),
							}

							op, err := client.CreateSchemaBundle(ctx, req)
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
						Usage: "update schema-bundles",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If set, ignore the safety checks when updating the Schema Bundle.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-bundle", Usage: "The ID of the schema bundle.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "schema_bundle.name" not yet supported.
							schema_bundle_name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("schema-bundle"))
							fmt.Printf("Executing update on %s\n", schema_bundle_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-bundle", Usage: "The ID of the schema bundle.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("schema-bundle"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetSchemaBundleRequest{
								Name: name,
							}

							resp, err := client.GetSchemaBundle(ctx, req)
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
						Usage: "list schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of schema bundles to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSchemaBundles` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListSchemaBundlesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSchemaBundles(ctx, req)
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
						Usage: "delete schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag of the schema bundle.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-bundle", Usage: "The ID of the schema bundle.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("schema-bundle"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSchemaBundle on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteSchemaBundleRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteSchemaBundle(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "describe",
						Usage: "describe snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/snapshots/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("snapshot"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetSnapshotRequest{
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
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of snapshots to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of `next_page_token` returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListSnapshotsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
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
						Name:  "delete",
						Usage: "delete snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The ID of the cluster.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/snapshots/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("snapshot"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSnapshot on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteSnapshotRequest{
								Name: name,
							}

							if err := client.DeleteSnapshot(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "tables",
				Usage: "Manage tables resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table-id", Usage: "The name by which the new table should be referred to within the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateTableRequest{
								Parent:  parent,
								TableId: cmd.String("table-id"),
							}

							resp, err := client.CreateTable(ctx, req)
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
						Usage: "create tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "source-snapshot", Usage: "The unique name of the snapshot from which to restore the table.", Required: true},
							&cli.StringFlag{Name: "table-id", Usage: "The name by which the new table should be referred to within the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CreateTableFromSnapshotRequest{
								Parent:         parent,
								TableId:        cmd.String("table-id"),
								SourceSnapshot: cmd.String("source-snapshot"),
							}

							op, err := client.CreateTableFromSnapshot(ctx, req)
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
						Usage: "list tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of `next_page_token` returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view to be applied to the returned tables' fields.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ListTablesRequest{
								Parent:    parent,
								View:      adminpb.Table_View(adminpb.Table_View_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTables(ctx, req)
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
						Usage: "describe tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view to be applied to the returned table's fields.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetTableRequest{
								Name: name,
								View: adminpb.Table_View(adminpb.Table_View_value[cmd.String("view")]),
							}

							resp, err := client.GetTable(ctx, req)
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
						Usage: "update tables",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If true, ignore safety checks when updating the table.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "table.name" not yet supported.
							table_name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing update on %s\n", table_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTable on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DeleteTableRequest{
								Name: name,
							}

							if err := client.DeleteTable(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "undelete",
						Usage: "undelete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.UndeleteTableRequest{
								Name: name,
							}

							op, err := client.UndeleteTable(ctx, req)
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
						Name:  "modify-column-families",
						Usage: "modify-column-families tables",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "ignore-warnings", Usage: "If true, ignore safety checks when modifying the column families.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.ModifyColumnFamiliesRequest{
								Name:           name,
								IgnoreWarnings: cmd.Bool("ignore-warnings"),
							}

							resp, err := client.ModifyColumnFamilies(ctx, req)
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
						Name:  "drop-row-range",
						Usage: "drop-row-range tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DropRowRange on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.DropRowRangeRequest{
								Name: name,
							}

							if err := client.DropRowRange(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "generate-consistency-token",
						Usage: "generate-consistency-token tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GenerateConsistencyTokenRequest{
								Name: name,
							}

							resp, err := client.GenerateConsistencyToken(ctx, req)
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
						Name:  "check-consistency",
						Usage: "check-consistency tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "consistency-token", Usage: "The token created using GenerateConsistencyToken for the Table.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.CheckConsistencyRequest{
								Name:             name,
								ConsistencyToken: cmd.String("consistency-token"),
							}

							resp, err := client.CheckConsistency(ctx, req)
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
						Name:  "snapshot",
						Usage: "snapshot tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cluster", Usage: "The name of the cluster where the snapshot will be created in.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "Description of the snapshot.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot-id", Usage: "The ID by which the new snapshot should be referred to within the.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.SnapshotTableRequest{
								Name:        name,
								Cluster:     cmd.String("cluster"),
								SnapshotId:  cmd.String("snapshot-id"),
								Description: cmd.String("description"),
							}

							op, err := client.SnapshotTable(ctx, req)
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
						Usage: "restore tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table-id", Usage: "The id of the table to create and restore to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.RestoreTableRequest{
								Parent:  parent,
								TableId: cmd.String("table-id"),
							}

							op, err := client.RestoreTable(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Usage: "set-iam-policy tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := admin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &adminpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
