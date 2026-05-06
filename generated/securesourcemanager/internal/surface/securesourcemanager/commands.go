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

package securesourcemanager

import (
	securesourcemanager "cloud.google.com/go/securesourcemanager/apiv1"
	"cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the securesourcemanager command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "securesourcemanager",
		Usage: "manage Secure Source Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "branch-rules",
				Usage: "Manage branch-rules resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create branch-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch-rule-id", Usage: "The branch rule id.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreateBranchRuleRequest{
								Parent:       parent,
								BranchRuleId: cmd.String("branch-rule-id"),
							}

							op, err := client.CreateBranchRule(ctx, req)
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
						Usage: "list branch-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListBranchRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBranchRules(ctx, req)
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
						Usage: "describe branch-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch-rule", Usage: "The ID of the branch rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("branch-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetBranchRuleRequest{
								Name: name,
							}

							resp, err := client.GetBranchRule(ctx, req)
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
						Usage: "update branch-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branch-rule", Usage: "The ID of the branch rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the review, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "branch_rule.name" not yet supported.
							branch_rule_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("branch-rule"))
							fmt.Printf("Executing update on %s\n", branch_rule_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete branch-rules",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the branch rule is not found, the request.", Required: false},
							&cli.StringFlag{Name: "branch-rule", Usage: "The ID of the branch rule.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("branch-rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBranchRule %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeleteBranchRuleRequest{
								Name:         name,
								AllowMissing: cmd.Bool("allow-missing"),
							}

							op, err := client.DeleteBranchRule(ctx, req)
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
				Name:  "hooks",
				Usage: "Manage hooks resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list hooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListHooksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListHooks(ctx, req)
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
						Usage: "describe hooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hook", Usage: "The ID of the hook.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("hook"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetHookRequest{
								Name: name,
							}

							resp, err := client.GetHook(ctx, req)
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
						Usage: "create hooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hook-id", Usage: "The ID to use for the hook, which will become the final component.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreateHookRequest{
								Parent: parent,
								HookId: cmd.String("hook-id"),
							}

							op, err := client.CreateHook(ctx, req)
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
						Usage: "update hooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hook", Usage: "The ID of the hook.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "hook.name" not yet supported.
							hook_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("hook"))
							fmt.Printf("Executing update on %s\n", hook_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete hooks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "hook", Usage: "The ID of the hook.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("hook"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteHook %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeleteHookRequest{
								Name: name,
							}

							op, err := client.DeleteHook(ctx, req)
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
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter for filtering results.", Required: false},
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
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListInstancesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInstances(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetInstanceRequest{
								Name: name,
							}

							resp, err := client.GetInstance(ctx, req)
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
						Usage: "create instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance-id", Usage: "ID of the instance to be created.", Required: true},
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
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreateInstanceRequest{
								Parent:     parent,
								InstanceId: cmd.String("instance-id"),
								RequestId:  cmd.String("request-id"),
							}

							op, err := client.CreateInstance(ctx, req)
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
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, will force the deletion of the instance.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInstance %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeleteInstanceRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteInstance(ctx, req)
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
						Usage: "set-iam-policy instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "issue-comments",
				Usage: "Manage issue-comments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create issue-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreateIssueCommentRequest{
								Parent: parent,
							}

							op, err := client.CreateIssueComment(ctx, req)
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
						Usage: "describe issue-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "issue-comment", Usage: "The ID of the issue comment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s/issueComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"), cmd.String("issue-comment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetIssueCommentRequest{
								Name: name,
							}

							resp, err := client.GetIssueComment(ctx, req)
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
						Usage: "list issue-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListIssueCommentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListIssueComments(ctx, req)
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
						Usage: "update issue-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "issue-comment", Usage: "The ID of the issue comment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "issue_comment.name" not yet supported.
							issue_comment_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s/issueComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"), cmd.String("issue-comment"))
							fmt.Printf("Executing update on %s\n", issue_comment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete issue-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "issue-comment", Usage: "The ID of the issue comment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s/issueComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"), cmd.String("issue-comment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIssueComment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeleteIssueCommentRequest{
								Name: name,
							}

							op, err := client.DeleteIssueComment(ctx, req)
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
				Name:  "issues",
				Usage: "Manage issues resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreateIssueRequest{
								Parent: parent,
							}

							op, err := client.CreateIssue(ctx, req)
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
						Usage: "describe issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetIssueRequest{
								Name: name,
							}

							resp, err := client.GetIssue(ctx, req)
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
						Usage: "list issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Used to filter the resulting issues list.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListIssuesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListIssues(ctx, req)
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
						Usage: "update issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "issue.name" not yet supported.
							issue_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							fmt.Printf("Executing update on %s\n", issue_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the issue.", Required: false},
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteIssue %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeleteIssueRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteIssue(ctx, req)
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
						Name:  "open",
						Usage: "open issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the issue.", Required: false},
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.OpenIssueRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.OpenIssue(ctx, req)
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
						Name:  "close",
						Usage: "close issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the issue.", Required: false},
							&cli.StringFlag{Name: "issue", Usage: "The ID of the issue.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("issue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CloseIssueRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.CloseIssue(ctx, req)
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
				Name:  "pull-request-comments",
				Usage: "Manage pull-request-comments resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe pull-request-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "pull-request-comment", Usage: "The ID of the pull request comment.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s/pullRequestComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"), cmd.String("pull-request-comment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetPullRequestCommentRequest{
								Name: name,
							}

							resp, err := client.GetPullRequestComment(ctx, req)
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
						Usage: "list pull-request-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListPullRequestCommentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPullRequestComments(ctx, req)
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
						Name:  "create",
						Usage: "create pull-request-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreatePullRequestCommentRequest{
								Parent: parent,
							}

							op, err := client.CreatePullRequestComment(ctx, req)
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
						Usage: "update pull-request-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "pull-request-comment", Usage: "The ID of the pull request comment.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "pull_request_comment.name" not yet supported.
							pull_request_comment_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s/pullRequestComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"), cmd.String("pull-request-comment"))
							fmt.Printf("Executing update on %s\n", pull_request_comment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete pull-request-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "pull-request-comment", Usage: "The ID of the pull request comment.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s/pullRequestComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"), cmd.String("pull-request-comment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePullRequestComment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeletePullRequestCommentRequest{
								Name: name,
							}

							op, err := client.DeletePullRequestComment(ctx, req)
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
						Name:  "batch-create",
						Usage: "batch-create pull-request-comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.BatchCreatePullRequestCommentsRequest{
								Parent: parent,
							}

							op, err := client.BatchCreatePullRequestComments(ctx, req)
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
						Name:  "resolve",
						Usage: "resolve pull-request-comments",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-fill", Usage: "If set, at least one comment in a thread is required, rest of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the pull request comments to resolve.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ResolvePullRequestCommentsRequest{
								Parent:   parent,
								Names:    cmd.StringSlice("names"),
								AutoFill: cmd.Bool("auto-fill"),
							}

							op, err := client.ResolvePullRequestComments(ctx, req)
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
						Name:  "unresolve",
						Usage: "unresolve pull-request-comments",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-fill", Usage: "If set, at least one comment in a thread is required, rest of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the pull request comments to unresolve.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.UnresolvePullRequestCommentsRequest{
								Parent:   parent,
								Names:    cmd.StringSlice("names"),
								AutoFill: cmd.Bool("auto-fill"),
							}

							op, err := client.UnresolvePullRequestComments(ctx, req)
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
				Name:  "pull-requests",
				Usage: "Manage pull-requests resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreatePullRequestRequest{
								Parent: parent,
							}

							op, err := client.CreatePullRequest(ctx, req)
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
						Usage: "describe pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetPullRequestRequest{
								Name: name,
							}

							resp, err := client.GetPullRequest(ctx, req)
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
						Usage: "list pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListPullRequestsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPullRequests(ctx, req)
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
						Usage: "update pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "pull_request.name" not yet supported.
							pull_request_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							fmt.Printf("Executing update on %s\n", pull_request_name)
							return nil
						},
					},

					{
						Name:  "merge",
						Usage: "merge pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.MergePullRequestRequest{
								Name: name,
							}

							op, err := client.MergePullRequest(ctx, req)
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
						Name:  "open",
						Usage: "open pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.OpenPullRequestRequest{
								Name: name,
							}

							op, err := client.OpenPullRequest(ctx, req)
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
						Name:  "close",
						Usage: "close pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ClosePullRequestRequest{
								Name: name,
							}

							op, err := client.ClosePullRequest(ctx, req)
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
						Usage: "list pull-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "pull-request", Usage: "The ID of the pull request.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"), cmd.String("pull-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListPullRequestFileDiffsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPullRequestFileDiffs(ctx, req)
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
				Name:  "repositories",
				Usage: "Manage repositories resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter results.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The name of the instance in which the repository is hosted,.", Required: false},
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
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.ListRepositoriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								Instance:  cmd.String("instance"),
							}

							limit := cmd.Int("limit")
							it := client.ListRepositories(ctx, req)
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
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetRepositoryRequest{
								Name: name,
							}

							resp, err := client.GetRepository(ctx, req)
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
						Usage: "create repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repository-id", Usage: "The ID to use for the repository, which will become the final.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.CreateRepositoryRequest{
								Parent:       parent,
								RepositoryId: cmd.String("repository-id"),
							}

							op, err := client.CreateRepository(ctx, req)
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
						Usage: "update repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "False by default.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "repository.name" not yet supported.
							repository_name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing update on %s\n", repository_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete repositories",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the repository is not found, the request will.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRepository %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.DeleteRepositoryRequest{
								Name:         name,
								AllowMissing: cmd.Bool("allow-missing"),
							}

							op, err := client.DeleteRepository(ctx, req)
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
						Name:  "describe",
						Usage: "describe repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicyRepo(ctx, req)
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
						Usage: "set-iam-policy repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicyRepo(ctx, req)
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
						Usage: "test-iam-permissions repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securesourcemanager.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securesourcemanagerpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissionsRepo(ctx, req)
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
						Name:  "fetch-tree",
						Usage: "fetch-tree repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "recursive", Usage: "If true, include all subfolders and their files in the response.", Required: false},
							&cli.StringFlag{Name: "ref", Usage: "`ref` can be a SHA-1 hash, a branch name, or a tag.", Required: false},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The format is.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							repository := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing fetch-tree on %s\n", repository)
							return nil
						},
					},

					{
						Name:  "fetch-blob",
						Usage: "fetch-blob repositories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "repositorie", Usage: "The ID of the repositorie.", Required: true},
							&cli.StringFlag{Name: "repository", Usage: "The format is.", Required: true},
							&cli.StringFlag{Name: "sha", Usage: "The SHA-1 hash of the blob to retrieve.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							repository := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repositorie"))
							fmt.Printf("Executing fetch-blob on %s\n", repository)
							return nil
						},
					},
				},
			},
		},
	}
}
