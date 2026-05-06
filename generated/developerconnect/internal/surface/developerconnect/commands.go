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

package developerconnect

import (
	developerconnect "cloud.google.com/go/developerconnect/apiv1"
	"cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the developerconnect command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "developerconnect",
		Usage: "manage Developer Connect API resources",
		Commands: []*cli.Command{
			{
				Name:  "account-connectors",
				Usage: "Manage account-connectors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list account-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.ListAccountConnectorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAccountConnectors(ctx, req)
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
						Usage: "describe account-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.GetAccountConnectorRequest{
								Name: name,
							}

							resp, err := client.GetAccountConnector(ctx, req)
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
						Name:  "create",
						Usage: "create account-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector-id", Usage: "The ID to use for the AccountConnector, which will become the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.CreateAccountConnectorRequest{
								Parent:             parent,
								AccountConnectorId: cmd.String("account-connector-id"),
								RequestId:          cmd.String("request-id"),
								ValidateOnly:       cmd.Bool("validate-only"),
							}

							op, err := client.CreateAccountConnector(ctx, req)
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

					{
						Name:  "update",
						Usage: "update account-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the accountConnector is not found a new.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "account_connector.name" not yet supported.
							account_connector_name := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							fmt.Printf("Executing update on %s\n", account_connector_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete account-connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the AccountConnectorn.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any Users from this AccountConnector will also.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAccountConnector %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.DeleteAccountConnectorRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
								Force:        cmd.Bool("force"),
							}

							op, err := client.DeleteAccountConnector(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "connections",
				Usage: "Manage connections resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.ListConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListConnections(ctx, req)
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
						Usage: "describe connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.GetConnectionRequest{
								Name: name,
							}

							resp, err := client.GetConnection(ctx, req)
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
						Name:  "create",
						Usage: "create connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.CreateConnectionRequest{
								Parent:       parent,
								ConnectionId: cmd.String("connection-id"),
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateConnection(ctx, req)
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

					{
						Name:  "update",
						Usage: "update connections",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the connection is not found a new connection.", Required: false},
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "connection.name" not yet supported.
							connection_name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing update on %s\n", connection_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The current etag of the Connection.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteConnection %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.DeleteConnectionRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteConnection(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "fetch-linkable-git-repositories",
						Usage: "fetch-linkable-git-repositories connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results to return in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page start.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							connection := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing fetch-linkable-git-repositories on %s\n", connection)
							return nil
						},
					},

					{
						Name:  "fetch-git-hub-installations",
						Usage: "fetch-git-hub-installations connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							connection := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing fetch-git-hub-installations on %s\n", connection)
							return nil
						},
					},
				},
			},
			{
				Name:  "deployment-events",
				Usage: "Manage deployment-events resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe deployment-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-event", Usage: "The ID of the deployment event.", Required: true},
							&cli.StringFlag{Name: "insights-config", Usage: "The ID of the insights config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/insightsConfigs/%s/deploymentEvents/%s", cmd.String("project"), cmd.String("location"), cmd.String("insights-config"), cmd.String("deployment-event"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.GetDeploymentEventRequest{
								Name: name,
							}

							resp, err := client.GetDeploymentEvent(ctx, req)
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
						Usage: "list deployment-events",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that matches a subset of the DeploymentEvents.", Required: false},
							&cli.StringFlag{Name: "insights-config", Usage: "The ID of the insights config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of deployment events to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDeploymentEvents`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/insightsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("insights-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.ListDeploymentEventsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeploymentEvents(ctx, req)
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
				Name:  "git-repository-links",
				Usage: "Manage git-repository-links resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "git-repository-link-id", Usage: "The ID to use for the repository, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.CreateGitRepositoryLinkRequest{
								Parent:              parent,
								GitRepositoryLinkId: cmd.String("git-repository-link-id"),
								RequestId:           cmd.String("request-id"),
								ValidateOnly:        cmd.Bool("validate-only"),
							}

							op, err := client.CreateGitRepositoryLink(ctx, req)
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

					{
						Name:  "delete",
						Usage: "delete git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.StringFlag{Name: "git-repository-link", Usage: "The ID of the git repository link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("git-repository-link"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGitRepositoryLink %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.DeleteGitRepositoryLinkRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteGitRepositoryLink(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.ListGitRepositoryLinksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGitRepositoryLinks(ctx, req)
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
						Usage: "describe git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "git-repository-link", Usage: "The ID of the git repository link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("git-repository-link"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.GetGitRepositoryLinkRequest{
								Name: name,
							}

							resp, err := client.GetGitRepositoryLink(ctx, req)
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
						Name:  "fetch-read-write-token",
						Usage: "fetch-read-write-token git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "git-repository-link", Usage: "The ID of the git repository link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							git_repository_link := fmt.Sprintf("projects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("git-repository-link"))
							fmt.Printf("Executing fetch-read-write-token on %s\n", git_repository_link)
							return nil
						},
					},

					{
						Name:  "fetch-read-token",
						Usage: "fetch-read-token git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "git-repository-link", Usage: "The ID of the git repository link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							git_repository_link := fmt.Sprintf("projects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("git-repository-link"))
							fmt.Printf("Executing fetch-read-token on %s\n", git_repository_link)
							return nil
						},
					},

					{
						Name:  "fetch-git-refs",
						Usage: "fetch-git-refs git-repository-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "connection", Usage: "The ID of the connection.", Required: true},
							&cli.StringFlag{Name: "git-repository-link", Usage: "The ID of the git repository link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results to return in the list.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page start.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "ref-type", Usage: "Type of refs to fetch.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							git_repository_link := fmt.Sprintf("projects/%s/locations/%s/connections/%s/gitRepositoryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"), cmd.String("git-repository-link"))
							fmt.Printf("Executing fetch-git-refs on %s\n", git_repository_link)
							return nil
						},
					},
				},
			},
			{
				Name:  "insights-configs",
				Usage: "Manage insights-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list insights-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.ListInsightsConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInsightsConfigs(ctx, req)
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
						Name:  "create",
						Usage: "create insights-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "insights-config-id", Usage: "ID of the requesting InsightsConfig.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.CreateInsightsConfigRequest{
								Parent:           parent,
								InsightsConfigId: cmd.String("insights-config-id"),
								ValidateOnly:     cmd.Bool("validate-only"),
							}

							op, err := client.CreateInsightsConfig(ctx, req)
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

					{
						Name:  "describe",
						Usage: "describe insights-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "insights-config", Usage: "The ID of the insights config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/insightsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("insights-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.GetInsightsConfigRequest{
								Name: name,
							}

							resp, err := client.GetInsightsConfig(ctx, req)
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
						Usage: "update insights-configs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the insightsConfig is not found a new.", Required: false},
							&cli.StringFlag{Name: "insights-config", Usage: "The ID of the insights config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "insights_config.name" not yet supported.
							insights_config_name := fmt.Sprintf("projects/%s/locations/%s/insightsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("insights-config"))
							fmt.Printf("Executing update on %s\n", insights_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete insights-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.StringFlag{Name: "insights-config", Usage: "The ID of the insights config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/insightsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("insights-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInsightsConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.DeleteInsightsConfigRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteInsightsConfig(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
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
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
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
				},
			},
			{
				Name:  "users",
				Usage: "Manage users resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch-access-token",
						Usage: "fetch-access-token users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							account_connector := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							fmt.Printf("Executing fetch-access-token on %s\n", account_connector)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.ListUsersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListUsers(ctx, req)
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
						Name:  "delete",
						Usage: "delete users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request, but do not actually post it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s/users/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"), cmd.String("user"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteUser %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.DeleteUserRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteUser(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "fetch-self",
						Usage: "fetch-self users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.FetchSelfRequest{
								Name: name,
							}

							resp, err := client.FetchSelf(ctx, req)
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
						Usage: "delete users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSelf %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := developerconnect.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &developerconnectpb.DeleteSelfRequest{
								Name: name,
							}

							op, err := client.DeleteSelf(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "start-o-auth-flow",
						Usage: "start-o-auth-flow users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							account_connector := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							fmt.Printf("Executing start-o-auth-flow on %s\n", account_connector)
							return nil
						},
					},

					{
						Name:  "finish-o-auth-flow",
						Usage: "finish-o-auth-flow users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-connector", Usage: "The ID of the account connector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							account_connector := fmt.Sprintf("projects/%s/locations/%s/accountConnectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("account-connector"))
							fmt.Printf("Executing finish-o-auth-flow on %s\n", account_connector)
							return nil
						},
					},
				},
			},
		},
	}
}
