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

package tasks

import (
	tasks "cloud.google.com/go/tasks/apiv2"
	"cloud.google.com/go/tasks/apiv2/taskspb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the cloudtasks command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudtasks",
		Usage: "manage Cloud Tasks API resources",
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
			{
				Name:  "queues",
				Usage: "Manage queues resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "`filter` can be used to specify a subset of queues.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Requested page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.ListQueuesRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListQueues(ctx, req)
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
						Usage: "describe queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.GetQueueRequest{
								Name: name,
							}

							resp, err := client.GetQueue(ctx, req)
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
						Usage: "create queues",
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
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.CreateQueueRequest{
								Parent: parent,
							}

							resp, err := client.CreateQueue(ctx, req)
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
						Usage: "update queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "queue.name" not yet supported.
							queue_name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing update on %s\n", queue_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteQueue on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.DeleteQueueRequest{
								Name: name,
							}

							if err := client.DeleteQueue(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "purge",
						Usage: "purge queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.PurgeQueueRequest{
								Name: name,
							}

							resp, err := client.PurgeQueue(ctx, req)
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
						Name:  "pause",
						Usage: "pause queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.PauseQueueRequest{
								Name: name,
							}

							resp, err := client.PauseQueue(ctx, req)
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
						Name:  "resume",
						Usage: "resume queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.ResumeQueueRequest{
								Name: name,
							}

							resp, err := client.ResumeQueue(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.GetIamPolicyRequest{
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
						Usage: "set-iam-policy queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.SetIamPolicyRequest{
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
						Usage: "test-iam-permissions queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.TestIamPermissionsRequest{
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
						Name:  "list",
						Usage: "list tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
							&cli.StringFlag{Name: "response-view", Usage: "The response_view specifies which subset of the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.ListTasksRequest{
								Parent:       parent,
								ResponseView: taskspb.Task_View(taskspb.Task_View_value[cmd.String("response-view")]),
								PageSize:     int32(cmd.Int("page-size")),
								PageToken:    cmd.String("page-token"),
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

					{
						Name:  "describe",
						Usage: "describe tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
							&cli.StringFlag{Name: "response-view", Usage: "The response_view specifies which subset of the.", Required: false},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.GetTaskRequest{
								Name:         name,
								ResponseView: taskspb.Task_View(taskspb.Task_View_value[cmd.String("response-view")]),
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
						Name:  "create",
						Usage: "create tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
							&cli.StringFlag{Name: "response-view", Usage: "The response_view specifies which subset of the.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.CreateTaskRequest{
								Parent:       parent,
								ResponseView: taskspb.Task_View(taskspb.Task_View_value[cmd.String("response-view")]),
							}

							resp, err := client.CreateTask(ctx, req)
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
						Usage: "delete tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteTask on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.DeleteTaskRequest{
								Name: name,
							}

							if err := client.DeleteTask(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "run",
						Usage: "run tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The ID of the queue.", Required: true},
							&cli.StringFlag{Name: "response-view", Usage: "The response_view specifies which subset of the.", Required: false},
							&cli.StringFlag{Name: "task", Usage: "The ID of the task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := tasks.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &taskspb.RunTaskRequest{
								Name:         name,
								ResponseView: taskspb.Task_View(taskspb.Task_View_value[cmd.String("response-view")]),
							}

							resp, err := client.RunTask(ctx, req)
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
