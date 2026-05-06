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

package cloudtrace

import (
	cloudtrace "cloud.google.com/go/cloudtrace/apiv2"
	"cloud.google.com/go/cloudtrace/apiv2/cloudtracepb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the cloudtrace command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudtrace",
		Usage: "manage Stackdriver Trace API resources",
		Commands: []*cli.Command{
			{
				Name:  "spans",
				Usage: "Manage spans resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create spans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-span-id", Usage: "The `[SPAN_ID]` of this span's parent span.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "span", Usage: "The ID of the span.", Required: true},
							&cli.StringFlag{Name: "span-kind", Usage: "Distinguishes between spans generated in a particular context.", Required: false},
							&cli.StringFlag{Name: "trace", Usage: "The ID of the trace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/traces/%s/spans/%s", cmd.String("project"), cmd.String("trace"), cmd.String("span"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudtrace.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudtracepb.Span{
								Name:         name,
								ParentSpanId: cmd.String("parent-span-id"),
								SpanKind:     cloudtracepb.Span_SpanKind(cloudtracepb.Span_SpanKind_value[cmd.String("span-kind")]),
							}

							resp, err := client.CreateSpan(ctx, req)
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
				Name:  "traces",
				Usage: "Manage traces resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-write",
						Usage: "batch-write traces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute BatchWriteSpans on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudtrace.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudtracepb.BatchWriteSpansRequest{
								Name: name,
							}

							if err := client.BatchWriteSpans(ctx, req); err != nil {
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
