package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	datafusion "cloud.google.com/go/datafusion/apiv1"
	"cloud.google.com/go/datafusion/apiv1/datafusionpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "datafusion",
				Usage: "manage Cloud Data Fusion API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "instances",
						Usage: "Manage instances resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &datafusionpb.GetInstanceRequest{Name: name}
									resp, err := client.GetInstance(ctx, req)
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
								Usage: "create instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance-id", Usage: "The instance id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "enable-stackdriver-logging", Usage: "The enable stackdriver logging.", Required: false},
									&cli.BoolFlag{Name: "enable-stackdriver-monitoring", Usage: "The enable stackdriver monitoring.", Required: false},
									&cli.BoolFlag{Name: "private-instance", Usage: "The private instance.", Required: false},
									&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: false},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "dataproc-service-account", Usage: "The dataproc service account.", Required: false},
									&cli.BoolFlag{Name: "enable-rbac", Usage: "The enable rbac.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &datafusionpb.CreateInstanceRequest{Parent: parent}
									req.InstanceId = cmd.String("instance-id")
									req.Instance = &datafusionpb.Instance{
										Description:                 cmd.String("description"),
										EnableStackdriverLogging:    cmd.Bool("enable-stackdriver-logging"),
										EnableStackdriverMonitoring: cmd.Bool("enable-stackdriver-monitoring"),
										PrivateInstance:             cmd.Bool("private-instance"),
										Zone:                        cmd.String("zone"),
										Version:                     cmd.String("version"),
										DisplayName:                 cmd.String("display-name"),
										DataprocServiceAccount:      cmd.String("dataproc-service-account"),
										EnableRbac:                  cmd.Bool("enable-rbac"),
									}
									op, err := client.CreateInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Usage: "delete instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &datafusionpb.DeleteInstanceRequest{Name: name}
									op, err := client.DeleteInstance(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "enable-stackdriver-logging", Usage: "The enable stackdriver logging.", Required: false},
									&cli.BoolFlag{Name: "enable-stackdriver-monitoring", Usage: "The enable stackdriver monitoring.", Required: false},
									&cli.BoolFlag{Name: "private-instance", Usage: "The private instance.", Required: false},
									&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: false},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "dataproc-service-account", Usage: "The dataproc service account.", Required: false},
									&cli.BoolFlag{Name: "enable-rbac", Usage: "The enable rbac.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &datafusionpb.UpdateInstanceRequest{}
									req.Instance = &datafusionpb.Instance{
										Name:                        name,
										Description:                 cmd.String("description"),
										EnableStackdriverLogging:    cmd.Bool("enable-stackdriver-logging"),
										EnableStackdriverMonitoring: cmd.Bool("enable-stackdriver-monitoring"),
										PrivateInstance:             cmd.Bool("private-instance"),
										Zone:                        cmd.String("zone"),
										Version:                     cmd.String("version"),
										DisplayName:                 cmd.String("display-name"),
										DataprocServiceAccount:      cmd.String("dataproc-service-account"),
										EnableRbac:                  cmd.Bool("enable-rbac"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("enable-stackdriver-logging") {
										paths = append(paths, "enable_stackdriver_logging")
									}
									if cmd.IsSet("enable-stackdriver-monitoring") {
										paths = append(paths, "enable_stackdriver_monitoring")
									}
									if cmd.IsSet("private-instance") {
										paths = append(paths, "private_instance")
									}
									if cmd.IsSet("zone") {
										paths = append(paths, "zone")
									}
									if cmd.IsSet("version") {
										paths = append(paths, "version")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("dataproc-service-account") {
										paths = append(paths, "dataproc_service_account")
									}
									if cmd.IsSet("enable-rbac") {
										paths = append(paths, "enable_rbac")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
								Name:  "restart",
								Usage: "restart instances",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &datafusionpb.RestartInstanceRequest{Name: name}
									op, err := client.RestartInstance(ctx, req)
									if err != nil {
										return err
									}
									resp, err := op.Wait(ctx)
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
						Name:  "operations",
						Usage: "Manage operations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &longrunningpb.ListOperationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListOperations(ctx, req)
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
								Usage: "describe operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.GetOperationRequest{Name: name}
									resp, err := client.GetOperation(ctx, req)
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
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "cancel",
								Usage: "cancel operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := datafusion.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.CancelOperationRequest{Name: name}
									if err := client.CancelOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Cancelled %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "versions",
						Usage: "Manage versions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list versions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
