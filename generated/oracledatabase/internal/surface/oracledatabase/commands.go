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

package oracledatabase

import (
	oracledatabase "cloud.google.com/go/oracledatabase/apiv1"
	"cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the oracledatabase command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "oracledatabase",
		Usage: "manage Oracle Database@Google Cloud API resources",
		Commands: []*cli.Command{
			{
				Name:  "autonomous-database-backups",
				Usage: "Manage autonomous-database-backups resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list autonomous-database-backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListAutonomousDatabaseBackupsRequest{
								Parent:    parent,
								Filter:    cmd.String("filter"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutonomousDatabaseBackups(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "autonomous-database-character-sets",
				Usage: "Manage autonomous-database-character-sets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list autonomous-database-character-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListAutonomousDatabaseCharacterSetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutonomousDatabaseCharacterSets(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "autonomous-databases",
				Usage: "Manage autonomous-databases resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListAutonomousDatabasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutonomousDatabases(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetAutonomousDatabaseRequest{
								Name: name,
							}

							resp, err := client.GetAutonomousDatabase(ctx, req)
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
						Usage: "create autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database-id", Usage: "The ID of the Autonomous Database to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateAutonomousDatabaseRequest{
								Parent:               parent,
								AutonomousDatabaseId: cmd.String("autonomous-database-id"),
								RequestId:            cmd.String("request-id"),
							}

							op, err := client.CreateAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "update autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "autonomous_database.name" not yet supported.
							autonomous_database_name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							fmt.Printf("Executing update on %s\n", autonomous_database_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteAutonomousDatabase %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteAutonomousDatabaseRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "restore",
						Usage: "restore autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.RestoreAutonomousDatabaseRequest{
								Name: name,
							}

							op, err := client.RestoreAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "generate-wallet",
						Usage: "generate-wallet autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.BoolFlag{Name: "is-regional", Usage: "True when requesting regional connection strings in PDB connect.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "password", Usage: "The password used to encrypt the keys inside the wallet.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "type", Usage: "The type of wallet generation for the Autonomous Database.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GenerateAutonomousDatabaseWalletRequest{
								Name:       name,
								Type:       oracledatabasepb.GenerateType(oracledatabasepb.GenerateType_value[cmd.String("type")]),
								IsRegional: cmd.Bool("is-regional"),
								Password:   cmd.String("password"),
							}

							resp, err := client.GenerateAutonomousDatabaseWallet(ctx, req)
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
						Name:  "stop",
						Usage: "stop autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.StopAutonomousDatabaseRequest{
								Name: name,
							}

							op, err := client.StopAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "start",
						Usage: "start autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.StartAutonomousDatabaseRequest{
								Name: name,
							}

							op, err := client.StartAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "restart",
						Usage: "restart autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.RestartAutonomousDatabaseRequest{
								Name: name,
							}

							op, err := client.RestartAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "switchover",
						Usage: "switchover autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "peer-autonomous-database", Usage: "The peer database name to switch over to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.SwitchoverAutonomousDatabaseRequest{
								Name:                   name,
								PeerAutonomousDatabase: cmd.String("peer-autonomous-database"),
							}

							op, err := client.SwitchoverAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Name:  "failover",
						Usage: "failover autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autonomous-database", Usage: "The ID of the autonomous database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "peer-autonomous-database", Usage: "The peer database name to fail over to.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.FailoverAutonomousDatabaseRequest{
								Name:                   name,
								PeerAutonomousDatabase: cmd.String("peer-autonomous-database"),
							}

							op, err := client.FailoverAutonomousDatabase(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "autonomous-db-versions",
				Usage: "Manage autonomous-db-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list autonomous-db-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListAutonomousDbVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAutonomousDbVersions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "cloud-exadata-infrastructures",
				Usage: "Manage cloud-exadata-infrastructures resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListCloudExadataInfrastructuresRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListCloudExadataInfrastructures(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-exadata-infrastructure", Usage: "The ID of the cloud exadata infrastructure.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud-exadata-infrastructure"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetCloudExadataInfrastructureRequest{
								Name: name,
							}

							resp, err := client.GetCloudExadataInfrastructure(ctx, req)
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
						Usage: "create cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-exadata-infrastructure-id", Usage: "The ID of the Exadata Infrastructure to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateCloudExadataInfrastructureRequest{
								Parent:                       parent,
								CloudExadataInfrastructureId: cmd.String("cloud-exadata-infrastructure-id"),
								RequestId:                    cmd.String("request-id"),
							}

							op, err := client.CreateCloudExadataInfrastructure(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-exadata-infrastructure", Usage: "The ID of the cloud exadata infrastructure.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all VM clusters for this Exadata Infrastructure.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud-exadata-infrastructure"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCloudExadataInfrastructure %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteCloudExadataInfrastructureRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteCloudExadataInfrastructure(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "cloud-vm-clusters",
				Usage: "Manage cloud-vm-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The number of VM clusters to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the page of results the server returns.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListCloudVmClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListCloudVmClusters(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-vm-cluster", Usage: "The ID of the cloud vm cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud-vm-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetCloudVmClusterRequest{
								Name: name,
							}

							resp, err := client.GetCloudVmCluster(ctx, req)
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
						Usage: "create cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-vm-cluster-id", Usage: "The ID of the VM Cluster to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateCloudVmClusterRequest{
								Parent:           parent,
								CloudVmClusterId: cmd.String("cloud-vm-cluster-id"),
								RequestId:        cmd.String("request-id"),
							}

							op, err := client.CreateCloudVmCluster(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-vm-cluster", Usage: "The ID of the cloud vm cluster.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, all child resources for the VM Cluster will be.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud-vm-cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteCloudVmCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteCloudVmClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Force:     cmd.Bool("force"),
							}

							op, err := client.DeleteCloudVmCluster(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "database-character-sets",
				Usage: "Manage database-character-sets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list database-character-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of DatabaseCharacterSets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDatabaseCharacterSetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDatabaseCharacterSets(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "databases",
				Usage: "Manage databases resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the requested page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDatabasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDatabases(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "database", Usage: "The ID of the database.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/databases/%s", cmd.String("project"), cmd.String("location"), cmd.String("database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetDatabaseRequest{
								Name: name,
							}

							resp, err := client.GetDatabase(ctx, req)
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
				Name:  "db-nodes",
				Usage: "Manage db-nodes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list db-nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-vm-cluster", Usage: "The ID of the cloud vm cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the node should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud-vm-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDbNodesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDbNodes(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "db-servers",
				Usage: "Manage db-servers resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list db-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cloud-exadata-infrastructure", Usage: "The ID of the cloud exadata infrastructure.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud-exadata-infrastructure"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDbServersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDbServers(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "db-system-initial-storage-sizes",
				Usage: "Manage db-system-initial-storage-sizes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list db-system-initial-storage-sizes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the requested page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDbSystemInitialStorageSizesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListDbSystemInitialStorageSizes(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "db-system-shapes",
				Usage: "Manage db-system-shapes resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list db-system-shapes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDbSystemShapesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDbSystemShapes(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "db-systems",
				Usage: "Manage db-systems resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDbSystemsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListDbSystems(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "db-system", Usage: "The ID of the db system.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dbSystems/%s", cmd.String("project"), cmd.String("location"), cmd.String("db-system"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetDbSystemRequest{
								Name: name,
							}

							resp, err := client.GetDbSystem(ctx, req)
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
						Usage: "create db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "db-system-id", Usage: "The ID of the DbSystem to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateDbSystemRequest{
								Parent:     parent,
								DbSystemId: cmd.String("db-system-id"),
								RequestId:  cmd.String("request-id"),
							}

							op, err := client.CreateDbSystem(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "db-system", Usage: "The ID of the db system.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dbSystems/%s", cmd.String("project"), cmd.String("location"), cmd.String("db-system"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteDbSystem %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteDbSystemRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteDbSystem(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "db-versions",
				Usage: "Manage db-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list db-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "Filter expression that matches a subset of the DbVersions to.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the requested page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListDbVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListDbVersions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "entitlements",
				Usage: "Manage entitlements resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListEntitlementsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListEntitlements(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "exadb-vm-clusters",
				Usage: "Manage exadb-vm-clusters resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListExadbVmClustersRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListExadbVmClusters(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exadb-vm-cluster", Usage: "The ID of the exadb vm cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb-vm-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetExadbVmClusterRequest{
								Name: name,
							}

							resp, err := client.GetExadbVmCluster(ctx, req)
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
						Usage: "create exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exadb-vm-cluster-id", Usage: "The ID of the ExadbVmCluster to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateExadbVmClusterRequest{
								Parent:           parent,
								ExadbVmClusterId: cmd.String("exadb-vm-cluster-id"),
								RequestId:        cmd.String("request-id"),
							}

							op, err := client.CreateExadbVmCluster(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exadb-vm-cluster", Usage: "The ID of the exadb vm cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb-vm-cluster"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteExadbVmCluster %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteExadbVmClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteExadbVmCluster(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exadb-vm-cluster", Usage: "The ID of the exadb vm cluster.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "exadb_vm_cluster.name" not yet supported.
							exadb_vm_cluster_name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb-vm-cluster"))
							fmt.Printf("Executing update on %s\n", exadb_vm_cluster_name)
							return nil
						},
					},

					{
						Name:  "remove-virtual-machine",
						Usage: "remove-virtual-machine exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exadb-vm-cluster", Usage: "The ID of the exadb vm cluster.", Required: true},
							&cli.StringSliceFlag{Name: "hostnames", Usage: "The list of host names of db nodes to be removed from the.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb-vm-cluster"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.RemoveVirtualMachineExadbVmClusterRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
								Hostnames: cmd.StringSlice("hostnames"),
							}

							op, err := client.RemoveVirtualMachineExadbVmCluster(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
				Name:  "exascale-db-storage-vaults",
				Usage: "Manage exascale-db-storage-vaults resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListExascaleDbStorageVaultsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListExascaleDbStorageVaults(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exascale-db-storage-vault", Usage: "The ID of the exascale db storage vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exascaleDbStorageVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("exascale-db-storage-vault"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetExascaleDbStorageVaultRequest{
								Name: name,
							}

							resp, err := client.GetExascaleDbStorageVault(ctx, req)
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
						Usage: "create exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exascale-db-storage-vault-id", Usage: "The ID of the ExascaleDbStorageVault to create.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateExascaleDbStorageVaultRequest{
								Parent:                   parent,
								ExascaleDbStorageVaultId: cmd.String("exascale-db-storage-vault-id"),
								RequestId:                cmd.String("request-id"),
							}

							op, err := client.CreateExascaleDbStorageVault(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "exascale-db-storage-vault", Usage: "The ID of the exascale db storage vault.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exascaleDbStorageVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("exascale-db-storage-vault"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteExascaleDbStorageVault %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteExascaleDbStorageVaultRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteExascaleDbStorageVault(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "gi-versions",
				Usage: "Manage gi-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list gi-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListGiVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListGiVersions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
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
				Name:  "minor-versions",
				Usage: "Manage minor-versions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list minor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "gi-version", Usage: "The ID of the gi version.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying the requested page of results to return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/giVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("gi-version"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListMinorVersionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMinorVersions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "odb-networks",
				Usage: "Manage odb-networks resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListOdbNetworksRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListOdbNetworks(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The ID of the odb network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb-network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetOdbNetworkRequest{
								Name: name,
							}

							resp, err := client.GetOdbNetwork(ctx, req)
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
						Usage: "create odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network-id", Usage: "The ID of the OdbNetwork to create.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateOdbNetworkRequest{
								Parent:       parent,
								OdbNetworkId: cmd.String("odb-network-id"),
								RequestId:    cmd.String("request-id"),
							}

							op, err := client.CreateOdbNetwork(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The ID of the odb network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb-network"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteOdbNetwork %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteOdbNetworkRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteOdbNetwork(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
			{
				Name:  "odb-subnets",
				Usage: "Manage odb-subnets resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The ID of the odb network.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "An expression for ordering the results of the request.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A token identifying a page of results the server should return.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb-network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListOdbSubnetsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
								OrderBy:   cmd.String("order-by"),
							}

							limit := cmd.Int("limit")
							it := client.ListOdbSubnets(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The ID of the odb network.", Required: true},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The ID of the odb subnet.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s/odbSubnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb-network"), cmd.String("odb-subnet"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetOdbSubnetRequest{
								Name: name,
							}

							resp, err := client.GetOdbSubnet(ctx, req)
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
						Usage: "create odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The ID of the odb network.", Required: true},
							&cli.StringFlag{Name: "odb-subnet-id", Usage: "The ID of the OdbSubnet to create.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb-network"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.CreateOdbSubnetRequest{
								Parent:      parent,
								OdbSubnetId: cmd.String("odb-subnet-id"),
								RequestId:   cmd.String("request-id"),
							}

							op, err := client.CreateOdbSubnet(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
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
						Usage: "delete odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The ID of the odb network.", Required: true},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The ID of the odb subnet.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional ID to identify the request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s/odbSubnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb-network"), cmd.String("odb-subnet"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("DeleteOdbSubnet %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.DeleteOdbSubnetRequest{
								Name:      name,
								RequestId: cmd.String("request-id"),
							}

							op, err := client.DeleteOdbSubnet(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Println("Done.")
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
				Name:  "pluggable-databases",
				Usage: "Manage pluggable-databases resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list pluggable-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of PluggableDatabases to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListPluggableDatabases`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.ListPluggableDatabasesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListPluggableDatabases(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe pluggable-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "pluggable-database", Usage: "The ID of the pluggable database.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pluggableDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("pluggable-database"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := oracledatabase.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &oracledatabasepb.GetPluggableDatabaseRequest{
								Name: name,
							}

							resp, err := client.GetPluggableDatabase(ctx, req)
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
