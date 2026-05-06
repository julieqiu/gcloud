package biglake

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud biglake command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "biglake",
		Usage: "manage BigLake API resources",
		Commands: []*cli.Command{
			{
				Name:  "catalogs",
				Usage: "Manage catalogs resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list catalogs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create catalogs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "failover",
						Usage: "failover catalogs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing failover...")
							return nil
						},
					},
				},
			},
			{
				Name:  "config",
				Usage: "Manage config resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe config",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
				},
			},
			{
				Name:  "credentials",
				Usage: "Manage credentials resources",
				Commands: []*cli.Command{
					{
						Name:  "load-iceberg-table-credentials",
						Usage: "load-iceberg-table-credentials credentials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							fmt.Printf("Executing load-iceberg-table-credentials on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "namespaces",
				Usage: "Manage namespaces resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list namespaces",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create namespaces",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "properties",
				Usage: "Manage properties resources",
				Commands: []*cli.Command{
					{
						Name:  "update",
						Usage: "update properties",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "register",
				Usage: "Manage register resources",
				Commands: []*cli.Command{
					{
						Name:  "register-iceberg-table",
						Usage: "register-iceberg-table register",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing register-iceberg-table on %s\n", parent)
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
						Name:  "list",
						Usage: "list tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update-iceberg-table",
						Usage: "update-iceberg-table tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							fmt.Printf("Executing update-iceberg-table on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
