package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	migration "cloud.google.com/go/bigquery/migration/apiv2"
	"cloud.google.com/go/bigquery/migration/apiv2/migrationpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "bigquerymigration",
				Usage: "manage BigQuery Migration API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "subtasks",
						Usage: "Manage subtasks resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe subtasks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
									&cli.StringFlag{Name: "subtask", Usage: "The subtask.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s/subtasks/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"), cmd.String("subtask"))
									client, err := migration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &migrationpb.GetMigrationSubtaskRequest{Name: name}
									resp, err := client.GetMigrationSubtask(ctx, req)
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
								Name:  "list",
								Usage: "list subtasks",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := migration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &migrationpb.ListMigrationSubtasksRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMigrationSubtasks(ctx, req)
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
						},
					},
					{
						Name:  "workflows",
						Usage: "Manage workflows resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create workflows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := migration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &migrationpb.CreateMigrationWorkflowRequest{Parent: parent}
									req.MigrationWorkflow = &migrationpb.MigrationWorkflow{
										DisplayName: cmd.String("display-name"),
									}
									resp, err := client.CreateMigrationWorkflow(ctx, req)
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
								Usage: "describe workflows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
									client, err := migration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &migrationpb.GetMigrationWorkflowRequest{Name: name}
									resp, err := client.GetMigrationWorkflow(ctx, req)
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
								Name:  "list",
								Usage: "list workflows",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete workflows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
									client, err := migration.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &migrationpb.DeleteMigrationWorkflowRequest{Name: name}
									if err := client.DeleteMigrationWorkflow(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "start",
								Usage: "start workflows",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "workflow", Usage: "The workflow.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/workflows/%s", cmd.String("project"), cmd.String("location"), cmd.String("workflow"))
									fmt.Printf("Executing start on %s\n", name)
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
