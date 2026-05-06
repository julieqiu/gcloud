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

package securitycenter

import (
	securitycenter "cloud.google.com/go/securitycenter/apiv2"
	"cloud.google.com/go/securitycenter/apiv2/securitycenterpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the securitycenter command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "securitycenter",
		Usage: "manage Security Command Center API resources",
		Commands: []*cli.Command{
			{
				Name:  "attack-paths",
				Usage: "Manage attack-paths resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list attack-paths",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression that filters the attack path in the response.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListAttackPathsResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "simulation", Usage: "The ID of the simulation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/simulations/%s", cmd.String("organization"), cmd.String("simulation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListAttackPathsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListAttackPaths(ctx, req)
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
				Name:  "big-query-exports",
				Usage: "Manage big-query-exports resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create big-query-exports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "big-query-export-id", Usage: "Unique identifier provided by the client within the parent scope.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.CreateBigQueryExportRequest{
								Parent:           parent,
								BigQueryExportId: cmd.String("big-query-export-id"),
							}

							resp, err := client.CreateBigQueryExport(ctx, req)
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
						Usage: "delete big-query-exports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "big-query-export", Usage: "The ID of the big query export.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", cmd.String("organization"), cmd.String("location"), cmd.String("big-query-export"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteBigQueryExport on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.DeleteBigQueryExportRequest{
								Name: name,
							}

							if err := client.DeleteBigQueryExport(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe big-query-exports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "big-query-export", Usage: "The ID of the big query export.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", cmd.String("organization"), cmd.String("location"), cmd.String("big-query-export"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetBigQueryExportRequest{
								Name: name,
							}

							resp, err := client.GetBigQueryExport(ctx, req)
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
						Usage: "list big-query-exports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of configs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListBigQueryExports` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListBigQueryExportsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBigQueryExports(ctx, req)
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
						Usage: "update big-query-exports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "big-query-export", Usage: "The ID of the big query export.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "big_query_export.name" not yet supported.
							big_query_export_name := fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", cmd.String("organization"), cmd.String("location"), cmd.String("big-query-export"))
							fmt.Printf("Executing update on %s\n", big_query_export_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "external-systems",
				Usage: "Manage external-systems resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update external-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-system", Usage: "The ID of the external system.", Required: true},
							&cli.StringFlag{Name: "finding", Usage: "The ID of the finding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "external_system.name" not yet supported.
							external_system_name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s/externalSystems/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"), cmd.String("external-system"))
							fmt.Printf("Executing update on %s\n", external_system_name)
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
						Name:  "bulk-mute",
						Usage: "bulk-mute findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Expression that identifies findings that should be updated.", Required: false},
							&cli.StringFlag{Name: "mute-state", Usage: "All findings matching the given filter will have their mute state.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.BulkMuteFindingsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								MuteState: securitycenterpb.BulkMuteFindingsRequest_MuteState(securitycenterpb.BulkMuteFindingsRequest_MuteState_value[cmd.String("mute-state")]),
							}

							op, err := client.BulkMuteFindings(ctx, req)
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
						Usage: "create findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "finding-id", Usage: "Unique identifier provided by the client within the parent scope.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/sources/%s/locations/%s", cmd.String("organization"), cmd.String("source"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.CreateFindingRequest{
								Parent:    parent,
								FindingId: cmd.String("finding-id"),
							}

							resp, err := client.CreateFinding(ctx, req)
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
						Name:  "group",
						Usage: "group findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Expression that defines the filter to apply across findings.", Required: false},
							&cli.StringFlag{Name: "group-by", Usage: "Expression that defines what assets fields to use for grouping.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `GroupFindingsResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GroupFindingsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								GroupBy:   cmd.String("group-by"),
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.GroupFindings(ctx, req)
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
						Name:  "list",
						Usage: "list findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Expression that defines the filter to apply across findings.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Expression that defines what fields and order to use for sorting.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListFindingsResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListFindingsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
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

					{
						Name:  "set-state",
						Usage: "set-state findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "finding", Usage: "The ID of the finding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
							&cli.StringFlag{Name: "state", Usage: "The desired State of the finding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.SetFindingStateRequest{
								Name:  name,
								State: securitycenterpb.Finding_State(securitycenterpb.Finding_State_value[cmd.String("state")]),
							}

							resp, err := client.SetFindingState(ctx, req)
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
						Name:  "set-mute",
						Usage: "set-mute findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "finding", Usage: "The ID of the finding.", Required: true},
							&cli.StringFlag{Name: "mute", Usage: "The desired state of the Mute.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.SetMuteRequest{
								Name: name,
								Mute: securitycenterpb.Finding_Mute(securitycenterpb.Finding_Mute_value[cmd.String("mute")]),
							}

							resp, err := client.SetMute(ctx, req)
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
						Usage: "update findings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "finding", Usage: "The ID of the finding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "finding.name" not yet supported.
							finding_name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
							fmt.Printf("Executing update on %s\n", finding_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "mute-configs",
				Usage: "Manage mute-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create mute-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mute-config-id", Usage: "Unique identifier provided by the client within the parent scope.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.CreateMuteConfigRequest{
								Parent:       parent,
								MuteConfigId: cmd.String("mute-config-id"),
							}

							resp, err := client.CreateMuteConfig(ctx, req)
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
						Usage: "delete mute-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "mute-config", Usage: "The ID of the mute config.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/muteConfigs/%s", cmd.String("organization"), cmd.String("mute-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteMuteConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.DeleteMuteConfigRequest{
								Name: name,
							}

							if err := client.DeleteMuteConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe mute-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "mute-config", Usage: "The ID of the mute config.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/muteConfigs/%s", cmd.String("organization"), cmd.String("mute-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetMuteConfigRequest{
								Name: name,
							}

							resp, err := client.GetMuteConfig(ctx, req)
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
						Usage: "list mute-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of configs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListMuteConfigs` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListMuteConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMuteConfigs(ctx, req)
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
						Usage: "update mute-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "mute-config", Usage: "The ID of the mute config.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "mute_config.name" not yet supported.
							mute_config_name := fmt.Sprintf("organizations/%s/muteConfigs/%s", cmd.String("organization"), cmd.String("mute-config"))
							fmt.Printf("Executing update on %s\n", mute_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "notification-configs",
				Usage: "Manage notification-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create notification-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "config-id", Usage: "Required.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.CreateNotificationConfigRequest{
								Parent:   parent,
								ConfigId: cmd.String("config-id"),
							}

							resp, err := client.CreateNotificationConfig(ctx, req)
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
						Usage: "delete notification-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notification-config", Usage: "The ID of the notification config.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/notificationConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteNotificationConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.DeleteNotificationConfigRequest{
								Name: name,
							}

							if err := client.DeleteNotificationConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe notification-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notification-config", Usage: "The ID of the notification config.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/notificationConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetNotificationConfigRequest{
								Name: name,
							}

							resp, err := client.GetNotificationConfig(ctx, req)
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
						Usage: "list notification-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListNotificationConfigsResponse`; indicates.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListNotificationConfigsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListNotificationConfigs(ctx, req)
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
						Usage: "update notification-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "notification-config", Usage: "The ID of the notification config.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "notification_config.name" not yet supported.
							notification_config_name := fmt.Sprintf("organizations/%s/locations/%s/notificationConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("notification-config"))
							fmt.Printf("Executing update on %s\n", notification_config_name)
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
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/operations", cmd.String("organization"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/operations/%s", cmd.String("organization"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/operations/%s", cmd.String("organization"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/operations/%s", cmd.String("organization"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "resource-value-configs",
				Usage: "Manage resource-value-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "batch-create",
						Usage: "batch-create resource-value-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.BatchCreateResourceValueConfigsRequest{
								Parent: parent,
							}

							resp, err := client.BatchCreateResourceValueConfigs(ctx, req)
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
						Usage: "delete resource-value-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "resource-value-config", Usage: "The ID of the resource value config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/resourceValueConfigs/%s", cmd.String("organization"), cmd.String("resource-value-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteResourceValueConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.DeleteResourceValueConfigRequest{
								Name: name,
							}

							if err := client.DeleteResourceValueConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe resource-value-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "resource-value-config", Usage: "The ID of the resource value config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/resourceValueConfigs/%s", cmd.String("organization"), cmd.String("resource-value-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetResourceValueConfigRequest{
								Name: name,
							}

							resp, err := client.GetResourceValueConfig(ctx, req)
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
						Usage: "list resource-value-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of configs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListResourceValueConfigs` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListResourceValueConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListResourceValueConfigs(ctx, req)
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
						Usage: "update resource-value-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "resource-value-config", Usage: "The ID of the resource value config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "resource_value_config.name" not yet supported.
							resource_value_config_name := fmt.Sprintf("organizations/%s/resourceValueConfigs/%s", cmd.String("organization"), cmd.String("resource-value-config"))
							fmt.Printf("Executing update on %s\n", resource_value_config_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "security-marks",
				Usage: "Manage security-marks resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update security-marks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "finding", Usage: "The ID of the finding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "security_marks.name" not yet supported.
							security_marks_name := fmt.Sprintf("organizations/%s/sources/%s/findings/%s/securityMarks", cmd.String("organization"), cmd.String("source"), cmd.String("finding"))
							fmt.Printf("Executing update on %s\n", security_marks_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "simulations",
				Usage: "Manage simulations resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe simulations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "simulation", Usage: "The ID of the simulation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/simulations/%s", cmd.String("organization"), cmd.String("simulation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetSimulationRequest{
								Name: name,
							}

							resp, err := client.GetSimulation(ctx, req)
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
				Name:  "sources",
				Usage: "Manage sources resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.CreateSourceRequest{
								Parent: parent,
							}

							resp, err := client.CreateSource(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Usage: "describe sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetSourceRequest{
								Name: name,
							}

							resp, err := client.GetSource(ctx, req)
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
						Usage: "list sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListSourcesResponse`; indicates.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListSourcesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListSources(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
						Usage: "update sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "source", Usage: "The ID of the source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "source.name" not yet supported.
							source_name := fmt.Sprintf("organizations/%s/sources/%s", cmd.String("organization"), cmd.String("source"))
							fmt.Printf("Executing update on %s\n", source_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "valued-resources",
				Usage: "Manage valued-resources resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe valued-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "simulation", Usage: "The ID of the simulation.", Required: true},
							&cli.StringFlag{Name: "valued-resource", Usage: "The ID of the valued resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/simulations/%s/valuedResources/%s", cmd.String("organization"), cmd.String("simulation"), cmd.String("valued-resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.GetValuedResourceRequest{
								Name: name,
							}

							resp, err := client.GetValuedResource(ctx, req)
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
						Usage: "list valued-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter expression that filters the valued resources in the response.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "The fields by which to order the valued resources response.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListValuedResourcesResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "simulation", Usage: "The ID of the simulation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/simulations/%s", cmd.String("organization"), cmd.String("simulation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securitycenter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securitycenterpb.ListValuedResourcesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListValuedResources(ctx, req)
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
		},
	}
}
