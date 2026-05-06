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

package gkerecommender

import (
	gkerecommender "cloud.google.com/go/gkerecommender/apiv1"
	"cloud.google.com/go/gkerecommender/apiv1/gkerecommenderpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the gkerecommender command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkerecommender",
		Usage: "manage GKE Recommender API resources",
		Commands: []*cli.Command{
			{
				Name:  "benchmarking-data",
				Usage: "Manage benchmarking-data resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch",
						Usage: "fetch benchmarking-data",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance-type", Usage: "The instance type to filter benchmarking data.", Required: false},
							&cli.StringFlag{Name: "pricing-model", Usage: "The pricing model to use for the benchmarking data.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkerecommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkerecommenderpb.FetchBenchmarkingDataRequest{
								InstanceType: cmd.String("instance-type"),
								PricingModel: cmd.String("pricing-model"),
							}

							resp, err := client.FetchBenchmarkingData(ctx, req)
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
				Name:  "model-server-versions",
				Usage: "Manage model-server-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch",
						Usage: "fetch model-server-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "model", Usage: "The model for which to list model server versions.", Required: true},
							&cli.StringFlag{Name: "model-server", Usage: "The model server for which to list versions.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkerecommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkerecommenderpb.FetchModelServerVersionsRequest{
								Model:       cmd.String("model"),
								ModelServer: cmd.String("model-server"),
								PageSize:    runtime.Ptr(int32(cmd.Int("page-size"))),
								PageToken:   runtime.Ptr(cmd.String("page-token")),
							}

							resp, err := client.FetchModelServerVersions(ctx, req)
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
				Name:  "model-servers",
				Usage: "Manage model-servers resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch",
						Usage: "fetch model-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "model", Usage: "The model for which to list model servers.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkerecommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkerecommenderpb.FetchModelServersRequest{
								Model:     cmd.String("model"),
								PageSize:  runtime.Ptr(int32(cmd.Int("page-size"))),
								PageToken: runtime.Ptr(cmd.String("page-token")),
							}

							resp, err := client.FetchModelServers(ctx, req)
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
				Name:  "models",
				Usage: "Manage models resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch",
						Usage: "fetch models",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkerecommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkerecommenderpb.FetchModelsRequest{
								PageSize:  runtime.Ptr(int32(cmd.Int("page-size"))),
								PageToken: runtime.Ptr(cmd.String("page-token")),
							}

							resp, err := client.FetchModels(ctx, req)
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
				Name:  "optimized-manifest",
				Usage: "Manage optimized-manifest resources",
				Commands: []*cli.Command{

					{
						Name:  "generate",
						Usage: "generate optimized-manifest",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "accelerator-type", Usage: "The accelerator type.", Required: true},
							&cli.StringFlag{Name: "kubernetes-namespace", Usage: "The kubernetes namespace to deploy the manifests in.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkerecommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkerecommenderpb.GenerateOptimizedManifestRequest{
								AcceleratorType:     cmd.String("accelerator-type"),
								KubernetesNamespace: cmd.String("kubernetes-namespace"),
							}

							resp, err := client.GenerateOptimizedManifest(ctx, req)
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
				Name:  "profiles",
				Usage: "Manage profiles resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch",
						Usage: "fetch profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "model", Usage: "The model to filter profiles by.", Required: false},
							&cli.StringFlag{Name: "model-server", Usage: "The model server to filter profiles by.", Required: false},
							&cli.StringFlag{Name: "model-server-version", Usage: "The model server version to filter profiles by.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The target number of results to return in a single response.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The value of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := gkerecommender.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &gkerecommenderpb.FetchProfilesRequest{
								Model:              cmd.String("model"),
								ModelServer:        cmd.String("model-server"),
								ModelServerVersion: cmd.String("model-server-version"),
								PageSize:           runtime.Ptr(int32(cmd.Int("page-size"))),
								PageToken:          runtime.Ptr(cmd.String("page-token")),
							}

							limit := cmd.Int("limit")
							it := client.FetchProfiles(ctx, req)
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
