package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	"cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
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
				Name:  "dialogflow",
				Usage: "manage Dialogflow API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "agents",
						Usage: "Manage agents resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListAgentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAgents(ctx, req)
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
								Usage: "describe agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetAgentRequest{Name: name}
									resp, err := client.GetAgent(ctx, req)
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
								Usage: "create agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "default-language-code", Usage: "The default language code.", Required: true},
									&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "avatar-uri", Usage: "The avatar uri.", Required: false},
									&cli.StringFlag{Name: "security-settings", Usage: "The security settings.", Required: false},
									&cli.BoolFlag{Name: "enable-stackdriver-logging", Usage: "The enable stackdriver logging.", Required: false},
									&cli.BoolFlag{Name: "enable-spell-correction", Usage: "The enable spell correction.", Required: false},
									&cli.BoolFlag{Name: "enable-multi-language-training", Usage: "The enable multi language training.", Required: false},
									&cli.BoolFlag{Name: "locked", Usage: "The locked.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateAgentRequest{Parent: parent}
									req.Agent = &cxpb.Agent{
										Name:                        cmd.String("name"),
										DisplayName:                 cmd.String("display-name"),
										DefaultLanguageCode:         cmd.String("default-language-code"),
										TimeZone:                    cmd.String("time-zone"),
										Description:                 cmd.String("description"),
										AvatarUri:                   cmd.String("avatar-uri"),
										SecuritySettings:            cmd.String("security-settings"),
										EnableStackdriverLogging:    cmd.Bool("enable-stackdriver-logging"),
										EnableSpellCorrection:       cmd.Bool("enable-spell-correction"),
										EnableMultiLanguageTraining: cmd.Bool("enable-multi-language-training"),
										Locked:                      cmd.Bool("locked"),
									}
									resp, err := client.CreateAgent(ctx, req)
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
								Usage: "update agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "default-language-code", Usage: "The default language code.", Required: false},
									&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "avatar-uri", Usage: "The avatar uri.", Required: false},
									&cli.StringFlag{Name: "security-settings", Usage: "The security settings.", Required: false},
									&cli.BoolFlag{Name: "enable-stackdriver-logging", Usage: "The enable stackdriver logging.", Required: false},
									&cli.BoolFlag{Name: "enable-spell-correction", Usage: "The enable spell correction.", Required: false},
									&cli.BoolFlag{Name: "enable-multi-language-training", Usage: "The enable multi language training.", Required: false},
									&cli.BoolFlag{Name: "locked", Usage: "The locked.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateAgentRequest{}
									req.Agent = &cxpb.Agent{
										Name:                        name,
										Name:                        cmd.String("name"),
										DisplayName:                 cmd.String("display-name"),
										DefaultLanguageCode:         cmd.String("default-language-code"),
										TimeZone:                    cmd.String("time-zone"),
										Description:                 cmd.String("description"),
										AvatarUri:                   cmd.String("avatar-uri"),
										SecuritySettings:            cmd.String("security-settings"),
										EnableStackdriverLogging:    cmd.Bool("enable-stackdriver-logging"),
										EnableSpellCorrection:       cmd.Bool("enable-spell-correction"),
										EnableMultiLanguageTraining: cmd.Bool("enable-multi-language-training"),
										Locked:                      cmd.Bool("locked"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("default-language-code") {
										paths = append(paths, "default_language_code")
									}
									if cmd.IsSet("time-zone") {
										paths = append(paths, "time_zone")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("avatar-uri") {
										paths = append(paths, "avatar_uri")
									}
									if cmd.IsSet("security-settings") {
										paths = append(paths, "security_settings")
									}
									if cmd.IsSet("enable-stackdriver-logging") {
										paths = append(paths, "enable_stackdriver_logging")
									}
									if cmd.IsSet("enable-spell-correction") {
										paths = append(paths, "enable_spell_correction")
									}
									if cmd.IsSet("enable-multi-language-training") {
										paths = append(paths, "enable_multi_language_training")
									}
									if cmd.IsSet("locked") {
										paths = append(paths, "locked")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateAgent(ctx, req)
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
								Usage: "delete agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteAgentRequest{Name: name}
									if err := client.DeleteAgent(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "export",
								Usage: "export agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "agent-uri", Usage: "The agent uri.", Required: false},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: false},
									&cli.BoolFlag{Name: "include-bigquery-export-settings", Usage: "The include bigquery export settings.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.ExportAgentRequest{Name: name}
									req.AgentUri = cmd.String("agent-uri")
									req.Environment = cmd.String("environment")
									req.IncludeBigqueryExportSettings = cmd.Bool("include-bigquery-export-settings")
									op, err := client.ExportAgent(ctx, req)
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
								Name:  "restore",
								Usage: "restore agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.RestoreAgentRequest{Name: name}
									op, err := client.RestoreAgent(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("RestoreAgent completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "validate",
								Usage: "validate agents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing validate on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "changelogs",
						Usage: "Manage changelogs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list changelogs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewChangelogsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListChangelogsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListChangelogs(ctx, req)
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
								Usage: "describe changelogs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "changelog", Usage: "The changelog.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/changelogs/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("changelog"))
									client, err := cx.NewChangelogsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetChangelogRequest{Name: name}
									resp, err := client.GetChangelog(ctx, req)
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
						Name:  "continuous-test-results",
						Usage: "Manage continuous-test-results resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list continuous-test-results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewEnvironmentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListContinuousTestResultsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListContinuousTestResults(ctx, req)
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
						Name:  "deployments",
						Usage: "Manage deployments resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewDeploymentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListDeploymentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDeployments(ctx, req)
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
								Usage: "describe deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("deployment"))
									client, err := cx.NewDeploymentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetDeploymentRequest{Name: name}
									resp, err := client.GetDeployment(ctx, req)
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
						Name:  "entity-types",
						Usage: "Manage entity-types resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "entity_type", Usage: "The entity_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("entity_type"))
									client, err := cx.NewEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetEntityTypeRequest{Name: name}
									resp, err := client.GetEntityType(ctx, req)
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
								Usage: "create entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.BoolFlag{Name: "enable-fuzzy-extraction", Usage: "The enable fuzzy extraction.", Required: false},
									&cli.BoolFlag{Name: "redact", Usage: "The redact.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateEntityTypeRequest{Parent: parent}
									req.LanguageCode = cmd.String("language-code")
									req.EntityType = &cxpb.EntityType{
										Name:                  cmd.String("name"),
										DisplayName:           cmd.String("display-name"),
										EnableFuzzyExtraction: cmd.Bool("enable-fuzzy-extraction"),
										Redact:                cmd.Bool("redact"),
									}
									resp, err := client.CreateEntityType(ctx, req)
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
								Usage: "update entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "entity_type", Usage: "The entity_type.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.BoolFlag{Name: "enable-fuzzy-extraction", Usage: "The enable fuzzy extraction.", Required: false},
									&cli.BoolFlag{Name: "redact", Usage: "The redact.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("entity_type"))
									client, err := cx.NewEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateEntityTypeRequest{}
									req.EntityType = &cxpb.EntityType{
										Name:                  name,
										Name:                  cmd.String("name"),
										DisplayName:           cmd.String("display-name"),
										EnableFuzzyExtraction: cmd.Bool("enable-fuzzy-extraction"),
										Redact:                cmd.Bool("redact"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("enable-fuzzy-extraction") {
										paths = append(paths, "enable_fuzzy_extraction")
									}
									if cmd.IsSet("redact") {
										paths = append(paths, "redact")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateEntityType(ctx, req)
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
								Usage: "delete entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "entity_type", Usage: "The entity_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("entity_type"))
									client, err := cx.NewEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteEntityTypeRequest{Name: name}
									if err := client.DeleteEntityType(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListEntityTypesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListEntityTypes(ctx, req)
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
								Name:  "export",
								Usage: "export entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing export on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "import",
								Usage: "import entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListSessionEntityTypesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSessionEntityTypes(ctx, req)
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
								Usage: "describe entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
									&cli.StringFlag{Name: "entity_type", Usage: "The entity_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"), cmd.String("entity_type"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetSessionEntityTypeRequest{Name: name}
									resp, err := client.GetSessionEntityType(ctx, req)
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
								Usage: "create entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateSessionEntityTypeRequest{Parent: parent}
									req.SessionEntityType = &cxpb.SessionEntityType{
										Name: cmd.String("name"),
									}
									resp, err := client.CreateSessionEntityType(ctx, req)
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
								Usage: "update entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
									&cli.StringFlag{Name: "entity_type", Usage: "The entity_type.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"), cmd.String("entity_type"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateSessionEntityTypeRequest{}
									req.SessionEntityType = &cxpb.SessionEntityType{
										Name: name,
										Name: cmd.String("name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSessionEntityType(ctx, req)
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
								Usage: "delete entity-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "session", Usage: "The session.", Required: true},
									&cli.StringFlag{Name: "entity_type", Usage: "The entity_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/sessions/%s/entityTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("session"), cmd.String("entity_type"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteSessionEntityTypeRequest{Name: name}
									if err := client.DeleteSessionEntityType(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "environments",
						Usage: "Manage environments resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list environments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewEnvironmentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListEnvironmentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListEnvironments(ctx, req)
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
								Usage: "describe environments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewEnvironmentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetEnvironmentRequest{Name: name}
									resp, err := client.GetEnvironment(ctx, req)
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
								Usage: "create environments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewEnvironmentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateEnvironmentRequest{Parent: parent}
									req.Environment = &cxpb.Environment{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateEnvironment(ctx, req)
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
								Usage: "update environments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewEnvironmentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateEnvironmentRequest{}
									req.Environment = &cxpb.Environment{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
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
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateEnvironment(ctx, req)
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
								Usage: "delete environments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewEnvironmentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteEnvironmentRequest{Name: name}
									if err := client.DeleteEnvironment(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "lookup-environment-history",
								Usage: "lookup-environment-history environments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									fmt.Printf("Executing lookup-environment-history on %s\n", name)
									return nil
								},
							},
							{
								Name:  "run-continuous-test",
								Usage: "run-continuous-test environments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing run-continuous-test...")
									return nil
								},
							},
							{
								Name:  "deploy-flow",
								Usage: "deploy-flow environments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing deploy-flow...")
									return nil
								},
							},
						},
					},
					{
						Name:  "examples",
						Usage: "Manage examples resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create examples",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewExamplesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateExampleRequest{Parent: parent}
									req.Example = &cxpb.Example{
										Name:         cmd.String("name"),
										DisplayName:  cmd.String("display-name"),
										Description:  cmd.String("description"),
										LanguageCode: cmd.String("language-code"),
									}
									resp, err := client.CreateExample(ctx, req)
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
								Usage: "delete examples",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "example", Usage: "The example.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("example"))
									client, err := cx.NewExamplesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteExampleRequest{Name: name}
									if err := client.DeleteExample(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list examples",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewExamplesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListExamplesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListExamples(ctx, req)
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
								Usage: "describe examples",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "example", Usage: "The example.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("example"))
									client, err := cx.NewExamplesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetExampleRequest{Name: name}
									resp, err := client.GetExample(ctx, req)
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
								Usage: "update examples",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "example", Usage: "The example.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/examples/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("example"))
									client, err := cx.NewExamplesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateExampleRequest{}
									req.Example = &cxpb.Example{
										Name:         name,
										Name:         cmd.String("name"),
										DisplayName:  cmd.String("display-name"),
										Description:  cmd.String("description"),
										LanguageCode: cmd.String("language-code"),
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
									if cmd.IsSet("language-code") {
										paths = append(paths, "language_code")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateExample(ctx, req)
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
						Name:  "experiments",
						Usage: "Manage experiments resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewExperimentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListExperimentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListExperiments(ctx, req)
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
								Usage: "describe experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "experiment", Usage: "The experiment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
									client, err := cx.NewExperimentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetExperimentRequest{Name: name}
									resp, err := client.GetExperiment(ctx, req)
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
								Usage: "create experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "rollout-failure-reason", Usage: "The rollout failure reason.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"))
									client, err := cx.NewExperimentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateExperimentRequest{Parent: parent}
									req.Experiment = &cxpb.Experiment{
										Name:                 cmd.String("name"),
										DisplayName:          cmd.String("display-name"),
										Description:          cmd.String("description"),
										RolloutFailureReason: cmd.String("rollout-failure-reason"),
									}
									resp, err := client.CreateExperiment(ctx, req)
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
								Usage: "update experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "experiment", Usage: "The experiment.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "rollout-failure-reason", Usage: "The rollout failure reason.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
									client, err := cx.NewExperimentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateExperimentRequest{}
									req.Experiment = &cxpb.Experiment{
										Name:                 name,
										Name:                 cmd.String("name"),
										DisplayName:          cmd.String("display-name"),
										Description:          cmd.String("description"),
										RolloutFailureReason: cmd.String("rollout-failure-reason"),
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
									if cmd.IsSet("rollout-failure-reason") {
										paths = append(paths, "rollout_failure_reason")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateExperiment(ctx, req)
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
								Usage: "delete experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "experiment", Usage: "The experiment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
									client, err := cx.NewExperimentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteExperimentRequest{Name: name}
									if err := client.DeleteExperiment(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "start",
								Usage: "start experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "experiment", Usage: "The experiment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
									fmt.Printf("Executing start on %s\n", name)
									return nil
								},
							},
							{
								Name:  "stop",
								Usage: "stop experiments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
									&cli.StringFlag{Name: "experiment", Usage: "The experiment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/environments/%s/experiments/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("environment"), cmd.String("experiment"))
									fmt.Printf("Executing stop on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "flows",
						Usage: "Manage flows resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "locked", Usage: "The locked.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateFlowRequest{Parent: parent}
									req.LanguageCode = cmd.String("language-code")
									req.Flow = &cxpb.Flow{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
										Locked:      cmd.Bool("locked"),
									}
									resp, err := client.CreateFlow(ctx, req)
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
								Usage: "delete flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteFlowRequest{Name: name}
									if err := client.DeleteFlow(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListFlowsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFlows(ctx, req)
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
								Usage: "describe flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetFlowRequest{Name: name}
									resp, err := client.GetFlow(ctx, req)
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
								Usage: "update flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "locked", Usage: "The locked.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateFlowRequest{}
									req.Flow = &cxpb.Flow{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
										Locked:      cmd.Bool("locked"),
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
									if cmd.IsSet("locked") {
										paths = append(paths, "locked")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateFlow(ctx, req)
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
								Name:  "train",
								Usage: "train flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.TrainFlowRequest{Name: name}
									op, err := client.TrainFlow(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("TrainFlow completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "validate",
								Usage: "validate flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									fmt.Printf("Executing validate on %s\n", name)
									return nil
								},
							},
							{
								Name:  "import",
								Usage: "import flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "export",
								Usage: "export flows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "flow-uri", Usage: "The flow uri.", Required: false},
									&cli.BoolFlag{Name: "include-referenced-flows", Usage: "The include referenced flows.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.ExportFlowRequest{Name: name}
									req.FlowUri = cmd.String("flow-uri")
									req.IncludeReferencedFlows = cmd.Bool("include-referenced-flows")
									op, err := client.ExportFlow(ctx, req)
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
						Name:  "generative-settings",
						Usage: "Manage generative-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe generative-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generativeSettings", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetGenerativeSettingsRequest{Name: name}
									resp, err := client.GetGenerativeSettings(ctx, req)
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
								Usage: "update generative-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generativeSettings", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateGenerativeSettingsRequest{}
									req.GenerativeSettings = &cxpb.GenerativeSettings{
										Name:         name,
										Name:         cmd.String("name"),
										LanguageCode: cmd.String("language-code"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("language-code") {
										paths = append(paths, "language_code")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateGenerativeSettings(ctx, req)
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
						Name:  "generators",
						Usage: "Manage generators resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list generators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewGeneratorsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListGeneratorsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGenerators(ctx, req)
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
								Usage: "describe generators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "generator", Usage: "The generator.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("generator"))
									client, err := cx.NewGeneratorsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetGeneratorRequest{Name: name}
									resp, err := client.GetGenerator(ctx, req)
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
								Usage: "create generators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewGeneratorsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateGeneratorRequest{Parent: parent}
									req.LanguageCode = cmd.String("language-code")
									req.Generator = &cxpb.Generator{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateGenerator(ctx, req)
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
								Usage: "update generators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "generator", Usage: "The generator.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("generator"))
									client, err := cx.NewGeneratorsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateGeneratorRequest{}
									req.Generator = &cxpb.Generator{
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
									resp, err := client.UpdateGenerator(ctx, req)
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
								Usage: "delete generators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "generator", Usage: "The generator.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("generator"))
									client, err := cx.NewGeneratorsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteGeneratorRequest{Name: name}
									if err := client.DeleteGenerator(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "intents",
						Usage: "Manage intents resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewIntentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListIntentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListIntents(ctx, req)
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
								Usage: "describe intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "intent", Usage: "The intent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/intents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("intent"))
									client, err := cx.NewIntentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetIntentRequest{Name: name}
									resp, err := client.GetIntent(ctx, req)
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
								Usage: "create intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: false},
									&cli.BoolFlag{Name: "is-fallback", Usage: "The is fallback.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "dtmf-pattern", Usage: "The dtmf pattern.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewIntentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateIntentRequest{Parent: parent}
									req.LanguageCode = cmd.String("language-code")
									req.Intent = &cxpb.Intent{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Priority:    int32(cmd.Int("priority")),
										IsFallback:  cmd.Bool("is-fallback"),
										Description: cmd.String("description"),
										DtmfPattern: cmd.String("dtmf-pattern"),
									}
									resp, err := client.CreateIntent(ctx, req)
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
								Usage: "update intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "intent", Usage: "The intent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: false},
									&cli.BoolFlag{Name: "is-fallback", Usage: "The is fallback.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "dtmf-pattern", Usage: "The dtmf pattern.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/intents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("intent"))
									client, err := cx.NewIntentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateIntentRequest{}
									req.Intent = &cxpb.Intent{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Priority:    int32(cmd.Int("priority")),
										IsFallback:  cmd.Bool("is-fallback"),
										Description: cmd.String("description"),
										DtmfPattern: cmd.String("dtmf-pattern"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("priority") {
										paths = append(paths, "priority")
									}
									if cmd.IsSet("is-fallback") {
										paths = append(paths, "is_fallback")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("dtmf-pattern") {
										paths = append(paths, "dtmf_pattern")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateIntent(ctx, req)
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
								Usage: "delete intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "intent", Usage: "The intent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/intents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("intent"))
									client, err := cx.NewIntentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteIntentRequest{Name: name}
									if err := client.DeleteIntent(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "import",
								Usage: "import intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "export",
								Usage: "export intents",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing export on %s\n", parent)
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
									client, err := cx.NewAgentsClient(ctx)
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
									client, err := cx.NewAgentsClient(ctx)
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
									client, err := cx.NewChangelogsClient(ctx)
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
									client, err := cx.NewChangelogsClient(ctx)
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
									client, err := cx.NewDeploymentsClient(ctx)
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
									client, err := cx.NewDeploymentsClient(ctx)
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
									client, err := cx.NewEntityTypesClient(ctx)
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
									client, err := cx.NewEntityTypesClient(ctx)
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
									client, err := cx.NewEnvironmentsClient(ctx)
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
									client, err := cx.NewEnvironmentsClient(ctx)
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
									client, err := cx.NewExamplesClient(ctx)
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
									client, err := cx.NewExamplesClient(ctx)
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
									client, err := cx.NewExperimentsClient(ctx)
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
									client, err := cx.NewExperimentsClient(ctx)
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
									client, err := cx.NewFlowsClient(ctx)
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
									client, err := cx.NewFlowsClient(ctx)
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
									client, err := cx.NewGeneratorsClient(ctx)
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
									client, err := cx.NewGeneratorsClient(ctx)
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
									client, err := cx.NewIntentsClient(ctx)
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
									client, err := cx.NewIntentsClient(ctx)
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
									client, err := cx.NewPagesClient(ctx)
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
									client, err := cx.NewPagesClient(ctx)
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
									client, err := cx.NewPlaybooksClient(ctx)
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
									client, err := cx.NewPlaybooksClient(ctx)
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
									client, err := cx.NewSecuritySettingsClient(ctx)
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
									client, err := cx.NewSecuritySettingsClient(ctx)
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
									client, err := cx.NewSessionsClient(ctx)
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
									client, err := cx.NewSessionsClient(ctx)
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
									client, err := cx.NewSessionEntityTypesClient(ctx)
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
									client, err := cx.NewSessionEntityTypesClient(ctx)
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
									client, err := cx.NewTestCasesClient(ctx)
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
									client, err := cx.NewTestCasesClient(ctx)
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
									client, err := cx.NewToolsClient(ctx)
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
									client, err := cx.NewToolsClient(ctx)
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
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
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
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
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
									client, err := cx.NewVersionsClient(ctx)
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
									client, err := cx.NewVersionsClient(ctx)
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
									client, err := cx.NewWebhooksClient(ctx)
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
									client, err := cx.NewWebhooksClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewAgentsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewAgentsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewAgentsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewChangelogsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewChangelogsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewChangelogsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewDeploymentsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewDeploymentsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewDeploymentsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewEntityTypesClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewEntityTypesClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewEntityTypesClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewEnvironmentsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewEnvironmentsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewEnvironmentsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewExamplesClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewExamplesClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewExamplesClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewExperimentsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewExperimentsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewExperimentsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewFlowsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewFlowsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewFlowsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewGeneratorsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewGeneratorsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewGeneratorsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewIntentsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewIntentsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewIntentsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewPagesClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewPagesClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewPagesClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewPlaybooksClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewPlaybooksClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewPlaybooksClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewSecuritySettingsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewSecuritySettingsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewSecuritySettingsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewSessionsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewSessionsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewSessionsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewSessionEntityTypesClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewTestCasesClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewTestCasesClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewTestCasesClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewToolsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewToolsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewToolsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewVersionsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewVersionsClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewVersionsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := cx.NewWebhooksClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewWebhooksClient(ctx)
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
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := cx.NewWebhooksClient(ctx)
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
						Name:  "pages",
						Usage: "Manage pages resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list pages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewPagesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListPagesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPages(ctx, req)
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
								Usage: "describe pages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "page", Usage: "The page.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("page"))
									client, err := cx.NewPagesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetPageRequest{Name: name}
									resp, err := client.GetPage(ctx, req)
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
								Usage: "create pages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewPagesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreatePageRequest{Parent: parent}
									req.LanguageCode = cmd.String("language-code")
									req.Page = &cxpb.Page{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreatePage(ctx, req)
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
								Usage: "update pages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "page", Usage: "The page.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("page"))
									client, err := cx.NewPagesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdatePageRequest{}
									req.Page = &cxpb.Page{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
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
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdatePage(ctx, req)
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
								Usage: "delete pages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "page", Usage: "The page.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("page"))
									client, err := cx.NewPagesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeletePageRequest{Name: name}
									if err := client.DeletePage(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "playbooks",
						Usage: "Manage playbooks resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "goal", Usage: "The goal.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreatePlaybookRequest{Parent: parent}
									req.Playbook = &cxpb.Playbook{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Goal:        cmd.String("goal"),
									}
									resp, err := client.CreatePlaybook(ctx, req)
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
								Usage: "delete playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeletePlaybookRequest{Name: name}
									if err := client.DeletePlaybook(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListPlaybooksRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPlaybooks(ctx, req)
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
								Usage: "describe playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetPlaybookRequest{Name: name}
									resp, err := client.GetPlaybook(ctx, req)
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
								Name:  "export",
								Usage: "export playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "playbook-uri", Usage: "The playbook uri.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.ExportPlaybookRequest{Name: name}
									req.PlaybookUri = cmd.String("playbook-uri")
									op, err := client.ExportPlaybook(ctx, req)
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
								Name:  "import",
								Usage: "import playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update playbooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "goal", Usage: "The goal.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdatePlaybookRequest{}
									req.Playbook = &cxpb.Playbook{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Goal:        cmd.String("goal"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("goal") {
										paths = append(paths, "goal")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdatePlaybook(ctx, req)
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
						Name:  "results",
						Usage: "Manage results resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "test_case", Usage: "The test_case.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test_case"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListTestCaseResultsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTestCaseResults(ctx, req)
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
								Usage: "describe results",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "test_case", Usage: "The test_case.", Required: true},
									&cli.StringFlag{Name: "result", Usage: "The result.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s/results/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test_case"), cmd.String("result"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetTestCaseResultRequest{Name: name}
									resp, err := client.GetTestCaseResult(ctx, req)
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
						Name:  "security-settings",
						Usage: "Manage security-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create security-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "inspect-template", Usage: "The inspect template.", Required: false},
									&cli.StringFlag{Name: "deidentify-template", Usage: "The deidentify template.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := cx.NewSecuritySettingsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateSecuritySettingsRequest{Parent: parent}
									req.SecuritySettings = &cxpb.SecuritySettings{
										Name:               cmd.String("name"),
										DisplayName:        cmd.String("display-name"),
										InspectTemplate:    cmd.String("inspect-template"),
										DeidentifyTemplate: cmd.String("deidentify-template"),
									}
									resp, err := client.CreateSecuritySettings(ctx, req)
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
								Usage: "describe security-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_settings", Usage: "The security_settings.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", cmd.String("project"), cmd.String("location"), cmd.String("security_settings"))
									client, err := cx.NewSecuritySettingsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetSecuritySettingsRequest{Name: name}
									resp, err := client.GetSecuritySettings(ctx, req)
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
								Usage: "update security-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_settings", Usage: "The security_settings.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "inspect-template", Usage: "The inspect template.", Required: false},
									&cli.StringFlag{Name: "deidentify-template", Usage: "The deidentify template.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", cmd.String("project"), cmd.String("location"), cmd.String("security_settings"))
									client, err := cx.NewSecuritySettingsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateSecuritySettingsRequest{}
									req.SecuritySettings = &cxpb.SecuritySettings{
										Name:               name,
										Name:               cmd.String("name"),
										DisplayName:        cmd.String("display-name"),
										InspectTemplate:    cmd.String("inspect-template"),
										DeidentifyTemplate: cmd.String("deidentify-template"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("inspect-template") {
										paths = append(paths, "inspect_template")
									}
									if cmd.IsSet("deidentify-template") {
										paths = append(paths, "deidentify_template")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSecuritySettings(ctx, req)
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
								Usage: "list security-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := cx.NewSecuritySettingsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListSecuritySettingsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSecuritySettings(ctx, req)
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
								Usage: "delete security-settings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_settings", Usage: "The security_settings.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/securitySettings/%s", cmd.String("project"), cmd.String("location"), cmd.String("security_settings"))
									client, err := cx.NewSecuritySettingsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteSecuritySettingsRequest{Name: name}
									if err := client.DeleteSecuritySettings(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "sessions",
						Usage: "Manage sessions resources",
						Commands: []*cli.Command{
							{
								Name:  "detect-intent",
								Usage: "detect-intent sessions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detect-intent...")
									return nil
								},
							},
							{
								Name:  "server-streaming-detect-intent",
								Usage: "server-streaming-detect-intent sessions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing server-streaming-detect-intent...")
									return nil
								},
							},
							{
								Name:  "match-intent",
								Usage: "match-intent sessions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing match-intent...")
									return nil
								},
							},
							{
								Name:  "fulfill-intent",
								Usage: "fulfill-intent sessions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fulfill-intent...")
									return nil
								},
							},
							{
								Name:  "submit-answer-feedback",
								Usage: "submit-answer-feedback sessions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing submit-answer-feedback...")
									return nil
								},
							},
						},
					},
					{
						Name:  "test-cases",
						Usage: "Manage test-cases resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListTestCasesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTestCases(ctx, req)
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
								Name:  "batch-delete",
								Usage: "batch-delete test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing batch-delete on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "test_case", Usage: "The test_case.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test_case"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetTestCaseRequest{Name: name}
									resp, err := client.GetTestCase(ctx, req)
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
								Usage: "create test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "notes", Usage: "The notes.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateTestCaseRequest{Parent: parent}
									req.TestCase = &cxpb.TestCase{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Notes:       cmd.String("notes"),
									}
									resp, err := client.CreateTestCase(ctx, req)
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
								Usage: "update test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "test_case", Usage: "The test_case.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "notes", Usage: "The notes.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test_case"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateTestCaseRequest{}
									req.TestCase = &cxpb.TestCase{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Notes:       cmd.String("notes"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("notes") {
										paths = append(paths, "notes")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateTestCase(ctx, req)
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
								Name:  "run",
								Usage: "run test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "test_case", Usage: "The test_case.", Required: true},
									&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/testCases/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("test_case"))
									client, err := cx.NewTestCasesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.RunTestCaseRequest{Name: name}
									req.Environment = cmd.String("environment")
									op, err := client.RunTestCase(ctx, req)
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
								Name:  "batch-run",
								Usage: "batch-run test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing batch-run on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "calculate-coverage",
								Usage: "calculate-coverage test-cases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing calculate-coverage...")
									return nil
								},
							},
							{
								Name:  "import",
								Usage: "import test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "export",
								Usage: "export test-cases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									fmt.Printf("Executing export on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "tools",
						Usage: "Manage tools resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create tools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateToolRequest{Parent: parent}
									req.Tool = &cxpb.Tool{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreateTool(ctx, req)
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
								Usage: "list tools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListToolsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTools(ctx, req)
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
								Usage: "describe tools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetToolRequest{Name: name}
									resp, err := client.GetTool(ctx, req)
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
								Usage: "update tools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateToolRequest{}
									req.Tool = &cxpb.Tool{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
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
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateTool(ctx, req)
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
								Usage: "delete tools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteToolRequest{Name: name}
									if err := client.DeleteTool(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "transition-route-groups",
						Usage: "Manage transition-route-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list transition-route-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListTransitionRouteGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTransitionRouteGroups(ctx, req)
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
								Usage: "describe transition-route-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "transition_route_group", Usage: "The transition_route_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/transitionRouteGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("transition_route_group"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetTransitionRouteGroupRequest{Name: name}
									resp, err := client.GetTransitionRouteGroup(ctx, req)
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
								Usage: "create transition-route-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateTransitionRouteGroupRequest{Parent: parent}
									req.LanguageCode = cmd.String("language-code")
									req.TransitionRouteGroup = &cxpb.TransitionRouteGroup{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateTransitionRouteGroup(ctx, req)
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
								Usage: "update transition-route-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "transition_route_group", Usage: "The transition_route_group.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/transitionRouteGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("transition_route_group"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateTransitionRouteGroupRequest{}
									req.TransitionRouteGroup = &cxpb.TransitionRouteGroup{
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
									resp, err := client.UpdateTransitionRouteGroup(ctx, req)
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
								Usage: "delete transition-route-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "transition_route_group", Usage: "The transition_route_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/transitionRouteGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("transition_route_group"))
									client, err := cx.NewTransitionRouteGroupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteTransitionRouteGroupRequest{Name: name}
									if err := client.DeleteTransitionRouteGroup(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "validation-result",
						Usage: "Manage validation-result resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe validation-result",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/validationResult", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewAgentsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetAgentValidationResultRequest{Name: name}
									resp, err := client.GetAgentValidationResult(ctx, req)
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
								Usage: "describe validation-result",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/validationResult", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewFlowsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetFlowValidationResultRequest{Name: name}
									resp, err := client.GetFlowValidationResult(ctx, req)
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
						Name:  "versions",
						Usage: "Manage versions resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreatePlaybookVersionRequest{Parent: parent}
									req.PlaybookVersion = &cxpb.PlaybookVersion{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreatePlaybookVersion(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("version"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetPlaybookVersionRequest{Name: name}
									resp, err := client.GetPlaybookVersion(ctx, req)
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
								Name:  "restore",
								Usage: "restore versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("version"))
									fmt.Printf("Executing restore on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListPlaybookVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPlaybookVersions(ctx, req)
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
								Usage: "delete versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "playbook", Usage: "The playbook.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("playbook"), cmd.String("version"))
									client, err := cx.NewPlaybooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeletePlaybookVersionRequest{Name: name}
									if err := client.DeletePlaybookVersion(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListToolVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListToolVersions(ctx, req)
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
								Usage: "create versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateToolVersionRequest{Parent: parent}
									req.ToolVersion = &cxpb.ToolVersion{
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateToolVersion(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"), cmd.String("version"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetToolVersionRequest{Name: name}
									resp, err := client.GetToolVersion(ctx, req)
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
								Usage: "delete versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"), cmd.String("version"))
									client, err := cx.NewToolsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteToolVersionRequest{Name: name}
									if err := client.DeleteToolVersion(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "restore",
								Usage: "restore versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "tool", Usage: "The tool.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("tool"), cmd.String("version"))
									fmt.Printf("Executing restore on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewVersionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVersions(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
									client, err := cx.NewVersionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetVersionRequest{Name: name}
									resp, err := client.GetVersion(ctx, req)
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
								Usage: "create versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"))
									client, err := cx.NewVersionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateVersionRequest{Parent: parent}
									req.Version = &cxpb.Version{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateVersion(ctx, req)
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
								Usage: "update versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
									client, err := cx.NewVersionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateVersionRequest{}
									req.Version = &cxpb.Version{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
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
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateVersion(ctx, req)
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
								Usage: "delete versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
									client, err := cx.NewVersionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteVersionRequest{Name: name}
									if err := client.DeleteVersion(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "load",
								Usage: "load versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "flow", Usage: "The flow.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
									&cli.BoolFlag{Name: "allow-override-agent-resources", Usage: "The allow override agent resources.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("flow"), cmd.String("version"))
									client, err := cx.NewVersionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.LoadVersionRequest{Name: name}
									req.AllowOverrideAgentResources = cmd.Bool("allow-override-agent-resources")
									op, err := client.LoadVersion(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("LoadVersion completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "compare-versions",
								Usage: "compare-versions versions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing compare-versions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "webhooks",
						Usage: "Manage webhooks resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list webhooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewWebhooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &cxpb.ListWebhooksRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWebhooks(ctx, req)
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
								Usage: "describe webhooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "webhook", Usage: "The webhook.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("webhook"))
									client, err := cx.NewWebhooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.GetWebhookRequest{Name: name}
									resp, err := client.GetWebhook(ctx, req)
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
								Usage: "create webhooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/agents/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"))
									client, err := cx.NewWebhooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.CreateWebhookRequest{Parent: parent}
									req.Webhook = &cxpb.Webhook{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Disabled:    cmd.Bool("disabled"),
									}
									resp, err := client.CreateWebhook(ctx, req)
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
								Usage: "update webhooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "webhook", Usage: "The webhook.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("webhook"))
									client, err := cx.NewWebhooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.UpdateWebhookRequest{}
									req.Webhook = &cxpb.Webhook{
										Name:        name,
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Disabled:    cmd.Bool("disabled"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateWebhook(ctx, req)
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
								Usage: "delete webhooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "agent", Usage: "The agent.", Required: true},
									&cli.StringFlag{Name: "webhook", Usage: "The webhook.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("agent"), cmd.String("webhook"))
									client, err := cx.NewWebhooksClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &cxpb.DeleteWebhookRequest{Name: name}
									if err := client.DeleteWebhook(ctx, req); err != nil {
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
