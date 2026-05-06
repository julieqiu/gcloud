package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
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
				Name:  "secretmanager",
				Usage: "manage Secret Manager API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
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
									client, err := secretmanager.NewClient(ctx)
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
									client, err := secretmanager.NewClient(ctx)
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
						Name:  "secrets",
						Usage: "Manage secrets resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list secrets",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &secretmanagerpb.ListSecretsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSecrets(ctx, req)
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
								Name:  "create",
								Usage: "create secrets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret-id", Usage: "The secret id.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &secretmanagerpb.CreateSecretRequest{Parent: parent}
									req.SecretId = cmd.String("secret-id")
									req.Secret = &secretmanagerpb.Secret{
										Etag: cmd.String("etag"),
									}
									resp, err := client.CreateSecret(ctx, req)
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
								Name:  "add-version",
								Usage: "add-version secrets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
									fmt.Printf("Executing add-version on %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe secrets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &secretmanagerpb.GetSecretRequest{Name: name}
									resp, err := client.GetSecret(ctx, req)
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
								Usage: "update secrets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &secretmanagerpb.UpdateSecretRequest{}
									req.Secret = &secretmanagerpb.Secret{
										Name: name,
										Etag: cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSecret(ctx, req)
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
								Usage: "delete secrets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s", cmd.String("project"), cmd.String("secret"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &secretmanagerpb.DeleteSecretRequest{Name: name}
									if err := client.DeleteSecret(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy secrets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy secrets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions secrets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "versions",
						Usage: "Manage versions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list versions",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &secretmanagerpb.ListSecretVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSecretVersions(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
									&cli.StringFlag{Name: "secret_version", Usage: "The secret_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("secret_version"))
									client, err := secretmanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &secretmanagerpb.GetSecretVersionRequest{Name: name}
									resp, err := client.GetSecretVersion(ctx, req)
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
								Name:  "access",
								Usage: "access versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
									&cli.StringFlag{Name: "secret_version", Usage: "The secret_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("secret_version"))
									fmt.Printf("Executing access on %s\n", name)
									return nil
								},
							},
							{
								Name:  "disable",
								Usage: "disable versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
									&cli.StringFlag{Name: "secret_version", Usage: "The secret_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("secret_version"))
									fmt.Printf("Executing disable on %s\n", name)
									return nil
								},
							},
							{
								Name:  "enable",
								Usage: "enable versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
									&cli.StringFlag{Name: "secret_version", Usage: "The secret_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("secret_version"))
									fmt.Printf("Executing enable on %s\n", name)
									return nil
								},
							},
							{
								Name:  "destroy",
								Usage: "destroy versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "secret", Usage: "The secret.", Required: true},
									&cli.StringFlag{Name: "secret_version", Usage: "The secret_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", cmd.String("project"), cmd.String("secret"), cmd.String("secret_version"))
									fmt.Printf("Executing destroy on %s\n", name)
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
