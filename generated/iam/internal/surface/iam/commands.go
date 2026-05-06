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

package iam

import (
	iam "cloud.google.com/go/iam/apiv3"
	"cloud.google.com/go/iam/apiv3/iampb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the iam command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "iam",
		Usage: "manage Identity and Access Management (IAM) API resources",
		Commands: []*cli.Command{
			{
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

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
				},
			},
			{
				Name:  "policy-bindings",
				Usage: "Manage policy-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-binding-id", Usage: "The ID to use for the policy binding, which will become the final.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the creation, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.CreatePolicyBindingRequest{
								Parent:          parent,
								PolicyBindingId: cmd.String("policy-binding-id"),
								ValidateOnly:    cmd.Bool("validate-only"),
							}

							op, err := client.CreatePolicyBinding(ctx, req)
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
						Usage: "describe policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-binding", Usage: "The ID of the policy binding.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/policyBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("policy-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.GetPolicyBindingRequest{
								Name: name,
							}

							resp, err := client.GetPolicyBinding(ctx, req)
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
						Usage: "update policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-binding", Usage: "The ID of the policy binding.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the update, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "policy_binding.name" not yet supported.
							policy_binding_name := fmt.Sprintf("projects/%s/locations/%s/policyBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("policy-binding"))
							fmt.Printf("Executing update on %s\n", policy_binding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag of the policy binding.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-binding", Usage: "The ID of the policy binding.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the deletion, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/policyBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("policy-binding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePolicyBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.DeletePolicyBindingRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeletePolicyBinding(ctx, req)
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
						Usage: "list policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of policy bindings to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListPolicyBindings` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.ListPolicyBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPolicyBindings(ctx, req)
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
						Name:  "search-target-policy-bindings",
						Usage: "search-target-policy-bindings policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of policy bindings to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target", Usage: "The target resource, which is bound to the policy in the binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.SearchTargetPolicyBindingsRequest{
								Parent:    parent,
								Target:    cmd.String("target"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchTargetPolicyBindings(ctx, req)
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
				Name:  "principal-access-boundary-policies",
				Usage: "Manage principal-access-boundary-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "principal-access-boundary-policy-id", Usage: "The ID to use for the principal access boundary policy, which.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the creation, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.CreatePrincipalAccessBoundaryPolicyRequest{
								Parent:                          parent,
								PrincipalAccessBoundaryPolicyId: cmd.String("principal-access-boundary-policy-id"),
								ValidateOnly:                    cmd.Bool("validate-only"),
							}

							op, err := client.CreatePrincipalAccessBoundaryPolicy(ctx, req)
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
						Usage: "describe principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "principal-access-boundary-policie", Usage: "The ID of the principal access boundary policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal-access-boundary-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.GetPrincipalAccessBoundaryPolicyRequest{
								Name: name,
							}

							resp, err := client.GetPrincipalAccessBoundaryPolicy(ctx, req)
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
						Usage: "update principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "principal-access-boundary-policie", Usage: "The ID of the principal access boundary policie.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the update, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "principal_access_boundary_policy.name" not yet supported.
							principal_access_boundary_policy_name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal-access-boundary-policie"))
							fmt.Printf("Executing update on %s\n", principal_access_boundary_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag of the principal access boundary policy.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, the request will force the deletion of the policy.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "principal-access-boundary-policie", Usage: "The ID of the principal access boundary policie.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the deletion, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal-access-boundary-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePrincipalAccessBoundaryPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.DeletePrincipalAccessBoundaryPolicyRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								ValidateOnly: cmd.Bool("validate-only"),
								Force:        cmd.Bool("force"),
							}

							op, err := client.DeletePrincipalAccessBoundaryPolicy(ctx, req)
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
						Usage: "list principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of principal access boundary policies to.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.ListPrincipalAccessBoundaryPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPrincipalAccessBoundaryPolicies(ctx, req)
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
						Name:  "search-policy-bindings",
						Usage: "search-policy-bindings principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of policy bindings to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "principal-access-boundary-policie", Usage: "The ID of the principal access boundary policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal-access-boundary-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := iam.NewIamPolicyClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &iampb.SearchPrincipalAccessBoundaryPolicyBindingsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchPrincipalAccessBoundaryPolicyBindings(ctx, req)
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
