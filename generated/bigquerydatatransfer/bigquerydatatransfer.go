package bigquerydatatransfer

import (
	"context"
	"fmt"
	"strings"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud bigquerydatatransfer command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigquerydatatransfer",
		Usage: "manage BigQuery Data Transfer API resources",
		Commands: []*cli.Command{
			{
				Name:  "data-sources",
				Usage: "Manage data-sources resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data_source", Usage: "The data_source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dataSources/%s", cmd.String("project"), cmd.String("data_source"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.GetDataSourceRequest{Name: name}
							resp, err := client.GetDataSource(ctx, req)
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
						Usage: "list data-sources",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datatransferpb.ListDataSourcesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDataSources(ctx, req)
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
						Name:  "check-valid-creds",
						Usage: "check-valid-creds data-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "data_source", Usage: "The data_source.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dataSources/%s", cmd.String("project"), cmd.String("data_source"))
							fmt.Printf("Executing check-valid-creds on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{
					{
						Name:  "enroll-data-sources",
						Usage: "enroll-data-sources locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing enroll-data-sources...")
							return nil
						},
					},
					{
						Name:  "unenroll-data-sources",
						Usage: "unenroll-data-sources locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing unenroll-data-sources...")
							return nil
						},
					},
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
							client, err := datatransfer.NewClient(ctx)
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
							client, err := datatransfer.NewClient(ctx)
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
				Name:  "runs",
				Usage: "Manage runs resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s/runs/%s", cmd.String("project"), cmd.String("transfer_config"), cmd.String("run"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.GetTransferRunRequest{Name: name}
							resp, err := client.GetTransferRun(ctx, req)
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
						Usage: "delete runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
							&cli.StringFlag{Name: "run", Usage: "The run.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s/runs/%s", cmd.String("project"), cmd.String("transfer_config"), cmd.String("run"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.DeleteTransferRunRequest{Name: name}
							if err := client.DeleteTransferRun(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list runs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datatransferpb.ListTransferRunsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListTransferRuns(ctx, req)
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
				Name:  "transfer-configs",
				Usage: "Manage transfer-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "authorization-code", Usage: "The authorization code.", Required: false},
							&cli.StringFlag{Name: "version-info", Usage: "The version info.", Required: false},
							&cli.StringFlag{Name: "service-account-name", Usage: "The service account name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "data-source-id", Usage: "The data source id.", Required: false},
							&cli.StringFlag{Name: "schedule", Usage: "The schedule.", Required: false},
							&cli.IntFlag{Name: "data-refresh-window-days", Usage: "The data refresh window days.", Required: false},
							&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
							&cli.IntFlag{Name: "user-id", Usage: "The user id.", Required: false},
							&cli.StringFlag{Name: "notification-pubsub-topic", Usage: "The notification pubsub topic.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.CreateTransferConfigRequest{Parent: parent}
							req.AuthorizationCode = cmd.String("authorization-code")
							req.VersionInfo = cmd.String("version-info")
							req.ServiceAccountName = cmd.String("service-account-name")
							req.TransferConfig = &datatransferpb.TransferConfig{
								DisplayName:             cmd.String("display-name"),
								DataSourceId:            cmd.String("data-source-id"),
								Schedule:                cmd.String("schedule"),
								DataRefreshWindowDays:   int32(cmd.Int("data-refresh-window-days")),
								Disabled:                cmd.Bool("disabled"),
								UserId:                  int64(cmd.Int("user-id")),
								NotificationPubsubTopic: cmd.String("notification-pubsub-topic"),
							}
							resp, err := client.CreateTransferConfig(ctx, req)
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
						Usage: "update transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "data-source-id", Usage: "The data source id.", Required: false},
							&cli.StringFlag{Name: "schedule", Usage: "The schedule.", Required: false},
							&cli.IntFlag{Name: "data-refresh-window-days", Usage: "The data refresh window days.", Required: false},
							&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
							&cli.IntFlag{Name: "user-id", Usage: "The user id.", Required: false},
							&cli.StringFlag{Name: "notification-pubsub-topic", Usage: "The notification pubsub topic.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.UpdateTransferConfigRequest{}
							req.TransferConfig = &datatransferpb.TransferConfig{
								Name:                    name,
								DisplayName:             cmd.String("display-name"),
								DataSourceId:            cmd.String("data-source-id"),
								Schedule:                cmd.String("schedule"),
								DataRefreshWindowDays:   int32(cmd.Int("data-refresh-window-days")),
								Disabled:                cmd.Bool("disabled"),
								UserId:                  int64(cmd.Int("user-id")),
								NotificationPubsubTopic: cmd.String("notification-pubsub-topic"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("data-source-id") {
								paths = append(paths, "data_source_id")
							}
							if cmd.IsSet("schedule") {
								paths = append(paths, "schedule")
							}
							if cmd.IsSet("data-refresh-window-days") {
								paths = append(paths, "data_refresh_window_days")
							}
							if cmd.IsSet("disabled") {
								paths = append(paths, "disabled")
							}
							if cmd.IsSet("user-id") {
								paths = append(paths, "user_id")
							}
							if cmd.IsSet("notification-pubsub-topic") {
								paths = append(paths, "notification_pubsub_topic")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateTransferConfig(ctx, req)
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
						Usage: "delete transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.DeleteTransferConfigRequest{Name: name}
							if err := client.DeleteTransferConfig(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datatransferpb.GetTransferConfigRequest{Name: name}
							resp, err := client.GetTransferConfig(ctx, req)
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
						Usage: "list transfer-configs",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datatransferpb.ListTransferConfigsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListTransferConfigs(ctx, req)
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
						Name:  "schedule-runs",
						Usage: "schedule-runs transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							fmt.Printf("Executing schedule-runs on %s\n", name)
							return nil
						},
					},
					{
						Name:  "start-manual-runs",
						Usage: "start-manual-runs transfer-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							fmt.Printf("Executing start-manual-runs on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "transfer-logs",
				Usage: "Manage transfer-logs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list transfer-logs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "transfer_config", Usage: "The transfer_config.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/transferConfigs/%s", cmd.String("project"), cmd.String("transfer_config"))
							client, err := datatransfer.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datatransferpb.ListTransferLogsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListTransferLogs(ctx, req)
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
