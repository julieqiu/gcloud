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

package confidentialcomputing

import (
	confidentialcomputing "cloud.google.com/go/confidentialcomputing/apiv1"
	"cloud.google.com/go/confidentialcomputing/apiv1/confidentialcomputingpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the confidentialcomputing command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "confidentialcomputing",
		Usage: "manage Confidential Computing API resources",
		Commands: []*cli.Command{
			{
				Name:  "challenges",
				Usage: "Manage challenges resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create challenges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := confidentialcomputing.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &confidentialcomputingpb.CreateChallengeRequest{
								Parent: parent,
							}

							resp, err := client.CreateChallenge(ctx, req)
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
						Name:  "verify-attestation",
						Usage: "verify-attestation challenges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attester", Usage: "An optional indicator of the attester, only applies to certain.", Required: false},
							&cli.StringFlag{Name: "challenge", Usage: "The ID of the challenge.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							challenge := fmt.Sprintf("projects/%s/locations/%s/challenges/%s", cmd.String("project"), cmd.String("location"), cmd.String("challenge"))
							fmt.Printf("Executing verify-attestation on %s\n", challenge)
							return nil
						},
					},

					{
						Name:  "verify-confidential-space",
						Usage: "verify-confidential-space challenges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "challenge", Usage: "The ID of the challenge.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							challenge := fmt.Sprintf("projects/%s/locations/%s/challenges/%s", cmd.String("project"), cmd.String("location"), cmd.String("challenge"))
							fmt.Printf("Executing verify-confidential-space on %s\n", challenge)
							return nil
						},
					},

					{
						Name:  "verify-confidential-gke",
						Usage: "verify-confidential-gke challenges",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "challenge", Usage: "The ID of the challenge.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							challenge := fmt.Sprintf("projects/%s/locations/%s/challenges/%s", cmd.String("project"), cmd.String("location"), cmd.String("challenge"))
							fmt.Printf("Executing verify-confidential-gke on %s\n", challenge)
							return nil
						},
					},
				},
			},
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
