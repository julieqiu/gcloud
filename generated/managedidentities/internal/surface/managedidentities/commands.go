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

package managedidentities

import (
	managedidentities "cloud.google.com/go/managedidentities/apiv1"
	"cloud.google.com/go/managedidentities/apiv1/managedidentitiespb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the managedidentities command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "managedidentities",
		Usage: "manage Managed Service for Microsoft Active Directory API resources",
		Commands: []*cli.Command{
			{
				Name:  "domains",
				Usage: "Manage domains resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain-name", Usage: "The fully qualified domain name.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.CreateMicrosoftAdDomainRequest{
								Parent:     parent,
								DomainName: cmd.String("domain-name"),
							}

							op, err := client.CreateMicrosoftAdDomain(ctx, req)
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
						Name:  "reset-admin-password",
						Usage: "reset-admin-password domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.ResetAdminPasswordRequest{
								Name: name,
							}

							resp, err := client.ResetAdminPassword(ctx, req)
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
						Usage: "list domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter specifying constraints of a list operation.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Specifies the ordering of results.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The `next_page_token` value returned from a previous ListDomainsRequest.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/global", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.ListDomainsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDomains(ctx, req)
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
						Usage: "describe domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.GetDomainRequest{
								Name: name,
							}

							resp, err := client.GetDomain(ctx, req)
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
						Usage: "update domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "domain.name" not yet supported.
							domain_name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							fmt.Printf("Executing update on %s\n", domain_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDomain %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.DeleteDomainRequest{
								Name: name,
							}

							op, err := client.DeleteDomain(ctx, req)
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
						Name:  "attach-trust",
						Usage: "attach-trust domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.AttachTrustRequest{
								Name: name,
							}

							op, err := client.AttachTrust(ctx, req)
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
						Name:  "reconfigure-trust",
						Usage: "reconfigure-trust domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "target-dns-ip-addresses", Usage: "The target DNS server IP addresses to resolve the remote domain involved.", Required: true},
							&cli.StringFlag{Name: "target-domain-name", Usage: "The fully-qualified target domain name which will be in trust with current.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.ReconfigureTrustRequest{
								Name:                 name,
								TargetDomainName:     cmd.String("target-domain-name"),
								TargetDnsIpAddresses: cmd.StringSlice("target-dns-ip-addresses"),
							}

							op, err := client.ReconfigureTrust(ctx, req)
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
						Name:  "detach-trust",
						Usage: "detach-trust domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.DetachTrustRequest{
								Name: name,
							}

							op, err := client.DetachTrust(ctx, req)
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
						Name:  "validate-trust",
						Usage: "validate-trust domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain", Usage: "The ID of the domain.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/domains/%s", cmd.String("project"), cmd.String("domain"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := managedidentities.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &managedidentitiespb.ValidateTrustRequest{
								Name: name,
							}

							op, err := client.ValidateTrust(ctx, req)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
