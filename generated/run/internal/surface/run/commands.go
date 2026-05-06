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

package run

import (
	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the run command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "manage Cloud Run Admin API resources",
		Commands: []*cli.Command{
			{
				Name:  "builds",
				Usage: "Manage builds resources",
				Commands: []*cli.Command{

					{
						Name:  "submit",
						Usage: "submit builds",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "client", Usage: "The client that initiated the build request.", Required: false},
							&cli.StringFlag{Name: "image-uri", Usage: "Artifact Registry URI to store the built image.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "machine-type", Usage: "The machine type from default pool to use for the build.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "release-track", Usage: "The release track of the client that initiated the build request.", Required: false},
							&cli.StringFlag{Name: "service-account", Usage: "The service account to use for the build.", Required: false},
							&cli.StringSliceFlag{Name: "tags", Usage: "Additional tags to annotate the build.", Required: false},
							&cli.StringFlag{Name: "worker-pool", Usage: "Name of the Cloud Build Custom Worker Pool that should be used to.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.SubmitBuildRequest{
								Parent:         parent,
								ImageUri:       cmd.String("image-uri"),
								ServiceAccount: cmd.String("service-account"),
								WorkerPool:     cmd.String("worker-pool"),
								Tags:           cmd.StringSlice("tags"),
								MachineType:    cmd.String("machine-type"),
								ReleaseTrack:   runpb.LaunchStage(runpb.LaunchStage_value[cmd.String("release-track")]),
								Client:         cmd.String("client"),
							}

							resp, err := client.SubmitBuild(ctx, req)
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
				Name:  "executions",
				Usage: "Manage executions resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetExecutionRequest{
								Name: name,
							}

							resp, err := client.GetExecution(ctx, req)
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
						Usage: "list executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Executions to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListExecutions.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with active ones.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListExecutionsRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListExecutions(ctx, req)
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
						Usage: "delete executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the resource.", Required: false},
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.DeleteExecutionRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteExecution(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the resource.", Required: false},
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.CancelExecutionRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.CancelExecution(ctx, req)
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
			{
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance-id", Usage: "The unique identifier for the Instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.CreateInstanceRequest{
								Parent:       parent,
								InstanceId:   cmd.String("instance-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateInstance(ctx, req)
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
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.DeleteInstanceRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteInstance(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetInstanceRequest{
								Name: name,
							}

							resp, err := client.GetInstance(ctx, req)
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
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Instances to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListInstances.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListInstancesRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListInstances(ctx, req)
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
						Name:  "stop",
						Usage: "stop instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the resource.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.StopInstanceRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.StopInstance(ctx, req)
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
						Name:  "start",
						Usage: "start instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the resource.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.StartInstanceRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.StartInstance(ctx, req)
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
			{
				Name:  "jobs",
				Usage: "Manage jobs resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job-id", Usage: "The unique identifier for the Job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.CreateJobRequest{
								Parent:       parent,
								JobId:        cmd.String("job-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateJob(ctx, req)
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
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetJobRequest{
								Name: name,
							}

							resp, err := client.GetJob(ctx, req)
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
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Jobs to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListJobs.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with active ones.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListJobsRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListJobs(ctx, req)
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
						Usage: "update jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and if the Job does not exist, it will create a.", Required: false},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "job.name" not yet supported.
							job_name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							fmt.Printf("Executing update on %s\n", job_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the.", Required: false},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.DeleteJobRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteJob(ctx, req)
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
						Name:  "run",
						Usage: "run jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the.", Required: false},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.RunJobRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.RunJob(ctx, req)
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
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
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

					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "revisions",
				Usage: "Manage revisions resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/services/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetRevisionRequest{
								Name: name,
							}

							resp, err := client.GetRevision(ctx, req)
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
						Usage: "list revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of revisions to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListRevisions.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with active ones.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListRevisionsRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListRevisions(ctx, req)
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
						Usage: "delete revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The ID of the revision.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/services/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"), cmd.String("revision"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.DeleteRevisionRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteRevision(ctx, req)
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
			{
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service-id", Usage: "The unique identifier for the Service.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.CreateServiceRequest{
								Parent:       parent,
								ServiceId:    cmd.String("service-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateService(ctx, req)
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
						Usage: "describe services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetServiceRequest{
								Name: name,
							}

							resp, err := client.GetService(ctx, req)
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
						Usage: "list services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Services to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListServices.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with active ones.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListServicesRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListServices(ctx, req)
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
						Usage: "update services",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and if the Service does not exist, it will create.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "service.name" not yet supported.
							service_name := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							fmt.Printf("Executing update on %s\n", service_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.DeleteServiceRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteService(ctx, req)
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
						Usage: "describe services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "service", Usage: "The ID of the service.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/services/%s", cmd.String("project"), cmd.String("location"), cmd.String("service"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
				Name:  "tasks",
				Usage: "Manage tasks resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s/executions/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("execution"), cmd.String("task"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetTaskRequest{
								Name: name,
							}

							resp, err := client.GetTask(ctx, req)
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
						Usage: "list tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "execution", Usage: "The ID of the execution.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The ID of the job.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of Tasks to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListTasks.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with active ones.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/jobs/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"), cmd.String("execution"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListTasksRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListTasks(ctx, req)
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
				Name:  "worker-pools",
				Usage: "Manage worker-pools resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
							&cli.StringFlag{Name: "worker-pool-id", Usage: "The unique identifier for the WorkerPool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.CreateWorkerPoolRequest{
								Parent:       parent,
								WorkerPoolId: cmd.String("worker-pool-id"),
								ValidateOnly: cmd.Bool("validate-only"),
							}

							op, err := client.CreateWorkerPool(ctx, req)
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
						Usage: "describe worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "worker-pool", Usage: "The ID of the worker pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("worker-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetWorkerPoolRequest{
								Name: name,
							}

							resp, err := client.GetWorkerPool(ctx, req)
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
						Usage: "list worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of WorkerPools to return in this call.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous call to ListWorkerPools.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "If true, returns deleted (but unexpired) resources along with active ones.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.ListWorkerPoolsRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListWorkerPools(ctx, req)
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
						Usage: "update worker-pools",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and if the WorkerPool does not exist, it will.", Required: false},
							&cli.BoolFlag{Name: "force-new-revision", Usage: "If set to true, a new revision will be created from the template.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated and default values.", Required: false},
							&cli.StringFlag{Name: "worker-pool", Usage: "The ID of the worker pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "worker_pool.name" not yet supported.
							worker_pool_name := fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("worker-pool"))
							fmt.Printf("Executing update on %s\n", worker_pool_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "A system-generated fingerprint for this version of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "Indicates that the request should be validated without actually.", Required: false},
							&cli.StringFlag{Name: "worker-pool", Usage: "The ID of the worker pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("worker-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.DeleteWorkerPoolRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteWorkerPool(ctx, req)
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
						Usage: "describe worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "worker-pool", Usage: "The ID of the worker pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("worker-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "worker-pool", Usage: "The ID of the worker pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("worker-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions worker-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "worker-pool", Usage: "The ID of the worker pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", cmd.String("project"), cmd.String("location"), cmd.String("worker-pool"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := run.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &runpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
