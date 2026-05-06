package memorystore

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	memorystore "cloud.google.com/go/memorystore/apiv1"
	"cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud memorystore command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "memorystore",
		Usage: "manage Memorystore API resources",
		Commands: []*cli.Command{
			{
				Name:  "backup-collections",
				Usage: "Manage backup-collections resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list backup-collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &memorystorepb.ListBackupCollectionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListBackupCollections(ctx, req)
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
						Usage: "describe backup-collections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_collection", Usage: "The backup_collection.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_collection"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.GetBackupCollectionRequest{Name: name}
							resp, err := client.GetBackupCollection(ctx, req)
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
				Name:  "backups",
				Usage: "Manage backups resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_collection", Usage: "The backup_collection.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_collection"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &memorystorepb.ListBackupsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListBackups(ctx, req)
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
						Usage: "describe backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_collection", Usage: "The backup_collection.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_collection"), cmd.String("backup"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.GetBackupRequest{Name: name}
							resp, err := client.GetBackup(ctx, req)
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
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_collection", Usage: "The backup_collection.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_collection"), cmd.String("backup"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.DeleteBackupRequest{Name: name}
							op, err := client.DeleteBackup(ctx, req)
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
						Name:  "export",
						Usage: "export backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_collection", Usage: "The backup_collection.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_collection"), cmd.String("backup"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.ExportBackupRequest{Name: name}
							op, err := client.ExportBackup(ctx, req)
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
				Name:  "certificate-authority",
				Usage: "Manage certificate-authority resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe certificate-authority",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.GetCertificateAuthorityRequest{Name: name}
							resp, err := client.GetCertificateAuthority(ctx, req)
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
				Name:  "instances",
				Usage: "Manage instances resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &memorystorepb.ListInstancesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListInstances(ctx, req)
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
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.GetInstanceRequest{Name: name}
							resp, err := client.GetInstance(ctx, req)
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
						Usage: "create instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance-id", Usage: "The instance id.", Required: true},
							&cli.IntFlag{Name: "replica-count", Usage: "The replica count.", Required: false},
							&cli.IntFlag{Name: "shard-count", Usage: "The shard count.", Required: false},
							&cli.StringFlag{Name: "engine-version", Usage: "The engine version.", Required: false},
							&cli.BoolFlag{Name: "deletion-protection-enabled", Usage: "The deletion protection enabled.", Required: false},
							&cli.BoolFlag{Name: "simulate-maintenance-event", Usage: "The simulate maintenance event.", Required: false},
							&cli.BoolFlag{Name: "ondemand-maintenance", Usage: "The ondemand maintenance.", Required: false},
							&cli.BoolFlag{Name: "async-instance-endpoints-deletion-enabled", Usage: "The async instance endpoints deletion enabled.", Required: false},
							&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: false},
							&cli.StringFlag{Name: "maintenance-version", Usage: "The maintenance version.", Required: false},
							&cli.BoolFlag{Name: "allow-fewer-zones-deployment", Usage: "The allow fewer zones deployment.", Required: false},
							&cli.StringFlag{Name: "server-ca-pool", Usage: "The server ca pool.", Required: false},
							&cli.BoolFlag{Name: "rotate-server-certificate", Usage: "The rotate server certificate.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.CreateInstanceRequest{Parent: parent}
							req.InstanceId = cmd.String("instance-id")
							req.Instance = &memorystorepb.Instance{
								ReplicaCount:                          int32(cmd.Int("replica-count")),
								ShardCount:                            int32(cmd.Int("shard-count")),
								EngineVersion:                         cmd.String("engine-version"),
								DeletionProtectionEnabled:             cmd.Bool("deletion-protection-enabled"),
								SimulateMaintenanceEvent:              cmd.Bool("simulate-maintenance-event"),
								OndemandMaintenance:                   cmd.Bool("ondemand-maintenance"),
								AsyncInstanceEndpointsDeletionEnabled: cmd.Bool("async-instance-endpoints-deletion-enabled"),
								KmsKey:                                cmd.String("kms-key"),
								MaintenanceVersion:                    cmd.String("maintenance-version"),
								AllowFewerZonesDeployment:             cmd.Bool("allow-fewer-zones-deployment"),
								ServerCaPool:                          cmd.String("server-ca-pool"),
								RotateServerCertificate:               cmd.Bool("rotate-server-certificate"),
							}
							op, err := client.CreateInstance(ctx, req)
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
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.IntFlag{Name: "replica-count", Usage: "The replica count.", Required: false},
							&cli.IntFlag{Name: "shard-count", Usage: "The shard count.", Required: false},
							&cli.StringFlag{Name: "engine-version", Usage: "The engine version.", Required: false},
							&cli.BoolFlag{Name: "deletion-protection-enabled", Usage: "The deletion protection enabled.", Required: false},
							&cli.BoolFlag{Name: "simulate-maintenance-event", Usage: "The simulate maintenance event.", Required: false},
							&cli.BoolFlag{Name: "ondemand-maintenance", Usage: "The ondemand maintenance.", Required: false},
							&cli.BoolFlag{Name: "async-instance-endpoints-deletion-enabled", Usage: "The async instance endpoints deletion enabled.", Required: false},
							&cli.StringFlag{Name: "kms-key", Usage: "The kms key.", Required: false},
							&cli.StringFlag{Name: "maintenance-version", Usage: "The maintenance version.", Required: false},
							&cli.BoolFlag{Name: "allow-fewer-zones-deployment", Usage: "The allow fewer zones deployment.", Required: false},
							&cli.StringFlag{Name: "server-ca-pool", Usage: "The server ca pool.", Required: false},
							&cli.BoolFlag{Name: "rotate-server-certificate", Usage: "The rotate server certificate.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.UpdateInstanceRequest{}
							req.Instance = &memorystorepb.Instance{
								Name:                                  name,
								ReplicaCount:                          int32(cmd.Int("replica-count")),
								ShardCount:                            int32(cmd.Int("shard-count")),
								EngineVersion:                         cmd.String("engine-version"),
								DeletionProtectionEnabled:             cmd.Bool("deletion-protection-enabled"),
								SimulateMaintenanceEvent:              cmd.Bool("simulate-maintenance-event"),
								OndemandMaintenance:                   cmd.Bool("ondemand-maintenance"),
								AsyncInstanceEndpointsDeletionEnabled: cmd.Bool("async-instance-endpoints-deletion-enabled"),
								KmsKey:                                cmd.String("kms-key"),
								MaintenanceVersion:                    cmd.String("maintenance-version"),
								AllowFewerZonesDeployment:             cmd.Bool("allow-fewer-zones-deployment"),
								ServerCaPool:                          cmd.String("server-ca-pool"),
								RotateServerCertificate:               cmd.Bool("rotate-server-certificate"),
							}
							var paths []string
							if cmd.IsSet("replica-count") {
								paths = append(paths, "replica_count")
							}
							if cmd.IsSet("shard-count") {
								paths = append(paths, "shard_count")
							}
							if cmd.IsSet("engine-version") {
								paths = append(paths, "engine_version")
							}
							if cmd.IsSet("deletion-protection-enabled") {
								paths = append(paths, "deletion_protection_enabled")
							}
							if cmd.IsSet("simulate-maintenance-event") {
								paths = append(paths, "simulate_maintenance_event")
							}
							if cmd.IsSet("ondemand-maintenance") {
								paths = append(paths, "ondemand_maintenance")
							}
							if cmd.IsSet("async-instance-endpoints-deletion-enabled") {
								paths = append(paths, "async_instance_endpoints_deletion_enabled")
							}
							if cmd.IsSet("kms-key") {
								paths = append(paths, "kms_key")
							}
							if cmd.IsSet("maintenance-version") {
								paths = append(paths, "maintenance_version")
							}
							if cmd.IsSet("allow-fewer-zones-deployment") {
								paths = append(paths, "allow_fewer_zones_deployment")
							}
							if cmd.IsSet("server-ca-pool") {
								paths = append(paths, "server_ca_pool")
							}
							if cmd.IsSet("rotate-server-certificate") {
								paths = append(paths, "rotate_server_certificate")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateInstance(ctx, req)
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
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.DeleteInstanceRequest{Name: name}
							op, err := client.DeleteInstance(ctx, req)
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
						Name:  "reschedule-maintenance",
						Usage: "reschedule-maintenance instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.RescheduleMaintenanceRequest{Name: name}
							op, err := client.RescheduleMaintenance(ctx, req)
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
						Name:  "backup",
						Usage: "backup instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "backup-id", Usage: "The backup id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.BackupInstanceRequest{Name: name}
							req.BackupId = cmd.String("backup-id")
							op, err := client.BackupInstance(ctx, req)
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
							client, err := memorystore.NewClient(ctx)
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
							client, err := memorystore.NewClient(ctx)
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
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := memorystore.NewClient(ctx)
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
							client, err := memorystore.NewClient(ctx)
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
							client, err := memorystore.NewClient(ctx)
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
							client, err := memorystore.NewClient(ctx)
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
				Name:  "shared-regional-certificate-authority",
				Usage: "Manage shared-regional-certificate-authority resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe shared-regional-certificate-authority",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/sharedRegionalCertificateAuthority", cmd.String("project"), cmd.String("location"))
							client, err := memorystore.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &memorystorepb.GetSharedRegionalCertificateAuthorityRequest{Name: name}
							resp, err := client.GetSharedRegionalCertificateAuthority(ctx, req)
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
