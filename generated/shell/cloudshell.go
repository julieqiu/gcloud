package cloudshell

import (
	"context"
	"fmt"

	shell "cloud.google.com/go/shell/apiv1"
	"cloud.google.com/go/shell/apiv1/shellpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

// Command returns the gcloud cloudshell command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudshell",
		Usage: "manage Cloud Shell API resources",
		Commands: []*cli.Command{
			{
				Name:  "environments",
				Usage: "Manage environments resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/environments/%s", cmd.String("user"), cmd.String("environment"))
							client, err := shell.NewCloudShellClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &shellpb.GetEnvironmentRequest{Name: name}
							resp, err := client.GetEnvironment(ctx, req)
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
						Name:  "start",
						Usage: "start environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing start...")
							return nil
						},
					},
					{
						Name:  "authorize",
						Usage: "authorize environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing authorize...")
							return nil
						},
					},
					{
						Name:  "add-public-key",
						Usage: "add-public-key environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing add-public-key...")
							return nil
						},
					},
					{
						Name:  "remove-public-key",
						Usage: "remove-public-key environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing remove-public-key...")
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
				},
			},
		},
	}
}
