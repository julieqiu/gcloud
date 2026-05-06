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

package orgpolicy

import (
	orgpolicy "cloud.google.com/go/orgpolicy/apiv2"
	"cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the orgpolicy command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "orgpolicy",
		Usage: "manage Organization Policy API resources",
		Commands: []*cli.Command{
			{
				Name:  "constraints",
				Usage: "Manage constraints resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list constraints",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Size of the pages to be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token used to retrieve the next page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.ListConstraintsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListConstraints(ctx, req)
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
				Name:  "custom-constraints",
				Usage: "Manage custom-constraints resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create custom-constraints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.CreateCustomConstraintRequest{
								Parent: parent,
							}

							resp, err := client.CreateCustomConstraint(ctx, req)
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
						Usage: "update custom-constraints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-constraint", Usage: "The ID of the custom constraint.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "custom_constraint.name" not yet supported.
							custom_constraint_name := fmt.Sprintf("organizations/%s/customConstraints/%s", cmd.String("organization"), cmd.String("custom-constraint"))
							fmt.Printf("Executing update on %s\n", custom_constraint_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe custom-constraints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-constraint", Usage: "The ID of the custom constraint.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/customConstraints/%s", cmd.String("organization"), cmd.String("custom-constraint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.GetCustomConstraintRequest{
								Name: name,
							}

							resp, err := client.GetCustomConstraint(ctx, req)
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
						Usage: "list custom-constraints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Size of the pages to be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token used to retrieve the next page.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.ListCustomConstraintsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCustomConstraints(ctx, req)
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
						Usage: "delete custom-constraints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-constraint", Usage: "The ID of the custom constraint.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/customConstraints/%s", cmd.String("organization"), cmd.String("custom-constraint"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCustomConstraint on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.DeleteCustomConstraintRequest{
								Name: name,
							}

							if err := client.DeleteCustomConstraint(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "policies",
				Usage: "Manage policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list policies",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Size of the pages to be returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token used to retrieve the next page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.ListPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPolicies(ctx, req)
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
						Usage: "describe policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policie", Usage: "The ID of the policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/policies/%s", cmd.String("project"), cmd.String("policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.GetPolicyRequest{
								Name: name,
							}

							resp, err := client.GetPolicy(ctx, req)
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
						Usage: "describe policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policie", Usage: "The ID of the policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/policies/%s", cmd.String("project"), cmd.String("policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.GetEffectivePolicyRequest{
								Name: name,
							}

							resp, err := client.GetEffectivePolicy(ctx, req)
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
						Usage: "create policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.CreatePolicyRequest{
								Parent: parent,
							}

							resp, err := client.CreatePolicy(ctx, req)
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
						Usage: "update policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policie", Usage: "The ID of the policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "policy.name" not yet supported.
							policy_name := fmt.Sprintf("projects/%s/policies/%s", cmd.String("project"), cmd.String("policie"))
							fmt.Printf("Executing update on %s\n", policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The current etag of policy.", Required: false},
							&cli.StringFlag{Name: "policie", Usage: "The ID of the policie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/policies/%s", cmd.String("project"), cmd.String("policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePolicy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := orgpolicy.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &orgpolicypb.DeletePolicyRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeletePolicy(ctx, req); err != nil {
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
