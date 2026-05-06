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

package kms

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudkms command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudkms",
		Usage: "manage Cloud Key Management Service (KMS) API resources",
		Commands: []*cli.Command{
			{
				Name:  "autokey-config",
				Usage: "Manage autokey-config resources",
				Commands: []*cli.Command{

					{
						Name:  "update",
						Usage: "update autokey-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "autokey_config.name" not yet supported.
							autokey_config_name := fmt.Sprintf("folders/%s/autokeyConfig", cmd.String("folder"))
							fmt.Printf("Executing update on %s\n", autokey_config_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe autokey-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "folder", Usage: "The ID of the folder.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("folders/%s/autokeyConfig", cmd.String("folder"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetAutokeyConfigRequest{
								Name: name,
							}

							resp, err := client.GetAutokeyConfig(ctx, req)
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
				Name:  "crypto-key-versions",
				Usage: "Manage crypto-key-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The fields to include in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListCryptoKeyVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								View:      kmspb.CryptoKeyVersion_CryptoKeyVersionView(kmspb.CryptoKeyVersion_CryptoKeyVersionView_value[cmd.String("view")]),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCryptoKeyVersions(ctx, req)
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
						Usage: "describe crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetCryptoKeyVersionRequest{
								Name: name,
							}

							resp, err := client.GetCryptoKeyVersion(ctx, req)
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
						Usage: "create crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateCryptoKeyVersionRequest{
								Parent: parent,
							}

							resp, err := client.CreateCryptoKeyVersion(ctx, req)
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
						Usage: "delete crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCryptoKeyVersion %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.DeleteCryptoKeyVersionRequest{
								Name: name,
							}

							op, err := client.DeleteCryptoKeyVersion(ctx, req)
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
						Name:  "import",
						Usage: "import crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "algorithm", Usage: "The.", Required: true},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The optional [name][google.", Required: false},
							&cli.StringFlag{Name: "import-job", Usage: "The [name][google.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wrapped-key", Usage: "The wrapped key material to import.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ImportCryptoKeyVersionRequest{
								Parent:           parent,
								CryptoKeyVersion: cmd.String("crypto-key-version"),
								Algorithm:        kmspb.CryptoKeyVersion_CryptoKeyVersionAlgorithm(kmspb.CryptoKeyVersion_CryptoKeyVersionAlgorithm_value[cmd.String("algorithm")]),
								ImportJob:        cmd.String("import-job"),
								WrappedKey:       []byte(cmd.String("wrapped-key")),
							}

							resp, err := client.ImportCryptoKeyVersion(ctx, req)
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
						Usage: "update crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "crypto_key_version.name" not yet supported.
							crypto_key_version_name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							fmt.Printf("Executing update on %s\n", crypto_key_version_name)
							return nil
						},
					},

					{
						Name:  "destroy",
						Usage: "destroy crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.DestroyCryptoKeyVersionRequest{
								Name: name,
							}

							resp, err := client.DestroyCryptoKeyVersion(ctx, req)
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
						Name:  "restore",
						Usage: "restore crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.RestoreCryptoKeyVersionRequest{
								Name: name,
							}

							resp, err := client.RestoreCryptoKeyVersion(ctx, req)
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
						Name:  "raw-encrypt",
						Usage: "raw-encrypt crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "additional-authenticated-data", Usage: "Optional data that, if specified, must also be provided during.", Required: false},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "initialization-vector", Usage: "A customer-supplied initialization vector that will be used for.", Required: false},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "plaintext", Usage: "The data to encrypt.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.RawEncryptRequest{
								Name:                        name,
								Plaintext:                   []byte(cmd.String("plaintext")),
								AdditionalAuthenticatedData: []byte(cmd.String("additional-authenticated-data")),
								InitializationVector:        []byte(cmd.String("initialization-vector")),
							}

							resp, err := client.RawEncrypt(ctx, req)
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
						Name:  "raw-decrypt",
						Usage: "raw-decrypt crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "additional-authenticated-data", Usage: "Optional data that must match the data originally supplied in.", Required: false},
							&cli.StringFlag{Name: "ciphertext", Usage: "The encrypted data originally returned in.", Required: true},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "initialization-vector", Usage: "The initialization vector (IV) used during encryption, which must.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "tag-length", Usage: "The length of the authentication tag that is appended to the end of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.RawDecryptRequest{
								Name:                        name,
								Ciphertext:                  []byte(cmd.String("ciphertext")),
								AdditionalAuthenticatedData: []byte(cmd.String("additional-authenticated-data")),
								InitializationVector:        []byte(cmd.String("initialization-vector")),
								TagLength:                   int32(cmd.Int("tag-length")),
							}

							resp, err := client.RawDecrypt(ctx, req)
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
						Name:  "asymmetric-sign",
						Usage: "asymmetric-sign crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "data", Usage: "The data to sign.", Required: false},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.AsymmetricSignRequest{
								Name: name,
								Data: []byte(cmd.String("data")),
							}

							resp, err := client.AsymmetricSign(ctx, req)
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
						Name:  "asymmetric-decrypt",
						Usage: "asymmetric-decrypt crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ciphertext", Usage: "The data encrypted with the named.", Required: true},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.AsymmetricDecryptRequest{
								Name:       name,
								Ciphertext: []byte(cmd.String("ciphertext")),
							}

							resp, err := client.AsymmetricDecrypt(ctx, req)
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
						Name:  "mac-sign",
						Usage: "mac-sign crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "data", Usage: "The data to sign.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.MacSignRequest{
								Name: name,
								Data: []byte(cmd.String("data")),
							}

							resp, err := client.MacSign(ctx, req)
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
						Name:  "mac-verify",
						Usage: "mac-verify crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "data", Usage: "The data used previously as a.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mac", Usage: "The signature to verify.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.MacVerifyRequest{
								Name: name,
								Data: []byte(cmd.String("data")),
								Mac:  []byte(cmd.String("mac")),
							}

							resp, err := client.MacVerify(ctx, req)
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
						Name:  "decapsulate",
						Usage: "decapsulate crypto-key-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ciphertext", Usage: "The ciphertext produced from encapsulation with the.", Required: true},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.DecapsulateRequest{
								Name:       name,
								Ciphertext: []byte(cmd.String("ciphertext")),
							}

							resp, err := client.Decapsulate(ctx, req)
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
				Name:  "crypto-keys",
				Usage: "Manage crypto-keys resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version-view", Usage: "The fields of the primary version to include in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListCryptoKeysRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								VersionView: kmspb.CryptoKeyVersion_CryptoKeyVersionView(kmspb.CryptoKeyVersion_CryptoKeyVersionView_value[cmd.String("version-view")]),
								Filter:      cmd.String("filter"),
								OrderBy:     cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCryptoKeys(ctx, req)
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
						Usage: "describe crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetCryptoKeyRequest{
								Name: name,
							}

							resp, err := client.GetCryptoKey(ctx, req)
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
						Usage: "create crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key-id", Usage: "It must be unique within a KeyRing and match the regular.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-initial-version-creation", Usage: "If set to true, the request will create a.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateCryptoKeyRequest{
								Parent:                     parent,
								CryptoKeyId:                cmd.String("crypto-key-id"),
								SkipInitialVersionCreation: cmd.Bool("skip-initial-version-creation"),
							}

							resp, err := client.CreateCryptoKey(ctx, req)
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
						Usage: "delete crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCryptoKey %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.DeleteCryptoKeyRequest{
								Name: name,
							}

							op, err := client.DeleteCryptoKey(ctx, req)
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
						Usage: "update crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "crypto_key.name" not yet supported.
							crypto_key_name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							fmt.Printf("Executing update on %s\n", crypto_key_name)
							return nil
						},
					},

					{
						Name:  "update-primary-version",
						Usage: "update-primary-version crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version-id", Usage: "The id of the child.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.UpdateCryptoKeyPrimaryVersionRequest{
								Name:               name,
								CryptoKeyVersionId: cmd.String("crypto-key-version-id"),
							}

							resp, err := client.UpdateCryptoKeyPrimaryVersion(ctx, req)
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
						Name:  "encrypt",
						Usage: "encrypt crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "additional-authenticated-data", Usage: "Optional data that, if specified, must also be provided during.", Required: false},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "plaintext", Usage: "The data to encrypt.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.EncryptRequest{
								Name:                        name,
								Plaintext:                   []byte(cmd.String("plaintext")),
								AdditionalAuthenticatedData: []byte(cmd.String("additional-authenticated-data")),
							}

							resp, err := client.Encrypt(ctx, req)
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
						Name:  "decrypt",
						Usage: "decrypt crypto-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "additional-authenticated-data", Usage: "Optional data that must match the data originally supplied in.", Required: false},
							&cli.StringFlag{Name: "ciphertext", Usage: "The encrypted data originally returned in.", Required: true},
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.DecryptRequest{
								Name:                        name,
								Ciphertext:                  []byte(cmd.String("ciphertext")),
								AdditionalAuthenticatedData: []byte(cmd.String("additional-authenticated-data")),
							}

							resp, err := client.Decrypt(ctx, req)
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
				Name:  "ekm-config",
				Usage: "Manage ekm-config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe ekm-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ekmConfig", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetEkmConfigRequest{
								Name: name,
							}

							resp, err := client.GetEkmConfig(ctx, req)
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
						Usage: "update ekm-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "ekm_config.name" not yet supported.
							ekm_config_name := fmt.Sprintf("projects/%s/locations/%s/ekmConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", ekm_config_name)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListEkmConnectionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListEkmConnections(ctx, req)
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
						Usage: "describe ekm-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ekm-connection", Usage: "The ID of the ekm connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("ekm-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetEkmConnectionRequest{
								Name: name,
							}

							resp, err := client.GetEkmConnection(ctx, req)
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
						Usage: "create ekm-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ekm-connection-id", Usage: "It must be unique within a location and match the regular.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateEkmConnectionRequest{
								Parent:          parent,
								EkmConnectionId: cmd.String("ekm-connection-id"),
							}

							resp, err := client.CreateEkmConnection(ctx, req)
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
						Usage: "update ekm-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ekm-connection", Usage: "The ID of the ekm connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "ekm_connection.name" not yet supported.
							ekm_connection_name := fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("ekm-connection"))
							fmt.Printf("Executing update on %s\n", ekm_connection_name)
							return nil
						},
					},

					{
						Name:  "verify-connectivity",
						Usage: "verify-connectivity ekm-connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ekm-connection", Usage: "The ID of the ekm connection.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/ekmConnections/%s", cmd.String("project"), cmd.String("location"), cmd.String("ekm-connection"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.VerifyConnectivityRequest{
								Name: name,
							}

							resp, err := client.VerifyConnectivity(ctx, req)
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
				Name:  "import-jobs",
				Usage: "Manage import-jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list import-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListImportJobsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListImportJobs(ctx, req)
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
						Usage: "describe import-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "import-job", Usage: "The ID of the import job.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/importJobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("import-job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetImportJobRequest{
								Name: name,
							}

							resp, err := client.GetImportJob(ctx, req)
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
						Usage: "create import-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "import-job-id", Usage: "It must be unique within a KeyRing and match the regular.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateImportJobRequest{
								Parent:      parent,
								ImportJobId: cmd.String("import-job-id"),
							}

							resp, err := client.CreateImportJob(ctx, req)
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
				Name:  "key-handles",
				Usage: "Manage key-handles resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create key-handles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-handle-id", Usage: "Id of the [KeyHandle][google.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateKeyHandleRequest{
								Parent:      parent,
								KeyHandleId: cmd.String("key-handle-id"),
							}

							op, err := client.CreateKeyHandle(ctx, req)
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
						Usage: "describe key-handles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-handle", Usage: "The ID of the key handle.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyHandles/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-handle"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetKeyHandleRequest{
								Name: name,
							}

							resp, err := client.GetKeyHandle(ctx, req)
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
						Usage: "list key-handles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter to apply when listing.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListKeyHandlesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListKeyHandles(ctx, req)
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
				Name:  "key-rings",
				Usage: "Manage key-rings resources",
				Commands: []*cli.Command{

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing test-iam-permissions on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListKeyRingsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListKeyRings(ctx, req)
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
						Usage: "describe key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetKeyRingRequest{
								Name: name,
							}

							resp, err := client.GetKeyRing(ctx, req)
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
						Usage: "create key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring-id", Usage: "It must be unique within a location and match the regular.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateKeyRingRequest{
								Parent:    parent,
								KeyRingId: cmd.String("key-ring-id"),
							}

							resp, err := client.CreateKeyRing(ctx, req)
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
						Usage: "set-iam-policy key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing set-iam-policy on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
							fmt.Printf("Executing describe on %s\n", resource)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions key-rings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"))
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
						Name:  "generate-random-bytes",
						Usage: "generate-random-bytes locations",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "length-bytes", Usage: "The length in bytes of the amount of randomness to retrieve.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "protection-level", Usage: "The [ProtectionLevel][google.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							location := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing generate-random-bytes on %s\n", location)
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

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
				Name:  "projects",
				Usage: "Manage projects resources",
				Commands: []*cli.Command{

					{
						Name:  "show-effective-autokey-config",
						Usage: "show-effective-autokey-config projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ShowEffectiveAutokeyConfigRequest{
								Parent: parent,
							}

							resp, err := client.ShowEffectiveAutokeyConfig(ctx, req)
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
				Name:  "proposals",
				Usage: "Manage proposals resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create proposals",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance-proposal-id", Usage: "It must be unique within a location and match the regular.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateSingleTenantHsmInstanceProposalRequest{
								Parent:                            parent,
								SingleTenantHsmInstanceProposalId: cmd.String("single-tenant-hsm-instance-proposal-id"),
							}

							op, err := client.CreateSingleTenantHsmInstanceProposal(ctx, req)
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
						Name:  "approve",
						Usage: "approve proposals",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "proposal", Usage: "The ID of the proposal.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"), cmd.String("proposal"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ApproveSingleTenantHsmInstanceProposalRequest{
								Name: name,
							}

							resp, err := client.ApproveSingleTenantHsmInstanceProposal(ctx, req)
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
						Name:  "execute",
						Usage: "execute proposals",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "proposal", Usage: "The ID of the proposal.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"), cmd.String("proposal"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ExecuteSingleTenantHsmInstanceProposalRequest{
								Name: name,
							}

							op, err := client.ExecuteSingleTenantHsmInstanceProposal(ctx, req)
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
						Usage: "describe proposals",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "proposal", Usage: "The ID of the proposal.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"), cmd.String("proposal"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetSingleTenantHsmInstanceProposalRequest{
								Name: name,
							}

							resp, err := client.GetSingleTenantHsmInstanceProposal(ctx, req)
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
						Usage: "list proposals",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If set to true,.", Required: false},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListSingleTenantHsmInstanceProposalsRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								Filter:      cmd.String("filter"),
								OrderBy:     cmd.String("order-by"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListSingleTenantHsmInstanceProposals(ctx, req)
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
						Usage: "delete proposals",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "proposal", Usage: "The ID of the proposal.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s/proposals/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"), cmd.String("proposal"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSingleTenantHsmInstanceProposal on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.DeleteSingleTenantHsmInstanceProposalRequest{
								Name: name,
							}

							if err := client.DeleteSingleTenantHsmInstanceProposal(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "crypto-key", Usage: "The ID of the crypto key.", Required: true},
							&cli.StringFlag{Name: "crypto-key-version", Usage: "The ID of the crypto key version.", Required: true},
							&cli.StringFlag{Name: "key-ring", Usage: "The ID of the key ring.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public-key-format", Usage: "The [PublicKey][google.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("key-ring"), cmd.String("crypto-key"), cmd.String("crypto-key-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetPublicKeyRequest{
								Name:            name,
								PublicKeyFormat: kmspb.PublicKey_PublicKeyFormat(kmspb.PublicKey_PublicKeyFormat_value[cmd.String("public-key-format")]),
							}

							resp, err := client.GetPublicKey(ctx, req)
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
				Name:  "retired-resources",
				Usage: "Manage retired-resources resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list retired-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListRetiredResourcesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRetiredResources(ctx, req)
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
						Usage: "describe retired-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "retired-resource", Usage: "The ID of the retired resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/retiredResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("retired-resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetRetiredResourceRequest{
								Name: name,
							}

							resp, err := client.GetRetiredResource(ctx, req)
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
				Name:  "single-tenant-hsm-instances",
				Usage: "Manage single-tenant-hsm-instances resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list single-tenant-hsm-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Only include resources that match the filter in the response.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Specify how the results should be sorted.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Optional limit on the number of.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Optional pagination token, returned earlier via.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If set to true,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.ListSingleTenantHsmInstancesRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								Filter:      cmd.String("filter"),
								OrderBy:     cmd.String("order-by"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListSingleTenantHsmInstances(ctx, req)
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
						Usage: "describe single-tenant-hsm-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance", Usage: "The ID of the single tenant hsm instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/singleTenantHsmInstances/%s", cmd.String("project"), cmd.String("location"), cmd.String("single-tenant-hsm-instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.GetSingleTenantHsmInstanceRequest{
								Name: name,
							}

							resp, err := client.GetSingleTenantHsmInstance(ctx, req)
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
						Usage: "create single-tenant-hsm-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "single-tenant-hsm-instance-id", Usage: "It must be unique within a location and match the regular.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := kms.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &kmspb.CreateSingleTenantHsmInstanceRequest{
								Parent:                    parent,
								SingleTenantHsmInstanceId: cmd.String("single-tenant-hsm-instance-id"),
							}

							op, err := client.CreateSingleTenantHsmInstance(ctx, req)
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
		},
	}
}
