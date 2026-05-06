package composer

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	service "cloud.google.com/go/orchestration/airflow/service/apiv1"
	"cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud composer command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "uuid", Usage: "The uuid.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.CreateEnvironmentRequest{Parent: parent}
							req.Environment = &servicepb.Environment{
								Uuid: cmd.String("uuid"),
							}
							op, err := client.CreateEnvironment(ctx, req)
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
						Name:  "describe",
						Usage: "describe environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update environments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "uuid", Usage: "The uuid.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.UpdateEnvironmentRequest{}
							req.Environment = &servicepb.Environment{
								Name: name,
								Uuid: cmd.String("uuid"),
							}
							var paths []string
							if cmd.IsSet("uuid") {
								paths = append(paths, "uuid")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateEnvironment(ctx, req)
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
						Usage: "delete environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing delete...")
							return nil
						},
					},
					{
						Name:  "execute-airflow-command",
						Usage: "execute-airflow-command environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing execute-airflow-command...")
							return nil
						},
					},
					{
						Name:  "stop-airflow-command",
						Usage: "stop-airflow-command environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing stop-airflow-command...")
							return nil
						},
					},
					{
						Name:  "poll-airflow-command",
						Usage: "poll-airflow-command environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing poll-airflow-command...")
							return nil
						},
					},
					{
						Name:  "check-upgrade",
						Usage: "check-upgrade environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing check-upgrade...")
							return nil
						},
					},
					{
						Name:  "save-snapshot",
						Usage: "save-snapshot environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing save-snapshot...")
							return nil
						},
					},
					{
						Name:  "load-snapshot",
						Usage: "load-snapshot environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing load-snapshot...")
							return nil
						},
					},
					{
						Name:  "database-failover",
						Usage: "database-failover environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing database-failover...")
							return nil
						},
					},
					{
						Name:  "fetch-database-properties",
						Usage: "fetch-database-properties environments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch-database-properties...")
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
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
							client, err := service.NewEnvironmentsClient(ctx)
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
							client, err := service.NewEnvironmentsClient(ctx)
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
							client, err := service.NewEnvironmentsClient(ctx)
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
							client, err := service.NewImageVersionsClient(ctx)
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
							client, err := service.NewImageVersionsClient(ctx)
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
							client, err := service.NewImageVersionsClient(ctx)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.CreateUserWorkloadsConfigMapRequest{Parent: parent}
							resp, err := client.CreateUserWorkloadsConfigMap(ctx, req)
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
						Usage: "describe user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "user_workloads_config_map", Usage: "The user_workloads_config_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user_workloads_config_map"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.GetUserWorkloadsConfigMapRequest{Name: name}
							resp, err := client.GetUserWorkloadsConfigMap(ctx, req)
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
						Usage: "list user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &servicepb.ListUserWorkloadsConfigMapsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListUserWorkloadsConfigMaps(ctx, req)
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
						Name:  "update",
						Usage: "update user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "user_workloads_config_map", Usage: "The user_workloads_config_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user_workloads_config_map"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.UpdateUserWorkloadsConfigMapRequest{}
							req.UserWorkloadsConfigMap = &servicepb.UserWorkloadsConfigMap{
								Name: name,
							}
							resp, err := client.UpdateUserWorkloadsConfigMap(ctx, req)
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
						Usage: "delete user-workloads-config-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "user_workloads_config_map", Usage: "The user_workloads_config_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsConfigMaps/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user_workloads_config_map"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.DeleteUserWorkloadsConfigMapRequest{Name: name}
							if err := client.DeleteUserWorkloadsConfigMap(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/environments/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.CreateUserWorkloadsSecretRequest{Parent: parent}
							resp, err := client.CreateUserWorkloadsSecret(ctx, req)
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
						Usage: "describe user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "user_workloads_secret", Usage: "The user_workloads_secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user_workloads_secret"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.GetUserWorkloadsSecretRequest{Name: name}
							resp, err := client.GetUserWorkloadsSecret(ctx, req)
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
						Usage: "list user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &servicepb.ListUserWorkloadsSecretsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListUserWorkloadsSecrets(ctx, req)
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
						Name:  "update",
						Usage: "update user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "user_workloads_secret", Usage: "The user_workloads_secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user_workloads_secret"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.UpdateUserWorkloadsSecretRequest{}
							req.UserWorkloadsSecret = &servicepb.UserWorkloadsSecret{
								Name: name,
							}
							resp, err := client.UpdateUserWorkloadsSecret(ctx, req)
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
						Usage: "delete user-workloads-secrets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "environment", Usage: "The environment.", Required: true},
							&cli.StringFlag{Name: "user_workloads_secret", Usage: "The user_workloads_secret.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", cmd.String("project"), cmd.String("location"), cmd.String("environment"), cmd.String("user_workloads_secret"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &servicepb.DeleteUserWorkloadsSecretRequest{Name: name}
							if err := client.DeleteUserWorkloadsSecret(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := service.NewEnvironmentsClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &servicepb.ListWorkloadsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListWorkloads(ctx, req)
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
		},
	}
}
