package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	metastore "cloud.google.com/go/metastore/apiv1"
	"cloud.google.com/go/metastore/apiv1/metastorepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "metastore",
				Usage: "manage Dataproc Metastore API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "backups",
						Usage: "Manage backups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &metastorepb.ListBackupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListBackups(ctx, req)
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
								Name:  "describe",
								Usage: "describe backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("backup"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.GetBackupRequest{Name: name}
									resp, err := client.GetBackup(ctx, req)
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
								Usage: "create backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "backup-id", Usage: "The backup id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.CreateBackupRequest{Parent: parent}
									req.BackupId = cmd.String("backup-id")
									req.Backup = &metastorepb.Backup{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateBackup(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "delete backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("backup"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.DeleteBackupRequest{Name: name}
									op, err := client.DeleteBackup(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "federations",
						Usage: "Manage federations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list federations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &metastorepb.ListFederationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFederations(ctx, req)
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
								Name:  "describe",
								Usage: "describe federations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "federation", Usage: "The federation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/federations/%s", cmd.String("project"), cmd.String("location"), cmd.String("federation"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.GetFederationRequest{Name: name}
									resp, err := client.GetFederation(ctx, req)
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
								Usage: "create federations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "federation-id", Usage: "The federation id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.CreateFederationRequest{Parent: parent}
									req.FederationId = cmd.String("federation-id")
									req.Federation = &metastorepb.Federation{
										Name:    cmd.String("name"),
										Version: cmd.String("version"),
									}
									op, err := client.CreateFederation(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "update",
								Usage: "update federations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "federation", Usage: "The federation.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/federations/%s", cmd.String("project"), cmd.String("location"), cmd.String("federation"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.UpdateFederationRequest{}
									req.Federation = &metastorepb.Federation{
										Name:    name,
										Name:    cmd.String("name"),
										Version: cmd.String("version"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("version") {
										paths = append(paths, "version")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateFederation(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "delete federations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "federation", Usage: "The federation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/federations/%s", cmd.String("project"), cmd.String("location"), cmd.String("federation"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.DeleteFederationRequest{Name: name}
									op, err := client.DeleteFederation(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
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
								Name:  "describe",
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
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
								Usage: "list locations",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
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
								Name:  "describe",
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
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
						Name:  "metadata-imports",
						Usage: "Manage metadata-imports resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list metadata-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &metastorepb.ListMetadataImportsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMetadataImports(ctx, req)
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
								Name:  "describe",
								Usage: "describe metadata-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "metadata_import", Usage: "The metadata_import.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s/metadataImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("metadata_import"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.GetMetadataImportRequest{Name: name}
									resp, err := client.GetMetadataImport(ctx, req)
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
								Usage: "create metadata-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "metadata-import-id", Usage: "The metadata import id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.CreateMetadataImportRequest{Parent: parent}
									req.MetadataImportId = cmd.String("metadata-import-id")
									req.MetadataImport = &metastorepb.MetadataImport{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateMetadataImport(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "update",
								Usage: "update metadata-imports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "metadata_import", Usage: "The metadata_import.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s/metadataImports/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("metadata_import"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.UpdateMetadataImportRequest{}
									req.MetadataImport = &metastorepb.MetadataImport{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMetadataImport(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
						Name:  "operations",
						Usage: "Manage operations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
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
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
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
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := metastore.NewDataprocMetastoreFederationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "services",
						Usage: "Manage services resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &metastorepb.ListServicesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListServices(ctx, req)
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
								Name:  "describe",
								Usage: "describe services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.GetServiceRequest{Name: name}
									resp, err := client.GetService(ctx, req)
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
								Usage: "create services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service-id", Usage: "The service id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.IntFlag{Name: "port", Usage: "The port.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.CreateServiceRequest{Parent: parent}
									req.ServiceId = cmd.String("service-id")
									req.Service = &metastorepb.Service{
										Name:    cmd.String("name"),
										Network: cmd.String("network"),
										Port:    int32(cmd.Int("port")),
									}
									op, err := client.CreateService(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "update",
								Usage: "update services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.IntFlag{Name: "port", Usage: "The port.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.UpdateServiceRequest{}
									req.Service = &metastorepb.Service{
										Name:    name,
										Name:    cmd.String("name"),
										Network: cmd.String("network"),
										Port:    int32(cmd.Int("port")),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									if cmd.IsSet("port") {
										paths = append(paths, "port")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateService(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "delete services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
									client, err := metastore.NewDataprocMetastoreClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &metastorepb.DeleteServiceRequest{Name: name}
									op, err := client.DeleteService(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "export-metadata",
								Usage: "export-metadata services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing export-metadata...")
									return nil
								},
							},
							{
								Name:  "restore",
								Usage: "restore services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing restore...")
									return nil
								},
							},
							{
								Name:  "query-metadata",
								Usage: "query-metadata services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-metadata...")
									return nil
								},
							},
							{
								Name:  "move-table-to-database",
								Usage: "move-table-to-database services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move-table-to-database...")
									return nil
								},
							},
							{
								Name:  "alter-location",
								Usage: "alter-location services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing alter-location...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
