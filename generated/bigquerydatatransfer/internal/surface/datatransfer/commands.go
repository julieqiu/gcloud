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

package datatransfer

import (
	datatransfer "cloud.google.com/go/datatransfer/apiv1"
	"cloud.google.com/go/datatransfer/apiv1/datatransferpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the bigquerydatatransfer command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigquerydatatransfer",
		Usage: "manage BigQuery Data Transfer API resources",
		Commands: []*cli.Command{
			{
				Name:  "data-sources",
				Usage: "Manage data-sources resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.GetDataSourceRequest{
								Name: name,
							}

							resp, err := client.GetDataSource(ctx, req)
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
						Usage: "list data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, which can be used to request a specific page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.ListDataSourcesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListDataSources(ctx, req)
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
						Name:  "check-valid-creds",
						Usage: "check-valid-creds data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-source", Usage: "The ID of the data source.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataSources/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.CheckValidCredsRequest{
								Name: name,
							}

							resp, err := client.CheckValidCreds(ctx, req)
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
						Name:  "enroll-data-sources",
						Usage: "enroll-data-sources locations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "data-source-ids", Usage: "Data sources that are enrolled.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute EnrollDataSources on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.EnrollDataSourcesRequest{
								Name:          name,
								DataSourceIds: cmd.StringSlice("data-source-ids"),
							}

							if err := client.EnrollDataSources(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "unenroll-data-sources",
						Usage: "unenroll-data-sources locations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "data-source-ids", Usage: "Data sources that are unenrolled.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute UnenrollDataSources on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.UnenrollDataSourcesRequest{
								Name:          name,
								DataSourceIds: cmd.StringSlice("data-source-ids"),
							}

							if err := client.UnenrollDataSources(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

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
				Name:  "runs",
				Usage: "Manage runs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"), cmd.String("run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.GetTransferRunRequest{
								Name: name,
							}

							resp, err := client.GetTransferRun(ctx, req)
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
						Usage: "delete runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"), cmd.String("run"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTransferRun on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.DeleteTransferRunRequest{
								Name: name,
							}

							if err := client.DeleteTransferRun(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, which can be used to request a specific page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run-attempt", Usage: "Indicates how run attempts are to be pulled.", Required: false},
							&cli.StringSliceFlag{Name: "states", Usage: "When specified, only transfer runs with requested states are returned.", Required: false},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.ListTransferRunsRequest{
								Parent:     parent,
								States:     cmd.StringSlice("states"),
								PageToken:  cmd.String("page-token"),
								PageSize:   int32(cmd.Int("page-size")),
								RunAttempt: datatransferpb.ListTransferRunsRequest_RunAttempt(datatransferpb.ListTransferRunsRequest_RunAttempt_value[cmd.String("run-attempt")]),
							}

							limit := cmd.Int("limit")
							it := client.ListTransferRuns(ctx, req)
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
				Name:  "transfer-configs",
				Usage: "Manage transfer-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-code", Usage: "Deprecated: Authorization code was required when.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-account-name", Usage: "Optional service account email.", Required: false},
							&cli.StringFlag{Name: "version-info", Usage: "Optional version info.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.CreateTransferConfigRequest{
								Parent:             parent,
								AuthorizationCode:  cmd.String("authorization-code"),
								VersionInfo:        cmd.String("version-info"),
								ServiceAccountName: cmd.String("service-account-name"),
							}

							resp, err := client.CreateTransferConfig(ctx, req)
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
						Usage: "update transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-code", Usage: "Deprecated: Authorization code was required when.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-account-name", Usage: "Optional service account email.", Required: false},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
							&cli.StringFlag{Name: "version-info", Usage: "Optional version info.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "transfer_config.name" not yet supported.
							transfer_config_name := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"))
							fmt.Printf("Executing update on %s\n", transfer_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTransferConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.DeleteTransferConfigRequest{
								Name: name,
							}

							if err := client.DeleteTransferConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.GetTransferConfigRequest{
								Name: name,
							}

							resp, err := client.GetTransferConfig(ctx, req)
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
						Usage: "list transfer-configs",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "data-source-ids", Usage: "When specified, only configurations of requested data sources are returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, which can be used to request a specific page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.ListTransferConfigsRequest{
								Parent:        parent,
								DataSourceIds: cmd.StringSlice("data-source-ids"),
								PageToken:     cmd.String("page-token"),
								PageSize:      int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListTransferConfigs(ctx, req)
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
						Name:  "schedule-runs",
						Usage: "schedule-runs transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.ScheduleTransferRunsRequest{
								Parent: parent,
							}

							resp, err := client.ScheduleTransferRuns(ctx, req)
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
						Name:  "start-manual-runs",
						Usage: "start-manual-runs transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.StartManualTransferRunsRequest{
								Parent: parent,
							}

							resp, err := client.StartManualTransferRuns(ctx, req)
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
				Name:  "transfer-logs",
				Usage: "Manage transfer-logs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list transfer-logs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "message-types", Usage: "Message types to return.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, which can be used to request a specific page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The ID of the run.", Required: true},
							&cli.StringFlag{Name: "transfer-config", Usage: "The ID of the transfer config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", cmd.String("project"), cmd.String("location"), cmd.String("transfer-config"), cmd.String("run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datatransfer.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datatransferpb.ListTransferLogsRequest{
								Parent:       parent,
								PageToken:    cmd.String("page-token"),
								PageSize:     int32(cmd.Int("page-size")),
								MessageTypes: cmd.StringSlice("message-types"),
							}

							limit := cmd.Int("limit")
							it := client.ListTransferLogs(ctx, req)
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
