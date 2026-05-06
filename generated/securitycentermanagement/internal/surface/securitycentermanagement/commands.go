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

package securitycentermanagement

import (
	securitycentermanagement "cloud.google.com/go/securitycentermanagement/apiv1"
	"cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the securitycentermanagement command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "securitycentermanagement",
		Usage: "manage Security Command Center Management API resources",
		Commands: []*cli.Command{
			{
				Name:  "effective-event-threat-detection-custom-modules",
				Usage: "Manage effective-event-threat-detection-custom-modules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list effective-event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListEffectiveEventThreatDetectionCustomModulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEffectiveEventThreatDetectionCustomModules(ctx, req)
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
						Usage: "describe effective-event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "effective-event-threat-detection-custom-module", Usage: "The ID of the effective event threat detection custom module.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/effectiveEventThreatDetectionCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("effective-event-threat-detection-custom-module"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.GetEffectiveEventThreatDetectionCustomModuleRequest{
								Name: name,
							}

							resp, err := client.GetEffectiveEventThreatDetectionCustomModule(ctx, req)
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
				Name:  "effective-security-health-analytics-custom-modules",
				Usage: "Manage effective-security-health-analytics-custom-modules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list effective-security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListEffectiveSecurityHealthAnalyticsCustomModulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEffectiveSecurityHealthAnalyticsCustomModules(ctx, req)
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
						Usage: "describe effective-security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "effective-security-health-analytics-custom-module", Usage: "The ID of the effective security health analytics custom module.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/effectiveSecurityHealthAnalyticsCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("effective-security-health-analytics-custom-module"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.GetEffectiveSecurityHealthAnalyticsCustomModuleRequest{
								Name: name,
							}

							resp, err := client.GetEffectiveSecurityHealthAnalyticsCustomModule(ctx, req)
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
				Name:  "event-threat-detection-custom-modules",
				Usage: "Manage event-threat-detection-custom-modules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of modules to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListEventThreatDetectionCustomModulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEventThreatDetectionCustomModules(ctx, req)
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
						Name:  "list",
						Usage: "list event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of modules to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListDescendantEventThreatDetectionCustomModulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDescendantEventThreatDetectionCustomModules(ctx, req)
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
						Usage: "describe event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "event-threat-detection-custom-module", Usage: "The ID of the event threat detection custom module.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/eventThreatDetectionCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("event-threat-detection-custom-module"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.GetEventThreatDetectionCustomModuleRequest{
								Name: name,
							}

							resp, err := client.GetEventThreatDetectionCustomModule(ctx, req)
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
						Usage: "create event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.CreateEventThreatDetectionCustomModuleRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							resp, err := client.CreateEventThreatDetectionCustomModule(ctx, req)
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
						Usage: "update event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "event-threat-detection-custom-module", Usage: "The ID of the event threat detection custom module.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "event_threat_detection_custom_module.name" not yet supported.
							event_threat_detection_custom_module_name := fmt.Sprintf("projects/%s/locations/%s/eventThreatDetectionCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("event-threat-detection-custom-module"))
							fmt.Printf("Executing update on %s\n", event_threat_detection_custom_module_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "event-threat-detection-custom-module", Usage: "The ID of the event threat detection custom module.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/eventThreatDetectionCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("event-threat-detection-custom-module"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEventThreatDetectionCustomModule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.DeleteEventThreatDetectionCustomModuleRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							if err := client.DeleteEventThreatDetectionCustomModule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "validate",
						Usage: "validate event-threat-detection-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "raw-text", Usage: "The raw text of the module's contents.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The type of the module.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ValidateEventThreatDetectionCustomModuleRequest{
								Parent:  parent,
								RawText: cmd.String("raw-text"),
								Type:    cmd.String("type"),
							}

							resp, err := client.ValidateEventThreatDetectionCustomModule(ctx, req)
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
				Name:  "security-center-services",
				Usage: "Manage security-center-services resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe security-center-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-center-service", Usage: "The ID of the security center service.", Required: true},
							&cli.BoolFlag{Name: "show-eligible-modules-only", Usage: "Set to `true` to show only modules that are in scope.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/securityCenterServices/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-center-service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.GetSecurityCenterServiceRequest{
								Name:                    name,
								ShowEligibleModulesOnly: cmd.Bool("show-eligible-modules-only"),
							}

							resp, err := client.GetSecurityCenterService(ctx, req)
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
						Usage: "list security-center-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-eligible-modules-only", Usage: "Flag that, when set, is used to filter the module settings that are shown.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListSecurityCenterServicesRequest{
								Parent:                  parent,
								PageSize:                int32(cmd.Int("page-size")),
								PageToken:               cmd.String("page-token"),
								ShowEligibleModulesOnly: cmd.Bool("show-eligible-modules-only"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecurityCenterServices(ctx, req)
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
						Name:  "update",
						Usage: "update security-center-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-center-service", Usage: "The ID of the security center service.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "security_center_service.name" not yet supported.
							security_center_service_name := fmt.Sprintf("projects/%s/locations/%s/securityCenterServices/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-center-service"))
							fmt.Printf("Executing update on %s\n", security_center_service_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "security-health-analytics-custom-modules",
				Usage: "Manage security-health-analytics-custom-modules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListSecurityHealthAnalyticsCustomModulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecurityHealthAnalyticsCustomModules(ctx, req)
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
						Name:  "list",
						Usage: "list security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A pagination token returned from a previous request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.ListDescendantSecurityHealthAnalyticsCustomModulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDescendantSecurityHealthAnalyticsCustomModules(ctx, req)
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
						Usage: "describe security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-health-analytics-custom-module", Usage: "The ID of the security health analytics custom module.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/securityHealthAnalyticsCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-health-analytics-custom-module"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.GetSecurityHealthAnalyticsCustomModuleRequest{
								Name: name,
							}

							resp, err := client.GetSecurityHealthAnalyticsCustomModule(ctx, req)
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
						Usage: "create security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.CreateSecurityHealthAnalyticsCustomModuleRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							resp, err := client.CreateSecurityHealthAnalyticsCustomModule(ctx, req)
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
						Usage: "update security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-health-analytics-custom-module", Usage: "The ID of the security health analytics custom module.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "security_health_analytics_custom_module.name" not yet supported.
							security_health_analytics_custom_module_name := fmt.Sprintf("projects/%s/locations/%s/securityHealthAnalyticsCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-health-analytics-custom-module"))
							fmt.Printf("Executing update on %s\n", security_health_analytics_custom_module_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security-health-analytics-custom-module", Usage: "The ID of the security health analytics custom module.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When set to `true`, the request will be validated (including IAM.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/securityHealthAnalyticsCustomModules/%s", cmd.String("project"), cmd.String("location"), cmd.String("security-health-analytics-custom-module"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSecurityHealthAnalyticsCustomModule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.DeleteSecurityHealthAnalyticsCustomModuleRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							if err := client.DeleteSecurityHealthAnalyticsCustomModule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "simulate",
						Usage: "simulate security-health-analytics-custom-modules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycentermanagement.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycentermanagementpb.SimulateSecurityHealthAnalyticsCustomModuleRequest{
								Parent: parent,
							}

							resp, err := client.SimulateSecurityHealthAnalyticsCustomModule(ctx, req)
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
