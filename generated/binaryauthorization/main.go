package main

import (
	"context"
	"fmt"
	"os"

	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
	"cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "binaryauthorization",
				Usage: "manage Binary Authorization API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "attestors",
						Usage: "Manage attestors resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create attestors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "attestor-id", Usage: "The attestor id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := binaryauthorization.NewBinauthzManagementServiceV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.CreateAttestorRequest{Parent: parent}
									req.AttestorId = cmd.String("attestor-id")
									req.Attestor = &binaryauthorizationpb.Attestor{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									resp, err := client.CreateAttestor(ctx, req)
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
								Usage: "describe attestors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "attestor", Usage: "The attestor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
									client, err := binaryauthorization.NewBinauthzManagementServiceV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.GetAttestorRequest{Name: name}
									resp, err := client.GetAttestor(ctx, req)
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
								Usage: "update attestors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "attestor", Usage: "The attestor.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
									client, err := binaryauthorization.NewBinauthzManagementServiceV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.UpdateAttestorRequest{}
									req.Attestor = &binaryauthorizationpb.Attestor{
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
									resp, err := client.UpdateAttestor(ctx, req)
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
								Usage: "list attestors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete attestors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "attestor", Usage: "The attestor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
									client, err := binaryauthorization.NewBinauthzManagementServiceV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.DeleteAttestorRequest{Name: name}
									if err := client.DeleteAttestor(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "validate-attestation-occurrence",
								Usage: "validate-attestation-occurrence attestors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing validate-attestation-occurrence...")
									return nil
								},
							},
						},
					},
					{
						Name:  "policy",
						Usage: "Manage policy resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/policy", cmd.String("project"))
									client, err := binaryauthorization.NewBinauthzManagementServiceV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.GetPolicyRequest{Name: name}
									resp, err := client.GetPolicy(ctx, req)
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
								Usage: "update policy",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/policy", cmd.String("project"))
									client, err := binaryauthorization.NewBinauthzManagementServiceV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.UpdatePolicyRequest{}
									req.Policy = &binaryauthorizationpb.Policy{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdatePolicy(ctx, req)
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
								Usage: "describe policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/policy", cmd.String("project"))
									client, err := binaryauthorization.NewSystemPolicyV1Client(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &binaryauthorizationpb.GetSystemPolicyRequest{Name: name}
									resp, err := client.GetSystemPolicy(ctx, req)
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
