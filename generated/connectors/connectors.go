package connectors

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud connectors command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "connectors",
		Usage: "manage Connectors API resources",
		Commands: []*cli.Command{
			{
				Name:  "connection-schema-metadata",
				Usage: "Manage connection-schema-metadata resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe connection-schema-metadata",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/connectionSchemaMetadata", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "refresh",
						Usage: "refresh connection-schema-metadata",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s/connectionSchemaMetadata", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing refresh on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "connections",
				Usage: "Manage connections resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list connections",
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
						Usage: "describe connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create connections",
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
						Usage: "update connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "connection", Usage: "The connection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/connections/%s", cmd.String("project"), cmd.String("location"), cmd.String("connection"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "connectors",
				Usage: "Manage connectors resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list connectors",
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
						Usage: "describe connectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "provider", Usage: "The provider.", Required: true},
							&cli.StringFlag{Name: "connector", Usage: "The connector.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/providers/%s/connectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("provider"), cmd.String("connector"))
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
				Name:  "providers",
				Usage: "Manage providers resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list providers",
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
						Usage: "describe providers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "provider", Usage: "The provider.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/providers/%s", cmd.String("project"), cmd.String("location"), cmd.String("provider"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "runtime-action-schemas",
				Usage: "Manage runtime-action-schemas resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list runtime-action-schemas",
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
				Name:  "runtime-config",
				Usage: "Manage runtime-config resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe runtime-config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/runtimeConfig", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "runtime-entity-schemas",
				Usage: "Manage runtime-entity-schemas resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list runtime-entity-schemas",
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
				Name:  "settings",
				Usage: "Manage settings resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe settings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/global/settings", cmd.String("project"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "versions",
				Usage: "Manage versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "provider", Usage: "The provider.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/providers/%s", cmd.String("project"), cmd.String("location"), cmd.String("provider"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "provider", Usage: "The provider.", Required: true},
							&cli.StringFlag{Name: "connector", Usage: "The connector.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/providers/%s/connectors/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("provider"), cmd.String("connector"), cmd.String("version"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
