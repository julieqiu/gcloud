package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
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
				Name:  "cloudkms",
				Usage: "manage Cloud Key Management Service (KMS) API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "autokey-config",
						Usage: "Manage autokey-config resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update autokey-config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
									&cli.StringFlag{Name: "key-project", Usage: "The key project.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s/autokeyConfig", cmd.String("folder"))
									client, err := kms.NewAutokeyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.UpdateAutokeyConfigRequest{}
									req.AutokeyConfig = &kmspb.AutokeyConfig{
										Name:       name,
										KeyProject: cmd.String("key-project"),
										Etag:       cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("key-project") {
										paths = append(paths, "key_project")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateAutokeyConfig(ctx, req)
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
								Usage: "describe autokey-config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s/autokeyConfig", cmd.String("folder"))
									client, err := kms.NewAutokeyAdminClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetAutokeyConfigRequest{Name: name}
									resp, err := client.GetAutokeyConfig(ctx, req)
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
						Name:  "crypto-key-versions",
						Usage: "Manage crypto-key-versions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &kmspb.ListCryptoKeyVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCryptoKeyVersions(ctx, req)
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
								Usage: "describe crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetCryptoKeyVersionRequest{Name: name}
									resp, err := client.GetCryptoKeyVersion(ctx, req)
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
								Usage: "create crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateCryptoKeyVersionRequest{Parent: parent}
									resp, err := client.CreateCryptoKeyVersion(ctx, req)
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
								Usage: "delete crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.DeleteCryptoKeyVersionRequest{Name: name}
									op, err := client.DeleteCryptoKeyVersion(ctx, req)
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
								Name:  "import",
								Usage: "import crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"))
									fmt.Printf("Executing import on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.UpdateCryptoKeyVersionRequest{}
									req.CryptoKeyVersion = &kmspb.CryptoKeyVersion{
										Name: name,
									}
									resp, err := client.UpdateCryptoKeyVersion(ctx, req)
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
								Name:  "destroy",
								Usage: "destroy crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing destroy on %s\n", name)
									return nil
								},
							},
							{
								Name:  "restore",
								Usage: "restore crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing restore on %s\n", name)
									return nil
								},
							},
							{
								Name:  "raw-encrypt",
								Usage: "raw-encrypt crypto-key-versions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing raw-encrypt...")
									return nil
								},
							},
							{
								Name:  "raw-decrypt",
								Usage: "raw-decrypt crypto-key-versions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing raw-decrypt...")
									return nil
								},
							},
							{
								Name:  "asymmetric-sign",
								Usage: "asymmetric-sign crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing asymmetric-sign on %s\n", name)
									return nil
								},
							},
							{
								Name:  "asymmetric-decrypt",
								Usage: "asymmetric-decrypt crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing asymmetric-decrypt on %s\n", name)
									return nil
								},
							},
							{
								Name:  "mac-sign",
								Usage: "mac-sign crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing mac-sign on %s\n", name)
									return nil
								},
							},
							{
								Name:  "mac-verify",
								Usage: "mac-verify crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing mac-verify on %s\n", name)
									return nil
								},
							},
							{
								Name:  "decapsulate",
								Usage: "decapsulate crypto-key-versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									fmt.Printf("Executing decapsulate on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "crypto-keys",
						Usage: "Manage crypto-keys resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &kmspb.ListCryptoKeysRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListCryptoKeys(ctx, req)
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
								Usage: "describe crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetCryptoKeyRequest{Name: name}
									resp, err := client.GetCryptoKey(ctx, req)
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
								Usage: "create crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto-key-id", Usage: "The crypto key id.", Required: true},
									&cli.BoolFlag{Name: "skip-initial-version-creation", Usage: "The skip initial version creation.", Required: false},
									&cli.BoolFlag{Name: "import-only", Usage: "The import only.", Required: false},
									&cli.StringFlag{Name: "crypto-key-backend", Usage: "The crypto key backend.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateCryptoKeyRequest{Parent: parent}
									req.CryptoKeyId = cmd.String("crypto-key-id")
									req.SkipInitialVersionCreation = cmd.Bool("skip-initial-version-creation")
									req.CryptoKey = &kmspb.CryptoKey{
										ImportOnly:       cmd.Bool("import-only"),
										CryptoKeyBackend: cmd.String("crypto-key-backend"),
									}
									resp, err := client.CreateCryptoKey(ctx, req)
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
								Usage: "delete crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.DeleteCryptoKeyRequest{Name: name}
									op, err := client.DeleteCryptoKey(ctx, req)
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
								Name:  "update",
								Usage: "update crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.BoolFlag{Name: "import-only", Usage: "The import only.", Required: false},
									&cli.StringFlag{Name: "crypto-key-backend", Usage: "The crypto key backend.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.UpdateCryptoKeyRequest{}
									req.CryptoKey = &kmspb.CryptoKey{
										Name:             name,
										ImportOnly:       cmd.Bool("import-only"),
										CryptoKeyBackend: cmd.String("crypto-key-backend"),
									}
									var paths []string
									if cmd.IsSet("import-only") {
										paths = append(paths, "import_only")
									}
									if cmd.IsSet("crypto-key-backend") {
										paths = append(paths, "crypto_key_backend")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateCryptoKey(ctx, req)
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
								Name:  "update-primary-version",
								Usage: "update-primary-version crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
									fmt.Printf("Executing update-primary-version on %s\n", name)
									return nil
								},
							},
							{
								Name:  "encrypt",
								Usage: "encrypt crypto-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing encrypt...")
									return nil
								},
							},
							{
								Name:  "decrypt",
								Usage: "decrypt crypto-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
									fmt.Printf("Executing decrypt on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "ekm-config",
						Usage: "Manage ekm-config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe ekm-config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/ekmConfig", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewEkmClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetEkmConfigRequest{Name: name}
									resp, err := client.GetEkmConfig(ctx, req)
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
								Usage: "update ekm-config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "default-ekm-connection", Usage: "The default ekm connection.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/ekmConfig", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewEkmClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.UpdateEkmConfigRequest{}
									req.EkmConfig = &kmspb.EkmConfig{
										Name:                 name,
										DefaultEkmConnection: cmd.String("default-ekm-connection"),
									}
									var paths []string
									if cmd.IsSet("default-ekm-connection") {
										paths = append(paths, "default_ekm_connection")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateEkmConfig(ctx, req)
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
						Name:  "ekm-connections",
						Usage: "Manage ekm-connections resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list ekm-connections",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ekm-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ekm_connection", Usage: "The ekm_connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("ekm_connection"))
									client, err := kms.NewEkmClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetEkmConnectionRequest{Name: name}
									resp, err := client.GetEkmConnection(ctx, req)
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
								Usage: "create ekm-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ekm-connection-id", Usage: "The ekm connection id.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "crypto-space-path", Usage: "The crypto space path.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewEkmClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateEkmConnectionRequest{Parent: parent}
									req.EkmConnectionId = cmd.String("ekm-connection-id")
									req.EkmConnection = &kmspb.EkmConnection{
										Etag:            cmd.String("etag"),
										CryptoSpacePath: cmd.String("crypto-space-path"),
									}
									resp, err := client.CreateEkmConnection(ctx, req)
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
								Usage: "update ekm-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ekm_connection", Usage: "The ekm_connection.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
									&cli.StringFlag{Name: "crypto-space-path", Usage: "The crypto space path.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("ekm_connection"))
									client, err := kms.NewEkmClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.UpdateEkmConnectionRequest{}
									req.EkmConnection = &kmspb.EkmConnection{
										Name:            name,
										Etag:            cmd.String("etag"),
										CryptoSpacePath: cmd.String("crypto-space-path"),
									}
									var paths []string
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									if cmd.IsSet("crypto-space-path") {
										paths = append(paths, "crypto_space_path")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateEkmConnection(ctx, req)
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
								Name:  "verify-connectivity",
								Usage: "verify-connectivity ekm-connections",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "ekm_connection", Usage: "The ekm_connection.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("ekm_connection"))
									fmt.Printf("Executing verify-connectivity on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "import-jobs",
						Usage: "Manage import-jobs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list import-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &kmspb.ListImportJobsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListImportJobs(ctx, req)
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
								Usage: "describe import-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "import_job", Usage: "The import_job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/importJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("import_job"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetImportJobRequest{Name: name}
									resp, err := client.GetImportJob(ctx, req)
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
								Usage: "create import-jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "import-job-id", Usage: "The import job id.", Required: true},
									&cli.StringFlag{Name: "crypto-key-backend", Usage: "The crypto key backend.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateImportJobRequest{Parent: parent}
									req.ImportJobId = cmd.String("import-job-id")
									req.ImportJob = &kmspb.ImportJob{
										CryptoKeyBackend: cmd.String("crypto-key-backend"),
									}
									resp, err := client.CreateImportJob(ctx, req)
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
						Name:  "key-handles",
						Usage: "Manage key-handles resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create key-handles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key-handle-id", Usage: "The key handle id.", Required: false},
									&cli.StringFlag{Name: "resource-type-selector", Usage: "The resource type selector.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewAutokeyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateKeyHandleRequest{Parent: parent}
									req.KeyHandleId = cmd.String("key-handle-id")
									req.KeyHandle = &kmspb.KeyHandle{
										ResourceTypeSelector: cmd.String("resource-type-selector"),
									}
									op, err := client.CreateKeyHandle(ctx, req)
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
								Usage: "describe key-handles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_handle", Usage: "The key_handle.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyHandles/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_handle"))
									client, err := kms.NewAutokeyClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetKeyHandleRequest{Name: name}
									resp, err := client.GetKeyHandle(ctx, req)
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
								Usage: "list key-handles",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "key-rings",
						Usage: "Manage key-rings resources",
						Commands: []*cli.Command{
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe key-rings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetKeyRingRequest{Name: name}
									resp, err := client.GetKeyRing(ctx, req)
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
								Usage: "create key-rings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key-ring-id", Usage: "The key ring id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateKeyRingRequest{Parent: parent}
									req.KeyRingId = cmd.String("key-ring-id")
									resp, err := client.CreateKeyRing(ctx, req)
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
								Name:  "set-iam-policy",
								Usage: "set-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions key-rings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
									client, err := kms.NewAutokeyClient(ctx)
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
									client, err := kms.NewAutokeyClient(ctx)
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
									client, err := kms.NewAutokeyAdminClient(ctx)
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
									client, err := kms.NewAutokeyAdminClient(ctx)
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
									client, err := kms.NewEkmClient(ctx)
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
									client, err := kms.NewEkmClient(ctx)
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
									client, err := kms.NewHsmManagementClient(ctx)
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
									client, err := kms.NewHsmManagementClient(ctx)
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
							{
								Name:  "generate-random-bytes",
								Usage: "generate-random-bytes locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-random-bytes...")
									return nil
								},
							},
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
									client, err := kms.NewKeyManagementClient(ctx)
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
									client, err := kms.NewKeyManagementClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := kms.NewAutokeyClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := kms.NewAutokeyAdminClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := kms.NewEkmClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := kms.NewHsmManagementClient(ctx)
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
								Name:  "describe",
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := kms.NewKeyManagementClient(ctx)
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
						},
					},
					{
						Name:  "projects",
						Usage: "Manage projects resources",
						Commands: []*cli.Command{
							{
								Name:  "show-effective-autokey-config",
								Usage: "show-effective-autokey-config projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing show-effective-autokey-config...")
									return nil
								},
							},
						},
					},
					{
						Name:  "proposals",
						Usage: "Manage proposals resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create proposals",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single_tenant_hsm_instance", Usage: "The single_tenant_hsm_instance.", Required: true},
									&cli.StringFlag{Name: "single-tenant-hsm-instance-proposal-id", Usage: "The single tenant hsm instance proposal id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s", cmd.String("project"), cmd.String("location"), cmd.String("single_tenant_hsm_instance"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateSingleTenantHsmInstanceProposalRequest{Parent: parent}
									req.SingleTenantHsmInstanceProposalId = cmd.String("single-tenant-hsm-instance-proposal-id")
									op, err := client.CreateSingleTenantHsmInstanceProposal(ctx, req)
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
								Name:  "approve",
								Usage: "approve proposals",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single_tenant_hsm_instance", Usage: "The single_tenant_hsm_instance.", Required: true},
									&cli.StringFlag{Name: "proposal", Usage: "The proposal.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single_tenant_hsm_instance"), cmd.String("proposal"))
									fmt.Printf("Executing approve on %s\n", name)
									return nil
								},
							},
							{
								Name:  "execute",
								Usage: "execute proposals",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single_tenant_hsm_instance", Usage: "The single_tenant_hsm_instance.", Required: true},
									&cli.StringFlag{Name: "proposal", Usage: "The proposal.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single_tenant_hsm_instance"), cmd.String("proposal"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.ExecuteSingleTenantHsmInstanceProposalRequest{Name: name}
									op, err := client.ExecuteSingleTenantHsmInstanceProposal(ctx, req)
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
								Usage: "describe proposals",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single_tenant_hsm_instance", Usage: "The single_tenant_hsm_instance.", Required: true},
									&cli.StringFlag{Name: "proposal", Usage: "The proposal.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single_tenant_hsm_instance"), cmd.String("proposal"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetSingleTenantHsmInstanceProposalRequest{Name: name}
									resp, err := client.GetSingleTenantHsmInstanceProposal(ctx, req)
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
								Usage: "list proposals",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &kmspb.ListSingleTenantHsmInstanceProposalsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSingleTenantHsmInstanceProposals(ctx, req)
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
								Name:  "delete",
								Usage: "delete proposals",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single_tenant_hsm_instance", Usage: "The single_tenant_hsm_instance.", Required: true},
									&cli.StringFlag{Name: "proposal", Usage: "The proposal.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single_tenant_hsm_instance"), cmd.String("proposal"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.DeleteSingleTenantHsmInstanceProposalRequest{Name: name}
									if err := client.DeleteSingleTenantHsmInstanceProposal(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "public-key",
						Usage: "Manage public-key resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe public-key",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
									&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
									&cli.StringFlag{Name: "crypto_key_version", Usage: "The crypto_key_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"), cmd.String("crypto_key_version"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetPublicKeyRequest{Name: name}
									resp, err := client.GetPublicKey(ctx, req)
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
						Name:  "retired-resources",
						Usage: "Manage retired-resources resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list retired-resources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &kmspb.ListRetiredResourcesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRetiredResources(ctx, req)
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
								Usage: "describe retired-resources",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "retired_resource", Usage: "The retired_resource.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/retiredResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("retired_resource"))
									client, err := kms.NewKeyManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetRetiredResourceRequest{Name: name}
									resp, err := client.GetRetiredResource(ctx, req)
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
						Name:  "single-tenant-hsm-instances",
						Usage: "Manage single-tenant-hsm-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list single-tenant-hsm-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe single-tenant-hsm-instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single_tenant_hsm_instance", Usage: "The single_tenant_hsm_instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s", cmd.String("project"), cmd.String("location"), cmd.String("single_tenant_hsm_instance"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.GetSingleTenantHsmInstanceRequest{Name: name}
									resp, err := client.GetSingleTenantHsmInstance(ctx, req)
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
								Usage: "create single-tenant-hsm-instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "single-tenant-hsm-instance-id", Usage: "The single tenant hsm instance id.", Required: false},
									&cli.BoolFlag{Name: "key-portability-enabled", Usage: "The key portability enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := kms.NewHsmManagementClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &kmspb.CreateSingleTenantHsmInstanceRequest{Parent: parent}
									req.SingleTenantHsmInstanceId = cmd.String("single-tenant-hsm-instance-id")
									req.SingleTenantHsmInstance = &kmspb.SingleTenantHsmInstance{
										KeyPortabilityEnabled: cmd.Bool("key-portability-enabled"),
									}
									op, err := client.CreateSingleTenantHsmInstance(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
