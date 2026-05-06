// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schemaregistry

import (
	schemaregistry "cloud.google.com/go/schemaregistry/apiv1"
	"cloud.google.com/go/schemaregistry/apiv1/schemaregistrypb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the managedkafka command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "managedkafka",
		Usage: "manage Managed Service for Apache Kafka API resources",
		Commands: []*cli.Command{
			{
				Name:  "compatibility",
				Usage: "Manage compatibility resources",
				Commands: []*cli.Command{

					{
						Name:  "check-compatibility",
						Usage: "check-compatibility compatibility",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "compatibility", Usage: "The ID of the compatibility.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The schema payload.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "schema-type", Usage: "The schema type of the schema.", Required: false},
							&cli.BoolFlag{Name: "verbose", Usage: "If true, the response will contain the compatibility check result.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/compatibility/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("compatibility"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.CheckCompatibilityRequest{
								Name:       name,
								SchemaType: runtime.Ptr(schemaregistrypb.Schema_SchemaType(schemaregistrypb.Schema_SchemaType_value[cmd.String("schema-type")])),
								Schema:     cmd.String("schema"),
								Verbose:    runtime.Ptr(cmd.Bool("verbose")),
							}

							resp, err := client.CheckCompatibility(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "config",
				Usage: "Manage config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "config", Usage: "The ID of the config.", Required: true},
							&cli.BoolFlag{Name: "default-to-global", Usage: "If true, the config will fall back to the config at the global.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/config/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetSchemaConfigRequest{
								Name:            name,
								DefaultToGlobal: runtime.Ptr(cmd.Bool("default-to-global")),
							}

							resp, err := client.GetSchemaConfig(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "compatibility", Usage: "The compatibility type of the schemas.", Required: true},
							&cli.StringFlag{Name: "config", Usage: "The ID of the config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.BoolFlag{Name: "normalize", Usage: "If true, the schema will be normalized before being stored or.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/config/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.UpdateSchemaConfigRequest{
								Name:          name,
								Compatibility: runtime.Ptr(schemaregistrypb.SchemaConfig_CompatibilityType(schemaregistrypb.SchemaConfig_CompatibilityType_value[cmd.String("compatibility")])),
								Normalize:     runtime.Ptr(cmd.Bool("normalize")),
							}

							resp, err := client.UpdateSchemaConfig(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "config", Usage: "The ID of the config.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/config/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("config"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.DeleteSchemaConfigRequest{
								Name: name,
							}

							resp, err := client.DeleteSchemaConfig(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "contexts",
				Usage: "Manage contexts resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "context", Usage: "The ID of the context.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("context"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetContextRequest{
								Name: name,
							}

							resp, err := client.GetContext(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list contexts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListContextsRequest{
								Parent: parent,
							}

							resp, err := client.ListContexts(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
				Name:  "mode",
				Usage: "Manage mode resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe mode",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mode", Usage: "The ID of the mode.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/mode/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("mode"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetSchemaModeRequest{
								Name: name,
							}

							resp, err := client.GetSchemaMode(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update mode",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mode", Usage: "The ID of the mode.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/mode/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("mode"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.UpdateSchemaModeRequest{
								Name: name,
							}

							resp, err := client.UpdateSchemaMode(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete mode",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "mode", Usage: "The ID of the mode.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/mode/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("mode"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.DeleteSchemaModeRequest{
								Name: name,
							}

							resp, err := client.DeleteSchemaMode(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
				Name:  "referencedby",
				Usage: "Manage referencedby resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list referencedby",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListReferencedSchemasRequest{
								Parent: parent,
							}

							resp, err := client.ListReferencedSchemas(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "schema",
				Usage: "Manage schema resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe schema",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "Used to limit the search for the schema ID to a specific subject,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetSchemaRequest{
								Name:    name,
								Subject: runtime.Ptr(cmd.String("subject")),
							}

							resp, err := client.GetRawSchema(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe schema",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deleted", Usage: "If true, no matter if the subject/version is soft-deleted or not,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetVersionRequest{
								Name:    name,
								Deleted: runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.GetRawSchemaVersion(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "schema-registries",
				Usage: "Manage schema-registries resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe schema-registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetSchemaRegistryRequest{
								Name: name,
							}

							resp, err := client.GetSchemaRegistry(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list schema-registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListSchemaRegistriesRequest{
								Parent: parent,
							}

							resp, err := client.ListSchemaRegistries(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create schema-registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registry-id", Usage: "The schema registry instance ID to use for this schema registry.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.CreateSchemaRegistryRequest{
								Parent:           parent,
								SchemaRegistryId: cmd.String("schema-registry-id"),
							}

							resp, err := client.CreateSchemaRegistry(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete schema-registries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteSchemaRegistry on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.DeleteSchemaRegistryRequest{
								Name: name,
							}

							if err := client.DeleteSchemaRegistry(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "schemas",
				Usage: "Manage schemas resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe schemas",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "Used to limit the search for the schema ID to a specific subject,.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetSchemaRequest{
								Name:    name,
								Subject: runtime.Ptr(cmd.String("subject")),
							}

							resp, err := client.GetSchema(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "subjects",
				Usage: "Manage subjects resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list subjects",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deleted", Usage: "If true, the response will include soft-deleted subjects.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject-prefix", Usage: "The context to filter the subjects by, in the format of.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListSubjectsRequest{
								Parent:        parent,
								SubjectPrefix: runtime.Ptr(cmd.String("subject-prefix")),
								Deleted:       runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.ListSubjects(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list subjects",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deleted", Usage: "If true, the response will include soft-deleted subjects.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The subject to filter the subjects by.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListSubjectsBySchemaIdRequest{
								Parent:  parent,
								Subject: runtime.Ptr(cmd.String("subject")),
								Deleted: runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.ListSubjectsBySchemaId(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete subjects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.BoolFlag{Name: "permanent", Usage: "If true, the subject and all associated metadata including the.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.DeleteSubjectRequest{
								Name:      name,
								Permanent: runtime.Ptr(cmd.Bool("permanent")),
							}

							resp, err := client.DeleteSubject(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "lookup-version",
						Usage: "lookup-version subjects",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deleted", Usage: "If true, soft-deleted versions will be included in lookup, no.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.BoolFlag{Name: "normalize", Usage: "If true, the schema will be normalized before being looked up.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The schema payload.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "schema-type", Usage: "The schema type of the schema.", Required: false},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.LookupVersionRequest{
								Parent:     parent,
								SchemaType: runtime.Ptr(schemaregistrypb.Schema_SchemaType(schemaregistrypb.Schema_SchemaType_value[cmd.String("schema-type")])),
								Schema:     cmd.String("schema"),
								Normalize:  runtime.Ptr(cmd.Bool("normalize")),
								Deleted:    runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.LookupVersion(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "types",
				Usage: "Manage types resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListSchemaTypesRequest{
								Parent: parent,
							}

							resp, err := client.ListSchemaTypes(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
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
							&cli.BoolFlag{Name: "deleted", Usage: "If true, the response will include soft-deleted versions of the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The ID of the schema.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The subject to filter the subjects by.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("schema"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListSchemaVersionsRequest{
								Parent:  parent,
								Subject: runtime.Ptr(cmd.String("subject")),
								Deleted: runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.ListSchemaVersions(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe versions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deleted", Usage: "If true, no matter if the subject/version is soft-deleted or not,.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.GetVersionRequest{
								Name:    name,
								Deleted: runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.GetVersion(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list versions",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deleted", Usage: "If true, the response will include soft-deleted versions of an.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.ListVersionsRequest{
								Parent:  parent,
								Deleted: runtime.Ptr(cmd.Bool("deleted")),
							}

							resp, err := client.ListVersions(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create versions",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "id", Usage: "The schema ID of the schema.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.BoolFlag{Name: "normalize", Usage: "If true, the schema will be normalized before being stored.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema", Usage: "The schema payload.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "schema-type", Usage: "The type of the schema.", Required: false},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
							&cli.IntFlag{Name: "version", Usage: "The version to create.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.CreateVersionRequest{
								Parent:     parent,
								Version:    runtime.Ptr(int32(cmd.Int("version"))),
								Id:         runtime.Ptr(int32(cmd.Int("id"))),
								SchemaType: runtime.Ptr(schemaregistrypb.Schema_SchemaType(schemaregistrypb.Schema_SchemaType_value[cmd.String("schema-type")])),
								Schema:     cmd.String("schema"),
								Normalize:  runtime.Ptr(cmd.Bool("normalize")),
							}

							resp, err := client.CreateVersion(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.BoolFlag{Name: "permanent", Usage: "If true, both the version and the referenced schema ID will be.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "schema-registrie", Usage: "The ID of the schema registrie.", Required: true},
							&cli.StringFlag{Name: "subject", Usage: "The ID of the subject.", Required: true},
							&cli.StringFlag{Name: "version", Usage: "The ID of the version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema-registrie"), cmd.String("subject"), cmd.String("version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := schemaregistry.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &schemaregistrypb.DeleteVersionRequest{
								Name:      name,
								Permanent: runtime.Ptr(cmd.Bool("permanent")),
							}

							resp, err := client.DeleteVersion(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
		},
	}
}
