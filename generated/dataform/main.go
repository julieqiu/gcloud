package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	dataform "cloud.google.com/go/dataform/apiv1"
	"cloud.google.com/go/dataform/apiv1/dataformpb"
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
				Name:  "dataform",
				Usage: "manage Dataform API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "compilation-results",
						Usage: "Manage compilation-results resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list compilation-results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataformpb.ListCompilationResultsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCompilationResults(ctx, req)
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
								Usage: "describe compilation-results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "compilation_result", Usage: "The compilation_result.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/compilationResults/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("compilation_result"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetCompilationResultRequest{Name: name}
									resp, err := client.GetCompilationResult(ctx, req)
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
								Usage: "create compilation-results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateCompilationResultRequest{Parent: parent}
									resp, err := client.CreateCompilationResult(ctx, req)
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
								Name:  "query",
								Usage: "query compilation-results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "compilation_result", Usage: "The compilation_result.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/compilationResults/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("compilation_result"))
									fmt.Printf("Executing query on %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetConfigRequest{Name: name}
									resp, err := client.GetConfig(ctx, req)
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
								Usage: "update config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "default-kms-key-name", Usage: "The default kms key name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.UpdateConfigRequest{}
									req.Config = &dataformpb.Config{
										Name:              name,
										DefaultKmsKeyName: cmd.String("default-kms-key-name"),
									}
									var paths []string
									if cmd.IsSet("default-kms-key-name") {
										paths = append(paths, "default_kms_key_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateConfig(ctx, req)
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
						Name:  "folders",
						Usage: "Manage folders resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetFolderRequest{Name: name}
									resp, err := client.GetFolder(ctx, req)
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
								Usage: "create folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "containing-folder", Usage: "The containing folder.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateFolderRequest{Parent: parent}
									req.Folder = &dataformpb.Folder{
										DisplayName:      cmd.String("display-name"),
										ContainingFolder: cmd.String("containing-folder"),
									}
									resp, err := client.CreateFolder(ctx, req)
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
								Usage: "update folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "containing-folder", Usage: "The containing folder.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.UpdateFolderRequest{}
									req.Folder = &dataformpb.Folder{
										Name:             name,
										DisplayName:      cmd.String("display-name"),
										ContainingFolder: cmd.String("containing-folder"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("containing-folder") {
										paths = append(paths, "containing_folder")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateFolder(ctx, req)
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
								Usage: "delete folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteFolderRequest{Name: name}
									if err := client.DeleteFolder(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete-tree",
								Usage: "delete-tree folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
									&cli.BoolFlag{Name: "force", Usage: "The force.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteFolderTreeRequest{Name: name}
									req.Force = cmd.Bool("force")
									op, err := client.DeleteFolderTree(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("DeleteFolderTree completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "query-folder-contents",
								Usage: "query-folder-contents folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-folder-contents...")
									return nil
								},
							},
							{
								Name:  "move",
								Usage: "move folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
									&cli.StringFlag{Name: "destination-containing-folder", Usage: "The destination containing folder.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/folders/%s", cmd.String("project"), cmd.String("location"), cmd.String("folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.MoveFolderRequest{Name: name}
									req.DestinationContainingFolder = cmd.String("destination-containing-folder")
									op, err := client.MoveFolder(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("MoveFolder completed on %s\n", name)
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
								Name:  "query-user-root-contents",
								Usage: "query-user-root-contents locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-user-root-contents...")
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
									client, err := dataform.NewClient(ctx)
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
									client, err := dataform.NewClient(ctx)
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
									client, err := dataform.NewClient(ctx)
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
									client, err := dataform.NewClient(ctx)
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
									client, err := dataform.NewClient(ctx)
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
									client, err := dataform.NewClient(ctx)
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
						Name:  "release-configs",
						Usage: "Manage release-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list release-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataformpb.ListReleaseConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListReleaseConfigs(ctx, req)
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
								Usage: "describe release-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "release_config", Usage: "The release_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("release_config"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetReleaseConfigRequest{Name: name}
									resp, err := client.GetReleaseConfig(ctx, req)
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
								Usage: "create release-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "release-config-id", Usage: "The release config id.", Required: true},
									&cli.StringFlag{Name: "git-commitish", Usage: "The git commitish.", Required: true},
									&cli.StringFlag{Name: "cron-schedule", Usage: "The cron schedule.", Required: false},
									&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
									&cli.StringFlag{Name: "release-compilation-result", Usage: "The release compilation result.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateReleaseConfigRequest{Parent: parent}
									req.ReleaseConfigId = cmd.String("release-config-id")
									req.ReleaseConfig = &dataformpb.ReleaseConfig{
										GitCommitish:             cmd.String("git-commitish"),
										CronSchedule:             cmd.String("cron-schedule"),
										TimeZone:                 cmd.String("time-zone"),
										ReleaseCompilationResult: cmd.String("release-compilation-result"),
										Disabled:                 cmd.Bool("disabled"),
									}
									resp, err := client.CreateReleaseConfig(ctx, req)
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
								Usage: "update release-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "release_config", Usage: "The release_config.", Required: true},
									&cli.StringFlag{Name: "git-commitish", Usage: "The git commitish.", Required: false},
									&cli.StringFlag{Name: "cron-schedule", Usage: "The cron schedule.", Required: false},
									&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
									&cli.StringFlag{Name: "release-compilation-result", Usage: "The release compilation result.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("release_config"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.UpdateReleaseConfigRequest{}
									req.ReleaseConfig = &dataformpb.ReleaseConfig{
										Name:                     name,
										GitCommitish:             cmd.String("git-commitish"),
										CronSchedule:             cmd.String("cron-schedule"),
										TimeZone:                 cmd.String("time-zone"),
										ReleaseCompilationResult: cmd.String("release-compilation-result"),
										Disabled:                 cmd.Bool("disabled"),
									}
									var paths []string
									if cmd.IsSet("git-commitish") {
										paths = append(paths, "git_commitish")
									}
									if cmd.IsSet("cron-schedule") {
										paths = append(paths, "cron_schedule")
									}
									if cmd.IsSet("time-zone") {
										paths = append(paths, "time_zone")
									}
									if cmd.IsSet("release-compilation-result") {
										paths = append(paths, "release_compilation_result")
									}
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateReleaseConfig(ctx, req)
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
								Usage: "delete release-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "release_config", Usage: "The release_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("release_config"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteReleaseConfigRequest{Name: name}
									if err := client.DeleteReleaseConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "repositories",
						Usage: "Manage repositories resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetRepositoryRequest{Name: name}
									resp, err := client.GetRepository(ctx, req)
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
								Usage: "create repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository-id", Usage: "The repository id.", Required: true},
									&cli.StringFlag{Name: "containing-folder", Usage: "The containing folder.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "npmrc-environment-variables-secret-version", Usage: "The npmrc environment variables secret version.", Required: false},
									&cli.BoolFlag{Name: "set-authenticated-user-admin", Usage: "The set authenticated user admin.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: false},
									&cli.StringFlag{Name: "kms-key-name", Usage: "The kms key name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateRepositoryRequest{Parent: parent}
									req.RepositoryId = cmd.String("repository-id")
									req.Repository = &dataformpb.Repository{
										ContainingFolder:                       cmd.String("containing-folder"),
										DisplayName:                            cmd.String("display-name"),
										NpmrcEnvironmentVariablesSecretVersion: cmd.String("npmrc-environment-variables-secret-version"),
										SetAuthenticatedUserAdmin:              cmd.Bool("set-authenticated-user-admin"),
										ServiceAccount:                         cmd.String("service-account"),
										KmsKeyName:                             cmd.String("kms-key-name"),
									}
									resp, err := client.CreateRepository(ctx, req)
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
								Usage: "update repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "containing-folder", Usage: "The containing folder.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "npmrc-environment-variables-secret-version", Usage: "The npmrc environment variables secret version.", Required: false},
									&cli.BoolFlag{Name: "set-authenticated-user-admin", Usage: "The set authenticated user admin.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: false},
									&cli.StringFlag{Name: "kms-key-name", Usage: "The kms key name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.UpdateRepositoryRequest{}
									req.Repository = &dataformpb.Repository{
										Name:                                   name,
										ContainingFolder:                       cmd.String("containing-folder"),
										DisplayName:                            cmd.String("display-name"),
										NpmrcEnvironmentVariablesSecretVersion: cmd.String("npmrc-environment-variables-secret-version"),
										SetAuthenticatedUserAdmin:              cmd.Bool("set-authenticated-user-admin"),
										ServiceAccount:                         cmd.String("service-account"),
										KmsKeyName:                             cmd.String("kms-key-name"),
									}
									var paths []string
									if cmd.IsSet("containing-folder") {
										paths = append(paths, "containing_folder")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("npmrc-environment-variables-secret-version") {
										paths = append(paths, "npmrc_environment_variables_secret_version")
									}
									if cmd.IsSet("set-authenticated-user-admin") {
										paths = append(paths, "set_authenticated_user_admin")
									}
									if cmd.IsSet("service-account") {
										paths = append(paths, "service_account")
									}
									if cmd.IsSet("kms-key-name") {
										paths = append(paths, "kms_key_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateRepository(ctx, req)
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
								Usage: "delete repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteRepositoryRequest{Name: name}
									if err := client.DeleteRepository(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "move",
								Usage: "move repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "destination-containing-folder", Usage: "The destination containing folder.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.MoveRepositoryRequest{Name: name}
									req.DestinationContainingFolder = cmd.String("destination-containing-folder")
									op, err := client.MoveRepository(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("MoveRepository completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "commit",
								Usage: "commit repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									fmt.Printf("Executing commit on %s\n", name)
									return nil
								},
							},
							{
								Name:  "read-file",
								Usage: "read-file repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									fmt.Printf("Executing read-file on %s\n", name)
									return nil
								},
							},
							{
								Name:  "query-directory-contents",
								Usage: "query-directory-contents repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									fmt.Printf("Executing query-directory-contents on %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-history",
								Usage: "fetch-history repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									fmt.Printf("Executing fetch-history on %s\n", name)
									return nil
								},
							},
							{
								Name:  "compute-access-token-status",
								Usage: "compute-access-token-status repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									fmt.Printf("Executing compute-access-token-status on %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-remote-branches",
								Usage: "fetch-remote-branches repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									fmt.Printf("Executing fetch-remote-branches on %s\n", name)
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "team-folders",
						Usage: "Manage team-folders resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe team-folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "team_folder", Usage: "The team_folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team_folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetTeamFolderRequest{Name: name}
									resp, err := client.GetTeamFolder(ctx, req)
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
								Usage: "create team-folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateTeamFolderRequest{Parent: parent}
									req.TeamFolder = &dataformpb.TeamFolder{
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateTeamFolder(ctx, req)
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
								Usage: "update team-folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "team_folder", Usage: "The team_folder.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team_folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.UpdateTeamFolderRequest{}
									req.TeamFolder = &dataformpb.TeamFolder{
										Name:        name,
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateTeamFolder(ctx, req)
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
								Usage: "delete team-folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "team_folder", Usage: "The team_folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team_folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteTeamFolderRequest{Name: name}
									if err := client.DeleteTeamFolder(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete-tree",
								Usage: "delete-tree team-folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "team_folder", Usage: "The team_folder.", Required: true},
									&cli.BoolFlag{Name: "force", Usage: "The force.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/teamFolders/%s", cmd.String("project"), cmd.String("location"), cmd.String("team_folder"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteTeamFolderTreeRequest{Name: name}
									req.Force = cmd.Bool("force")
									op, err := client.DeleteTeamFolderTree(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("DeleteTeamFolderTree completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "query-contents",
								Usage: "query-contents team-folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-contents...")
									return nil
								},
							},
							{
								Name:  "search",
								Usage: "search team-folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search...")
									return nil
								},
							},
						},
					},
					{
						Name:  "workflow-configs",
						Usage: "Manage workflow-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list workflow-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataformpb.ListWorkflowConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWorkflowConfigs(ctx, req)
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
								Usage: "describe workflow-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_config", Usage: "The workflow_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_config"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetWorkflowConfigRequest{Name: name}
									resp, err := client.GetWorkflowConfig(ctx, req)
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
								Usage: "create workflow-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow-config-id", Usage: "The workflow config id.", Required: true},
									&cli.StringFlag{Name: "release-config", Usage: "The release config.", Required: true},
									&cli.StringFlag{Name: "cron-schedule", Usage: "The cron schedule.", Required: false},
									&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateWorkflowConfigRequest{Parent: parent}
									req.WorkflowConfigId = cmd.String("workflow-config-id")
									req.WorkflowConfig = &dataformpb.WorkflowConfig{
										ReleaseConfig: cmd.String("release-config"),
										CronSchedule:  cmd.String("cron-schedule"),
										TimeZone:      cmd.String("time-zone"),
										Disabled:      cmd.Bool("disabled"),
									}
									resp, err := client.CreateWorkflowConfig(ctx, req)
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
								Usage: "update workflow-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_config", Usage: "The workflow_config.", Required: true},
									&cli.StringFlag{Name: "release-config", Usage: "The release config.", Required: false},
									&cli.StringFlag{Name: "cron-schedule", Usage: "The cron schedule.", Required: false},
									&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_config"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.UpdateWorkflowConfigRequest{}
									req.WorkflowConfig = &dataformpb.WorkflowConfig{
										Name:          name,
										ReleaseConfig: cmd.String("release-config"),
										CronSchedule:  cmd.String("cron-schedule"),
										TimeZone:      cmd.String("time-zone"),
										Disabled:      cmd.Bool("disabled"),
									}
									var paths []string
									if cmd.IsSet("release-config") {
										paths = append(paths, "release_config")
									}
									if cmd.IsSet("cron-schedule") {
										paths = append(paths, "cron_schedule")
									}
									if cmd.IsSet("time-zone") {
										paths = append(paths, "time_zone")
									}
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateWorkflowConfig(ctx, req)
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
								Usage: "delete workflow-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_config", Usage: "The workflow_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_config"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteWorkflowConfigRequest{Name: name}
									if err := client.DeleteWorkflowConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "workflow-invocations",
						Usage: "Manage workflow-invocations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list workflow-invocations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataformpb.ListWorkflowInvocationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWorkflowInvocations(ctx, req)
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
								Usage: "describe workflow-invocations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_invocation", Usage: "The workflow_invocation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_invocation"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetWorkflowInvocationRequest{Name: name}
									resp, err := client.GetWorkflowInvocation(ctx, req)
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
								Usage: "create workflow-invocations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateWorkflowInvocationRequest{Parent: parent}
									resp, err := client.CreateWorkflowInvocation(ctx, req)
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
								Usage: "delete workflow-invocations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_invocation", Usage: "The workflow_invocation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_invocation"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteWorkflowInvocationRequest{Name: name}
									if err := client.DeleteWorkflowInvocation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel workflow-invocations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_invocation", Usage: "The workflow_invocation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_invocation"))
									fmt.Printf("Executing cancel on %s\n", name)
									return nil
								},
							},
							{
								Name:  "query",
								Usage: "query workflow-invocations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workflow_invocation", Usage: "The workflow_invocation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowInvocations/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workflow_invocation"))
									fmt.Printf("Executing query on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "workspaces",
						Usage: "Manage workspaces resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &dataformpb.ListWorkspacesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWorkspaces(ctx, req)
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
								Usage: "describe workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.GetWorkspaceRequest{Name: name}
									resp, err := client.GetWorkspace(ctx, req)
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
								Usage: "create workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace-id", Usage: "The workspace id.", Required: true},
									&cli.BoolFlag{Name: "disable-moves", Usage: "The disable moves.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.CreateWorkspaceRequest{Parent: parent}
									req.WorkspaceId = cmd.String("workspace-id")
									req.Workspace = &dataformpb.Workspace{
										DisableMoves: cmd.Bool("disable-moves"),
									}
									resp, err := client.CreateWorkspace(ctx, req)
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
								Usage: "delete workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									client, err := dataform.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &dataformpb.DeleteWorkspaceRequest{Name: name}
									if err := client.DeleteWorkspace(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "install-npm-packages",
								Usage: "install-npm-packages workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing install-npm-packages...")
									return nil
								},
							},
							{
								Name:  "pull",
								Usage: "pull workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									fmt.Printf("Executing pull on %s\n", name)
									return nil
								},
							},
							{
								Name:  "push",
								Usage: "push workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									fmt.Printf("Executing push on %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-file-git-statuses",
								Usage: "fetch-file-git-statuses workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									fmt.Printf("Executing fetch-file-git-statuses on %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-git-ahead-behind",
								Usage: "fetch-git-ahead-behind workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									fmt.Printf("Executing fetch-git-ahead-behind on %s\n", name)
									return nil
								},
							},
							{
								Name:  "commit",
								Usage: "commit workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									fmt.Printf("Executing commit on %s\n", name)
									return nil
								},
							},
							{
								Name:  "reset",
								Usage: "reset workspaces",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "workspace", Usage: "The workspace.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workspaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("workspace"))
									fmt.Printf("Executing reset on %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-file-diff",
								Usage: "fetch-file-diff workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-file-diff...")
									return nil
								},
							},
							{
								Name:  "query-directory-contents",
								Usage: "query-directory-contents workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query-directory-contents...")
									return nil
								},
							},
							{
								Name:  "search-files",
								Usage: "search-files workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search-files...")
									return nil
								},
							},
							{
								Name:  "make-directory",
								Usage: "make-directory workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing make-directory...")
									return nil
								},
							},
							{
								Name:  "remove-directory",
								Usage: "remove-directory workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-directory...")
									return nil
								},
							},
							{
								Name:  "move-directory",
								Usage: "move-directory workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move-directory...")
									return nil
								},
							},
							{
								Name:  "read-file",
								Usage: "read-file workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing read-file...")
									return nil
								},
							},
							{
								Name:  "remove-file",
								Usage: "remove-file workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-file...")
									return nil
								},
							},
							{
								Name:  "move-file",
								Usage: "move-file workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move-file...")
									return nil
								},
							},
							{
								Name:  "write-file",
								Usage: "write-file workspaces",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing write-file...")
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
