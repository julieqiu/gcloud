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

package oslogin

import (
	oslogin "cloud.google.com/go/oslogin/apiv1"
	"cloud.google.com/go/oslogin/apiv1/osloginpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the oslogin command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "oslogin",
		Usage: "manage Cloud OS Login API resources",
		Commands: []*cli.Command{
			{
				Name:  "login-profile",
				Usage: "Manage login-profile resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe login-profile",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project-id", Usage: "The project ID of the Google Cloud Platform project.", Required: false},
							&cli.StringFlag{Name: "system-id", Usage: "A system ID for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s", cmd.String("user"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.GetLoginProfileRequest{
								Name:      name,
								ProjectId: cmd.String("project-id"),
								SystemId:  cmd.String("system-id"),
							}

							resp, err := client.GetLoginProfile(ctx, req)
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
				Name:  "projects",
				Usage: "Manage projects resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete projects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/projects/%s", cmd.String("user"), cmd.String("project"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeletePosixAccount on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.DeletePosixAccountRequest{
								Name: name,
							}

							if err := client.DeletePosixAccount(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("users/%s", cmd.String("user"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.CreateSshPublicKeyRequest{
								Parent: parent,
							}

							resp, err := client.CreateSshPublicKey(ctx, req)
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
						Usage: "delete ssh-public-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ssh-public-key", Usage: "The ID of the ssh public key.", Required: true},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("ssh-public-key"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSshPublicKey on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.DeleteSshPublicKeyRequest{
								Name: name,
							}

							if err := client.DeleteSshPublicKey(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ssh-public-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ssh-public-key", Usage: "The ID of the ssh public key.", Required: true},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("ssh-public-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.GetSshPublicKeyRequest{
								Name: name,
							}

							resp, err := client.GetSshPublicKey(ctx, req)
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
						Usage: "update ssh-public-keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "ssh-public-key", Usage: "The ID of the ssh public key.", Required: true},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/sshPublicKeys/%s", cmd.String("user"), cmd.String("ssh-public-key"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.UpdateSshPublicKeyRequest{
								Name: name,
							}

							resp, err := client.UpdateSshPublicKey(ctx, req)
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
				Name:  "users",
				Usage: "Manage users resources",
				Commands: []*cli.Command{

					{
						Name:  "import-ssh-public-key",
						Usage: "import-ssh-public-key users",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project-id", Usage: "The project ID of the Google Cloud Platform project.", Required: false},
							&cli.StringSliceFlag{Name: "regions", Usage: "The regions to which to assert that the key was written.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("users/%s", cmd.String("user"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oslogin.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &osloginpb.ImportSshPublicKeyRequest{
								Parent:    parent,
								ProjectId: cmd.String("project-id"),
								Regions:   cmd.StringSlice("regions"),
							}

							resp, err := client.ImportSshPublicKey(ctx, req)
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
