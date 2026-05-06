package oracledatabase

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	oracledatabase "cloud.google.com/go/oracledatabase/apiv1"
	"cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud oracledatabase command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListAutonomousDatabaseBackupsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAutonomousDatabaseBackups(ctx, req)
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
				Name:  "autonomous-database-character-sets",
				Usage: "Manage autonomous-database-character-sets resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list autonomous-database-character-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListAutonomousDatabaseCharacterSetsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAutonomousDatabaseCharacterSets(ctx, req)
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
				Name:  "autonomous-databases",
				Usage: "Manage autonomous-databases resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListAutonomousDatabasesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAutonomousDatabases(ctx, req)
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
						Usage: "describe autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetAutonomousDatabaseRequest{Name: name}
							resp, err := client.GetAutonomousDatabase(ctx, req)
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
						Usage: "create autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous-database-id", Usage: "The autonomous database id.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "admin-password", Usage: "The admin password.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.StringFlag{Name: "cidr", Usage: "The cidr.", Required: false},
							&cli.StringFlag{Name: "odb-network", Usage: "The odb network.", Required: false},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The odb subnet.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateAutonomousDatabaseRequest{Parent: parent}
							req.AutonomousDatabaseId = cmd.String("autonomous-database-id")
							req.AutonomousDatabase = &oracledatabasepb.AutonomousDatabase{
								Database:      cmd.String("database"),
								DisplayName:   cmd.String("display-name"),
								AdminPassword: cmd.String("admin-password"),
								Network:       cmd.String("network"),
								Cidr:          cmd.String("cidr"),
								OdbNetwork:    cmd.String("odb-network"),
								OdbSubnet:     cmd.String("odb-subnet"),
							}
							op, err := client.CreateAutonomousDatabase(ctx, req)
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
						Usage: "update autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "admin-password", Usage: "The admin password.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.StringFlag{Name: "cidr", Usage: "The cidr.", Required: false},
							&cli.StringFlag{Name: "odb-network", Usage: "The odb network.", Required: false},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The odb subnet.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.UpdateAutonomousDatabaseRequest{}
							req.AutonomousDatabase = &oracledatabasepb.AutonomousDatabase{
								Name:          name,
								Database:      cmd.String("database"),
								DisplayName:   cmd.String("display-name"),
								AdminPassword: cmd.String("admin-password"),
								Network:       cmd.String("network"),
								Cidr:          cmd.String("cidr"),
								OdbNetwork:    cmd.String("odb-network"),
								OdbSubnet:     cmd.String("odb-subnet"),
							}
							var paths []string
							if cmd.IsSet("database") {
								paths = append(paths, "database")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("admin-password") {
								paths = append(paths, "admin_password")
							}
							if cmd.IsSet("network") {
								paths = append(paths, "network")
							}
							if cmd.IsSet("cidr") {
								paths = append(paths, "cidr")
							}
							if cmd.IsSet("odb-network") {
								paths = append(paths, "odb_network")
							}
							if cmd.IsSet("odb-subnet") {
								paths = append(paths, "odb_subnet")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateAutonomousDatabase(ctx, req)
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
						Usage: "delete autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteAutonomousDatabaseRequest{Name: name}
							op, err := client.DeleteAutonomousDatabase(ctx, req)
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
					{
						Name:  "restore",
						Usage: "restore autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.RestoreAutonomousDatabaseRequest{Name: name}
							op, err := client.RestoreAutonomousDatabase(ctx, req)
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
						Name:  "generate-wallet",
						Usage: "generate-wallet autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							fmt.Printf("Executing generate-wallet on %s\n", name)
							return nil
						},
					},
					{
						Name:  "stop",
						Usage: "stop autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.StopAutonomousDatabaseRequest{Name: name}
							op, err := client.StopAutonomousDatabase(ctx, req)
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
						Name:  "start",
						Usage: "start autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.StartAutonomousDatabaseRequest{Name: name}
							op, err := client.StartAutonomousDatabase(ctx, req)
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
						Name:  "restart",
						Usage: "restart autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.RestartAutonomousDatabaseRequest{Name: name}
							op, err := client.RestartAutonomousDatabase(ctx, req)
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
						Name:  "switchover",
						Usage: "switchover autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
							&cli.StringFlag{Name: "peer-autonomous-database", Usage: "The peer autonomous database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.SwitchoverAutonomousDatabaseRequest{Name: name}
							req.PeerAutonomousDatabase = cmd.String("peer-autonomous-database")
							op, err := client.SwitchoverAutonomousDatabase(ctx, req)
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
						Name:  "failover",
						Usage: "failover autonomous-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "autonomous_database", Usage: "The autonomous_database.", Required: true},
							&cli.StringFlag{Name: "peer-autonomous-database", Usage: "The peer autonomous database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/autonomousDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("autonomous_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.FailoverAutonomousDatabaseRequest{Name: name}
							req.PeerAutonomousDatabase = cmd.String("peer-autonomous-database")
							op, err := client.FailoverAutonomousDatabase(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListAutonomousDbVersionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAutonomousDbVersions(ctx, req)
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
				Name:  "cloud-exadata-infrastructures",
				Usage: "Manage cloud-exadata-infrastructures resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListCloudExadataInfrastructuresRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListCloudExadataInfrastructures(ctx, req)
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
						Usage: "describe cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud_exadata_infrastructure", Usage: "The cloud_exadata_infrastructure.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud_exadata_infrastructure"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetCloudExadataInfrastructureRequest{Name: name}
							resp, err := client.GetCloudExadataInfrastructure(ctx, req)
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
						Usage: "create cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud-exadata-infrastructure-id", Usage: "The cloud exadata infrastructure id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "gcp-oracle-zone", Usage: "The gcp oracle zone.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateCloudExadataInfrastructureRequest{Parent: parent}
							req.CloudExadataInfrastructureId = cmd.String("cloud-exadata-infrastructure-id")
							req.CloudExadataInfrastructure = &oracledatabasepb.CloudExadataInfrastructure{
								DisplayName:   cmd.String("display-name"),
								GcpOracleZone: cmd.String("gcp-oracle-zone"),
							}
							op, err := client.CreateCloudExadataInfrastructure(ctx, req)
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
						Usage: "delete cloud-exadata-infrastructures",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud_exadata_infrastructure", Usage: "The cloud_exadata_infrastructure.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud_exadata_infrastructure"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteCloudExadataInfrastructureRequest{Name: name}
							op, err := client.DeleteCloudExadataInfrastructure(ctx, req)
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
				Name:  "cloud-vm-clusters",
				Usage: "Manage cloud-vm-clusters resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListCloudVmClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListCloudVmClusters(ctx, req)
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
						Usage: "describe cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud_vm_cluster", Usage: "The cloud_vm_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetCloudVmClusterRequest{Name: name}
							resp, err := client.GetCloudVmCluster(ctx, req)
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
						Usage: "create cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud-vm-cluster-id", Usage: "The cloud vm cluster id.", Required: true},
							&cli.StringFlag{Name: "exadata-infrastructure", Usage: "The exadata infrastructure.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "cidr", Usage: "The cidr.", Required: false},
							&cli.StringFlag{Name: "backup-subnet-cidr", Usage: "The backup subnet cidr.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
							&cli.StringFlag{Name: "odb-network", Usage: "The odb network.", Required: false},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The odb subnet.", Required: false},
							&cli.StringFlag{Name: "backup-odb-subnet", Usage: "The backup odb subnet.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateCloudVmClusterRequest{Parent: parent}
							req.CloudVmClusterId = cmd.String("cloud-vm-cluster-id")
							req.CloudVmCluster = &oracledatabasepb.CloudVmCluster{
								ExadataInfrastructure: cmd.String("exadata-infrastructure"),
								DisplayName:           cmd.String("display-name"),
								Cidr:                  cmd.String("cidr"),
								BackupSubnetCidr:      cmd.String("backup-subnet-cidr"),
								Network:               cmd.String("network"),
								OdbNetwork:            cmd.String("odb-network"),
								OdbSubnet:             cmd.String("odb-subnet"),
								BackupOdbSubnet:       cmd.String("backup-odb-subnet"),
							}
							op, err := client.CreateCloudVmCluster(ctx, req)
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
						Usage: "delete cloud-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud_vm_cluster", Usage: "The cloud_vm_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteCloudVmClusterRequest{Name: name}
							op, err := client.DeleteCloudVmCluster(ctx, req)
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
				Name:  "database-character-sets",
				Usage: "Manage database-character-sets resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list database-character-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDatabaseCharacterSetsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDatabaseCharacterSets(ctx, req)
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
				Name:  "databases",
				Usage: "Manage databases resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDatabasesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDatabases(ctx, req)
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
						Usage: "describe databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "database", Usage: "The database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/databases/%s", cmd.String("project"), cmd.String("location"), cmd.String("database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetDatabaseRequest{Name: name}
							resp, err := client.GetDatabase(ctx, req)
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
				Name:  "db-nodes",
				Usage: "Manage db-nodes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list db-nodes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud_vm_cluster", Usage: "The cloud_vm_cluster.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/cloudVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDbNodesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDbNodes(ctx, req)
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
				Name:  "db-servers",
				Usage: "Manage db-servers resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list db-servers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "cloud_exadata_infrastructure", Usage: "The cloud_exadata_infrastructure.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/cloudExadataInfrastructures/%s", cmd.String("project"), cmd.String("location"), cmd.String("cloud_exadata_infrastructure"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDbServersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDbServers(ctx, req)
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
				Name:  "db-system-initial-storage-sizes",
				Usage: "Manage db-system-initial-storage-sizes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list db-system-initial-storage-sizes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDbSystemInitialStorageSizesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDbSystemInitialStorageSizes(ctx, req)
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
				Name:  "db-system-shapes",
				Usage: "Manage db-system-shapes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list db-system-shapes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDbSystemShapesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDbSystemShapes(ctx, req)
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
				Name:  "db-systems",
				Usage: "Manage db-systems resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDbSystemsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDbSystems(ctx, req)
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
						Usage: "describe db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "db_system", Usage: "The db_system.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dbSystems/%s", cmd.String("project"), cmd.String("location"), cmd.String("db_system"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetDbSystemRequest{Name: name}
							resp, err := client.GetDbSystem(ctx, req)
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
						Usage: "create db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "db-system-id", Usage: "The db system id.", Required: true},
							&cli.StringFlag{Name: "gcp-oracle-zone", Usage: "The gcp oracle zone.", Required: false},
							&cli.StringFlag{Name: "odb-network", Usage: "The odb network.", Required: false},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The odb subnet.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateDbSystemRequest{Parent: parent}
							req.DbSystemId = cmd.String("db-system-id")
							req.DbSystem = &oracledatabasepb.DbSystem{
								GcpOracleZone: cmd.String("gcp-oracle-zone"),
								OdbNetwork:    cmd.String("odb-network"),
								OdbSubnet:     cmd.String("odb-subnet"),
								DisplayName:   cmd.String("display-name"),
							}
							op, err := client.CreateDbSystem(ctx, req)
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
						Usage: "delete db-systems",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "db_system", Usage: "The db_system.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dbSystems/%s", cmd.String("project"), cmd.String("location"), cmd.String("db_system"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteDbSystemRequest{Name: name}
							op, err := client.DeleteDbSystem(ctx, req)
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
				Name:  "db-versions",
				Usage: "Manage db-versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list db-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListDbVersionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListDbVersions(ctx, req)
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
				Name:  "entitlements",
				Usage: "Manage entitlements resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list entitlements",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListEntitlementsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListEntitlements(ctx, req)
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
				Name:  "exadb-vm-clusters",
				Usage: "Manage exadb-vm-clusters resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListExadbVmClustersRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListExadbVmClusters(ctx, req)
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
						Usage: "describe exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exadb_vm_cluster", Usage: "The exadb_vm_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetExadbVmClusterRequest{Name: name}
							resp, err := client.GetExadbVmCluster(ctx, req)
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
						Usage: "create exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exadb-vm-cluster-id", Usage: "The exadb vm cluster id.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The odb network.", Required: false},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The odb subnet.", Required: true},
							&cli.StringFlag{Name: "backup-odb-subnet", Usage: "The backup odb subnet.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateExadbVmClusterRequest{Parent: parent}
							req.ExadbVmClusterId = cmd.String("exadb-vm-cluster-id")
							req.ExadbVmCluster = &oracledatabasepb.ExadbVmCluster{
								OdbNetwork:      cmd.String("odb-network"),
								OdbSubnet:       cmd.String("odb-subnet"),
								BackupOdbSubnet: cmd.String("backup-odb-subnet"),
								DisplayName:     cmd.String("display-name"),
							}
							op, err := client.CreateExadbVmCluster(ctx, req)
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
						Usage: "delete exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exadb_vm_cluster", Usage: "The exadb_vm_cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteExadbVmClusterRequest{Name: name}
							op, err := client.DeleteExadbVmCluster(ctx, req)
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
					{
						Name:  "update",
						Usage: "update exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exadb_vm_cluster", Usage: "The exadb_vm_cluster.", Required: true},
							&cli.StringFlag{Name: "odb-network", Usage: "The odb network.", Required: false},
							&cli.StringFlag{Name: "odb-subnet", Usage: "The odb subnet.", Required: false},
							&cli.StringFlag{Name: "backup-odb-subnet", Usage: "The backup odb subnet.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.UpdateExadbVmClusterRequest{}
							req.ExadbVmCluster = &oracledatabasepb.ExadbVmCluster{
								Name:            name,
								OdbNetwork:      cmd.String("odb-network"),
								OdbSubnet:       cmd.String("odb-subnet"),
								BackupOdbSubnet: cmd.String("backup-odb-subnet"),
								DisplayName:     cmd.String("display-name"),
							}
							var paths []string
							if cmd.IsSet("odb-network") {
								paths = append(paths, "odb_network")
							}
							if cmd.IsSet("odb-subnet") {
								paths = append(paths, "odb_subnet")
							}
							if cmd.IsSet("backup-odb-subnet") {
								paths = append(paths, "backup_odb_subnet")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateExadbVmCluster(ctx, req)
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
						Name:  "remove-virtual-machine",
						Usage: "remove-virtual-machine exadb-vm-clusters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exadb_vm_cluster", Usage: "The exadb_vm_cluster.", Required: true},
							&cli.StringSliceFlag{Name: "hostnames", Usage: "The hostnames.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exadbVmClusters/%s", cmd.String("project"), cmd.String("location"), cmd.String("exadb_vm_cluster"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.RemoveVirtualMachineExadbVmClusterRequest{Name: name}
							req.Hostnames = cmd.StringSlice("hostnames")
							op, err := client.RemoveVirtualMachineExadbVmCluster(ctx, req)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListExascaleDbStorageVaultsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListExascaleDbStorageVaults(ctx, req)
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
						Usage: "describe exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exascale_db_storage_vault", Usage: "The exascale_db_storage_vault.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exascaleDbStorageVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("exascale_db_storage_vault"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetExascaleDbStorageVaultRequest{Name: name}
							resp, err := client.GetExascaleDbStorageVault(ctx, req)
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
						Usage: "create exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exascale-db-storage-vault-id", Usage: "The exascale db storage vault id.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "gcp-oracle-zone", Usage: "The gcp oracle zone.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateExascaleDbStorageVaultRequest{Parent: parent}
							req.ExascaleDbStorageVaultId = cmd.String("exascale-db-storage-vault-id")
							req.ExascaleDbStorageVault = &oracledatabasepb.ExascaleDbStorageVault{
								DisplayName:   cmd.String("display-name"),
								GcpOracleZone: cmd.String("gcp-oracle-zone"),
							}
							op, err := client.CreateExascaleDbStorageVault(ctx, req)
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
						Usage: "delete exascale-db-storage-vaults",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "exascale_db_storage_vault", Usage: "The exascale_db_storage_vault.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/exascaleDbStorageVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("exascale_db_storage_vault"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteExascaleDbStorageVaultRequest{Name: name}
							op, err := client.DeleteExascaleDbStorageVault(ctx, req)
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
				Name:  "gi-versions",
				Usage: "Manage gi-versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list gi-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListGiVersionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListGiVersions(ctx, req)
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
							client, err := oracledatabase.NewClient(ctx)
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
							client, err := oracledatabase.NewClient(ctx)
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
				Name:  "minor-versions",
				Usage: "Manage minor-versions resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list minor-versions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "gi_version", Usage: "The gi_version.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/giVersions/%s", cmd.String("project"), cmd.String("location"), cmd.String("gi_version"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListMinorVersionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListMinorVersions(ctx, req)
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
				Name:  "odb-networks",
				Usage: "Manage odb-networks resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListOdbNetworksRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListOdbNetworks(ctx, req)
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
						Usage: "describe odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb_network", Usage: "The odb_network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb_network"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetOdbNetworkRequest{Name: name}
							resp, err := client.GetOdbNetwork(ctx, req)
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
						Usage: "create odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb-network-id", Usage: "The odb network id.", Required: true},
							&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
							&cli.StringFlag{Name: "gcp-oracle-zone", Usage: "The gcp oracle zone.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateOdbNetworkRequest{Parent: parent}
							req.OdbNetworkId = cmd.String("odb-network-id")
							req.OdbNetwork = &oracledatabasepb.OdbNetwork{
								Network:       cmd.String("network"),
								GcpOracleZone: cmd.String("gcp-oracle-zone"),
							}
							op, err := client.CreateOdbNetwork(ctx, req)
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
						Usage: "delete odb-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb_network", Usage: "The odb_network.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb_network"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteOdbNetworkRequest{Name: name}
							op, err := client.DeleteOdbNetwork(ctx, req)
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
				Name:  "odb-subnets",
				Usage: "Manage odb-subnets resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb_network", Usage: "The odb_network.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb_network"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListOdbSubnetsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListOdbSubnets(ctx, req)
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
						Usage: "describe odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb_network", Usage: "The odb_network.", Required: true},
							&cli.StringFlag{Name: "odb_subnet", Usage: "The odb_subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s/odbSubnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb_network"), cmd.String("odb_subnet"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetOdbSubnetRequest{Name: name}
							resp, err := client.GetOdbSubnet(ctx, req)
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
						Usage: "create odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb_network", Usage: "The odb_network.", Required: true},
							&cli.StringFlag{Name: "odb-subnet-id", Usage: "The odb subnet id.", Required: true},
							&cli.StringFlag{Name: "cidr-range", Usage: "The cidr range.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb_network"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.CreateOdbSubnetRequest{Parent: parent}
							req.OdbSubnetId = cmd.String("odb-subnet-id")
							req.OdbSubnet = &oracledatabasepb.OdbSubnet{
								CidrRange: cmd.String("cidr-range"),
							}
							op, err := client.CreateOdbSubnet(ctx, req)
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
						Usage: "delete odb-subnets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "odb_network", Usage: "The odb_network.", Required: true},
							&cli.StringFlag{Name: "odb_subnet", Usage: "The odb_subnet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/odbNetworks/%s/odbSubnets/%s", cmd.String("project"), cmd.String("location"), cmd.String("odb_network"), cmd.String("odb_subnet"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.DeleteOdbSubnetRequest{Name: name}
							op, err := client.DeleteOdbSubnet(ctx, req)
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
							client, err := oracledatabase.NewClient(ctx)
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
							client, err := oracledatabase.NewClient(ctx)
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
							client, err := oracledatabase.NewClient(ctx)
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
							client, err := oracledatabase.NewClient(ctx)
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
				Name:  "pluggable-databases",
				Usage: "Manage pluggable-databases resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list pluggable-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &oracledatabasepb.ListPluggableDatabasesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPluggableDatabases(ctx, req)
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
						Usage: "describe pluggable-databases",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "pluggable_database", Usage: "The pluggable_database.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/pluggableDatabases/%s", cmd.String("project"), cmd.String("location"), cmd.String("pluggable_database"))
							client, err := oracledatabase.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &oracledatabasepb.GetPluggableDatabaseRequest{Name: name}
							resp, err := client.GetPluggableDatabase(ctx, req)
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
