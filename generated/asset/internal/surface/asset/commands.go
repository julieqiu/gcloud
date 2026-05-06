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

package asset

import (
	asset "cloud.google.com/go/asset/apiv1"
	"cloud.google.com/go/asset/apiv1/assetpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudasset command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudasset",
		Usage: "manage Cloud Asset API resources",
		Commands: []*cli.Command{
			{
				Name:  "analysis-query",
				Usage: "Manage analysis-query resources",
				Commands: []*cli.Command{

					{
						Name:  "analyze-iam-policy",
						Usage: "analyze-iam-policy analysis-query",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "saved-analysis-query", Usage: "The name of a saved query, which must be in the format of:.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "analysis_query.scope" not yet supported.
							fmt.Println("Executing analyze-iam-policy...")
							return nil
						},
					},

					{
						Name:  "analyze-iam-policy-longrunning",
						Usage: "analyze-iam-policy-longrunning analysis-query",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "saved-analysis-query", Usage: "The name of a saved query, which must be in the format of:.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "analysis_query.scope" not yet supported.
							fmt.Println("Executing analyze-iam-policy-longrunning...")
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
						Name:  "list",
						Usage: "list assets",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "asset-types", Usage: "A list of asset types to take a snapshot for.", Required: false},
							&cli.StringFlag{Name: "content-type", Usage: "Asset content type.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of assets to be returned in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `next_page_token` returned from the previous `ListAssetsResponse`, or.", Required: false},
							&cli.StringSliceFlag{Name: "relationship-types", Usage: "A list of relationship types to output, for example:.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.ListAssetsRequest{
								AssetTypes:        cmd.StringSlice("asset-types"),
								ContentType:       assetpb.ContentType(assetpb.ContentType_value[cmd.String("content-type")]),
								PageSize:          int32(cmd.Int("page-size")),
								PageToken:         cmd.String("page-token"),
								RelationshipTypes: cmd.StringSlice("relationship-types"),
							}

							limit := cmd.Int("limit")
							it := client.ListAssets(ctx, req)
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
				Name:  "effective-iam-policies",
				Usage: "Manage effective-iam-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-get",
						Usage: "batch-get effective-iam-policies",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "names", Usage: "The names refer to the [full_resource_names].", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "Only IAM policies on or below the scope will be returned.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-get...")
							return nil
						},
					},
				},
			},
			{
				Name:  "feeds",
				Usage: "Manage feeds resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feed-id", Usage: "This is the client-assigned asset feed identifier and it needs to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.CreateFeedRequest{
								FeedId: cmd.String("feed-id"),
							}

							resp, err := client.CreateFeed(ctx, req)
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
						Usage: "describe feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feed", Usage: "The ID of the feed.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("feeds/%s", cmd.String("feed"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.GetFeedRequest{
								Name: name,
							}

							resp, err := client.GetFeed(ctx, req)
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
						Usage: "list feeds",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.ListFeedsRequest{}

							resp, err := client.ListFeeds(ctx, req)
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
						Usage: "update feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feed", Usage: "The ID of the feed.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "feed.name" not yet supported.
							feed_name := fmt.Sprintf("feeds/%s", cmd.String("feed"))
							fmt.Printf("Executing update on %s\n", feed_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feed", Usage: "The ID of the feed.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("feeds/%s", cmd.String("feed"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteFeed on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.DeleteFeedRequest{
								Name: name,
							}

							if err := client.DeleteFeed(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "parent",
				Usage: "Manage parent resources",
				Commands: []*cli.Command{

					{
						Name:  "export-assets",
						Usage: "export-assets parent",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "asset-types", Usage: "A list of asset types to take a snapshot for.", Required: false},
							&cli.StringFlag{Name: "content-type", Usage: "Asset content type.", Required: false},
							&cli.StringSliceFlag{Name: "relationship-types", Usage: "A list of relationship types to export, for example:.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.ExportAssetsRequest{
								AssetTypes:        cmd.StringSlice("asset-types"),
								ContentType:       assetpb.ContentType(assetpb.ContentType_value[cmd.String("content-type")]),
								RelationshipTypes: cmd.StringSlice("relationship-types"),
							}

							op, err := client.ExportAssets(ctx, req)
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
						Name:  "batch-get-assets-history",
						Usage: "batch-get-assets-history parent",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "asset-names", Usage: "A list of the full names of the assets.", Required: false},
							&cli.StringFlag{Name: "content-type", Usage: "The content type.", Required: false},
							&cli.StringSliceFlag{Name: "relationship-types", Usage: "A list of relationship types to output, for example:.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.BatchGetAssetsHistoryRequest{
								AssetNames:        cmd.StringSlice("asset-names"),
								ContentType:       assetpb.ContentType(assetpb.ContentType_value[cmd.String("content-type")]),
								RelationshipTypes: cmd.StringSlice("relationship-types"),
							}

							resp, err := client.BatchGetAssetsHistory(ctx, req)
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
						Name:  "query-assets",
						Usage: "query-assets parent",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of rows to return in the results.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from previous `QueryAssets`.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.QueryAssetsRequest{
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.QueryAssets(ctx, req)
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
				Name:  "resource",
				Usage: "Manage resource resources",
				Commands: []*cli.Command{

					{
						Name:  "analyze-move",
						Usage: "analyze-move resource",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination-parent", Usage: "Name of the Google Cloud folder or organization to reparent the.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Analysis view indicating what information should be included in the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.AnalyzeMoveRequest{
								DestinationParent: cmd.String("destination-parent"),
								View:              assetpb.AnalyzeMoveRequest_AnalysisView(assetpb.AnalyzeMoveRequest_AnalysisView_value[cmd.String("view")]),
							}

							resp, err := client.AnalyzeMove(ctx, req)
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
				Name:  "saved-queries",
				Usage: "Manage saved-queries resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "saved-query-id", Usage: "The ID to use for the saved query, which must be unique in the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.CreateSavedQueryRequest{
								SavedQueryId: cmd.String("saved-query-id"),
							}

							resp, err := client.CreateSavedQuery(ctx, req)
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
						Usage: "describe saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "saved-querie", Usage: "The ID of the saved querie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("savedQueries/%s", cmd.String("saved-querie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.GetSavedQueryRequest{
								Name: name,
							}

							resp, err := client.GetSavedQuery(ctx, req)
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
						Usage: "list saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The expression to filter resources.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of saved queries to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListSavedQueries` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.ListSavedQueriesRequest{
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSavedQueries(ctx, req)
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
						Usage: "update saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "saved-querie", Usage: "The ID of the saved querie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "saved_query.name" not yet supported.
							saved_query_name := fmt.Sprintf("savedQueries/%s", cmd.String("saved-querie"))
							fmt.Printf("Executing update on %s\n", saved_query_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete saved-queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "saved-querie", Usage: "The ID of the saved querie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("savedQueries/%s", cmd.String("saved-querie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSavedQuery on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := asset.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assetpb.DeleteSavedQueryRequest{
								Name: name,
							}

							if err := client.DeleteSavedQuery(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "scope",
				Usage: "Manage scope resources",
				Commands: []*cli.Command{

					{
						Name:  "search-all-resources",
						Usage: "search-all-resources scope",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "asset-types", Usage: "A list of asset types that this request searches for.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields specifying the sorting order of.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size for search result pagination.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
							&cli.StringFlag{Name: "query", Usage: "The query statement.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "A scope can be a project, a folder, or an organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search-all-resources...")
							return nil
						},
					},

					{
						Name:  "search-all-iam-policies",
						Usage: "search-all-iam-policies scope",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "asset-types", Usage: "A list of asset types that the IAM policies are attached to.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "A comma-separated list of fields specifying the sorting order of.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size for search result pagination.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, retrieve the next batch of results from the preceding.", Required: false},
							&cli.StringFlag{Name: "query", Usage: "The query statement.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "A scope can be a project, a folder, or an organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search-all-iam-policies...")
							return nil
						},
					},

					{
						Name:  "analyze-org-policies",
						Usage: "analyze-org-policies scope",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "constraint", Usage: "The name of the constraint to analyze organization policies for.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The expression to filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The pagination token to retrieve the next page.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "The organization to scope the request.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing analyze-org-policies...")
							return nil
						},
					},

					{
						Name:  "analyze-org-policy-governed-containers",
						Usage: "analyze-org-policy-governed-containers scope",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "constraint", Usage: "The name of the constraint to analyze governed containers for.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The expression to filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The pagination token to retrieve the next page.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "The organization to scope the request.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing analyze-org-policy-governed-containers...")
							return nil
						},
					},

					{
						Name:  "analyze-org-policy-governed-assets",
						Usage: "analyze-org-policy-governed-assets scope",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "constraint", Usage: "The name of the constraint to analyze governed assets for.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The expression to filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The pagination token to retrieve the next page.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "The organization to scope the request.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing analyze-org-policy-governed-assets...")
							return nil
						},
					},
				},
			},
		},
	}
}
