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

package datapolicies

import (
	datapolicies "cloud.google.com/go/datapolicies/apiv2"
	"cloud.google.com/go/datapolicies/apiv2/datapoliciespb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the bigquerydatapolicy command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigquerydatapolicy",
		Usage: "manage BigQuery Data Policy API resources",
		Commands: []*cli.Command{
			{
				Name:  "data-policies",
				Usage: "Manage data-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policy-id", Usage: "User-assigned (human readable) ID of the data policy that needs.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.CreateDataPolicyRequest{
								Parent:       parent,
								DataPolicyId: cmd.String("data-policy-id"),
							}

							resp, err := client.CreateDataPolicy(ctx, req)
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
						Name:  "add-grantees",
						Usage: "add-grantees data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "data-policy", Usage: "Resource name of this data policy, in the format of.", Required: true},
							&cli.StringSliceFlag{Name: "grantees", Usage: "IAM principal that should be granted Fine Grained Access to the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							data_policy := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							fmt.Printf("Executing add-grantees on %s\n", data_policy)
							return nil
						},
					},

					{
						Name:  "remove-grantees",
						Usage: "remove-grantees data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "data-policy", Usage: "Resource name of this data policy, in the format of.", Required: true},
							&cli.StringSliceFlag{Name: "grantees", Usage: "IAM principal that should be revoked from Fine Grained Access to.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							data_policy := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							fmt.Printf("Executing remove-grantees on %s\n", data_policy)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update data-policies",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the data policy is not found, a new data.", Required: false},
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_policy.name" not yet supported.
							data_policy_name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							fmt.Printf("Executing update on %s\n", data_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDataPolicy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.DeleteDataPolicyRequest{
								Name: name,
							}

							if err := client.DeleteDataPolicy(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.GetDataPolicyRequest{
								Name: name,
							}

							resp, err := client.GetDataPolicy(ctx, req)
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
						Usage: "list data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filters the data policies by policy tags that they.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of data policies to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `nextPageToken` value returned from a previous list request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.ListDataPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataPolicies(ctx, req)
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
						Usage: "get-iam-policy data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.GetIamPolicyRequest{
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
						Usage: "set-iam-policy data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.SetIamPolicyRequest{
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
						Usage: "test-iam-permissions data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-policie", Usage: "The ID of the data policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datapolicies.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datapoliciespb.TestIamPermissionsRequest{
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
		},
	}
}
