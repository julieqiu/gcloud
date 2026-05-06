package main

import (
	"context"
	"fmt"
	"os"

	publicca "cloud.google.com/go/security/publicca/apiv1"
	"cloud.google.com/go/security/publicca/apiv1/publiccapb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "publicca",
				Usage: "manage Public Certificate Authority API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "external-account-keys",
						Usage: "Manage external-account-keys resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create external-account-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := publicca.NewPublicCertificateAuthorityClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &publiccapb.CreateExternalAccountKeyRequest{Parent: parent}
									resp, err := client.CreateExternalAccountKey(ctx, req)
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
