package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	parametermanager "cloud.google.com/go/parametermanager/apiv1"
	"cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "parametermanager",
				Usage: "manage Parameter Manager API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
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
									client, err := parametermanager.NewClient(ctx)
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
									client, err := parametermanager.NewClient(ctx)
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
						Name:  "parameters",
						Usage: "Manage parameters resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list parameters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &parametermanagerpb.ListParametersRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListParameters(ctx, req)
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
								Usage: "describe parameters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.GetParameterRequest{Name: name}
									resp, err := client.GetParameter(ctx, req)
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
								Usage: "create parameters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter-id", Usage: "The parameter id.", Required: true},
									&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.CreateParameterRequest{Parent: parent}
									req.ParameterId = cmd.String("parameter-id")
									req.Parameter = &parametermanagerpb.Parameter{
										KmsKey: cmd.String("kms-key"),
									}
									resp, err := client.CreateParameter(ctx, req)
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
								Usage: "update parameters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.UpdateParameterRequest{}
									req.Parameter = &parametermanagerpb.Parameter{
										Name:   name,
										KmsKey: cmd.String("kms-key"),
									}
									var paths []string
									if cmd.IsSet("kms-key") {
										paths = append(paths, "kms_key")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateParameter(ctx, req)
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
								Usage: "delete parameters",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.DeleteParameterRequest{Name: name}
									if err := client.DeleteParameter(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/parameters/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &parametermanagerpb.ListParameterVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListParameterVersions(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.StringFlag{Name: "parameter_version", Usage: "The parameter_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"), cmd.String("parameter_version"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.GetParameterVersionRequest{Name: name}
									resp, err := client.GetParameterVersion(ctx, req)
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
								Name:  "render",
								Usage: "render versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.StringFlag{Name: "parameter_version", Usage: "The parameter_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"), cmd.String("parameter_version"))
									fmt.Printf("Executing render on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.StringFlag{Name: "parameter-version-id", Usage: "The parameter version id.", Required: true},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/parameters/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.CreateParameterVersionRequest{Parent: parent}
									req.ParameterVersionId = cmd.String("parameter-version-id")
									req.ParameterVersion = &parametermanagerpb.ParameterVersion{
										Disabled: cmd.Bool("disabled"),
									}
									resp, err := client.CreateParameterVersion(ctx, req)
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
								Usage: "update versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.StringFlag{Name: "parameter_version", Usage: "The parameter_version.", Required: true},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"), cmd.String("parameter_version"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.UpdateParameterVersionRequest{}
									req.ParameterVersion = &parametermanagerpb.ParameterVersion{
										Name:     name,
										Disabled: cmd.Bool("disabled"),
									}
									var paths []string
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateParameterVersion(ctx, req)
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
								Usage: "delete versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "parameter", Usage: "The parameter.", Required: true},
									&cli.StringFlag{Name: "parameter_version", Usage: "The parameter_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/parameters/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("parameter"), cmd.String("parameter_version"))
									client, err := parametermanager.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &parametermanagerpb.DeleteParameterVersionRequest{Name: name}
									if err := client.DeleteParameterVersion(ctx, req); err != nil {
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
