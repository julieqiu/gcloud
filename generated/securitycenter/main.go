package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	securitycenter "cloud.google.com/go/securitycenter/apiv2"
	"cloud.google.com/go/securitycenter/apiv2/securitycenterpb"
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
				Name:  "securitycenter",
				Usage: "manage Security Command Center API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "attack-paths",
						Usage: "Manage attack-paths resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list attack-paths",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "simulation", Usage: "The simulation.", Required: true},
									&cli.StringFlag{Name: "valued_resource", Usage: "The valued_resource.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/simulations/%s/valuedResources/%s", cmd.String("organization"), cmd.String("simulation"), cmd.String("valued_resource"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListAttackPathsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAttackPaths(ctx, req)
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
						},
					},
					{
						Name:  "big-query-exports",
						Usage: "Manage big-query-exports resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create big-query-exports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "big-query-export-id", Usage: "The big query export id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
									&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.CreateBigQueryExportRequest{Parent: parent}
									req.BigQueryExportId = cmd.String("big-query-export-id")
									req.BigQueryExport = &securitycenterpb.BigQueryExport{
										Description: cmd.String("description"),
										Filter:      cmd.String("filter"),
										Dataset:     cmd.String("dataset"),
									}
									resp, err := client.CreateBigQueryExport(ctx, req)
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
								Usage: "delete big-query-exports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "export", Usage: "The export.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", cmd.String("organization"), cmd.String("location"), cmd.String("export"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.DeleteBigQueryExportRequest{Name: name}
									if err := client.DeleteBigQueryExport(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe big-query-exports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "export", Usage: "The export.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", cmd.String("organization"), cmd.String("location"), cmd.String("export"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetBigQueryExportRequest{Name: name}
									resp, err := client.GetBigQueryExport(ctx, req)
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
								Usage: "list big-query-exports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListBigQueryExportsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListBigQueryExports(ctx, req)
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
								Name:  "update",
								Usage: "update big-query-exports",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "export", Usage: "The export.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
									&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", cmd.String("organization"), cmd.String("location"), cmd.String("export"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateBigQueryExportRequest{}
									req.BigQueryExport = &securitycenterpb.BigQueryExport{
										Name:        name,
										Description: cmd.String("description"),
										Filter:      cmd.String("filter"),
										Dataset:     cmd.String("dataset"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("filter") {
										paths = append(paths, "filter")
									}
									if cmd.IsSet("dataset") {
										paths = append(paths, "dataset")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateBigQueryExport(ctx, req)
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
						Name:  "external-systems",
						Usage: "Manage external-systems resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update external-systems",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "finding", Usage: "The finding.", Required: true},
									&cli.StringFlag{Name: "externalsystem", Usage: "The externalsystem.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "external-uid", Usage: "The external uid.", Required: false},
									&cli.StringFlag{Name: "status", Usage: "The status.", Required: false},
									&cli.StringFlag{Name: "case-uri", Usage: "The case uri.", Required: false},
									&cli.StringFlag{Name: "case-priority", Usage: "The case priority.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s/externalSystems/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"), cmd.String("externalsystem"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateExternalSystemRequest{}
									req.ExternalSystem = &securitycenterpb.ExternalSystem{
										Name:         name,
										Name:         cmd.String("name"),
										ExternalUid:  cmd.String("external-uid"),
										Status:       cmd.String("status"),
										CaseUri:      cmd.String("case-uri"),
										CasePriority: cmd.String("case-priority"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("external-uid") {
										paths = append(paths, "external_uid")
									}
									if cmd.IsSet("status") {
										paths = append(paths, "status")
									}
									if cmd.IsSet("case-uri") {
										paths = append(paths, "case_uri")
									}
									if cmd.IsSet("case-priority") {
										paths = append(paths, "case_priority")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateExternalSystem(ctx, req)
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
						Name:  "findings",
						Usage: "Manage findings resources",
						Commands: []*cli.Command{
							{
								Name:  "bulk-mute",
								Usage: "bulk-mute findings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing bulk-mute...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create findings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "finding-id", Usage: "The finding id.", Required: true},
									&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: false},
									&cli.StringFlag{Name: "resource-name", Usage: "The resource name.", Required: false},
									&cli.StringFlag{Name: "category", Usage: "The category.", Required: false},
									&cli.StringFlag{Name: "external-uri", Usage: "The external uri.", Required: false},
									&cli.StringFlag{Name: "mute-initiator", Usage: "The mute initiator.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "next-steps", Usage: "The next steps.", Required: false},
									&cli.StringFlag{Name: "module-name", Usage: "The module name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.CreateFindingRequest{Parent: parent}
									req.FindingId = cmd.String("finding-id")
									req.Finding = &securitycenterpb.Finding{
										Parent:        cmd.String("parent"),
										ResourceName:  cmd.String("resource-name"),
										Category:      cmd.String("category"),
										ExternalUri:   cmd.String("external-uri"),
										MuteInitiator: cmd.String("mute-initiator"),
										Description:   cmd.String("description"),
										NextSteps:     cmd.String("next-steps"),
										ModuleName:    cmd.String("module-name"),
									}
									resp, err := client.CreateFinding(ctx, req)
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
								Name:  "group",
								Usage: "group findings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									fmt.Printf("Executing group on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list findings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListFindingsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFindings(ctx, req)
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
								Name:  "set-state",
								Usage: "set-state findings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "finding", Usage: "The finding.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
									fmt.Printf("Executing set-state on %s\n", name)
									return nil
								},
							},
							{
								Name:  "set-mute",
								Usage: "set-mute findings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "finding", Usage: "The finding.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
									fmt.Printf("Executing set-mute on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update findings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "finding", Usage: "The finding.", Required: true},
									&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: false},
									&cli.StringFlag{Name: "resource-name", Usage: "The resource name.", Required: false},
									&cli.StringFlag{Name: "category", Usage: "The category.", Required: false},
									&cli.StringFlag{Name: "external-uri", Usage: "The external uri.", Required: false},
									&cli.StringFlag{Name: "mute-initiator", Usage: "The mute initiator.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "next-steps", Usage: "The next steps.", Required: false},
									&cli.StringFlag{Name: "module-name", Usage: "The module name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateFindingRequest{}
									req.Finding = &securitycenterpb.Finding{
										Name:          name,
										Parent:        cmd.String("parent"),
										ResourceName:  cmd.String("resource-name"),
										Category:      cmd.String("category"),
										ExternalUri:   cmd.String("external-uri"),
										MuteInitiator: cmd.String("mute-initiator"),
										Description:   cmd.String("description"),
										NextSteps:     cmd.String("next-steps"),
										ModuleName:    cmd.String("module-name"),
									}
									var paths []string
									if cmd.IsSet("parent") {
										paths = append(paths, "parent")
									}
									if cmd.IsSet("resource-name") {
										paths = append(paths, "resource_name")
									}
									if cmd.IsSet("category") {
										paths = append(paths, "category")
									}
									if cmd.IsSet("external-uri") {
										paths = append(paths, "external_uri")
									}
									if cmd.IsSet("mute-initiator") {
										paths = append(paths, "mute_initiator")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("next-steps") {
										paths = append(paths, "next_steps")
									}
									if cmd.IsSet("module-name") {
										paths = append(paths, "module_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateFinding(ctx, req)
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
						Name:  "mute-configs",
						Usage: "Manage mute-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create mute-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "mute-config-id", Usage: "The mute config id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.CreateMuteConfigRequest{Parent: parent}
									req.MuteConfigId = cmd.String("mute-config-id")
									req.MuteConfig = &securitycenterpb.MuteConfig{
										Description: cmd.String("description"),
										Filter:      cmd.String("filter"),
									}
									resp, err := client.CreateMuteConfig(ctx, req)
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
								Usage: "delete mute-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "mute_config", Usage: "The mute_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/muteConfigs/%s", cmd.String("organization"), cmd.String("mute_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.DeleteMuteConfigRequest{Name: name}
									if err := client.DeleteMuteConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe mute-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "mute_config", Usage: "The mute_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/muteConfigs/%s", cmd.String("organization"), cmd.String("mute_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetMuteConfigRequest{Name: name}
									resp, err := client.GetMuteConfig(ctx, req)
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
								Usage: "list mute-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListMuteConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMuteConfigs(ctx, req)
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
								Name:  "update",
								Usage: "update mute-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "mute_config", Usage: "The mute_config.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/muteConfigs/%s", cmd.String("organization"), cmd.String("mute_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateMuteConfigRequest{}
									req.MuteConfig = &securitycenterpb.MuteConfig{
										Name:        name,
										Description: cmd.String("description"),
										Filter:      cmd.String("filter"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("filter") {
										paths = append(paths, "filter")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateMuteConfig(ctx, req)
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
						Name:  "notification-configs",
						Usage: "Manage notification-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create notification-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "config-id", Usage: "The config id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "pubsub-topic", Usage: "The pubsub topic.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.CreateNotificationConfigRequest{Parent: parent}
									req.ConfigId = cmd.String("config-id")
									req.NotificationConfig = &securitycenterpb.NotificationConfig{
										Description: cmd.String("description"),
										PubsubTopic: cmd.String("pubsub-topic"),
									}
									resp, err := client.CreateNotificationConfig(ctx, req)
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
								Usage: "delete notification-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "notification_config", Usage: "The notification_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/notificationConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.DeleteNotificationConfigRequest{Name: name}
									if err := client.DeleteNotificationConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe notification-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "notification_config", Usage: "The notification_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/notificationConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetNotificationConfigRequest{Name: name}
									resp, err := client.GetNotificationConfig(ctx, req)
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
								Usage: "list notification-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListNotificationConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListNotificationConfigs(ctx, req)
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
								Name:  "update",
								Usage: "update notification-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "notification_config", Usage: "The notification_config.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "pubsub-topic", Usage: "The pubsub topic.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/notificationConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateNotificationConfigRequest{}
									req.NotificationConfig = &securitycenterpb.NotificationConfig{
										Name:        name,
										Description: cmd.String("description"),
										PubsubTopic: cmd.String("pubsub-topic"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("pubsub-topic") {
										paths = append(paths, "pubsub_topic")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateNotificationConfig(ctx, req)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/operations", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/operations/%s", cmd.String("organization"), cmd.String("operation"))
									client, err := securitycenter.NewClient(ctx)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/operations/%s", cmd.String("organization"), cmd.String("operation"))
									client, err := securitycenter.NewClient(ctx)
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
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/operations/%s", cmd.String("organization"), cmd.String("operation"))
									client, err := securitycenter.NewClient(ctx)
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
						Name:  "resource-value-configs",
						Usage: "Manage resource-value-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "batch-create",
								Usage: "batch-create resource-value-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									fmt.Printf("Executing batch-create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete resource-value-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "resource_value_config", Usage: "The resource_value_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/resourceValueConfigs/%s", cmd.String("organization"), cmd.String("resource_value_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.DeleteResourceValueConfigRequest{Name: name}
									if err := client.DeleteResourceValueConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe resource-value-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "resource_value_config", Usage: "The resource_value_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/resourceValueConfigs/%s", cmd.String("organization"), cmd.String("resource_value_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetResourceValueConfigRequest{Name: name}
									resp, err := client.GetResourceValueConfig(ctx, req)
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
								Usage: "list resource-value-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListResourceValueConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListResourceValueConfigs(ctx, req)
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
								Name:  "update",
								Usage: "update resource-value-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "resource_value_config", Usage: "The resource_value_config.", Required: true},
									&cli.StringFlag{Name: "resource-type", Usage: "The resource type.", Required: false},
									&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/resourceValueConfigs/%s", cmd.String("organization"), cmd.String("resource_value_config"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateResourceValueConfigRequest{}
									req.ResourceValueConfig = &securitycenterpb.ResourceValueConfig{
										Name:         name,
										ResourceType: cmd.String("resource-type"),
										Scope:        cmd.String("scope"),
										Description:  cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("resource-type") {
										paths = append(paths, "resource_type")
									}
									if cmd.IsSet("scope") {
										paths = append(paths, "scope")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateResourceValueConfig(ctx, req)
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
						Name:  "security-marks",
						Usage: "Manage security-marks resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update security-marks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "canonical-name", Usage: "The canonical name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/assets/%s/securityMarks", cmd.String("organization"), cmd.String("asset"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateSecurityMarksRequest{}
									req.SecurityMarks = &securitycenterpb.SecurityMarks{
										Name:          name,
										Name:          cmd.String("name"),
										CanonicalName: cmd.String("canonical-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("canonical-name") {
										paths = append(paths, "canonical_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSecurityMarks(ctx, req)
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
						Name:  "simulations",
						Usage: "Manage simulations resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe simulations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "simulation", Usage: "The simulation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/simulations/%s", cmd.String("organization"), cmd.String("simulation"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetSimulationRequest{Name: name}
									resp, err := client.GetSimulation(ctx, req)
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
						Name:  "sources",
						Usage: "Manage sources resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "canonical-name", Usage: "The canonical name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.CreateSourceRequest{Parent: parent}
									req.Source = &securitycenterpb.Source{
										Name:          cmd.String("name"),
										DisplayName:   cmd.String("display-name"),
										Description:   cmd.String("description"),
										CanonicalName: cmd.String("canonical-name"),
									}
									resp, err := client.CreateSource(ctx, req)
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
								Name:  "get-iam-policy",
								Usage: "get-iam-policy sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetSourceRequest{Name: name}
									resp, err := client.GetSource(ctx, req)
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
								Usage: "list sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListSourcesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSources(ctx, req)
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
								Name:  "set-iam-policy",
								Usage: "set-iam-policy sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update sources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "source", Usage: "The source.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "canonical-name", Usage: "The canonical name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.UpdateSourceRequest{}
									req.Source = &securitycenterpb.Source{
										Name:          name,
										Name:          cmd.String("name"),
										DisplayName:   cmd.String("display-name"),
										Description:   cmd.String("description"),
										CanonicalName: cmd.String("canonical-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("canonical-name") {
										paths = append(paths, "canonical_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSource(ctx, req)
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
						Name:  "valued-resources",
						Usage: "Manage valued-resources resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe valued-resources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "simulation", Usage: "The simulation.", Required: true},
									&cli.StringFlag{Name: "valued_resource", Usage: "The valued_resource.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/simulations/%s/valuedResources/%s", cmd.String("organization"), cmd.String("simulation"), cmd.String("valued_resource"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securitycenterpb.GetValuedResourceRequest{Name: name}
									resp, err := client.GetValuedResource(ctx, req)
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
								Usage: "list valued-resources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "simulation", Usage: "The simulation.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/simulations/%s", cmd.String("organization"), cmd.String("simulation"))
									client, err := securitycenter.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securitycenterpb.ListValuedResourcesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListValuedResources(ctx, req)
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
