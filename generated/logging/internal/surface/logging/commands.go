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

package logging

import (
	logging "cloud.google.com/go/logging/apiv2"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the logging command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "logging",
		Usage: "manage Cloud Logging API resources",
		Commands: []*cli.Command{
			{
				Name:  "buckets",
				Usage: "Manage buckets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s", cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListBucketsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListBuckets(ctx, req)
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
						Usage: "describe buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.GetBucketRequest{
								Name: name,
							}

							resp, err := client.GetBucket(ctx, req)
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
						Usage: "create buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket-id", Usage: "A client-assigned identifier such as `\"my-bucket\"`.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s", cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateBucketRequest{
								Parent:   parent,
								BucketId: cmd.String("bucket-id"),
							}

							op, err := client.CreateBucketAsync(ctx, req)
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
						Name:  "update-async",
						Usage: "update-async buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UpdateBucketRequest{
								Name: name,
							}

							op, err := client.UpdateBucketAsync(ctx, req)
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
						Name:  "create",
						Usage: "create buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket-id", Usage: "A client-assigned identifier such as `\"my-bucket\"`.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s", cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateBucketRequest{
								Parent:   parent,
								BucketId: cmd.String("bucket-id"),
							}

							resp, err := client.CreateBucket(ctx, req)
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
						Usage: "update buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UpdateBucketRequest{
								Name: name,
							}

							resp, err := client.UpdateBucket(ctx, req)
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
						Usage: "delete buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteBucket on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.DeleteBucketRequest{
								Name: name,
							}

							if err := client.DeleteBucket(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "undelete",
						Usage: "undelete buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute UndeleteBucket on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UndeleteBucketRequest{
								Name: name,
							}

							if err := client.UndeleteBucket(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "cmek-settings",
				Usage: "Manage cmek-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe cmek-settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.GetCmekSettingsRequest{}

							resp, err := client.GetCmekSettings(ctx, req)
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
						Usage: "update cmek-settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UpdateCmekSettingsRequest{}

							resp, err := client.UpdateCmekSettings(ctx, req)
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
				Name:  "entries",
				Usage: "Manage entries resources",
				Commands: []*cli.Command{

					{
						Name:  "write",
						Usage: "write entries",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "dry-run", Usage: "If true, the request should expect normal response, but the.", Required: false},
							&cli.StringFlag{Name: "log-name", Usage: "A default log resource name that is assigned to all log entries.", Required: false},
							&cli.BoolFlag{Name: "partial-success", Usage: "Whether a batch's valid entries should be written even if some.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.WriteLogEntriesRequest{
								LogName:        cmd.String("log-name"),
								PartialSuccess: cmd.Bool("partial-success"),
								DryRun:         cmd.Bool("dry-run"),
							}

							resp, err := client.WriteLogEntries(ctx, req)
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
						Usage: "list entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only log entries that match the filter are returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "How the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
							&cli.StringSliceFlag{Name: "resource-names", Usage: "Names of one or more parent resources from which to.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListLogEntriesRequest{
								ResourceNames: cmd.StringSlice("resource-names"),
								Filter:        cmd.String("filter"),
								OrderBy:       cmd.String("order-by"),
								PageSize:      int32(cmd.Int("page-size")),
								PageToken:     cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListLogEntries(ctx, req)
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
						Name:  "tail",
						Usage: "tail entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only log entries that match the filter are returned.", Required: false},
							&cli.StringSliceFlag{Name: "resource-names", Usage: "Name of a parent resource from which to retrieve log entries:.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.TailLogEntriesRequest{
								ResourceNames: cmd.StringSlice("resource-names"),
								Filter:        cmd.String("filter"),
							}

							resp, err := client.TailLogEntries(ctx, req)
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
						Name:  "copy",
						Usage: "copy entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "destination", Usage: "Destination to which to copy log entries.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter specifying which log entries to copy.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CopyLogEntriesRequest{
								Filter:      cmd.String("filter"),
								Destination: cmd.String("destination"),
							}

							op, err := client.CopyLogEntries(ctx, req)
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
				},
			},
			{
				Name:  "exclusions",
				Usage: "Manage exclusions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list exclusions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListExclusionsRequest{
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListExclusions(ctx, req)
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
						Usage: "describe exclusions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exclusion", Usage: "The ID of the exclusion.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("exclusions/%s", cmd.String("exclusion"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.GetExclusionRequest{
								Name: name,
							}

							resp, err := client.GetExclusion(ctx, req)
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
						Usage: "create exclusions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateExclusionRequest{}

							resp, err := client.CreateExclusion(ctx, req)
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
						Usage: "update exclusions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exclusion", Usage: "The ID of the exclusion.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("exclusions/%s", cmd.String("exclusion"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UpdateExclusionRequest{
								Name: name,
							}

							resp, err := client.UpdateExclusion(ctx, req)
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
						Usage: "delete exclusions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exclusion", Usage: "The ID of the exclusion.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("exclusions/%s", cmd.String("exclusion"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteExclusion on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.DeleteExclusionRequest{
								Name: name,
							}

							if err := client.DeleteExclusion(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "links",
				Usage: "Manage links resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "link-id", Usage: "The ID to use for the link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateLinkRequest{
								Parent: parent,
								LinkId: cmd.String("link-id"),
							}

							op, err := client.CreateLink(ctx, req)
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
						Name:  "delete",
						Usage: "delete links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "link", Usage: "The ID of the link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s/links/%s", cmd.String("location"), cmd.String("bucket"), cmd.String("link"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteLink %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.DeleteLinkRequest{
								Name: name,
							}

							op, err := client.DeleteLink(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListLinksRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListLinks(ctx, req)
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
						Usage: "describe links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "link", Usage: "The ID of the link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s/links/%s", cmd.String("location"), cmd.String("bucket"), cmd.String("link"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.GetLinkRequest{
								Name: name,
							}

							resp, err := client.GetLink(ctx, req)
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
				Name:  "logs",
				Usage: "Manage logs resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete logs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "log", Usage: "The ID of the log.", Required: true},
							&cli.StringFlag{Name: "log-name", Usage: "The resource name of the log to delete:.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							log_name := fmt.Sprintf("projects/%s/logs/%s", cmd.String("project"), cmd.String("log"))
							fmt.Printf("Executing delete on %s\n", log_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list logs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
							&cli.StringSliceFlag{Name: "resource-names", Usage: "List of resource names to list logs for:.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListLogsRequest{
								ResourceNames: cmd.StringSlice("resource-names"),
								PageSize:      int32(cmd.Int("page-size")),
								PageToken:     cmd.String("page-token"),
							}

							resp, err := client.ListLogs(ctx, req)
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
				Name:  "metrics",
				Usage: "Manage metrics resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list metrics",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListLogMetricsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListLogMetrics(ctx, req)
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
						Usage: "describe metrics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metric", Usage: "The ID of the metric.", Required: true},
							&cli.StringFlag{Name: "metric-name", Usage: "The resource name of the desired metric:.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							metric_name := fmt.Sprintf("projects/%s/metrics/%s", cmd.String("project"), cmd.String("metric"))
							fmt.Printf("Executing describe on %s\n", metric_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create metrics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateLogMetricRequest{
								Parent: parent,
							}

							resp, err := client.CreateLogMetric(ctx, req)
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
						Usage: "update metrics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metric", Usage: "The ID of the metric.", Required: true},
							&cli.StringFlag{Name: "metric-name", Usage: "The resource name of the metric to update:.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							metric_name := fmt.Sprintf("projects/%s/metrics/%s", cmd.String("project"), cmd.String("metric"))
							fmt.Printf("Executing update on %s\n", metric_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete metrics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metric", Usage: "The ID of the metric.", Required: true},
							&cli.StringFlag{Name: "metric-name", Usage: "The resource name of the metric to delete:.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							metric_name := fmt.Sprintf("projects/%s/metrics/%s", cmd.String("project"), cmd.String("metric"))
							fmt.Printf("Executing delete on %s\n", metric_name)
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
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListMonitoredResourceDescriptorsRequest{
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
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s", cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/operations/%s", cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/operations/%s", cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s", cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/operations/%s", cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/operations/%s", cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s", cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/operations/%s", cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/operations/%s", cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "settings",
				Usage: "Manage settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.GetSettingsRequest{}

							resp, err := client.GetSettings(ctx, req)
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
						Usage: "update settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UpdateSettingsRequest{}

							resp, err := client.UpdateSettings(ctx, req)
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
				Name:  "sinks",
				Usage: "Manage sinks resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list sinks",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListSinksRequest{
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListSinks(ctx, req)
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
						Usage: "describe sinks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "sink", Usage: "The ID of the sink.", Required: true},
							&cli.StringFlag{Name: "sink-name", Usage: "The resource name of the sink:.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							sink_name := fmt.Sprintf("sinks/%s", cmd.String("sink"))
							fmt.Printf("Executing describe on %s\n", sink_name)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create sinks",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "unique-writer-identity", Usage: "Determines the kind of IAM identity returned as `writer_identity`.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateSinkRequest{
								UniqueWriterIdentity: cmd.Bool("unique-writer-identity"),
							}

							resp, err := client.CreateSink(ctx, req)
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
						Usage: "update sinks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "sink", Usage: "The ID of the sink.", Required: true},
							&cli.StringFlag{Name: "sink-name", Usage: "The full resource name of the sink to update, including the.", Required: true},
							&cli.BoolFlag{Name: "unique-writer-identity", Usage: "See [sinks.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							sink_name := fmt.Sprintf("sinks/%s", cmd.String("sink"))
							fmt.Printf("Executing update on %s\n", sink_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete sinks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "sink", Usage: "The ID of the sink.", Required: true},
							&cli.StringFlag{Name: "sink-name", Usage: "The full resource name of the sink to delete, including the.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							sink_name := fmt.Sprintf("sinks/%s", cmd.String("sink"))
							fmt.Printf("Executing delete on %s\n", sink_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "views",
				Usage: "Manage views resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, then retrieve the next batch of results from the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.ListViewsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListViews(ctx, req)
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
						Usage: "describe views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The ID of the view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s/views/%s", cmd.String("location"), cmd.String("bucket"), cmd.String("view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.GetViewRequest{
								Name: name,
							}

							resp, err := client.GetView(ctx, req)
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
						Usage: "create views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "view-id", Usage: "A client-assigned identifier such as `\"my-view\"`.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("locations/%s/buckets/%s", cmd.String("location"), cmd.String("bucket"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.CreateViewRequest{
								Parent: parent,
								ViewId: cmd.String("view-id"),
							}

							resp, err := client.CreateView(ctx, req)
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
						Usage: "update views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The ID of the view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s/views/%s", cmd.String("location"), cmd.String("bucket"), cmd.String("view"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.UpdateViewRequest{
								Name: name,
							}

							resp, err := client.UpdateView(ctx, req)
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
						Usage: "delete views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "bucket", Usage: "The ID of the bucket.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The ID of the view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/buckets/%s/views/%s", cmd.String("location"), cmd.String("bucket"), cmd.String("view"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteView on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := logging.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &loggingpb.DeleteViewRequest{
								Name: name,
							}

							if err := client.DeleteView(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
		},
	}
}
