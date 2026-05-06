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

package essentialcontacts

import (
	essentialcontacts "cloud.google.com/go/essentialcontacts/apiv1"
	"cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the essentialcontacts command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "essentialcontacts",
		Usage: "manage Essential Contacts API resources",
		Commands: []*cli.Command{
			{
				Name:  "contacts",
				Usage: "Manage contacts resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := essentialcontacts.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &essentialcontactspb.CreateContactRequest{
								Parent: parent,
							}

							resp, err := client.CreateContact(ctx, req)
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
						Usage: "update contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contact", Usage: "The ID of the contact.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "contact.name" not yet supported.
							contact_name := fmt.Sprintf("projects/%s/contacts/%s", cmd.String("project"), cmd.String("contact"))
							fmt.Printf("Executing update on %s\n", contact_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list contacts",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, retrieves the next batch of results from the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := essentialcontacts.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &essentialcontactspb.ListContactsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListContacts(ctx, req)
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
						Usage: "describe contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contact", Usage: "The ID of the contact.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/contacts/%s", cmd.String("project"), cmd.String("contact"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := essentialcontacts.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &essentialcontactspb.GetContactRequest{
								Name: name,
							}

							resp, err := client.GetContact(ctx, req)
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
						Usage: "delete contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contact", Usage: "The ID of the contact.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/contacts/%s", cmd.String("project"), cmd.String("contact"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteContact on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := essentialcontacts.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &essentialcontactspb.DeleteContactRequest{
								Name: name,
							}

							if err := client.DeleteContact(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "compute",
						Usage: "compute contacts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "notification-categories", Usage: "The categories of notifications to compute contacts for.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return from this request.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If present, retrieves the next batch of results from the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := essentialcontacts.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &essentialcontactspb.ComputeContactsRequest{
								Parent:                 parent,
								NotificationCategories: cmd.StringSlice("notification-categories"),
								PageSize:               int32(cmd.Int("page-size")),
								PageToken:              cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ComputeContacts(ctx, req)
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
						Name:  "send-test-message",
						Usage: "send-test-message contacts",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "contacts", Usage: "The list of names of the contacts to send a test message to.", Required: true},
							&cli.StringFlag{Name: "notification-category", Usage: "The notification category to send the test message for.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s", cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute SendTestMessage on %s?", resource)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := essentialcontacts.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &essentialcontactspb.SendTestMessageRequest{
								Resource:             resource,
								Contacts:             cmd.StringSlice("contacts"),
								NotificationCategory: essentialcontactspb.NotificationCategory(essentialcontactspb.NotificationCategory_value[cmd.String("notification-category")]),
							}

							if err := client.SendTestMessage(ctx, req); err != nil {
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
