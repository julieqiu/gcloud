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

package recommender

import (
	recommender "cloud.google.com/go/recommender/apiv1"
	"cloud.google.com/go/recommender/apiv1/recommenderpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the recommender command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "recommender",
		Usage: "manage Recommender API resources",
		Commands: []*cli.Command{
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
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("recommender"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.GetRecommenderConfigRequest{
								Name: name,
							}

							resp, err := client.GetRecommenderConfig(ctx, req)
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
						Usage: "update config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, validate the request and preview the change, but do not actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "recommender_config.name" not yet supported.
							recommender_config_name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("recommender"))
							fmt.Printf("Executing update on %s\n", recommender_config_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "insight-type", Usage: "The ID of the insight type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("insight-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.GetInsightTypeConfigRequest{
								Name: name,
							}

							resp, err := client.GetInsightTypeConfig(ctx, req)
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
						Usage: "update config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "insight-type", Usage: "The ID of the insight type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, validate the request and preview the change, but do not actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "insight_type_config.name" not yet supported.
							insight_type_config_name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("insight-type"))
							fmt.Printf("Executing update on %s\n", insight_type_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "insights",
				Usage: "Manage insights resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list insights",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression to restrict the insights returned.", Required: false},
							&cli.StringFlag{Name: "insight-type", Usage: "The ID of the insight type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, retrieves the next batch of results from the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("insight-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.ListInsightsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListInsights(ctx, req)
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
						Usage: "describe insights",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "insight", Usage: "The ID of the insight.", Required: true},
							&cli.StringFlag{Name: "insight-type", Usage: "The ID of the insight type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/insights/%s", cmd.String("project"), cmd.String("location"), cmd.String("insight-type"), cmd.String("insight"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.GetInsightRequest{
								Name: name,
							}

							resp, err := client.GetInsight(ctx, req)
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
						Name:  "mark-accepted",
						Usage: "mark-accepted insights",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Fingerprint of the Insight.", Required: true},
							&cli.StringFlag{Name: "insight", Usage: "The ID of the insight.", Required: true},
							&cli.StringFlag{Name: "insight-type", Usage: "The ID of the insight type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/insights/%s", cmd.String("project"), cmd.String("location"), cmd.String("insight-type"), cmd.String("insight"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.MarkInsightAcceptedRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.MarkInsightAccepted(ctx, req)
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
				Name:  "recommendations",
				Usage: "Manage recommendations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list recommendations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression to restrict the recommendations returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, retrieves the next batch of results from the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.ListRecommendationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListRecommendations(ctx, req)
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
						Usage: "describe recommendations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommendation", Usage: "The ID of the recommendation.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.GetRecommendationRequest{
								Name: name,
							}

							resp, err := client.GetRecommendation(ctx, req)
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
						Name:  "mark-dismissed",
						Usage: "mark-dismissed recommendations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Fingerprint of the Recommendation.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommendation", Usage: "The ID of the recommendation.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.MarkRecommendationDismissedRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.MarkRecommendationDismissed(ctx, req)
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
						Name:  "mark-claimed",
						Usage: "mark-claimed recommendations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Fingerprint of the Recommendation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommendation", Usage: "The ID of the recommendation.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.MarkRecommendationClaimedRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.MarkRecommendationClaimed(ctx, req)
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
						Name:  "mark-succeeded",
						Usage: "mark-succeeded recommendations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Fingerprint of the Recommendation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommendation", Usage: "The ID of the recommendation.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.MarkRecommendationSucceededRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.MarkRecommendationSucceeded(ctx, req)
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
						Name:  "mark-failed",
						Usage: "mark-failed recommendations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Fingerprint of the Recommendation.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recommendation", Usage: "The ID of the recommendation.", Required: true},
							&cli.StringFlag{Name: "recommender", Usage: "The ID of the recommender.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recommenderpb.MarkRecommendationFailedRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							resp, err := client.MarkRecommendationFailed(ctx, req)
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
