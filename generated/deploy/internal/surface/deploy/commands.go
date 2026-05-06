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

package deploy

import (
	deploy "cloud.google.com/go/deploy/apiv1"
	"cloud.google.com/go/deploy/apiv1/deploypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the clouddeploy command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "clouddeploy",
		Usage: "manage Cloud Deploy API resources",
		Commands: []*cli.Command{
			{
				Name:  "automation-runs",
				Usage: "Manage automation-runs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe automation-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "automation-run", Usage: "The ID of the automation run.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automationRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("automation-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetAutomationRunRequest{
								Name: name,
							}

							resp, err := client.GetAutomationRun(ctx, req)
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
						Usage: "list automation-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter automationRuns to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of automationRuns to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListAutomationRuns` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListAutomationRunsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutomationRuns(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel automation-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "automation-run", Usage: "The ID of the automation run.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automationRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("automation-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CancelAutomationRunRequest{
								Name: name,
							}

							resp, err := client.CancelAutomationRun(ctx, req)
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
				Name:  "automations",
				Usage: "Manage automations resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create automations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "automation-id", Usage: "ID of the `Automation`.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateAutomationRequest{
								Parent:       parent,
								AutomationId: cmd.String("automation-id"),
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateAutomation(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "update",
						Usage: "update automations",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, updating a `Automation` that does not exist will.", Required: false},
							&cli.StringFlag{Name: "automation", Usage: "The ID of the automation.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "automation.name" not yet supported.
							automation_name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("automation"))
							fmt.Printf("Executing update on %s\n", automation_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete automations",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, then deleting an already deleted or non-existing.", Required: false},
							&cli.StringFlag{Name: "automation", Usage: "The ID of the automation.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The weak etag of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and verify whether the resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("automation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAutomation %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.DeleteAutomationRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteAutomation(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe automations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "automation", Usage: "The ID of the automation.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/automations/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("automation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetAutomationRequest{
								Name: name,
							}

							resp, err := client.GetAutomation(ctx, req)
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
						Usage: "list automations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter automations to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of automations to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListAutomations` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListAutomationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutomations(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetConfigRequest{
								Name: name,
							}

							resp, err := client.GetConfig(ctx, req)
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
				Name:  "custom-target-types",
				Usage: "Manage custom-target-types resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list custom-target-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter custom target types to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `CustomTargetType` objects to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListCustomTargetTypes`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListCustomTargetTypesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCustomTargetTypes(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe custom-target-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-target-type", Usage: "The ID of the custom target type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-target-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetCustomTargetTypeRequest{
								Name: name,
							}

							resp, err := client.GetCustomTargetType(ctx, req)
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
						Usage: "create custom-target-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-target-type-id", Usage: "ID of the `CustomTargetType`.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateCustomTargetTypeRequest{
								Parent:             parent,
								CustomTargetTypeId: cmd.String("custom-target-type-id"),
								RequestId:          cmd.String("request-id"),
								ValidateOnly:       cmd.Bool("validate-only"),
							}

							op, err := client.CreateCustomTargetType(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "update",
						Usage: "update custom-target-types",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, updating a `CustomTargetType` that does not exist.", Required: false},
							&cli.StringFlag{Name: "custom-target-type", Usage: "The ID of the custom target type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "custom_target_type.name" not yet supported.
							custom_target_type_name := fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-target-type"))
							fmt.Printf("Executing update on %s\n", custom_target_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete custom-target-types",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, then deleting an already deleted or non-existing.", Required: false},
							&cli.StringFlag{Name: "custom-target-type", Usage: "The ID of the custom target type.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated but no actual change is.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customTargetTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-target-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCustomTargetType %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.DeleteCustomTargetTypeRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteCustomTargetType(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "filter", Usage: "Filter pipelines to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of pipelines to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDeliveryPipelines` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListDeliveryPipelinesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeliveryPipelines(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe delivery-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetDeliveryPipelineRequest{
								Name: name,
							}

							resp, err := client.GetDeliveryPipeline(ctx, req)
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
						Usage: "create delivery-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline-id", Usage: "ID of the `DeliveryPipeline`.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateDeliveryPipelineRequest{
								Parent:             parent,
								DeliveryPipelineId: cmd.String("delivery-pipeline-id"),
								RequestId:          cmd.String("request-id"),
								ValidateOnly:       cmd.Bool("validate-only"),
							}

							op, err := client.CreateDeliveryPipeline(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "update",
						Usage: "update delivery-pipelines",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, updating a `DeliveryPipeline` that does not exist.", Required: false},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "delivery_pipeline.name" not yet supported.
							delivery_pipeline_name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							fmt.Printf("Executing update on %s\n", delivery_pipeline_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete delivery-pipelines",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, then deleting an already deleted or non-existing.", Required: false},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all child resources under this pipeline will also.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDeliveryPipeline %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.DeleteDeliveryPipelineRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
								Force:        cmd.Bool("force"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteDeliveryPipeline(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "rollback-target",
						Usage: "rollback-target delivery-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-id", Usage: "ID of the `Release` to roll back to.", Required: false},
							&cli.StringFlag{Name: "rollout-id", Usage: "ID of the rollback `Rollout` to create.", Required: true},
							&cli.StringFlag{Name: "rollout-to-roll-back", Usage: "If provided, this must be the latest `Rollout` that is on the.", Required: false},
							&cli.StringFlag{Name: "target-id", Usage: "ID of the `Target` that is being rolled back.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.RollbackTargetRequest{
								Name:                 name,
								TargetId:             cmd.String("target-id"),
								RolloutId:            cmd.String("rollout-id"),
								ReleaseId:            cmd.String("release-id"),
								RolloutToRollBack:    cmd.String("rollout-to-roll-back"),
								ValidateOnly:         cmd.Bool("validate-only"),
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
							}

							resp, err := client.RollbackTarget(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy delivery-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe delivery-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions delivery-pipelines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
							&cli.StringFlag{Name: "deploy-policy-id", Usage: "ID of the `DeployPolicy`.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateDeployPolicyRequest{
								Parent:         parent,
								DeployPolicyId: cmd.String("deploy-policy-id"),
								RequestId:      cmd.String("request-id"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.CreateDeployPolicy(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "update",
						Usage: "update deploy-policies",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, updating a `DeployPolicy` that does not exist.", Required: false},
							&cli.StringFlag{Name: "deploy-policie", Usage: "The ID of the deploy policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deploy_policy.name" not yet supported.
							deploy_policy_name := fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("deploy-policie"))
							fmt.Printf("Executing update on %s\n", deploy_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete deploy-policies",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, then deleting an already deleted or non-existing.", Required: false},
							&cli.StringFlag{Name: "deploy-policie", Usage: "The ID of the deploy policie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("deploy-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDeployPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.DeleteDeployPolicyRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteDeployPolicy(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list deploy-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter deploy policies to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of deploy policies to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDeployPolicies` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListDeployPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeployPolicies(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe deploy-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deploy-policie", Usage: "The ID of the deploy policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deployPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("deploy-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetDeployPolicyRequest{
								Name: name,
							}

							resp, err := client.GetDeployPolicy(ctx, req)
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
				Name:  "job-runs",
				Usage: "Manage job-runs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list job-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter results to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `JobRun` objects to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListJobRuns` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListJobRunsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListJobRuns(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe job-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "job-run", Usage: "The ID of the job run.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s/jobRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"), cmd.String("job-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetJobRunRequest{
								Name: name,
							}

							resp, err := client.GetJobRun(ctx, req)
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
						Name:  "terminate",
						Usage: "terminate job-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "job-run", Usage: "The ID of the job run.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s/jobRuns/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"), cmd.String("job-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.TerminateJobRunRequest{
								Name:                 name,
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
							}

							resp, err := client.TerminateJobRun(ctx, req)
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
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
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
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter releases to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `Release` objects to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListReleases` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListReleasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListReleases(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetReleaseRequest{
								Name: name,
							}

							resp, err := client.GetRelease(ctx, req)
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
						Usage: "create releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-id", Usage: "ID of the `Release`.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateReleaseRequest{
								Parent:               parent,
								ReleaseId:            cmd.String("release-id"),
								RequestId:            cmd.String("request-id"),
								ValidateOnly:         cmd.Bool("validate-only"),
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
							}

							op, err := client.CreateRelease(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "abandon",
						Usage: "abandon releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.AbandonReleaseRequest{
								Name: name,
							}

							resp, err := client.AbandonRelease(ctx, req)
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
				Name:  "rollouts",
				Usage: "Manage rollouts resources",
				Commands: []*cli.Command{

					{
						Name:  "approve",
						Usage: "approve rollouts",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "approved", Usage: "True = approve; false = reject.", Required: true},
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ApproveRolloutRequest{
								Name:                 name,
								Approved:             cmd.Bool("approved"),
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
							}

							resp, err := client.ApproveRollout(ctx, req)
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
						Name:  "advance",
						Usage: "advance rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "phase-id", Usage: "The phase ID to advance the `Rollout` to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.AdvanceRolloutRequest{
								Name:                 name,
								PhaseId:              cmd.String("phase-id"),
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
							}

							resp, err := client.AdvanceRollout(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CancelRolloutRequest{
								Name:                 name,
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
							}

							resp, err := client.CancelRollout(ctx, req)
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
						Usage: "list rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter rollouts to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `Rollout` objects to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListRollouts` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListRolloutsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRollouts(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetRolloutRequest{
								Name: name,
							}

							resp, err := client.GetRollout(ctx, req)
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
						Usage: "create rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "rollout-id", Usage: "ID of the `Rollout`.", Required: true},
							&cli.StringFlag{Name: "starting-phase-id", Usage: "The starting phase ID for the `Rollout`.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateRolloutRequest{
								Parent:               parent,
								RolloutId:            cmd.String("rollout-id"),
								RequestId:            cmd.String("request-id"),
								ValidateOnly:         cmd.Bool("validate-only"),
								OverrideDeployPolicy: cmd.StringSlice("override-deploy-policy"),
								StartingPhaseId:      cmd.String("starting-phase-id"),
							}

							op, err := client.CreateRollout(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "ignore-job",
						Usage: "ignore-job rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "job-id", Usage: "The job ID for the Job to ignore.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "phase-id", Usage: "The phase ID the Job to ignore belongs to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							rollout := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							fmt.Printf("Executing ignore-job on %s\n", rollout)
							return nil
						},
					},

					{
						Name:  "retry-job",
						Usage: "retry-job rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "delivery-pipeline", Usage: "The ID of the delivery pipeline.", Required: true},
							&cli.StringFlag{Name: "job-id", Usage: "The job ID for the Job to retry.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "override-deploy-policy", Usage: "Deploy policies to override.", Required: false},
							&cli.StringFlag{Name: "phase-id", Usage: "The phase ID the Job to retry belongs to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							rollout := fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s/releases/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("delivery-pipeline"), cmd.String("release"), cmd.String("rollout"))
							fmt.Printf("Executing retry-job on %s\n", rollout)
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
							&cli.StringFlag{Name: "filter", Usage: "Filter targets to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Field to sort by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of `Target` objects to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListTargets` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.ListTargetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTargets(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe targets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target", Usage: "The ID of the target.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/targets/%s", cmd.String("project"), cmd.String("location"), cmd.String("target"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.GetTargetRequest{
								Name: name,
							}

							resp, err := client.GetTarget(ctx, req)
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
						Usage: "create targets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target-id", Usage: "ID of the `Target`.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.CreateTargetRequest{
								Parent:       parent,
								TargetId:     cmd.String("target-id"),
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateTarget(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "update",
						Usage: "update targets",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, updating a `Target` that does not exist will.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target", Usage: "The ID of the target.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, the request is validated and the user is provided.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "target.name" not yet supported.
							target_name := fmt.Sprintf("projects/%s/locations/%s/targets/%s", cmd.String("project"), cmd.String("location"), cmd.String("target"))
							fmt.Printf("Executing update on %s\n", target_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete targets",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, then deleting an already deleted or non-existing.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target", Usage: "The ID of the target.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/targets/%s", cmd.String("project"), cmd.String("location"), cmd.String("target"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTarget %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := deploy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &deploypb.DeleteTargetRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								AllowMissing: cmd.Bool("allow-missing"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteTarget(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
		},
	}
}
