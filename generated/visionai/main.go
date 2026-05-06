package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	visionai "cloud.google.com/go/visionai/apiv1"
	"cloud.google.com/go/visionai/apiv1/visionaipb"
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
				Name:  "visionai",
				Usage: "manage Vision AI API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "analyses",
						Usage: "Manage analyses resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list analyses",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListAnalysesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAnalyses(ctx, req)
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
								Usage: "describe analyses",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("analysis"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetAnalysisRequest{Name: name}
									resp, err := client.GetAnalysis(ctx, req)
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
								Usage: "create analyses",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "analysis-id", Usage: "The analysis id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.BoolFlag{Name: "disable-event-watch", Usage: "The disable event watch.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateAnalysisRequest{Parent: parent}
									req.AnalysisId = cmd.String("analysis-id")
									req.Analysis = &visionaipb.Analysis{
										Name:              cmd.String("name"),
										DisableEventWatch: cmd.Bool("disable-event-watch"),
									}
									op, err := client.CreateAnalysis(ctx, req)
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
								Usage: "update analyses",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.BoolFlag{Name: "disable-event-watch", Usage: "The disable event watch.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("analysis"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateAnalysisRequest{}
									req.Analysis = &visionaipb.Analysis{
										Name:              name,
										Name:              cmd.String("name"),
										DisableEventWatch: cmd.Bool("disable-event-watch"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("disable-event-watch") {
										paths = append(paths, "disable_event_watch")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAnalysis(ctx, req)
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
								Usage: "delete analyses",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("analysis"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteAnalysisRequest{Name: name}
									op, err := client.DeleteAnalysis(ctx, req)
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
						Name:  "annotations",
						Usage: "Manage annotations resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create annotations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "annotation-id", Usage: "The annotation id.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateAnnotationRequest{Parent: parent}
									req.AnnotationId = cmd.String("annotation-id")
									req.Annotation = &visionaipb.Annotation{
										Name: cmd.String("name"),
									}
									resp, err := client.CreateAnnotation(ctx, req)
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
								Usage: "describe annotations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "annotation", Usage: "The annotation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s/annotations/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"), cmd.String("annotation"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetAnnotationRequest{Name: name}
									resp, err := client.GetAnnotation(ctx, req)
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
								Usage: "list annotations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListAnnotationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAnnotations(ctx, req)
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
								Usage: "update annotations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "annotation", Usage: "The annotation.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s/annotations/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"), cmd.String("annotation"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateAnnotationRequest{}
									req.Annotation = &visionaipb.Annotation{
										Name: name,
										Name: cmd.String("name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateAnnotation(ctx, req)
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
								Usage: "delete annotations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "annotation", Usage: "The annotation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s/annotations/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"), cmd.String("annotation"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteAnnotationRequest{Name: name}
									if err := client.DeleteAnnotation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "applications",
						Usage: "Manage applications resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListApplicationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListApplications(ctx, req)
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
								Usage: "describe applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetApplicationRequest{Name: name}
									resp, err := client.GetApplication(ctx, req)
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
								Usage: "create applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application-id", Usage: "The application id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateApplicationRequest{Parent: parent}
									req.ApplicationId = cmd.String("application-id")
									req.Application = &visionaipb.Application{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateApplication(ctx, req)
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
								Usage: "update applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateApplicationRequest{}
									req.Application = &visionaipb.Application{
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
									op, err := client.UpdateApplication(ctx, req)
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
								Usage: "delete applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteApplicationRequest{Name: name}
									op, err := client.DeleteApplication(ctx, req)
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
								Name:  "deploy",
								Usage: "deploy applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.BoolFlag{Name: "enable-monitoring", Usage: "The enable monitoring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeployApplicationRequest{Name: name}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.EnableMonitoring = cmd.Bool("enable-monitoring")
									op, err := client.DeployApplication(ctx, req)
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
								Name:  "undeploy",
								Usage: "undeploy applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UndeployApplicationRequest{Name: name}
									op, err := client.UndeployApplication(ctx, req)
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
								Name:  "add-stream-input",
								Usage: "add-stream-input applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.AddApplicationStreamInputRequest{Name: name}
									op, err := client.AddApplicationStreamInput(ctx, req)
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
								Name:  "remove-stream-input",
								Usage: "remove-stream-input applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.RemoveApplicationStreamInputRequest{Name: name}
									op, err := client.RemoveApplicationStreamInput(ctx, req)
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
								Name:  "update-stream-input",
								Usage: "update-stream-input applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.BoolFlag{Name: "allow-missing", Usage: "The allow missing.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateApplicationStreamInputRequest{Name: name}
									req.AllowMissing = cmd.Bool("allow-missing")
									op, err := client.UpdateApplicationStreamInput(ctx, req)
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
								Name:  "create",
								Usage: "create applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateApplicationInstancesRequest{Parent: parent}
									req.Name = cmd.String("name")
									op, err := client.CreateApplicationInstances(ctx, req)
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
								Usage: "delete applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteApplicationInstancesRequest{Name: name}
									op, err := client.DeleteApplicationInstances(ctx, req)
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
								Name:  "update-application-instances",
								Usage: "update-application-instances applications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.BoolFlag{Name: "allow-missing", Usage: "The allow missing.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateApplicationInstancesRequest{Name: name}
									req.AllowMissing = cmd.Bool("allow-missing")
									op, err := client.UpdateApplicationInstances(ctx, req)
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
						Name:  "assets",
						Usage: "Manage assets resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset-id", Usage: "The asset id.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateAssetRequest{Parent: parent}
									req.AssetId = cmd.String("asset-id")
									req.Asset = &visionaipb.Asset{
										Name: cmd.String("name"),
									}
									resp, err := client.CreateAsset(ctx, req)
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
								Usage: "update assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateAssetRequest{}
									req.Asset = &visionaipb.Asset{
										Name: name,
										Name: cmd.String("name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateAsset(ctx, req)
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
								Usage: "describe assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetAssetRequest{Name: name}
									resp, err := client.GetAsset(ctx, req)
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
								Usage: "list assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListAssetsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAssets(ctx, req)
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
								Usage: "delete assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteAssetRequest{Name: name}
									op, err := client.DeleteAsset(ctx, req)
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
								Name:  "upload",
								Usage: "upload assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UploadAssetRequest{Name: name}
									op, err := client.UploadAsset(ctx, req)
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
								Name:  "generate-retrieval-url",
								Usage: "generate-retrieval-url assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									fmt.Printf("Executing generate-retrieval-url on %s\n", name)
									return nil
								},
							},
							{
								Name:  "analyze",
								Usage: "analyze assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.AnalyzeAssetRequest{Name: name}
									op, err := client.AnalyzeAsset(ctx, req)
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
								Name:  "index",
								Usage: "index assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "index", Usage: "The index.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.IndexAssetRequest{Name: name}
									req.Index = cmd.String("index")
									op, err := client.IndexAsset(ctx, req)
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
								Name:  "remove-index",
								Usage: "remove-index assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
									&cli.StringFlag{Name: "index", Usage: "The index.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.RemoveIndexAssetRequest{Name: name}
									req.Index = cmd.String("index")
									op, err := client.RemoveIndexAsset(ctx, req)
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
								Name:  "clip",
								Usage: "clip assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									fmt.Printf("Executing clip on %s\n", name)
									return nil
								},
							},
							{
								Name:  "generate-hls-uri",
								Usage: "generate-hls-uri assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "asset", Usage: "The asset.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/assets/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("asset"))
									fmt.Printf("Executing generate-hls-uri on %s\n", name)
									return nil
								},
							},
							{
								Name:  "import",
								Usage: "import assets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project_number"), cmd.String("location"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "channels",
						Usage: "Manage channels resources",
						Commands: []*cli.Command{
							{
								Name:  "materialize-channel",
								Usage: "materialize-channel channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									fmt.Printf("Executing materialize-channel on %s\n", parent)
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
								Name:  "health-check",
								Usage: "health-check clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing health-check...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list clusters",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe clusters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetClusterRequest{Name: name}
									resp, err := client.GetCluster(ctx, req)
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
								Usage: "create clusters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster-id", Usage: "The cluster id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateClusterRequest{Parent: parent}
									req.ClusterId = cmd.String("cluster-id")
									op, err := client.CreateCluster(ctx, req)
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
								Usage: "update clusters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateClusterRequest{}
									req.Cluster = &visionaipb.Cluster{
										Name: name,
									}
									op, err := client.UpdateCluster(ctx, req)
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
								Usage: "delete clusters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteClusterRequest{Name: name}
									op, err := client.DeleteCluster(ctx, req)
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
						Name:  "collections",
						Usage: "Manage collections resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create collections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "collection-id", Usage: "The collection id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateCollectionRequest{Parent: parent}
									req.CollectionId = cmd.String("collection-id")
									req.Collection = &visionaipb.Collection{
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateCollection(ctx, req)
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
								Usage: "delete collections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "collection", Usage: "The collection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("collection"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteCollectionRequest{Name: name}
									op, err := client.DeleteCollection(ctx, req)
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
								Name:  "describe",
								Usage: "describe collections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "collection", Usage: "The collection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("collection"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetCollectionRequest{Name: name}
									resp, err := client.GetCollection(ctx, req)
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
								Usage: "update collections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "collection", Usage: "The collection.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/collections/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("collection"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateCollectionRequest{}
									req.Collection = &visionaipb.Collection{
										Name:        name,
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateCollection(ctx, req)
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
								Usage: "list collections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListCollectionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCollections(ctx, req)
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
								Name:  "add-collection-item",
								Usage: "add-collection-item collections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-collection-item...")
									return nil
								},
							},
							{
								Name:  "remove-collection-item",
								Usage: "remove-collection-item collections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-collection-item...")
									return nil
								},
							},
							{
								Name:  "view-collection-items",
								Usage: "view-collection-items collections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing view-collection-items...")
									return nil
								},
							},
						},
					},
					{
						Name:  "corpora",
						Usage: "Manage corpora resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create corpora",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project_number"), cmd.String("location"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateCorpusRequest{Parent: parent}
									req.Corpus = &visionaipb.Corpus{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateCorpus(ctx, req)
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
								Usage: "describe corpora",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetCorpusRequest{Name: name}
									resp, err := client.GetCorpus(ctx, req)
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
								Usage: "update corpora",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateCorpusRequest{}
									req.Corpus = &visionaipb.Corpus{
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
									resp, err := client.UpdateCorpus(ctx, req)
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
								Usage: "list corpora",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete corpora",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteCorpusRequest{Name: name}
									if err := client.DeleteCorpus(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "analyze",
								Usage: "analyze corpora",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.AnalyzeCorpusRequest{Name: name}
									op, err := client.AnalyzeCorpus(ctx, req)
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
								Name:  "search-assets",
								Usage: "search-assets corpora",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search-assets...")
									return nil
								},
							},
						},
					},
					{
						Name:  "data-schemas",
						Usage: "Manage data-schemas resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create data-schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateDataSchemaRequest{Parent: parent}
									req.DataSchema = &visionaipb.DataSchema{
										Name: cmd.String("name"),
										Key:  cmd.String("key"),
									}
									resp, err := client.CreateDataSchema(ctx, req)
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
								Usage: "update data-schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "data_schema", Usage: "The data_schema.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/dataSchemas/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("data_schema"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateDataSchemaRequest{}
									req.DataSchema = &visionaipb.DataSchema{
										Name: name,
										Name: cmd.String("name"),
										Key:  cmd.String("key"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("key") {
										paths = append(paths, "key")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateDataSchema(ctx, req)
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
								Usage: "describe data-schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "data_schema", Usage: "The data_schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/dataSchemas/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("data_schema"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetDataSchemaRequest{Name: name}
									resp, err := client.GetDataSchema(ctx, req)
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
								Usage: "delete data-schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "data_schema", Usage: "The data_schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/dataSchemas/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("data_schema"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteDataSchemaRequest{Name: name}
									if err := client.DeleteDataSchema(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list data-schemas",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListDataSchemasRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDataSchemas(ctx, req)
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
						Name:  "drafts",
						Usage: "Manage drafts resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list drafts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListDraftsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDrafts(ctx, req)
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
								Usage: "describe drafts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.StringFlag{Name: "draft", Usage: "The draft.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/drafts/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("draft"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetDraftRequest{Name: name}
									resp, err := client.GetDraft(ctx, req)
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
								Usage: "create drafts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.StringFlag{Name: "draft-id", Usage: "The draft id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateDraftRequest{Parent: parent}
									req.DraftId = cmd.String("draft-id")
									req.Draft = &visionaipb.Draft{
										Name:        cmd.String("name"),
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateDraft(ctx, req)
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
								Usage: "update drafts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.StringFlag{Name: "draft", Usage: "The draft.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/drafts/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("draft"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateDraftRequest{}
									req.Draft = &visionaipb.Draft{
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
									op, err := client.UpdateDraft(ctx, req)
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
								Usage: "delete drafts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.StringFlag{Name: "draft", Usage: "The draft.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/drafts/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("draft"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteDraftRequest{Name: name}
									op, err := client.DeleteDraft(ctx, req)
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
						Name:  "events",
						Usage: "Manage events resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list events",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListEventsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListEvents(ctx, req)
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
								Usage: "describe events",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "event", Usage: "The event.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/events/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("event"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetEventRequest{Name: name}
									resp, err := client.GetEvent(ctx, req)
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
								Usage: "create events",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "event-id", Usage: "The event id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateEventRequest{Parent: parent}
									req.EventId = cmd.String("event-id")
									req.Event = &visionaipb.Event{
										Name: cmd.String("name"),
									}
									op, err := client.CreateEvent(ctx, req)
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
								Usage: "update events",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "event", Usage: "The event.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/events/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("event"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateEventRequest{}
									req.Event = &visionaipb.Event{
										Name: name,
										Name: cmd.String("name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateEvent(ctx, req)
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
								Usage: "delete events",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "event", Usage: "The event.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/events/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("event"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteEventRequest{Name: name}
									op, err := client.DeleteEvent(ctx, req)
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
						Name:  "index-endpoints",
						Usage: "Manage index-endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "search-index-endpoint",
								Usage: "search-index-endpoint index-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search-index-endpoint...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create index-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "index-endpoint-id", Usage: "The index endpoint id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateIndexEndpointRequest{Parent: parent}
									req.IndexEndpointId = cmd.String("index-endpoint-id")
									req.IndexEndpoint = &visionaipb.IndexEndpoint{
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateIndexEndpoint(ctx, req)
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
								Usage: "describe index-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "index_endpoint", Usage: "The index_endpoint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index_endpoint"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetIndexEndpointRequest{Name: name}
									resp, err := client.GetIndexEndpoint(ctx, req)
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
								Usage: "list index-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update index-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "index_endpoint", Usage: "The index_endpoint.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index_endpoint"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateIndexEndpointRequest{}
									req.IndexEndpoint = &visionaipb.IndexEndpoint{
										Name:        name,
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateIndexEndpoint(ctx, req)
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
								Usage: "delete index-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "index_endpoint", Usage: "The index_endpoint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("index_endpoint"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteIndexEndpointRequest{Name: name}
									op, err := client.DeleteIndexEndpoint(ctx, req)
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
								Name:  "deploy-index",
								Usage: "deploy-index index-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing deploy-index...")
									return nil
								},
							},
							{
								Name:  "undeploy-index",
								Usage: "undeploy-index index-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing undeploy-index...")
									return nil
								},
							},
						},
					},
					{
						Name:  "indexes",
						Usage: "Manage indexes resources",
						Commands: []*cli.Command{
							{
								Name:  "view-assets",
								Usage: "view-assets indexes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing view-assets...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create indexes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "index-id", Usage: "The index id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateIndexRequest{Parent: parent}
									req.IndexId = cmd.String("index-id")
									req.Index = &visionaipb.Index{
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateIndex(ctx, req)
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
								Usage: "update indexes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "index", Usage: "The index.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("index"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateIndexRequest{}
									req.Index = &visionaipb.Index{
										Name:        name,
										DisplayName: cmd.String("display-name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateIndex(ctx, req)
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
								Usage: "describe indexes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "index", Usage: "The index.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("index"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetIndexRequest{Name: name}
									resp, err := client.GetIndex(ctx, req)
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
								Usage: "list indexes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListIndexesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListIndexes(ctx, req)
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
								Usage: "delete indexes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "index", Usage: "The index.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/indexes/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("index"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteIndexRequest{Name: name}
									op, err := client.DeleteIndex(ctx, req)
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
						Name:  "instances",
						Usage: "Manage instances resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/applications/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListInstancesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInstances(ctx, req)
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
								Usage: "describe instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "application", Usage: "The application.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/applications/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("application"), cmd.String("instance"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetInstanceRequest{Name: name}
									resp, err := client.GetInstance(ctx, req)
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
									client, err := visionai.NewHealthCheckClient(ctx)
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
									client, err := visionai.NewHealthCheckClient(ctx)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "resolve-operator-info",
								Usage: "resolve-operator-info locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resolve-operator-info...")
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
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
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
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
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
									client, err := visionai.NewAppPlatformClient(ctx)
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
									client, err := visionai.NewAppPlatformClient(ctx)
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
									client, err := visionai.NewStreamingClient(ctx)
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
									client, err := visionai.NewStreamingClient(ctx)
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
									client, err := visionai.NewStreamsClient(ctx)
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
									client, err := visionai.NewStreamsClient(ctx)
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
									client, err := visionai.NewWarehouseClient(ctx)
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
									client, err := visionai.NewWarehouseClient(ctx)
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
									client, err := visionai.NewHealthCheckClient(ctx)
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
									client, err := visionai.NewHealthCheckClient(ctx)
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
									client, err := visionai.NewHealthCheckClient(ctx)
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
									client, err := visionai.NewHealthCheckClient(ctx)
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
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
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
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
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
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
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
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
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
									client, err := visionai.NewAppPlatformClient(ctx)
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
									client, err := visionai.NewAppPlatformClient(ctx)
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
									client, err := visionai.NewAppPlatformClient(ctx)
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
									client, err := visionai.NewAppPlatformClient(ctx)
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
									client, err := visionai.NewStreamingClient(ctx)
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
									client, err := visionai.NewStreamingClient(ctx)
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
									client, err := visionai.NewStreamingClient(ctx)
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
									client, err := visionai.NewStreamingClient(ctx)
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
									client, err := visionai.NewStreamsClient(ctx)
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
									client, err := visionai.NewStreamsClient(ctx)
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
									client, err := visionai.NewStreamsClient(ctx)
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
									client, err := visionai.NewStreamsClient(ctx)
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
									client, err := visionai.NewWarehouseClient(ctx)
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
									client, err := visionai.NewWarehouseClient(ctx)
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
									client, err := visionai.NewWarehouseClient(ctx)
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
									client, err := visionai.NewWarehouseClient(ctx)
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
						Name:  "operators",
						Usage: "Manage operators resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list operators",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe operators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operator", Usage: "The operator.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operators/%s", cmd.String("project"), cmd.String("location"), cmd.String("operator"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetOperatorRequest{Name: name}
									resp, err := client.GetOperator(ctx, req)
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
								Usage: "create operators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operator-id", Usage: "The operator id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "docker-image", Usage: "The docker image.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateOperatorRequest{Parent: parent}
									req.OperatorId = cmd.String("operator-id")
									req.Operator = &visionaipb.Operator{
										Name:        cmd.String("name"),
										DockerImage: cmd.String("docker-image"),
									}
									op, err := client.CreateOperator(ctx, req)
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
								Usage: "update operators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operator", Usage: "The operator.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "docker-image", Usage: "The docker image.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operators/%s", cmd.String("project"), cmd.String("location"), cmd.String("operator"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateOperatorRequest{}
									req.Operator = &visionaipb.Operator{
										Name:        name,
										Name:        cmd.String("name"),
										DockerImage: cmd.String("docker-image"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("docker-image") {
										paths = append(paths, "docker_image")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateOperator(ctx, req)
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
								Usage: "delete operators",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operator", Usage: "The operator.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operators/%s", cmd.String("project"), cmd.String("location"), cmd.String("operator"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteOperatorRequest{Name: name}
									op, err := client.DeleteOperator(ctx, req)
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
						Name:  "processes",
						Usage: "Manage processes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list processes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListProcessesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListProcesses(ctx, req)
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
								Usage: "describe processes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "process", Usage: "The process.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/processes/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("process"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetProcessRequest{Name: name}
									resp, err := client.GetProcess(ctx, req)
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
								Usage: "create processes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "process-id", Usage: "The process id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: true},
									&cli.StringFlag{Name: "event-id", Usage: "The event id.", Required: false},
									&cli.StringFlag{Name: "batch-id", Usage: "The batch id.", Required: false},
									&cli.IntFlag{Name: "retry-count", Usage: "The retry count.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateProcessRequest{Parent: parent}
									req.ProcessId = cmd.String("process-id")
									req.Process = &visionaipb.Process{
										Name:       cmd.String("name"),
										Analysis:   cmd.String("analysis"),
										EventId:    cmd.String("event-id"),
										BatchId:    cmd.String("batch-id"),
										RetryCount: int32(cmd.Int("retry-count")),
									}
									op, err := client.CreateProcess(ctx, req)
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
								Usage: "update processes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "process", Usage: "The process.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: false},
									&cli.StringFlag{Name: "event-id", Usage: "The event id.", Required: false},
									&cli.StringFlag{Name: "batch-id", Usage: "The batch id.", Required: false},
									&cli.IntFlag{Name: "retry-count", Usage: "The retry count.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/processes/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("process"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateProcessRequest{}
									req.Process = &visionaipb.Process{
										Name:       name,
										Name:       cmd.String("name"),
										Analysis:   cmd.String("analysis"),
										EventId:    cmd.String("event-id"),
										BatchId:    cmd.String("batch-id"),
										RetryCount: int32(cmd.Int("retry-count")),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("analysis") {
										paths = append(paths, "analysis")
									}
									if cmd.IsSet("event-id") {
										paths = append(paths, "event_id")
									}
									if cmd.IsSet("batch-id") {
										paths = append(paths, "batch_id")
									}
									if cmd.IsSet("retry-count") {
										paths = append(paths, "retry_count")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateProcess(ctx, req)
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
								Usage: "delete processes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "process", Usage: "The process.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/processes/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("process"))
									client, err := visionai.NewLiveVideoAnalyticsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteProcessRequest{Name: name}
									op, err := client.DeleteProcess(ctx, req)
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
								Name:  "batch-run",
								Usage: "batch-run processes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing batch-run on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "processors",
						Usage: "Manage processors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list processors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListProcessorsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListProcessors(ctx, req)
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
								Name:  "prebuilt",
								Usage: "prebuilt processors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									fmt.Printf("Executing prebuilt on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe processors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "processor", Usage: "The processor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetProcessorRequest{Name: name}
									resp, err := client.GetProcessor(ctx, req)
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
								Usage: "create processors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "processor-id", Usage: "The processor id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "supports-post-processing", Usage: "The supports post processing.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateProcessorRequest{Parent: parent}
									req.ProcessorId = cmd.String("processor-id")
									req.Processor = &visionaipb.Processor{
										Name:                   cmd.String("name"),
										DisplayName:            cmd.String("display-name"),
										Description:            cmd.String("description"),
										SupportsPostProcessing: cmd.Bool("supports-post-processing"),
									}
									op, err := client.CreateProcessor(ctx, req)
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
								Usage: "update processors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "processor", Usage: "The processor.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "supports-post-processing", Usage: "The supports post processing.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateProcessorRequest{}
									req.Processor = &visionaipb.Processor{
										Name:                   name,
										Name:                   cmd.String("name"),
										DisplayName:            cmd.String("display-name"),
										Description:            cmd.String("description"),
										SupportsPostProcessing: cmd.Bool("supports-post-processing"),
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
									if cmd.IsSet("supports-post-processing") {
										paths = append(paths, "supports_post_processing")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateProcessor(ctx, req)
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
								Usage: "delete processors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "processor", Usage: "The processor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/processors/%s", cmd.String("project"), cmd.String("location"), cmd.String("processor"))
									client, err := visionai.NewAppPlatformClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteProcessorRequest{Name: name}
									op, err := client.DeleteProcessor(ctx, req)
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
						Name:  "search-configs",
						Usage: "Manage search-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create search-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search-config-id", Usage: "The search config id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateSearchConfigRequest{Parent: parent}
									req.SearchConfigId = cmd.String("search-config-id")
									req.SearchConfig = &visionaipb.SearchConfig{
										Name: cmd.String("name"),
									}
									resp, err := client.CreateSearchConfig(ctx, req)
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
								Usage: "update search-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search_config", Usage: "The search_config.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchConfigs/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("search_config"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateSearchConfigRequest{}
									req.SearchConfig = &visionaipb.SearchConfig{
										Name: name,
										Name: cmd.String("name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSearchConfig(ctx, req)
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
								Usage: "describe search-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search_config", Usage: "The search_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchConfigs/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("search_config"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetSearchConfigRequest{Name: name}
									resp, err := client.GetSearchConfig(ctx, req)
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
								Usage: "delete search-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search_config", Usage: "The search_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchConfigs/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("search_config"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteSearchConfigRequest{Name: name}
									if err := client.DeleteSearchConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list search-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListSearchConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSearchConfigs(ctx, req)
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
						Name:  "search-hypernyms",
						Usage: "Manage search-hypernyms resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create search-hypernyms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search-hypernym-id", Usage: "The search hypernym id.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "hypernym", Usage: "The hypernym.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateSearchHypernymRequest{Parent: parent}
									req.SearchHypernymId = cmd.String("search-hypernym-id")
									req.SearchHypernym = &visionaipb.SearchHypernym{
										Name:     cmd.String("name"),
										Hypernym: cmd.String("hypernym"),
									}
									resp, err := client.CreateSearchHypernym(ctx, req)
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
								Usage: "update search-hypernyms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search_hypernym", Usage: "The search_hypernym.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "hypernym", Usage: "The hypernym.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchHypernyms/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("search_hypernym"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateSearchHypernymRequest{}
									req.SearchHypernym = &visionaipb.SearchHypernym{
										Name:     name,
										Name:     cmd.String("name"),
										Hypernym: cmd.String("hypernym"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("hypernym") {
										paths = append(paths, "hypernym")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSearchHypernym(ctx, req)
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
								Usage: "describe search-hypernyms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search_hypernym", Usage: "The search_hypernym.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchHypernyms/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("search_hypernym"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetSearchHypernymRequest{Name: name}
									resp, err := client.GetSearchHypernym(ctx, req)
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
								Usage: "delete search-hypernyms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.StringFlag{Name: "search_hypernym", Usage: "The search_hypernym.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/corpora/%s/searchHypernyms/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"), cmd.String("search_hypernym"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteSearchHypernymRequest{Name: name}
									if err := client.DeleteSearchHypernym(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list search-hypernyms",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "project_number", Usage: "The project_number.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "corpus", Usage: "The corpus.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/corpora/%s", cmd.String("project_number"), cmd.String("location"), cmd.String("corpus"))
									client, err := visionai.NewWarehouseClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListSearchHypernymsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSearchHypernyms(ctx, req)
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
						Name:  "series",
						Usage: "Manage series resources",
						Commands: []*cli.Command{
							{
								Name:  "acquire-lease",
								Usage: "acquire-lease series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing acquire-lease...")
									return nil
								},
							},
							{
								Name:  "renew-lease",
								Usage: "renew-lease series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing renew-lease...")
									return nil
								},
							},
							{
								Name:  "release-lease",
								Usage: "release-lease series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing release-lease...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list series",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListSeriesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSeries(ctx, req)
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
								Usage: "describe series",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "series", Usage: "The series.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("series"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetSeriesRequest{Name: name}
									resp, err := client.GetSeries(ctx, req)
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
								Usage: "create series",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "series-id", Usage: "The series id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "stream", Usage: "The stream.", Required: true},
									&cli.StringFlag{Name: "event", Usage: "The event.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateSeriesRequest{Parent: parent}
									req.SeriesId = cmd.String("series-id")
									req.Series = &visionaipb.Series{
										Name:   cmd.String("name"),
										Stream: cmd.String("stream"),
										Event:  cmd.String("event"),
									}
									op, err := client.CreateSeries(ctx, req)
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
								Usage: "update series",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "series", Usage: "The series.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "stream", Usage: "The stream.", Required: false},
									&cli.StringFlag{Name: "event", Usage: "The event.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("series"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateSeriesRequest{}
									req.Series = &visionaipb.Series{
										Name:   name,
										Name:   cmd.String("name"),
										Stream: cmd.String("stream"),
										Event:  cmd.String("event"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("stream") {
										paths = append(paths, "stream")
									}
									if cmd.IsSet("event") {
										paths = append(paths, "event")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateSeries(ctx, req)
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
								Usage: "delete series",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "series", Usage: "The series.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/series/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("series"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteSeriesRequest{Name: name}
									op, err := client.DeleteSeries(ctx, req)
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
						Name:  "streams",
						Usage: "Manage streams resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list streams",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &visionaipb.ListStreamsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListStreams(ctx, req)
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
								Usage: "describe streams",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "stream", Usage: "The stream.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.GetStreamRequest{Name: name}
									resp, err := client.GetStream(ctx, req)
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
								Usage: "create streams",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "stream-id", Usage: "The stream id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.BoolFlag{Name: "enable-hls-playback", Usage: "The enable hls playback.", Required: false},
									&cli.StringFlag{Name: "media-warehouse-asset", Usage: "The media warehouse asset.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.CreateStreamRequest{Parent: parent}
									req.StreamId = cmd.String("stream-id")
									req.Stream = &visionaipb.Stream{
										Name:                cmd.String("name"),
										DisplayName:         cmd.String("display-name"),
										EnableHlsPlayback:   cmd.Bool("enable-hls-playback"),
										MediaWarehouseAsset: cmd.String("media-warehouse-asset"),
									}
									op, err := client.CreateStream(ctx, req)
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
								Usage: "update streams",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "stream", Usage: "The stream.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.BoolFlag{Name: "enable-hls-playback", Usage: "The enable hls playback.", Required: false},
									&cli.StringFlag{Name: "media-warehouse-asset", Usage: "The media warehouse asset.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.UpdateStreamRequest{}
									req.Stream = &visionaipb.Stream{
										Name:                name,
										Name:                cmd.String("name"),
										DisplayName:         cmd.String("display-name"),
										EnableHlsPlayback:   cmd.Bool("enable-hls-playback"),
										MediaWarehouseAsset: cmd.String("media-warehouse-asset"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("enable-hls-playback") {
										paths = append(paths, "enable_hls_playback")
									}
									if cmd.IsSet("media-warehouse-asset") {
										paths = append(paths, "media_warehouse_asset")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateStream(ctx, req)
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
								Usage: "delete streams",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
									&cli.StringFlag{Name: "stream", Usage: "The stream.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/streams/%s", cmd.String("project"), cmd.String("location"), cmd.String("cluster"), cmd.String("stream"))
									client, err := visionai.NewStreamsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &visionaipb.DeleteStreamRequest{Name: name}
									op, err := client.DeleteStream(ctx, req)
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
								Name:  "get-thumbnail",
								Usage: "get-thumbnail streams",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-thumbnail...")
									return nil
								},
							},
							{
								Name:  "generate-stream-hls-token",
								Usage: "generate-stream-hls-token streams",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-stream-hls-token...")
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
