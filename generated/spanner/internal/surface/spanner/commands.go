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

package spanner

import (
	spanner "cloud.google.com/go/spanner/apiv1"
	"cloud.google.com/go/spanner/apiv1/spannerpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the spanner command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "spanner",
		Usage: "manage Cloud Spanner API resources",
		Commands: []*cli.Command{
			{
				Name:  "databases",
				Usage: "Manage databases resources",
				Commands: []*cli.Command{

					{
						Name:  "cache-update",
						Usage: "cache-update databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "max-range-count", Usage: "The maximum number of ranges to return in the response.", Required: false},
							&cli.IntFlag{Name: "max-recipe-count", Usage: "The maximum number of key recipes to return in the response.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing cache-update on %s\n", database)
							return nil
						},
					},
				},
			},
			{
				Name:  "sessions",
				Usage: "Manage sessions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing create on %s\n", database)
							return nil
						},
					},

					{
						Name:  "batch-create",
						Usage: "batch-create sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "session-count", Usage: "The number of sessions to be created in this batch call.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing batch-create on %s\n", database)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := spanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &spannerpb.GetSessionRequest{
								Name: name,
							}

							resp, err := client.GetSession(ctx, req)
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
						Usage: "list sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of sessions to be returned in the response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "If non-empty, `page_token` should contain a.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							database := fmt.Sprintf("projects/%s/instances/%s/databases/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"))
							fmt.Printf("Executing list on %s\n", database)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSession on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := spanner.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &spannerpb.DeleteSessionRequest{
								Name: name,
							}

							if err := client.DeleteSession(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "execute-sql",
						Usage: "execute-sql sessions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "data-boost-enabled", Usage: "If this is for a partitioned query and this field is set to `true`, the.", Required: false},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.BoolFlag{Name: "last-statement", Usage: "If set to `true`, this statement marks the end of the.", Required: false},
							&cli.StringFlag{Name: "partition-token", Usage: "If present, results are restricted to the specified partition.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-mode", Usage: "Used to control the amount of debugging information returned in.", Required: false},
							&cli.StringFlag{Name: "resume-token", Usage: "If this request is resuming a previously interrupted SQL statement.", Required: false},
							&cli.IntFlag{Name: "seqno", Usage: "A per-transaction sequence number used to identify this request.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "sql", Usage: "The SQL string.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing execute-sql on %s\n", session)
							return nil
						},
					},

					{
						Name:  "execute-streaming-sql",
						Usage: "execute-streaming-sql sessions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "data-boost-enabled", Usage: "If this is for a partitioned query and this field is set to `true`, the.", Required: false},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.BoolFlag{Name: "last-statement", Usage: "If set to `true`, this statement marks the end of the.", Required: false},
							&cli.StringFlag{Name: "partition-token", Usage: "If present, results are restricted to the specified partition.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-mode", Usage: "Used to control the amount of debugging information returned in.", Required: false},
							&cli.StringFlag{Name: "resume-token", Usage: "If this request is resuming a previously interrupted SQL statement.", Required: false},
							&cli.IntFlag{Name: "seqno", Usage: "A per-transaction sequence number used to identify this request.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "sql", Usage: "The SQL string.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing execute-streaming-sql on %s\n", session)
							return nil
						},
					},

					{
						Name:  "execute-batch-dml",
						Usage: "execute-batch-dml sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.BoolFlag{Name: "last-statements", Usage: "If set to `true`, this request marks the end of the transaction.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "seqno", Usage: "A per-transaction sequence number used to identify this request.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing execute-batch-dml on %s\n", session)
							return nil
						},
					},

					{
						Name:  "read",
						Usage: "read sessions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "columns", Usage: "The columns of [table][google.", Required: true},
							&cli.BoolFlag{Name: "data-boost-enabled", Usage: "If this is for a partitioned read and this field is set to `true`, the.", Required: false},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "index", Usage: "If non-empty, the name of an index on.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "If greater than zero, only the first `limit` rows are yielded.", Required: false},
							&cli.StringFlag{Name: "lock-hint", Usage: "Lock Hint for the request, it can only be used with read-write.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Order for the returned rows.", Required: false},
							&cli.StringFlag{Name: "partition-token", Usage: "If present, results are restricted to the specified partition.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resume-token", Usage: "If this request is resuming a previously interrupted read,.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The name of the table in the database to be read.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing read on %s\n", session)
							return nil
						},
					},

					{
						Name:  "streaming-read",
						Usage: "streaming-read sessions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "columns", Usage: "The columns of [table][google.", Required: true},
							&cli.BoolFlag{Name: "data-boost-enabled", Usage: "If this is for a partitioned read and this field is set to `true`, the.", Required: false},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "index", Usage: "If non-empty, the name of an index on.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "If greater than zero, only the first `limit` rows are yielded.", Required: false},
							&cli.StringFlag{Name: "lock-hint", Usage: "Lock Hint for the request, it can only be used with read-write.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Order for the returned rows.", Required: false},
							&cli.StringFlag{Name: "partition-token", Usage: "If present, results are restricted to the specified partition.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resume-token", Usage: "If this request is resuming a previously interrupted read,.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The name of the table in the database to be read.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing streaming-read on %s\n", session)
							return nil
						},
					},

					{
						Name:  "begin-transaction",
						Usage: "begin-transaction sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing begin-transaction on %s\n", session)
							return nil
						},
					},

					{
						Name:  "commit",
						Usage: "commit sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-commit-stats", Usage: "If `true`, then statistics related to the transaction is included in.", Required: false},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing commit on %s\n", session)
							return nil
						},
					},

					{
						Name:  "rollback",
						Usage: "rollback sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "transaction-id", Usage: "The transaction to roll back.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing rollback on %s\n", session)
							return nil
						},
					},

					{
						Name:  "partition-query",
						Usage: "partition-query sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "sql", Usage: "The query request to generate partitions for.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing partition-query on %s\n", session)
							return nil
						},
					},

					{
						Name:  "partition-read",
						Usage: "partition-read sessions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "columns", Usage: "The columns of [table][google.", Required: false},
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "index", Usage: "If non-empty, the name of an index on.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
							&cli.StringFlag{Name: "table", Usage: "The name of the table in the database to be read.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing partition-read on %s\n", session)
							return nil
						},
					},

					{
						Name:  "batch-write",
						Usage: "batch-write sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.BoolFlag{Name: "exclude-txn-from-change-streams", Usage: "If you don't set the `exclude_txn_from_change_streams` option or.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "session", Usage: "The ID of the session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							session := fmt.Sprintf("projects/%s/instances/%s/databases/%s/sessions/%s", cmd.String("project"), cmd.String("instance"), cmd.String("database"), cmd.String("session"))
							fmt.Printf("Executing batch-write on %s\n", session)
							return nil
						},
					},
				},
			},
		},
	}
}
