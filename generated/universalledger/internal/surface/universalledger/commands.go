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

package universalledger

import (
	universalledger "cloud.google.com/go/universalledger/apiv1"
	"cloud.google.com/go/universalledger/apiv1/universalledgerpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the universalledger command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "universalledger",
		Usage: "manage Universal Ledger API resources",
		Commands: []*cli.Command{
			{
				Name:  "endpoints",
				Usage: "Manage endpoints resources",
				Commands: []*cli.Command{

					{
						Name:  "submit-transaction",
						Usage: "submit-transaction endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serialized-signed-transaction", Usage: "A protobuf serialized.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing submit-transaction on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of endpoints to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListEndpoints` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := universalledger.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &universalledgerpb.ListEndpointsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEndpoints(ctx, req)
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
						Usage: "describe endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := universalledger.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &universalledgerpb.GetEndpointRequest{
								Name: name,
							}

							resp, err := client.GetEndpoint(ctx, req)
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
						Name:  "submit-operational-transaction",
						Usage: "submit-operational-transaction endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "serialized-signed-operational-transaction", Usage: "A protobuf serialized.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing submit-operational-transaction on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "query-transaction-state",
						Usage: "query-transaction-state endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transaction-digest-hex", Usage: "The transaction digest to get the state of.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing query-transaction-state on %s\n", endpoint)
							return nil
						},
					},

					{
						Name:  "query-account",
						Usage: "query-account endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "account-id", Usage: "The account ID to get information about.", Required: true},
							&cli.StringFlag{Name: "endpoint", Usage: "The ID of the endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "round-id", Usage: "The ID of the execution round (similar to block \"height\") at.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							endpoint := fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint"))
							fmt.Printf("Executing query-account on %s\n", endpoint)
							return nil
						},
					},
				},
			},
		},
	}
}
