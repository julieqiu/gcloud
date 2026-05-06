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

package networksecurity

import (
	networksecurity "cloud.google.com/go/networksecurity/apiv1"
	"cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the networksecurity command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "networksecurity",
		Usage: "manage Network Security API resources",
		Commands: []*cli.Command{
			{
				Name:  "address-groups",
				Usage: "Manage address-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of AddressGroups to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListAddressGroupsRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListAddressGroups(ctx, req)
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
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetAddressGroupRequest{
								Name: name,
							}

							resp, err := client.GetAddressGroup(ctx, req)
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
						Usage: "create address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group-id", Usage: "Short name of the AddressGroup resource to be created.", Required: true},
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
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateAddressGroupRequest{
								Parent:         parent,
								AddressGroupId: cmd.String("address-group-id"),
								RequestId:      cmd.String("request-id"),
							}

							op, err := client.CreateAddressGroup(ctx, req)
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
						Usage: "update address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "address_group.name" not yet supported.
							address_group_name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing update on %s\n", address_group_name)
							return nil
						},
					},

					{
						Name:  "add-items",
						Usage: "add-items address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringSliceFlag{Name: "items", Usage: "List of items to add.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing add-items on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "remove-items",
						Usage: "remove-items address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringSliceFlag{Name: "items", Usage: "List of items to remove.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing remove-items on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "clone-items",
						Usage: "clone-items address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-address-group", Usage: "Source address group to clone items from.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing clone-items on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAddressGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteAddressGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteAddressGroup(ctx, req)
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
						Usage: "list address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of references to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing list on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of AddressGroups to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "If true, allow partial responses for multi-regional Aggregated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListAddressGroupsRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListAddressGroups(ctx, req)
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
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetAddressGroupRequest{
								Name: name,
							}

							resp, err := client.GetAddressGroup(ctx, req)
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
						Usage: "create address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group-id", Usage: "Short name of the AddressGroup resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateAddressGroupRequest{
								Parent:         parent,
								AddressGroupId: cmd.String("address-group-id"),
								RequestId:      cmd.String("request-id"),
							}

							op, err := client.CreateAddressGroup(ctx, req)
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
						Usage: "update address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "address_group.name" not yet supported.
							address_group_name := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing update on %s\n", address_group_name)
							return nil
						},
					},

					{
						Name:  "add-items",
						Usage: "add-items address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringSliceFlag{Name: "items", Usage: "List of items to add.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing add-items on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "remove-items",
						Usage: "remove-items address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringSliceFlag{Name: "items", Usage: "List of items to remove.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing remove-items on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "clone-items",
						Usage: "clone-items address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-address-group", Usage: "Source address group to clone items from.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing clone-items on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAddressGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteAddressGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteAddressGroup(ctx, req)
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
						Usage: "list address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of references to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							address_group := fmt.Sprintf("organizations/%s/locations/%s/addressGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing list on %s\n", address_group)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-group", Usage: "The ID of the address group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address-group"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "authorization-policies",
				Usage: "Manage authorization-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list authorization-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of AuthorizationPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListAuthorizationPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAuthorizationPolicies(ctx, req)
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
						Usage: "describe authorization-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-policie", Usage: "The ID of the authorization policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetAuthorizationPolicyRequest{
								Name: name,
							}

							resp, err := client.GetAuthorizationPolicy(ctx, req)
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
						Usage: "create authorization-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-policy-id", Usage: "Short name of the AuthorizationPolicy resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateAuthorizationPolicyRequest{
								Parent:                parent,
								AuthorizationPolicyId: cmd.String("authorization-policy-id"),
							}

							op, err := client.CreateAuthorizationPolicy(ctx, req)
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
						Usage: "update authorization-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-policie", Usage: "The ID of the authorization policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "authorization_policy.name" not yet supported.
							authorization_policy_name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization-policie"))
							fmt.Printf("Executing update on %s\n", authorization_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete authorization-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-policie", Usage: "The ID of the authorization policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAuthorizationPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteAuthorizationPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteAuthorizationPolicy(ctx, req)
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
				Name:  "authz-policies",
				Usage: "Manage authz-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list authz-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results that the server returns.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListAuthzPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAuthzPolicies(ctx, req)
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
						Usage: "describe authz-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-policie", Usage: "The ID of the authz policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetAuthzPolicyRequest{
								Name: name,
							}

							resp, err := client.GetAuthzPolicy(ctx, req)
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
						Usage: "create authz-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-policy-id", Usage: "User-provided ID of the `AuthzPolicy` resource to be created.", Required: true},
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
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateAuthzPolicyRequest{
								Parent:        parent,
								AuthzPolicyId: cmd.String("authz-policy-id"),
								RequestId:     cmd.String("request-id"),
							}

							op, err := client.CreateAuthzPolicy(ctx, req)
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
						Usage: "update authz-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-policie", Usage: "The ID of the authz policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "authz_policy.name" not yet supported.
							authz_policy_name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz-policie"))
							fmt.Printf("Executing update on %s\n", authz_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete authz-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authz-policie", Usage: "The ID of the authz policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAuthzPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteAuthzPolicyRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteAuthzPolicy(ctx, req)
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
				Name:  "backend-authentication-configs",
				Usage: "Manage backend-authentication-configs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list backend-authentication-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of BackendAuthenticationConfigs to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListBackendAuthenticationConfigsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListBackendAuthenticationConfigsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListBackendAuthenticationConfigs(ctx, req)
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
						Usage: "describe backend-authentication-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend-authentication-config", Usage: "The ID of the backend authentication config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend-authentication-config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetBackendAuthenticationConfigRequest{
								Name: name,
							}

							resp, err := client.GetBackendAuthenticationConfig(ctx, req)
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
						Usage: "create backend-authentication-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend-authentication-config-id", Usage: "Short name of the BackendAuthenticationConfig resource to be.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateBackendAuthenticationConfigRequest{
								Parent:                        parent,
								BackendAuthenticationConfigId: cmd.String("backend-authentication-config-id"),
							}

							op, err := client.CreateBackendAuthenticationConfig(ctx, req)
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
						Usage: "update backend-authentication-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend-authentication-config", Usage: "The ID of the backend authentication config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "backend_authentication_config.name" not yet supported.
							backend_authentication_config_name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend-authentication-config"))
							fmt.Printf("Executing update on %s\n", backend_authentication_config_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backend-authentication-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend-authentication-config", Usage: "The ID of the backend authentication config.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "Etag of the resource.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend-authentication-config"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteBackendAuthenticationConfig %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteBackendAuthenticationConfigRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteBackendAuthenticationConfig(ctx, req)
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
				Name:  "client-tls-policies",
				Usage: "Manage client-tls-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list client-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of ClientTlsPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListClientTlsPoliciesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListClientTlsPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListClientTlsPolicies(ctx, req)
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
						Usage: "describe client-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "client-tls-policie", Usage: "The ID of the client tls policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client-tls-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetClientTlsPolicyRequest{
								Name: name,
							}

							resp, err := client.GetClientTlsPolicy(ctx, req)
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
						Usage: "create client-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "client-tls-policy-id", Usage: "Short name of the ClientTlsPolicy resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateClientTlsPolicyRequest{
								Parent:            parent,
								ClientTlsPolicyId: cmd.String("client-tls-policy-id"),
							}

							op, err := client.CreateClientTlsPolicy(ctx, req)
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
						Usage: "update client-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "client-tls-policie", Usage: "The ID of the client tls policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "client_tls_policy.name" not yet supported.
							client_tls_policy_name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client-tls-policie"))
							fmt.Printf("Executing update on %s\n", client_tls_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete client-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "client-tls-policie", Usage: "The ID of the client tls policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client-tls-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteClientTlsPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteClientTlsPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteClientTlsPolicy(ctx, req)
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
				Name:  "dns-threat-detectors",
				Usage: "Manage dns-threat-detectors resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list dns-threat-detectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListDnsThreatDetectorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDnsThreatDetectors(ctx, req)
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
						Usage: "describe dns-threat-detectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dns-threat-detector", Usage: "The ID of the dns threat detector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns-threat-detector"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetDnsThreatDetectorRequest{
								Name: name,
							}

							resp, err := client.GetDnsThreatDetector(ctx, req)
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
						Usage: "create dns-threat-detectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dns-threat-detector-id", Usage: "The ID of the requesting DnsThreatDetector object.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateDnsThreatDetectorRequest{
								Parent:              parent,
								DnsThreatDetectorId: cmd.String("dns-threat-detector-id"),
							}

							resp, err := client.CreateDnsThreatDetector(ctx, req)
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
						Usage: "update dns-threat-detectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dns-threat-detector", Usage: "The ID of the dns threat detector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "dns_threat_detector.name" not yet supported.
							dns_threat_detector_name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns-threat-detector"))
							fmt.Printf("Executing update on %s\n", dns_threat_detector_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete dns-threat-detectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dns-threat-detector", Usage: "The ID of the dns threat detector.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns-threat-detector"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDnsThreatDetector on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteDnsThreatDetectorRequest{
								Name: name,
							}

							if err := client.DeleteDnsThreatDetector(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "firewall-endpoint-associations",
				Usage: "Manage firewall-endpoint-associations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list firewall-endpoint-associations",
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
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListFirewallEndpointAssociationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFirewallEndpointAssociations(ctx, req)
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
						Usage: "describe firewall-endpoint-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint-association", Usage: "The ID of the firewall endpoint association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall-endpoint-association"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetFirewallEndpointAssociationRequest{
								Name: name,
							}

							resp, err := client.GetFirewallEndpointAssociation(ctx, req)
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
						Usage: "create firewall-endpoint-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint-association-id", Usage: "Id of the requesting object.", Required: false},
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
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateFirewallEndpointAssociationRequest{
								Parent:                        parent,
								FirewallEndpointAssociationId: cmd.String("firewall-endpoint-association-id"),
								RequestId:                     cmd.String("request-id"),
							}

							op, err := client.CreateFirewallEndpointAssociation(ctx, req)
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
						Usage: "delete firewall-endpoint-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint-association", Usage: "The ID of the firewall endpoint association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall-endpoint-association"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFirewallEndpointAssociation %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteFirewallEndpointAssociationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteFirewallEndpointAssociation(ctx, req)
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
						Usage: "update firewall-endpoint-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint-association", Usage: "The ID of the firewall endpoint association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "firewall_endpoint_association.name" not yet supported.
							firewall_endpoint_association_name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall-endpoint-association"))
							fmt.Printf("Executing update on %s\n", firewall_endpoint_association_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "firewall-endpoints",
				Usage: "Manage firewall-endpoints resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list firewall-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filtering results.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Hint for how to order the results.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListFirewallEndpointsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListFirewallEndpoints(ctx, req)
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
						Usage: "describe firewall-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint", Usage: "The ID of the firewall endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall-endpoint"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetFirewallEndpointRequest{
								Name: name,
							}

							resp, err := client.GetFirewallEndpoint(ctx, req)
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
						Usage: "create firewall-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint-id", Usage: "Id of the requesting object.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateFirewallEndpointRequest{
								Parent:             parent,
								FirewallEndpointId: cmd.String("firewall-endpoint-id"),
								RequestId:          cmd.String("request-id"),
							}

							op, err := client.CreateFirewallEndpoint(ctx, req)
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
						Usage: "delete firewall-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint", Usage: "The ID of the firewall endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall-endpoint"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteFirewallEndpoint %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteFirewallEndpointRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteFirewallEndpoint(ctx, req)
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
						Usage: "update firewall-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall-endpoint", Usage: "The ID of the firewall endpoint.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "firewall_endpoint.name" not yet supported.
							firewall_endpoint_name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall-endpoint"))
							fmt.Printf("Executing update on %s\n", firewall_endpoint_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "gateway-security-policies",
				Usage: "Manage gateway-security-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list gateway-security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of GatewaySecurityPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListGatewaySecurityPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGatewaySecurityPolicies(ctx, req)
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
						Usage: "describe gateway-security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetGatewaySecurityPolicyRequest{
								Name: name,
							}

							resp, err := client.GetGatewaySecurityPolicy(ctx, req)
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
						Usage: "create gateway-security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policy-id", Usage: "Short name of the GatewaySecurityPolicy resource to be created.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateGatewaySecurityPolicyRequest{
								Parent:                  parent,
								GatewaySecurityPolicyId: cmd.String("gateway-security-policy-id"),
							}

							op, err := client.CreateGatewaySecurityPolicy(ctx, req)
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
						Usage: "update gateway-security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "gateway_security_policy.name" not yet supported.
							gateway_security_policy_name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"))
							fmt.Printf("Executing update on %s\n", gateway_security_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete gateway-security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGatewaySecurityPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteGatewaySecurityPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteGatewaySecurityPolicy(ctx, req)
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
				Name:  "intercept-deployment-groups",
				Usage: "Manage intercept-deployment-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list intercept-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListInterceptDeploymentGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInterceptDeploymentGroups(ctx, req)
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
						Usage: "describe intercept-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The ID of the intercept deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetInterceptDeploymentGroupRequest{
								Name: name,
							}

							resp, err := client.GetInterceptDeploymentGroup(ctx, req)
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
						Usage: "create intercept-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment-group-id", Usage: "The ID to use for the new deployment group, which will become the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateInterceptDeploymentGroupRequest{
								Parent:                     parent,
								InterceptDeploymentGroupId: cmd.String("intercept-deployment-group-id"),
								RequestId:                  cmd.String("request-id"),
							}

							op, err := client.CreateInterceptDeploymentGroup(ctx, req)
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
						Usage: "update intercept-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The ID of the intercept deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "intercept_deployment_group.name" not yet supported.
							intercept_deployment_group_name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-deployment-group"))
							fmt.Printf("Executing update on %s\n", intercept_deployment_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete intercept-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The ID of the intercept deployment group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-deployment-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInterceptDeploymentGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteInterceptDeploymentGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInterceptDeploymentGroup(ctx, req)
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
				Name:  "intercept-deployments",
				Usage: "Manage intercept-deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list intercept-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListInterceptDeployments`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListInterceptDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInterceptDeployments(ctx, req)
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
						Usage: "describe intercept-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment", Usage: "The ID of the intercept deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetInterceptDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetInterceptDeployment(ctx, req)
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
						Usage: "create intercept-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment-id", Usage: "The ID to use for the new deployment, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateInterceptDeploymentRequest{
								Parent:                parent,
								InterceptDeploymentId: cmd.String("intercept-deployment-id"),
								RequestId:             cmd.String("request-id"),
							}

							op, err := client.CreateInterceptDeployment(ctx, req)
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
						Usage: "update intercept-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment", Usage: "The ID of the intercept deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "intercept_deployment.name" not yet supported.
							intercept_deployment_name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-deployment"))
							fmt.Printf("Executing update on %s\n", intercept_deployment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete intercept-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-deployment", Usage: "The ID of the intercept deployment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInterceptDeployment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteInterceptDeploymentRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInterceptDeployment(ctx, req)
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
				Name:  "intercept-endpoint-group-associations",
				Usage: "Manage intercept-endpoint-group-associations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list intercept-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListInterceptEndpointGroupAssociationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInterceptEndpointGroupAssociations(ctx, req)
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
						Usage: "describe intercept-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group-association", Usage: "The ID of the intercept endpoint group association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-endpoint-group-association"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetInterceptEndpointGroupAssociationRequest{
								Name: name,
							}

							resp, err := client.GetInterceptEndpointGroupAssociation(ctx, req)
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
						Usage: "create intercept-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group-association-id", Usage: "The ID to use for the new association, which will become the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateInterceptEndpointGroupAssociationRequest{
								Parent:                              parent,
								InterceptEndpointGroupAssociationId: cmd.String("intercept-endpoint-group-association-id"),
								RequestId:                           cmd.String("request-id"),
							}

							op, err := client.CreateInterceptEndpointGroupAssociation(ctx, req)
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
						Usage: "update intercept-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group-association", Usage: "The ID of the intercept endpoint group association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "intercept_endpoint_group_association.name" not yet supported.
							intercept_endpoint_group_association_name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-endpoint-group-association"))
							fmt.Printf("Executing update on %s\n", intercept_endpoint_group_association_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete intercept-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group-association", Usage: "The ID of the intercept endpoint group association.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-endpoint-group-association"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInterceptEndpointGroupAssociation %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteInterceptEndpointGroupAssociationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInterceptEndpointGroupAssociation(ctx, req)
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
				Name:  "intercept-endpoint-groups",
				Usage: "Manage intercept-endpoint-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list intercept-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListInterceptEndpointGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListInterceptEndpointGroups(ctx, req)
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
						Usage: "describe intercept-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group", Usage: "The ID of the intercept endpoint group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-endpoint-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetInterceptEndpointGroupRequest{
								Name: name,
							}

							resp, err := client.GetInterceptEndpointGroup(ctx, req)
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
						Usage: "create intercept-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group-id", Usage: "The ID to use for the endpoint group, which will become the final.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateInterceptEndpointGroupRequest{
								Parent:                   parent,
								InterceptEndpointGroupId: cmd.String("intercept-endpoint-group-id"),
								RequestId:                cmd.String("request-id"),
							}

							op, err := client.CreateInterceptEndpointGroup(ctx, req)
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
						Usage: "update intercept-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group", Usage: "The ID of the intercept endpoint group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "intercept_endpoint_group.name" not yet supported.
							intercept_endpoint_group_name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-endpoint-group"))
							fmt.Printf("Executing update on %s\n", intercept_endpoint_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete intercept-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "intercept-endpoint-group", Usage: "The ID of the intercept endpoint group.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept-endpoint-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInterceptEndpointGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteInterceptEndpointGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteInterceptEndpointGroup(ctx, req)
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
				Name:  "mirroring-deployment-groups",
				Usage: "Manage mirroring-deployment-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list mirroring-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListMirroringDeploymentGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMirroringDeploymentGroups(ctx, req)
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
						Usage: "describe mirroring-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The ID of the mirroring deployment group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-deployment-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetMirroringDeploymentGroupRequest{
								Name: name,
							}

							resp, err := client.GetMirroringDeploymentGroup(ctx, req)
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
						Usage: "create mirroring-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment-group-id", Usage: "The ID to use for the new deployment group, which will become the.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateMirroringDeploymentGroupRequest{
								Parent:                     parent,
								MirroringDeploymentGroupId: cmd.String("mirroring-deployment-group-id"),
								RequestId:                  cmd.String("request-id"),
							}

							op, err := client.CreateMirroringDeploymentGroup(ctx, req)
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
						Usage: "update mirroring-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The ID of the mirroring deployment group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "mirroring_deployment_group.name" not yet supported.
							mirroring_deployment_group_name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-deployment-group"))
							fmt.Printf("Executing update on %s\n", mirroring_deployment_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete mirroring-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The ID of the mirroring deployment group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-deployment-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMirroringDeploymentGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteMirroringDeploymentGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteMirroringDeploymentGroup(ctx, req)
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
				Name:  "mirroring-deployments",
				Usage: "Manage mirroring-deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list mirroring-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListMirroringDeployments`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListMirroringDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMirroringDeployments(ctx, req)
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
						Usage: "describe mirroring-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment", Usage: "The ID of the mirroring deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetMirroringDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetMirroringDeployment(ctx, req)
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
						Usage: "create mirroring-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment-id", Usage: "The ID to use for the new deployment, which will become the final.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateMirroringDeploymentRequest{
								Parent:                parent,
								MirroringDeploymentId: cmd.String("mirroring-deployment-id"),
								RequestId:             cmd.String("request-id"),
							}

							op, err := client.CreateMirroringDeployment(ctx, req)
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
						Usage: "update mirroring-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment", Usage: "The ID of the mirroring deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "mirroring_deployment.name" not yet supported.
							mirroring_deployment_name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-deployment"))
							fmt.Printf("Executing update on %s\n", mirroring_deployment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete mirroring-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-deployment", Usage: "The ID of the mirroring deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMirroringDeployment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteMirroringDeploymentRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteMirroringDeployment(ctx, req)
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
				Name:  "mirroring-endpoint-group-associations",
				Usage: "Manage mirroring-endpoint-group-associations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list mirroring-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListMirroringEndpointGroupAssociationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMirroringEndpointGroupAssociations(ctx, req)
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
						Usage: "describe mirroring-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group-association", Usage: "The ID of the mirroring endpoint group association.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-endpoint-group-association"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetMirroringEndpointGroupAssociationRequest{
								Name: name,
							}

							resp, err := client.GetMirroringEndpointGroupAssociation(ctx, req)
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
						Usage: "create mirroring-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group-association-id", Usage: "The ID to use for the new association, which will become the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateMirroringEndpointGroupAssociationRequest{
								Parent:                              parent,
								MirroringEndpointGroupAssociationId: cmd.String("mirroring-endpoint-group-association-id"),
								RequestId:                           cmd.String("request-id"),
							}

							op, err := client.CreateMirroringEndpointGroupAssociation(ctx, req)
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
						Usage: "update mirroring-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group-association", Usage: "The ID of the mirroring endpoint group association.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "mirroring_endpoint_group_association.name" not yet supported.
							mirroring_endpoint_group_association_name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-endpoint-group-association"))
							fmt.Printf("Executing update on %s\n", mirroring_endpoint_group_association_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete mirroring-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group-association", Usage: "The ID of the mirroring endpoint group association.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-endpoint-group-association"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMirroringEndpointGroupAssociation %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteMirroringEndpointGroupAssociationRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteMirroringEndpointGroupAssociation(ctx, req)
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
				Name:  "mirroring-endpoint-groups",
				Usage: "Manage mirroring-endpoint-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list mirroring-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sort expression.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListMirroringEndpointGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListMirroringEndpointGroups(ctx, req)
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
						Usage: "describe mirroring-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group", Usage: "The ID of the mirroring endpoint group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-endpoint-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetMirroringEndpointGroupRequest{
								Name: name,
							}

							resp, err := client.GetMirroringEndpointGroup(ctx, req)
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
						Usage: "create mirroring-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group-id", Usage: "The ID to use for the endpoint group, which will become the final.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateMirroringEndpointGroupRequest{
								Parent:                   parent,
								MirroringEndpointGroupId: cmd.String("mirroring-endpoint-group-id"),
								RequestId:                cmd.String("request-id"),
							}

							op, err := client.CreateMirroringEndpointGroup(ctx, req)
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
						Usage: "update mirroring-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group", Usage: "The ID of the mirroring endpoint group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "mirroring_endpoint_group.name" not yet supported.
							mirroring_endpoint_group_name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-endpoint-group"))
							fmt.Printf("Executing update on %s\n", mirroring_endpoint_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete mirroring-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mirroring-endpoint-group", Usage: "The ID of the mirroring endpoint group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "A unique identifier for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring-endpoint-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteMirroringEndpointGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteMirroringEndpointGroupRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteMirroringEndpointGroup(ctx, req)
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
				},
			},
			{
				Name:  "rules",
				Usage: "Manage rules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of GatewaySecurityPolicyRules to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListGatewaySecurityPolicyRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListGatewaySecurityPolicyRules(ctx, req)
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
						Usage: "describe rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"), cmd.String("rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetGatewaySecurityPolicyRuleRequest{
								Name: name,
							}

							resp, err := client.GetGatewaySecurityPolicyRule(ctx, req)
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
						Usage: "create rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "gateway-security-policy-rule-id", Usage: "The ID to use for the rule, which will become the final component of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateGatewaySecurityPolicyRuleRequest{
								Parent:                      parent,
								GatewaySecurityPolicyRuleId: cmd.String("gateway-security-policy-rule-id"),
							}

							op, err := client.CreateGatewaySecurityPolicyRule(ctx, req)
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
						Usage: "update rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "gateway_security_policy_rule.name" not yet supported.
							gateway_security_policy_rule_name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"), cmd.String("rule"))
							fmt.Printf("Executing update on %s\n", gateway_security_policy_rule_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "gateway-security-policie", Usage: "The ID of the gateway security policie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway-security-policie"), cmd.String("rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteGatewaySecurityPolicyRule %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteGatewaySecurityPolicyRuleRequest{
								Name: name,
							}

							op, err := client.DeleteGatewaySecurityPolicyRule(ctx, req)
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
				Name:  "security-profile-groups",
				Usage: "Manage security-profile-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list security-profile-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of SecurityProfileGroups to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListSecurityProfileGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecurityProfileGroups(ctx, req)
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
						Usage: "describe security-profile-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile-group", Usage: "The ID of the security profile group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security-profile-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetSecurityProfileGroupRequest{
								Name: name,
							}

							resp, err := client.GetSecurityProfileGroup(ctx, req)
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
						Usage: "create security-profile-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile-group-id", Usage: "Short name of the SecurityProfileGroup resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateSecurityProfileGroupRequest{
								Parent:                 parent,
								SecurityProfileGroupId: cmd.String("security-profile-group-id"),
							}

							op, err := client.CreateSecurityProfileGroup(ctx, req)
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
						Usage: "update security-profile-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile-group", Usage: "The ID of the security profile group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "security_profile_group.name" not yet supported.
							security_profile_group_name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security-profile-group"))
							fmt.Printf("Executing update on %s\n", security_profile_group_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete security-profile-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If client provided etag is out of date, delete will return.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile-group", Usage: "The ID of the security profile group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security-profile-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSecurityProfileGroup %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteSecurityProfileGroupRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteSecurityProfileGroup(ctx, req)
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
				Name:  "security-profiles",
				Usage: "Manage security-profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list security-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of SecurityProfiles to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListSecurityProfilesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListSecurityProfiles(ctx, req)
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
						Usage: "describe security-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile", Usage: "The ID of the security profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security-profile"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetSecurityProfileRequest{
								Name: name,
							}

							resp, err := client.GetSecurityProfile(ctx, req)
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
						Usage: "create security-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile-id", Usage: "Short name of the SecurityProfile resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateSecurityProfileRequest{
								Parent:            parent,
								SecurityProfileId: cmd.String("security-profile-id"),
							}

							op, err := client.CreateSecurityProfile(ctx, req)
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
						Usage: "update security-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile", Usage: "The ID of the security profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "security_profile.name" not yet supported.
							security_profile_name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security-profile"))
							fmt.Printf("Executing update on %s\n", security_profile_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete security-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "If client provided etag is out of date, delete will return.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "security-profile", Usage: "The ID of the security profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security-profile"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteSecurityProfile %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteSecurityProfileRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeleteSecurityProfile(ctx, req)
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
				Name:  "server-tls-policies",
				Usage: "Manage server-tls-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list server-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of ServerTlsPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListServerTlsPoliciesResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Setting this field to `true` will opt the request into returning.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListServerTlsPoliciesRequest{
								Parent:               parent,
								PageSize:             int32(cmd.Int("page-size")),
								PageToken:            cmd.String("page-token"),
								ReturnPartialSuccess: cmd.Bool("return-partial-success"),
							}

							limit := cmd.Int("limit")
							it := client.ListServerTlsPolicies(ctx, req)
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
						Usage: "describe server-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "server-tls-policie", Usage: "The ID of the server tls policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server-tls-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetServerTlsPolicyRequest{
								Name: name,
							}

							resp, err := client.GetServerTlsPolicy(ctx, req)
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
						Usage: "create server-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "server-tls-policy-id", Usage: "Short name of the ServerTlsPolicy resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateServerTlsPolicyRequest{
								Parent:            parent,
								ServerTlsPolicyId: cmd.String("server-tls-policy-id"),
							}

							op, err := client.CreateServerTlsPolicy(ctx, req)
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
						Usage: "update server-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "server-tls-policie", Usage: "The ID of the server tls policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "server_tls_policy.name" not yet supported.
							server_tls_policy_name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server-tls-policie"))
							fmt.Printf("Executing update on %s\n", server_tls_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete server-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "server-tls-policie", Usage: "The ID of the server tls policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server-tls-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteServerTlsPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteServerTlsPolicyRequest{
								Name: name,
							}

							op, err := client.DeleteServerTlsPolicy(ctx, req)
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
				Name:  "tls-inspection-policies",
				Usage: "Manage tls-inspection-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list tls-inspection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of TlsInspectionPolicies to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListTlsInspectionPoliciesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListTlsInspectionPolicies(ctx, req)
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
						Usage: "describe tls-inspection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-inspection-policie", Usage: "The ID of the tls inspection policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls-inspection-policie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetTlsInspectionPolicyRequest{
								Name: name,
							}

							resp, err := client.GetTlsInspectionPolicy(ctx, req)
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
						Usage: "create tls-inspection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-inspection-policy-id", Usage: "Short name of the TlsInspectionPolicy resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateTlsInspectionPolicyRequest{
								Parent:                parent,
								TlsInspectionPolicyId: cmd.String("tls-inspection-policy-id"),
							}

							op, err := client.CreateTlsInspectionPolicy(ctx, req)
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
						Usage: "update tls-inspection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-inspection-policie", Usage: "The ID of the tls inspection policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "tls_inspection_policy.name" not yet supported.
							tls_inspection_policy_name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls-inspection-policie"))
							fmt.Printf("Executing update on %s\n", tls_inspection_policy_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tls-inspection-policies",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force", Usage: "If set to true, any rules for this TlsInspectionPolicy will also be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "tls-inspection-policie", Usage: "The ID of the tls inspection policie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls-inspection-policie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteTlsInspectionPolicy %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteTlsInspectionPolicyRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							op, err := client.DeleteTlsInspectionPolicy(ctx, req)
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
				Name:  "url-lists",
				Usage: "Manage url-lists resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list url-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of UrlLists to return per call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value returned by the last `ListUrlListsResponse`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.ListUrlListsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUrlLists(ctx, req)
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
						Usage: "describe url-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "url-list", Usage: "The ID of the url list.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url-list"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.GetUrlListRequest{
								Name: name,
							}

							resp, err := client.GetUrlList(ctx, req)
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
						Usage: "create url-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "url-list-id", Usage: "Short name of the UrlList resource to be created.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.CreateUrlListRequest{
								Parent:    parent,
								UrlListId: cmd.String("url-list-id"),
							}

							op, err := client.CreateUrlList(ctx, req)
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
						Usage: "update url-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "url-list", Usage: "The ID of the url list.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "url_list.name" not yet supported.
							url_list_name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url-list"))
							fmt.Printf("Executing update on %s\n", url_list_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete url-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "url-list", Usage: "The ID of the url list.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url-list"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteUrlList %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := networksecurity.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &networksecuritypb.DeleteUrlListRequest{
								Name: name,
							}

							op, err := client.DeleteUrlList(ctx, req)
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
