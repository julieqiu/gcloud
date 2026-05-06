package translate

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud translate command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "translate",
		Usage: "manage Cloud Translation API resources",
		Commands: []*cli.Command{
			{
				Name:  "adaptive-mt-datasets",
				Usage: "Manage adaptive-mt-datasets resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list adaptive-mt-datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "import-adaptive-mt-file",
						Usage: "import-adaptive-mt-file adaptive-mt-datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing import-adaptive-mt-file on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "adaptive-mt-files",
				Usage: "Manage adaptive-mt-files resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe adaptive-mt-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
							&cli.StringFlag{Name: "file", Usage: "The file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s/adaptiveMtFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("file"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete adaptive-mt-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
							&cli.StringFlag{Name: "file", Usage: "The file.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s/adaptiveMtFiles/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"), cmd.String("file"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list adaptive-mt-files",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "adaptive-mt-sentences",
				Usage: "Manage adaptive-mt-sentences resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list adaptive-mt-sentences",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/adaptiveMtDatasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing list on %s\n", parent)
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
						Name:  "create",
						Usage: "create datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataset", Usage: "The dataset.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataset"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "import-data",
						Usage: "import-data datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing import-data...")
							return nil
						},
					},
					{
						Name:  "export-data",
						Usage: "export-data datasets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export-data...")
							return nil
						},
					},
				},
			},
			{
				Name:  "examples",
				Usage: "Manage examples resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list examples",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "glossaries",
				Usage: "Manage glossaries resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list glossaries",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete glossaries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "glossary-entries",
				Usage: "Manage glossary-entries resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
							&cli.StringFlag{Name: "glossary_entry", Usage: "The glossary_entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/glossaryEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"), cmd.String("glossary_entry"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
							&cli.StringFlag{Name: "glossary_entry", Usage: "The glossary_entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/glossaryEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"), cmd.String("glossary_entry"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete glossary-entries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "glossary", Usage: "The glossary.", Required: true},
							&cli.StringFlag{Name: "glossary_entry", Usage: "The glossary_entry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/glossaries/%s/glossaryEntries/%s", cmd.String("project"), cmd.String("location"), cmd.String("glossary"), cmd.String("glossary_entry"))
							fmt.Printf("Executing delete on %s\n", name)
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
						Name:  "translate-text",
						Usage: "translate-text locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing translate-text...")
							return nil
						},
					},
					{
						Name:  "romanize-text",
						Usage: "romanize-text locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing romanize-text...")
							return nil
						},
					},
					{
						Name:  "detect-language",
						Usage: "detect-language locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing detect-language...")
							return nil
						},
					},
					{
						Name:  "translate-document",
						Usage: "translate-document locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing translate-document...")
							return nil
						},
					},
					{
						Name:  "batch-translate-text",
						Usage: "batch-translate-text locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-translate-text...")
							return nil
						},
					},
					{
						Name:  "batch-translate-document",
						Usage: "batch-translate-document locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing batch-translate-document...")
							return nil
						},
					},
					{
						Name:  "adaptive-mt-translate",
						Usage: "adaptive-mt-translate locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing adaptive-mt-translate...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
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
						Name:  "create",
						Usage: "create models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list models",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "model", Usage: "The model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/models/%s", cmd.String("project"), cmd.String("location"), cmd.String("model"))
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
					{
						Name:  "wait",
						Usage: "wait operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "supported-languages",
				Usage: "Manage supported-languages resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe supported-languages",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
							return nil
						},
					},
				},
			},
		},
	}
}
