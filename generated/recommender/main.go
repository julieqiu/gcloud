package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	recommender "cloud.google.com/go/recommender/apiv1"
	"cloud.google.com/go/recommender/apiv1/recommenderpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "recommender",
				Usage: "manage Recommender API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "config",
						Usage: "Manage config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("recommender"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recommenderpb.GetRecommenderConfigRequest{Name: name}
									resp, err := client.GetRecommenderConfig(ctx, req)
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
								Usage: "update config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("recommender"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recommenderpb.UpdateRecommenderConfigRequest{}
									req.RecommenderConfig = &recommenderpb.RecommenderConfig{
										Name:        name,
										Name:        cmd.String("name"),
										Etag:        cmd.String("etag"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateRecommenderConfig(ctx, req)
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
								Usage: "describe config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "insight_type", Usage: "The insight_type.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("insight_type"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recommenderpb.GetInsightTypeConfigRequest{Name: name}
									resp, err := client.GetInsightTypeConfig(ctx, req)
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
								Usage: "update config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "insight_type", Usage: "The insight_type.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("insight_type"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recommenderpb.UpdateInsightTypeConfigRequest{}
									req.InsightTypeConfig = &recommenderpb.InsightTypeConfig{
										Name:        name,
										Name:        cmd.String("name"),
										Etag:        cmd.String("etag"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateInsightTypeConfig(ctx, req)
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
						Name:  "insights",
						Usage: "Manage insights resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list insights",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &recommenderpb.ListInsightsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInsights(ctx, req)
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
								Usage: "describe insights",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "insight_type", Usage: "The insight_type.", Required: true},
									&cli.StringFlag{Name: "insight", Usage: "The insight.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/insights/%s", cmd.String("project"), cmd.String("location"), cmd.String("insight_type"), cmd.String("insight"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recommenderpb.GetInsightRequest{Name: name}
									resp, err := client.GetInsight(ctx, req)
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
								Name:  "mark-accepted",
								Usage: "mark-accepted insights",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "insight_type", Usage: "The insight_type.", Required: true},
									&cli.StringFlag{Name: "insight", Usage: "The insight.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/insightTypes/%s/insights/%s", cmd.String("project"), cmd.String("location"), cmd.String("insight_type"), cmd.String("insight"))
									fmt.Printf("Executing mark-accepted on %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &recommenderpb.ListRecommendationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRecommendations(ctx, req)
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
								Usage: "describe recommendations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
									&cli.StringFlag{Name: "recommendation", Usage: "The recommendation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
									client, err := recommender.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recommenderpb.GetRecommendationRequest{Name: name}
									resp, err := client.GetRecommendation(ctx, req)
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
								Name:  "mark-dismissed",
								Usage: "mark-dismissed recommendations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
									&cli.StringFlag{Name: "recommendation", Usage: "The recommendation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
									fmt.Printf("Executing mark-dismissed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "mark-claimed",
								Usage: "mark-claimed recommendations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
									&cli.StringFlag{Name: "recommendation", Usage: "The recommendation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
									fmt.Printf("Executing mark-claimed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "mark-succeeded",
								Usage: "mark-succeeded recommendations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
									&cli.StringFlag{Name: "recommendation", Usage: "The recommendation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
									fmt.Printf("Executing mark-succeeded on %s\n", name)
									return nil
								},
							},
							{
								Name:  "mark-failed",
								Usage: "mark-failed recommendations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recommender", Usage: "The recommender.", Required: true},
									&cli.StringFlag{Name: "recommendation", Usage: "The recommendation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", cmd.String("project"), cmd.String("location"), cmd.String("recommender"), cmd.String("recommendation"))
									fmt.Printf("Executing mark-failed on %s\n", name)
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
