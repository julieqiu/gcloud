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

package service

import (
	service "cloud.google.com/go/service/apiv1"
	"cloud.google.com/go/service/apiv1/servicepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the composer command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "composer",
		Usage: "manage Cloud Composer API resources",
		Commands: []*cli.Command{
			{
				Name:  "environments",
				Usage: "Manage environments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create environments",
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
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.CreateEnvironmentRequest{
								Parent: parent,
							}

							op, err := client.CreateEnvironment(ctx, req)
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
						Name:  "describe",
						Usage: "describe environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.GetEnvironmentRequest{
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
						Name:  "list",
						Usage: "list environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of environments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.ListEnvironmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEnvironments(ctx, req)
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
						Name:  "update",
						Usage: "update environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.UpdateEnvironmentRequest{
								Name: name,
							}

							op, err := client.UpdateEnvironment(ctx, req)
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
						Name:  "delete",
						Usage: "delete environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteEnvironment %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.DeleteEnvironmentRequest{
								Name: name,
							}

							op, err := client.DeleteEnvironment(ctx, req)
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
						Name:  "execute-airflow-command",
						Usage: "execute-airflow-command environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "command", Usage: "Airflow command.", Required: false},
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "parameters", Usage: "Parameters for the Airflow command/subcommand as an array of arguments.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "subcommand", Usage: "Airflow subcommand.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing execute-airflow-command on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "stop-airflow-command",
						Usage: "stop-airflow-command environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "execution-id", Usage: "The unique ID of the command execution.", Required: false},
							&cli.BoolFlag{Name: "force", Usage: "If true, the execution is terminated forcefully (SIGKILL).", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pod", Usage: "The name of the pod where the command is executed.", Required: false},
							&cli.StringFlag{Name: "pod-namespace", Usage: "The namespace of the pod where the command is executed.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing stop-airflow-command on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "poll-airflow-command",
						Usage: "poll-airflow-command environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "execution-id", Usage: "The unique ID of the command execution.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "next-line-number", Usage: "Line number from which new logs should be fetched.", Required: false},
							&cli.StringFlag{Name: "pod", Usage: "The name of the pod where the command is executed.", Required: false},
							&cli.StringFlag{Name: "pod-namespace", Usage: "The namespace of the pod where the command is executed.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing poll-airflow-command on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "check-upgrade",
						Usage: "check-upgrade environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "image-version", Usage: "The version of the software running in the environment.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing check-upgrade on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "save-snapshot",
						Usage: "save-snapshot environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot-location", Usage: "Location in a Cloud Storage where the snapshot is going to be stored, e.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing save-snapshot on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "load-snapshot",
						Usage: "load-snapshot environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "skip-airflow-overrides-setting", Usage: "Whether or not to skip setting Airflow overrides when loading the.", Required: false},
							&cli.BoolFlag{Name: "skip-environment-variables-setting", Usage: "Whether or not to skip setting environment variables when loading the.", Required: false},
							&cli.BoolFlag{Name: "skip-gcs-data-copying", Usage: "Whether or not to skip copying Cloud Storage data when loading the.", Required: false},
							&cli.BoolFlag{Name: "skip-pypi-packages-installation", Usage: "Whether or not to skip installing Pypi packages when loading the.", Required: false},
							&cli.StringFlag{Name: "snapshot-path", Usage: "A Cloud Storage path to a snapshot to load, e.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing load-snapshot on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "database-failover",
						Usage: "database-failover environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing database-failover on %s\n", environment)
							return nil
						},
					},

					{
						Name:  "fetch-database-properties",
						Usage: "fetch-database-properties environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							environment := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							fmt.Printf("Executing fetch-database-properties on %s\n", environment)
							return nil
						},
					},
				},
			},
			{
				Name:  "image-versions",
				Usage: "Manage image-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list image-versions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "include-past-releases", Usage: "Whether or not image versions from old releases should be included.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of image_versions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.ListImageVersionsRequest{
								Parent:              parent,
								PageSize:            int32(cmd.Int("page-size")),
								PageToken:           cmd.String("page-token"),
								IncludePastReleases: cmd.Bool("include-past-releases"),
							}

							limit := cmd.Int("limit")
							it := client.ListImageVersions(ctx, req)
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
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
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
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "user-workloads-config-maps",
				Usage: "Manage user-workloads-config-maps resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.CreateUserWorkloadsConfigMapRequest{
								Parent: parent,
							}

							resp, err := client.CreateUserWorkloadsConfigMap(ctx, req)
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
						Usage: "describe user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-workloads-config-map", Usage: "The ID of the user workloads config map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user-workloads-config-map"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.GetUserWorkloadsConfigMapRequest{
								Name: name,
							}

							resp, err := client.GetUserWorkloadsConfigMap(ctx, req)
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
						Usage: "list user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of ConfigMaps to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.ListUserWorkloadsConfigMapsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUserWorkloadsConfigMaps(ctx, req)
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
						Name:  "update",
						Usage: "update user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-workloads-config-map", Usage: "The ID of the user workloads config map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "user_workloads_config_map.name" not yet supported.
							user_workloads_config_map_name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user-workloads-config-map"))
							fmt.Printf("Executing update on %s\n", user_workloads_config_map_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-workloads-config-map", Usage: "The ID of the user workloads config map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user-workloads-config-map"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteUserWorkloadsConfigMap on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.DeleteUserWorkloadsConfigMapRequest{
								Name: name,
							}

							if err := client.DeleteUserWorkloadsConfigMap(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "user-workloads-secrets",
				Usage: "Manage user-workloads-secrets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.CreateUserWorkloadsSecretRequest{
								Parent: parent,
							}

							resp, err := client.CreateUserWorkloadsSecret(ctx, req)
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
						Usage: "describe user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-workloads-secret", Usage: "The ID of the user workloads secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user-workloads-secret"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.GetUserWorkloadsSecretRequest{
								Name: name,
							}

							resp, err := client.GetUserWorkloadsSecret(ctx, req)
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
						Usage: "list user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Secrets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.ListUserWorkloadsSecretsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListUserWorkloadsSecrets(ctx, req)
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
						Name:  "update",
						Usage: "update user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-workloads-secret", Usage: "The ID of the user workloads secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "user_workloads_secret.name" not yet supported.
							user_workloads_secret_name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user-workloads-secret"))
							fmt.Printf("Executing update on %s\n", user_workloads_secret_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "user-workloads-secret", Usage: "The ID of the user workloads secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user-workloads-secret"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteUserWorkloadsSecret on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.DeleteUserWorkloadsSecretRequest{
								Name: name,
							}

							if err := client.DeleteUserWorkloadsSecret(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
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
						Name:  "list",
						Usage: "list workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "environment", Usage: "The ID of the environment.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The list filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of environments to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request,.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := service.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &servicepb.ListWorkloadsRequest{
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
