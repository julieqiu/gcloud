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

package datacatalog

import (
	datacatalog "cloud.google.com/go/datacatalog/apiv1"
	"cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the datacatalog command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "datacatalog",
		Usage: "manage Google Cloud Data Catalog API resources",
		Commands: []*cli.Command{
			{
				Name:  "catalog",
				Usage: "Manage catalog resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search catalog",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "admin-search", Usage: "If set, use searchAll permission granted on organizations from.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the order of results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Upper bound on the number of results you can get in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token that, if specified, returns the next page of.", Required: false},
							&cli.StringFlag{Name: "query", Usage: "The query string with a minimum of 3 characters and specific.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.SearchCatalogRequest{
								Query:       cmd.String("query"),
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								OrderBy:     cmd.String("order-by"),
								AdminSearch: cmd.Bool("admin-search"),
							}

							limit := cmd.Int("limit")
							it := client.SearchCatalog(ctx, req)
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
				Name:  "entries",
				Usage: "Manage entries resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "entry-id", Usage: "The ID of the entry to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreateEntryRequest{
								Parent:  parent,
								EntryId: cmd.String("entry-id"),
							}

							resp, err := client.CreateEntry(ctx, req)
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
						Usage: "update entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entry.name" not yet supported.
							entry_name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							fmt.Printf("Executing update on %s\n", entry_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEntry on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeleteEntryRequest{
								Name: name,
							}

							if err := client.DeleteEntry(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetEntryRequest{
								Name: name,
							}

							resp, err := client.GetEntry(ctx, req)
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
						Name:  "lookup",
						Usage: "lookup entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "Location where the lookup should be performed.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "Project where the lookup should be performed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.LookupEntryRequest{
								Project:  cmd.String("project"),
								Location: cmd.String("location"),
							}

							resp, err := client.LookupEntry(ctx, req)
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
						Usage: "list entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token that specifies the next page to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ListEntriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntries(ctx, req)
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
						Name:  "modify-entry-overview",
						Usage: "modify-entry-overview entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ModifyEntryOverviewRequest{
								Name: name,
							}

							resp, err := client.ModifyEntryOverview(ctx, req)
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
						Name:  "modify-entry-contacts",
						Usage: "modify-entry-contacts entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ModifyEntryContactsRequest{
								Name: name,
							}

							resp, err := client.ModifyEntryContacts(ctx, req)
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
						Name:  "star",
						Usage: "star entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.StarEntryRequest{
								Name: name,
							}

							resp, err := client.StarEntry(ctx, req)
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
						Name:  "unstar",
						Usage: "unstar entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.UnstarEntryRequest{
								Name: name,
							}

							resp, err := client.UnstarEntry(ctx, req)
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
						Name:  "import",
						Usage: "import entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "job-id", Usage: "(Optional) Dataplex task job id, if specified will be used as.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ImportEntriesRequest{
								Parent: parent,
								JobId:  cmd.String("job-id"),
							}

							op, err := client.ImportEntries(ctx, req)
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
				Name:  "entry-groups",
				Usage: "Manage entry-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group-id", Usage: "The ID of the entry group to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreateEntryGroupRequest{
								Parent:       parent,
								EntryGroupId: cmd.String("entry-group-id"),
							}

							resp, err := client.CreateEntryGroup(ctx, req)
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
						Usage: "describe entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetEntryGroupRequest{
								Name: name,
							}

							resp, err := client.GetEntryGroup(ctx, req)
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
						Usage: "update entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entry_group.name" not yet supported.
							entry_group_name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							fmt.Printf("Executing update on %s\n", entry_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If true, deletes all entries in the entry group.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEntryGroup on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeleteEntryGroupRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteEntryGroup(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token that specifies the next page to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ListEntryGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntryGroups(ctx, req)
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
				Name:  "enum-values",
				Usage: "Manage enum-values resources",
				Commands: []*cli.Command{

					{
						Name:  "rename",
						Usage: "rename enum-values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "enum-value", Usage: "The ID of the enum value.", Required: true},
							&cli.StringFlag{Name: "field", Usage: "The ID of the field.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-enum-value-display-name", Usage: "The new display name of the enum value.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s/enumValues/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"), cmd.String("field"), cmd.String("enum-value"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.RenameTagTemplateFieldEnumValueRequest{
								Name:                    name,
								NewEnumValueDisplayName: cmd.String("new-enum-value-display-name"),
							}

							resp, err := client.RenameTagTemplateFieldEnumValue(ctx, req)
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
				Name:  "fields",
				Usage: "Manage fields resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
							&cli.StringFlag{Name: "tag-template-field-id", Usage: "The ID of the tag template field to create.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreateTagTemplateFieldRequest{
								Parent:             parent,
								TagTemplateFieldId: cmd.String("tag-template-field-id"),
							}

							resp, err := client.CreateTagTemplateField(ctx, req)
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
						Usage: "update fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "field", Usage: "The ID of the field.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"), cmd.String("field"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.UpdateTagTemplateFieldRequest{
								Name: name,
							}

							resp, err := client.UpdateTagTemplateField(ctx, req)
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
						Name:  "rename",
						Usage: "rename fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "field", Usage: "The ID of the field.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "new-tag-template-field-id", Usage: "The new ID of this tag template field.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"), cmd.String("field"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.RenameTagTemplateFieldRequest{
								Name:                  name,
								NewTagTemplateFieldId: cmd.String("new-tag-template-field-id"),
							}

							resp, err := client.RenameTagTemplateField(ctx, req)
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
						Usage: "delete fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "field", Usage: "The ID of the field.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If true, deletes this field from any tags that use it.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"), cmd.String("field"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTagTemplateField on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeleteTagTemplateFieldRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteTagTemplateField(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "set-config",
						Usage: "set-config locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.SetConfigRequest{
								Name: name,
							}

							resp, err := client.SetConfig(ctx, req)
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
						Name:  "retrieve-config",
						Usage: "retrieve-config locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.RetrieveConfigRequest{
								Name: name,
							}

							resp, err := client.RetrieveConfig(ctx, req)
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
						Name:  "retrieve-effective-config",
						Usage: "retrieve-effective-config locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.RetrieveEffectiveConfigRequest{
								Name: name,
							}

							resp, err := client.RetrieveEffectiveConfig(ctx, req)
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
				Name:  "policy-tags",
				Usage: "Manage policy-tags resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreatePolicyTagRequest{
								Parent: parent,
							}

							resp, err := client.CreatePolicyTag(ctx, req)
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
						Usage: "delete policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-tag", Usage: "The ID of the policy tag.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"), cmd.String("policy-tag"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePolicyTag on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeletePolicyTagRequest{
								Name: name,
							}

							if err := client.DeletePolicyTag(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-tag", Usage: "The ID of the policy tag.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "policy_tag.name" not yet supported.
							policy_tag_name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"), cmd.String("policy-tag"))
							fmt.Printf("Executing update on %s\n", policy_tag_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The pagination token of the next results page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ListPolicyTagsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPolicyTags(ctx, req)
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
						Usage: "describe policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "policy-tag", Usage: "The ID of the policy tag.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"), cmd.String("policy-tag"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetPolicyTagRequest{
								Name: name,
							}

							resp, err := client.GetPolicyTag(ctx, req)
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
				Name:  "tag-templates",
				Usage: "Manage tag-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template-id", Usage: "The ID of the tag template to create.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreateTagTemplateRequest{
								Parent:        parent,
								TagTemplateId: cmd.String("tag-template-id"),
							}

							resp, err := client.CreateTagTemplate(ctx, req)
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
						Usage: "describe tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetTagTemplateRequest{
								Name: name,
							}

							resp, err := client.GetTagTemplate(ctx, req)
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
						Usage: "update tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tag_template.name" not yet supported.
							tag_template_name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							fmt.Printf("Executing update on %s\n", tag_template_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tag-templates",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If true, deletes all tags that use this template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTagTemplate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeleteTagTemplateRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteTagTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.SetIamPolicyRequest{
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
						Usage: "get-iam-policy tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetIamPolicyRequest{
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
						Usage: "test-iam-permissions tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The ID of the tag template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.TestIamPermissionsRequest{
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
				Name:  "tags",
				Usage: "Manage tags resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreateTagRequest{
								Parent: parent,
							}

							resp, err := client.CreateTag(ctx, req)
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
						Usage: "update tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The ID of the tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tag.name" not yet supported.
							tag_name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"), cmd.String("tag"))
							fmt.Printf("Executing update on %s\n", tag_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The ID of the tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"), cmd.String("tag"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTag on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeleteTagRequest{
								Name: name,
							}

							if err := client.DeleteTag(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of tags to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token that specifies the next page to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ListTagsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTags(ctx, req)
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
						Name:  "reconcile",
						Usage: "reconcile tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.BoolFlag{Name: "force-delete-missing", Usage: "If set to `true`, deletes entry tags related to a tag template.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tag-template", Usage: "The name of the tag template, which is used for reconciliation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ReconcileTagsRequest{
								Parent:             parent,
								TagTemplate:        cmd.String("tag-template"),
								ForceDeleteMissing: cmd.Bool("force-delete-missing"),
							}

							op, err := client.ReconcileTags(ctx, req)
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
				Name:  "taxonomies",
				Usage: "Manage taxonomies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.CreateTaxonomyRequest{
								Parent: parent,
							}

							resp, err := client.CreateTaxonomy(ctx, req)
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
						Usage: "delete taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTaxonomy on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.DeleteTaxonomyRequest{
								Name: name,
							}

							if err := client.DeleteTaxonomy(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "taxonomy.name" not yet supported.
							taxonomy_name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							fmt.Printf("Executing update on %s\n", taxonomy_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Supported field for filter is 'service' and value is 'dataplex'.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The pagination token of the next results page.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ListTaxonomiesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListTaxonomies(ctx, req)
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
						Usage: "describe taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetTaxonomyRequest{
								Name: name,
							}

							resp, err := client.GetTaxonomy(ctx, req)
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
						Usage: "get-iam-policy taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.GetIamPolicyRequest{
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
						Usage: "set-iam-policy taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.SetIamPolicyRequest{
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
						Usage: "test-iam-permissions taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.TestIamPermissionsRequest{
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

					{
						Name:  "replace",
						Usage: "replace taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The ID of the taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ReplaceTaxonomyRequest{
								Name: name,
							}

							resp, err := client.ReplaceTaxonomy(ctx, req)
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
						Name:  "import",
						Usage: "import taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ImportTaxonomiesRequest{
								Parent: parent,
							}

							resp, err := client.ImportTaxonomies(ctx, req)
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
						Name:  "export",
						Usage: "export taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "taxonomies", Usage: "Resource names of the taxonomies to export.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := datacatalog.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &datacatalogpb.ExportTaxonomiesRequest{
								Parent:     parent,
								Taxonomies: cmd.StringSlice("taxonomies"),
							}

							resp, err := client.ExportTaxonomies(ctx, req)
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
