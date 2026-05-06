package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	deploy "cloud.google.com/go/deploy/apiv1"
	"cloud.google.com/go/deploy/apiv1/deploypb"
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
				Name:  "clouddeploy",
				Usage: "manage Cloud Deploy API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "automation-runs",
						Usage: "Manage automation-runs resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe automation-runs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "automation_run", Usage: "The automation_run.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automationRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("automation_run"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetAutomationRunRequest{Name: name}
									resp, err := client.GetAutomationRun(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "list automation-runs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListAutomationRunsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAutomationRuns(ctx, req)
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
								Name:  "cancel",
								Usage: "cancel automation-runs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "automation_run", Usage: "The automation_run.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automationRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("automation_run"))
									fmt.Printf("Executing cancel on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "automations",
						Usage: "Manage automations resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create automations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "automation-id", Usage: "The automation id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.BoolFlag{Name: "suspended", Usage: "The suspended.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateAutomationRequest{Parent: parent}
									req.AutomationId = cmd.String("automation-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.Automation = &deploypb.Automation{
										Description:    cmd.String("description"),
										Etag:           cmd.String("etag"),
										Suspended:      cmd.Bool("suspended"),
										ServiceAccount: cmd.String("service-account"),
									}
									op, err := client.CreateAutomation(ctx, req)
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
								Usage: "update automations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "automation", Usage: "The automation.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.BoolFlag{Name: "suspended", Usage: "The suspended.", Required: false},
									&cli.StringFlag{Name: "service-account", Usage: "The service account.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("automation"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.UpdateAutomationRequest{}
									req.Automation = &deploypb.Automation{
										Name:           name,
										Description:    cmd.String("description"),
										Etag:           cmd.String("etag"),
										Suspended:      cmd.Bool("suspended"),
										ServiceAccount: cmd.String("service-account"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									if cmd.IsSet("suspended") {
										paths = append(paths, "suspended")
									}
									if cmd.IsSet("service-account") {
										paths = append(paths, "service_account")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAutomation(ctx, req)
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
								Usage: "delete automations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "automation", Usage: "The automation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("automation"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.DeleteAutomationRequest{Name: name}
									op, err := client.DeleteAutomation(ctx, req)
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
								Usage: "describe automations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "automation", Usage: "The automation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("automation"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetAutomationRequest{Name: name}
									resp, err := client.GetAutomation(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "list automations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListAutomationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAutomations(ctx, req)
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
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetConfigRequest{Name: name}
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
						},
					},
					{
						Name:  "custom-target-types",
						Usage: "Manage custom-target-types resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list custom-target-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListCustomTargetTypesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCustomTargetTypes(ctx, req)
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
								Usage: "describe custom-target-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_target_type", Usage: "The custom_target_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_target_type"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetCustomTargetTypeRequest{Name: name}
									resp, err := client.GetCustomTargetType(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "create custom-target-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom-target-type-id", Usage: "The custom target type id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateCustomTargetTypeRequest{Parent: parent}
									req.CustomTargetTypeId = cmd.String("custom-target-type-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.CustomTargetType = &deploypb.CustomTargetType{
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
									}
									op, err := client.CreateCustomTargetType(ctx, req)
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
								Usage: "update custom-target-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_target_type", Usage: "The custom_target_type.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_target_type"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.UpdateCustomTargetTypeRequest{}
									req.CustomTargetType = &deploypb.CustomTargetType{
										Name:        name,
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateCustomTargetType(ctx, req)
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
								Usage: "delete custom-target-types",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_target_type", Usage: "The custom_target_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_target_type"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.DeleteCustomTargetTypeRequest{Name: name}
									op, err := client.DeleteCustomTargetType(ctx, req)
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
						Name:  "delivery-pipelines",
						Usage: "Manage delivery-pipelines resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list delivery-pipelines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListDeliveryPipelinesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDeliveryPipelines(ctx, req)
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
								Usage: "describe delivery-pipelines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetDeliveryPipelineRequest{Name: name}
									resp, err := client.GetDeliveryPipeline(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "create delivery-pipelines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery-pipeline-id", Usage: "The delivery pipeline id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.BoolFlag{Name: "suspended", Usage: "The suspended.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateDeliveryPipelineRequest{Parent: parent}
									req.DeliveryPipelineId = cmd.String("delivery-pipeline-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.DeliveryPipeline = &deploypb.DeliveryPipeline{
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
										Suspended:   cmd.Bool("suspended"),
									}
									op, err := client.CreateDeliveryPipeline(ctx, req)
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
								Usage: "update delivery-pipelines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.BoolFlag{Name: "suspended", Usage: "The suspended.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.UpdateDeliveryPipelineRequest{}
									req.DeliveryPipeline = &deploypb.DeliveryPipeline{
										Name:        name,
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
										Suspended:   cmd.Bool("suspended"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									if cmd.IsSet("suspended") {
										paths = append(paths, "suspended")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateDeliveryPipeline(ctx, req)
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
								Usage: "delete delivery-pipelines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.DeleteDeliveryPipelineRequest{Name: name}
									op, err := client.DeleteDeliveryPipeline(ctx, req)
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
								Name:  "rollback-target",
								Usage: "rollback-target delivery-pipelines",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									fmt.Printf("Executing rollback-target on %s\n", name)
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy delivery-pipelines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy delivery-pipelines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions delivery-pipelines",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "deploy-policies",
						Usage: "Manage deploy-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create deploy-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deploy-policy-id", Usage: "The deploy policy id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "suspended", Usage: "The suspended.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateDeployPolicyRequest{Parent: parent}
									req.DeployPolicyId = cmd.String("deploy-policy-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.DeployPolicy = &deploypb.DeployPolicy{
										Description: cmd.String("description"),
										Suspended:   cmd.Bool("suspended"),
										Etag:        cmd.String("etag"),
									}
									op, err := client.CreateDeployPolicy(ctx, req)
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
								Usage: "update deploy-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deploy_policy", Usage: "The deploy_policy.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "suspended", Usage: "The suspended.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("deploy_policy"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.UpdateDeployPolicyRequest{}
									req.DeployPolicy = &deploypb.DeployPolicy{
										Name:        name,
										Description: cmd.String("description"),
										Suspended:   cmd.Bool("suspended"),
										Etag:        cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("suspended") {
										paths = append(paths, "suspended")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateDeployPolicy(ctx, req)
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
								Usage: "delete deploy-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deploy_policy", Usage: "The deploy_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("deploy_policy"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.DeleteDeployPolicyRequest{Name: name}
									op, err := client.DeleteDeployPolicy(ctx, req)
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
								Name:  "list",
								Usage: "list deploy-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListDeployPoliciesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDeployPolicies(ctx, req)
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
								Usage: "describe deploy-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "deploy_policy", Usage: "The deploy_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("deploy_policy"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetDeployPolicyRequest{Name: name}
									resp, err := client.GetDeployPolicy(ctx, req)
									if err != nil {
										return err
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
						Name:  "job-runs",
						Usage: "Manage job-runs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list job-runs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListJobRunsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListJobRuns(ctx, req)
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
								Usage: "describe job-runs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
									&cli.StringFlag{Name: "job_run", Usage: "The job_run.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s/jobRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"), cmd.String("rollout"), cmd.String("job_run"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetJobRunRequest{Name: name}
									resp, err := client.GetJobRun(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "terminate",
								Usage: "terminate job-runs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
									&cli.StringFlag{Name: "job_run", Usage: "The job_run.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s/jobRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"), cmd.String("rollout"), cmd.String("job_run"))
									fmt.Printf("Executing terminate on %s\n", name)
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
									client, err := deploy.NewCloudDeployClient(ctx)
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
									client, err := deploy.NewCloudDeployClient(ctx)
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
									client, err := deploy.NewCloudDeployClient(ctx)
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
									client, err := deploy.NewCloudDeployClient(ctx)
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
									client, err := deploy.NewCloudDeployClient(ctx)
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
									client, err := deploy.NewCloudDeployClient(ctx)
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
						Name:  "releases",
						Usage: "Manage releases resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListReleasesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListReleases(ctx, req)
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
								Usage: "describe releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetReleaseRequest{Name: name}
									resp, err := client.GetRelease(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "create releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release-id", Usage: "The release id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "skaffold-config-uri", Usage: "The skaffold config uri.", Required: false},
									&cli.StringFlag{Name: "skaffold-config-path", Usage: "The skaffold config path.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "skaffold-version", Usage: "The skaffold version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateReleaseRequest{Parent: parent}
									req.ReleaseId = cmd.String("release-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.Release = &deploypb.Release{
										Description:        cmd.String("description"),
										SkaffoldConfigUri:  cmd.String("skaffold-config-uri"),
										SkaffoldConfigPath: cmd.String("skaffold-config-path"),
										Etag:               cmd.String("etag"),
										SkaffoldVersion:    cmd.String("skaffold-version"),
									}
									op, err := client.CreateRelease(ctx, req)
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
								Name:  "abandon",
								Usage: "abandon releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"))
									fmt.Printf("Executing abandon on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "rollouts",
						Usage: "Manage rollouts resources",
						Commands: []*cli.Command{
							{
								Name:  "approve",
								Usage: "approve rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"), cmd.String("rollout"))
									fmt.Printf("Executing approve on %s\n", name)
									return nil
								},
							},
							{
								Name:  "advance",
								Usage: "advance rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"), cmd.String("rollout"))
									fmt.Printf("Executing advance on %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"), cmd.String("rollout"))
									fmt.Printf("Executing cancel on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListRolloutsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRollouts(ctx, req)
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
								Usage: "describe rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"), cmd.String("rollout"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetRolloutRequest{Name: name}
									resp, err := client.GetRollout(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "create rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "delivery_pipeline", Usage: "The delivery_pipeline.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "rollout-id", Usage: "The rollout id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "starting-phase-id", Usage: "The starting phase id.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "target-id", Usage: "The target id.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery_pipeline"), cmd.String("release"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateRolloutRequest{Parent: parent}
									req.RolloutId = cmd.String("rollout-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.StartingPhaseId = cmd.String("starting-phase-id")
									req.Rollout = &deploypb.Rollout{
										Description: cmd.String("description"),
										TargetId:    cmd.String("target-id"),
										Etag:        cmd.String("etag"),
									}
									op, err := client.CreateRollout(ctx, req)
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
								Name:  "ignore-job",
								Usage: "ignore-job rollouts",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing ignore-job...")
									return nil
								},
							},
							{
								Name:  "retry-job",
								Usage: "retry-job rollouts",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing retry-job...")
									return nil
								},
							},
						},
					},
					{
						Name:  "targets",
						Usage: "Manage targets resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list targets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &deploypb.ListTargetsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTargets(ctx, req)
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
								Usage: "describe targets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target", Usage: "The target.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/targets/%s", cmd.String("project"), cmd.String("location"), cmd.String("target"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.GetTargetRequest{Name: name}
									resp, err := client.GetTarget(ctx, req)
									if err != nil {
										return err
									}
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
								Usage: "create targets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target-id", Usage: "The target id.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "require-approval", Usage: "The require approval.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.CreateTargetRequest{Parent: parent}
									req.TargetId = cmd.String("target-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.Target = &deploypb.Target{
										Description:     cmd.String("description"),
										RequireApproval: cmd.Bool("require-approval"),
										Etag:            cmd.String("etag"),
									}
									op, err := client.CreateTarget(ctx, req)
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
								Usage: "update targets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target", Usage: "The target.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "require-approval", Usage: "The require approval.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/targets/%s", cmd.String("project"), cmd.String("location"), cmd.String("target"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.UpdateTargetRequest{}
									req.Target = &deploypb.Target{
										Name:            name,
										Description:     cmd.String("description"),
										RequireApproval: cmd.Bool("require-approval"),
										Etag:            cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("require-approval") {
										paths = append(paths, "require_approval")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTarget(ctx, req)
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
								Usage: "delete targets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "target", Usage: "The target.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/targets/%s", cmd.String("project"), cmd.String("location"), cmd.String("target"))
									client, err := deploy.NewCloudDeployClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &deploypb.DeleteTargetRequest{Name: name}
									op, err := client.DeleteTarget(ctx, req)
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
