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

package binaryauthorization

import (
	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
	"cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the binaryauthorization command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "binaryauthorization",
		Usage: "manage Binary Authorization API resources",
		Commands: []*cli.Command{
			{
				Name:  "attestors",
				Usage: "Manage attestors resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create attestors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attestor-id", Usage: "The [attestors][google.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := binaryauthorization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &binaryauthorizationpb.CreateAttestorRequest{
								Parent:     parent,
								AttestorId: cmd.String("attestor-id"),
							}

							resp, err := client.CreateAttestor(ctx, req)
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
						Name:  "describe",
						Usage: "describe attestors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attestor", Usage: "The ID of the attestor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := binaryauthorization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &binaryauthorizationpb.GetAttestorRequest{
								Name: name,
							}

							resp, err := client.GetAttestor(ctx, req)
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
						Usage: "update attestors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attestor", Usage: "The ID of the attestor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "attestor.name" not yet supported.
							attestor_name := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
							fmt.Printf("Executing update on %s\n", attestor_name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list attestors",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := binaryauthorization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &binaryauthorizationpb.ListAttestorsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAttestors(ctx, req)
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
						Usage: "delete attestors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attestor", Usage: "The ID of the attestor.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAttestor on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := binaryauthorization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &binaryauthorizationpb.DeleteAttestorRequest{
								Name: name,
							}

							if err := client.DeleteAttestor(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "validate-attestation-occurrence",
						Usage: "validate-attestation-occurrence attestors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "attestor", Usage: "The ID of the attestor.", Required: true},
							&cli.StringFlag{Name: "occurrence-note", Usage: "The resource name of the [Note][grafeas.", Required: true},
							&cli.StringFlag{Name: "occurrence-resource-uri", Usage: "The URI of the artifact (e.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							attestor := fmt.Sprintf("projects/%s/attestors/%s", cmd.String("project"), cmd.String("attestor"))
							fmt.Printf("Executing validate-attestation-occurrence on %s\n", attestor)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/policy", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := binaryauthorization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &binaryauthorizationpb.GetPolicyRequest{
								Name: name,
							}

							resp, err := client.GetPolicy(ctx, req)
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
						Usage: "update policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "policy.name" not yet supported.
							policy_name := fmt.Sprintf("projects/%s/policy", cmd.String("project"))
							fmt.Printf("Executing update on %s\n", policy_name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("locations/%s/policy", cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := binaryauthorization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &binaryauthorizationpb.GetSystemPolicyRequest{
								Name: name,
							}

							resp, err := client.GetSystemPolicy(ctx, req)
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
