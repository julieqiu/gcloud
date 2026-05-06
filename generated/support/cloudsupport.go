package cloudsupport

import (
	"context"
	"fmt"
	"strings"

	support "cloud.google.com/go/support/apiv2"
	"cloud.google.com/go/support/apiv2/supportpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud cloudsupport command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudsupport",
		Usage: "manage Google Cloud Support API resources",
		Commands: []*cli.Command{
			{
				Name:  "attachments",
				Usage: "Manage attachments resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							client, err := support.NewCaseAttachmentClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &supportpb.ListAttachmentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAttachments(ctx, req)
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
				Name:  "case-classifications",
				Usage: "Manage case-classifications resources",
				Commands: []*cli.Command{
					{
						Name:  "search",
						Usage: "search case-classifications",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search...")
							return nil
						},
					},
				},
			},
			{
				Name:  "cases",
				Usage: "Manage cases resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "case", Usage: "The case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/cases/%s", cmd.String("organization"), cmd.String("case"))
							client, err := support.NewCaseClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &supportpb.GetCaseRequest{Name: name}
							resp, err := client.GetCase(ctx, req)
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
						Usage: "list cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							client, err := support.NewCaseClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &supportpb.ListCasesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListCases(ctx, req)
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
						Name:  "search",
						Usage: "search cases",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
							&cli.StringFlag{Name: "contact-email", Usage: "The contact email.", Required: false},
							&cli.BoolFlag{Name: "escalated", Usage: "The escalated.", Required: false},
							&cli.BoolFlag{Name: "test-case", Usage: "The test case.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							client, err := support.NewCaseClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &supportpb.CreateCaseRequest{Parent: parent}
							req.Case = &supportpb.Case{
								DisplayName:  cmd.String("display-name"),
								Description:  cmd.String("description"),
								TimeZone:     cmd.String("time-zone"),
								ContactEmail: cmd.String("contact-email"),
								Escalated:    cmd.Bool("escalated"),
								TestCase:     cmd.Bool("test-case"),
								LanguageCode: cmd.String("language-code"),
							}
							resp, err := client.CreateCase(ctx, req)
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
						Usage: "update cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "case", Usage: "The case.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "time-zone", Usage: "The time zone.", Required: false},
							&cli.StringFlag{Name: "contact-email", Usage: "The contact email.", Required: false},
							&cli.BoolFlag{Name: "escalated", Usage: "The escalated.", Required: false},
							&cli.BoolFlag{Name: "test-case", Usage: "The test case.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/cases/%s", cmd.String("organization"), cmd.String("case"))
							client, err := support.NewCaseClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &supportpb.UpdateCaseRequest{}
							req.Case = &supportpb.Case{
								Name:         name,
								DisplayName:  cmd.String("display-name"),
								Description:  cmd.String("description"),
								TimeZone:     cmd.String("time-zone"),
								ContactEmail: cmd.String("contact-email"),
								Escalated:    cmd.Bool("escalated"),
								TestCase:     cmd.Bool("test-case"),
								LanguageCode: cmd.String("language-code"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("time-zone") {
								paths = append(paths, "time_zone")
							}
							if cmd.IsSet("contact-email") {
								paths = append(paths, "contact_email")
							}
							if cmd.IsSet("escalated") {
								paths = append(paths, "escalated")
							}
							if cmd.IsSet("test-case") {
								paths = append(paths, "test_case")
							}
							if cmd.IsSet("language-code") {
								paths = append(paths, "language_code")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateCase(ctx, req)
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
						Name:  "escalate",
						Usage: "escalate cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "case", Usage: "The case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/cases/%s", cmd.String("organization"), cmd.String("case"))
							fmt.Printf("Executing escalate on %s\n", name)
							return nil
						},
					},
					{
						Name:  "close",
						Usage: "close cases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "case", Usage: "The case.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/cases/%s", cmd.String("organization"), cmd.String("case"))
							fmt.Printf("Executing close on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "comments",
				Usage: "Manage comments resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							client, err := support.NewCommentClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &supportpb.ListCommentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListComments(ctx, req)
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
						Usage: "create comments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "case", Usage: "The case.", Required: true},
							&cli.StringFlag{Name: "body", Usage: "The body.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/cases/%s", cmd.String("organization"), cmd.String("case"))
							client, err := support.NewCommentClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &supportpb.CreateCommentRequest{Parent: parent}
							req.Comment = &supportpb.Comment{
								Body: cmd.String("body"),
							}
							resp, err := client.CreateComment(ctx, req)
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
		},
	}
}
