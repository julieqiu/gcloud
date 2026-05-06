package assuredworkloads

import (
	"context"
	"fmt"
	"strings"

	assuredworkloads "cloud.google.com/go/assuredworkloads/apiv1"
	"cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud assuredworkloads command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "assuredworkloads",
		Usage: "manage Assured Workloads API resources",
		Commands: []*cli.Command{
			{
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							client, err := assuredworkloads.NewClient(ctx)
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
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/operations/%s", cmd.String("organization"), cmd.String("location"), cmd.String("operation"))
							client, err := assuredworkloads.NewClient(ctx)
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
				Name:  "workloads",
				Usage: "Manage workloads resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "billing-account", Usage: "The billing account.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "provisioned-resources-parent", Usage: "The provisioned resources parent.", Required: false},
							&cli.BoolFlag{Name: "enable-sovereign-controls", Usage: "The enable sovereign controls.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							client, err := assuredworkloads.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &assuredworkloadspb.CreateWorkloadRequest{Parent: parent}
							req.ExternalId = cmd.String("external-id")
							req.Workload = &assuredworkloadspb.Workload{
								Name:                       cmd.String("name"),
								DisplayName:                cmd.String("display-name"),
								BillingAccount:             cmd.String("billing-account"),
								Etag:                       cmd.String("etag"),
								ProvisionedResourcesParent: cmd.String("provisioned-resources-parent"),
								EnableSovereignControls:    cmd.Bool("enable-sovereign-controls"),
							}
							op, err := client.CreateWorkload(ctx, req)
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
						Usage: "update workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "billing-account", Usage: "The billing account.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
							&cli.StringFlag{Name: "provisioned-resources-parent", Usage: "The provisioned resources parent.", Required: false},
							&cli.BoolFlag{Name: "enable-sovereign-controls", Usage: "The enable sovereign controls.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							client, err := assuredworkloads.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &assuredworkloadspb.UpdateWorkloadRequest{}
							req.Workload = &assuredworkloadspb.Workload{
								Name:                       name,
								Name:                       cmd.String("name"),
								DisplayName:                cmd.String("display-name"),
								BillingAccount:             cmd.String("billing-account"),
								Etag:                       cmd.String("etag"),
								ProvisionedResourcesParent: cmd.String("provisioned-resources-parent"),
								EnableSovereignControls:    cmd.Bool("enable-sovereign-controls"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("billing-account") {
								paths = append(paths, "billing_account")
							}
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							if cmd.IsSet("provisioned-resources-parent") {
								paths = append(paths, "provisioned_resources_parent")
							}
							if cmd.IsSet("enable-sovereign-controls") {
								paths = append(paths, "enable_sovereign_controls")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateWorkload(ctx, req)
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
						Name:  "restrict-allowed-resources",
						Usage: "restrict-allowed-resources workloads",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing restrict-allowed-resources...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							client, err := assuredworkloads.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &assuredworkloadspb.DeleteWorkloadRequest{Name: name}
							if err := client.DeleteWorkload(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "workload", Usage: "The workload.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/workloads/%s", cmd.String("organization"), cmd.String("location"), cmd.String("workload"))
							client, err := assuredworkloads.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &assuredworkloadspb.GetWorkloadRequest{Name: name}
							resp, err := client.GetWorkload(ctx, req)
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
						Usage: "list workloads",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							client, err := assuredworkloads.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &assuredworkloadspb.ListWorkloadsRequest{Parent: parent}
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
