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

package accessapproval

import (
	accessapproval "cloud.google.com/go/accessapproval/apiv1"
	"cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the accessapproval command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "accessapproval",
		Usage: "manage Access Approval API resources",
		Commands: []*cli.Command{
			{
				Name:  "access-approval-settings",
				Usage: "Manage access-approval-settings resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe access-approval-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/accessApprovalSettings", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.GetAccessApprovalSettingsMessage{
								Name: name,
							}

							resp, err := client.GetAccessApprovalSettings(ctx, req)
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
						Usage: "update access-approval-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "settings.name" not yet supported.
							settings_name := fmt.Sprintf("projects/%s/accessApprovalSettings", cmd.String("project"))
							fmt.Printf("Executing update on %s\n", settings_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete access-approval-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/accessApprovalSettings", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAccessApprovalSettings on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.DeleteAccessApprovalSettingsMessage{
								Name: name,
							}

							if err := client.DeleteAccessApprovalSettings(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "approval-requests",
				Usage: "Manage approval-requests resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list approval-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter on the type of approval requests to retrieve.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.ListApprovalRequestsMessage{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListApprovalRequests(ctx, req)
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
						Usage: "describe approval-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "approval-request", Usage: "The ID of the approval request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.GetApprovalRequestMessage{
								Name: name,
							}

							resp, err := client.GetApprovalRequest(ctx, req)
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
						Usage: "approve approval-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "approval-request", Usage: "The ID of the approval request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.ApproveApprovalRequestMessage{
								Name: name,
							}

							resp, err := client.ApproveApprovalRequest(ctx, req)
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
						Name:  "dismiss",
						Usage: "dismiss approval-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "approval-request", Usage: "The ID of the approval request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.DismissApprovalRequestMessage{
								Name: name,
							}

							resp, err := client.DismissApprovalRequest(ctx, req)
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
						Name:  "invalidate",
						Usage: "invalidate approval-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "approval-request", Usage: "The ID of the approval request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/approvalRequests/%s", cmd.String("project"), cmd.String("approval-request"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.InvalidateApprovalRequestMessage{
								Name: name,
							}

							resp, err := client.InvalidateApprovalRequest(ctx, req)
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
				Name:  "service-account",
				Usage: "Manage service-account resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe service-account",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/serviceAccount", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := accessapproval.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &accessapprovalpb.GetAccessApprovalServiceAccountMessage{
								Name: name,
							}

							resp, err := client.GetAccessApprovalServiceAccount(ctx, req)
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
