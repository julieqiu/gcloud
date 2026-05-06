package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	privateca "cloud.google.com/go/security/privateca/apiv1"
	"cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "privateca",
				Usage: "manage Certificate Authority API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "ca-pools",
						Usage: "Manage ca-pools resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create ca-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca-pool-id", Usage: "The ca pool id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.CreateCaPoolRequest{Parent: parent}
									req.CaPoolId = cmd.String("ca-pool-id")
									op, err := client.CreateCaPool(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update ca-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.UpdateCaPoolRequest{}
									req.CaPool = &privatecapb.CaPool{
										Name: name,
									}
									op, err := client.UpdateCaPool(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ca-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.GetCaPoolRequest{Name: name}
									resp, err := client.GetCaPool(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list ca-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete ca-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.DeleteCaPoolRequest{Name: name}
									op, err := client.DeleteCaPool(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "fetch-ca-certs",
								Usage: "fetch-ca-certs ca-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch-ca-certs...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy ca-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy ca-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions ca-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
									&cli.StringFlag{Name: "pem-ca-certificate", Usage: "The pem ca certificate.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.ActivateCertificateAuthorityRequest{Name: name}
									req.PemCaCertificate = cmd.String("pem-ca-certificate")
									op, err := client.ActivateCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate-authority-id", Usage: "The certificate authority id.", Required: true},
									&cli.StringFlag{Name: "gcs-bucket", Usage: "The gcs bucket.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.CreateCertificateAuthorityRequest{Parent: parent}
									req.CertificateAuthorityId = cmd.String("certificate-authority-id")
									req.CertificateAuthority = &privatecapb.CertificateAuthority{
										GcsBucket: cmd.String("gcs-bucket"),
									}
									op, err := client.CreateCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "disable",
								Usage: "disable certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
									&cli.BoolFlag{Name: "ignore-dependent-resources", Usage: "The ignore dependent resources.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.DisableCertificateAuthorityRequest{Name: name}
									req.IgnoreDependentResources = cmd.Bool("ignore-dependent-resources")
									op, err := client.DisableCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "enable",
								Usage: "enable certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.EnableCertificateAuthorityRequest{Name: name}
									op, err := client.EnableCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "fetch",
								Usage: "fetch certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									fmt.Printf("Executing fetch on %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.GetCertificateAuthorityRequest{Name: name}
									resp, err := client.GetCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &privatecapb.ListCertificateAuthoritiesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCertificateAuthorities(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "undelete",
								Usage: "undelete certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.UndeleteCertificateAuthorityRequest{Name: name}
									op, err := client.UndeleteCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.DeleteCertificateAuthorityRequest{Name: name}
									op, err := client.DeleteCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update certificate-authorities",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
									&cli.StringFlag{Name: "gcs-bucket", Usage: "The gcs bucket.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.UpdateCertificateAuthorityRequest{}
									req.CertificateAuthority = &privatecapb.CertificateAuthority{
										Name:      name,
										GcsBucket: cmd.String("gcs-bucket"),
									}
									var paths []string
									if cmd.IsSet("gcs-bucket") {
										paths = append(paths, "gcs_bucket")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateCertificateAuthority(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
									&cli.StringFlag{Name: "certificate_revocation_list", Usage: "The certificate_revocation_list.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s/certificateRevocationLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"), cmd.String("certificate_revocation_list"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.GetCertificateRevocationListRequest{Name: name}
									resp, err := client.GetCertificateRevocationList(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list certificate-revocation-lists",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &privatecapb.ListCertificateRevocationListsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCertificateRevocationLists(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update certificate-revocation-lists",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate_authority", Usage: "The certificate_authority.", Required: true},
									&cli.StringFlag{Name: "certificate_revocation_list", Usage: "The certificate_revocation_list.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificateAuthorities/%s/certificateRevocationLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate_authority"), cmd.String("certificate_revocation_list"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.UpdateCertificateRevocationListRequest{}
									req.CertificateRevocationList = &privatecapb.CertificateRevocationList{
										Name: name,
									}
									op, err := client.UpdateCertificateRevocationList(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "certificate-template-id", Usage: "The certificate template id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.CreateCertificateTemplateRequest{Parent: parent}
									req.CertificateTemplateId = cmd.String("certificate-template-id")
									req.CertificateTemplate = &privatecapb.CertificateTemplate{
										Description: cmd.String("description"),
									}
									op, err := client.CreateCertificateTemplate(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete certificate-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "certificate_template", Usage: "The certificate_template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_template"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.DeleteCertificateTemplateRequest{Name: name}
									op, err := client.DeleteCertificateTemplate(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe certificate-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "certificate_template", Usage: "The certificate_template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_template"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.GetCertificateTemplateRequest{Name: name}
									resp, err := client.GetCertificateTemplate(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list certificate-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update certificate-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "certificate_template", Usage: "The certificate_template.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_template"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.UpdateCertificateTemplateRequest{}
									req.CertificateTemplate = &privatecapb.CertificateTemplate{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateCertificateTemplate(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate-id", Usage: "The certificate id.", Required: false},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "issuing-certificate-authority-id", Usage: "The issuing certificate authority id.", Required: false},
									&cli.StringFlag{Name: "certificate-template", Usage: "The certificate template.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/caPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.CreateCertificateRequest{Parent: parent}
									req.CertificateId = cmd.String("certificate-id")
									req.ValidateOnly = cmd.Bool("validate-only")
									req.IssuingCertificateAuthorityId = cmd.String("issuing-certificate-authority-id")
									req.Certificate = &privatecapb.Certificate{
										CertificateTemplate: cmd.String("certificate-template"),
									}
									resp, err := client.CreateCertificate(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe certificates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate", Usage: "The certificate.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.GetCertificateRequest{Name: name}
									resp, err := client.GetCertificate(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list certificates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &privatecapb.ListCertificatesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCertificates(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "revoke",
								Usage: "revoke certificates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate", Usage: "The certificate.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate"))
									fmt.Printf("Executing revoke on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update certificates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ca_pool", Usage: "The ca_pool.", Required: true},
									&cli.StringFlag{Name: "certificate", Usage: "The certificate.", Required: true},
									&cli.StringFlag{Name: "certificate-template", Usage: "The certificate template.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/caPools/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("ca_pool"), cmd.String("certificate"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &privatecapb.UpdateCertificateRequest{}
									req.Certificate = &privatecapb.Certificate{
										Name:                name,
										CertificateTemplate: cmd.String("certificate-template"),
									}
									var paths []string
									if cmd.IsSet("certificate-template") {
										paths = append(paths, "certificate_template")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateCertificate(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
									limit := cmd.Int("limit")
									count := 0
									for {
										if limit > 0 && count >= limit {
											break
										}
										resp, err := it.Next()
										if err == iterator.Done {
											break
										}
										if err != nil {
											return err
										}
										out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
										if err != nil {
											return err
										}
										if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
											continue
										}
										if cmd.Bool("uri") {
											fmt.Println(resp.GetName())
										} else {
											fmt.Println(string(out))
										}
										count++
									}
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
									if err != nil {
										return err
									}
									out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
									if err != nil {
										return err
									}
									fmt.Println(string(out))
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := privateca.NewCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
