package analyticshub

import (
	"context"
	"fmt"
	"strings"

	analyticshub "cloud.google.com/go/bigquery/analyticshub/apiv1"
	"cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud analyticshub command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "analyticshub",
		Usage: "manage Analytics Hub API resources",
		Commands: []*cli.Command{
			{
				Name:  "data-exchanges",
				Usage: "Manage data-exchanges resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &analyticshubpb.ListDataExchangesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDataExchanges(ctx, req)
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
						Name:  "list",
						Usage: "list data-exchanges",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.GetDataExchangeRequest{Name: name}
							resp, err := client.GetDataExchange(ctx, req)
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
						Usage: "create data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data-exchange-id", Usage: "The data exchange id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "primary-contact", Usage: "The primary contact.", Required: false},
							&cli.StringFlag{Name: "documentation", Usage: "The documentation.", Required: false},
							&cli.BoolFlag{Name: "log-linked-dataset-query-user-email", Usage: "The log linked dataset query user email.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.CreateDataExchangeRequest{Parent: parent}
							req.DataExchangeId = cmd.String("data-exchange-id")
							req.DataExchange = &analyticshubpb.DataExchange{
								DisplayName:                    cmd.String("display-name"),
								Description:                    cmd.String("description"),
								PrimaryContact:                 cmd.String("primary-contact"),
								Documentation:                  cmd.String("documentation"),
								LogLinkedDatasetQueryUserEmail: cmd.Bool("log-linked-dataset-query-user-email"),
							}
							resp, err := client.CreateDataExchange(ctx, req)
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
						Usage: "update data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "primary-contact", Usage: "The primary contact.", Required: false},
							&cli.StringFlag{Name: "documentation", Usage: "The documentation.", Required: false},
							&cli.BoolFlag{Name: "log-linked-dataset-query-user-email", Usage: "The log linked dataset query user email.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.UpdateDataExchangeRequest{}
							req.DataExchange = &analyticshubpb.DataExchange{
								Name:                           name,
								DisplayName:                    cmd.String("display-name"),
								Description:                    cmd.String("description"),
								PrimaryContact:                 cmd.String("primary-contact"),
								Documentation:                  cmd.String("documentation"),
								LogLinkedDatasetQueryUserEmail: cmd.Bool("log-linked-dataset-query-user-email"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("primary-contact") {
								paths = append(paths, "primary_contact")
							}
							if cmd.IsSet("documentation") {
								paths = append(paths, "documentation")
							}
							if cmd.IsSet("log-linked-dataset-query-user-email") {
								paths = append(paths, "log_linked_dataset_query_user_email")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateDataExchange(ctx, req)
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
						Usage: "delete data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.DeleteDataExchangeRequest{Name: name}
							if err := client.DeleteDataExchange(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "subscribe",
						Usage: "subscribe data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "destination", Usage: "The destination.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The subscription.", Required: true},
							&cli.StringFlag{Name: "subscriber-contact", Usage: "The subscriber contact.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.SubscribeDataExchangeRequest{Name: name}
							req.Destination = cmd.String("destination")
							req.Subscription = cmd.String("subscription")
							req.SubscriberContact = cmd.String("subscriber-contact")
							op, err := client.SubscribeDataExchange(ctx, req)
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
						Usage: "list data-exchanges",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataExchange", Usage: "The dataExchange.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataExchange"))
							client, err := analyticshub.NewClient(ctx)
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
						Usage: "set-iam-policy data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataExchange", Usage: "The dataExchange.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataExchange"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions data-exchanges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataExchange", Usage: "The dataExchange.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataExchange"))
							client, err := analyticshub.NewClient(ctx)
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
				Name:  "listings",
				Usage: "Manage listings resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &analyticshubpb.ListListingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListListings(ctx, req)
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
						Usage: "describe listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The listing.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("listing"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.GetListingRequest{Name: name}
							resp, err := client.GetListing(ctx, req)
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
						Usage: "create listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "listing-id", Usage: "The listing id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "primary-contact", Usage: "The primary contact.", Required: false},
							&cli.StringFlag{Name: "documentation", Usage: "The documentation.", Required: false},
							&cli.StringFlag{Name: "request-access", Usage: "The request access.", Required: false},
							&cli.BoolFlag{Name: "log-linked-dataset-query-user-email", Usage: "The log linked dataset query user email.", Required: false},
							&cli.BoolFlag{Name: "allow-only-metadata-sharing", Usage: "The allow only metadata sharing.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.CreateListingRequest{Parent: parent}
							req.ListingId = cmd.String("listing-id")
							req.Listing = &analyticshubpb.Listing{
								DisplayName:                    cmd.String("display-name"),
								Description:                    cmd.String("description"),
								PrimaryContact:                 cmd.String("primary-contact"),
								Documentation:                  cmd.String("documentation"),
								RequestAccess:                  cmd.String("request-access"),
								LogLinkedDatasetQueryUserEmail: cmd.Bool("log-linked-dataset-query-user-email"),
								AllowOnlyMetadataSharing:       cmd.Bool("allow-only-metadata-sharing"),
							}
							resp, err := client.CreateListing(ctx, req)
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
						Usage: "update listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The listing.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "primary-contact", Usage: "The primary contact.", Required: false},
							&cli.StringFlag{Name: "documentation", Usage: "The documentation.", Required: false},
							&cli.StringFlag{Name: "request-access", Usage: "The request access.", Required: false},
							&cli.BoolFlag{Name: "log-linked-dataset-query-user-email", Usage: "The log linked dataset query user email.", Required: false},
							&cli.BoolFlag{Name: "allow-only-metadata-sharing", Usage: "The allow only metadata sharing.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("listing"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.UpdateListingRequest{}
							req.Listing = &analyticshubpb.Listing{
								Name:                           name,
								DisplayName:                    cmd.String("display-name"),
								Description:                    cmd.String("description"),
								PrimaryContact:                 cmd.String("primary-contact"),
								Documentation:                  cmd.String("documentation"),
								RequestAccess:                  cmd.String("request-access"),
								LogLinkedDatasetQueryUserEmail: cmd.Bool("log-linked-dataset-query-user-email"),
								AllowOnlyMetadataSharing:       cmd.Bool("allow-only-metadata-sharing"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("primary-contact") {
								paths = append(paths, "primary_contact")
							}
							if cmd.IsSet("documentation") {
								paths = append(paths, "documentation")
							}
							if cmd.IsSet("request-access") {
								paths = append(paths, "request_access")
							}
							if cmd.IsSet("log-linked-dataset-query-user-email") {
								paths = append(paths, "log_linked_dataset_query_user_email")
							}
							if cmd.IsSet("allow-only-metadata-sharing") {
								paths = append(paths, "allow_only_metadata_sharing")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateListing(ctx, req)
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
						Usage: "delete listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The listing.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("listing"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.DeleteListingRequest{Name: name}
							if err := client.DeleteListing(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "subscribe",
						Usage: "subscribe listings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "listing", Usage: "The listing.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/listings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("listing"))
							fmt.Printf("Executing subscribe on %s\n", name)
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
						Name:  "describe",
						Usage: "describe operations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
				},
			},
			{
				Name:  "query-templates",
				Usage: "Manage query-templates resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "query-template-id", Usage: "The query template id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "proposer", Usage: "The proposer.", Required: false},
							&cli.StringFlag{Name: "primary-contact", Usage: "The primary contact.", Required: false},
							&cli.StringFlag{Name: "documentation", Usage: "The documentation.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.CreateQueryTemplateRequest{Parent: parent}
							req.QueryTemplateId = cmd.String("query-template-id")
							req.QueryTemplate = &analyticshubpb.QueryTemplate{
								DisplayName:    cmd.String("display-name"),
								Description:    cmd.String("description"),
								Proposer:       cmd.String("proposer"),
								PrimaryContact: cmd.String("primary-contact"),
								Documentation:  cmd.String("documentation"),
							}
							resp, err := client.CreateQueryTemplate(ctx, req)
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
						Usage: "describe query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "query_template", Usage: "The query_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("query_template"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.GetQueryTemplateRequest{Name: name}
							resp, err := client.GetQueryTemplate(ctx, req)
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
						Usage: "list query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &analyticshubpb.ListQueryTemplatesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListQueryTemplates(ctx, req)
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
						Usage: "update query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "query_template", Usage: "The query_template.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "proposer", Usage: "The proposer.", Required: false},
							&cli.StringFlag{Name: "primary-contact", Usage: "The primary contact.", Required: false},
							&cli.StringFlag{Name: "documentation", Usage: "The documentation.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("query_template"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.UpdateQueryTemplateRequest{}
							req.QueryTemplate = &analyticshubpb.QueryTemplate{
								Name:           name,
								DisplayName:    cmd.String("display-name"),
								Description:    cmd.String("description"),
								Proposer:       cmd.String("proposer"),
								PrimaryContact: cmd.String("primary-contact"),
								Documentation:  cmd.String("documentation"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("proposer") {
								paths = append(paths, "proposer")
							}
							if cmd.IsSet("primary-contact") {
								paths = append(paths, "primary_contact")
							}
							if cmd.IsSet("documentation") {
								paths = append(paths, "documentation")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateQueryTemplate(ctx, req)
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
						Usage: "delete query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "query_template", Usage: "The query_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("query_template"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.DeleteQueryTemplateRequest{Name: name}
							if err := client.DeleteQueryTemplate(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "submit",
						Usage: "submit query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "query_template", Usage: "The query_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("query_template"))
							fmt.Printf("Executing submit on %s\n", name)
							return nil
						},
					},
					{
						Name:  "approve",
						Usage: "approve query-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_exchange", Usage: "The data_exchange.", Required: true},
							&cli.StringFlag{Name: "query_template", Usage: "The query_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s/queryTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_exchange"), cmd.String("query_template"))
							fmt.Printf("Executing approve on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "subscriptions",
				Usage: "Manage subscriptions resources",
				Commands: []*cli.Command{
					{
						Name:  "refresh",
						Usage: "refresh subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.RefreshSubscriptionRequest{Name: name}
							op, err := client.RefreshSubscription(ctx, req)
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
						Usage: "describe subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.GetSubscriptionRequest{Name: name}
							resp, err := client.GetSubscription(ctx, req)
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
						Usage: "list subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &analyticshubpb.ListSubscriptionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListSubscriptions(ctx, req)
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
						Name:  "revoke",
						Usage: "revoke subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							fmt.Printf("Executing revoke on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/subscriptions/%s", cmd.String("project"), cmd.String("location"), cmd.String("subscription"))
							client, err := analyticshub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &analyticshubpb.DeleteSubscriptionRequest{Name: name}
							op, err := client.DeleteSubscription(ctx, req)
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
		},
	}
}
