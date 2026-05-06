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

package configdelivery

import (
	configdelivery "cloud.google.com/go/configdelivery/apiv1"
	"cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the configdelivery command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "configdelivery",
		Usage: "manage Config Delivery API resources",
		Commands: []*cli.Command{
			{
				Name:  "fleet-packages",
				Usage: "Manage fleet-packages resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list fleet-packages",
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
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.ListFleetPackagesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFleetPackages(ctx, req)
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
						Usage: "describe fleet-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.GetFleetPackageRequest{
								Name: name,
							}

							resp, err := client.GetFleetPackage(ctx, req)
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
						Usage: "create fleet-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.CreateFleetPackageRequest{
								Parent:         parent,
								FleetPackageId: cmd.String("fleet-package-id"),
								RequestId:      cmd.String("request-id"),
							}

							op, err := client.CreateFleetPackage(ctx, req)
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
						Usage: "update fleet-packages",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "fleet_package.name" not yet supported.
							fleet_package_name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"))
							fmt.Printf("Executing update on %s\n", fleet_package_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete fleet-packages",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, then deleting an already deleted or non existing.", Required: false},
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any rollouts for this FleetPackage will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFleetPackage %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.DeleteFleetPackageRequest{
								Name:         name,
								RequestId:    cmd.String("request-id"),
								Force:        cmd.Bool("force"),
								AllowMissing: cmd.Bool("allow-missing"),
							}

							op, err := client.DeleteFleetPackage(ctx, req)
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
				},
			},
			{
				Name:  "releases",
				Usage: "Manage releases resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.ListReleasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListReleases(ctx, req)
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
						Usage: "describe releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.GetReleaseRequest{
								Name: name,
							}

							resp, err := client.GetRelease(ctx, req)
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
						Usage: "create releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.CreateReleaseRequest{
								Parent:    parent,
								ReleaseId: cmd.String("release-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateRelease(ctx, req)
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
						Usage: "update releases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "release.name" not yet supported.
							release_name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"))
							fmt.Printf("Executing update on %s\n", release_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete releases",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any variants of this release will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRelease %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.DeleteReleaseRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteRelease(ctx, req)
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
				Name:  "resource-bundles",
				Usage: "Manage resource-bundles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list resource-bundles",
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
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.ListResourceBundlesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListResourceBundles(ctx, req)
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
						Usage: "describe resource-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.GetResourceBundleRequest{
								Name: name,
							}

							resp, err := client.GetResourceBundle(ctx, req)
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
						Usage: "create resource-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle-id", Usage: "Id of the requesting object.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.CreateResourceBundleRequest{
								Parent:           parent,
								ResourceBundleId: cmd.String("resource-bundle-id"),
								RequestId:        cmd.String("request-id"),
							}

							op, err := client.CreateResourceBundle(ctx, req)
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
						Usage: "update resource-bundles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "resource_bundle.name" not yet supported.
							resource_bundle_name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"))
							fmt.Printf("Executing update on %s\n", resource_bundle_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete resource-bundles",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any releases of this resource bundle will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteResourceBundle %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.DeleteResourceBundleRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteResourceBundle(ctx, req)
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
				Name:  "rollouts",
				Usage: "Manage rollouts resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.ListRolloutsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListRollouts(ctx, req)
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
						Usage: "describe rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.GetRolloutRequest{
								Name: name,
							}

							resp, err := client.GetRollout(ctx, req)
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
						Name:  "suspend",
						Usage: "suspend rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "Reason for suspension.", Required: false},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.SuspendRolloutRequest{
								Name:   name,
								Reason: cmd.String("reason"),
							}

							op, err := client.SuspendRollout(ctx, req)
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
						Name:  "resume",
						Usage: "resume rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "Reason for resuming the rollout.", Required: false},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.ResumeRolloutRequest{
								Name:   name,
								Reason: cmd.String("reason"),
							}

							op, err := client.ResumeRollout(ctx, req)
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
						Name:  "abort",
						Usage: "abort rollouts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "fleet-package", Usage: "The ID of the fleet package.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "Reason for aborting.", Required: false},
							&cli.StringFlag{Name: "rollout", Usage: "The ID of the rollout.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet-package"), cmd.String("rollout"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.AbortRolloutRequest{
								Name:   name,
								Reason: cmd.String("reason"),
							}

							op, err := client.AbortRollout(ctx, req)
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
				},
			},
			{
				Name:  "variants",
				Usage: "Manage variants resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list variants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.ListVariantsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListVariants(ctx, req)
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
						Usage: "describe variants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
							&cli.StringFlag{Name: "variant", Usage: "The ID of the variant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"), cmd.String("variant"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.GetVariantRequest{
								Name: name,
							}

							resp, err := client.GetVariant(ctx, req)
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
						Usage: "create variants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
							&cli.StringFlag{Name: "variant-id", Usage: "Id of the requesting object.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.CreateVariantRequest{
								Parent:    parent,
								VariantId: cmd.String("variant-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateVariant(ctx, req)
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
						Usage: "update variants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
							&cli.StringFlag{Name: "variant", Usage: "The ID of the variant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "variant.name" not yet supported.
							variant_name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"), cmd.String("variant"))
							fmt.Printf("Executing update on %s\n", variant_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete variants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release", Usage: "The ID of the release.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource-bundle", Usage: "The ID of the resource bundle.", Required: true},
							&cli.StringFlag{Name: "variant", Usage: "The ID of the variant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource-bundle"), cmd.String("release"), cmd.String("variant"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteVariant %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := configdelivery.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &configdeliverypb.DeleteVariantRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteVariant(ctx, req)
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
		},
	}
}
