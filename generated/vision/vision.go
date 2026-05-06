package vision

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud vision command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "vision",
		Usage: "manage Cloud Vision API resources",
		Commands: []*cli.Command{
			{
				Name:  "files",
				Usage: "Manage files resources",
				Commands: []*cli.Command{
					{
						Name:  "annotate",
						Usage: "annotate files",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing annotate...")
							return nil
						},
					},
					{
						Name:  "async-batch-annotate",
						Usage: "async-batch-annotate files",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing async-batch-annotate...")
							return nil
						},
					},
				},
			},
			{
				Name:  "images",
				Usage: "Manage images resources",
				Commands: []*cli.Command{
					{
						Name:  "annotate",
						Usage: "annotate images",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing annotate...")
							return nil
						},
					},
					{
						Name:  "async-batch-annotate",
						Usage: "async-batch-annotate images",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing async-batch-annotate...")
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "product-sets",
				Usage: "Manage product-sets resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create product-sets",
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
						Usage: "list product-sets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "add-product",
						Usage: "add-product product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							fmt.Printf("Executing add-product on %s\n", name)
							return nil
						},
					},
					{
						Name:  "remove-product",
						Usage: "remove-product product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							fmt.Printf("Executing remove-product on %s\n", name)
							return nil
						},
					},
					{
						Name:  "import",
						Usage: "import product-sets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing import...")
							return nil
						},
					},
				},
			},
			{
				Name:  "products",
				Usage: "Manage products resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create products",
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
						Usage: "list products",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list products",
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
						Name:  "purge",
						Usage: "purge products",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing purge...")
							return nil
						},
					},
				},
			},
			{
				Name:  "reference-images",
				Usage: "Manage reference-images resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create reference-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete reference-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
							&cli.StringFlag{Name: "reference_image", Usage: "The reference_image.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"), cmd.String("reference_image"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list reference-images",
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
						Usage: "describe reference-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
							&cli.StringFlag{Name: "reference_image", Usage: "The reference_image.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"), cmd.String("reference_image"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
