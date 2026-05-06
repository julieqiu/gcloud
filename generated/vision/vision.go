package vision

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/apiv1/visionpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
							client, err := vision.NewImageAnnotatorClient(ctx)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := vision.NewProductSearchClient(ctx)
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
							&cli.StringFlag{Name: "product-set-id", Usage: "The product set id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.CreateProductSetRequest{Parent: parent}
							req.ProductSetId = cmd.String("product-set-id")
							req.ProductSet = &visionpb.ProductSet{
								Name:        cmd.String("name"),
								DisplayName: cmd.String("display-name"),
							}
							resp, err := client.CreateProductSet(ctx, req)
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
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.GetProductSetRequest{Name: name}
							resp, err := client.GetProductSet(ctx, req)
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
						Usage: "update product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.UpdateProductSetRequest{}
							req.ProductSet = &visionpb.ProductSet{
								Name:        name,
								Name:        cmd.String("name"),
								DisplayName: cmd.String("display-name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateProductSet(ctx, req)
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
						Usage: "delete product-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product_set", Usage: "The product_set.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/productSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("product_set"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.DeleteProductSetRequest{Name: name}
							if err := client.DeleteProductSet(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
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
							&cli.StringFlag{Name: "product-id", Usage: "The product id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "product-category", Usage: "The product category.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.CreateProductRequest{Parent: parent}
							req.ProductId = cmd.String("product-id")
							req.Product = &visionpb.Product{
								Name:            cmd.String("name"),
								DisplayName:     cmd.String("display-name"),
								Description:     cmd.String("description"),
								ProductCategory: cmd.String("product-category"),
							}
							resp, err := client.CreateProduct(ctx, req)
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
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.GetProductRequest{Name: name}
							resp, err := client.GetProduct(ctx, req)
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
						Usage: "update products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "product-category", Usage: "The product category.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.UpdateProductRequest{}
							req.Product = &visionpb.Product{
								Name:            name,
								Name:            cmd.String("name"),
								DisplayName:     cmd.String("display-name"),
								Description:     cmd.String("description"),
								ProductCategory: cmd.String("product-category"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("product-category") {
								paths = append(paths, "product_category")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateProduct(ctx, req)
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
						Usage: "delete products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.DeleteProductRequest{Name: name}
							if err := client.DeleteProduct(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list products",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &visionpb.ListProductsInProductSetRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListProductsInProductSet(ctx, req)
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
							&cli.StringFlag{Name: "reference-image-id", Usage: "The reference image id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "uri", Usage: "The uri.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/products/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.CreateReferenceImageRequest{Parent: parent}
							req.ReferenceImageId = cmd.String("reference-image-id")
							req.ReferenceImage = &visionpb.ReferenceImage{
								Name: cmd.String("name"),
								Uri:  cmd.String("uri"),
							}
							resp, err := client.CreateReferenceImage(ctx, req)
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
						Usage: "delete reference-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
							&cli.StringFlag{Name: "reference_image", Usage: "The reference_image.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"), cmd.String("reference_image"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.DeleteReferenceImageRequest{Name: name}
							if err := client.DeleteReferenceImage(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list reference-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &visionpb.ListReferenceImagesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListReferenceImages(ctx, req)
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
						Usage: "describe reference-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "product", Usage: "The product.", Required: true},
							&cli.StringFlag{Name: "reference_image", Usage: "The reference_image.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", cmd.String("project"), cmd.String("location"), cmd.String("product"), cmd.String("reference_image"))
							client, err := vision.NewProductSearchClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &visionpb.GetReferenceImageRequest{Name: name}
							resp, err := client.GetReferenceImage(ctx, req)
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
		},
	}
}
