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

package pubsub

import (
	pubsub "cloud.google.com/go/pubsub/apiv1"
	"cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the pubsub command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "pubsub",
		Usage: "manage Cloud Pub/Sub API resources",
		Commands: []*cli.Command{
			{
				Name:  "schemas",
				Usage: "Manage schemas resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-id", Usage: "The ID to use for the schema, which will become the final component of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.CreateSchemaRequest{
								Parent:   parent,
								SchemaId: cmd.String("schema-id"),
							}

							resp, err := client.CreateSchema(ctx, req)
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
						Usage: "describe schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The set of fields to return in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.GetSchemaRequest{
								Name: name,
								View: pubsubpb.SchemaView(pubsubpb.SchemaView_value[cmd.String("view")]),
							}

							resp, err := client.GetSchema(ctx, req)
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
						Usage: "list schemas",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of schemas to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListSchemasResponse`; indicates that.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The set of Schema fields to return in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.ListSchemasRequest{
								Parent:    parent,
								View:      pubsubpb.SchemaView(pubsubpb.SchemaView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSchemas(ctx, req)
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
						Name:  "list",
						Usage: "list schemas",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of revisions to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token, received from a previous ListSchemaRevisions call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The set of Schema fields to return in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.ListSchemaRevisionsRequest{
								Name:      name,
								View:      pubsubpb.SchemaView(pubsubpb.SchemaView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSchemaRevisions(ctx, req)
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
						Name:  "commit",
						Usage: "commit schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.CommitSchemaRequest{
								Name: name,
							}

							resp, err := client.CommitSchema(ctx, req)
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
						Name:  "rollback",
						Usage: "rollback schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision-id", Usage: "The revision ID to roll back to.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.RollbackSchemaRequest{
								Name:       name,
								RevisionId: cmd.String("revision-id"),
							}

							resp, err := client.RollbackSchema(ctx, req)
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
						Usage: "delete schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision-id", Usage: "This field is deprecated and should not be used for specifying.", Required: false},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.DeleteSchemaRevisionRequest{
								Name:       name,
								RevisionId: cmd.String("revision-id"),
							}

							resp, err := client.DeleteSchemaRevision(ctx, req)
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
						Usage: "delete schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/schemas/%s", cmd.String("project"), cmd.String("schema"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSchema on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.DeleteSchemaRequest{
								Name: name,
							}

							if err := client.DeleteSchema(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "validate",
						Usage: "validate schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.ValidateSchemaRequest{
								Parent: parent,
							}

							resp, err := client.ValidateSchema(ctx, req)
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
						Name:  "validate-message",
						Usage: "validate-message schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encoding", Usage: "The encoding expected for messages.", Required: false},
							&cli.StringFlag{Name: "message", Usage: "Message to validate against the provided `schema_spec`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.ValidateMessageRequest{
								Parent:   parent,
								Message:  []byte(cmd.String("message")),
								Encoding: pubsubpb.Encoding(pubsubpb.Encoding_value[cmd.String("encoding")]),
							}

							resp, err := client.ValidateMessage(ctx, req)
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
				Name:  "snapshots",
				Usage: "Manage snapshots resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of snapshot names to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListTopicSnapshotsResponse`;.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							topic := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing list on %s\n", topic)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							snapshot := fmt.Sprintf("projects/%s/snapshots/%s", cmd.String("project"), cmd.String("snapshot"))
							fmt.Printf("Executing describe on %s\n", snapshot)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of snapshots to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListSnapshotsResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "create-snapshot",
						Usage: "create-snapshot snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The subscription whose backlog the snapshot retains.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/snapshots/%s", cmd.String("project"), cmd.String("snapshot"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.CreateSnapshotRequest{
								Name:         name,
								Subscription: cmd.String("subscription"),
							}

							resp, err := client.CreateSnapshot(ctx, req)
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
						Usage: "update snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "snapshot.name" not yet supported.
							snapshot_name := fmt.Sprintf("projects/%s/snapshots/%s", cmd.String("project"), cmd.String("snapshot"))
							fmt.Printf("Executing update on %s\n", snapshot_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							snapshot := fmt.Sprintf("projects/%s/snapshots/%s", cmd.String("project"), cmd.String("snapshot"))
							fmt.Printf("Executing delete on %s\n", snapshot)
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
						Name:  "list",
						Usage: "list subscriptions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of subscription names to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListTopicSubscriptionsResponse`;.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							topic := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing list on %s\n", topic)
							return nil
						},
					},

					{
						Name:  "detach",
						Usage: "detach subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing detach on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions subscriptions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "create-subscription",
						Usage: "create-subscription subscriptions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "ack-deadline-seconds", Usage: "The approximate amount of time (on a best-effort basis) Pub/Sub.", Required: false},
							&cli.BoolFlag{Name: "detached", Usage: "Indicates whether the subscription is detached from its topic.", Required: false},
							&cli.BoolFlag{Name: "enable-exactly-once-delivery", Usage: "If true, Pub/Sub provides the following guarantees for the.", Required: false},
							&cli.BoolFlag{Name: "enable-message-ordering", Usage: "If true, messages published with the same `ordering_key` in.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "An expression written in the Pub/Sub [filter.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "retain-acked-messages", Usage: "Indicates whether to retain acknowledged messages.", Required: false},
							&cli.StringFlag{Name: "state", Usage: "An output-only field indicating whether or not the.", Required: false},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The name of the topic from which this subscription is receiving.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.Subscription{
								Name:                      name,
								Topic:                     cmd.String("topic"),
								AckDeadlineSeconds:        int32(cmd.Int("ack-deadline-seconds")),
								RetainAckedMessages:       cmd.Bool("retain-acked-messages"),
								EnableMessageOrdering:     cmd.Bool("enable-message-ordering"),
								Filter:                    cmd.String("filter"),
								Detached:                  cmd.Bool("detached"),
								EnableExactlyOnceDelivery: cmd.Bool("enable-exactly-once-delivery"),
								State:                     pubsubpb.Subscription_State(pubsubpb.Subscription_State_value[cmd.String("state")]),
							}

							resp, err := client.CreateSubscription(ctx, req)
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
						Usage: "describe subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing describe on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "subscription.name" not yet supported.
							subscription_name := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing update on %s\n", subscription_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list subscriptions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of subscriptions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListSubscriptionsResponse`;.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing delete on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "modify-ack-deadline",
						Usage: "modify-ack-deadline subscriptions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "ack-deadline-seconds", Usage: "The new ack deadline with respect to the time this request was.", Required: true},
							&cli.StringSliceFlag{Name: "ack-ids", Usage: "List of acknowledgment IDs.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing modify-ack-deadline on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "acknowledge",
						Usage: "acknowledge subscriptions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "ack-ids", Usage: "The acknowledgment ID for the messages being acknowledged that.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing acknowledge on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "pull",
						Usage: "pull subscriptions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "max-messages", Usage: "The maximum number of messages to return for this request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-immediately", Usage: "If this field set to true, the system will respond immediately.", Required: false},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing pull on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "modify-push-config",
						Usage: "modify-push-config subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing modify-push-config on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "seek",
						Usage: "seek subscriptions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							subscription := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing seek on %s\n", subscription)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions subscriptions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions subscriptions",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subscription", Usage: "The ID of the subscription.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/subscriptions/%s", cmd.String("project"), cmd.String("subscription"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "topics",
				Usage: "Manage topics resources",
				Commands: []*cli.Command{

					{
						Name:  "create-topic",
						Usage: "create-topic topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "kms-key-name", Usage: "The resource name of the Cloud KMS CryptoKey to be used to.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "satisfies-pzs", Usage: "Reserved for future use.", Required: false},
							&cli.StringFlag{Name: "state", Usage: "An output-only field indicating the state of the topic.", Required: false},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := pubsub.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &pubsubpb.Topic{
								Name:         name,
								KmsKeyName:   cmd.String("kms-key-name"),
								SatisfiesPzs: cmd.Bool("satisfies-pzs"),
								State:        pubsubpb.Topic_State(pubsubpb.Topic_State_value[cmd.String("state")]),
							}

							resp, err := client.CreateTopic(ctx, req)
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
						Usage: "update topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "topic.name" not yet supported.
							topic_name := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing update on %s\n", topic_name)
							return nil
						},
					},

					{
						Name:  "publish",
						Usage: "publish topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							topic := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing publish on %s\n", topic)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							topic := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing describe on %s\n", topic)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list topics",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of topics to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListTopicsResponse`; indicates.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							topic := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing delete on %s\n", topic)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe topics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "topic", Usage: "The ID of the topic.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/topics/%s", cmd.String("project"), cmd.String("topic"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},
				},
			},
		},
	}
}
