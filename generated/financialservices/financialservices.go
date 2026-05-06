package financialservices

import (
	"context"
	"fmt"
	"strings"

	financialservices "cloud.google.com/go/financialservices/apiv1"
	"cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud financialservices command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "financialservices",
		Usage: "manage Financial Services API resources",
		Commands: []*cli.Command{
			{
				Name:  "backtest-results",
				Usage: "Manage backtest-results resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list backtest-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &financialservicespb.ListBacktestResultsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListBacktestResults(ctx, req)
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
						Usage: "describe backtest-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "backtest_result", Usage: "The backtest_result.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/backtestResults/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("backtest_result"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetBacktestResultRequest{Name: name}
							resp, err := client.GetBacktestResult(ctx, req)
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
						Usage: "create backtest-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "backtest-result-id", Usage: "The backtest result id.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
							&cli.IntFlag{Name: "backtest-periods", Usage: "The backtest periods.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.CreateBacktestResultRequest{Parent: parent}
							req.BacktestResultId = cmd.String("backtest-result-id")
							req.BacktestResult = &financialservicespb.BacktestResult{
								Dataset:         cmd.String("dataset"),
								Model:           cmd.String("model"),
								BacktestPeriods: int32(cmd.Int("backtest-periods")),
							}
							op, err := client.CreateBacktestResult(ctx, req)
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
						Name:  "update",
						Usage: "update backtest-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "backtest_result", Usage: "The backtest_result.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: false},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: false},
							&cli.IntFlag{Name: "backtest-periods", Usage: "The backtest periods.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/backtestResults/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("backtest_result"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.UpdateBacktestResultRequest{}
							req.BacktestResult = &financialservicespb.BacktestResult{
								Name:            name,
								Dataset:         cmd.String("dataset"),
								Model:           cmd.String("model"),
								BacktestPeriods: int32(cmd.Int("backtest-periods")),
							}
							var paths []string
							if cmd.IsSet("dataset") {
								paths = append(paths, "dataset")
							}
							if cmd.IsSet("model") {
								paths = append(paths, "model")
							}
							if cmd.IsSet("backtest-periods") {
								paths = append(paths, "backtest_periods")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateBacktestResult(ctx, req)
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
						Name:  "export-metadata",
						Usage: "export-metadata backtest-results",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export-metadata...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete backtest-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "backtest_result", Usage: "The backtest_result.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/backtestResults/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("backtest_result"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.DeleteBacktestResultRequest{Name: name}
							op, err := client.DeleteBacktestResult(ctx, req)
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
				},
			},
			{
				Name:  "datasets",
				Usage: "Manage datasets resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &financialservicespb.ListDatasetsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDatasets(ctx, req)
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
						Usage: "describe datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/datasets/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("dataset"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetDatasetRequest{Name: name}
							resp, err := client.GetDataset(ctx, req)
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
						Usage: "create datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "dataset-id", Usage: "The dataset id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.CreateDatasetRequest{Parent: parent}
							req.DatasetId = cmd.String("dataset-id")
							op, err := client.CreateDataset(ctx, req)
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
						Name:  "update",
						Usage: "update datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/datasets/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("dataset"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.UpdateDatasetRequest{}
							req.Dataset = &financialservicespb.Dataset{
								Name: name,
							}
							op, err := client.UpdateDataset(ctx, req)
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
						Usage: "delete datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/datasets/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("dataset"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.DeleteDatasetRequest{Name: name}
							op, err := client.DeleteDataset(ctx, req)
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
				},
			},
			{
				Name:  "engine-configs",
				Usage: "Manage engine-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list engine-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &financialservicespb.ListEngineConfigsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListEngineConfigs(ctx, req)
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
						Usage: "describe engine-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "engine_config", Usage: "The engine_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/engineConfigs/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("engine_config"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetEngineConfigRequest{Name: name}
							resp, err := client.GetEngineConfig(ctx, req)
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
						Usage: "create engine-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "engine-config-id", Usage: "The engine config id.", Required: true},
							&cli.StringFlag{Name: "engine-version", Usage: "The engine version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.CreateEngineConfigRequest{Parent: parent}
							req.EngineConfigId = cmd.String("engine-config-id")
							req.EngineConfig = &financialservicespb.EngineConfig{
								EngineVersion: cmd.String("engine-version"),
							}
							op, err := client.CreateEngineConfig(ctx, req)
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
						Name:  "update",
						Usage: "update engine-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "engine_config", Usage: "The engine_config.", Required: true},
							&cli.StringFlag{Name: "engine-version", Usage: "The engine version.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/engineConfigs/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("engine_config"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.UpdateEngineConfigRequest{}
							req.EngineConfig = &financialservicespb.EngineConfig{
								Name:          name,
								EngineVersion: cmd.String("engine-version"),
							}
							var paths []string
							if cmd.IsSet("engine-version") {
								paths = append(paths, "engine_version")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateEngineConfig(ctx, req)
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
						Name:  "export-metadata",
						Usage: "export-metadata engine-configs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export-metadata...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete engine-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "engine_config", Usage: "The engine_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/engineConfigs/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("engine_config"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.DeleteEngineConfigRequest{Name: name}
							op, err := client.DeleteEngineConfig(ctx, req)
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
				},
			},
			{
				Name:  "engine-versions",
				Usage: "Manage engine-versions resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe engine-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "engine_version", Usage: "The engine_version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/engineVersions/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("engine_version"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetEngineVersionRequest{Name: name}
							resp, err := client.GetEngineVersion(ctx, req)
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
						Usage: "list engine-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &financialservicespb.ListEngineVersionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListEngineVersions(ctx, req)
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
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetInstanceRequest{Name: name}
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
							&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.CreateInstanceRequest{Parent: parent}
							req.InstanceId = cmd.String("instance-id")
							req.Instance = &financialservicespb.Instance{
								KmsKey: cmd.String("kms-key"),
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
						Name:  "update",
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.UpdateInstanceRequest{}
							req.Instance = &financialservicespb.Instance{
								Name:   name,
								KmsKey: cmd.String("kms-key"),
							}
							var paths []string
							if cmd.IsSet("kms-key") {
								paths = append(paths, "kms_key")
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
						Name:  "delete",
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.DeleteInstanceRequest{Name: name}
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
						Name:  "import-registered-parties",
						Usage: "import-registered-parties instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringSliceFlag{Name: "party-tables", Usage: "The party tables.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.ImportRegisteredPartiesRequest{Name: name}
							req.PartyTables = cmd.StringSlice("party-tables")
							req.ValidateOnly = cmd.Bool("validate-only")
							op, err := client.ImportRegisteredParties(ctx, req)
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
						Name:  "export-registered-parties",
						Usage: "export-registered-parties instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.ExportRegisteredPartiesRequest{Name: name}
							op, err := client.ExportRegisteredParties(ctx, req)
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
							client, err := financialservices.NewAMLClient(ctx)
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
							client, err := financialservices.NewAMLClient(ctx)
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
				Name:  "models",
				Usage: "Manage models resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &financialservicespb.ListModelsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListModels(ctx, req)
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
						Usage: "describe models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/models/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("model"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetModelRequest{Name: name}
							resp, err := client.GetModel(ctx, req)
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
						Usage: "create models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "model-id", Usage: "The model id.", Required: true},
							&cli.StringFlag{Name: "engine-config", Usage: "The engine config.", Required: true},
							&cli.StringFlag{Name: "primary-dataset", Usage: "The primary dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.CreateModelRequest{Parent: parent}
							req.ModelId = cmd.String("model-id")
							req.Model = &financialservicespb.Model{
								EngineConfig:   cmd.String("engine-config"),
								PrimaryDataset: cmd.String("primary-dataset"),
							}
							op, err := client.CreateModel(ctx, req)
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
						Name:  "update",
						Usage: "update models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
							&cli.StringFlag{Name: "engine-config", Usage: "The engine config.", Required: false},
							&cli.StringFlag{Name: "primary-dataset", Usage: "The primary dataset.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/models/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("model"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.UpdateModelRequest{}
							req.Model = &financialservicespb.Model{
								Name:           name,
								EngineConfig:   cmd.String("engine-config"),
								PrimaryDataset: cmd.String("primary-dataset"),
							}
							var paths []string
							if cmd.IsSet("engine-config") {
								paths = append(paths, "engine_config")
							}
							if cmd.IsSet("primary-dataset") {
								paths = append(paths, "primary_dataset")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateModel(ctx, req)
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
						Name:  "export-metadata",
						Usage: "export-metadata models",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export-metadata...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/models/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("model"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.DeleteModelRequest{Name: name}
							op, err := client.DeleteModel(ctx, req)
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
							client, err := financialservices.NewAMLClient(ctx)
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
							client, err := financialservices.NewAMLClient(ctx)
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
							client, err := financialservices.NewAMLClient(ctx)
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
							client, err := financialservices.NewAMLClient(ctx)
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
				Name:  "prediction-results",
				Usage: "Manage prediction-results resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list prediction-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &financialservicespb.ListPredictionResultsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPredictionResults(ctx, req)
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
						Usage: "describe prediction-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "prediction_result", Usage: "The prediction_result.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/predictionResults/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("prediction_result"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.GetPredictionResultRequest{Name: name}
							resp, err := client.GetPredictionResult(ctx, req)
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
						Usage: "create prediction-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "prediction-result-id", Usage: "The prediction result id.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
							&cli.IntFlag{Name: "prediction-periods", Usage: "The prediction periods.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.CreatePredictionResultRequest{Parent: parent}
							req.PredictionResultId = cmd.String("prediction-result-id")
							req.PredictionResult = &financialservicespb.PredictionResult{
								Dataset:           cmd.String("dataset"),
								Model:             cmd.String("model"),
								PredictionPeriods: int32(cmd.Int("prediction-periods")),
							}
							op, err := client.CreatePredictionResult(ctx, req)
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
						Name:  "update",
						Usage: "update prediction-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "prediction_result", Usage: "The prediction_result.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: false},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: false},
							&cli.IntFlag{Name: "prediction-periods", Usage: "The prediction periods.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/predictionResults/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("prediction_result"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.UpdatePredictionResultRequest{}
							req.PredictionResult = &financialservicespb.PredictionResult{
								Name:              name,
								Dataset:           cmd.String("dataset"),
								Model:             cmd.String("model"),
								PredictionPeriods: int32(cmd.Int("prediction-periods")),
							}
							var paths []string
							if cmd.IsSet("dataset") {
								paths = append(paths, "dataset")
							}
							if cmd.IsSet("model") {
								paths = append(paths, "model")
							}
							if cmd.IsSet("prediction-periods") {
								paths = append(paths, "prediction_periods")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdatePredictionResult(ctx, req)
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
						Name:  "export-metadata",
						Usage: "export-metadata prediction-results",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export-metadata...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete prediction-results",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_num", Usage: "The project_num.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "prediction_result", Usage: "The prediction_result.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s/predictionResults/%s", cmd.String("project_num"), cmd.String("location"), cmd.String("instance"), cmd.String("prediction_result"))
							client, err := financialservices.NewAMLClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &financialservicespb.DeletePredictionResultRequest{Name: name}
							op, err := client.DeletePredictionResult(ctx, req)
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
				},
			},
		},
	}
}
