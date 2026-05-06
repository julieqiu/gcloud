package language

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud language command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "language",
		Usage: "manage Cloud Natural Language API resources",
		Commands: []*cli.Command{
			{
				Name:  "documents",
				Usage: "Manage documents resources",
				Commands: []*cli.Command{
					{
						Name:  "analyze-sentiment",
						Usage: "analyze-sentiment documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing analyze-sentiment...")
							return nil
						},
					},
					{
						Name:  "analyze-entities",
						Usage: "analyze-entities documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing analyze-entities...")
							return nil
						},
					},
					{
						Name:  "classify-text",
						Usage: "classify-text documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing classify-text...")
							return nil
						},
					},
					{
						Name:  "moderate-text",
						Usage: "moderate-text documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing moderate-text...")
							return nil
						},
					},
					{
						Name:  "annotate-text",
						Usage: "annotate-text documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing annotate-text...")
							return nil
						},
					},
				},
			},
		},
	}
}
