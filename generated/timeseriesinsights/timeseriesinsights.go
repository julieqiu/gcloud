package timeseriesinsights

import (
	"context"
	"fmt"

	timeseriesinsights "cloud.google.com/go/timeseriesinsights/apiv1"
	"cloud.google.com/go/timeseriesinsights/apiv1/timeseriesinsightspb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
)

// Command returns the gcloud timeseriesinsights command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "timeseriesinsights",
		Usage: "manage Timeseries Insights API resources",
		Commands: []*cli.Command{
			{
				Name:  "datasets",
				Usage: "Manage datasets resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := timeseriesinsights.NewTimeseriesInsightsControllerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &timeseriesinsightspb.CreateDataSetRequest{Parent: parent}
							req.Dataset = &timeseriesinsightspb.DataSet{
								Name: cmd.String("name"),
							}
							resp, err := client.CreateDataSet(ctx, req)
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
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/datasets/%s", cmd.String("project"), cmd.String("dataset"))
							client, err := timeseriesinsights.NewTimeseriesInsightsControllerClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &timeseriesinsightspb.DeleteDataSetRequest{Name: name}
							if err := client.DeleteDataSet(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "append-events",
						Usage: "append-events datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing append-events...")
							return nil
						},
					},
					{
						Name:  "query",
						Usage: "query datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/datasets/%s", cmd.String("project"), cmd.String("dataset"))
							fmt.Printf("Executing query on %s\n", name)
							return nil
						},
					},
					{
						Name:  "evaluate-slice",
						Usage: "evaluate-slice datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing evaluate-slice...")
							return nil
						},
					},
					{
						Name:  "evaluate-timeseries",
						Usage: "evaluate-timeseries datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing evaluate-timeseries...")
							return nil
						},
					},
				},
			},
		},
	}
}
