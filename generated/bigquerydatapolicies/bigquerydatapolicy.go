package bigquerydatapolicy

import (
	"context"
	"fmt"
	"strings"

	datapolicies "cloud.google.com/go/bigquery/datapolicies/apiv2"
	"cloud.google.com/go/bigquery/datapolicies/apiv2/datapoliciespb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud bigquerydatapolicy command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigquerydatapolicy",
		Usage: "manage BigQuery Data Policy API resources",
		Commands: []*cli.Command{
			{
				Name:  "data-policies",
				Usage: "Manage data-policies resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data-policy-id", Usage: "The data policy id.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datapoliciespb.CreateDataPolicyRequest{Parent: parent}
							req.DataPolicyId = cmd.String("data-policy-id")
							req.DataPolicy = &datapoliciespb.DataPolicy{
								Etag: cmd.String("etag"),
							}
							resp, err := client.CreateDataPolicy(ctx, req)
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
						Name:  "add-grantees",
						Usage: "add-grantees data-policies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing add-grantees...")
							return nil
						},
					},
					{
						Name:  "remove-grantees",
						Usage: "remove-grantees data-policies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing remove-grantees...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_policy", Usage: "The data_policy.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_policy"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datapoliciespb.UpdateDataPolicyRequest{}
							req.DataPolicy = &datapoliciespb.DataPolicy{
								Name: name,
								Etag: cmd.String("etag"),
							}
							var paths []string
							if cmd.IsSet("etag") {
								paths = append(paths, "etag")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateDataPolicy(ctx, req)
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
						Usage: "delete data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_policy", Usage: "The data_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_policy"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datapoliciespb.DeleteDataPolicyRequest{Name: name}
							if err := client.DeleteDataPolicy(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "data_policy", Usage: "The data_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("data_policy"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &datapoliciespb.GetDataPolicyRequest{Name: name}
							resp, err := client.GetDataPolicy(ctx, req)
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
						Usage: "list data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &datapoliciespb.ListDataPoliciesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDataPolicies(ctx, req)
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
						Name:  "get-iam-policy",
						Usage: "get-iam-policy data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataPolicie", Usage: "The dataPolicie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataPolicie"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataPolicie", Usage: "The dataPolicie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataPolicie"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions data-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dataPolicie", Usage: "The dataPolicie.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dataPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("dataPolicie"))
							client, err := datapolicies.NewDataPolicyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
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
