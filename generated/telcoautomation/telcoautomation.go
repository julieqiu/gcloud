package telcoautomation

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	telcoautomation "cloud.google.com/go/telcoautomation/apiv1"
	"cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud telcoautomation command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "telcoautomation",
		Usage: "manage Telco Automation API resources",
		Commands: []*cli.Command{
			{
				Name:  "blueprints",
				Usage: "Manage blueprints resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint-id", Usage: "The blueprint id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "source-blueprint", Usage: "The source blueprint.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.CreateBlueprintRequest{Parent: parent}
							req.BlueprintId = cmd.String("blueprint-id")
							req.Blueprint = &telcoautomationpb.Blueprint{
								Name:            cmd.String("name"),
								SourceBlueprint: cmd.String("source-blueprint"),
								DisplayName:     cmd.String("display-name"),
							}
							resp, err := client.CreateBlueprint(ctx, req)
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
						Usage: "update blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "source-blueprint", Usage: "The source blueprint.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.UpdateBlueprintRequest{}
							req.Blueprint = &telcoautomationpb.Blueprint{
								Name:            name,
								Name:            cmd.String("name"),
								SourceBlueprint: cmd.String("source-blueprint"),
								DisplayName:     cmd.String("display-name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("source-blueprint") {
								paths = append(paths, "source_blueprint")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateBlueprint(ctx, req)
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
						Usage: "describe blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.GetBlueprintRequest{Name: name}
							resp, err := client.GetBlueprint(ctx, req)
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
						Usage: "delete blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.DeleteBlueprintRequest{Name: name}
							if err := client.DeleteBlueprint(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListBlueprintsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListBlueprints(ctx, req)
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
						Name:  "approve",
						Usage: "approve blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							fmt.Printf("Executing approve on %s\n", name)
							return nil
						},
					},
					{
						Name:  "propose",
						Usage: "propose blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							fmt.Printf("Executing propose on %s\n", name)
							return nil
						},
					},
					{
						Name:  "reject",
						Usage: "reject blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							fmt.Printf("Executing reject on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListBlueprintRevisionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListBlueprintRevisions(ctx, req)
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
						Name:  "search-revisions",
						Usage: "search-revisions blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							fmt.Printf("Executing search-revisions on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "discard",
						Usage: "discard blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "blueprint", Usage: "The blueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/blueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("blueprint"))
							fmt.Printf("Executing discard on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "deployments",
				Usage: "Manage deployments resources",
				Commands: []*cli.Command{
					{
						Name:  "search-revisions",
						Usage: "search-revisions deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							fmt.Printf("Executing search-revisions on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment-id", Usage: "The deployment id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "source-blueprint-revision", Usage: "The source blueprint revision.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "workload-cluster", Usage: "The workload cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.CreateDeploymentRequest{Parent: parent}
							req.DeploymentId = cmd.String("deployment-id")
							req.Deployment = &telcoautomationpb.Deployment{
								Name:                    cmd.String("name"),
								SourceBlueprintRevision: cmd.String("source-blueprint-revision"),
								DisplayName:             cmd.String("display-name"),
								WorkloadCluster:         cmd.String("workload-cluster"),
							}
							resp, err := client.CreateDeployment(ctx, req)
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
						Usage: "update deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "source-blueprint-revision", Usage: "The source blueprint revision.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "workload-cluster", Usage: "The workload cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.UpdateDeploymentRequest{}
							req.Deployment = &telcoautomationpb.Deployment{
								Name:                    name,
								Name:                    cmd.String("name"),
								SourceBlueprintRevision: cmd.String("source-blueprint-revision"),
								DisplayName:             cmd.String("display-name"),
								WorkloadCluster:         cmd.String("workload-cluster"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("source-blueprint-revision") {
								paths = append(paths, "source_blueprint_revision")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("workload-cluster") {
								paths = append(paths, "workload_cluster")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateDeployment(ctx, req)
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
						Usage: "describe deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.GetDeploymentRequest{Name: name}
							resp, err := client.GetDeployment(ctx, req)
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
						Name:  "remove",
						Usage: "remove deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							fmt.Printf("Executing remove on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListDeploymentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDeployments(ctx, req)
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
						Name:  "list",
						Usage: "list deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListDeploymentRevisionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDeploymentRevisions(ctx, req)
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
						Name:  "discard",
						Usage: "discard deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							fmt.Printf("Executing discard on %s\n", name)
							return nil
						},
					},
					{
						Name:  "apply",
						Usage: "apply deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							fmt.Printf("Executing apply on %s\n", name)
							return nil
						},
					},
					{
						Name:  "compute-deployment-status",
						Usage: "compute-deployment-status deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							fmt.Printf("Executing compute-deployment-status on %s\n", name)
							return nil
						},
					},
					{
						Name:  "rollback",
						Usage: "rollback deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							fmt.Printf("Executing rollback on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "edge-slms",
				Usage: "Manage edge-slms resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListEdgeSlmsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListEdgeSlms(ctx, req)
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
						Usage: "describe edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "edge_slm", Usage: "The edge_slm.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/edgeSlms/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge_slm"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.GetEdgeSlmRequest{Name: name}
							resp, err := client.GetEdgeSlm(ctx, req)
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
						Usage: "create edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "edge-slm-id", Usage: "The edge slm id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "orchestration-cluster", Usage: "The orchestration cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.CreateEdgeSlmRequest{Parent: parent}
							req.EdgeSlmId = cmd.String("edge-slm-id")
							req.EdgeSlm = &telcoautomationpb.EdgeSlm{
								Name:                 cmd.String("name"),
								OrchestrationCluster: cmd.String("orchestration-cluster"),
							}
							op, err := client.CreateEdgeSlm(ctx, req)
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
						Usage: "delete edge-slms",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "edge_slm", Usage: "The edge_slm.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/edgeSlms/%s", cmd.String("project"), cmd.String("location"), cmd.String("edge_slm"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.DeleteEdgeSlmRequest{Name: name}
							op, err := client.DeleteEdgeSlm(ctx, req)
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
				Name:  "hydrated-deployments",
				Usage: "Manage hydrated-deployments resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
							&cli.StringFlag{Name: "hydrated_deployment", Usage: "The hydrated_deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s/hydratedDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"), cmd.String("hydrated_deployment"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.GetHydratedDeploymentRequest{Name: name}
							resp, err := client.GetHydratedDeployment(ctx, req)
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
						Usage: "list hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListHydratedDeploymentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListHydratedDeployments(ctx, req)
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
						Name:  "update",
						Usage: "update hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
							&cli.StringFlag{Name: "hydrated_deployment", Usage: "The hydrated_deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s/hydratedDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"), cmd.String("hydrated_deployment"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.UpdateHydratedDeploymentRequest{}
							req.HydratedDeployment = &telcoautomationpb.HydratedDeployment{
								Name: name,
							}
							resp, err := client.UpdateHydratedDeployment(ctx, req)
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
						Name:  "apply",
						Usage: "apply hydrated-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
							&cli.StringFlag{Name: "deployment", Usage: "The deployment.", Required: true},
							&cli.StringFlag{Name: "hydrated_deployment", Usage: "The hydrated_deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s/deployments/%s/hydratedDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"), cmd.String("deployment"), cmd.String("hydrated_deployment"))
							fmt.Printf("Executing apply on %s\n", name)
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
							parent := fmt.Sprintf("projects/%s/locations", cmd.String("project"))
							client, err := telcoautomation.NewClient(ctx)
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
							client, err := telcoautomation.NewClient(ctx)
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
							parent := fmt.Sprintf("projects/%s/locations/%s/operations", cmd.String("project"), cmd.String("location"))
							client, err := telcoautomation.NewClient(ctx)
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
							client, err := telcoautomation.NewClient(ctx)
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
							client, err := telcoautomation.NewClient(ctx)
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
							client, err := telcoautomation.NewClient(ctx)
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
				Name:  "orchestration-clusters",
				Usage: "Manage orchestration-clusters resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListOrchestrationClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListOrchestrationClusters(ctx, req)
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
						Usage: "describe orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.GetOrchestrationClusterRequest{Name: name}
							resp, err := client.GetOrchestrationCluster(ctx, req)
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
						Usage: "create orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration-cluster-id", Usage: "The orchestration cluster id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.CreateOrchestrationClusterRequest{Parent: parent}
							req.OrchestrationClusterId = cmd.String("orchestration-cluster-id")
							req.OrchestrationCluster = &telcoautomationpb.OrchestrationCluster{
								Name: cmd.String("name"),
							}
							op, err := client.CreateOrchestrationCluster(ctx, req)
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
						Usage: "delete orchestration-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "orchestration_cluster", Usage: "The orchestration_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/orchestrationClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("orchestration_cluster"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.DeleteOrchestrationClusterRequest{Name: name}
							op, err := client.DeleteOrchestrationCluster(ctx, req)
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
				Name:  "public-blueprints",
				Usage: "Manage public-blueprints resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list public-blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &telcoautomationpb.ListPublicBlueprintsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPublicBlueprints(ctx, req)
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
						Usage: "describe public-blueprints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "public_lueprint", Usage: "The public_lueprint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/publicBlueprints/%s", cmd.String("project"), cmd.String("location"), cmd.String("public_lueprint"))
							client, err := telcoautomation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &telcoautomationpb.GetPublicBlueprintRequest{Name: name}
							resp, err := client.GetPublicBlueprint(ctx, req)
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
