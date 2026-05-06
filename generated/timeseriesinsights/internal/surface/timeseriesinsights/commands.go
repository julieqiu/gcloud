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

package timeseriesinsights

import (
	timeseriesinsights "cloud.google.com/go/timeseriesinsights/apiv1"
	"cloud.google.com/go/timeseriesinsights/apiv1/timeseriesinsightspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the timeseriesinsights command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "timeseriesinsights",
		Usage: "manage Timeseries Insights API resources",
		Commands: []*cli.Command{
			{
				Name:  "datasets",
				Usage: "Manage datasets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results to return in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token to provide to skip to a particular spot in the list.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := timeseriesinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &timeseriesinsightspb.ListDataSetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataSets(ctx, req)
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
						Usage: "create datasets",
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
							client, err := timeseriesinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &timeseriesinsightspb.CreateDataSetRequest{
								Parent: parent,
							}

							resp, err := client.CreateDataSet(ctx, req)
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
						Usage: "delete datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDataSet on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := timeseriesinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &timeseriesinsightspb.DeleteDataSetRequest{
								Name: name,
							}

							if err := client.DeleteDataSet(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "append-events",
						Usage: "append-events datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							dataset := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing append-events on %s\n", dataset)
							return nil
						},
					},

					{
						Name:  "query",
						Usage: "query datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "num-returned-slices", Usage: "How many slices are returned in.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-timeseries", Usage: "If specified, we will return the actual and forecasted time for all.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := timeseriesinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &timeseriesinsightspb.QueryDataSetRequest{
								Name:              name,
								NumReturnedSlices: runtime.Ptr(int32(cmd.Int("num-returned-slices"))),
								ReturnTimeseries:  cmd.Bool("return-timeseries"),
							}

							resp, err := client.QueryDataSet(ctx, req)
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
						Name:  "evaluate-slice",
						Usage: "evaluate-slice datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The ID of the dataset.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							dataset := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing evaluate-slice on %s\n", dataset)
							return nil
						},
					},

					{
						Name:  "evaluate-timeseries",
						Usage: "evaluate-timeseries datasets",
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
							client, err := timeseriesinsights.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &timeseriesinsightspb.EvaluateTimeseriesRequest{
								Parent: parent,
							}

							resp, err := client.EvaluateTimeseries(ctx, req)
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
