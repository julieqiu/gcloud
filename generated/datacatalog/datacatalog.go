package datacatalog

import (
	"context"
	"fmt"
	"strings"

	datacatalog "cloud.google.com/go/datacatalog/apiv1"
	"cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud datacatalog command tree.
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search...")
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry-id", Usage: "The entry id.", Required: true},
							&cli.StringFlag{Name: "linked-resource", Usage: "The linked resource.", Required: false},
							&cli.StringFlag{Name: "fully-qualified-name", Usage: "The fully qualified name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreateEntryRequest{Parent: parent}
							req.EntryId = cmd.String("entry-id")
							req.Entry = &datacatalogpb.Entry{
								LinkedResource:     cmd.String("linked-resource"),
								FullyQualifiedName: cmd.String("fully-qualified-name"),
								DisplayName:        cmd.String("display-name"),
								Description:        cmd.String("description"),
							}
							resp, err := client.CreateEntry(ctx, req)
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
						Usage: "update entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
							&cli.StringFlag{Name: "linked-resource", Usage: "The linked resource.", Required: false},
							&cli.StringFlag{Name: "fully-qualified-name", Usage: "The fully qualified name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdateEntryRequest{}
							req.Entry = &datacatalogpb.Entry{
								Name:               name,
								LinkedResource:     cmd.String("linked-resource"),
								FullyQualifiedName: cmd.String("fully-qualified-name"),
								DisplayName:        cmd.String("display-name"),
								Description:        cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("linked-resource") {
								paths = append(paths, "linked_resource")
							}
							if cmd.IsSet("fully-qualified-name") {
								paths = append(paths, "fully_qualified_name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateEntry(ctx, req)
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
						Usage: "delete entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeleteEntryRequest{Name: name}
							if err := client.DeleteEntry(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.GetEntryRequest{Name: name}
							resp, err := client.GetEntry(ctx, req)
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
						Name:  "lookup",
						Usage: "lookup entries",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing lookup...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datacatalogpb.ListEntriesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListEntries(ctx, req)
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
						Name:  "modify-entry-overview",
						Usage: "modify-entry-overview entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							fmt.Printf("Executing modify-entry-overview on %s\n", name)
							return nil
						},
					},
					{
						Name:  "modify-entry-contacts",
						Usage: "modify-entry-contacts entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							fmt.Printf("Executing modify-entry-contacts on %s\n", name)
							return nil
						},
					},
					{
						Name:  "star",
						Usage: "star entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							fmt.Printf("Executing star on %s\n", name)
							return nil
						},
					},
					{
						Name:  "unstar",
						Usage: "unstar entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							fmt.Printf("Executing unstar on %s\n", name)
							return nil
						},
					},
					{
						Name:  "import",
						Usage: "import entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"))
							fmt.Printf("Executing import on %s\n", parent)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry-group-id", Usage: "The entry group id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.BoolFlag{Name: "transferred-to-dataplex", Usage: "The transferred to dataplex.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreateEntryGroupRequest{Parent: parent}
							req.EntryGroupId = cmd.String("entry-group-id")
							req.EntryGroup = &datacatalogpb.EntryGroup{
								DisplayName:           cmd.String("display-name"),
								Description:           cmd.String("description"),
								TransferredToDataplex: cmd.Bool("transferred-to-dataplex"),
							}
							resp, err := client.CreateEntryGroup(ctx, req)
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
						Usage: "describe entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.GetEntryGroupRequest{Name: name}
							resp, err := client.GetEntryGroup(ctx, req)
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
						Usage: "update entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.BoolFlag{Name: "transferred-to-dataplex", Usage: "The transferred to dataplex.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdateEntryGroupRequest{}
							req.EntryGroup = &datacatalogpb.EntryGroup{
								Name:                  name,
								DisplayName:           cmd.String("display-name"),
								Description:           cmd.String("description"),
								TransferredToDataplex: cmd.Bool("transferred-to-dataplex"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("transferred-to-dataplex") {
								paths = append(paths, "transferred_to_dataplex")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateEntryGroup(ctx, req)
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
						Usage: "delete entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeleteEntryGroupRequest{Name: name}
							if err := client.DeleteEntryGroup(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datacatalogpb.ListEntryGroupsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListEntryGroups(ctx, req)
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
				Name:  "enum-values",
				Usage: "Manage enum-values resources",
				Commands: []*cli.Command{
					{
						Name:  "rename",
						Usage: "rename enum-values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
							&cli.StringFlag{Name: "tag_template_field_id", Usage: "The tag_template_field_id.", Required: true},
							&cli.StringFlag{Name: "enum_value_display_name", Usage: "The enum_value_display_name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s/enumValues/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"), cmd.String("tag_template_field_id"), cmd.String("enum_value_display_name"))
							fmt.Printf("Executing rename on %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
							&cli.StringFlag{Name: "tag-template-field-id", Usage: "The tag template field id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.BoolFlag{Name: "is-required", Usage: "The is required.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.IntFlag{Name: "order", Usage: "The order.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreateTagTemplateFieldRequest{Parent: parent}
							req.TagTemplateFieldId = cmd.String("tag-template-field-id")
							req.TagTemplateField = &datacatalogpb.TagTemplateField{
								DisplayName: cmd.String("display-name"),
								IsRequired:  cmd.Bool("is-required"),
								Description: cmd.String("description"),
								Order:       int32(cmd.Int("order")),
							}
							resp, err := client.CreateTagTemplateField(ctx, req)
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
						Usage: "update fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
							&cli.StringFlag{Name: "field", Usage: "The field.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.BoolFlag{Name: "is-required", Usage: "The is required.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.IntFlag{Name: "order", Usage: "The order.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"), cmd.String("field"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdateTagTemplateFieldRequest{}
							req.TagTemplateField = &datacatalogpb.TagTemplateField{
								Name:        name,
								DisplayName: cmd.String("display-name"),
								IsRequired:  cmd.Bool("is-required"),
								Description: cmd.String("description"),
								Order:       int32(cmd.Int("order")),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("is-required") {
								paths = append(paths, "is_required")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("order") {
								paths = append(paths, "order")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateTagTemplateField(ctx, req)
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
						Name:  "rename",
						Usage: "rename fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
							&cli.StringFlag{Name: "field", Usage: "The field.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"), cmd.String("field"))
							fmt.Printf("Executing rename on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete fields",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
							&cli.StringFlag{Name: "field", Usage: "The field.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s/fields/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"), cmd.String("field"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeleteTagTemplateFieldRequest{Name: name}
							if err := client.DeleteTagTemplateField(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing set-config...")
							return nil
						},
					},
					{
						Name:  "retrieve-config",
						Usage: "retrieve-config locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing retrieve-config...")
							return nil
						},
					},
					{
						Name:  "retrieve-effective-config",
						Usage: "retrieve-effective-config locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing retrieve-effective-config...")
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
							client, err := datacatalog.NewClient(ctx)
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
							client, err := datacatalog.NewClient(ctx)
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
							client, err := datacatalog.NewClient(ctx)
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
							client, err := datacatalog.NewClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerSerializationClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerSerializationClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerSerializationClient(ctx)
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
							client, err := datacatalog.NewPolicyTagManagerSerializationClient(ctx)
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
				Name:  "policy-tags",
				Usage: "Manage policy-tags resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "parent-policy-tag", Usage: "The parent policy tag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreatePolicyTagRequest{Parent: parent}
							req.PolicyTag = &datacatalogpb.PolicyTag{
								DisplayName:     cmd.String("display-name"),
								Description:     cmd.String("description"),
								ParentPolicyTag: cmd.String("parent-policy-tag"),
							}
							resp, err := client.CreatePolicyTag(ctx, req)
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
						Usage: "delete policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
							&cli.StringFlag{Name: "policy_tag", Usage: "The policy_tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"), cmd.String("policy_tag"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeletePolicyTagRequest{Name: name}
							if err := client.DeletePolicyTag(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
							&cli.StringFlag{Name: "policy_tag", Usage: "The policy_tag.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "parent-policy-tag", Usage: "The parent policy tag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"), cmd.String("policy_tag"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdatePolicyTagRequest{}
							req.PolicyTag = &datacatalogpb.PolicyTag{
								Name:            name,
								DisplayName:     cmd.String("display-name"),
								Description:     cmd.String("description"),
								ParentPolicyTag: cmd.String("parent-policy-tag"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("parent-policy-tag") {
								paths = append(paths, "parent_policy_tag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdatePolicyTag(ctx, req)
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
						Usage: "list policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datacatalogpb.ListPolicyTagsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPolicyTags(ctx, req)
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
						Usage: "describe policy-tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
							&cli.StringFlag{Name: "policy_tag", Usage: "The policy_tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"), cmd.String("policy_tag"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.GetPolicyTagRequest{Name: name}
							resp, err := client.GetPolicyTag(ctx, req)
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
				Name:  "tag-templates",
				Usage: "Manage tag-templates resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag-template-id", Usage: "The tag template id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.BoolFlag{Name: "is-publicly-readable", Usage: "The is publicly readable.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreateTagTemplateRequest{Parent: parent}
							req.TagTemplateId = cmd.String("tag-template-id")
							req.TagTemplate = &datacatalogpb.TagTemplate{
								DisplayName:        cmd.String("display-name"),
								IsPubliclyReadable: cmd.Bool("is-publicly-readable"),
							}
							resp, err := client.CreateTagTemplate(ctx, req)
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
						Usage: "describe tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.GetTagTemplateRequest{Name: name}
							resp, err := client.GetTagTemplate(ctx, req)
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
						Usage: "update tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.BoolFlag{Name: "is-publicly-readable", Usage: "The is publicly readable.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdateTagTemplateRequest{}
							req.TagTemplate = &datacatalogpb.TagTemplate{
								Name:               name,
								DisplayName:        cmd.String("display-name"),
								IsPubliclyReadable: cmd.Bool("is-publicly-readable"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("is-publicly-readable") {
								paths = append(paths, "is_publicly_readable")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateTagTemplate(ctx, req)
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
						Usage: "delete tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tag_template", Usage: "The tag_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tag_template"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeleteTagTemplateRequest{Name: name}
							if err := client.DeleteTagTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tagTemplate", Usage: "The tagTemplate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tagTemplate"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tagTemplate", Usage: "The tagTemplate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tagTemplate"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions tag-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tagTemplate", Usage: "The tagTemplate.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tagTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("tagTemplate"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
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
				Name:  "tags",
				Usage: "Manage tags resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
							&cli.StringFlag{Name: "template", Usage: "The template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreateTagRequest{Parent: parent}
							req.Tag = &datacatalogpb.Tag{
								Template: cmd.String("template"),
							}
							resp, err := client.CreateTag(ctx, req)
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
						Usage: "update tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The tag.", Required: true},
							&cli.StringFlag{Name: "template", Usage: "The template.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"), cmd.String("tag"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdateTagRequest{}
							req.Tag = &datacatalogpb.Tag{
								Name:     name,
								Template: cmd.String("template"),
							}
							var paths []string
							if cmd.IsSet("template") {
								paths = append(paths, "template")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateTag(ctx, req)
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
						Usage: "delete tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
							&cli.StringFlag{Name: "tag", Usage: "The tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s/tags/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"), cmd.String("tag"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeleteTagRequest{Name: name}
							if err := client.DeleteTag(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							client, err := datacatalog.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datacatalogpb.ListTagsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListTags(ctx, req)
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
						Name:  "reconcile",
						Usage: "reconcile tags",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "entry_group", Usage: "The entry_group.", Required: true},
							&cli.StringFlag{Name: "entry", Usage: "The entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry_group"), cmd.String("entry"))
							fmt.Printf("Executing reconcile on %s\n", parent)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.CreateTaxonomyRequest{Parent: parent}
							req.Taxonomy = &datacatalogpb.Taxonomy{
								DisplayName: cmd.String("display-name"),
								Description: cmd.String("description"),
							}
							resp, err := client.CreateTaxonomy(ctx, req)
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
						Usage: "delete taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.DeleteTaxonomyRequest{Name: name}
							if err := client.DeleteTaxonomy(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.UpdateTaxonomyRequest{}
							req.Taxonomy = &datacatalogpb.Taxonomy{
								Name:        name,
								DisplayName: cmd.String("display-name"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateTaxonomy(ctx, req)
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
						Usage: "list taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datacatalogpb.ListTaxonomiesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListTaxonomies(ctx, req)
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
						Usage: "describe taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datacatalogpb.GetTaxonomyRequest{Name: name}
							resp, err := client.GetTaxonomy(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The taxonomie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomie", Usage: "The taxonomie.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomie"))
							client, err := datacatalog.NewPolicyTagManagerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
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
						Name:  "replace",
						Usage: "replace taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "taxonomy", Usage: "The taxonomy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("taxonomy"))
							fmt.Printf("Executing replace on %s\n", name)
							return nil
						},
					},
					{
						Name:  "import",
						Usage: "import taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing import on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing export on %s\n", parent)
							return nil
						},
					},
				},
			},
		},
	}
}
