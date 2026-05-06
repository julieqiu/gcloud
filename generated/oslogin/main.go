package main

import (
	"context"
	"fmt"
	"os"

	oslogin "cloud.google.com/go/oslogin/apiv1"
	"cloud.google.com/go/oslogin/apiv1/osloginpb"
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
				Name:  "oslogin",
				Usage: "manage Cloud OS Login API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "login-profile",
						Usage: "Manage login-profile resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe login-profile",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("users/%s/projects/%s", cmd.String("user"), cmd.String("project"))
									client, err := oslogin.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osloginpb.GetLoginProfileRequest{Name: name}
									resp, err := client.GetLoginProfile(ctx, req)
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
								Name:  "delete",
								Usage: "delete projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("users/%s/projects/%s", cmd.String("user"), cmd.String("project"))
									client, err := oslogin.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osloginpb.DeletePosixAccountRequest{Name: name}
									if err := client.DeletePosixAccount(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "ssh-public-keys",
						Usage: "Manage ssh-public-keys resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create ssh-public-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: false},
									&cli.IntFlag{Name: "expiration-time-usec", Usage: "The expiration time usec.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("users/%s", cmd.String("user"))
									client, err := oslogin.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osloginpb.CreateSshPublicKeyRequest{Parent: parent}
									req.SshPublicKey = &osloginpb.SshPublicKey{
										Key:                cmd.String("key"),
										ExpirationTimeUsec: int64(cmd.Int("expiration-time-usec")),
									}
									resp, err := client.CreateSshPublicKey(ctx, req)
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
								Usage: "delete ssh-public-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
									&cli.StringFlag{Name: "fingerprint", Usage: "The fingerprint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("fingerprint"))
									client, err := oslogin.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osloginpb.DeleteSshPublicKeyRequest{Name: name}
									if err := client.DeleteSshPublicKey(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ssh-public-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
									&cli.StringFlag{Name: "fingerprint", Usage: "The fingerprint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("fingerprint"))
									client, err := oslogin.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osloginpb.GetSshPublicKeyRequest{Name: name}
									resp, err := client.GetSshPublicKey(ctx, req)
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
								Usage: "update ssh-public-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
									&cli.StringFlag{Name: "fingerprint", Usage: "The fingerprint.", Required: true},
									&cli.StringFlag{Name: "key", Usage: "The key.", Required: false},
									&cli.IntFlag{Name: "expiration-time-usec", Usage: "The expiration time usec.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("fingerprint"))
									client, err := oslogin.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &osloginpb.UpdateSshPublicKeyRequest{}
									req.SshPublicKey = &osloginpb.SshPublicKey{
										Name:               name,
										Key:                cmd.String("key"),
										ExpirationTimeUsec: int64(cmd.Int("expiration-time-usec")),
									}
									var paths []string
									if cmd.IsSet("key") {
										paths = append(paths, "key")
									}
									if cmd.IsSet("expiration-time-usec") {
										paths = append(paths, "expiration_time_usec")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateSshPublicKey(ctx, req)
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
						Name:  "users",
						Usage: "Manage users resources",
						Commands: []*cli.Command{
							{
								Name:  "import-ssh-public-key",
								Usage: "import-ssh-public-key users",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "user", Usage: "The user.", Required: true},
									&cli.StringFlag{Name: "fingerprint", Usage: "The fingerprint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("fingerprint"))
									fmt.Printf("Executing import-ssh-public-key on %s\n", name)
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
