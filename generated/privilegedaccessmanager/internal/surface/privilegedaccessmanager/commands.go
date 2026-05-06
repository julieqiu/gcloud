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

package privilegedaccessmanager

import (
	privilegedaccessmanager "cloud.google.com/go/privilegedaccessmanager/apiv1"
	"cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the privilegedaccessmanager command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "privilegedaccessmanager",
		Usage: "manage Privileged Access Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "entitlements",
				Usage: "Manage entitlements resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
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
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.ListEntitlementsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntitlements(ctx, req)
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
						Usage: "search entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "caller-access-type", Usage: "Only entitlements where the calling user has this access are.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only entitlements matching this filter are returned in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
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
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.SearchEntitlementsRequest{
								Parent:           parent,
								CallerAccessType: privilegedaccessmanagerpb.SearchEntitlementsRequest_CallerAccessType(privilegedaccessmanagerpb.SearchEntitlementsRequest_CallerAccessType_value[cmd.String("caller-access-type")]),
								Filter:           cmd.String("filter"),
								PageSize:         int32(cmd.Int("page-size")),
								PageToken:        cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchEntitlements(ctx, req)
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
						Usage: "describe entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.GetEntitlementRequest{
								Name: name,
							}

							resp, err := client.GetEntitlement(ctx, req)
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
						Usage: "create entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement-id", Usage: "The ID to use for this entitlement.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.CreateEntitlementRequest{
								Parent:        parent,
								EntitlementId: cmd.String("entitlement-id"),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.CreateEntitlement(ctx, req)
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
						Usage: "delete entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any child grant under this entitlement is also.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.DeleteEntitlementRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteEntitlement(ctx, req)
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
						Usage: "update entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entitlement.name" not yet supported.
							entitlement_name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
							fmt.Printf("Executing update on %s\n", entitlement_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "grants",
				Usage: "Manage grants resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.ListGrantsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGrants(ctx, req)
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
						Usage: "search grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "caller-relationship", Usage: "Only grants which the caller is related to by this relationship.", Required: true},
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only grants matching this filter are returned in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.SearchGrantsRequest{
								Parent:             parent,
								CallerRelationship: privilegedaccessmanagerpb.SearchGrantsRequest_CallerRelationshipType(privilegedaccessmanagerpb.SearchGrantsRequest_CallerRelationshipType_value[cmd.String("caller-relationship")]),
								Filter:             cmd.String("filter"),
								PageSize:           int32(cmd.Int("page-size")),
								PageToken:          cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchGrants(ctx, req)
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
						Usage: "describe grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "grant", Usage: "The ID of the grant.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.GetGrantRequest{
								Name: name,
							}

							resp, err := client.GetGrant(ctx, req)
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
						Usage: "create grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.CreateGrantRequest{
								Parent:    parent,
								RequestId: cmd.String("request-id"),
							}

							resp, err := client.CreateGrant(ctx, req)
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
						Usage: "approve grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "grant", Usage: "The ID of the grant.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "The reason for approving this grant.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.ApproveGrantRequest{
								Name:   name,
								Reason: cmd.String("reason"),
							}

							resp, err := client.ApproveGrant(ctx, req)
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
						Name:  "deny",
						Usage: "deny grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "grant", Usage: "The ID of the grant.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "The reason for denying this grant.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.DenyGrantRequest{
								Name:   name,
								Reason: cmd.String("reason"),
							}

							resp, err := client.DenyGrant(ctx, req)
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
						Name:  "revoke",
						Usage: "revoke grants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitlement", Usage: "The ID of the entitlement.", Required: true},
							&cli.StringFlag{Name: "grant", Usage: "The ID of the grant.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "The reason for revoking this grant.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entitlements/%s/grants/%s", cmd.String("project"), cmd.String("location"), cmd.String("entitlement"), cmd.String("grant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.RevokeGrantRequest{
								Name:   name,
								Reason: cmd.String("reason"),
							}

							op, err := client.RevokeGrant(ctx, req)
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
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "check-onboarding-status",
						Usage: "check-onboarding-status locations",
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
							client, err := privilegedaccessmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privilegedaccessmanagerpb.CheckOnboardingStatusRequest{
								Parent: parent,
							}

							resp, err := client.CheckOnboardingStatus(ctx, req)
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
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
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
				},
			},
		},
	}
}
