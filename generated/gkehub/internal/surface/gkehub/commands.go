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

package gkehub

import (
	gkehub "cloud.google.com/go/gkehub/apiv1"
	"cloud.google.com/go/gkehub/apiv1/gkehubpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the gkehub command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkehub",
		Usage: "manage GKE Hub resources",
		Commands: []*cli.Command{
			{
				Name:  "bindings",
				Usage: "Manage bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "binding", Usage: "The ID of the binding.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetMembershipBindingRequest{
								Name: name,
							}

							resp, err := client.GetMembershipBinding(ctx, req)
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
						Usage: "create bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "membership-binding-id", Usage: "The ID to use for the MembershipBinding.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateMembershipBindingRequest{
								Parent:              parent,
								MembershipBindingId: cmd.String("membership-binding-id"),
							}

							op, err := client.CreateMembershipBinding(ctx, req)
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
						Name:  "update",
						Usage: "update bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "binding", Usage: "The ID of the binding.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "membership_binding.name" not yet supported.
							membership_binding_name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("binding"))
							fmt.Printf("Executing update on %s\n", membership_binding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "binding", Usage: "The ID of the binding.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("binding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMembershipBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteMembershipBindingRequest{
								Name: name,
							}

							op, err := client.DeleteMembershipBinding(ctx, req)
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
						Usage: "list bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists MembershipBindings that match the filter expression,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListMembershipBindings` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListMembershipBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMembershipBindings(ctx, req)
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
				Name:  "features",
				Usage: "Manage features resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists Features that match the filter expression, following the syntax.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "One or more fields to compare and use to sort the output.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListFeatures` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If set to true, the response will return partial results when.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListFeaturesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								Filter:               cmd.String("filter"),
								OrderBy:              cmd.String("order-by"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListFeatures(ctx, req)
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
						Usage: "describe features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If set to true, the response will return partial results when.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetFeatureRequest{
								Name:                 name,
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							resp, err := client.GetFeature(ctx, req)
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
						Usage: "create features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature-id", Usage: "The ID of the feature to create.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateFeatureRequest{
								Parent:    parent,
								FeatureId: cmd.String("feature-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateFeature(ctx, req)
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
						Usage: "delete features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, the delete will ignore any outstanding resources for.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFeature %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteFeatureRequest{
								Name:      name,
								Force:     cmd.Bool("force"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteFeature(ctx, req)
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
						Name:  "update",
						Usage: "update features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "feature", Usage: "The ID of the feature.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.UpdateFeatureRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.UpdateFeature(ctx, req)
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
				Name:  "fleets",
				Usage: "Manage fleets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create fleets",
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
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateFleetRequest{
								Parent: parent,
							}

							op, err := client.CreateFleet(ctx, req)
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
						Usage: "describe fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet", Usage: "The ID of the fleet.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetFleetRequest{
								Name: name,
							}

							resp, err := client.GetFleet(ctx, req)
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
						Usage: "update fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet", Usage: "The ID of the fleet.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "fleet.name" not yet supported.
							fleet_name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							fmt.Printf("Executing update on %s\n", fleet_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet", Usage: "The ID of the fleet.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFleet %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteFleetRequest{
								Name: name,
							}

							op, err := client.DeleteFleet(ctx, req)
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
						Usage: "list fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of fleets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListFleets` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListFleetsRequest{
								Parent:    parent,
								PageToken: cmd.String("page-token"),
								PageSize:  int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.ListFleets(ctx, req)
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
				Name:  "memberships",
				Usage: "Manage memberships resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists Memberships that match the filter expression, following the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "One or more fields to compare and use to sort the output.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListMemberships` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListMembershipsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMemberships(ctx, req)
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
						Usage: "describe memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetMembershipRequest{
								Name: name,
							}

							resp, err := client.GetMembership(ctx, req)
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
						Usage: "create memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership-id", Usage: "Client chosen ID for the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateMembershipRequest{
								Parent:       parent,
								MembershipId: cmd.String("membership-id"),
								RequestId:    cmd.String("request-id"),
							}

							op, err := client.CreateMembership(ctx, req)
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
						Usage: "delete memberships",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any subresource from this Membership will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMembership %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteMembershipRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteMembership(ctx, req)
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
						Name:  "update",
						Usage: "update memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.UpdateMembershipRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.UpdateMembership(ctx, req)
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
						Name:  "generate-connect-manifest",
						Usage: "generate-connect-manifest memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image-pull-secret-content", Usage: "The image pull secret content for the registry, if not public.", Required: false},
							&cli.BoolFlag{Name: "is-upgrade", Usage: "If true, generate the resources for upgrade only.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "Namespace for GKE Connect agent resources.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "proxy", Usage: "URI of a proxy if connectivity from the agent to.", Required: false},
							&cli.StringFlag{Name: "registry", Usage: "The registry to fetch the connect agent image from.", Required: false},
							&cli.StringFlag{Name: "version", Usage: "The Connect agent version to use.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GenerateConnectManifestRequest{
								Name:                   name,
								Namespace:              cmd.String("namespace"),
								Proxy:                  []byte(cmd.String("proxy")),
								Version:                cmd.String("version"),
								IsUpgrade:              cmd.Bool("is-upgrade"),
								Registry:               cmd.String("registry"),
								ImagePullSecretContent: []byte(cmd.String("image-pull-secret-content")),
							}

							resp, err := client.GenerateConnectManifest(ctx, req)
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
				Name:  "namespaces",
				Usage: "Manage namespaces resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetScopeNamespaceRequest{
								Name: name,
							}

							resp, err := client.GetScopeNamespace(ctx, req)
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
						Usage: "create namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
							&cli.StringFlag{Name: "scope-namespace-id", Usage: "Client chosen ID for the Namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateScopeNamespaceRequest{
								Parent:           parent,
								ScopeNamespaceId: cmd.String("scope-namespace-id"),
							}

							op, err := client.CreateScopeNamespace(ctx, req)
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
						Name:  "update",
						Usage: "update namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "scope_namespace.name" not yet supported.
							scope_namespace_name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							fmt.Printf("Executing update on %s\n", scope_namespace_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The ID of the namespace.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteScopeNamespace %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteScopeNamespaceRequest{
								Name: name,
							}

							op, err := client.DeleteScopeNamespace(ctx, req)
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
						Usage: "list namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListFeatures` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListScopeNamespacesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListScopeNamespaces(ctx, req)
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "rbacrolebindings",
				Usage: "Manage rbacrolebindings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The ID of the rbacrolebinding.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetScopeRBACRoleBindingRequest{
								Name: name,
							}

							resp, err := client.GetScopeRBACRoleBinding(ctx, req)
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
						Usage: "create rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding-id", Usage: "Client chosen ID for the RBACRoleBinding.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateScopeRBACRoleBindingRequest{
								Parent:            parent,
								RbacrolebindingId: cmd.String("rbacrolebinding-id"),
							}

							op, err := client.CreateScopeRBACRoleBinding(ctx, req)
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
						Name:  "update",
						Usage: "update rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The ID of the rbacrolebinding.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "rbacrolebinding.name" not yet supported.
							rbacrolebinding_name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							fmt.Printf("Executing update on %s\n", rbacrolebinding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The ID of the rbacrolebinding.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteScopeRBACRoleBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteScopeRBACRoleBindingRequest{
								Name: name,
							}

							op, err := client.DeleteScopeRBACRoleBinding(ctx, req)
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
						Usage: "list rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListScopeRBACRoleBindings`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListScopeRBACRoleBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListScopeRBACRoleBindings(ctx, req)
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
						Usage: "describe rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The ID of the rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("rbacrolebinding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetMembershipRBACRoleBindingRequest{
								Name: name,
							}

							resp, err := client.GetMembershipRBACRoleBinding(ctx, req)
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
						Usage: "create rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding-id", Usage: "Client chosen ID for the RBACRoleBinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateMembershipRBACRoleBindingRequest{
								Parent:            parent,
								RbacrolebindingId: cmd.String("rbacrolebinding-id"),
							}

							op, err := client.CreateMembershipRBACRoleBinding(ctx, req)
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
						Name:  "update",
						Usage: "update rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The ID of the rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "rbacrolebinding.name" not yet supported.
							rbacrolebinding_name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("rbacrolebinding"))
							fmt.Printf("Executing update on %s\n", rbacrolebinding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The ID of the rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("rbacrolebinding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMembershipRBACRoleBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteMembershipRBACRoleBindingRequest{
								Name: name,
							}

							op, err := client.DeleteMembershipRBACRoleBinding(ctx, req)
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
						Usage: "list rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListMembershipRBACRoleBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListMembershipRBACRoleBindings(ctx, req)
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
						Name:  "generate-membership-rbac-role-binding-yaml",
						Usage: "generate-membership-rbac-role-binding-yaml rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding-id", Usage: "Client chosen ID for the RBACRoleBinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GenerateMembershipRBACRoleBindingYAMLRequest{
								Parent:            parent,
								RbacrolebindingId: cmd.String("rbacrolebinding-id"),
							}

							resp, err := client.GenerateMembershipRBACRoleBindingYAML(ctx, req)
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
				Name:  "scopes",
				Usage: "Manage scopes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Lists Memberships that match the filter expression, following the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListBoundMemberships` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
							&cli.StringFlag{Name: "scope-name", Usage: "Name of the Scope, in the format.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							scope_name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing list on %s\n", scope_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.GetScopeRequest{
								Name: name,
							}

							resp, err := client.GetScope(ctx, req)
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
						Usage: "create scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope-id", Usage: "Client chosen ID for the Scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.CreateScopeRequest{
								Parent:  parent,
								ScopeId: cmd.String("scope-id"),
							}

							op, err := client.CreateScope(ctx, req)
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
						Name:  "update",
						Usage: "update scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "scope.name" not yet supported.
							scope_name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing update on %s\n", scope_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The ID of the scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteScope %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.DeleteScopeRequest{
								Name: name,
							}

							op, err := client.DeleteScope(ctx, req)
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
						Usage: "list scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListScopes` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListScopesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListScopes(ctx, req)
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
						Usage: "list scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "When requesting a 'page' of resources, `page_size` specifies.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Token returned by previous call to `ListPermittedScopes` which.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkehub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkehubpb.ListPermittedScopesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPermittedScopes(ctx, req)
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
