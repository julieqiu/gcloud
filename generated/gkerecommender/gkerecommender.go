package gkerecommender

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud gkerecommender command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkerecommender",
		Usage: "manage GKE Recommender API resources",
		Commands: []*cli.Command{
			{
				Name:  "benchmarking-data",
				Usage: "Manage benchmarking-data resources",
				Commands: []*cli.Command{
					{
						Name:  "fetch",
						Usage: "fetch benchmarking-data",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch...")
							return nil
						},
					},
				},
			},
			{
				Name:  "model-server-versions",
				Usage: "Manage model-server-versions resources",
				Commands: []*cli.Command{
					{
						Name:  "fetch",
						Usage: "fetch model-server-versions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch...")
							return nil
						},
					},
				},
			},
			{
				Name:  "model-servers",
				Usage: "Manage model-servers resources",
				Commands: []*cli.Command{
					{
						Name:  "fetch",
						Usage: "fetch model-servers",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch...")
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
						Name:  "fetch",
						Usage: "fetch models",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch...")
							return nil
						},
					},
				},
			},
			{
				Name:  "optimized-manifest",
				Usage: "Manage optimized-manifest resources",
				Commands: []*cli.Command{
					{
						Name:  "generate",
						Usage: "generate optimized-manifest",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate...")
							return nil
						},
					},
				},
			},
			{
				Name:  "profiles",
				Usage: "Manage profiles resources",
				Commands: []*cli.Command{
					{
						Name:  "fetch",
						Usage: "fetch profiles",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing fetch...")
							return nil
						},
					},
				},
			},
		},
	}
}
