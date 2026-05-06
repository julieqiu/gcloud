package timeseriesinsights

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
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
							fmt.Printf("Executing delete on %s\n", name)
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
