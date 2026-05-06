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

package assuredworkloads

import (
	assuredworkloads "cloud.google.com/go/assuredworkloads/apiv1"
	"cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the assuredworkloads command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "assuredworkloads",
		Usage: "manage Assured Workloads API resources",
		Commands: []*cli.Command{
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
				},
			},
			{
				Name:  "workloads",
				Usage: "Manage workloads resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external-id", Usage: "A identifier associated with the workload and underlying projects which.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := assuredworkloads.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assuredworkloadspb.CreateWorkloadRequest{
								Parent:     parent,
								ExternalId: cmd.String("external-id"),
							}

							op, err := client.CreateWorkload(ctx, req)
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
						Usage: "update workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "workload.name" not yet supported.
							workload_name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							fmt.Printf("Executing update on %s\n", workload_name)
							return nil
						},
					},

					{
						Name:  "restrict-allowed-resources",
						Usage: "restrict-allowed-resources workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "restriction-type", Usage: "The type of restriction for using gcp products in the Workload environment.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := assuredworkloads.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assuredworkloadspb.RestrictAllowedResourcesRequest{
								Name:            name,
								RestrictionType: assuredworkloadspb.RestrictAllowedResourcesRequest_RestrictionType(assuredworkloadspb.RestrictAllowedResourcesRequest_RestrictionType_value[cmd.String("restriction-type")]),
							}

							resp, err := client.RestrictAllowedResources(ctx, req)
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
						Usage: "delete workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "The etag of the workload.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteWorkload on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := assuredworkloads.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assuredworkloadspb.DeleteWorkloadRequest{
								Name: name,
								Etag: cmd.String("etag"),
							}

							if err := client.DeleteWorkload(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The ID of the workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := assuredworkloads.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assuredworkloadspb.GetWorkloadRequest{
								Name: name,
							}

							resp, err := client.GetWorkload(ctx, req)
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
						Usage: "list workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A custom filter for filtering by properties of a workload.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "organization", Usage: "The ID of the organization.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token returned from previous request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := assuredworkloads.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &assuredworkloadspb.ListWorkloadsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkloads(ctx, req)
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
				},
			},
		},
	}
}
