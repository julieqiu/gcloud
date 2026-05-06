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

package privateca

import (
	privateca "cloud.google.com/go/privateca/apiv1"
	"cloud.google.com/go/privateca/apiv1/privatecapb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the privateca command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "privateca",
		Usage: "manage Certificate Authority API resources",
		Commands: []*cli.Command{
			{
				Name:  "ca-pools",
				Usage: "Manage ca-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool-id", Usage: "It must be unique within a location and match the regular.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.CreateCaPoolRequest{
								Parent:    parent,
								CaPoolId:  cmd.String("ca-pool-id"),
								RequestId: cmd.String("request-id"),
							}

							op, err := client.CreateCaPool(ctx, req)
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
						Usage: "update ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "ca_pool.name" not yet supported.
							ca_pool_name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							fmt.Printf("Executing update on %s\n", ca_pool_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.GetCaPoolRequest{
								Name: name,
							}

							resp, err := client.GetCaPool(ctx, req)
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
						Usage: "list ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.ListCaPoolsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCaPools(ctx, req)
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
						Usage: "delete ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.BoolFlag{Name: "ignore-dependent-resources", Usage: "This field allows this pool to be deleted even if it's being.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCaPool %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.DeleteCaPoolRequest{
								Name:                     name,
								RequestId:                cmd.String("request-id"),
								IgnoreDependentResources: cmd.Bool("ignore-dependent-resources"),
							}

							op, err := client.DeleteCaPool(ctx, req)
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
						Name:  "fetch-ca-certs",
						Usage: "fetch-ca-certs ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							ca_pool := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							fmt.Printf("Executing fetch-ca-certs on %s\n", ca_pool)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions ca-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},
				},
			},
			{
				Name:  "certificate-authorities",
				Usage: "Manage certificate-authorities resources",
				Commands: []*cli.Command{

					{
						Name:  "activate",
						Usage: "activate certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pem-ca-certificate", Usage: "The signed CA certificate issued from.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.ActivateCertificateAuthorityRequest{
								Name:             name,
								PemCaCertificate: cmd.String("pem-ca-certificate"),
								RequestId:        cmd.String("request-id"),
							}

							op, err := client.ActivateCertificateAuthority(ctx, req)
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
						Name:  "create",
						Usage: "create certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authority-id", Usage: "It must be unique within a location and match the regular.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.CreateCertificateAuthorityRequest{
								Parent:                 parent,
								CertificateAuthorityId: cmd.String("certificate-authority-id"),
								RequestId:              cmd.String("request-id"),
							}

							op, err := client.CreateCertificateAuthority(ctx, req)
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
						Name:  "disable",
						Usage: "disable certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.BoolFlag{Name: "ignore-dependent-resources", Usage: "This field allows this CA to be disabled even if it's being.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.DisableCertificateAuthorityRequest{
								Name:                     name,
								RequestId:                cmd.String("request-id"),
								IgnoreDependentResources: cmd.Bool("ignore-dependent-resources"),
							}

							op, err := client.DisableCertificateAuthority(ctx, req)
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
						Name:  "enable",
						Usage: "enable certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.EnableCertificateAuthorityRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.EnableCertificateAuthority(ctx, req)
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
						Name:  "fetch",
						Usage: "fetch certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.FetchCertificateAuthorityCsrRequest{
								Name: name,
							}

							resp, err := client.FetchCertificateAuthorityCsr(ctx, req)
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
						Usage: "describe certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.GetCertificateAuthorityRequest{
								Name: name,
							}

							resp, err := client.GetCertificateAuthority(ctx, req)
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
						Usage: "list certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.ListCertificateAuthoritiesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCertificateAuthorities(ctx, req)
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
						Name:  "undelete",
						Usage: "undelete certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.UndeleteCertificateAuthorityRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.UndeleteCertificateAuthority(ctx, req)
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
						Usage: "delete certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.BoolFlag{Name: "ignore-active-certificates", Usage: "This field allows the CA to be deleted even if the CA has.", Required: false},
							&cli.BoolFlag{Name: "ignore-dependent-resources", Usage: "This field allows this CA to be deleted even if it's being.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "skip-grace-period", Usage: "If this flag is set, the Certificate Authority will be deleted as.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.DeleteCertificateAuthorityRequest{
								Name:                     name,
								RequestId:                cmd.String("request-id"),
								IgnoreActiveCertificates: cmd.Bool("ignore-active-certificates"),
								SkipGracePeriod:          cmd.Bool("skip-grace-period"),
								IgnoreDependentResources: cmd.Bool("ignore-dependent-resources"),
							}

							op, err := client.DeleteCertificateAuthority(ctx, req)
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
						Usage: "update certificate-authorities",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "certificate_authority.name" not yet supported.
							certificate_authority_name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							fmt.Printf("Executing update on %s\n", certificate_authority_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "certificate-revocation-lists",
				Usage: "Manage certificate-revocation-lists resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe certificate-revocation-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "certificate-revocation-list", Usage: "The ID of the certificate revocation list.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s/certificateRevocationLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"), cmd.String("certificate-revocation-list"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.GetCertificateRevocationListRequest{
								Name: name,
							}

							resp, err := client.GetCertificateRevocationList(ctx, req)
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
						Usage: "list certificate-revocation-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.ListCertificateRevocationListsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCertificateRevocationLists(ctx, req)
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
						Usage: "update certificate-revocation-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-authoritie", Usage: "The ID of the certificate authoritie.", Required: true},
							&cli.StringFlag{Name: "certificate-revocation-list", Usage: "The ID of the certificate revocation list.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "certificate_revocation_list.name" not yet supported.
							certificate_revocation_list_name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s/certificateRevocationLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate-authoritie"), cmd.String("certificate-revocation-list"))
							fmt.Printf("Executing update on %s\n", certificate_revocation_list_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "certificate-templates",
				Usage: "Manage certificate-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create certificate-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "certificate-template-id", Usage: "It must be unique within a location and match the regular.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.CreateCertificateTemplateRequest{
								Parent:                parent,
								CertificateTemplateId: cmd.String("certificate-template-id"),
								RequestId:             cmd.String("request-id"),
							}

							op, err := client.CreateCertificateTemplate(ctx, req)
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
						Usage: "delete certificate-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "certificate-template", Usage: "The ID of the certificate template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate-template"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCertificateTemplate %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.DeleteCertificateTemplateRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteCertificateTemplate(ctx, req)
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
						Usage: "describe certificate-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "certificate-template", Usage: "The ID of the certificate template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.GetCertificateTemplateRequest{
								Name: name,
							}

							resp, err := client.GetCertificateTemplate(ctx, req)
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
						Usage: "list certificate-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.ListCertificateTemplatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCertificateTemplates(ctx, req)
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
						Usage: "update certificate-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "certificate-template", Usage: "The ID of the certificate template.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "certificate_template.name" not yet supported.
							certificate_template_name := fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate-template"))
							fmt.Printf("Executing update on %s\n", certificate_template_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "certificates",
				Usage: "Manage certificates resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate-id", Usage: "It must be unique within a location and match the regular.", Required: false},
							&cli.StringFlag{Name: "issuing-certificate-authority-id", Usage: "The resource ID of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If this is true, no.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.CreateCertificateRequest{
								Parent:                        parent,
								CertificateId:                 cmd.String("certificate-id"),
								RequestId:                     cmd.String("request-id"),
								ValidateOnly:                  cmd.Bool("validate-only"),
								IssuingCertificateAuthorityId: cmd.String("issuing-certificate-authority-id"),
							}

							resp, err := client.CreateCertificate(ctx, req)
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
						Usage: "describe certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate", Usage: "The ID of the certificate.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.GetCertificateRequest{
								Name: name,
							}

							resp, err := client.GetCertificate(ctx, req)
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
						Usage: "list certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.ListCertificatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCertificates(ctx, req)
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
						Name:  "revoke",
						Usage: "revoke certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate", Usage: "The ID of the certificate.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reason", Usage: "The.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := privateca.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &privatecapb.RevokeCertificateRequest{
								Name:      name,
								Reason:    privatecapb.RevocationReason(privatecapb.RevocationReason_value[cmd.String("reason")]),
								RequestId: cmd.String("request-id"),
							}

							resp, err := client.RevokeCertificate(ctx, req)
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
						Usage: "update certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ca-pool", Usage: "The ID of the ca pool.", Required: true},
							&cli.StringFlag{Name: "certificate", Usage: "The ID of the certificate.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "certificate.name" not yet supported.
							certificate_name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca-pool"), cmd.String("certificate"))
							fmt.Printf("Executing update on %s\n", certificate_name)
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
		},
	}
}
