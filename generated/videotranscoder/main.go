package main

import (
	"context"
	"fmt"
	"os"

	transcoder "cloud.google.com/go/video/transcoder/apiv1"
	"cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "transcoder",
				Usage: "manage Transcoder API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "job-templates",
						Usage: "Manage job-templates resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create job-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job-template-id", Usage: "The job template id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := transcoder.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &transcoderpb.CreateJobTemplateRequest{Parent: parent}
									req.JobTemplateId = cmd.String("job-template-id")
									req.JobTemplate = &transcoderpb.JobTemplate{
										Name: cmd.String("name"),
									}
									resp, err := client.CreateJobTemplate(ctx, req)
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
								Usage: "list job-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe job-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job_template", Usage: "The job_template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/jobTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("job_template"))
									client, err := transcoder.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &transcoderpb.GetJobTemplateRequest{Name: name}
									resp, err := client.GetJobTemplate(ctx, req)
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
								Usage: "delete job-templates",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job_template", Usage: "The job_template.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/jobTemplates/%s", cmd.String("project"), cmd.String("location"), cmd.String("job_template"))
									client, err := transcoder.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &transcoderpb.DeleteJobTemplateRequest{Name: name}
									if err := client.DeleteJobTemplate(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "input-uri", Usage: "The input uri.", Required: false},
									&cli.StringFlag{Name: "output-uri", Usage: "The output uri.", Required: false},
									&cli.IntFlag{Name: "ttl-after-completion-days", Usage: "The ttl after completion days.", Required: false},
									&cli.IntFlag{Name: "batch-mode-priority", Usage: "The batch mode priority.", Required: false},
									&cli.BoolFlag{Name: "fill-content-gaps", Usage: "The fill content gaps.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := transcoder.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &transcoderpb.CreateJobRequest{Parent: parent}
									req.Job = &transcoderpb.Job{
										Name:                   cmd.String("name"),
										InputUri:               cmd.String("input-uri"),
										OutputUri:              cmd.String("output-uri"),
										TtlAfterCompletionDays: int32(cmd.Int("ttl-after-completion-days")),
										BatchModePriority:      int32(cmd.Int("batch-mode-priority")),
										FillContentGaps:        cmd.Bool("fill-content-gaps"),
									}
									resp, err := client.CreateJob(ctx, req)
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
								Usage: "list jobs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
									client, err := transcoder.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &transcoderpb.GetJobRequest{Name: name}
									resp, err := client.GetJob(ctx, req)
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
								Usage: "delete jobs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", cmd.String("project"), cmd.String("location"), cmd.String("job"))
									client, err := transcoder.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &transcoderpb.DeleteJobRequest{Name: name}
									if err := client.DeleteJob(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
