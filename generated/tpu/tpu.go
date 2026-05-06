package tpu

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud tpu command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "tpu",
		Usage: "manage Cloud TPU API resources",
		Commands: []*cli.Command{
			{
				Name:  "accelerator-types",
				Usage: "Manage accelerator-types resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list accelerator-types",
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
						Usage: "describe accelerator-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "accelerator_type", Usage: "The accelerator_type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/acceleratorTypes/%s", cmd.String("project"), cmd.String("location"), cmd.String("accelerator_type"))
							fmt.Printf("Executing describe on %s\n", name)
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
						Name:  "generate-service-identity",
						Usage: "generate-service-identity locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing generate-service-identity...")
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
				Name:  "nodes",
				Usage: "Manage nodes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list nodes",
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
						Usage: "describe nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create nodes",
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
						Usage: "delete nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "stop",
						Usage: "stop nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node"))
							fmt.Printf("Executing stop on %s\n", name)
							return nil
						},
					},
					{
						Name:  "start",
						Usage: "start nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node"))
							fmt.Printf("Executing start on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-guest-attributes",
						Usage: "get-guest-attributes nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "node", Usage: "The node.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/nodes/%s", cmd.String("project"), cmd.String("location"), cmd.String("node"))
							fmt.Printf("Executing get-guest-attributes on %s\n", name)
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
				},
			},
			{
				Name:  "queued-resources",
				Usage: "Manage queued-resources resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list queued-resources",
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
						Usage: "describe queued-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queued_resource", Usage: "The queued_resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queuedResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("queued_resource"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create queued-resources",
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
						Usage: "delete queued-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queued_resource", Usage: "The queued_resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queuedResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("queued_resource"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "reset",
						Usage: "reset queued-resources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "queued_resource", Usage: "The queued_resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/queuedResources/%s", cmd.String("project"), cmd.String("location"), cmd.String("queued_resource"))
							fmt.Printf("Executing reset on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "runtime-versions",
				Usage: "Manage runtime-versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list runtime-versions",
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
						Usage: "describe runtime-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "runtime_version", Usage: "The runtime_version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/runtimeVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("runtime_version"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
