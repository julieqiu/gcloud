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

package biglake

import (
	biglake "cloud.google.com/go/biglake/apiv1"
	"cloud.google.com/go/biglake/apiv1/biglakepb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the biglake command tree for inclusion under the gcloud root.
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.GetIcebergCatalogRequest{
								Name: name,
							}

							resp, err := client.GetIcebergCatalog(ctx, req)
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
						Usage: "list catalogs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of catalogs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous `ListIcebergCatalogs`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view of the catalog to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.ListIcebergCatalogsRequest{
								Parent:    parent,
								View:      biglakepb.ListIcebergCatalogsRequest_CatalogView(biglakepb.ListIcebergCatalogsRequest_CatalogView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.ListIcebergCatalogs(ctx, req)
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
						Usage: "delete catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIcebergCatalog on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.DeleteIcebergCatalogRequest{
								Name: name,
							}

							if err := client.DeleteIcebergCatalog(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "iceberg_catalog.name" not yet supported.
							iceberg_catalog_name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing update on %s\n", iceberg_catalog_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "iceberg-catalog-id", Usage: "The name of the catalog.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.CreateIcebergCatalogRequest{
								Parent:           parent,
								IcebergCatalogId: cmd.String("iceberg-catalog-id"),
							}

							resp, err := client.CreateIcebergCatalog(ctx, req)
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
						Name:  "failover",
						Usage: "failover catalogs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "primary-replica", Usage: "The region being assigned as the new primary replica region.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, only validate the request, but do not perform the update.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.FailoverIcebergCatalogRequest{
								Name:           name,
								PrimaryReplica: cmd.String("primary-replica"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							resp, err := client.FailoverIcebergCatalog(ctx, req)
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
				Name:  "config",
				Usage: "Manage config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "warehouse", Usage: "Warehouse location or identifier to request from the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.GetIcebergCatalogConfigRequest{
								Warehouse: cmd.String("warehouse"),
							}

							resp, err := client.GetIcebergCatalogConfig(ctx, req)
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
				Name:  "credentials",
				Usage: "Manage credentials resources",
				Commands: []*cli.Command{

					{
						Name:  "load-iceberg-table-credentials",
						Usage: "load-iceberg-table-credentials credentials",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshots", Usage: "What snapshot to get.", Required: false},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.GetIcebergTableRequest{
								Name:      name,
								Snapshots: cmd.String("snapshots"),
							}

							resp, err := client.LoadIcebergTableCredentials(ctx, req)
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
				Name:  "namespaces",
				Usage: "Manage namespaces resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api-parent", Usage: "The parent from the resource path.", Required: true},
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "For servers that support pagination, this signals an upper bound.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "PageToken.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							api_parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							fmt.Printf("Executing list on %s\n", api_parent)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.GetIcebergNamespaceRequest{
								Name: name,
							}

							resp, err := client.GetIcebergNamespace(ctx, req)
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
						Usage: "create namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.CreateIcebergNamespaceRequest{
								Parent: parent,
							}

							resp, err := client.CreateIcebergNamespace(ctx, req)
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
						Usage: "delete namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIcebergNamespace on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.DeleteIcebergNamespaceRequest{
								Name: name,
							}

							if err := client.DeleteIcebergNamespace(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.UpdateIcebergNamespaceRequest{
								Name: name,
							}

							resp, err := client.UpdateIcebergNamespace(ctx, req)
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
				Name:  "register",
				Usage: "Manage register resources",
				Commands: []*cli.Command{

					{
						Name:  "register-iceberg-table",
						Usage: "register-iceberg-table register",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "metadata-location", Usage: "The metadata location of the table.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "overwrite", Usage: "Whether to overwrite the table if it already exists.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.RegisterIcebergTableRequest{
								Parent:           parent,
								MetadataLocation: cmd.String("metadata-location"),
								Overwrite:        cmd.String("overwrite"),
							}

							resp, err := client.RegisterIcebergTable(ctx, req)
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
				Name:  "tables",
				Usage: "Manage tables resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Page size for pagination.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "PageToken for pagination.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.ListIcebergTableIdentifiersRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							resp, err := client.ListIcebergTableIdentifiers(ctx, req)
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
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.CreateIcebergTableRequest{
								Parent: parent,
							}

							resp, err := client.CreateIcebergTable(ctx, req)
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
						Usage: "delete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "purge-requested", Usage: "If true we'll delete both the table and the data.", Required: false},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIcebergTable on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.DeleteIcebergTableRequest{
								Name:           name,
								PurgeRequested: cmd.Bool("purge-requested"),
							}

							if err := client.DeleteIcebergTable(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshots", Usage: "What snapshot to get.", Required: false},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.GetIcebergTableRequest{
								Name:      name,
								Snapshots: cmd.String("snapshots"),
							}

							resp, err := client.GetIcebergTable(ctx, req)
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
						Name:  "update-iceberg-table",
						Usage: "update-iceberg-table tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "catalog", Usage: "The ID of the catalog.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The ID of the table.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := biglake.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &biglakepb.UpdateIcebergTableRequest{
								Name: name,
							}

							resp, err := client.UpdateIcebergTable(ctx, req)
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
