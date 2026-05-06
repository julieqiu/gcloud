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

package monitoring

import (
	monitoring "cloud.google.com/go/monitoring/apiv3"
	"cloud.google.com/go/monitoring/apiv3/monitoringpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the monitoring command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "monitoring",
		Usage: "manage Cloud Monitoring API resources",
		Commands: []*cli.Command{
			{
				Name:  "alert-policies",
				Usage: "Manage alert-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list alert-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If provided, this field specifies the criteria that must be met.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields by which to sort the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListAlertPoliciesRequest{
								Name:      name,
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAlertPolicies(ctx, req)
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
						Usage: "describe alert-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "alert-policie", Usage: "The ID of the alert policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/alertPolicies/%s", cmd.String("project"), cmd.String("alert-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetAlertPolicyRequest{
								Name: name,
							}

							resp, err := client.GetAlertPolicy(ctx, req)
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
						Usage: "create alert-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateAlertPolicyRequest{
								Name: name,
							}

							resp, err := client.CreateAlertPolicy(ctx, req)
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
						Name:  "delete",
						Usage: "delete alert-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "alert-policie", Usage: "The ID of the alert policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/alertPolicies/%s", cmd.String("project"), cmd.String("alert-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAlertPolicy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteAlertPolicyRequest{
								Name: name,
							}

							if err := client.DeleteAlertPolicy(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update alert-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "alert-policie", Usage: "The ID of the alert policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "alert_policy.name" not yet supported.
							alert_policy_name := fmt.Sprintf("projects/%s/alertPolicies/%s", cmd.String("project"), cmd.String("alert-policie"))
							fmt.Printf("Executing update on %s\n", alert_policy_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "groups",
				Usage: "Manage groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list groups",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "A positive number that is the maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `next_page_token` value.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListGroupsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGroups(ctx, req)
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
						Usage: "describe groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetGroupRequest{
								Name: name,
							}

							resp, err := client.GetGroup(ctx, req)
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
						Usage: "create groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, validate this request but do not create the group.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateGroupRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							resp, err := client.CreateGroup(ctx, req)
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
						Usage: "update groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, validate this request but do not update the existing group.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "group.name" not yet supported.
							group_name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
							fmt.Printf("Executing update on %s\n", group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "recursive", Usage: "If this field is true, then the request means to delete a group with all.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGroup on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteGroupRequest{
								Name:      name,
								Recursive: cmd.Bool("recursive"),
							}

							if err := client.DeleteGroup(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "members",
				Usage: "Manage members resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list members",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An optional [list.", Required: false},
							&cli.StringFlag{Name: "group", Usage: "The ID of the group.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "A positive number that is the maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `next_page_token` value.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListGroupMembersRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListGroupMembers(ctx, req)
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
				Name:  "metric-descriptors",
				Usage: "Manage metric-descriptors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list metric-descriptors",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "active-only", Usage: "If true, only metrics and monitored resource types that have.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "If this field is empty, all custom and.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "A positive number that is the maximum number of results to.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListMetricDescriptorsRequest{
								Name:       name,
								Filter:     cmd.String("filter"),
								PageSize:   int32(cmd.Int("page-size")),
								PageToken:  cmd.String("page-token"),
								ActiveOnly: cmd.Bool("active-only"),
							}

							limit := cmd.Int("limit")
							it := client.ListMetricDescriptors(ctx, req)
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
						Usage: "describe metric-descriptors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metric-descriptor", Usage: "The ID of the metric descriptor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/metricDescriptors/%s", cmd.String("project"), cmd.String("metric-descriptor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetMetricDescriptorRequest{
								Name: name,
							}

							resp, err := client.GetMetricDescriptor(ctx, req)
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
						Usage: "create metric-descriptors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateMetricDescriptorRequest{
								Name: name,
							}

							resp, err := client.CreateMetricDescriptor(ctx, req)
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
						Name:  "delete",
						Usage: "delete metric-descriptors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metric-descriptor", Usage: "The ID of the metric descriptor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/metricDescriptors/%s", cmd.String("project"), cmd.String("metric-descriptor"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteMetricDescriptor on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteMetricDescriptorRequest{
								Name: name,
							}

							if err := client.DeleteMetricDescriptor(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "monitored-resource-descriptors",
				Usage: "Manage monitored-resource-descriptors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list monitored-resource-descriptors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An optional [filter](https://cloud.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "A positive number that is the maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListMonitoredResourceDescriptorsRequest{
								Name:      name,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMonitoredResourceDescriptors(ctx, req)
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
						Usage: "describe monitored-resource-descriptors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "monitored-resource-descriptor", Usage: "The ID of the monitored resource descriptor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/monitoredResourceDescriptors/%s", cmd.String("project"), cmd.String("monitored-resource-descriptor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetMonitoredResourceDescriptorRequest{
								Name: name,
							}

							resp, err := client.GetMonitoredResourceDescriptor(ctx, req)
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
				Name:  "notification-channel-descriptors",
				Usage: "Manage notification-channel-descriptors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list notification-channel-descriptors",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If non-empty, `page_token` must contain a value returned as the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListNotificationChannelDescriptorsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListNotificationChannelDescriptors(ctx, req)
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
						Usage: "describe notification-channel-descriptors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification-channel-descriptor", Usage: "The ID of the notification channel descriptor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/notificationChannelDescriptors/%s", cmd.String("project"), cmd.String("notification-channel-descriptor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetNotificationChannelDescriptorRequest{
								Name: name,
							}

							resp, err := client.GetNotificationChannelDescriptor(ctx, req)
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
				Name:  "notification-channels",
				Usage: "Manage notification-channels resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If provided, this field specifies the criteria that must be met.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields by which to sort the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If non-empty, `page_token` must contain a value returned as the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListNotificationChannelsRequest{
								Name:      name,
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListNotificationChannels(ctx, req)
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
						Usage: "describe notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification-channel", Usage: "The ID of the notification channel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetNotificationChannelRequest{
								Name: name,
							}

							resp, err := client.GetNotificationChannel(ctx, req)
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
						Usage: "create notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateNotificationChannelRequest{
								Name: name,
							}

							resp, err := client.CreateNotificationChannel(ctx, req)
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
						Usage: "update notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification-channel", Usage: "The ID of the notification channel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "notification_channel.name" not yet supported.
							notification_channel_name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification-channel"))
							fmt.Printf("Executing update on %s\n", notification_channel_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete notification-channels",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If true, the notification channel will be deleted regardless of its.", Required: false},
							&cli.StringFlag{Name: "notification-channel", Usage: "The ID of the notification channel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification-channel"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteNotificationChannel on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteNotificationChannelRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteNotificationChannel(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "send-verification-code",
						Usage: "send-verification-code notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification-channel", Usage: "The ID of the notification channel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification-channel"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute SendNotificationChannelVerificationCode on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.SendNotificationChannelVerificationCodeRequest{
								Name: name,
							}

							if err := client.SendNotificationChannelVerificationCode(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "get-verification-code",
						Usage: "get-verification-code notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification-channel", Usage: "The ID of the notification channel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetNotificationChannelVerificationCodeRequest{
								Name: name,
							}

							resp, err := client.GetNotificationChannelVerificationCode(ctx, req)
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
						Name:  "verify",
						Usage: "verify notification-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "code", Usage: "The verification code that was delivered to the channel as.", Required: true},
							&cli.StringFlag{Name: "notification-channel", Usage: "The ID of the notification channel.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification-channel"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.VerifyNotificationChannelRequest{
								Name: name,
								Code: cmd.String("code"),
							}

							resp, err := client.VerifyNotificationChannel(ctx, req)
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
				Name:  "service-level-objectives",
				Usage: "Manage service-level-objectives resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create service-level-objectives",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "service-level-objective-id", Usage: "The ServiceLevelObjective id to use for this.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateServiceLevelObjectiveRequest{
								Parent:                  parent,
								ServiceLevelObjectiveId: cmd.String("service-level-objective-id"),
							}

							resp, err := client.CreateServiceLevelObjective(ctx, req)
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
						Name:  "describe",
						Usage: "describe service-level-objectives",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "service-level-objective", Usage: "The ID of the service level objective.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "View of the `ServiceLevelObjective` to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s/serviceLevelObjectives/%s", cmd.String("service"), cmd.String("service-level-objective"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetServiceLevelObjectiveRequest{
								Name: name,
								View: monitoringpb.ServiceLevelObjective_View(monitoringpb.ServiceLevelObjective_View_value[cmd.String("view")]),
							}

							resp, err := client.GetServiceLevelObjective(ctx, req)
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
						Usage: "list service-level-objectives",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter specifying what `ServiceLevelObjective`s to return.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "A non-negative number that is the maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "View of the `ServiceLevelObjective`s to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListServiceLevelObjectivesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								View:      monitoringpb.ServiceLevelObjective_View(monitoringpb.ServiceLevelObjective_View_value[cmd.String("view")]),
							}

							limit := cmd.Int("limit")
							it := client.ListServiceLevelObjectives(ctx, req)
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
						Usage: "update service-level-objectives",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "service-level-objective", Usage: "The ID of the service level objective.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_level_objective.name" not yet supported.
							service_level_objective_name := fmt.Sprintf("services/%s/serviceLevelObjectives/%s", cmd.String("service"), cmd.String("service-level-objective"))
							fmt.Printf("Executing update on %s\n", service_level_objective_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-level-objectives",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "service-level-objective", Usage: "The ID of the service level objective.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s/serviceLevelObjectives/%s", cmd.String("service"), cmd.String("service-level-objective"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteServiceLevelObjective on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteServiceLevelObjectiveRequest{
								Name: name,
							}

							if err := client.DeleteServiceLevelObjective(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "create",
						Usage: "create services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service-id", Usage: "The Service id to use for this Service.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateServiceRequest{
								ServiceId: cmd.String("service-id"),
							}

							resp, err := client.CreateService(ctx, req)
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
						Name:  "describe",
						Usage: "describe services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetServiceRequest{
								Name: name,
							}

							resp, err := client.GetService(ctx, req)
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
						Usage: "list services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter specifying what `Service`s to return.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "A non-negative number that is the maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListServicesRequest{
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListServices(ctx, req)
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
						Usage: "update services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service.name" not yet supported.
							service_name := fmt.Sprintf("services/%s", cmd.String("service"))
							fmt.Printf("Executing update on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("services/%s", cmd.String("service"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteService on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteServiceRequest{
								Name: name,
							}

							if err := client.DeleteService(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "snoozes",
				Usage: "Manage snoozes resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create snoozes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateSnoozeRequest{
								Parent: parent,
							}

							resp, err := client.CreateSnooze(ctx, req)
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
						Usage: "list snoozes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Optional filter to restrict results to the given criteria.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return for a single query.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `next_page_token` from a previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListSnoozesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSnoozes(ctx, req)
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
						Usage: "describe snoozes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snooze", Usage: "The ID of the snooze.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/snoozes/%s", cmd.String("project"), cmd.String("snooze"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetSnoozeRequest{
								Name: name,
							}

							resp, err := client.GetSnooze(ctx, req)
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
						Usage: "update snoozes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snooze", Usage: "The ID of the snooze.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "snooze.name" not yet supported.
							snooze_name := fmt.Sprintf("projects/%s/snoozes/%s", cmd.String("project"), cmd.String("snooze"))
							fmt.Printf("Executing update on %s\n", snooze_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "time-series",
				Usage: "Manage time-series resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A [monitoring.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Unsupported: must be left blank.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "A positive number that is the maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specifies which information is returned about the time series.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListTimeSeriesRequest{
								Name:      name,
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								View:      monitoringpb.ListTimeSeriesRequest_TimeSeriesView(monitoringpb.ListTimeSeriesRequest_TimeSeriesView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTimeSeries(ctx, req)
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
						Name:  "create",
						Usage: "create time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CreateTimeSeries on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateTimeSeriesRequest{
								Name: name,
							}

							if err := client.CreateTimeSeries(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create time-series",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CreateServiceTimeSeries on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateTimeSeriesRequest{
								Name: name,
							}

							if err := client.CreateServiceTimeSeries(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "query",
						Usage: "query time-series",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "A positive number that is the maximum number of time_series_data to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "The query in the [Monitoring Query.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.QueryTimeSeriesRequest{
								Name:      name,
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.QueryTimeSeries(ctx, req)
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
				Name:  "uptime-check-configs",
				Usage: "Manage uptime-check-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list uptime-check-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "If provided, this field specifies the criteria that must be met by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListUptimeCheckConfigsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUptimeCheckConfigs(ctx, req)
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
						Usage: "describe uptime-check-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "uptime-check-config", Usage: "The ID of the uptime check config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", cmd.String("project"), cmd.String("uptime-check-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.GetUptimeCheckConfigRequest{
								Name: name,
							}

							resp, err := client.GetUptimeCheckConfig(ctx, req)
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
						Usage: "create uptime-check-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.CreateUptimeCheckConfigRequest{
								Parent: parent,
							}

							resp, err := client.CreateUptimeCheckConfig(ctx, req)
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
						Usage: "update uptime-check-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "uptime-check-config", Usage: "The ID of the uptime check config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "uptime_check_config.name" not yet supported.
							uptime_check_config_name := fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", cmd.String("project"), cmd.String("uptime-check-config"))
							fmt.Printf("Executing update on %s\n", uptime_check_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete uptime-check-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "uptime-check-config", Usage: "The ID of the uptime check config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", cmd.String("project"), cmd.String("uptime-check-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteUptimeCheckConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.DeleteUptimeCheckConfigRequest{
								Name: name,
							}

							if err := client.DeleteUptimeCheckConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "uptime-check-ips",
				Usage: "Manage uptime-check-ips resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list uptime-check-ips",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If this field is not empty then it must contain the `nextPageToken` value.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := monitoring.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &monitoringpb.ListUptimeCheckIpsRequest{
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUptimeCheckIps(ctx, req)
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
		},
	}
}
