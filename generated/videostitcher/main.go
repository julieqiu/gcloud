package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	stitcher "cloud.google.com/go/video/stitcher/apiv1"
	"cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
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
				Name:  "videostitcher",
				Usage: "manage Video Stitcher API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "cdn-keys",
						Usage: "Manage cdn-keys resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create cdn-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cdn-key-id", Usage: "The cdn key id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "hostname", Usage: "The hostname.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.CreateCdnKeyRequest{Parent: parent}
									req.CdnKeyId = cmd.String("cdn-key-id")
									req.CdnKey = &stitcherpb.CdnKey{
										Name:     cmd.String("name"),
										Hostname: cmd.String("hostname"),
									}
									op, err := client.CreateCdnKey(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "list cdn-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListCdnKeysRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCdnKeys(ctx, req)
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
								Usage: "describe cdn-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cdn_key", Usage: "The cdn_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/cdnKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("cdn_key"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetCdnKeyRequest{Name: name}
									resp, err := client.GetCdnKey(ctx, req)
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
								Usage: "delete cdn-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cdn_key", Usage: "The cdn_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/cdnKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("cdn_key"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.DeleteCdnKeyRequest{Name: name}
									op, err := client.DeleteCdnKey(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update cdn-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "cdn_key", Usage: "The cdn_key.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "hostname", Usage: "The hostname.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/cdnKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("cdn_key"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.UpdateCdnKeyRequest{}
									req.CdnKey = &stitcherpb.CdnKey{
										Name:     name,
										Name:     cmd.String("name"),
										Hostname: cmd.String("hostname"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("hostname") {
										paths = append(paths, "hostname")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateCdnKey(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
						Name:  "live-ad-tag-details",
						Usage: "Manage live-ad-tag-details resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list live-ad-tag-details",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live_session", Usage: "The live_session.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/liveSessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("live_session"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListLiveAdTagDetailsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLiveAdTagDetails(ctx, req)
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
								Usage: "describe live-ad-tag-details",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live_session", Usage: "The live_session.", Required: true},
									&cli.StringFlag{Name: "live_ad_tag_detail", Usage: "The live_ad_tag_detail.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/liveSessions/%s/liveAdTagDetails/%s", cmd.String("project"), cmd.String("location"), cmd.String("live_session"), cmd.String("live_ad_tag_detail"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetLiveAdTagDetailRequest{Name: name}
									resp, err := client.GetLiveAdTagDetail(ctx, req)
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
						Name:  "live-configs",
						Usage: "Manage live-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create live-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live-config-id", Usage: "The live config id.", Required: true},
									&cli.StringFlag{Name: "source-uri", Usage: "The source uri.", Required: true},
									&cli.StringFlag{Name: "ad-tag-uri", Usage: "The ad tag uri.", Required: false},
									&cli.StringFlag{Name: "default-slate", Usage: "The default slate.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.CreateLiveConfigRequest{Parent: parent}
									req.LiveConfigId = cmd.String("live-config-id")
									req.LiveConfig = &stitcherpb.LiveConfig{
										SourceUri:    cmd.String("source-uri"),
										AdTagUri:     cmd.String("ad-tag-uri"),
										DefaultSlate: cmd.String("default-slate"),
									}
									op, err := client.CreateLiveConfig(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "list live-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListLiveConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLiveConfigs(ctx, req)
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
								Usage: "describe live-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live_config", Usage: "The live_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/liveConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("live_config"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetLiveConfigRequest{Name: name}
									resp, err := client.GetLiveConfig(ctx, req)
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
								Usage: "delete live-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live_config", Usage: "The live_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/liveConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("live_config"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.DeleteLiveConfigRequest{Name: name}
									op, err := client.DeleteLiveConfig(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update live-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live_config", Usage: "The live_config.", Required: true},
									&cli.StringFlag{Name: "source-uri", Usage: "The source uri.", Required: false},
									&cli.StringFlag{Name: "ad-tag-uri", Usage: "The ad tag uri.", Required: false},
									&cli.StringFlag{Name: "default-slate", Usage: "The default slate.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/liveConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("live_config"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.UpdateLiveConfigRequest{}
									req.LiveConfig = &stitcherpb.LiveConfig{
										Name:         name,
										SourceUri:    cmd.String("source-uri"),
										AdTagUri:     cmd.String("ad-tag-uri"),
										DefaultSlate: cmd.String("default-slate"),
									}
									var paths []string
									if cmd.IsSet("source-uri") {
										paths = append(paths, "source_uri")
									}
									if cmd.IsSet("ad-tag-uri") {
										paths = append(paths, "ad_tag_uri")
									}
									if cmd.IsSet("default-slate") {
										paths = append(paths, "default_slate")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateLiveConfig(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
						Name:  "live-sessions",
						Usage: "Manage live-sessions resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create live-sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live-config", Usage: "The live config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.CreateLiveSessionRequest{Parent: parent}
									req.LiveSession = &stitcherpb.LiveSession{
										LiveConfig: cmd.String("live-config"),
									}
									resp, err := client.CreateLiveSession(ctx, req)
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
								Usage: "describe live-sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "live_session", Usage: "The live_session.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/liveSessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("live_session"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetLiveSessionRequest{Name: name}
									resp, err := client.GetLiveSession(ctx, req)
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
						Name:  "operations",
						Usage: "Manage operations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
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
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
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
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "slates",
						Usage: "Manage slates resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create slates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "slate-id", Usage: "The slate id.", Required: true},
									&cli.StringFlag{Name: "uri", Usage: "The uri.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.CreateSlateRequest{Parent: parent}
									req.SlateId = cmd.String("slate-id")
									req.Slate = &stitcherpb.Slate{
										Uri: cmd.String("uri"),
									}
									op, err := client.CreateSlate(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "list slates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListSlatesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSlates(ctx, req)
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
								Usage: "describe slates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "slate", Usage: "The slate.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/slates/%s", cmd.String("project"), cmd.String("location"), cmd.String("slate"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetSlateRequest{Name: name}
									resp, err := client.GetSlate(ctx, req)
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
								Usage: "update slates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "slate", Usage: "The slate.", Required: true},
									&cli.StringFlag{Name: "uri", Usage: "The uri.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/slates/%s", cmd.String("project"), cmd.String("location"), cmd.String("slate"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.UpdateSlateRequest{}
									req.Slate = &stitcherpb.Slate{
										Name: name,
										Uri:  cmd.String("uri"),
									}
									var paths []string
									if cmd.IsSet("uri") {
										paths = append(paths, "uri")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateSlate(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "delete slates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "slate", Usage: "The slate.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/slates/%s", cmd.String("project"), cmd.String("location"), cmd.String("slate"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.DeleteSlateRequest{Name: name}
									op, err := client.DeleteSlate(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "vod-ad-tag-details",
						Usage: "Manage vod-ad-tag-details resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list vod-ad-tag-details",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_session", Usage: "The vod_session.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/vodSessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_session"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListVodAdTagDetailsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVodAdTagDetails(ctx, req)
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
								Usage: "describe vod-ad-tag-details",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_session", Usage: "The vod_session.", Required: true},
									&cli.StringFlag{Name: "vod_ad_tag_detail", Usage: "The vod_ad_tag_detail.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vodSessions/%s/vodAdTagDetails/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_session"), cmd.String("vod_ad_tag_detail"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetVodAdTagDetailRequest{Name: name}
									resp, err := client.GetVodAdTagDetail(ctx, req)
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
						Name:  "vod-configs",
						Usage: "Manage vod-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create vod-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod-config-id", Usage: "The vod config id.", Required: true},
									&cli.StringFlag{Name: "source-uri", Usage: "The source uri.", Required: true},
									&cli.StringFlag{Name: "ad-tag-uri", Usage: "The ad tag uri.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.CreateVodConfigRequest{Parent: parent}
									req.VodConfigId = cmd.String("vod-config-id")
									req.VodConfig = &stitcherpb.VodConfig{
										SourceUri: cmd.String("source-uri"),
										AdTagUri:  cmd.String("ad-tag-uri"),
									}
									op, err := client.CreateVodConfig(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "list vod-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListVodConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVodConfigs(ctx, req)
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
								Usage: "describe vod-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_config", Usage: "The vod_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vodConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_config"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetVodConfigRequest{Name: name}
									resp, err := client.GetVodConfig(ctx, req)
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
								Usage: "delete vod-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_config", Usage: "The vod_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vodConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_config"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.DeleteVodConfigRequest{Name: name}
									op, err := client.DeleteVodConfig(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update vod-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_config", Usage: "The vod_config.", Required: true},
									&cli.StringFlag{Name: "source-uri", Usage: "The source uri.", Required: false},
									&cli.StringFlag{Name: "ad-tag-uri", Usage: "The ad tag uri.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vodConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_config"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.UpdateVodConfigRequest{}
									req.VodConfig = &stitcherpb.VodConfig{
										Name:      name,
										SourceUri: cmd.String("source-uri"),
										AdTagUri:  cmd.String("ad-tag-uri"),
									}
									var paths []string
									if cmd.IsSet("source-uri") {
										paths = append(paths, "source_uri")
									}
									if cmd.IsSet("ad-tag-uri") {
										paths = append(paths, "ad_tag_uri")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateVodConfig(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
						Name:  "vod-sessions",
						Usage: "Manage vod-sessions resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create vod-sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "source-uri", Usage: "The source uri.", Required: false},
									&cli.StringFlag{Name: "ad-tag-uri", Usage: "The ad tag uri.", Required: false},
									&cli.StringFlag{Name: "vod-config", Usage: "The vod config.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.CreateVodSessionRequest{Parent: parent}
									req.VodSession = &stitcherpb.VodSession{
										SourceUri: cmd.String("source-uri"),
										AdTagUri:  cmd.String("ad-tag-uri"),
										VodConfig: cmd.String("vod-config"),
									}
									resp, err := client.CreateVodSession(ctx, req)
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
								Usage: "describe vod-sessions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_session", Usage: "The vod_session.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vodSessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_session"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetVodSessionRequest{Name: name}
									resp, err := client.GetVodSession(ctx, req)
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
						Name:  "vod-stitch-details",
						Usage: "Manage vod-stitch-details resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list vod-stitch-details",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_session", Usage: "The vod_session.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/vodSessions/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_session"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &stitcherpb.ListVodStitchDetailsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVodStitchDetails(ctx, req)
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
								Usage: "describe vod-stitch-details",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "vod_session", Usage: "The vod_session.", Required: true},
									&cli.StringFlag{Name: "vod_stitch_detail", Usage: "The vod_stitch_detail.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/vodSessions/%s/vodStitchDetails/%s", cmd.String("project"), cmd.String("location"), cmd.String("vod_session"), cmd.String("vod_stitch_detail"))
									client, err := stitcher.NewVideoStitcherClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &stitcherpb.GetVodStitchDetailRequest{Name: name}
									resp, err := client.GetVodStitchDetail(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
