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

package gsuiteaddons

import (
	gsuiteaddons "cloud.google.com/go/gsuiteaddons/apiv1"
	"cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the gsuiteaddons command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gsuiteaddons",
		Usage: "manage Google Workspace add-ons API resources",
		Commands: []*cli.Command{
			{
				Name:  "authorization",
				Usage: "Manage authorization resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe authorization",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/authorization", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.GetAuthorizationRequest{
								Name: name,
							}

							resp, err := client.GetAuthorization(ctx, req)
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
				Name:  "deployments",
				Usage: "Manage deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment-id", Usage: "The id to use for this deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.CreateDeploymentRequest{
								Parent:       parent,
								DeploymentId: cmd.String("deployment-id"),
							}

							resp, err := client.CreateDeployment(ctx, req)
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
						Name:  "replace-deployment",
						Usage: "replace-deployment deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "deployment.name" not yet supported.
							deployment_name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
							fmt.Printf("Executing replace-deployment on %s\n", deployment_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.GetDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetDeployment(ctx, req)
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
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of deployments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListDeployments` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.ListDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDeployments(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag of the deployment to delete.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteDeployment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.DeleteDeploymentRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteDeployment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "install",
						Usage: "install deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute InstallDeployment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.InstallDeploymentRequest{
								Name: name,
							}

							if err := client.InstallDeployment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "uninstall",
						Usage: "uninstall deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deployments/%s", cmd.String("project"), cmd.String("deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute UninstallDeployment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.UninstallDeploymentRequest{
								Name: name,
							}

							if err := client.UninstallDeployment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
							&cli.StringFlag{Name: "deployment", Usage: "The ID of the deployment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deployments/%s/installStatus", cmd.String("project"), cmd.String("deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gsuiteaddons.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gsuiteaddonspb.GetInstallStatusRequest{
								Name: name,
							}

							resp, err := client.GetInstallStatus(ctx, req)
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
