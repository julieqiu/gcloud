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

package websecurityscanner

import (
	websecurityscanner "cloud.google.com/go/websecurityscanner/apiv1"
	"cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the websecurityscanner command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "websecurityscanner",
		Usage: "manage Web Security Scanner API resources",
		Commands: []*cli.Command{
			{
				Name:  "crawled-urls",
				Usage: "Manage crawled-urls resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list crawled-urls",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of CrawledUrls to return, can be limited by server.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to be returned.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
							&cli.StringFlag{Name: "scan-run", Usage: "The ID of the scan run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/scanConfigs/%s/scanRuns/%s", cmd.String("project"), cmd.String("scan-config"), cmd.String("scan-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.ListCrawledUrlsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListCrawledUrls(ctx, req)
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
				Name:  "finding-type-stats",
				Usage: "Manage finding-type-stats resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list finding-type-stats",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
							&cli.StringFlag{Name: "scan-run", Usage: "The ID of the scan run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/scanConfigs/%s/scanRuns/%s", cmd.String("project"), cmd.String("scan-config"), cmd.String("scan-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.ListFindingTypeStatsRequest{
								Parent: parent,
							}

							resp, err := client.ListFindingTypeStats(ctx, req)
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
				Name:  "findings",
				Usage: "Manage findings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "finding", Usage: "The ID of the finding.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
							&cli.StringFlag{Name: "scan-run", Usage: "The ID of the scan run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/scanConfigs/%s/scanRuns/%s/findings/%s", cmd.String("project"), cmd.String("scan-config"), cmd.String("scan-run"), cmd.String("finding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.GetFindingRequest{
								Name: name,
							}

							resp, err := client.GetFinding(ctx, req)
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
						Usage: "list findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Findings to return, can be limited by server.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to be returned.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
							&cli.StringFlag{Name: "scan-run", Usage: "The ID of the scan run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/scanConfigs/%s/scanRuns/%s", cmd.String("project"), cmd.String("scan-config"), cmd.String("scan-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.ListFindingsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListFindings(ctx, req)
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
				Name:  "scan-configs",
				Usage: "Manage scan-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create scan-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.CreateScanConfigRequest{
								Parent: parent,
							}

							resp, err := client.CreateScanConfig(ctx, req)
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
						Usage: "delete scan-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/scanConfigs/%s", cmd.String("project"), cmd.String("scan-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteScanConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.DeleteScanConfigRequest{
								Name: name,
							}

							if err := client.DeleteScanConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe scan-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/scanConfigs/%s", cmd.String("project"), cmd.String("scan-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.GetScanConfigRequest{
								Name: name,
							}

							resp, err := client.GetScanConfig(ctx, req)
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
						Usage: "list scan-configs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of ScanConfigs to return, can be limited by server.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to be returned.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.ListScanConfigsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListScanConfigs(ctx, req)
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
						Usage: "update scan-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "scan_config.name" not yet supported.
							scan_config_name := fmt.Sprintf("projects/%s/scanConfigs/%s", cmd.String("project"), cmd.String("scan-config"))
							fmt.Printf("Executing update on %s\n", scan_config_name)
							return nil
						},
					},

					{
						Name:  "start",
						Usage: "start scan-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/scanConfigs/%s", cmd.String("project"), cmd.String("scan-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.StartScanRunRequest{
								Name: name,
							}

							resp, err := client.StartScanRun(ctx, req)
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
				Name:  "scan-runs",
				Usage: "Manage scan-runs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe scan-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
							&cli.StringFlag{Name: "scan-run", Usage: "The ID of the scan run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/scanConfigs/%s/scanRuns/%s", cmd.String("project"), cmd.String("scan-config"), cmd.String("scan-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.GetScanRunRequest{
								Name: name,
							}

							resp, err := client.GetScanRun(ctx, req)
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
						Usage: "list scan-runs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of ScanRuns to return, can be limited by server.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results to be returned.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/scanConfigs/%s", cmd.String("project"), cmd.String("scan-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.ListScanRunsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListScanRuns(ctx, req)
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
						Name:  "stop",
						Usage: "stop scan-runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scan-config", Usage: "The ID of the scan config.", Required: true},
							&cli.StringFlag{Name: "scan-run", Usage: "The ID of the scan run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/scanConfigs/%s/scanRuns/%s", cmd.String("project"), cmd.String("scan-config"), cmd.String("scan-run"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := websecurityscanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &websecurityscannerpb.StopScanRunRequest{
								Name: name,
							}

							resp, err := client.StopScanRun(ctx, req)
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
