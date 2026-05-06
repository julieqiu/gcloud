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

package inventory

import (
	inventory "cloud.google.com/go/inventory/apiv1"
	"cloud.google.com/go/inventory/apiv1/inventorypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the kmsinventory command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "kmsinventory",
		Usage: "manage KMS Inventory API resources",
		Commands: []*cli.Command{
			{
				Name:  "crypto-keys",
				Usage: "Manage crypto-keys resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list crypto-keys",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of keys to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pass this into a subsequent request in order to receive the next.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := inventory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &inventorypb.ListCryptoKeysRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCryptoKeys(ctx, req)
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
				Name:  "protected-resources",
				Usage: "Manage protected-resources resources",
				Commands: []*cli.Command{

					{
						Name:  "search",
						Usage: "search protected-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The resource name of the.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of resources to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringSliceFlag{Name: "resource-types", Usage: "A list of resource types that this request searches for.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "A scope can be an organization or a project.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							scope := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing search on %s\n", scope)
							return nil
						},
					},
				},
			},
			{
				Name:  "protected-resources-summary",
				Usage: "Manage protected-resources-summary resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe protected-resources-summary",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "fallback-scope", Usage: "The scope to use if the kms organization service account is not.", Required: false},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := inventory.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &inventorypb.GetProtectedResourcesSummaryRequest{
								Name:          name,
								FallbackScope: inventorypb.FallbackScope(inventorypb.FallbackScope_value[cmd.String("fallback-scope")]),
							}

							resp, err := client.GetProtectedResourcesSummary(ctx, req)
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
