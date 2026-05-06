package main

import (
	"context"
	"fmt"
	"os"

	gsuiteaddons "cloud.google.com/go/gsuiteaddons/apiv1"
	"cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "gsuiteaddons",
				Usage: "manage Google Workspace add-ons API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "authorization",
						Usage: "Manage authorization resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe authorization",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/authorization", cmd.String("project"))
									client, err := gsuiteaddons.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &gsuiteaddonspb.GetAuthorizationRequest{Name: name}
									resp, err := client.GetAuthorization(ctx, req)
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
						Name:  "deployments",
						Usage: "Manage deployments resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment-id", Usage: "The deployment id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := gsuiteaddons.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &gsuiteaddonspb.CreateDeploymentRequest{Parent: parent}
									req.DeploymentId = cmd.String("deployment-id")
									req.Deployment = &gsuiteaddonspb.Deployment{
										Name: cmd.String("name"),
										Etag: cmd.String("etag"),
									}
									resp, err := client.CreateDeployment(ctx, req)
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
								Name:  "replace-deployment",
								Usage: "replace-deployment deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
									fmt.Printf("Executing replace-deployment on %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
									client, err := gsuiteaddons.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &gsuiteaddonspb.GetDeploymentRequest{Name: name}
									resp, err := client.GetDeployment(ctx, req)
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
								Usage: "list deployments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
									client, err := gsuiteaddons.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &gsuiteaddonspb.DeleteDeploymentRequest{Name: name}
									if err := client.DeleteDeployment(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "install",
								Usage: "install deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
									fmt.Printf("Executing install on %s\n", name)
									return nil
								},
							},
							{
								Name:  "uninstall",
								Usage: "uninstall deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
									fmt.Printf("Executing uninstall on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "install-status",
						Usage: "Manage install-status resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe install-status",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/deployments/%s/installStatus", cmd.String("project"), cmd.String("deployment"))
									client, err := gsuiteaddons.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &gsuiteaddonspb.GetInstallStatusRequest{Name: name}
									resp, err := client.GetInstallStatus(ctx, req)
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
