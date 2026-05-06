package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1"
	"cloud.google.com/go/recaptchaenterprise/apiv1/recaptchaenterprisepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "recaptchaenterprise",
				Usage: "manage reCAPTCHA Enterprise API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "assessments",
						Usage: "Manage assessments resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create assessments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.CreateAssessmentRequest{Parent: parent}
									resp, err := client.CreateAssessment(ctx, req)
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
								Name:  "annotate",
								Usage: "annotate assessments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "assessment", Usage: "The assessment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/assessments/%s", cmd.String("project"), cmd.String("assessment"))
									fmt.Printf("Executing annotate on %s\n", name)
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
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "path", Usage: "The path.", Required: false},
									&cli.StringFlag{Name: "condition", Usage: "The condition.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.CreateFirewallPolicyRequest{Parent: parent}
									req.FirewallPolicy = &recaptchaenterprisepb.FirewallPolicy{
										Description: cmd.String("description"),
										Path:        cmd.String("path"),
										Condition:   cmd.String("condition"),
									}
									resp, err := client.CreateFirewallPolicy(ctx, req)
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
								Usage: "list firewallpolicies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe firewallpolicies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "firewallpolicy", Usage: "The firewallpolicy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicy"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.GetFirewallPolicyRequest{Name: name}
									resp, err := client.GetFirewallPolicy(ctx, req)
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
								Usage: "update firewallpolicies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "firewallpolicy", Usage: "The firewallpolicy.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "path", Usage: "The path.", Required: false},
									&cli.StringFlag{Name: "condition", Usage: "The condition.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicy"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.UpdateFirewallPolicyRequest{}
									req.FirewallPolicy = &recaptchaenterprisepb.FirewallPolicy{
										Name:        name,
										Description: cmd.String("description"),
										Path:        cmd.String("path"),
										Condition:   cmd.String("condition"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("path") {
										paths = append(paths, "path")
									}
									if cmd.IsSet("condition") {
										paths = append(paths, "condition")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateFirewallPolicy(ctx, req)
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
								Usage: "delete firewallpolicies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "firewallpolicy", Usage: "The firewallpolicy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicy"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.DeleteFirewallPolicyRequest{Name: name}
									if err := client.DeleteFirewallPolicy(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "reorder",
								Usage: "reorder firewallpolicies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reorder...")
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
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.CreateKeyRequest{Parent: parent}
									req.Key = &recaptchaenterprisepb.Key{
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateKey(ctx, req)
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
								Usage: "list keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "retrieve-legacy-secret-key",
								Usage: "retrieve-legacy-secret-key keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing retrieve-legacy-secret-key...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.GetKeyRequest{Name: name}
									resp, err := client.GetKey(ctx, req)
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
								Usage: "update keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.UpdateKeyRequest{}
									req.Key = &recaptchaenterprisepb.Key{
										Name:        name,
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateKey(ctx, req)
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
								Usage: "delete keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.DeleteKeyRequest{Name: name}
									if err := client.DeleteKey(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "migrate",
								Usage: "migrate keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
									fmt.Printf("Executing migrate on %s\n", name)
									return nil
								},
							},
							{
								Name:  "add-ip-override",
								Usage: "add-ip-override keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
									fmt.Printf("Executing add-ip-override on %s\n", name)
									return nil
								},
							},
							{
								Name:  "remove-ip-override",
								Usage: "remove-ip-override keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
									fmt.Printf("Executing remove-ip-override on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list keys",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &recaptchaenterprisepb.ListIpOverridesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListIpOverrides(ctx, req)
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
						Name:  "memberships",
						Usage: "Manage memberships resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list memberships",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "relatedaccountgroup", Usage: "The relatedaccountgroup.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/relatedaccountgroups/%s", cmd.String("project"), cmd.String("relatedaccountgroup"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &recaptchaenterprisepb.ListRelatedAccountGroupMembershipsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRelatedAccountGroupMemberships(ctx, req)
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
						Name:  "metrics",
						Usage: "Manage metrics resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe metrics",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/keys/%s/metrics", cmd.String("project"), cmd.String("key"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &recaptchaenterprisepb.GetMetricsRequest{Name: name}
									resp, err := client.GetMetrics(ctx, req)
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
						Name:  "relatedaccountgroupmemberships",
						Usage: "Manage relatedaccountgroupmemberships resources",
						Commands: []*cli.Command{
							{
								Name:  "search",
								Usage: "search relatedaccountgroupmemberships",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search...")
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := recaptchaenterprise.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &recaptchaenterprisepb.ListRelatedAccountGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRelatedAccountGroups(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
