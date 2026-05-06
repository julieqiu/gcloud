package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	asset "cloud.google.com/go/asset/apiv1"
	"cloud.google.com/go/asset/apiv1/assetpb"
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
				Name:  "cloudasset",
				Usage: "manage Cloud Asset API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "assets",
						Usage: "Manage assets resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list assets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
									&cli.StringFlag{Name: "feed-id", Usage: "The feed id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.CreateFeedRequest{Parent: parent}
									req.FeedId = cmd.String("feed-id")
									req.Feed = &assetpb.Feed{
										Name: cmd.String("name"),
									}
									resp, err := client.CreateFeed(ctx, req)
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
								Usage: "describe feeds",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "feed", Usage: "The feed.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/feeds/%s", cmd.String("project"), cmd.String("feed"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.GetFeedRequest{Name: name}
									resp, err := client.GetFeed(ctx, req)
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
								Usage: "list feeds",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update feeds",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "feed", Usage: "The feed.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/feeds/%s", cmd.String("project"), cmd.String("feed"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.UpdateFeedRequest{}
									req.Feed = &assetpb.Feed{
										Name: name,
										Name: cmd.String("name"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateFeed(ctx, req)
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
								Usage: "delete feeds",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "feed", Usage: "The feed.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/feeds/%s", cmd.String("project"), cmd.String("feed"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.DeleteFeedRequest{Name: name}
									if err := client.DeleteFeed(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
									&cli.StringFlag{Name: "saved-query-id", Usage: "The saved query id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.CreateSavedQueryRequest{Parent: parent}
									req.SavedQueryId = cmd.String("saved-query-id")
									req.SavedQuery = &assetpb.SavedQuery{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreateSavedQuery(ctx, req)
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
								Usage: "describe saved-queries",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "saved_query", Usage: "The saved_query.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/savedQueries/%s", cmd.String("project"), cmd.String("saved_query"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.GetSavedQueryRequest{Name: name}
									resp, err := client.GetSavedQuery(ctx, req)
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
								Usage: "list saved-queries",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &assetpb.ListSavedQueriesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSavedQueries(ctx, req)
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
								Name:  "update",
								Usage: "update saved-queries",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "saved_query", Usage: "The saved_query.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/savedQueries/%s", cmd.String("project"), cmd.String("saved_query"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.UpdateSavedQueryRequest{}
									req.SavedQuery = &assetpb.SavedQuery{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSavedQuery(ctx, req)
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
								Usage: "delete saved-queries",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "saved_query", Usage: "The saved_query.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/savedQueries/%s", cmd.String("project"), cmd.String("saved_query"))
									client, err := asset.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &assetpb.DeleteSavedQueryRequest{Name: name}
									if err := client.DeleteSavedQuery(ctx, req); err != nil {
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
