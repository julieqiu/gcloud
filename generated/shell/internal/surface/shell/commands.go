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

package shell

import (
	shell "cloud.google.com/go/shell/apiv1"
	"cloud.google.com/go/shell/apiv1/shellpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the cloudshell command tree for inclusion under the gcloud root.
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
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/environments/%s", cmd.String("user"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := shell.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &shellpb.GetEnvironmentRequest{
								Name: name,
							}

							resp, err := client.GetEnvironment(ctx, req)
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
						Usage: "start environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-token", Usage: "The initial access token passed to the environment.", Required: false},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringSliceFlag{Name: "public-keys", Usage: "Public keys that should be added to the environment before it is started.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/environments/%s", cmd.String("user"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := shell.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &shellpb.StartEnvironmentRequest{
								Name:        name,
								AccessToken: cmd.String("access-token"),
								PublicKeys:  cmd.StringSlice("public-keys"),
							}

							op, err := client.StartEnvironment(ctx, req)
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
						Name:  "authorize",
						Usage: "authorize environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-token", Usage: "The OAuth access token that should be sent to the environment.", Required: false},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "id-token", Usage: "The OAuth ID token that should be sent to the environment.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("users/%s/environments/%s", cmd.String("user"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := shell.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &shellpb.AuthorizeEnvironmentRequest{
								Name:        name,
								AccessToken: cmd.String("access-token"),
								IdToken:     cmd.String("id-token"),
							}

							op, err := client.AuthorizeEnvironment(ctx, req)
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
						Name:  "add-public-key",
						Usage: "add-public-key environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "key", Usage: "Key that should be added to the environment.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("users/%s/environments/%s", cmd.String("user"), cmd.String("environment"))
							fmt.Printf("Executing add-public-key on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "remove-public-key",
						Usage: "remove-public-key environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "key", Usage: "Key that should be removed from the environment.", Required: false},
							&cli.StringFlag{Name: "user", Usage: "The ID of the user.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("users/%s/environments/%s", cmd.String("user"), cmd.String("environment"))
							fmt.Printf("Executing remove-public-key on %s\n", environment)
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
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("operations/%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
