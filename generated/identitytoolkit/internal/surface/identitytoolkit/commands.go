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

package identitytoolkit

import (
	identitytoolkit "cloud.google.com/go/identitytoolkit/apiv2"
	"cloud.google.com/go/identitytoolkit/apiv2/identitytoolkitpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the identitytoolkit command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "identitytoolkit",
		Usage: "manage Identity Toolkit API resources",
		Commands: []*cli.Command{
			{
				Name:  "mfa-enrollment",
				Usage: "Manage mfa-enrollment resources",
				Commands: []*cli.Command{

					{
						Name:  "finalize",
						Usage: "finalize mfa-enrollment",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "display-name", Usage: "Display name which is entered by users to distinguish between different.", Required: false},
							&cli.StringFlag{Name: "id-token", Usage: "ID token.", Required: true},
							&cli.StringFlag{Name: "tenant-id", Usage: "The ID of the Identity Platform tenant that the user enrolling MFA belongs.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := identitytoolkit.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &identitytoolkitpb.FinalizeMfaEnrollmentRequest{
								IdToken:     cmd.String("id-token"),
								DisplayName: cmd.String("display-name"),
								TenantId:    cmd.String("tenant-id"),
							}

							resp, err := client.FinalizeMfaEnrollment(ctx, req)
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
						Name:  "start",
						Usage: "start mfa-enrollment",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "id-token", Usage: "User's ID token.", Required: true},
							&cli.StringFlag{Name: "tenant-id", Usage: "The ID of the Identity Platform tenant that the user enrolling MFA belongs.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := identitytoolkit.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &identitytoolkitpb.StartMfaEnrollmentRequest{
								IdToken:  cmd.String("id-token"),
								TenantId: cmd.String("tenant-id"),
							}

							resp, err := client.StartMfaEnrollment(ctx, req)
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
						Name:  "withdraw",
						Usage: "withdraw mfa-enrollment",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "id-token", Usage: "User's ID token.", Required: true},
							&cli.StringFlag{Name: "mfa-enrollment-id", Usage: "MFA enrollment id from a current MFA enrollment.", Required: true},
							&cli.StringFlag{Name: "tenant-id", Usage: "The ID of the Identity Platform tenant that the user unenrolling MFA.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := identitytoolkit.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &identitytoolkitpb.WithdrawMfaRequest{
								IdToken:         cmd.String("id-token"),
								MfaEnrollmentId: cmd.String("mfa-enrollment-id"),
								TenantId:        cmd.String("tenant-id"),
							}

							resp, err := client.WithdrawMfa(ctx, req)
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
				Name:  "mfa-sign-in",
				Usage: "Manage mfa-sign-in resources",
				Commands: []*cli.Command{

					{
						Name:  "finalize",
						Usage: "finalize mfa-sign-in",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "mfa-pending-credential", Usage: "Pending credential from first factor sign-in.", Required: true},
							&cli.StringFlag{Name: "tenant-id", Usage: "The ID of the Identity Platform tenant the user is signing in to.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := identitytoolkit.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &identitytoolkitpb.FinalizeMfaSignInRequest{
								MfaPendingCredential: cmd.String("mfa-pending-credential"),
								TenantId:             cmd.String("tenant-id"),
							}

							resp, err := client.FinalizeMfaSignIn(ctx, req)
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
						Name:  "start",
						Usage: "start mfa-sign-in",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "mfa-enrollment-id", Usage: "MFA enrollment id from the user's list of current MFA enrollments.", Required: true},
							&cli.StringFlag{Name: "mfa-pending-credential", Usage: "Pending credential from first factor sign-in.", Required: true},
							&cli.StringFlag{Name: "tenant-id", Usage: "The ID of the Identity Platform tenant the user is signing in to.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := identitytoolkit.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &identitytoolkitpb.StartMfaSignInRequest{
								MfaPendingCredential: cmd.String("mfa-pending-credential"),
								MfaEnrollmentId:      cmd.String("mfa-enrollment-id"),
								TenantId:             cmd.String("tenant-id"),
							}

							resp, err := client.StartMfaSignIn(ctx, req)
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
