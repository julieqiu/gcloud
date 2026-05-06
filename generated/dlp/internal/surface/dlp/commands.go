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

package dlp

import (
	dlp "cloud.google.com/go/dlp/apiv2"
	"cloud.google.com/go/dlp/apiv2/dlppb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the dlp command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "dlp",
		Usage: "manage Sensitive Data Protection (DLP) resources",
		Commands: []*cli.Command{
			{
				Name:  "column-data-profiles",
				Usage: "Manage column-data-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list column-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Allows filtering.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by, followed by `asc` or `desc`.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListColumnDataProfilesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListColumnDataProfiles(ctx, req)
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
						Usage: "describe column-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "column-data-profile", Usage: "The ID of the column data profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/columnDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("column-data-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetColumnDataProfileRequest{
								Name: name,
							}

							resp, err := client.GetColumnDataProfile(ctx, req)
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
				Name:  "connections",
				Usage: "Manage connections resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create connections",
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
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateConnectionRequest{
								Parent: parent,
							}

							resp, err := client.CreateConnection(ctx, req)
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
						Usage: "describe connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetConnectionRequest{
								Name: name,
							}

							resp, err := client.GetConnection(ctx, req)
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
						Usage: "list connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Supported field/value: `state` - MISSING|AVAILABLE|ERROR.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results per page, max 1000.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token from a previous page to return the next set of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListConnections(ctx, req)
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
						Name:  "search",
						Usage: "search connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Supported field/value: - `state` - MISSING|AVAILABLE|ERROR.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results per page, max 1000.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token from a previous page to return the next set of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.SearchConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.SearchConnections(ctx, req)
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
						Name:  "delete",
						Usage: "delete connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteConnection on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteConnectionRequest{
								Name: name,
							}

							if err := client.DeleteConnection(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.UpdateConnectionRequest{
								Name: name,
							}

							resp, err := client.UpdateConnection(ctx, req)
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
				Name:  "content",
				Usage: "Manage content resources",
				Commands: []*cli.Command{

					{
						Name:  "inspect",
						Usage: "inspect content",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "inspect-template-name", Usage: "Template to use.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.InspectContentRequest{
								Parent:              parent,
								InspectTemplateName: cmd.String("inspect-template-name"),
								LocationId:          cmd.String("location-id"),
							}

							resp, err := client.InspectContent(ctx, req)
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
						Name:  "deidentify",
						Usage: "deidentify content",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deidentify-template-name", Usage: "Template to use.", Required: false},
							&cli.StringFlag{Name: "inspect-template-name", Usage: "Template to use.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeidentifyContentRequest{
								Parent:                 parent,
								InspectTemplateName:    cmd.String("inspect-template-name"),
								DeidentifyTemplateName: cmd.String("deidentify-template-name"),
								LocationId:             cmd.String("location-id"),
							}

							resp, err := client.DeidentifyContent(ctx, req)
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
						Name:  "reidentify",
						Usage: "reidentify content",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "inspect-template-name", Usage: "Template to use.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reidentify-template-name", Usage: "Template to use.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ReidentifyContentRequest{
								Parent:                 parent,
								InspectTemplateName:    cmd.String("inspect-template-name"),
								ReidentifyTemplateName: cmd.String("reidentify-template-name"),
								LocationId:             cmd.String("location-id"),
							}

							resp, err := client.ReidentifyContent(ctx, req)
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
				Name:  "deidentify-templates",
				Usage: "Manage deidentify-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "template-id", Usage: "The template id can contain uppercase and lowercase letters,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateDeidentifyTemplateRequest{
								Parent:     parent,
								TemplateId: cmd.String("template-id"),
								LocationId: cmd.String("location-id"),
							}

							resp, err := client.CreateDeidentifyTemplate(ctx, req)
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
						Usage: "update deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deidentify-template", Usage: "The ID of the deidentify template.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/deidentifyTemplates/%s", cmd.String("organization"), cmd.String("deidentify-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.UpdateDeidentifyTemplateRequest{
								Name: name,
							}

							resp, err := client.UpdateDeidentifyTemplate(ctx, req)
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
						Usage: "describe deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deidentify-template", Usage: "The ID of the deidentify template.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/deidentifyTemplates/%s", cmd.String("organization"), cmd.String("deidentify-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetDeidentifyTemplateRequest{
								Name: name,
							}

							resp, err := client.GetDeidentifyTemplate(ctx, req)
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
						Usage: "list deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by,.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListDeidentifyTemplatesRequest{
								Parent:     parent,
								PageToken:  cmd.String("page-token"),
								PageSize:   int32(cmd.Int("page-size")),
								OrderBy:    cmd.String("order-by"),
								LocationId: cmd.String("location-id"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeidentifyTemplates(ctx, req)
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
						Name:  "delete",
						Usage: "delete deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deidentify-template", Usage: "The ID of the deidentify template.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/deidentifyTemplates/%s", cmd.String("organization"), cmd.String("deidentify-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDeidentifyTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteDeidentifyTemplateRequest{
								Name: name,
							}

							if err := client.DeleteDeidentifyTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "discovery-configs",
				Usage: "Manage discovery-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "config-id", Usage: "The config ID can contain uppercase and lowercase letters,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateDiscoveryConfigRequest{
								Parent:   parent,
								ConfigId: cmd.String("config-id"),
							}

							resp, err := client.CreateDiscoveryConfig(ctx, req)
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
						Usage: "update discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "discovery-config", Usage: "The ID of the discovery config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/discoveryConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("discovery-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.UpdateDiscoveryConfigRequest{
								Name: name,
							}

							resp, err := client.UpdateDiscoveryConfig(ctx, req)
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
						Usage: "describe discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "discovery-config", Usage: "The ID of the discovery config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/discoveryConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("discovery-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetDiscoveryConfigRequest{
								Name: name,
							}

							resp, err := client.GetDiscoveryConfig(ctx, req)
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
						Usage: "list discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of config fields to order by,.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListDiscoveryConfigsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDiscoveryConfigs(ctx, req)
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
						Name:  "delete",
						Usage: "delete discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "discovery-config", Usage: "The ID of the discovery config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/discoveryConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("discovery-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDiscoveryConfig on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteDiscoveryConfigRequest{
								Name: name,
							}

							if err := client.DeleteDiscoveryConfig(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "dlp-jobs",
				Usage: "Manage dlp-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-id", Usage: "The job id can contain uppercase and lowercase letters,.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateDlpJobRequest{
								Parent:     parent,
								JobId:      cmd.String("job-id"),
								LocationId: cmd.String("location-id"),
							}

							resp, err := client.CreateDlpJob(ctx, req)
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
						Usage: "list dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Allows filtering.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by,.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The standard list page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The standard list page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The type of job.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListDlpJobsRequest{
								Parent:     parent,
								Filter:     cmd.String("filter"),
								PageSize:   int32(cmd.Int("page-size")),
								PageToken:  cmd.String("page-token"),
								Type:       dlppb.DlpJobType(dlppb.DlpJobType_value[cmd.String("type")]),
								OrderBy:    cmd.String("order-by"),
								LocationId: cmd.String("location-id"),
							}

							limit := cmd.Int("limit")
							it := client.ListDlpJobs(ctx, req)
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
						Usage: "describe dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp-job", Usage: "The ID of the dlp job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetDlpJobRequest{
								Name: name,
							}

							resp, err := client.GetDlpJob(ctx, req)
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
						Usage: "delete dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp-job", Usage: "The ID of the dlp job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDlpJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteDlpJobRequest{
								Name: name,
							}

							if err := client.DeleteDlpJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp-job", Usage: "The ID of the dlp job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelDlpJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CancelDlpJobRequest{
								Name: name,
							}

							if err := client.CancelDlpJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "hybrid-inspect",
						Usage: "hybrid-inspect dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp-job", Usage: "The ID of the dlp job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dlpJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("dlp-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.HybridInspectDlpJobRequest{
								Name: name,
							}

							resp, err := client.HybridInspectDlpJob(ctx, req)
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
						Name:  "finish",
						Usage: "finish dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp-job", Usage: "The ID of the dlp job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dlpJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("dlp-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute FinishDlpJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.FinishDlpJobRequest{
								Name: name,
							}

							if err := client.FinishDlpJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "file-store-data-profiles",
				Usage: "Manage file-store-data-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list file-store-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Allows filtering.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by, followed by `asc` or.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListFileStoreDataProfilesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListFileStoreDataProfiles(ctx, req)
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
						Usage: "describe file-store-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "file-store-data-profile", Usage: "The ID of the file store data profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/fileStoreDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("file-store-data-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetFileStoreDataProfileRequest{
								Name: name,
							}

							resp, err := client.GetFileStoreDataProfile(ctx, req)
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
						Usage: "delete file-store-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "file-store-data-profile", Usage: "The ID of the file store data profile.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/fileStoreDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("file-store-data-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteFileStoreDataProfile on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteFileStoreDataProfileRequest{
								Name: name,
							}

							if err := client.DeleteFileStoreDataProfile(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "image",
				Usage: "Manage image resources",
				Commands: []*cli.Command{

					{
						Name:  "redact",
						Usage: "redact image",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deidentify-template", Usage: "The full resource name of the de-identification template to use.", Required: false},
							&cli.BoolFlag{Name: "include-findings", Usage: "Whether the response should include findings along with the redacted.", Required: false},
							&cli.StringFlag{Name: "inspect-template", Usage: "The full resource name of the inspection template to use.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.RedactImageRequest{
								Parent:             parent,
								LocationId:         cmd.String("location-id"),
								IncludeFindings:    cmd.Bool("include-findings"),
								InspectTemplate:    cmd.String("inspect-template"),
								DeidentifyTemplate: cmd.String("deidentify-template"),
							}

							resp, err := client.RedactImage(ctx, req)
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
				Name:  "info-types",
				Usage: "Manage info-types resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "filter to only return infoTypes supported by certain parts of the.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "BCP-47 language code for localized infoType friendly.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListInfoTypesRequest{
								LanguageCode: cmd.String("language-code"),
								Filter:       cmd.String("filter"),
								LocationId:   cmd.String("location-id"),
							}

							resp, err := client.ListInfoTypes(ctx, req)
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
				Name:  "inspect-templates",
				Usage: "Manage inspect-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "template-id", Usage: "The template id can contain uppercase and lowercase letters,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateInspectTemplateRequest{
								Parent:     parent,
								TemplateId: cmd.String("template-id"),
							}

							resp, err := client.CreateInspectTemplate(ctx, req)
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
						Usage: "update inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "inspect-template", Usage: "The ID of the inspect template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/inspectTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("inspect-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.UpdateInspectTemplateRequest{
								Name: name,
							}

							resp, err := client.UpdateInspectTemplate(ctx, req)
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
						Usage: "describe inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "inspect-template", Usage: "The ID of the inspect template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/inspectTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("inspect-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetInspectTemplateRequest{
								Name: name,
							}

							resp, err := client.GetInspectTemplate(ctx, req)
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
						Usage: "list inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by,.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListInspectTemplatesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInspectTemplates(ctx, req)
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
						Name:  "delete",
						Usage: "delete inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "inspect-template", Usage: "The ID of the inspect template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/inspectTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("inspect-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteInspectTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteInspectTemplateRequest{
								Name: name,
							}

							if err := client.DeleteInspectTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "job-triggers",
				Usage: "Manage job-triggers resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "trigger-id", Usage: "The trigger id can contain uppercase and lowercase letters,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateJobTriggerRequest{
								Parent:     parent,
								TriggerId:  cmd.String("trigger-id"),
								LocationId: cmd.String("location-id"),
							}

							resp, err := client.CreateJobTrigger(ctx, req)
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
						Usage: "update job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-trigger", Usage: "The ID of the job trigger.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job-trigger"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.UpdateJobTriggerRequest{
								Name: name,
							}

							resp, err := client.UpdateJobTrigger(ctx, req)
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
						Name:  "hybrid-inspect",
						Usage: "hybrid-inspect job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-trigger", Usage: "The ID of the job trigger.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobTriggers/%s", cmd.String("project"), cmd.String("location"), cmd.String("job-trigger"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.HybridInspectJobTriggerRequest{
								Name: name,
							}

							resp, err := client.HybridInspectJobTrigger(ctx, req)
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
						Usage: "describe job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-trigger", Usage: "The ID of the job trigger.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job-trigger"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetJobTriggerRequest{
								Name: name,
							}

							resp, err := client.GetJobTrigger(ctx, req)
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
						Usage: "list job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Allows filtering.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of triggeredJob fields to order by,.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The type of jobs.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListJobTriggersRequest{
								Parent:     parent,
								PageToken:  cmd.String("page-token"),
								PageSize:   int32(cmd.Int("page-size")),
								OrderBy:    cmd.String("order-by"),
								Filter:     cmd.String("filter"),
								Type:       dlppb.DlpJobType(dlppb.DlpJobType_value[cmd.String("type")]),
								LocationId: cmd.String("location-id"),
							}

							limit := cmd.Int("limit")
							it := client.ListJobTriggers(ctx, req)
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
						Name:  "delete",
						Usage: "delete job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-trigger", Usage: "The ID of the job trigger.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job-trigger"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteJobTrigger on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteJobTriggerRequest{
								Name: name,
							}

							if err := client.DeleteJobTrigger(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "activate",
						Usage: "activate job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-trigger", Usage: "The ID of the job trigger.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job-trigger"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ActivateJobTriggerRequest{
								Name: name,
							}

							resp, err := client.ActivateJobTrigger(ctx, req)
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
				Name:  "project-data-profiles",
				Usage: "Manage project-data-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list project-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Allows filtering.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by, followed by `asc` or `desc`.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListProjectDataProfilesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListProjectDataProfiles(ctx, req)
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
						Usage: "describe project-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "project-data-profile", Usage: "The ID of the project data profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/projectDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("project-data-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetProjectDataProfileRequest{
								Name: name,
							}

							resp, err := client.GetProjectDataProfile(ctx, req)
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
				Name:  "stored-info-types",
				Usage: "Manage stored-info-types resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "stored-info-type-id", Usage: "The storedInfoType ID can contain uppercase and lowercase letters,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.CreateStoredInfoTypeRequest{
								Parent:           parent,
								StoredInfoTypeId: cmd.String("stored-info-type-id"),
								LocationId:       cmd.String("location-id"),
							}

							resp, err := client.CreateStoredInfoType(ctx, req)
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
						Usage: "update stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "stored-info-type", Usage: "The ID of the stored info type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/storedInfoTypes/%s", cmd.String("organization"), cmd.String("stored-info-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.UpdateStoredInfoTypeRequest{
								Name: name,
							}

							resp, err := client.UpdateStoredInfoType(ctx, req)
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
						Usage: "describe stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "stored-info-type", Usage: "The ID of the stored info type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/storedInfoTypes/%s", cmd.String("organization"), cmd.String("stored-info-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetStoredInfoTypeRequest{
								Name: name,
							}

							resp, err := client.GetStoredInfoType(ctx, req)
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
						Usage: "list stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location-id", Usage: "Deprecated.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by,.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListStoredInfoTypesRequest{
								Parent:     parent,
								PageToken:  cmd.String("page-token"),
								PageSize:   int32(cmd.Int("page-size")),
								OrderBy:    cmd.String("order-by"),
								LocationId: cmd.String("location-id"),
							}

							limit := cmd.Int("limit")
							it := client.ListStoredInfoTypes(ctx, req)
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
						Name:  "delete",
						Usage: "delete stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "stored-info-type", Usage: "The ID of the stored info type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/storedInfoTypes/%s", cmd.String("organization"), cmd.String("stored-info-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteStoredInfoType on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteStoredInfoTypeRequest{
								Name: name,
							}

							if err := client.DeleteStoredInfoType(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "table-data-profiles",
				Usage: "Manage table-data-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list table-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Allows filtering.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Comma-separated list of fields to order by, followed by `asc` or `desc`.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token to continue retrieval.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.ListTableDataProfilesRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
								OrderBy:   cmd.String("order-by"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListTableDataProfiles(ctx, req)
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
						Usage: "describe table-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "table-data-profile", Usage: "The ID of the table data profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/tableDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("table-data-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.GetTableDataProfileRequest{
								Name: name,
							}

							resp, err := client.GetTableDataProfile(ctx, req)
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
						Usage: "delete table-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "table-data-profile", Usage: "The ID of the table data profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/tableDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("table-data-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTableDataProfile on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dlp.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dlppb.DeleteTableDataProfileRequest{
								Name: name,
							}

							if err := client.DeleteTableDataProfile(ctx, req); err != nil {
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
