package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	clouddms "cloud.google.com/go/clouddms/apiv1"
	"cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
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
				Name:  "datamigration",
				Usage: "manage Database Migration API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "connection-profiles",
						Usage: "Manage connection-profiles resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list connection-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &clouddmspb.ListConnectionProfilesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListConnectionProfiles(ctx, req)
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
								Usage: "describe connection-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection_profile", Usage: "The connection_profile.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection_profile"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.GetConnectionProfileRequest{Name: name}
									resp, err := client.GetConnectionProfile(ctx, req)
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
								Usage: "create connection-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection-profile-id", Usage: "The connection profile id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.BoolFlag{Name: "skip-validation", Usage: "The skip validation.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.CreateConnectionProfileRequest{Parent: parent}
									req.ConnectionProfileId = cmd.String("connection-profile-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.SkipValidation = cmd.Bool("skip-validation")
									req.ConnectionProfile = &clouddmspb.ConnectionProfile{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreateConnectionProfile(ctx, req)
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
								Usage: "update connection-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection_profile", Usage: "The connection_profile.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection_profile"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.UpdateConnectionProfileRequest{}
									req.ConnectionProfile = &clouddmspb.ConnectionProfile{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateConnectionProfile(ctx, req)
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
								Usage: "delete connection-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "connection_profile", Usage: "The connection_profile.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/connectionProfiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection_profile"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.DeleteConnectionProfileRequest{Name: name}
									op, err := client.DeleteConnectionProfile(ctx, req)
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
								Name:  "set-iam-policy",
								Usage: "set-iam-policy connection-profiles",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy connection-profiles",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "conversion-workspaces",
						Usage: "Manage conversion-workspaces resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.GetConversionWorkspaceRequest{Name: name}
									resp, err := client.GetConversionWorkspace(ctx, req)
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
								Usage: "list conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &clouddmspb.ListConversionWorkspacesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListConversionWorkspaces(ctx, req)
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
								Usage: "create conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion-workspace-id", Usage: "The conversion workspace id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.CreateConversionWorkspaceRequest{Parent: parent}
									req.ConversionWorkspaceId = cmd.String("conversion-workspace-id")
									req.ConversionWorkspace = &clouddmspb.ConversionWorkspace{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreateConversionWorkspace(ctx, req)
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
								Usage: "update conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.UpdateConversionWorkspaceRequest{}
									req.ConversionWorkspace = &clouddmspb.ConversionWorkspace{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateConversionWorkspace(ctx, req)
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
								Usage: "delete conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.DeleteConversionWorkspaceRequest{Name: name}
									op, err := client.DeleteConversionWorkspace(ctx, req)
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
								Name:  "seed",
								Usage: "seed conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.BoolFlag{Name: "auto-commit", Usage: "The auto commit.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.SeedConversionWorkspaceRequest{Name: name}
									req.AutoCommit = cmd.Bool("auto-commit")
									op, err := client.SeedConversionWorkspace(ctx, req)
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
								Name:  "convert",
								Usage: "convert conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.BoolFlag{Name: "auto-commit", Usage: "The auto commit.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
									&cli.BoolFlag{Name: "convert-full-path", Usage: "The convert full path.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.ConvertConversionWorkspaceRequest{Name: name}
									req.AutoCommit = cmd.Bool("auto-commit")
									req.Filter = cmd.String("filter")
									req.ConvertFullPath = cmd.Bool("convert-full-path")
									op, err := client.ConvertConversionWorkspace(ctx, req)
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
								Name:  "commit",
								Usage: "commit conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.StringFlag{Name: "commit-name", Usage: "The commit name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.CommitConversionWorkspaceRequest{Name: name}
									req.CommitName = cmd.String("commit-name")
									op, err := client.CommitConversionWorkspace(ctx, req)
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
								Name:  "rollback",
								Usage: "rollback conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.RollbackConversionWorkspaceRequest{Name: name}
									op, err := client.RollbackConversionWorkspace(ctx, req)
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
								Name:  "apply",
								Usage: "apply conversion-workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
									&cli.BoolFlag{Name: "dry-run", Usage: "The dry run.", Required: false},
									&cli.BoolFlag{Name: "auto-commit", Usage: "The auto commit.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.ApplyConversionWorkspaceRequest{Name: name}
									req.Filter = cmd.String("filter")
									req.DryRun = cmd.Bool("dry-run")
									req.AutoCommit = cmd.Bool("auto-commit")
									op, err := client.ApplyConversionWorkspace(ctx, req)
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
								Name:  "describe-database-entities",
								Usage: "describe-database-entities conversion-workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe-database-entities...")
									return nil
								},
							},
							{
								Name:  "search-background-jobs",
								Usage: "search-background-jobs conversion-workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search-background-jobs...")
									return nil
								},
							},
							{
								Name:  "describe-conversion-workspace-revisions",
								Usage: "describe-conversion-workspace-revisions conversion-workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe-conversion-workspace-revisions...")
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
								Name:  "fetch-static-ips",
								Usage: "fetch-static-ips locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-static-ips...")
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
									client, err := clouddms.NewDataMigrationClient(ctx)
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
									client, err := clouddms.NewDataMigrationClient(ctx)
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
						Name:  "mapping-rules",
						Usage: "Manage mapping-rules resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create mapping-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.StringFlag{Name: "mapping-rule-id", Usage: "The mapping rule id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.IntFlag{Name: "rule-order", Usage: "The rule order.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.CreateMappingRuleRequest{Parent: parent}
									req.MappingRuleId = cmd.String("mapping-rule-id")
									req.MappingRule = &clouddmspb.MappingRule{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										RuleOrder:   int64(cmd.Int("rule-order")),
									}
									resp, err := client.CreateMappingRule(ctx, req)
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
								Usage: "delete mapping-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.StringFlag{Name: "mapping_rule", Usage: "The mapping_rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s/mappingRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"), cmd.String("mapping_rule"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.DeleteMappingRuleRequest{Name: name}
									if err := client.DeleteMappingRule(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list mapping-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &clouddmspb.ListMappingRulesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMappingRules(ctx, req)
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
								Usage: "describe mapping-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "conversion_workspace", Usage: "The conversion_workspace.", Required: true},
									&cli.StringFlag{Name: "mapping_rule", Usage: "The mapping_rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/conversionWorkspaces/%s/mappingRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversion_workspace"), cmd.String("mapping_rule"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.GetMappingRuleRequest{Name: name}
									resp, err := client.GetMappingRule(ctx, req)
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
								Name:  "import",
								Usage: "import mapping-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "migration-jobs",
						Usage: "Manage migration-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &clouddmspb.ListMigrationJobsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMigrationJobs(ctx, req)
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
								Usage: "describe migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.GetMigrationJobRequest{Name: name}
									resp, err := client.GetMigrationJob(ctx, req)
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
								Usage: "create migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration-job-id", Usage: "The migration job id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "dump-path", Usage: "The dump path.", Required: false},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "destination", Usage: "The destination.", Required: true},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
									&cli.StringFlag{Name: "cmek-key-name", Usage: "The cmek key name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.CreateMigrationJobRequest{Parent: parent}
									req.MigrationJobId = cmd.String("migration-job-id")
									req.MigrationJob = &clouddmspb.MigrationJob{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										DumpPath:    cmd.String("dump-path"),
										Source:      cmd.String("source"),
										Destination: cmd.String("destination"),
										Filter:      cmd.String("filter"),
										CmekKeyName: cmd.String("cmek-key-name"),
									}
									op, err := client.CreateMigrationJob(ctx, req)
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
								Usage: "update migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "dump-path", Usage: "The dump path.", Required: false},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: false},
									&cli.StringFlag{Name: "destination", Usage: "The destination.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
									&cli.StringFlag{Name: "cmek-key-name", Usage: "The cmek key name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.UpdateMigrationJobRequest{}
									req.MigrationJob = &clouddmspb.MigrationJob{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										DumpPath:    cmd.String("dump-path"),
										Source:      cmd.String("source"),
										Destination: cmd.String("destination"),
										Filter:      cmd.String("filter"),
										CmekKeyName: cmd.String("cmek-key-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("dump-path") {
										paths = append(paths, "dump_path")
									}
									if cmd.IsSet("source") {
										paths = append(paths, "source")
									}
									if cmd.IsSet("destination") {
										paths = append(paths, "destination")
									}
									if cmd.IsSet("filter") {
										paths = append(paths, "filter")
									}
									if cmd.IsSet("cmek-key-name") {
										paths = append(paths, "cmek_key_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMigrationJob(ctx, req)
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
								Usage: "delete migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.DeleteMigrationJobRequest{Name: name}
									op, err := client.DeleteMigrationJob(ctx, req)
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
								Name:  "start",
								Usage: "start migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
									&cli.BoolFlag{Name: "skip-validation", Usage: "The skip validation.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.StartMigrationJobRequest{Name: name}
									req.SkipValidation = cmd.Bool("skip-validation")
									op, err := client.StartMigrationJob(ctx, req)
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
								Name:  "stop",
								Usage: "stop migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.StopMigrationJobRequest{Name: name}
									op, err := client.StopMigrationJob(ctx, req)
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
								Name:  "resume",
								Usage: "resume migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.ResumeMigrationJobRequest{Name: name}
									op, err := client.ResumeMigrationJob(ctx, req)
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
								Name:  "promote",
								Usage: "promote migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.PromoteMigrationJobRequest{Name: name}
									op, err := client.PromoteMigrationJob(ctx, req)
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
								Name:  "verify",
								Usage: "verify migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.VerifyMigrationJobRequest{Name: name}
									op, err := client.VerifyMigrationJob(ctx, req)
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
								Name:  "restart",
								Usage: "restart migration-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "migration_job", Usage: "The migration_job.", Required: true},
									&cli.BoolFlag{Name: "skip-validation", Usage: "The skip validation.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/migrationJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("migration_job"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.RestartMigrationJobRequest{Name: name}
									req.SkipValidation = cmd.Bool("skip-validation")
									op, err := client.RestartMigrationJob(ctx, req)
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
								Name:  "generate-ssh-script",
								Usage: "generate-ssh-script migration-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-ssh-script...")
									return nil
								},
							},
							{
								Name:  "generate-tcp-proxy-script",
								Usage: "generate-tcp-proxy-script migration-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-tcp-proxy-script...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions migration-jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
									client, err := clouddms.NewDataMigrationClient(ctx)
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
									client, err := clouddms.NewDataMigrationClient(ctx)
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
									client, err := clouddms.NewDataMigrationClient(ctx)
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
									client, err := clouddms.NewDataMigrationClient(ctx)
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
						Name:  "private-connections",
						Usage: "Manage private-connections resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create private-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "private-connection-id", Usage: "The private connection id.", Required: true},
									&cli.BoolFlag{Name: "skip-validation", Usage: "The skip validation.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.CreatePrivateConnectionRequest{Parent: parent}
									req.PrivateConnectionId = cmd.String("private-connection-id")
									req.SkipValidation = cmd.Bool("skip-validation")
									req.PrivateConnection = &clouddmspb.PrivateConnection{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreatePrivateConnection(ctx, req)
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
								Name:  "describe",
								Usage: "describe private-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "private_connection", Usage: "The private_connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_connection"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.GetPrivateConnectionRequest{Name: name}
									resp, err := client.GetPrivateConnection(ctx, req)
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
								Usage: "list private-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &clouddmspb.ListPrivateConnectionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPrivateConnections(ctx, req)
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
								Name:  "delete",
								Usage: "delete private-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "private_connection", Usage: "The private_connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("private_connection"))
									client, err := clouddms.NewDataMigrationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &clouddmspb.DeletePrivateConnectionRequest{Name: name}
									op, err := client.DeletePrivateConnection(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
