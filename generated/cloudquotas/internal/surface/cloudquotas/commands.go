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

package cloudquotas

import (
	cloudquotas "cloud.google.com/go/cloudquotas/apiv1"
	"cloud.google.com/go/cloudquotas/apiv1/cloudquotaspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudquotas command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudquotas",
		Usage: "manage Cloud Quotas API resources",
		Commands: []*cli.Command{
			{
				Name:  "quota-infos",
				Usage: "Manage quota-infos resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list quota-infos",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudquotas.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudquotaspb.ListQuotaInfosRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListQuotaInfos(ctx, req)
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
						Usage: "describe quota-infos",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-info", Usage: "The ID of the quota info.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/services/%s/quotaInfos/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("quota-info"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudquotas.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudquotaspb.GetQuotaInfoRequest{
								Name: name,
							}

							resp, err := client.GetQuotaInfo(ctx, req)
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
				Name:  "quota-preferences",
				Usage: "Manage quota-preferences resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list quota-preferences",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter result QuotaPreferences by their state, type,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "How to order of the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudquotas.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudquotaspb.ListQuotaPreferencesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListQuotaPreferences(ctx, req)
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
						Usage: "describe quota-preferences",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-preference", Usage: "The ID of the quota preference.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/quotaPreferences/%s", cmd.String("project"), cmd.String("location"), cmd.String("quota-preference"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudquotas.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudquotaspb.GetQuotaPreferenceRequest{
								Name: name,
							}

							resp, err := client.GetQuotaPreference(ctx, req)
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
						Usage: "create quota-preferences",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "ignore-safety-checks", Usage: "The list of quota safety checks to be ignored.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-preference-id", Usage: "Id of the requesting object, must be unique under its parent.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := cloudquotas.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &cloudquotaspb.CreateQuotaPreferenceRequest{
								Parent:             parent,
								QuotaPreferenceId:  cmd.String("quota-preference-id"),
								IgnoreSafetyChecks: cmd.StringSlice("ignore-safety-checks"),
							}

							resp, err := client.CreateQuotaPreference(ctx, req)
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
						Usage: "update quota-preferences",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the quota preference is not found, a new one.", Required: false},
							&cli.StringSliceFlag{Name: "ignore-safety-checks", Usage: "The list of quota safety checks to be ignored.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "quota-preference", Usage: "The ID of the quota preference.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set to true, validate the request, but do not actually update.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "quota_preference.name" not yet supported.
							quota_preference_name := fmt.Sprintf("projects/%s/locations/%s/quotaPreferences/%s", cmd.String("project"), cmd.String("location"), cmd.String("quota-preference"))
							fmt.Printf("Executing update on %s\n", quota_preference_name)
							return nil
						},
					},
				},
			},
		},
	}
}
