package essentialcontacts

import (
	"context"
	"fmt"
	"strings"

	essentialcontacts "cloud.google.com/go/essentialcontacts/apiv1"
	"cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud essentialcontacts command tree.
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
							&cli.StringFlag{Name: "email", Usage: "The email.", Required: true},
							&cli.StringFlag{Name: "language-tag", Usage: "The language tag.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := essentialcontacts.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &essentialcontactspb.CreateContactRequest{Parent: parent}
							req.Contact = &essentialcontactspb.Contact{
								Email:       cmd.String("email"),
								LanguageTag: cmd.String("language-tag"),
							}
							resp, err := client.CreateContact(ctx, req)
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
						Usage: "update contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contact", Usage: "The contact.", Required: true},
							&cli.StringFlag{Name: "email", Usage: "The email.", Required: false},
							&cli.StringFlag{Name: "language-tag", Usage: "The language tag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/contacts/%s", cmd.String("project"), cmd.String("contact"))
							client, err := essentialcontacts.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &essentialcontactspb.UpdateContactRequest{}
							req.Contact = &essentialcontactspb.Contact{
								Name:        name,
								Email:       cmd.String("email"),
								LanguageTag: cmd.String("language-tag"),
							}
							var paths []string
							if cmd.IsSet("email") {
								paths = append(paths, "email")
							}
							if cmd.IsSet("language-tag") {
								paths = append(paths, "language_tag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateContact(ctx, req)
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
						Usage: "list contacts",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := essentialcontacts.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &essentialcontactspb.ListContactsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListContacts(ctx, req)
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
						Usage: "describe contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contact", Usage: "The contact.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/contacts/%s", cmd.String("project"), cmd.String("contact"))
							client, err := essentialcontacts.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &essentialcontactspb.GetContactRequest{Name: name}
							resp, err := client.GetContact(ctx, req)
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
						Usage: "delete contacts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "contact", Usage: "The contact.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/contacts/%s", cmd.String("project"), cmd.String("contact"))
							client, err := essentialcontacts.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &essentialcontactspb.DeleteContactRequest{Name: name}
							if err := client.DeleteContact(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "compute",
						Usage: "compute contacts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing compute on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "send-test-message",
						Usage: "send-test-message contacts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing send-test-message...")
							return nil
						},
					},
				},
			},
		},
	}
}
