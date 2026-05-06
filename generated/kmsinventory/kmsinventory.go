package kmsinventory

import (
	"context"
	"fmt"

	inventory "cloud.google.com/go/kms/inventory/apiv1"
	"cloud.google.com/go/kms/inventory/apiv1/inventorypb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

// Command returns the gcloud kmsinventory command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "kmsinventory",
		Usage: "manage KMS Inventory API resources",
		Commands: []*cli.Command{
			{
				Name:  "crypto-keys",
				Usage: "Manage crypto-keys resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list crypto-keys",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "protected-resources",
				Usage: "Manage protected-resources resources",
				Commands: []*cli.Command{
					{
						Name:  "search",
						Usage: "search protected-resources",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search...")
							return nil
						},
					},
				},
			},
			{
				Name:  "protected-resources-summary",
				Usage: "Manage protected-resources-summary resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe protected-resources-summary",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "key_ring", Usage: "The key_ring.", Required: true},
							&cli.StringFlag{Name: "crypto_key", Usage: "The crypto_key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/protectedResourcesSummary", cmd.String("project"), cmd.String("location"), cmd.String("key_ring"), cmd.String("crypto_key"))
							client, err := inventory.NewKeyTrackingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &inventorypb.GetProtectedResourcesSummaryRequest{Name: name}
							resp, err := client.GetProtectedResourcesSummary(ctx, req)
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
	}
}
