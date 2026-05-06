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

package appengine

import (
	appengine "cloud.google.com/go/appengine/apiv1"
	"cloud.google.com/go/appengine/apiv1/appenginepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the appengine command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "appengine",
		Usage: "manage App Engine Admin API resources",
		Commands: []*cli.Command{
			{
				Name:  "apps",
				Usage: "Manage apps resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetApplicationRequest{
								Name: name,
							}

							resp, err := client.GetApplication(ctx, req)
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
						Usage: "create apps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.CreateApplicationRequest{}

							op, err := client.CreateApplication(ctx, req)
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
						Usage: "update apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.UpdateApplicationRequest{
								Name: name,
							}

							op, err := client.UpdateApplication(ctx, req)
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
						Name:  "repair",
						Usage: "repair apps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.RepairApplicationRequest{
								Name: name,
							}

							op, err := client.RepairApplication(ctx, req)
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
				Name:  "authorized-certificates",
				Usage: "Manage authorized-certificates resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list authorized-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
							&cli.StringFlag{Name: "view", Usage: "Controls the set of fields returned in the `LIST` response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListAuthorizedCertificatesRequest{
								Parent:    parent,
								View:      appenginepb.AuthorizedCertificateView(appenginepb.AuthorizedCertificateView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAuthorizedCertificates(ctx, req)
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
						Usage: "describe authorized-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "authorized-certificate", Usage: "The ID of the authorized certificate.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Controls the set of fields returned in the `GET` response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/authorizedCertificates/%s", cmd.String("app"), cmd.String("authorized-certificate"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetAuthorizedCertificateRequest{
								Name: name,
								View: appenginepb.AuthorizedCertificateView(appenginepb.AuthorizedCertificateView_value[cmd.String("view")]),
							}

							resp, err := client.GetAuthorizedCertificate(ctx, req)
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
						Usage: "create authorized-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.CreateAuthorizedCertificateRequest{
								Parent: parent,
							}

							resp, err := client.CreateAuthorizedCertificate(ctx, req)
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
						Usage: "update authorized-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "authorized-certificate", Usage: "The ID of the authorized certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/authorizedCertificates/%s", cmd.String("app"), cmd.String("authorized-certificate"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.UpdateAuthorizedCertificateRequest{
								Name: name,
							}

							resp, err := client.UpdateAuthorizedCertificate(ctx, req)
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
						Usage: "delete authorized-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "authorized-certificate", Usage: "The ID of the authorized certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/authorizedCertificates/%s", cmd.String("app"), cmd.String("authorized-certificate"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAuthorizedCertificate on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DeleteAuthorizedCertificateRequest{
								Name: name,
							}

							if err := client.DeleteAuthorizedCertificate(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "authorized-domains",
				Usage: "Manage authorized-domains resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list authorized-domains",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListAuthorizedDomainsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAuthorizedDomains(ctx, req)
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
				Name:  "domain-mappings",
				Usage: "Manage domain-mappings resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list domain-mappings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListDomainMappingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDomainMappings(ctx, req)
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
						Usage: "describe domain-mappings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "domain-mapping", Usage: "The ID of the domain mapping.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/domainMappings/%s", cmd.String("app"), cmd.String("domain-mapping"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetDomainMappingRequest{
								Name: name,
							}

							resp, err := client.GetDomainMapping(ctx, req)
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
						Usage: "create domain-mappings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "override-strategy", Usage: "Whether the domain creation should override any existing mappings for this.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.CreateDomainMappingRequest{
								Parent:           parent,
								OverrideStrategy: appenginepb.DomainOverrideStrategy(appenginepb.DomainOverrideStrategy_value[cmd.String("override-strategy")]),
							}

							op, err := client.CreateDomainMapping(ctx, req)
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
						Usage: "update domain-mappings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "domain-mapping", Usage: "The ID of the domain mapping.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/domainMappings/%s", cmd.String("app"), cmd.String("domain-mapping"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.UpdateDomainMappingRequest{
								Name: name,
							}

							op, err := client.UpdateDomainMapping(ctx, req)
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
						Usage: "delete domain-mappings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "domain-mapping", Usage: "The ID of the domain mapping.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/domainMappings/%s", cmd.String("app"), cmd.String("domain-mapping"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDomainMapping %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DeleteDomainMappingRequest{
								Name: name,
							}

							op, err := client.DeleteDomainMapping(ctx, req)
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
				Name:  "ingress-rules",
				Usage: "Manage ingress-rules resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list ingress-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "matching-address", Usage: "A valid IP Address.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListIngressRulesRequest{
								Parent:          parent,
								PageSize:        int32(cmd.Int("page-size")),
								PageToken:       cmd.String("page-token"),
								MatchingAddress: cmd.String("matching-address"),
							}

							limit := cmd.Int("limit")
							it := client.ListIngressRules(ctx, req)
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
						Name:  "batch-update",
						Usage: "batch-update ingress-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/firewall/ingressRules", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.BatchUpdateIngressRulesRequest{
								Name: name,
							}

							resp, err := client.BatchUpdateIngressRules(ctx, req)
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
						Usage: "create ingress-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.CreateIngressRuleRequest{
								Parent: parent,
							}

							resp, err := client.CreateIngressRule(ctx, req)
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
						Usage: "describe ingress-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "ingress-rule", Usage: "The ID of the ingress rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/firewall/ingressRules/%s", cmd.String("app"), cmd.String("ingress-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetIngressRuleRequest{
								Name: name,
							}

							resp, err := client.GetIngressRule(ctx, req)
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
						Usage: "update ingress-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "ingress-rule", Usage: "The ID of the ingress rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/firewall/ingressRules/%s", cmd.String("app"), cmd.String("ingress-rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.UpdateIngressRuleRequest{
								Name: name,
							}

							resp, err := client.UpdateIngressRule(ctx, req)
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
						Usage: "delete ingress-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "ingress-rule", Usage: "The ID of the ingress rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/firewall/ingressRules/%s", cmd.String("app"), cmd.String("ingress-rule"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteIngressRule on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DeleteIngressRuleRequest{
								Name: name,
							}

							if err := client.DeleteIngressRule(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s/services/%s/versions/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListInstancesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListInstances(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s/versions/%s/instances/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetInstanceRequest{
								Name: name,
							}

							resp, err := client.GetInstance(ctx, req)
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
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s/versions/%s/instances/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"), cmd.String("instance"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteInstance %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DeleteInstanceRequest{
								Name: name,
							}

							op, err := client.DeleteInstance(ctx, req)
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
						Name:  "debug",
						Usage: "debug instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "ssh-key", Usage: "Public SSH key to add to the instance.", Required: false},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s/versions/%s/instances/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DebugInstanceRequest{
								Name:   name,
								SshKey: cmd.String("ssh-key"),
							}

							op, err := client.DebugInstance(ctx, req)
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
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s", cmd.String("app"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/operations/%s", cmd.String("app"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s", cmd.String("app"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListServicesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListServices(ctx, req)
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
						Usage: "describe services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s", cmd.String("app"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetServiceRequest{
								Name: name,
							}

							resp, err := client.GetService(ctx, req)
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
						Usage: "update services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.BoolFlag{Name: "migrate-traffic", Usage: "Set to `true` to gradually shift traffic to one or more versions that you.", Required: false},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s", cmd.String("app"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.UpdateServiceRequest{
								Name:           name,
								MigrateTraffic: cmd.Bool("migrate-traffic"),
							}

							op, err := client.UpdateService(ctx, req)
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
						Usage: "delete services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s", cmd.String("app"), cmd.String("service"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteService %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DeleteServiceRequest{
								Name: name,
							}

							op, err := client.DeleteService(ctx, req)
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
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum results to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Continuation token for fetching the next page of results.", Required: false},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Controls the set of fields returned in the `List` response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s/services/%s", cmd.String("app"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.ListVersionsRequest{
								Parent:    parent,
								View:      appenginepb.VersionView(appenginepb.VersionView_value[cmd.String("view")]),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListVersions(ctx, req)
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
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Controls the set of fields returned in the `Get` response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s/versions/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.GetVersionRequest{
								Name: name,
								View: appenginepb.VersionView(appenginepb.VersionView_value[cmd.String("view")]),
							}

							resp, err := client.GetVersion(ctx, req)
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
						Usage: "create versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("apps/%s/services/%s", cmd.String("app"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.CreateVersionRequest{
								Parent: parent,
							}

							op, err := client.CreateVersion(ctx, req)
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
						Usage: "update versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s/versions/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.UpdateVersionRequest{
								Name: name,
							}

							op, err := client.UpdateVersion(ctx, req)
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
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "app", Usage: "The ID of the app.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("apps/%s/services/%s/versions/%s", cmd.String("app"), cmd.String("service"), cmd.String("version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := appengine.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &appenginepb.DeleteVersionRequest{
								Name: name,
							}

							op, err := client.DeleteVersion(ctx, req)
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
