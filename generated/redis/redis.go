package redis

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	redis "cloud.google.com/go/redis/apiv1"
	"cloud.google.com/go/redis/apiv1/redispb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud redis command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "redis",
		Usage: "manage Google Cloud Memorystore for Redis API resources",
		Commands: []*cli.Command{
			{
				Name:  "auth-string",
				Usage: "Manage auth-string resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe auth-string",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.GetInstanceAuthStringRequest{Name: name}
							resp, err := client.GetInstanceAuthString(ctx, req)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
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
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.GetInstanceRequest{Name: name}
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
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "The location id.", Required: false},
							&cli.StringFlag{Name: "alternative-location-id", Usage: "The alternative location id.", Required: false},
							&cli.StringFlag{Name: "redis-version", Usage: "The redis version.", Required: false},
							&cli.StringFlag{Name: "reserved-ip-range", Usage: "The reserved ip range.", Required: false},
							&cli.StringFlag{Name: "secondary-ip-range", Usage: "The secondary ip range.", Required: false},
							&cli.IntFlag{Name: "memory-size-gb", Usage: "The memory size gb.", Required: true},
							&cli.StringFlag{Name: "authorized-network", Usage: "The authorized network.", Required: false},
							&cli.BoolFlag{Name: "auth-enabled", Usage: "The auth enabled.", Required: false},
							&cli.IntFlag{Name: "replica-count", Usage: "The replica count.", Required: false},
							&cli.StringFlag{Name: "customer-managed-key", Usage: "The customer managed key.", Required: false},
							&cli.StringFlag{Name: "maintenance-version", Usage: "The maintenance version.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.CreateInstanceRequest{Parent: parent}
							req.InstanceId = cmd.String("instance-id")
							req.Instance = &redispb.Instance{
								Name:                  cmd.String("name"),
								DisplayName:           cmd.String("display-name"),
								LocationId:            cmd.String("location-id"),
								AlternativeLocationId: cmd.String("alternative-location-id"),
								RedisVersion:          cmd.String("redis-version"),
								ReservedIpRange:       cmd.String("reserved-ip-range"),
								SecondaryIpRange:      cmd.String("secondary-ip-range"),
								MemorySizeGb:          int32(cmd.Int("memory-size-gb")),
								AuthorizedNetwork:     cmd.String("authorized-network"),
								AuthEnabled:           cmd.Bool("auth-enabled"),
								ReplicaCount:          int32(cmd.Int("replica-count")),
								CustomerManagedKey:    cmd.String("customer-managed-key"),
								MaintenanceVersion:    cmd.String("maintenance-version"),
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
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "location-id", Usage: "The location id.", Required: false},
							&cli.StringFlag{Name: "alternative-location-id", Usage: "The alternative location id.", Required: false},
							&cli.StringFlag{Name: "redis-version", Usage: "The redis version.", Required: false},
							&cli.StringFlag{Name: "reserved-ip-range", Usage: "The reserved ip range.", Required: false},
							&cli.StringFlag{Name: "secondary-ip-range", Usage: "The secondary ip range.", Required: false},
							&cli.IntFlag{Name: "memory-size-gb", Usage: "The memory size gb.", Required: false},
							&cli.StringFlag{Name: "authorized-network", Usage: "The authorized network.", Required: false},
							&cli.BoolFlag{Name: "auth-enabled", Usage: "The auth enabled.", Required: false},
							&cli.IntFlag{Name: "replica-count", Usage: "The replica count.", Required: false},
							&cli.StringFlag{Name: "customer-managed-key", Usage: "The customer managed key.", Required: false},
							&cli.StringFlag{Name: "maintenance-version", Usage: "The maintenance version.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.UpdateInstanceRequest{}
							req.Instance = &redispb.Instance{
								Name:                  name,
								Name:                  cmd.String("name"),
								DisplayName:           cmd.String("display-name"),
								LocationId:            cmd.String("location-id"),
								AlternativeLocationId: cmd.String("alternative-location-id"),
								RedisVersion:          cmd.String("redis-version"),
								ReservedIpRange:       cmd.String("reserved-ip-range"),
								SecondaryIpRange:      cmd.String("secondary-ip-range"),
								MemorySizeGb:          int32(cmd.Int("memory-size-gb")),
								AuthorizedNetwork:     cmd.String("authorized-network"),
								AuthEnabled:           cmd.Bool("auth-enabled"),
								ReplicaCount:          int32(cmd.Int("replica-count")),
								CustomerManagedKey:    cmd.String("customer-managed-key"),
								MaintenanceVersion:    cmd.String("maintenance-version"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("location-id") {
								paths = append(paths, "location_id")
							}
							if cmd.IsSet("alternative-location-id") {
								paths = append(paths, "alternative_location_id")
							}
							if cmd.IsSet("redis-version") {
								paths = append(paths, "redis_version")
							}
							if cmd.IsSet("reserved-ip-range") {
								paths = append(paths, "reserved_ip_range")
							}
							if cmd.IsSet("secondary-ip-range") {
								paths = append(paths, "secondary_ip_range")
							}
							if cmd.IsSet("memory-size-gb") {
								paths = append(paths, "memory_size_gb")
							}
							if cmd.IsSet("authorized-network") {
								paths = append(paths, "authorized_network")
							}
							if cmd.IsSet("auth-enabled") {
								paths = append(paths, "auth_enabled")
							}
							if cmd.IsSet("replica-count") {
								paths = append(paths, "replica_count")
							}
							if cmd.IsSet("customer-managed-key") {
								paths = append(paths, "customer_managed_key")
							}
							if cmd.IsSet("maintenance-version") {
								paths = append(paths, "maintenance_version")
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
						Name:  "upgrade",
						Usage: "upgrade instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
							&cli.StringFlag{Name: "redis-version", Usage: "The redis version.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.UpgradeInstanceRequest{Name: name}
							req.RedisVersion = cmd.String("redis-version")
							op, err := client.UpgradeInstance(ctx, req)
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
						Name:  "import",
						Usage: "import instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing import...")
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export instances",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
							return nil
						},
					},
					{
						Name:  "failover",
						Usage: "failover instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/instances/%s", cmd.String("project"), cmd.String("location"), cmd.String("instance"))
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.FailoverInstanceRequest{Name: name}
							op, err := client.FailoverInstance(ctx, req)
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
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.DeleteInstanceRequest{Name: name}
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
							client, err := redis.NewCloudRedisClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &redispb.RescheduleMaintenanceRequest{Name: name}
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
							client, err := redis.NewCloudRedisClient(ctx)
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
							client, err := redis.NewCloudRedisClient(ctx)
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
							client, err := redis.NewCloudRedisClient(ctx)
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
							client, err := redis.NewCloudRedisClient(ctx)
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
							client, err := redis.NewCloudRedisClient(ctx)
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
							client, err := redis.NewCloudRedisClient(ctx)
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
		},
	}
}
