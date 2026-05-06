package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	securesourcemanager "cloud.google.com/go/securesourcemanager/apiv1"
	"cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "securesourcemanager",
				Usage: "manage Secure Source Manager API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "branch-rules",
						Usage: "Manage branch-rules resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create branch-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "branch-rule-id", Usage: "The branch rule id.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "include-pattern", Usage: "The include pattern.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
									&cli.BoolFlag{Name: "require-pull-request", Usage: "The require pull request.", Required: false},
									&cli.IntFlag{Name: "minimum-reviews-count", Usage: "The minimum reviews count.", Required: false},
									&cli.IntFlag{Name: "minimum-approvals-count", Usage: "The minimum approvals count.", Required: false},
									&cli.BoolFlag{Name: "require-code-owner-approval", Usage: "The require code owner approval.", Required: false},
									&cli.BoolFlag{Name: "require-comments-resolved", Usage: "The require comments resolved.", Required: false},
									&cli.BoolFlag{Name: "allow-stale-reviews", Usage: "The allow stale reviews.", Required: false},
									&cli.BoolFlag{Name: "require-linear-history", Usage: "The require linear history.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreateBranchRuleRequest{Parent: parent}
									req.BranchRuleId = cmd.String("branch-rule-id")
									req.BranchRule = &securesourcemanagerpb.BranchRule{
										Etag:                     cmd.String("etag"),
										IncludePattern:           cmd.String("include-pattern"),
										Disabled:                 cmd.Bool("disabled"),
										RequirePullRequest:       cmd.Bool("require-pull-request"),
										MinimumReviewsCount:      int32(cmd.Int("minimum-reviews-count")),
										MinimumApprovalsCount:    int32(cmd.Int("minimum-approvals-count")),
										RequireCodeOwnerApproval: cmd.Bool("require-code-owner-approval"),
										RequireCommentsResolved:  cmd.Bool("require-comments-resolved"),
										AllowStaleReviews:        cmd.Bool("allow-stale-reviews"),
										RequireLinearHistory:     cmd.Bool("require-linear-history"),
									}
									op, err := client.CreateBranchRule(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list branch-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListBranchRulesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListBranchRules(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe branch-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "branch_rule", Usage: "The branch_rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("branch_rule"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetBranchRuleRequest{Name: name}
									resp, err := client.GetBranchRule(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update branch-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "branch_rule", Usage: "The branch_rule.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "include-pattern", Usage: "The include pattern.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
									&cli.BoolFlag{Name: "require-pull-request", Usage: "The require pull request.", Required: false},
									&cli.IntFlag{Name: "minimum-reviews-count", Usage: "The minimum reviews count.", Required: false},
									&cli.IntFlag{Name: "minimum-approvals-count", Usage: "The minimum approvals count.", Required: false},
									&cli.BoolFlag{Name: "require-code-owner-approval", Usage: "The require code owner approval.", Required: false},
									&cli.BoolFlag{Name: "require-comments-resolved", Usage: "The require comments resolved.", Required: false},
									&cli.BoolFlag{Name: "allow-stale-reviews", Usage: "The allow stale reviews.", Required: false},
									&cli.BoolFlag{Name: "require-linear-history", Usage: "The require linear history.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("branch_rule"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdateBranchRuleRequest{}
									req.BranchRule = &securesourcemanagerpb.BranchRule{
										Name:                     name,
										Etag:                     cmd.String("etag"),
										IncludePattern:           cmd.String("include-pattern"),
										Disabled:                 cmd.Bool("disabled"),
										RequirePullRequest:       cmd.Bool("require-pull-request"),
										MinimumReviewsCount:      int32(cmd.Int("minimum-reviews-count")),
										MinimumApprovalsCount:    int32(cmd.Int("minimum-approvals-count")),
										RequireCodeOwnerApproval: cmd.Bool("require-code-owner-approval"),
										RequireCommentsResolved:  cmd.Bool("require-comments-resolved"),
										AllowStaleReviews:        cmd.Bool("allow-stale-reviews"),
										RequireLinearHistory:     cmd.Bool("require-linear-history"),
									}
									var paths []string
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									if cmd.IsSet("include-pattern") {
										paths = append(paths, "include_pattern")
									}
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									if cmd.IsSet("require-pull-request") {
										paths = append(paths, "require_pull_request")
									}
									if cmd.IsSet("minimum-reviews-count") {
										paths = append(paths, "minimum_reviews_count")
									}
									if cmd.IsSet("minimum-approvals-count") {
										paths = append(paths, "minimum_approvals_count")
									}
									if cmd.IsSet("require-code-owner-approval") {
										paths = append(paths, "require_code_owner_approval")
									}
									if cmd.IsSet("require-comments-resolved") {
										paths = append(paths, "require_comments_resolved")
									}
									if cmd.IsSet("allow-stale-reviews") {
										paths = append(paths, "allow_stale_reviews")
									}
									if cmd.IsSet("require-linear-history") {
										paths = append(paths, "require_linear_history")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateBranchRule(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete branch-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "branch_rule", Usage: "The branch_rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/branchRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("branch_rule"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeleteBranchRuleRequest{Name: name}
									op, err := client.DeleteBranchRule(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListHooksRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListHooks(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe hooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "hook", Usage: "The hook.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("hook"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetHookRequest{Name: name}
									resp, err := client.GetHook(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create hooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "hook-id", Usage: "The hook id.", Required: true},
									&cli.StringFlag{Name: "target-uri", Usage: "The target uri.", Required: true},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
									&cli.StringFlag{Name: "sensitive-query-string", Usage: "The sensitive query string.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreateHookRequest{Parent: parent}
									req.HookId = cmd.String("hook-id")
									req.Hook = &securesourcemanagerpb.Hook{
										TargetUri:            cmd.String("target-uri"),
										Disabled:             cmd.Bool("disabled"),
										SensitiveQueryString: cmd.String("sensitive-query-string"),
									}
									op, err := client.CreateHook(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update hooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "hook", Usage: "The hook.", Required: true},
									&cli.StringFlag{Name: "target-uri", Usage: "The target uri.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
									&cli.StringFlag{Name: "sensitive-query-string", Usage: "The sensitive query string.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("hook"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdateHookRequest{}
									req.Hook = &securesourcemanagerpb.Hook{
										Name:                 name,
										TargetUri:            cmd.String("target-uri"),
										Disabled:             cmd.Bool("disabled"),
										SensitiveQueryString: cmd.String("sensitive-query-string"),
									}
									var paths []string
									if cmd.IsSet("target-uri") {
										paths = append(paths, "target_uri")
									}
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									if cmd.IsSet("sensitive-query-string") {
										paths = append(paths, "sensitive_query_string")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateHook(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete hooks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "hook", Usage: "The hook.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/hooks/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("hook"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeleteHookRequest{Name: name}
									op, err := client.DeleteHook(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListInstancesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInstances(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetInstanceRequest{Name: name}
									resp, err := client.GetInstance(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance-id", Usage: "The instance id.", Required: true},
									&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreateInstanceRequest{Parent: parent}
									req.InstanceId = cmd.String("instance-id")
									req.Instance = &securesourcemanagerpb.Instance{
										KmsKey: cmd.String("kms-key"),
									}
									op, err := client.CreateInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeleteInstanceRequest{Name: name}
									op, err := client.DeleteInstance(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "body", Usage: "The body.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreateIssueCommentRequest{Parent: parent}
									req.IssueComment = &securesourcemanagerpb.IssueComment{
										Body: cmd.String("body"),
									}
									op, err := client.CreateIssueComment(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe issue-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "comment", Usage: "The comment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s/issueComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"), cmd.String("comment"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetIssueCommentRequest{Name: name}
									resp, err := client.GetIssueComment(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list issue-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListIssueCommentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListIssueComments(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update issue-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "comment", Usage: "The comment.", Required: true},
									&cli.StringFlag{Name: "body", Usage: "The body.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s/issueComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"), cmd.String("comment"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdateIssueCommentRequest{}
									req.IssueComment = &securesourcemanagerpb.IssueComment{
										Name: name,
										Body: cmd.String("body"),
									}
									var paths []string
									if cmd.IsSet("body") {
										paths = append(paths, "body")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateIssueComment(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete issue-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "comment", Usage: "The comment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s/issueComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"), cmd.String("comment"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeleteIssueCommentRequest{Name: name}
									op, err := client.DeleteIssueComment(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "title", Usage: "The title.", Required: true},
									&cli.StringFlag{Name: "body", Usage: "The body.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreateIssueRequest{Parent: parent}
									req.Issue = &securesourcemanagerpb.Issue{
										Title: cmd.String("title"),
										Body:  cmd.String("body"),
										Etag:  cmd.String("etag"),
									}
									op, err := client.CreateIssue(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe issues",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetIssueRequest{Name: name}
									resp, err := client.GetIssue(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list issues",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListIssuesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListIssues(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update issues",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "title", Usage: "The title.", Required: false},
									&cli.StringFlag{Name: "body", Usage: "The body.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdateIssueRequest{}
									req.Issue = &securesourcemanagerpb.Issue{
										Name:  name,
										Title: cmd.String("title"),
										Body:  cmd.String("body"),
										Etag:  cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("title") {
										paths = append(paths, "title")
									}
									if cmd.IsSet("body") {
										paths = append(paths, "body")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateIssue(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete issues",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeleteIssueRequest{Name: name}
									op, err := client.DeleteIssue(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "open",
								Usage: "open issues",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.OpenIssueRequest{Name: name}
									req.Etag = cmd.String("etag")
									op, err := client.OpenIssue(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "close",
								Usage: "close issues",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("issue"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CloseIssueRequest{Name: name}
									req.Etag = cmd.String("etag")
									op, err := client.CloseIssue(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
									&cli.StringFlag{Name: "comment", Usage: "The comment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s/pullRequestComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"), cmd.String("comment"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetPullRequestCommentRequest{Name: name}
									resp, err := client.GetPullRequestComment(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListPullRequestCommentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPullRequestComments(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreatePullRequestCommentRequest{Parent: parent}
									op, err := client.CreatePullRequestComment(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
									&cli.StringFlag{Name: "comment", Usage: "The comment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s/pullRequestComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"), cmd.String("comment"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdatePullRequestCommentRequest{}
									req.PullRequestComment = &securesourcemanagerpb.PullRequestComment{
										Name: name,
									}
									op, err := client.UpdatePullRequestComment(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
									&cli.StringFlag{Name: "comment", Usage: "The comment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s/pullRequestComments/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"), cmd.String("comment"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeletePullRequestCommentRequest{Name: name}
									op, err := client.DeletePullRequestComment(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "batch-create",
								Usage: "batch-create pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									fmt.Printf("Executing batch-create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "resolve",
								Usage: "resolve pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									fmt.Printf("Executing resolve on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "unresolve",
								Usage: "unresolve pull-request-comments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									fmt.Printf("Executing unresolve on %s\n", parent)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "title", Usage: "The title.", Required: true},
									&cli.StringFlag{Name: "body", Usage: "The body.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreatePullRequestRequest{Parent: parent}
									req.PullRequest = &securesourcemanagerpb.PullRequest{
										Title: cmd.String("title"),
										Body:  cmd.String("body"),
									}
									op, err := client.CreatePullRequest(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetPullRequestRequest{Name: name}
									resp, err := client.GetPullRequest(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListPullRequestsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPullRequests(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
									&cli.StringFlag{Name: "title", Usage: "The title.", Required: false},
									&cli.StringFlag{Name: "body", Usage: "The body.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdatePullRequestRequest{}
									req.PullRequest = &securesourcemanagerpb.PullRequest{
										Name:  name,
										Title: cmd.String("title"),
										Body:  cmd.String("body"),
									}
									var paths []string
									if cmd.IsSet("title") {
										paths = append(paths, "title")
									}
									if cmd.IsSet("body") {
										paths = append(paths, "body")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdatePullRequest(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "merge",
								Usage: "merge pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.MergePullRequestRequest{Name: name}
									op, err := client.MergePullRequest(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "open",
								Usage: "open pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.OpenPullRequestRequest{Name: name}
									op, err := client.OpenPullRequest(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "close",
								Usage: "close pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "pull_request", Usage: "The pull_request.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s/pullRequests/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"), cmd.String("pull_request"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.ClosePullRequestRequest{Name: name}
									op, err := client.ClosePullRequest(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list pull-requests",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListPullRequestFileDiffsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListPullRequestFileDiffs(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &securesourcemanagerpb.ListRepositoriesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRepositories(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.GetRepositoryRequest{Name: name}
									resp, err := client.GetRepository(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository-id", Usage: "The repository id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.CreateRepositoryRequest{Parent: parent}
									req.RepositoryId = cmd.String("repository-id")
									req.Repository = &securesourcemanagerpb.Repository{
										Description: cmd.String("description"),
										Instance:    cmd.String("instance"),
										Etag:        cmd.String("etag"),
									}
									op, err := client.CreateRepository(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.UpdateRepositoryRequest{}
									req.Repository = &securesourcemanagerpb.Repository{
										Name:        name,
										Description: cmd.String("description"),
										Instance:    cmd.String("instance"),
										Etag:        cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("instance") {
										paths = append(paths, "instance")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateRepository(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete repositories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "repository", Usage: "The repository.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", cmd.String("project"), cmd.String("location"), cmd.String("repository"))
									client, err := securesourcemanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &securesourcemanagerpb.DeleteRepositoryRequest{Name: name}
									op, err := client.DeleteRepository(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "get-iam-policy-repo",
								Usage: "get-iam-policy-repo repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy-repo...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy-repo",
								Usage: "set-iam-policy-repo repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy-repo...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions-repo",
								Usage: "test-iam-permissions-repo repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions-repo...")
									return nil
								},
							},
							{
								Name:  "fetch-tree",
								Usage: "fetch-tree repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-tree...")
									return nil
								},
							},
							{
								Name:  "fetch-blob",
								Usage: "fetch-blob repositories",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-blob...")
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
