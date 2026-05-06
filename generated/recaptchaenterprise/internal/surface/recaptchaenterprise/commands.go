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

package recaptchaenterprise

import (
	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1"
	"cloud.google.com/go/recaptchaenterprise/apiv1/recaptchaenterprisepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the recaptchaenterprise command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "recaptchaenterprise",
		Usage: "manage reCAPTCHA Enterprise API resources",
		Commands: []*cli.Command{
			{
				Name:  "assessments",
				Usage: "Manage assessments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create assessments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.CreateAssessmentRequest{
								Parent: parent,
							}

							resp, err := client.CreateAssessment(ctx, req)
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
						Name:  "annotate",
						Usage: "annotate assessments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-id", Usage: "A stable account identifier to apply to the assessment.", Required: false},
							&cli.StringFlag{Name: "annotation", Usage: "The annotation that is assigned to the Event.", Required: false},
							&cli.StringFlag{Name: "assessment", Usage: "The ID of the assessment.", Required: true},
							&cli.StringFlag{Name: "hashed-account-id", Usage: "A stable hashed account identifier to apply to the assessment.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "reasons", Usage: "Reasons for the annotation that are assigned to the event.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/assessments/%s", cmd.String("project"), cmd.String("assessment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.AnnotateAssessmentRequest{
								Name:            name,
								Annotation:      recaptchaenterprisepb.AnnotateAssessmentRequest_Annotation(recaptchaenterprisepb.AnnotateAssessmentRequest_Annotation_value[cmd.String("annotation")]),
								Reasons:         cmd.StringSlice("reasons"),
								AccountId:       cmd.String("account-id"),
								HashedAccountId: []byte(cmd.String("hashed-account-id")),
							}

							resp, err := client.AnnotateAssessment(ctx, req)
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
				Name:  "firewallpolicies",
				Usage: "Manage firewallpolicies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.CreateFirewallPolicyRequest{
								Parent: parent,
							}

							resp, err := client.CreateFirewallPolicy(ctx, req)
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
						Usage: "list firewallpolicies",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of policies to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.ListFirewallPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListFirewallPolicies(ctx, req)
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
						Usage: "describe firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewallpolicie", Usage: "The ID of the firewallpolicie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.GetFirewallPolicyRequest{
								Name: name,
							}

							resp, err := client.GetFirewallPolicy(ctx, req)
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
						Usage: "update firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewallpolicie", Usage: "The ID of the firewallpolicie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "firewall_policy.name" not yet supported.
							firewall_policy_name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicie"))
							fmt.Printf("Executing update on %s\n", firewall_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewallpolicie", Usage: "The ID of the firewallpolicie.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteFirewallPolicy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.DeleteFirewallPolicyRequest{
								Name: name,
							}

							if err := client.DeleteFirewallPolicy(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "reorder",
						Usage: "reorder firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "names", Usage: "A list containing all policy names, in the new order.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.ReorderFirewallPoliciesRequest{
								Parent: parent,
								Names:  cmd.StringSlice("names"),
							}

							resp, err := client.ReorderFirewallPolicies(ctx, req)
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
				Name:  "keys",
				Usage: "Manage keys resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.CreateKeyRequest{
								Parent: parent,
							}

							resp, err := client.CreateKey(ctx, req)
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
						Usage: "list keys",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of keys to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.ListKeysRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListKeys(ctx, req)
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
						Name:  "retrieve-legacy-secret-key",
						Usage: "retrieve-legacy-secret-key keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							key := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing retrieve-legacy-secret-key on %s\n", key)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.GetKeyRequest{
								Name: name,
							}

							resp, err := client.GetKey(ctx, req)
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
						Usage: "update keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "key.name" not yet supported.
							key_name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing update on %s\n", key_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteKey on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.DeleteKeyRequest{
								Name: name,
							}

							if err := client.DeleteKey(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "migrate",
						Usage: "migrate keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-billing-check", Usage: "If true, skips the billing check.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.MigrateKeyRequest{
								Name:             name,
								SkipBillingCheck: cmd.Bool("skip-billing-check"),
							}

							resp, err := client.MigrateKey(ctx, req)
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
						Name:  "add-ip-override",
						Usage: "add-ip-override keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.AddIpOverrideRequest{
								Name: name,
							}

							resp, err := client.AddIpOverride(ctx, req)
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
						Name:  "remove-ip-override",
						Usage: "remove-ip-override keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.RemoveIpOverrideRequest{
								Name: name,
							}

							resp, err := client.RemoveIpOverride(ctx, req)
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
						Usage: "list keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of overrides to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.ListIpOverridesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListIpOverrides(ctx, req)
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
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of accounts to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "relatedaccountgroup", Usage: "The ID of the relatedaccountgroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/relatedaccountgroups/%s", cmd.String("project"), cmd.String("relatedaccountgroup"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.ListRelatedAccountGroupMembershipsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRelatedAccountGroupMemberships(ctx, req)
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
				Name:  "metrics",
				Usage: "Manage metrics resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe metrics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The ID of the key.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s/metrics", cmd.String("project"), cmd.String("key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.GetMetricsRequest{
								Name: name,
							}

							resp, err := client.GetMetrics(ctx, req)
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
				Name:  "relatedaccountgroupmemberships",
				Usage: "Manage relatedaccountgroupmemberships resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search relatedaccountgroupmemberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-id", Usage: "The unique stable account identifier used to search connections.", Required: false},
							&cli.StringFlag{Name: "hashed-account-id", Usage: "Deprecated: use `account_id` instead.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of groups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing search on %s\n", project)
							return nil
						},
					},
				},
			},
			{
				Name:  "relatedaccountgroups",
				Usage: "Manage relatedaccountgroups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list relatedaccountgroups",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of groups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListRelatedAccountGroups`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := recaptchaenterprise.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &recaptchaenterprisepb.ListRelatedAccountGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRelatedAccountGroups(ctx, req)
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
