package certificatemanager

import (
	"context"
	"fmt"
	"strings"

	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
	"cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud certificatemanager command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "certificatemanager",
		Usage: "manage Certificate Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "certificate-issuance-configs",
				Usage: "Manage certificate-issuance-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list certificate-issuance-configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe certificate-issuance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_issuance_config", Usage: "The certificate_issuance_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateIssuanceConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_issuance_config"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.GetCertificateIssuanceConfigRequest{Name: name}
							resp, err := client.GetCertificateIssuanceConfig(ctx, req)
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
						Usage: "create certificate-issuance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate-issuance-config-id", Usage: "The certificate issuance config id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.IntFlag{Name: "rotation-window-percentage", Usage: "The rotation window percentage.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.CreateCertificateIssuanceConfigRequest{Parent: parent}
							req.CertificateIssuanceConfigId = cmd.String("certificate-issuance-config-id")
							req.CertificateIssuanceConfig = &certificatemanagerpb.CertificateIssuanceConfig{
								Name:                     cmd.String("name"),
								Description:              cmd.String("description"),
								RotationWindowPercentage: int32(cmd.Int("rotation-window-percentage")),
							}
							op, err := client.CreateCertificateIssuanceConfig(ctx, req)
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
						Usage: "delete certificate-issuance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_issuance_config", Usage: "The certificate_issuance_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateIssuanceConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_issuance_config"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.DeleteCertificateIssuanceConfigRequest{Name: name}
							op, err := client.DeleteCertificateIssuanceConfig(ctx, req)
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
				},
			},
			{
				Name:  "certificate-map-entries",
				Usage: "Manage certificate-map-entries resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list certificate-map-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &certificatemanagerpb.ListCertificateMapEntriesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListCertificateMapEntries(ctx, req)
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
						Usage: "describe certificate-map-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
							&cli.StringFlag{Name: "certificate_map_entry", Usage: "The certificate_map_entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s/certificateMapEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"), cmd.String("certificate_map_entry"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.GetCertificateMapEntryRequest{Name: name}
							resp, err := client.GetCertificateMapEntry(ctx, req)
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
						Usage: "create certificate-map-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
							&cli.StringFlag{Name: "certificate-map-entry-id", Usage: "The certificate map entry id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.CreateCertificateMapEntryRequest{Parent: parent}
							req.CertificateMapEntryId = cmd.String("certificate-map-entry-id")
							req.CertificateMapEntry = &certificatemanagerpb.CertificateMapEntry{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							op, err := client.CreateCertificateMapEntry(ctx, req)
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
						Usage: "update certificate-map-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
							&cli.StringFlag{Name: "certificate_map_entry", Usage: "The certificate_map_entry.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s/certificateMapEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"), cmd.String("certificate_map_entry"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.UpdateCertificateMapEntryRequest{}
							req.CertificateMapEntry = &certificatemanagerpb.CertificateMapEntry{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateCertificateMapEntry(ctx, req)
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
						Usage: "delete certificate-map-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
							&cli.StringFlag{Name: "certificate_map_entry", Usage: "The certificate_map_entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s/certificateMapEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"), cmd.String("certificate_map_entry"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.DeleteCertificateMapEntryRequest{Name: name}
							op, err := client.DeleteCertificateMapEntry(ctx, req)
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
				},
			},
			{
				Name:  "certificate-maps",
				Usage: "Manage certificate-maps resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list certificate-maps",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe certificate-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.GetCertificateMapRequest{Name: name}
							resp, err := client.GetCertificateMap(ctx, req)
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
						Usage: "create certificate-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate-map-id", Usage: "The certificate map id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.CreateCertificateMapRequest{Parent: parent}
							req.CertificateMapId = cmd.String("certificate-map-id")
							req.CertificateMap = &certificatemanagerpb.CertificateMap{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							op, err := client.CreateCertificateMap(ctx, req)
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
						Usage: "update certificate-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.UpdateCertificateMapRequest{}
							req.CertificateMap = &certificatemanagerpb.CertificateMap{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateCertificateMap(ctx, req)
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
						Usage: "delete certificate-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate_map", Usage: "The certificate_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificateMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate_map"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.DeleteCertificateMapRequest{Name: name}
							op, err := client.DeleteCertificateMap(ctx, req)
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
				},
			},
			{
				Name:  "certificates",
				Usage: "Manage certificates resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list certificates",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate", Usage: "The certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.GetCertificateRequest{Name: name}
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
						Name:  "create",
						Usage: "create certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate-id", Usage: "The certificate id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.CreateCertificateRequest{Parent: parent}
							req.CertificateId = cmd.String("certificate-id")
							req.Certificate = &certificatemanagerpb.Certificate{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							op, err := client.CreateCertificate(ctx, req)
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
						Usage: "update certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate", Usage: "The certificate.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.UpdateCertificateRequest{}
							req.Certificate = &certificatemanagerpb.Certificate{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateCertificate(ctx, req)
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
						Usage: "delete certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "certificate", Usage: "The certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/certificates/%s", cmd.String("project"), cmd.String("location"), cmd.String("certificate"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.DeleteCertificateRequest{Name: name}
							op, err := client.DeleteCertificate(ctx, req)
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
				},
			},
			{
				Name:  "dns-authorizations",
				Usage: "Manage dns-authorizations resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list dns-authorizations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe dns-authorizations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dns_authorization", Usage: "The dns_authorization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsAuthorizations/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_authorization"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.GetDnsAuthorizationRequest{Name: name}
							resp, err := client.GetDnsAuthorization(ctx, req)
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
						Usage: "create dns-authorizations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dns-authorization-id", Usage: "The dns authorization id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "domain", Usage: "The domain.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.CreateDnsAuthorizationRequest{Parent: parent}
							req.DnsAuthorizationId = cmd.String("dns-authorization-id")
							req.DnsAuthorization = &certificatemanagerpb.DnsAuthorization{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Domain:      cmd.String("domain"),
							}
							op, err := client.CreateDnsAuthorization(ctx, req)
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
						Usage: "update dns-authorizations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dns_authorization", Usage: "The dns_authorization.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "domain", Usage: "The domain.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsAuthorizations/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_authorization"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.UpdateDnsAuthorizationRequest{}
							req.DnsAuthorization = &certificatemanagerpb.DnsAuthorization{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Domain:      cmd.String("domain"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("domain") {
								paths = append(paths, "domain")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateDnsAuthorization(ctx, req)
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
						Usage: "delete dns-authorizations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dns_authorization", Usage: "The dns_authorization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsAuthorizations/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_authorization"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.DeleteDnsAuthorizationRequest{Name: name}
							op, err := client.DeleteDnsAuthorization(ctx, req)
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
							client, err := certificatemanager.NewClient(ctx)
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
							client, err := certificatemanager.NewClient(ctx)
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
							client, err := certificatemanager.NewClient(ctx)
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
							client, err := certificatemanager.NewClient(ctx)
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
							client, err := certificatemanager.NewClient(ctx)
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
							client, err := certificatemanager.NewClient(ctx)
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
			{
				Name:  "trust-configs",
				Usage: "Manage trust-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list trust-configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe trust-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "trust_config", Usage: "The trust_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/trustConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("trust_config"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.GetTrustConfigRequest{Name: name}
							resp, err := client.GetTrustConfig(ctx, req)
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
						Usage: "create trust-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "trust-config-id", Usage: "The trust config id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.CreateTrustConfigRequest{Parent: parent}
							req.TrustConfigId = cmd.String("trust-config-id")
							req.TrustConfig = &certificatemanagerpb.TrustConfig{
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Etag:        cmd.String("etag"),
							}
							op, err := client.CreateTrustConfig(ctx, req)
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
						Usage: "update trust-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "trust_config", Usage: "The trust_config.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/trustConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("trust_config"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.UpdateTrustConfigRequest{}
							req.TrustConfig = &certificatemanagerpb.TrustConfig{
								Name:        name,
								Name:        cmd.String("name"),
								Description: cmd.String("description"),
								Etag:        cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateTrustConfig(ctx, req)
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
						Usage: "delete trust-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "trust_config", Usage: "The trust_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/trustConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("trust_config"))
							client, err := certificatemanager.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &certificatemanagerpb.DeleteTrustConfigRequest{Name: name}
							op, err := client.DeleteTrustConfig(ctx, req)
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
				},
			},
		},
	}
}
