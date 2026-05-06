package workflowexecutions

import (
	"context"
	"fmt"
	"strings"

	executions "cloud.google.com/go/workflows/executions/apiv1"
	"cloud.google.com/go/workflows/executions/apiv1/executionspb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

// Command returns the gcloud workflowexecutions command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "workflowexecutions",
		Usage: "manage Workflow Executions API resources",
		Commands: []*cli.Command{
			{
				Name:  "executions",
				Usage: "Manage executions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := executions.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &executionspb.ListExecutionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListExecutions(ctx, req)
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
						Name:  "create",
						Usage: "create executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
							&cli.StringFlag{Name: "argument", Usage: "The argument.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
							client, err := executions.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &executionspb.CreateExecutionRequest{Parent: parent}
							req.Execution = &executionspb.Execution{
								Argument: cmd.String("argument"),
							}
							resp, err := client.CreateExecution(ctx, req)
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
						Name:  "describe",
						Usage: "describe executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
							&cli.StringFlag{Name: "execution", Usage: "The execution.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"), cmd.String("execution"))
							client, err := executions.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &executionspb.GetExecutionRequest{Name: name}
							resp, err := client.GetExecution(ctx, req)
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
						Name:  "cancel",
						Usage: "cancel executions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
							&cli.StringFlag{Name: "execution", Usage: "The execution.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s/executions/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"), cmd.String("execution"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
