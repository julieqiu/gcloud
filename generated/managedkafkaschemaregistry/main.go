package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	schemaregistry "cloud.google.com/go/managedkafka/schemaregistry/apiv1"
	"cloud.google.com/go/managedkafka/schemaregistry/apiv1/schemaregistrypb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "managedkafka",
				Usage: "manage Managed Service for Apache Kafka API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "compatibility",
						Usage: "Manage compatibility resources",
						Commands: []*cli.Command{
							{
								Name:  "check-compatibility",
								Usage: "check-compatibility compatibility",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing check-compatibility...")
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetSchemaConfigRequest{Name: name}
									resp, err := client.GetSchemaConfig(ctx, req)
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
								Usage: "update config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.UpdateSchemaConfigRequest{}
									resp, err := client.UpdateSchemaConfig(ctx, req)
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
								Usage: "delete config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/config", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.DeleteSchemaConfigRequest{Name: name}
									if err := client.DeleteSchemaConfig(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "context", Usage: "The context.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/contexts/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("context"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetContextRequest{Name: name}
									resp, err := client.GetContext(ctx, req)
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
								Usage: "list contexts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &schemaregistrypb.ListContextsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListContexts(ctx, req)
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
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
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
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
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
						Name:  "mode",
						Usage: "Manage mode resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe mode",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/mode", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetSchemaModeRequest{Name: name}
									resp, err := client.GetSchemaMode(ctx, req)
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
								Usage: "update mode",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/mode", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.UpdateSchemaModeRequest{}
									resp, err := client.UpdateSchemaMode(ctx, req)
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
								Usage: "delete mode",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/mode", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.DeleteSchemaModeRequest{Name: name}
									if err := client.DeleteSchemaMode(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
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
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
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
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
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
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
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
						Name:  "referencedby",
						Usage: "Manage referencedby resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list referencedby",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "subject", Usage: "The subject.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("subject"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &schemaregistrypb.ListReferencedSchemasRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListReferencedSchemas(ctx, req)
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
					{
						Name:  "schema",
						Usage: "Manage schema resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe schema",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas/ids/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("schema"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetSchemaRequest{Name: name}
									resp, err := client.GetRawSchema(ctx, req)
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
								Usage: "describe schema",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "subject", Usage: "The subject.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("subject"), cmd.String("version"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetVersionRequest{Name: name}
									resp, err := client.GetRawSchemaVersion(ctx, req)
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
						Name:  "schema-registries",
						Usage: "Manage schema-registries resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe schema-registries",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetSchemaRegistryRequest{Name: name}
									resp, err := client.GetSchemaRegistry(ctx, req)
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
								Usage: "list schema-registries",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create schema-registries",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema-registry-id", Usage: "The schema registry id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.CreateSchemaRegistryRequest{Parent: parent}
									req.SchemaRegistryId = cmd.String("schema-registry-id")
									resp, err := client.CreateSchemaRegistry(ctx, req)
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
								Usage: "delete schema-registries",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.DeleteSchemaRegistryRequest{Name: name}
									if err := client.DeleteSchemaRegistry(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas/ids/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("schema"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetSchemaRequest{Name: name}
									resp, err := client.GetSchema(ctx, req)
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
						Name:  "subjects",
						Usage: "Manage subjects resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list subjects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list subjects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete subjects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "subject", Usage: "The subject.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("subject"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.DeleteSubjectRequest{Name: name}
									if err := client.DeleteSubject(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
							{
								Name:  "lookup-version",
								Usage: "lookup-version subjects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "subject", Usage: "The subject.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("subject"))
									fmt.Printf("Executing lookup-version on %s\n", name)
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/schemas", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &schemaregistrypb.ListSchemaVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSchemaVersions(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "subject", Usage: "The subject.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("subject"), cmd.String("version"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.GetVersionRequest{Name: name}
									resp, err := client.GetVersion(ctx, req)
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
								Usage: "list versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &schemaregistrypb.ListVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVersions(ctx, req)
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
								Name:  "create",
								Usage: "create versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.IntFlag{Name: "version", Usage: "The version.", Required: false},
									&cli.IntFlag{Name: "id", Usage: "The id.", Required: false},
									&cli.StringFlag{Name: "schema", Usage: "The schema.", Required: true},
									&cli.BoolFlag{Name: "normalize", Usage: "The normalize.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.CreateVersionRequest{Parent: parent}
									req.Version = int32(cmd.Int("version"))
									req.Id = int32(cmd.Int("id"))
									req.Schema = cmd.String("schema")
									req.Normalize = cmd.Bool("normalize")
									resp, err := client.CreateVersion(ctx, req)
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
								Usage: "delete versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "schema_registry", Usage: "The schema_registry.", Required: true},
									&cli.StringFlag{Name: "subject", Usage: "The subject.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/schemaRegistries/%s/subjects/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("schema_registry"), cmd.String("subject"), cmd.String("version"))
									client, err := schemaregistry.NewManagedSchemaRegistryClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &schemaregistrypb.DeleteVersionRequest{Name: name}
									if err := client.DeleteVersion(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
