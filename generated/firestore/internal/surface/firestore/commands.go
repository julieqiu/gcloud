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

package firestore

import (
	firestore "cloud.google.com/go/firestore/apiv1"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the firestore command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "firestore",
		Usage: "manage Cloud Firestore API resources",
		Commands: []*cli.Command{
			{
				Name:  "documents",
				Usage: "Manage documents resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/documents/%s", cmd.String("project"), cmd.String("database"), cmd.String("document"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := firestore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &firestorepb.GetDocumentRequest{
								Name: name,
							}

							resp, err := client.GetDocument(ctx, req)
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
						Usage: "list documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection_id", Usage: "The ID of the collection_id.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The optional ordering of the documents to return.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of documents to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDocuments` response.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-missing", Usage: "If the list should show missing documents.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s/documents/%s/%s", cmd.String("project"), cmd.String("database"), cmd.String("document"), cmd.String("collection_id"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "document.name" not yet supported.
							document_name := fmt.Sprintf("projects/%s/databases/%s/documents/%s", cmd.String("project"), cmd.String("database"), cmd.String("document"))
							fmt.Printf("Executing update on %s\n", document_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/documents/%s", cmd.String("project"), cmd.String("database"), cmd.String("document"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDocument on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := firestore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &firestorepb.DeleteDocumentRequest{
								Name: name,
							}

							if err := client.DeleteDocument(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "batch-get",
						Usage: "batch-get documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringSliceFlag{Name: "documents", Usage: "The names of the documents to retrieve.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing batch-get on %s\n", database)
							return nil
						},
					},

					{
						Name:  "begin-transaction",
						Usage: "begin-transaction documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing begin-transaction on %s\n", database)
							return nil
						},
					},

					{
						Name:  "commit",
						Usage: "commit documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transaction", Usage: "If set, applies all writes in this transaction, and commits it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing commit on %s\n", database)
							return nil
						},
					},

					{
						Name:  "rollback",
						Usage: "rollback documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transaction", Usage: "The transaction to roll back.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing rollback on %s\n", database)
							return nil
						},
					},

					{
						Name:  "run-query",
						Usage: "run-query documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s/documents", cmd.String("project"), cmd.String("database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := firestore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &firestorepb.RunQueryRequest{
								Parent: parent,
							}

							resp, err := client.RunQuery(ctx, req)
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
						Name:  "execute-pipeline",
						Usage: "execute-pipeline documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing execute-pipeline on %s\n", database)
							return nil
						},
					},

					{
						Name:  "run-aggregation-query",
						Usage: "run-aggregation-query documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s/documents", cmd.String("project"), cmd.String("database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := firestore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &firestorepb.RunAggregationQueryRequest{
								Parent: parent,
							}

							resp, err := client.RunAggregationQuery(ctx, req)
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
						Name:  "partition-query",
						Usage: "partition-query documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of partitions to return in this call, subject to.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `next_page_token` value returned from a previous call to.", Required: false},
							&cli.IntFlag{Name: "partition-count", Usage: "The desired maximum number of partition points.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s/documents", cmd.String("project"), cmd.String("database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := firestore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &firestorepb.PartitionQueryRequest{
								Parent:         parent,
								PartitionCount: cmd.Int("partition-count"),
								PageToken:      cmd.String("page-token"),
								PageSize:       int32(cmd.Int("page-size")),
							}

							limit := cmd.Int("limit")
							it := client.PartitionQuery(ctx, req)
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
						Name:  "write",
						Usage: "write documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "stream-id", Usage: "The ID of the write stream to resume.", Required: false},
							&cli.StringFlag{Name: "stream-token", Usage: "A stream token that was previously sent by the server.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing write on %s\n", database)
							return nil
						},
					},

					{
						Name:  "listen",
						Usage: "listen documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing listen on %s\n", database)
							return nil
						},
					},

					{
						Name:  "list-collection-ids",
						Usage: "list-collection-ids documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s/documents", cmd.String("project"), cmd.String("database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := firestore.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &firestorepb.ListCollectionIdsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							resp, err := client.ListCollectionIds(ctx, req)
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
						Name:  "batch-write",
						Usage: "batch-write documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing batch-write on %s\n", database)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "collection_id", Usage: "The ID of the collection_id.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "document", Usage: "The ID of the document.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/databases/%s/documents/%s/%s", cmd.String("project"), cmd.String("database"), cmd.String("document"), cmd.String("collection_id"))
							fmt.Printf("Executing create on %s\n", parent)
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
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s", cmd.String("project"), cmd.String("database"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/operations/%s", cmd.String("project"), cmd.String("database"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/operations/%s", cmd.String("project"), cmd.String("database"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/databases/%s/operations/%s", cmd.String("project"), cmd.String("database"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
