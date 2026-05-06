package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	biglake "cloud.google.com/go/biglake/apiv1"
	"cloud.google.com/go/biglake/apiv1/biglakepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "biglake",
				Usage: "manage BigLake API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
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
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.GetIcebergCatalogRequest{Name: name}
									resp, err := client.GetIcebergCatalog(ctx, req)
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
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.DeleteIcebergCatalogRequest{Name: name}
									if err := client.DeleteIcebergCatalog(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update catalogs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
									&cli.StringFlag{Name: "default-location", Usage: "The default location.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.UpdateIcebergCatalogRequest{}
									req.IcebergCatalog = &biglakepb.IcebergCatalog{
										Name:            name,
										DefaultLocation: cmd.String("default-location"),
									}
									var paths []string
									if cmd.IsSet("default-location") {
										paths = append(paths, "default_location")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateIcebergCatalog(ctx, req)
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
								Usage: "create catalogs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "iceberg-catalog-id", Usage: "The iceberg catalog id.", Required: true},
									&cli.StringFlag{Name: "default-location", Usage: "The default location.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.CreateIcebergCatalogRequest{Parent: parent}
									req.IcebergCatalogId = cmd.String("iceberg-catalog-id")
									req.IcebergCatalog = &biglakepb.IcebergCatalog{
										DefaultLocation: cmd.String("default-location"),
									}
									resp, err := client.CreateIcebergCatalog(ctx, req)
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
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.GetIcebergNamespaceRequest{Name: name}
									resp, err := client.GetIcebergNamespace(ctx, req)
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
								Usage: "create namespaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.CreateIcebergNamespaceRequest{Parent: parent}
									resp, err := client.CreateIcebergNamespace(ctx, req)
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
								Usage: "delete namespaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
									&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.DeleteIcebergNamespaceRequest{Name: name}
									if err := client.DeleteIcebergNamespace(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.UpdateIcebergNamespaceRequest{}
									resp, err := client.UpdateIcebergNamespace(ctx, req)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &biglakepb.ListIcebergTableIdentifiersRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListIcebergTableIdentifiers(ctx, req)
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
								Name:  "create",
								Usage: "create tables",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/catalogs/%s", cmd.String("project"), cmd.String("catalog"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.CreateIcebergTableRequest{Parent: parent}
									resp, err := client.CreateIcebergTable(ctx, req)
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
								Usage: "delete tables",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "catalog", Usage: "The catalog.", Required: true},
									&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
									&cli.StringFlag{Name: "table", Usage: "The table.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/%s/tables/%s", cmd.String("project"), cmd.String("catalog"), cmd.String("namespace"), cmd.String("table"))
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.DeleteIcebergTableRequest{Name: name}
									if err := client.DeleteIcebergTable(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									client, err := biglake.NewIcebergCatalogClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &biglakepb.GetIcebergTableRequest{Name: name}
									resp, err := client.GetIcebergTable(ctx, req)
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
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
