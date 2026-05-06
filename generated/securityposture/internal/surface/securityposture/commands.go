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

package securityposture

import (
	securityposture "cloud.google.com/go/securityposture/apiv1"
	"cloud.google.com/go/securityposture/apiv1/securityposturepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the securityposture command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "securityposture",
		Usage: "manage Security Posture API resources",
		Commands: []*cli.Command{
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
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
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/operations/%s", cmd.String("organization"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/operations/%s", cmd.String("organization"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/operations/%s", cmd.String("organization"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "posture-deployments",
				Usage: "Manage posture-deployments resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list posture-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied on the resource, defined by EBNF grammar.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.ListPostureDeploymentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPostureDeployments(ctx, req)
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
						Name:  "describe",
						Usage: "describe posture-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-deployment", Usage: "The ID of the posture deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/postureDeployments/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture-deployment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.GetPostureDeploymentRequest{
								Name: name,
							}

							resp, err := client.GetPostureDeployment(ctx, req)
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
						Name:  "create",
						Usage: "create posture-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-deployment-id", Usage: "User provided identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.CreatePostureDeploymentRequest{
								Parent:              parent,
								PostureDeploymentId: cmd.String("posture-deployment-id"),
							}

							op, err := client.CreatePostureDeployment(ctx, req)
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
						Name:  "update",
						Usage: "update posture-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-deployment", Usage: "The ID of the posture deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "posture_deployment.name" not yet supported.
							posture_deployment_name := fmt.Sprintf("organizations/%s/locations/%s/postureDeployments/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture-deployment"))
							fmt.Printf("Executing update on %s\n", posture_deployment_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete posture-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Etag value of the PostureDeployment to be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-deployment", Usage: "The ID of the posture deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/postureDeployments/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture-deployment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePostureDeployment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.DeletePostureDeploymentRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeletePostureDeployment(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "posture-templates",
				Usage: "Manage posture-templates resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list posture-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter to be applied on the resource, defined by EBNF grammar.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.ListPostureTemplatesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPostureTemplates(ctx, req)
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
						Name:  "describe",
						Usage: "describe posture-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-template", Usage: "The ID of the posture template.", Required: true},
							&cli.StringFlag{Name: "revision-id", Usage: "Specific revision_id of a Posture Template.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/postureTemplates/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture-template"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.GetPostureTemplateRequest{
								Name:       name,
								RevisionId: cmd.String("revision-id"),
							}

							resp, err := client.GetPostureTemplate(ctx, req)
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
				Name:  "postures",
				Usage: "Manage postures resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.ListPosturesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPostures(ctx, req)
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
						Name:  "list",
						Usage: "list postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "posture", Usage: "The ID of the posture.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/postures/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.ListPostureRevisionsRequest{
								Name:      name,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListPostureRevisions(ctx, req)
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
						Name:  "describe",
						Usage: "describe postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture", Usage: "The ID of the posture.", Required: true},
							&cli.StringFlag{Name: "revision-id", Usage: "Posture revision which needs to be retrieved.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/postures/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.GetPostureRequest{
								Name:       name,
								RevisionId: cmd.String("revision-id"),
							}

							resp, err := client.GetPosture(ctx, req)
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
						Name:  "create",
						Usage: "create postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-id", Usage: "User provided identifier.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.CreatePostureRequest{
								Parent:    parent,
								PostureId: cmd.String("posture-id"),
							}

							op, err := client.CreatePosture(ctx, req)
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
						Name:  "update",
						Usage: "update postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture", Usage: "The ID of the posture.", Required: true},
							&cli.StringFlag{Name: "revision-id", Usage: "Posture revision which needs to be updated.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "posture.name" not yet supported.
							posture_name := fmt.Sprintf("organizations/%s/locations/%s/postures/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture"))
							fmt.Printf("Executing update on %s\n", posture_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "Etag value of the Posture to be deleted.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture", Usage: "The ID of the posture.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/postures/%s", cmd.String("organization"), cmd.String("location"), cmd.String("posture"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeletePosture %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.DeletePostureRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							op, err := client.DeletePosture(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "extract",
						Usage: "extract postures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "posture-id", Usage: "User provided identifier.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "Workload from which the policies are to be extracted, it should.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := securityposture.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &securityposturepb.ExtractPostureRequest{
								Parent:    parent,
								PostureId: cmd.String("posture-id"),
								Workload:  cmd.String("workload"),
							}

							op, err := client.ExtractPosture(ctx, req)
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
				},
			},
		},
	}
}
