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

package gateway

import (
	gateway "cloud.google.com/go/gateway/apiv1"
	"cloud.google.com/go/gateway/apiv1/gatewaypb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the connectgateway command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "connectgateway",
		Usage: "manage Connect Gateway API resources",
		Commands: []*cli.Command{
			{
				Name:  "memberships",
				Usage: "Manage memberships resources",
				Commands: []*cli.Command{

					{
						Name:  "generate-credentials",
						Usage: "generate-credentials memberships",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force-use-agent", Usage: "Whether to force the use of Connect Agent-based transport.", Required: false},
							&cli.StringFlag{Name: "kubernetes-namespace", Usage: "The namespace to use in the kubeconfig context.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The ID of the membership.", Required: true},
							&cli.StringFlag{Name: "operating-system", Usage: "The operating system where the kubeconfig will be used.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The Connect Gateway version to be used in the resulting.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gateway.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gatewaypb.GenerateCredentialsRequest{
								Name:                name,
								ForceUseAgent:       cmd.Bool("force-use-agent"),
								Version:             cmd.String("version"),
								KubernetesNamespace: cmd.String("kubernetes-namespace"),
								OperatingSystem:     gatewaypb.GenerateCredentialsRequest_OperatingSystem(gatewaypb.GenerateCredentialsRequest_OperatingSystem_value[cmd.String("operating-system")]),
							}

							resp, err := client.GenerateCredentials(ctx, req)
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
