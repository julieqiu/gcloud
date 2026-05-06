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

package dataplex

import (
	dataplex "cloud.google.com/go/dataplex/apiv1"
	"cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the dataplex command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "dataplex",
		Usage: "manage Cloud Dataplex API resources",
		Commands: []*cli.Command{
			{
				Name:  "actions",
				Usage: "Manage actions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list actions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of actions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListLakeActions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListLakeActionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListLakeActions(ctx, req)
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
						Usage: "list actions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of actions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListZoneActions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListZoneActionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListZoneActions(ctx, req)
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
						Usage: "list actions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of actions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListAssetActions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListAssetActionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAssetActions(ctx, req)
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
				Name:  "aspect-types",
				Usage: "Manage aspect-types resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create aspect-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aspect-type-id", Usage: "AspectType identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateAspectTypeRequest{
								Parent:       parent,
								AspectTypeId: cmd.String("aspect-type-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateAspectType(ctx, req)
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
						Usage: "update aspect-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aspect-type", Usage: "The ID of the aspect type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "aspect_type.name" not yet supported.
							aspect_type_name := fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("aspect-type"))
							fmt.Printf("Executing update on %s\n", aspect_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete aspect-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aspect-type", Usage: "The ID of the aspect type.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If the client provided etag value does not match the current etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("aspect-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAspectType %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteAspectTypeRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteAspectType(ctx, req)
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
						Usage: "list aspect-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Orders the result by `name` or `create_time` fields.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of AspectTypes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListAspectTypes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListAspectTypesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAspectTypes(ctx, req)
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
						Usage: "describe aspect-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "aspect-type", Usage: "The ID of the aspect type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("aspect-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetAspectTypeRequest{
								Name: name,
							}

							resp, err := client.GetAspectType(ctx, req)
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
				Name:  "assets",
				Usage: "Manage assets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset-id", Usage: "Asset identifier.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateAssetRequest{
								Parent:       parent,
								AssetId:      cmd.String("asset-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateAsset(ctx, req)
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
						Usage: "update assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "asset.name" not yet supported.
							asset_name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("asset"))
							fmt.Printf("Executing update on %s\n", asset_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("asset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAsset %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteAssetRequest{
								Name: name,
							}

							op, err := client.DeleteAsset(ctx, req)
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
						Usage: "list assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of asset to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListAssets` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListAssetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAssets(ctx, req)
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
						Usage: "describe assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "asset", Usage: "The ID of the asset.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetAssetRequest{
								Name: name,
							}

							resp, err := client.GetAsset(ctx, req)
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
				Name:  "attributes",
				Usage: "Manage attributes resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create attributes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-attribute-id", Usage: "DataAttribute identifier.", Required: true},
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateDataAttributeRequest{
								Parent:          parent,
								DataAttributeId: cmd.String("data-attribute-id"),
								ValidateOnly:    cmd.Bool("validate-only"),
							}

							op, err := client.CreateDataAttribute(ctx, req)
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
						Usage: "update attributes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attribute", Usage: "The ID of the attribute.", Required: true},
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_attribute.name" not yet supported.
							data_attribute_name := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s/attributes/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"), cmd.String("attribute"))
							fmt.Printf("Executing update on %s\n", data_attribute_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete attributes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attribute", Usage: "The ID of the attribute.", Required: true},
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If the client provided etag value does not match the current etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s/attributes/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"), cmd.String("attribute"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataAttribute %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteDataAttributeRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteDataAttribute(ctx, req)
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
						Usage: "list attributes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of DataAttributes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListDataAttributes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataAttributesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataAttributes(ctx, req)
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
						Usage: "describe attributes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attribute", Usage: "The ID of the attribute.", Required: true},
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s/attributes/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"), cmd.String("attribute"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataAttributeRequest{
								Name: name,
							}

							resp, err := client.GetDataAttribute(ctx, req)
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
				Name:  "categories",
				Usage: "Manage categories resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create categories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "category-id", Usage: "GlossaryCategory identifier.", Required: true},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateGlossaryCategoryRequest{
								Parent:     parent,
								CategoryId: cmd.String("category-id"),
							}

							resp, err := client.CreateGlossaryCategory(ctx, req)
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
						Usage: "update categories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "categorie", Usage: "The ID of the categorie.", Required: true},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "category.name" not yet supported.
							category_name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/categories/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("categorie"))
							fmt.Printf("Executing update on %s\n", category_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete categories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "categorie", Usage: "The ID of the categorie.", Required: true},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/categories/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("categorie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGlossaryCategory on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteGlossaryCategoryRequest{
								Name: name,
							}

							if err := client.DeleteGlossaryCategory(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe categories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "categorie", Usage: "The ID of the categorie.", Required: true},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/categories/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("categorie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetGlossaryCategoryRequest{
								Name: name,
							}

							resp, err := client.GetGlossaryCategory(ctx, req)
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
						Usage: "list categories",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that filters GlossaryCategories listed in the.", Required: false},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by expression that orders GlossaryCategories listed in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of GlossaryCategories to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListGlossaryCategories`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListGlossaryCategoriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGlossaryCategories(ctx, req)
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
				Name:  "data-assets",
				Usage: "Manage data-assets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-asset-id", Usage: "The ID of the data asset to create.", Required: false},
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually creating the data asset.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateDataAssetRequest{
								Parent:       parent,
								DataAssetId:  cmd.String("data-asset-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateDataAsset(ctx, req)
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
						Usage: "update data-assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-asset", Usage: "The ID of the data asset.", Required: true},
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually updating the data asset.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_asset.name" not yet supported.
							data_asset_name := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s/dataAssets/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"), cmd.String("data-asset"))
							fmt.Printf("Executing update on %s\n", data_asset_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete data-assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-asset", Usage: "The ID of the data asset.", Required: true},
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag of the data asset.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually deleting the data asset.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s/dataAssets/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"), cmd.String("data-asset"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataAsset %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteDataAssetRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteDataAsset(ctx, req)
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
						Name:  "describe",
						Usage: "describe data-assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-asset", Usage: "The ID of the data asset.", Required: true},
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s/dataAssets/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"), cmd.String("data-asset"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataAssetRequest{
								Name: name,
							}

							resp, err := client.GetDataAsset(ctx, req)
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
						Usage: "list data-assets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that filters data assets listed in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by expression that orders data assets listed in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of data assets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDataAssets` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataAssetsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataAssets(ctx, req)
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
				Name:  "data-attribute-bindings",
				Usage: "Manage data-attribute-bindings resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-attribute-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-attribute-binding-id", Usage: "DataAttributeBinding identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateDataAttributeBindingRequest{
								Parent:                 parent,
								DataAttributeBindingId: cmd.String("data-attribute-binding-id"),
								ValidateOnly:           cmd.Bool("validate-only"),
							}

							op, err := client.CreateDataAttributeBinding(ctx, req)
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
						Usage: "update data-attribute-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-attribute-binding", Usage: "The ID of the data attribute binding.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_attribute_binding.name" not yet supported.
							data_attribute_binding_name := fmt.Sprintf("projects/%s/locations/%s/dataAttributeBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-attribute-binding"))
							fmt.Printf("Executing update on %s\n", data_attribute_binding_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete data-attribute-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-attribute-binding", Usage: "The ID of the data attribute binding.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If the client provided etag value does not match the current etag.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataAttributeBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-attribute-binding"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataAttributeBinding %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteDataAttributeBindingRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteDataAttributeBinding(ctx, req)
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
						Usage: "list data-attribute-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of DataAttributeBindings to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListDataAttributeBindings`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataAttributeBindingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataAttributeBindings(ctx, req)
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
						Usage: "describe data-attribute-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-attribute-binding", Usage: "The ID of the data attribute binding.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataAttributeBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-attribute-binding"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataAttributeBindingRequest{
								Name: name,
							}

							resp, err := client.GetDataAttributeBinding(ctx, req)
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
				Name:  "data-products",
				Usage: "Manage data-products resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-product-id", Usage: "The ID of the data product to create.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually creating the data product.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateDataProductRequest{
								Parent:        parent,
								DataProductId: cmd.String("data-product-id"),
								ValidateOnly:  cmd.Bool("validate-only"),
							}

							op, err := client.CreateDataProduct(ctx, req)
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
						Usage: "delete data-products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag of the data product.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually deleting the data product.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataProduct %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteDataProductRequest{
								Name:         name,
								Etag:         cmd.String("etag"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.DeleteDataProduct(ctx, req)
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
						Name:  "describe",
						Usage: "describe data-products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataProductRequest{
								Name: name,
							}

							resp, err := client.GetDataProduct(ctx, req)
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
						Usage: "list data-products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that filters data products listed in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by expression that orders data products listed in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of data products to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDataProducts` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataProductsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataProducts(ctx, req)
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
						Name:  "update",
						Usage: "update data-products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-product", Usage: "The ID of the data product.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually updating the data product.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_product.name" not yet supported.
							data_product_name := fmt.Sprintf("projects/%s/locations/%s/dataProducts/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-product"))
							fmt.Printf("Executing update on %s\n", data_product_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "data-scans",
				Usage: "Manage data-scans resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan-id", Usage: "DataScan identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateDataScanRequest{
								Parent:       parent,
								DataScanId:   cmd.String("data-scan-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateDataScan(ctx, req)
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
						Usage: "update data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_scan.name" not yet supported.
							data_scan_name := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"))
							fmt.Printf("Executing update on %s\n", data_scan_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any child resources of this data scan will also.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataScan %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteDataScanRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteDataScan(ctx, req)
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
						Name:  "describe",
						Usage: "describe data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Select the DataScan view to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataScanRequest{
								Name: name,
								View: dataplexpb.GetDataScanRequest_DataScanView(dataplexpb.GetDataScanRequest_DataScanView_value[cmd.String("view")]),
							}

							resp, err := client.GetDataScan(ctx, req)
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
						Usage: "list data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields (`name` or `create_time`) for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of dataScans to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListDataScans` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataScansRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataScans(ctx, req)
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
						Name:  "run",
						Usage: "run data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.RunDataScanRequest{
								Name: name,
							}

							resp, err := client.RunDataScan(ctx, req)
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
						Name:  "generate-data-quality-rules",
						Usage: "generate-data-quality-rules data-scans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GenerateDataQualityRulesRequest{
								Name: name,
							}

							resp, err := client.GenerateDataQualityRules(ctx, req)
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
				Name:  "data-taxonomies",
				Usage: "Manage data-taxonomies resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create data-taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-taxonomy-id", Usage: "DataTaxonomy identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateDataTaxonomyRequest{
								Parent:         parent,
								DataTaxonomyId: cmd.String("data-taxonomy-id"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.CreateDataTaxonomy(ctx, req)
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
						Usage: "update data-taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "data_taxonomy.name" not yet supported.
							data_taxonomy_name := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"))
							fmt.Printf("Executing update on %s\n", data_taxonomy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete data-taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If the client provided etag value does not match the current etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDataTaxonomy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteDataTaxonomyRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteDataTaxonomy(ctx, req)
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
						Usage: "list data-taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of DataTaxonomies to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous ` ListDataTaxonomies` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataTaxonomiesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataTaxonomies(ctx, req)
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
						Usage: "describe data-taxonomies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-taxonomie", Usage: "The ID of the data taxonomie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataTaxonomies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-taxonomie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataTaxonomyRequest{
								Name: name,
							}

							resp, err := client.GetDataTaxonomy(ctx, req)
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
				Name:  "encryption-configs",
				Usage: "Manage encryption-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create encryption-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encryption-config-id", Usage: "The ID of the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateEncryptionConfigRequest{
								Parent:             parent,
								EncryptionConfigId: cmd.String("encryption-config-id"),
							}

							op, err := client.CreateEncryptionConfig(ctx, req)
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
						Usage: "update encryption-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encryption-config", Usage: "The ID of the encryption config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "encryption_config.name" not yet supported.
							encryption_config_name := fmt.Sprintf("organizations/%s/locations/%s/encryptionConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("encryption-config"))
							fmt.Printf("Executing update on %s\n", encryption_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete encryption-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encryption-config", Usage: "The ID of the encryption config.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "Etag of the EncryptionConfig.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/encryptionConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("encryption-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEncryptionConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteEncryptionConfigRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteEncryptionConfig(ctx, req)
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
						Usage: "list encryption-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter the EncryptionConfigs to be returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of EncryptionConfigs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListEncryptionConfigs` call.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListEncryptionConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEncryptionConfigs(ctx, req)
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
						Usage: "describe encryption-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encryption-config", Usage: "The ID of the encryption config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/encryptionConfigs/%s", cmd.String("organization"), cmd.String("location"), cmd.String("encryption-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetEncryptionConfigRequest{
								Name: name,
							}

							resp, err := client.GetEncryptionConfig(ctx, req)
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
				Name:  "entities",
				Usage: "Manage entities resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateEntityRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							resp, err := client.CreateEntity(ctx, req)
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
						Usage: "update entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entity.name" not yet supported.
							entity_name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"))
							fmt.Printf("Executing update on %s\n", entity_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag associated with the entity, which can be retrieved with.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteEntity on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteEntityRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteEntity(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Used to select the subset of entity information to return.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetEntityRequest{
								Name: name,
								View: dataplexpb.GetEntityRequest_EntityView(dataplexpb.GetEntityRequest_EntityView_value[cmd.String("view")]),
							}

							resp, err := client.GetEntity(ctx, req)
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
						Usage: "list entities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The following filter parameters can be added to the URL to limit.", Required: false},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of entities to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListEntities` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specify the entity view to make a partial list request.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListEntitiesRequest{
								Parent:    parent,
								View:      dataplexpb.ListEntitiesRequest_EntityView(dataplexpb.ListEntitiesRequest_EntityView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntities(ctx, req)
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
				Name:  "entries",
				Usage: "Manage entries resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "entry-id", Usage: "Entry identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateEntryRequest{
								Parent:  parent,
								EntryId: cmd.String("entry-id"),
							}

							resp, err := client.CreateEntry(ctx, req)
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
						Usage: "update entries",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true and the entry doesn't exist, the service will.", Required: false},
							&cli.StringSliceFlag{Name: "aspect-keys", Usage: "The map keys of the Aspects which the service should modify.", Required: false},
							&cli.BoolFlag{Name: "delete-missing-aspects", Usage: "If set to true and the aspect_keys specify aspect ranges, the.", Required: false},
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entry.name" not yet supported.
							entry_name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							fmt.Printf("Executing update on %s\n", entry_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteEntryRequest{
								Name: name,
							}

							resp, err := client.DeleteEntry(ctx, req)
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
						Usage: "list entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter on the entries to return.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListEntries` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListEntriesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntries(ctx, req)
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
						Usage: "describe entries",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "aspect-types", Usage: "Limits the aspects returned to the provided aspect types.", Required: false},
							&cli.StringFlag{Name: "entrie", Usage: "The ID of the entrie.", Required: true},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "paths", Usage: "Limits the aspects returned to those associated with the provided.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "View to control which parts of an entry the service should.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entries/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetEntryRequest{
								Name:        name,
								View:        dataplexpb.EntryView(dataplexpb.EntryView_value[cmd.String("view")]),
								AspectTypes: cmd.StringSlice("aspect-types"),
								Paths:       cmd.StringSlice("paths"),
							}

							resp, err := client.GetEntry(ctx, req)
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
				Name:  "entry-groups",
				Usage: "Manage entry-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group-id", Usage: "EntryGroup identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateEntryGroupRequest{
								Parent:       parent,
								EntryGroupId: cmd.String("entry-group-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateEntryGroup(ctx, req)
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
						Usage: "update entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request, without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entry_group.name" not yet supported.
							entry_group_name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							fmt.Printf("Executing update on %s\n", entry_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If the client provided etag value does not match the current etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEntryGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteEntryGroupRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteEntryGroup(ctx, req)
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
						Usage: "list entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of EntryGroups to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListEntryGroups` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListEntryGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntryGroups(ctx, req)
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
						Usage: "describe entry-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetEntryGroupRequest{
								Name: name,
							}

							resp, err := client.GetEntryGroup(ctx, req)
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
				Name:  "entry-links",
				Usage: "Manage entry-links resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entry-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "entry-link-id", Usage: "Entry Link identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateEntryLinkRequest{
								Parent:      parent,
								EntryLinkId: cmd.String("entry-link-id"),
							}

							resp, err := client.CreateEntryLink(ctx, req)
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
						Usage: "update entry-links",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true and the entry link doesn't exist, the service will.", Required: false},
							&cli.StringSliceFlag{Name: "aspect-keys", Usage: "The map keys of the Aspects which the service should modify.", Required: false},
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "entry-link", Usage: "The ID of the entry link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entry_link.name" not yet supported.
							entry_link_name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entry-link"))
							fmt.Printf("Executing update on %s\n", entry_link_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entry-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "entry-link", Usage: "The ID of the entry link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entry-link"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteEntryLinkRequest{
								Name: name,
							}

							resp, err := client.DeleteEntryLink(ctx, req)
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
						Usage: "describe entry-links",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-group", Usage: "The ID of the entry group.", Required: true},
							&cli.StringFlag{Name: "entry-link", Usage: "The ID of the entry link.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s/entryLinks/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-group"), cmd.String("entry-link"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetEntryLinkRequest{
								Name: name,
							}

							resp, err := client.GetEntryLink(ctx, req)
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
				Name:  "entry-types",
				Usage: "Manage entry-types resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create entry-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-type-id", Usage: "EntryType identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateEntryTypeRequest{
								Parent:       parent,
								EntryTypeId:  cmd.String("entry-type-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateEntryType(ctx, req)
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
						Usage: "update entry-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-type", Usage: "The ID of the entry type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "entry_type.name" not yet supported.
							entry_type_name := fmt.Sprintf("projects/%s/locations/%s/entryTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-type"))
							fmt.Printf("Executing update on %s\n", entry_type_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete entry-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-type", Usage: "The ID of the entry type.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "If the client provided etag value does not match the current etag.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-type"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEntryType %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteEntryTypeRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteEntryType(ctx, req)
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
						Usage: "list entry-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Orders the result by `name` or `create_time` fields.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of EntryTypes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListEntryTypes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListEntryTypesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntryTypes(ctx, req)
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
						Usage: "describe entry-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry-type", Usage: "The ID of the entry type.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/entryTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("entry-type"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetEntryTypeRequest{
								Name: name,
							}

							resp, err := client.GetEntryType(ctx, req)
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
				Name:  "glossaries",
				Usage: "Manage glossaries resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossary-id", Usage: "Glossary ID: Glossary identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually creating the Glossary.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateGlossaryRequest{
								Parent:       parent,
								GlossaryId:   cmd.String("glossary-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateGlossary(ctx, req)
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
						Usage: "update glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validates the request without actually updating the Glossary.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "glossary.name" not yet supported.
							glossary_name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							fmt.Printf("Executing update on %s\n", glossary_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag of the Glossary.", Required: false},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGlossary %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteGlossaryRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteGlossary(ctx, req)
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
						Name:  "describe",
						Usage: "describe glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetGlossaryRequest{
								Name: name,
							}

							resp, err := client.GetGlossary(ctx, req)
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
						Usage: "list glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that filters Glossaries listed in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by expression that orders Glossaries listed in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Glossaries to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListGlossaries` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListGlossariesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGlossaries(ctx, req)
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
				Name:  "jobs",
				Usage: "Manage jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Select the DataScanJob view to return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetDataScanJobRequest{
								Name: name,
								View: dataplexpb.GetDataScanJobRequest_DataScanJobView(dataplexpb.GetDataScanJobRequest_DataScanJobView_value[cmd.String("view")]),
							}

							resp, err := client.GetDataScanJob(ctx, req)
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
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data-scan", Usage: "The ID of the data scan.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the ListDataScanJobs.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of DataScanJobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListDataScanJobs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/dataScans/%s", cmd.String("project"), cmd.String("location"), cmd.String("data-scan"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListDataScanJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDataScanJobs(ctx, req)
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
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of jobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListJobs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListJobs(ctx, req)
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
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetJobRequest{
								Name: name,
							}

							resp, err := client.GetJob(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"), cmd.String("job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CancelJobRequest{
								Name: name,
							}

							if err := client.CancelJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "lakes",
				Usage: "Manage lakes resources",
				Commands: []*cli.Command{

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake-id", Usage: "Lake identifier.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateLakeRequest{
								Parent:       parent,
								LakeId:       cmd.String("lake-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateLake(ctx, req)
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
						Usage: "update lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "lake.name" not yet supported.
							lake_name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing update on %s\n", lake_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteLake %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteLakeRequest{
								Name: name,
							}

							op, err := client.DeleteLake(ctx, req)
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
						Usage: "list lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Lakes to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListLakes` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListLakesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListLakes(ctx, req)
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
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetLakeRequest{
								Name: name,
							}

							resp, err := client.GetLake(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions lakes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
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
						Name:  "lookup-entry",
						Usage: "lookup-entry locations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "aspect-types", Usage: "Limits the aspects returned to the provided aspect types.", Required: false},
							&cli.StringFlag{Name: "entry", Usage: "The resource name of the Entry:.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "paths", Usage: "Limits the aspects returned to those associated with the provided.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "View to control which parts of an entry the service should.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.LookupEntryRequest{
								Name:        name,
								View:        dataplexpb.EntryView(dataplexpb.EntryView_value[cmd.String("view")]),
								AspectTypes: cmd.StringSlice("aspect-types"),
								Paths:       cmd.StringSlice("paths"),
								Entry:       cmd.String("entry"),
							}

							resp, err := client.LookupEntry(ctx, req)
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
						Name:  "search-entries",
						Usage: "search-entries locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the ordering of results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results in the search page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `SearchEntries` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "The query against which entries in scope should be matched.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope under which the search should be operating.", Required: false},
							&cli.BoolFlag{Name: "semantic-search", Usage: "Specifies whether the search should understand the meaning and.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.SearchEntriesRequest{
								Name:           name,
								Query:          cmd.String("query"),
								PageSize:       int32(cmd.Int("page-size")),
								PageToken:      cmd.String("page-token"),
								OrderBy:        cmd.String("order-by"),
								Scope:          cmd.String("scope"),
								SemanticSearch: cmd.Bool("semantic-search"),
							}

							limit := cmd.Int("limit")
							it := client.SearchEntries(ctx, req)
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
						Name:  "lookup-entry-links",
						Usage: "lookup-entry-links locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entry", Usage: "The resource name of the referred Entry.", Required: true},
							&cli.StringSliceFlag{Name: "entry-link-types", Usage: "Entry link types to filter the response by.", Required: false},
							&cli.StringFlag{Name: "entry-mode", Usage: "Mode of entry reference.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of EntryLinks to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `LookupEntryLinks` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.LookupEntryLinksRequest{
								Name:           name,
								Entry:          cmd.String("entry"),
								EntryMode:      dataplexpb.LookupEntryLinksRequest_EntryMode(dataplexpb.LookupEntryLinksRequest_EntryMode_value[cmd.String("entry-mode")]),
								EntryLinkTypes: cmd.StringSlice("entry-link-types"),
								PageSize:       int32(cmd.Int("page-size")),
								PageToken:      cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.LookupEntryLinks(ctx, req)
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
						Name:  "lookup-context",
						Usage: "lookup-context locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "resources", Usage: "The entry names to lookup context for.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.LookupContextRequest{
								Name:      name,
								Resources: cmd.StringSlice("resources"),
							}

							resp, err := client.LookupContext(ctx, req)
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
				Name:  "metadata-feeds",
				Usage: "Manage metadata-feeds resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create metadata-feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-feed-id", Usage: "The metadata job ID.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateMetadataFeedRequest{
								Parent:         parent,
								MetadataFeedId: cmd.String("metadata-feed-id"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.CreateMetadataFeed(ctx, req)
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
						Usage: "describe metadata-feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-feed", Usage: "The ID of the metadata feed.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataFeeds/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-feed"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetMetadataFeedRequest{
								Name: name,
							}

							resp, err := client.GetMetadataFeed(ctx, req)
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
						Usage: "list metadata-feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to sort the results by, either `name` or `create_time`.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of metadata feeds to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token received from a previous `ListMetadataFeeds` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListMetadataFeedsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMetadataFeeds(ctx, req)
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
						Usage: "delete metadata-feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-feed", Usage: "The ID of the metadata feed.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataFeeds/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-feed"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMetadataFeed %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteMetadataFeedRequest{
								Name: name,
							}

							op, err := client.DeleteMetadataFeed(ctx, req)
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
						Name:  "update",
						Usage: "update metadata-feeds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-feed", Usage: "The ID of the metadata feed.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "metadata_feed.name" not yet supported.
							metadata_feed_name := fmt.Sprintf("projects/%s/locations/%s/metadataFeeds/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-feed"))
							fmt.Printf("Executing update on %s\n", metadata_feed_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "metadata-jobs",
				Usage: "Manage metadata-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create metadata-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-job-id", Usage: "The metadata job ID.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The service validates the request without performing any.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateMetadataJobRequest{
								Parent:        parent,
								MetadataJobId: cmd.String("metadata-job-id"),
								ValidateOnly:  cmd.Bool("validate-only"),
							}

							op, err := client.CreateMetadataJob(ctx, req)
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
						Usage: "describe metadata-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-job", Usage: "The ID of the metadata job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetMetadataJobRequest{
								Name: name,
							}

							resp, err := client.GetMetadataJob(ctx, req)
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
						Usage: "list metadata-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "The field to sort the results by, either `name` or `create_time`.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of metadata jobs to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token received from a previous `ListMetadataJobs` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListMetadataJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMetadataJobs(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel metadata-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "metadata-job", Usage: "The ID of the metadata job.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/metadataJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("metadata-job"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute CancelMetadataJob on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CancelMetadataJobRequest{
								Name: name,
							}

							if err := client.CancelMetadataJob(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
				Name:  "partitions",
				Usage: "Manage partitions resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create partitions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreatePartitionRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
							}

							resp, err := client.CreatePartition(ctx, req)
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
						Usage: "delete partitions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag associated with the partition.", Required: false},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "partition", Usage: "The ID of the partition.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s/partitions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"), cmd.String("partition"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePartition on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeletePartitionRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeletePartition(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe partitions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "partition", Usage: "The ID of the partition.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s/partitions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"), cmd.String("partition"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetPartitionRequest{
								Name: name,
							}

							resp, err := client.GetPartition(ctx, req)
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
						Usage: "list partitions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "entitie", Usage: "The ID of the entitie.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Filter the partitions returned to the caller using a key value.", Required: false},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of partitions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListPartitions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"), cmd.String("entitie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListPartitionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPartitions(ctx, req)
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
				Name:  "tasks",
				Usage: "Manage tasks resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task-id", Usage: "Task identifier.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateTaskRequest{
								Parent:       parent,
								TaskId:       cmd.String("task-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateTask(ctx, req)
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
						Usage: "update tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "task.name" not yet supported.
							task_name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"))
							fmt.Printf("Executing update on %s\n", task_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTask %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteTaskRequest{
								Name: name,
							}

							op, err := client.DeleteTask(ctx, req)
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
						Usage: "list tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of tasks to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListZones` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListTasksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListTasks(ctx, req)
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
						Usage: "describe tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetTaskRequest{
								Name: name,
							}

							resp, err := client.GetTask(ctx, req)
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
						Name:  "run",
						Usage: "run tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("task"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.RunTaskRequest{
								Name: name,
							}

							resp, err := client.RunTask(ctx, req)
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
				Name:  "terms",
				Usage: "Manage terms resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create terms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "term-id", Usage: "GlossaryTerm identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateGlossaryTermRequest{
								Parent: parent,
								TermId: cmd.String("term-id"),
							}

							resp, err := client.CreateGlossaryTerm(ctx, req)
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
						Usage: "update terms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "term", Usage: "The ID of the term.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "term.name" not yet supported.
							term_name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/terms/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("term"))
							fmt.Printf("Executing update on %s\n", term_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete terms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "term", Usage: "The ID of the term.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/terms/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("term"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteGlossaryTerm on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteGlossaryTermRequest{
								Name: name,
							}

							if err := client.DeleteGlossaryTerm(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe terms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "term", Usage: "The ID of the term.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/terms/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"), cmd.String("term"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetGlossaryTermRequest{
								Name: name,
							}

							resp, err := client.GetGlossaryTerm(ctx, req)
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
						Usage: "list terms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that filters GlossaryTerms listed in the.", Required: false},
							&cli.StringFlag{Name: "glossarie", Usage: "The ID of the glossarie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by expression that orders GlossaryTerms listed in the.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of GlossaryTerms to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListGlossaryTerms` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossarie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListGlossaryTermsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListGlossaryTerms(ctx, req)
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
				Name:  "zones",
				Usage: "Manage zones resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone-id", Usage: "Zone identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.CreateZoneRequest{
								Parent:       parent,
								ZoneId:       cmd.String("zone-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateZone(ctx, req)
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
						Usage: "update zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Only validate the request, but do not perform mutations.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "zone.name" not yet supported.
							zone_name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							fmt.Printf("Executing update on %s\n", zone_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteZone %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.DeleteZoneRequest{
								Name: name,
							}

							op, err := client.DeleteZone(ctx, req)
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
						Usage: "list zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter request.", Required: false},
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Order by fields for the result.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of zones to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token received from a previous `ListZones` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/lakes/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.ListZonesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListZones(ctx, req)
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
						Usage: "describe zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "lake", Usage: "The ID of the lake.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", cmd.String("project"), cmd.String("location"), cmd.String("lake"), cmd.String("zone"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := dataplex.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &dataplexpb.GetZoneRequest{
								Name: name,
							}

							resp, err := client.GetZone(ctx, req)
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
