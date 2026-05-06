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

package accesscontextmanager

import (
	accesscontextmanager "cloud.google.com/go/accesscontextmanager/apiv1"
	"cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the accesscontextmanager command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "accesscontextmanager",
		Usage: "manage Access Context Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "access-levels",
				Usage: "Manage access-levels resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-level-format", Usage: "Whether to return `BasicLevels` in the Cloud Common Expression language, as.", Required: false},
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of [Access Levels].", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Next page token for the next batch of [Access Level].", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.ListAccessLevelsRequest{
								Parent:            parent,
								PageSize:          int32(cmd.Int("page-size")),
								PageToken:         cmd.String("page-token"),
								AccessLevelFormat: accesscontextmanagerpb.LevelFormat(accesscontextmanagerpb.LevelFormat_value[cmd.String("access-level-format")]),
							}

							limit := cmd.Int("limit")
							it := client.ListAccessLevels(ctx, req)
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
						Usage: "describe access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-level", Usage: "The ID of the access level.", Required: true},
							&cli.StringFlag{Name: "access-level-format", Usage: "Whether to return `BasicLevels` in the Cloud Common Expression.", Required: false},
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/accessLevels/%s", cmd.String("access-policie"), cmd.String("access-level"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.GetAccessLevelRequest{
								Name:              name,
								AccessLevelFormat: accesscontextmanagerpb.LevelFormat(accesscontextmanagerpb.LevelFormat_value[cmd.String("access-level-format")]),
							}

							resp, err := client.GetAccessLevel(ctx, req)
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
						Usage: "create access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.CreateAccessLevelRequest{
								Parent: parent,
							}

							op, err := client.CreateAccessLevel(ctx, req)
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
						Usage: "update access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-level", Usage: "The ID of the access level.", Required: true},
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "access_level.name" not yet supported.
							access_level_name := fmt.Sprintf("accessPolicies/%s/accessLevels/%s", cmd.String("access-policie"), cmd.String("access-level"))
							fmt.Printf("Executing update on %s\n", access_level_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-level", Usage: "The ID of the access level.", Required: true},
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/accessLevels/%s", cmd.String("access-policie"), cmd.String("access-level"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAccessLevel %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.DeleteAccessLevelRequest{
								Name: name,
							}

							op, err := client.DeleteAccessLevel(ctx, req)
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
						Name:  "replace-all",
						Usage: "replace-all access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag for the version of the [Access Policy].", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.ReplaceAccessLevelsRequest{
								Parent: parent,
								Etag:   cmd.String("etag"),
							}

							op, err := client.ReplaceAccessLevels(ctx, req)
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
				Name:  "access-policies",
				Usage: "Manage access-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list access-policies",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Number of AccessPolicy instances to include in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Next page token for the next batch of AccessPolicy instances.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.ListAccessPoliciesRequest{
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAccessPolicies(ctx, req)
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
						Usage: "describe access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.GetAccessPolicyRequest{
								Name: name,
							}

							resp, err := client.GetAccessPolicy(ctx, req)
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
						Usage: "create access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "An opaque identifier for the current version of the.", Required: false},
							&cli.StringSliceFlag{Name: "scopes", Usage: "The scopes of a policy define which resources an ACM policy can restrict,.", Required: false},
							&cli.StringFlag{Name: "title", Usage: "Human readable title.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.AccessPolicy{
								Title:  cmd.String("title"),
								Scopes: cmd.StringSlice("scopes"),
								Etag:   cmd.String("etag"),
							}

							op, err := client.CreateAccessPolicy(ctx, req)
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
						Usage: "update access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "policy.name" not yet supported.
							policy_name := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							fmt.Printf("Executing update on %s\n", policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAccessPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.DeleteAccessPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteAccessPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.SetIamPolicyRequest{
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.GetIamPolicyRequest{
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.TestIamPermissionsRequest{
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
				Name:  "gcp-user-access-bindings",
				Usage: "Manage gcp-user-access-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If left blank, returns the first page.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.ListGcpUserAccessBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGcpUserAccessBindings(ctx, req)
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
						Usage: "describe gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gcp-user-access-binding", Usage: "The ID of the gcp user access binding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/%s", cmd.String("organization"), cmd.String("gcp-user-access-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.GetGcpUserAccessBindingRequest{
								Name: name,
							}

							resp, err := client.GetGcpUserAccessBinding(ctx, req)
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
						Usage: "create gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.CreateGcpUserAccessBindingRequest{
								Parent: parent,
							}

							op, err := client.CreateGcpUserAccessBinding(ctx, req)
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
						Usage: "update gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gcp-user-access-binding", Usage: "The ID of the gcp user access binding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "gcp_user_access_binding.name" not yet supported.
							gcp_user_access_binding_name := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/%s", cmd.String("organization"), cmd.String("gcp-user-access-binding"))
							fmt.Printf("Executing update on %s\n", gcp_user_access_binding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gcp-user-access-binding", Usage: "The ID of the gcp user access binding.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/%s", cmd.String("organization"), cmd.String("gcp-user-access-binding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGcpUserAccessBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.DeleteGcpUserAccessBindingRequest{
								Name: name,
							}

							op, err := client.DeleteGcpUserAccessBinding(ctx, req)
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
				Name:  "service-perimeters",
				Usage: "Manage service-perimeters resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of [Service Perimeters].", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Next page token for the next batch of [Service Perimeter].", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.ListServicePerimetersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListServicePerimeters(ctx, req)
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
						Usage: "describe service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringFlag{Name: "service-perimeter", Usage: "The ID of the service perimeter.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", cmd.String("access-policie"), cmd.String("service-perimeter"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.GetServicePerimeterRequest{
								Name: name,
							}

							resp, err := client.GetServicePerimeter(ctx, req)
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
						Usage: "create service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.CreateServicePerimeterRequest{
								Parent: parent,
							}

							op, err := client.CreateServicePerimeter(ctx, req)
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
						Usage: "update service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringFlag{Name: "service-perimeter", Usage: "The ID of the service perimeter.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service_perimeter.name" not yet supported.
							service_perimeter_name := fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", cmd.String("access-policie"), cmd.String("service-perimeter"))
							fmt.Printf("Executing update on %s\n", service_perimeter_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringFlag{Name: "service-perimeter", Usage: "The ID of the service perimeter.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", cmd.String("access-policie"), cmd.String("service-perimeter"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServicePerimeter %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.DeleteServicePerimeterRequest{
								Name: name,
							}

							op, err := client.DeleteServicePerimeter(ctx, req)
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
						Name:  "replace-all",
						Usage: "replace-all service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag for the version of the [Access Policy].", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.ReplaceServicePerimetersRequest{
								Parent: parent,
								Etag:   cmd.String("etag"),
							}

							op, err := client.ReplaceServicePerimeters(ctx, req)
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
						Name:  "commit",
						Usage: "commit service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-policie", Usage: "The ID of the access policie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag for the version of the [Access Policy].", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accesscontextmanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accesscontextmanagerpb.CommitServicePerimetersRequest{
								Parent: parent,
								Etag:   cmd.String("etag"),
							}

							op, err := client.CommitServicePerimeters(ctx, req)
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
