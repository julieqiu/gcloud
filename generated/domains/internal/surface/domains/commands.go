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

package domains

import (
	domains "cloud.google.com/go/domains/apiv1"
	"cloud.google.com/go/domains/apiv1/domainspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the domains command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "domains",
		Usage: "manage Cloud Domains API resources",
		Commands: []*cli.Command{
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
				},
			},
			{
				Name:  "registrations",
				Usage: "Manage registrations resources",
				Commands: []*cli.Command{

					{
						Name:  "search-domains",
						Usage: "search-domains registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "String used to search for available domain names.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing search-domains on %s\n", location)
							return nil
						},
					},

					{
						Name:  "retrieve-register-parameters",
						Usage: "retrieve-register-parameters registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain-name", Usage: "The domain name.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing retrieve-register-parameters on %s\n", location)
							return nil
						},
					},

					{
						Name:  "register",
						Usage: "register registrations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "contact-notices", Usage: "The list of contact notices that the caller acknowledges.", Required: false},
							&cli.StringSliceFlag{Name: "domain-notices", Usage: "The list of domain notices that you acknowledge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "When true, only validation is performed, without actually registering.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := domains.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &domainspb.RegisterDomainRequest{
								Parent:         parent,
								DomainNotices:  cmd.StringSlice("domain-notices"),
								ContactNotices: cmd.StringSlice("contact-notices"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.RegisterDomain(ctx, req)
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
						Name:  "retrieve-transfer-parameters",
						Usage: "retrieve-transfer-parameters registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "domain-name", Usage: "The domain name.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing retrieve-transfer-parameters on %s\n", location)
							return nil
						},
					},

					{
						Name:  "transfer",
						Usage: "transfer registrations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "contact-notices", Usage: "The list of contact notices that you acknowledge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validate the request without actually transferring the domain.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := domains.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &domainspb.TransferDomainRequest{
								Parent:         parent,
								ContactNotices: cmd.StringSlice("contact-notices"),
								ValidateOnly:   cmd.Bool("validate-only"),
							}

							op, err := client.TransferDomain(ctx, req)
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
						Name:  "list",
						Usage: "list registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression to restrict the `Registration`s returned.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of results to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "When set to the `next_page_token` from a prior response, provides the next.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := domains.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &domainspb.ListRegistrationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListRegistrations(ctx, req)
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
						Usage: "describe registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := domains.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &domainspb.GetRegistrationRequest{
								Name: name,
							}

							resp, err := client.GetRegistration(ctx, req)
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
						Usage: "update registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "registration.name" not yet supported.
							registration_name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing update on %s\n", registration_name)
							return nil
						},
					},

					{
						Name:  "configure-management-settings",
						Usage: "configure-management-settings registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							registration := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing configure-management-settings on %s\n", registration)
							return nil
						},
					},

					{
						Name:  "configure-dns-settings",
						Usage: "configure-dns-settings registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validate the request without actually updating the DNS settings.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							registration := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing configure-dns-settings on %s\n", registration)
							return nil
						},
					},

					{
						Name:  "configure-contact-settings",
						Usage: "configure-contact-settings registrations",
						Flags: []cli.Flag{
							&cli.StringSliceFlag{Name: "contact-notices", Usage: "The list of contact notices that the caller acknowledges.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Validate the request without actually updating the contact settings.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							registration := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing configure-contact-settings on %s\n", registration)
							return nil
						},
					},

					{
						Name:  "export",
						Usage: "export registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := domains.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &domainspb.ExportRegistrationRequest{
								Name: name,
							}

							op, err := client.ExportRegistration(ctx, req)
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
						Usage: "delete registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteRegistration %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := domains.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &domainspb.DeleteRegistrationRequest{
								Name: name,
							}

							op, err := client.DeleteRegistration(ctx, req)
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
						Name:  "retrieve-authorization-code",
						Usage: "retrieve-authorization-code registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							registration := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing retrieve-authorization-code on %s\n", registration)
							return nil
						},
					},

					{
						Name:  "reset-authorization-code",
						Usage: "reset-authorization-code registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The ID of the registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							registration := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing reset-authorization-code on %s\n", registration)
							return nil
						},
					},
				},
			},
		},
	}
}
