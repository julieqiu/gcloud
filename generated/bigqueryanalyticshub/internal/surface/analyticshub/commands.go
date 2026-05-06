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

package analyticshub

import (
	analyticshub "cloud.google.com/go/analyticshub/apiv1"
	"cloud.google.com/go/analyticshub/apiv1/analyticshubpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the analyticshub command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "analyticshub",
		Usage: "manage Analytics Hub API resources",
		Commands: []*cli.Command{
			{
				Name:  "data-exchanges",
				Usage: "Manage data-exchanges resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.ListDataExchangesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataExchanges(ctx, req)
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
						Usage: "list data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							organization := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", organization)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.GetDataExchangeRequest{
								Name: name,
							}

							resp, err := client.GetDataExchange(ctx, req)
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
						Usage: "create data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange-id", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.CreateDataExchangeRequest{
								Parent:         parent,
								DataExchangeId: cmd.String("data-exchange-id"),
							}

							resp, err := client.CreateDataExchange(ctx, req)
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
						Usage: "update data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_exchange.name" not yet supported.
							data_exchange_name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							fmt.Printf("Executing update on %s\n", data_exchange_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDataExchange on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.DeleteDataExchangeRequest{
								Name: name,
							}

							if err := client.DeleteDataExchange(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "subscribe",
						Usage: "subscribe data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "destination", Usage: "The parent resource path of the Subscription.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscriber-contact", Usage: "Email of the subscriber.", Required: false},
							&cli.StringFlag{Name: "subscription", Usage: "Name of the subscription to create.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.SubscribeDataExchangeRequest{
								Name:              name,
								Destination:       cmd.String("destination"),
								Subscription:      cmd.String("subscription"),
								SubscriberContact: cmd.String("subscriber-contact"),
							}

							op, err := client.SubscribeDataExchange(ctx, req)
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
						Name:  "list",
						Usage: "list data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.BoolFlag{Name: "include-deleted-subscriptions", Usage: "If selected, includes deleted subscriptions in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.ListSharedResourceSubscriptionsRequest{
								Resource:                    resource,
								IncludeDeletedSubscriptions: cmd.Bool("include-deleted-subscriptions"),
								PageSize:                    int32(cmd.Int("page-size")),
								PageToken:                   cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSharedResourceSubscriptions(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.GetIamPolicyRequest{
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.SetIamPolicyRequest{
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
						Usage: "test-iam-permissions data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.TestIamPermissionsRequest{
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
				},
			},
			{
				Name:  "listings",
				Usage: "Manage listings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.ListListingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListListings(ctx, req)
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
						Usage: "describe listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The ID of the listing.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("listing"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.GetListingRequest{
								Name: name,
							}

							resp, err := client.GetListing(ctx, req)
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
						Usage: "create listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "listing-id", Usage: "The ID of the listing to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.CreateListingRequest{
								Parent:    parent,
								ListingId: cmd.String("listing-id"),
							}

							resp, err := client.CreateListing(ctx, req)
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
						Usage: "update listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The ID of the listing.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "listing.name" not yet supported.
							listing_name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("listing"))
							fmt.Printf("Executing update on %s\n", listing_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.BoolFlag{Name: "delete-commercial", Usage: "If the listing is commercial then this field must be set to true,.", Required: false},
							&cli.StringFlag{Name: "listing", Usage: "The ID of the listing.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("listing"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteListing on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.DeleteListingRequest{
								Name:             name,
								DeleteCommercial: cmd.Bool("delete-commercial"),
							}

							if err := client.DeleteListing(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "subscribe",
						Usage: "subscribe listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The ID of the listing.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("listing"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.SubscribeListingRequest{
								Name: name,
							}

							resp, err := client.SubscribeListing(ctx, req)
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
				Name:  "query-templates",
				Usage: "Manage query-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-template-id", Usage: "The ID of the QueryTemplate to create.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.CreateQueryTemplateRequest{
								Parent:          parent,
								QueryTemplateId: cmd.String("query-template-id"),
							}

							resp, err := client.CreateQueryTemplate(ctx, req)
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
						Usage: "describe query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-template", Usage: "The ID of the query template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("query-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.GetQueryTemplateRequest{
								Name: name,
							}

							resp, err := client.GetQueryTemplate(ctx, req)
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
						Usage: "list query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.ListQueryTemplatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListQueryTemplates(ctx, req)
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
						Usage: "update query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-template", Usage: "The ID of the query template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "query_template.name" not yet supported.
							query_template_name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("query-template"))
							fmt.Printf("Executing update on %s\n", query_template_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-template", Usage: "The ID of the query template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("query-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteQueryTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.DeleteQueryTemplateRequest{
								Name: name,
							}

							if err := client.DeleteQueryTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "submit",
						Usage: "submit query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-template", Usage: "The ID of the query template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("query-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.SubmitQueryTemplateRequest{
								Name: name,
							}

							resp, err := client.SubmitQueryTemplate(ctx, req)
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
						Name:  "approve",
						Usage: "approve query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-exchange", Usage: "The ID of the data exchange.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-template", Usage: "The ID of the query template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-exchange"), cmd.String("query-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.ApproveQueryTemplateRequest{
								Name: name,
							}

							resp, err := client.ApproveQueryTemplate(ctx, req)
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
				Name:  "subscriptions",
				Usage: "Manage subscriptions resources",
				Commands: []*cli.Command{

					{
						Name:  "refresh",
						Usage: "refresh subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.RefreshSubscriptionRequest{
								Name: name,
							}

							op, err := client.RefreshSubscription(ctx, req)
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
						Name:  "describe",
						Usage: "describe subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.GetSubscriptionRequest{
								Name: name,
							}

							resp, err := client.GetSubscription(ctx, req)
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
						Usage: "list subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.ListSubscriptionsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSubscriptions(ctx, req)
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
						Name:  "revoke",
						Usage: "revoke subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "revoke-commercial", Usage: "If the subscription is commercial then this field must be set to.", Required: false},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.RevokeSubscriptionRequest{
								Name:             name,
								RevokeCommercial: cmd.Bool("revoke-commercial"),
							}

							resp, err := client.RevokeSubscription(ctx, req)
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
						Usage: "delete subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSubscription %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := analyticshub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &analyticshubpb.DeleteSubscriptionRequest{
								Name: name,
							}

							op, err := client.DeleteSubscription(ctx, req)
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
				},
			},
		},
	}
}
