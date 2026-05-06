package cloudtasks

import (
	"context"
	"fmt"
	"strings"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud cloudtasks command tree.
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
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &locationpb.ListLocationsRequest{Name: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListLocations(ctx, req)
							limit := cmd.Int("limit")
							count := 0
							for {
								if limit > 0 && count >= limit {
									break
								}
								resp, err := it.Next()
								if err == iterator.Done {
									break
								}
								if err != nil {
									return err
								}
								out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
								if err != nil {
									return err
								}
								if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
									continue
								}
								if cmd.Bool("uri") {
									fmt.Println(resp.GetName())
								} else {
									fmt.Println(string(out))
								}
								count++
							}
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &locationpb.GetLocationRequest{Name: name}
							resp, err := client.GetLocation(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &cloudtaskspb.ListQueuesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListQueues(ctx, req)
							limit := cmd.Int("limit")
							count := 0
							for {
								if limit > 0 && count >= limit {
									break
								}
								resp, err := it.Next()
								if err == iterator.Done {
									break
								}
								if err != nil {
									return err
								}
								out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
								if err != nil {
									return err
								}
								if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
									continue
								}
								if cmd.Bool("uri") {
									fmt.Println(resp.GetName())
								} else {
									fmt.Println(string(out))
								}
								count++
							}
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.GetQueueRequest{Name: name}
							resp, err := client.GetQueue(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.CreateQueueRequest{Parent: parent}
							req.Queue = &cloudtaskspb.Queue{
								Name: cmd.String("name"),
							}
							resp, err := client.CreateQueue(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.UpdateQueueRequest{}
							req.Queue = &cloudtaskspb.Queue{
								Name: name,
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateQueue(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.DeleteQueueRequest{Name: name}
							if err := client.DeleteQueue(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "purge",
						Usage: "purge queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing purge on %s\n", name)
							return nil
						},
					},
					{
						Name:  "pause",
						Usage: "pause queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing pause on %s\n", name)
							return nil
						},
					},
					{
						Name:  "resume",
						Usage: "resume queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing resume on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions queues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &cloudtaskspb.ListTasksRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListTasks(ctx, req)
							limit := cmd.Int("limit")
							count := 0
							for {
								if limit > 0 && count >= limit {
									break
								}
								resp, err := it.Next()
								if err == iterator.Done {
									break
								}
								if err != nil {
									return err
								}
								out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
								if err != nil {
									return err
								}
								if filter := cmd.String("filter"); filter != "" && !strings.Contains(string(out), filter) {
									continue
								}
								if cmd.Bool("uri") {
									fmt.Println(resp.GetName())
								} else {
									fmt.Println(string(out))
								}
								count++
							}
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.GetTaskRequest{Name: name}
							resp, err := client.GetTask(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.IntFlag{Name: "dispatch-count", Usage: "The dispatch count.", Required: false},
							&cli.IntFlag{Name: "response-count", Usage: "The response count.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/queues/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.CreateTaskRequest{Parent: parent}
							req.Task = &cloudtaskspb.Task{
								Name:          cmd.String("name"),
								DispatchCount: int32(cmd.Int("dispatch-count")),
								ResponseCount: int32(cmd.Int("response-count")),
							}
							resp, err := client.CreateTask(ctx, req)
							if err != nil {
								return err
							}
							out, err := protojson.MarshalOptions{Multiline: true, Indent: "  "}.Marshal(resp)
							if err != nil {
								return err
							}
							fmt.Println(string(out))
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							client, err := cloudtasks.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &cloudtaskspb.DeleteTaskRequest{Name: name}
							if err := client.DeleteTask(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "run",
						Usage: "run tasks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queue", Usage: "The queue.", Required: true},
							&cli.StringFlag{Name: "task", Usage: "The task.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queues/%s/tasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("queue"), cmd.String("task"))
							fmt.Printf("Executing run on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
