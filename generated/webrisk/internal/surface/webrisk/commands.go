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

package webrisk

import (
	webrisk "cloud.google.com/go/webrisk/apiv1"
	"cloud.google.com/go/webrisk/apiv1/webriskpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the webrisk command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "webrisk",
		Usage: "manage Web Risk API resources",
		Commands: []*cli.Command{
			{
				Name:  "hashes",
				Usage: "Manage hashes resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search hashes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hash-prefix", Usage: "A hash prefix, consisting of the most significant 4-32 bytes of a SHA256.", Required: false},
							&cli.StringSliceFlag{Name: "threat-types", Usage: "The ThreatLists to search in.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := webrisk.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &webriskpb.SearchHashesRequest{
								HashPrefix:  []byte(cmd.String("hash-prefix")),
								ThreatTypes: cmd.StringSlice("threat-types"),
							}

							resp, err := client.SearchHashes(ctx, req)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "submissions",
				Usage: "Manage submissions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create submissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := webrisk.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &webriskpb.CreateSubmissionRequest{
								Parent: parent,
							}

							resp, err := client.CreateSubmission(ctx, req)
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
				Name:  "threat-lists",
				Usage: "Manage threat-lists resources",
				Commands: []*cli.Command{

					{
						Name:  "compute-diff",
						Usage: "compute-diff threat-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "threat-type", Usage: "The threat list to update.", Required: true},
							&cli.StringFlag{Name: "version-token", Usage: "The current version token of the client for the requested list (the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := webrisk.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &webriskpb.ComputeThreatListDiffRequest{
								ThreatType:   webriskpb.ThreatType(webriskpb.ThreatType_value[cmd.String("threat-type")]),
								VersionToken: []byte(cmd.String("version-token")),
							}

							resp, err := client.ComputeThreatListDiff(ctx, req)
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
				Name:  "uris",
				Usage: "Manage uris resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search uris",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "threat-types", Usage: "The ThreatLists to search in.", Required: true},
							&cli.StringFlag{Name: "uri", Usage: "The URI to be checked for matches.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := webrisk.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &webriskpb.SearchUrisRequest{
								Uri:         cmd.String("uri"),
								ThreatTypes: cmd.StringSlice("threat-types"),
							}

							resp, err := client.SearchUris(ctx, req)
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
						Name:  "submit",
						Usage: "submit uris",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := webrisk.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &webriskpb.SubmitUriRequest{
								Parent: parent,
							}

							op, err := client.SubmitUri(ctx, req)
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
		},
	}
}
