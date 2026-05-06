package bigtableadmin

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud bigtableadmin command tree.
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "app_profile", Usage: "The app_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("app_profile"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list app-profiles",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "app_profile", Usage: "The app_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("app_profile"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete app-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "app_profile", Usage: "The app_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("app_profile"))
							fmt.Printf("Executing delete on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
							&cli.StringFlag{Name: "authorized_view", Usage: "The authorized_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("authorized_view"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
							&cli.StringFlag{Name: "authorized_view", Usage: "The authorized_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("authorized_view"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete authorized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
							&cli.StringFlag{Name: "authorized_view", Usage: "The authorized_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("authorized_view"))
							fmt.Printf("Executing delete on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("backup"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("backup"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("backup"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "copy",
						Usage: "copy backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing copy on %s\n", parent)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update clusters",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "partial-update-cluster",
						Usage: "partial-update-cluster clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							fmt.Printf("Executing partial-update-cluster on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"))
							fmt.Printf("Executing delete on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", parent)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing update...")
							return nil
						},
					},
					{
						Name:  "partial-update-instance",
						Usage: "partial-update-instance instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing partial-update-instance on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "logical_view", Usage: "The logical_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("logical_view"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "logical_view", Usage: "The logical_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("logical_view"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete logical-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "logical_view", Usage: "The logical_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("logical_view"))
							fmt.Printf("Executing delete on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "materialized_view", Usage: "The materialized_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("materialized_view"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "materialized_view", Usage: "The materialized_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("materialized_view"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete materialized-views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "materialized_view", Usage: "The materialized_view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", cmd.String("project"), cmd.String("instance"), cmd.String("materialized_view"))
							fmt.Printf("Executing delete on %s\n", name)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
					{
						Name:  "cancel",
						Usage: "cancel operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing cancel...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
					{
						Name:  "cancel",
						Usage: "cancel operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing cancel...")
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
							&cli.StringFlag{Name: "schema_bundle", Usage: "The schema_bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("schema_bundle"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
							&cli.StringFlag{Name: "schema_bundle", Usage: "The schema_bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("schema_bundle"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete schema-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
							&cli.StringFlag{Name: "schema_bundle", Usage: "The schema_bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s/schemaBundles/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"), cmd.String("schema_bundle"))
							fmt.Printf("Executing delete on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/snapshots/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("snapshot"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/clusters/%s/snapshots/%s", cmd.String("project"), cmd.String("instance"), cmd.String("cluster"), cmd.String("snapshot"))
							fmt.Printf("Executing delete on %s\n", name)
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
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/instances/%s", cmd.String("project"), cmd.String("instance"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tables",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list tables",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "undelete",
						Usage: "undelete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing undelete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "modify-column-families",
						Usage: "modify-column-families tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing modify-column-families on %s\n", name)
							return nil
						},
					},
					{
						Name:  "drop-row-range",
						Usage: "drop-row-range tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing drop-row-range on %s\n", name)
							return nil
						},
					},
					{
						Name:  "generate-consistency-token",
						Usage: "generate-consistency-token tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing generate-consistency-token on %s\n", name)
							return nil
						},
					},
					{
						Name:  "check-consistency",
						Usage: "check-consistency tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing check-consistency on %s\n", name)
							return nil
						},
					},
					{
						Name:  "snapshot",
						Usage: "snapshot tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing snapshot on %s\n", name)
							return nil
						},
					},
					{
						Name:  "restore",
						Usage: "restore tables",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing restore on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/tables/%s", cmd.String("project"), cmd.String("instance"), cmd.String("table"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
