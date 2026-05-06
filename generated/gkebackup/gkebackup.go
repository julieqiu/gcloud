package gkebackup

import (
	"context"
	"fmt"
	"strings"

	gkebackup "cloud.google.com/go/gkebackup/apiv1"
	"cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud gkebackup command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkebackup",
		Usage: "manage Backup for GKE API resources",
		Commands: []*cli.Command{
			{
				Name:  "backup-channels",
				Usage: "Manage backup-channels resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup-channel-id", Usage: "The backup channel id.", Required: false},
							&cli.StringFlag{Name: "destination-project", Usage: "The destination project.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.CreateBackupChannelRequest{Parent: parent}
							req.BackupChannelId = cmd.String("backup-channel-id")
							req.BackupChannel = &gkebackuppb.BackupChannel{
								DestinationProject: cmd.String("destination-project"),
								Description:        cmd.String("description"),
							}
							op, err := client.CreateBackupChannel(ctx, req)
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
						Name:  "list",
						Usage: "list backup-channels",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_channel", Usage: "The backup_channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_channel"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetBackupChannelRequest{Name: name}
							resp, err := client.GetBackupChannel(ctx, req)
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
						Usage: "update backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_channel", Usage: "The backup_channel.", Required: true},
							&cli.StringFlag{Name: "destination-project", Usage: "The destination project.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_channel"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.UpdateBackupChannelRequest{}
							req.BackupChannel = &gkebackuppb.BackupChannel{
								Name:               name,
								DestinationProject: cmd.String("destination-project"),
								Description:        cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("destination-project") {
								paths = append(paths, "destination_project")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateBackupChannel(ctx, req)
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
						Usage: "delete backup-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_channel", Usage: "The backup_channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_channel"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.DeleteBackupChannelRequest{Name: name}
							op, err := client.DeleteBackupChannel(ctx, req)
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
				Name:  "backup-plan-bindings",
				Usage: "Manage backup-plan-bindings resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list backup-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkebackuppb.ListBackupPlanBindingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListBackupPlanBindings(ctx, req)
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
						Usage: "describe backup-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_channel", Usage: "The backup_channel.", Required: true},
							&cli.StringFlag{Name: "backup_plan_binding", Usage: "The backup_plan_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s/backupPlanBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_channel"), cmd.String("backup_plan_binding"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetBackupPlanBindingRequest{Name: name}
							resp, err := client.GetBackupPlanBinding(ctx, req)
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
				Name:  "backup-plans",
				Usage: "Manage backup-plans resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup-plan-id", Usage: "The backup plan id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
							&cli.BoolFlag{Name: "deactivated", Usage: "The deactivated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.CreateBackupPlanRequest{Parent: parent}
							req.BackupPlanId = cmd.String("backup-plan-id")
							req.BackupPlan = &gkebackuppb.BackupPlan{
								Description: cmd.String("description"),
								Cluster:     cmd.String("cluster"),
								Deactivated: cmd.Bool("deactivated"),
							}
							op, err := client.CreateBackupPlan(ctx, req)
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
						Name:  "list",
						Usage: "list backup-plans",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetBackupPlanRequest{Name: name}
							resp, err := client.GetBackupPlan(ctx, req)
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
						Usage: "update backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: false},
							&cli.BoolFlag{Name: "deactivated", Usage: "The deactivated.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.UpdateBackupPlanRequest{}
							req.BackupPlan = &gkebackuppb.BackupPlan{
								Name:        name,
								Description: cmd.String("description"),
								Cluster:     cmd.String("cluster"),
								Deactivated: cmd.Bool("deactivated"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("cluster") {
								paths = append(paths, "cluster")
							}
							if cmd.IsSet("deactivated") {
								paths = append(paths, "deactivated")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateBackupPlan(ctx, req)
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
						Usage: "delete backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.DeleteBackupPlanRequest{Name: name}
							op, err := client.DeleteBackupPlan(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backupPlan", Usage: "The backupPlan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backupPlan"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backupPlan", Usage: "The backupPlan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backupPlan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions backup-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backupPlan", Usage: "The backupPlan.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backupPlan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
			{
				Name:  "backups",
				Usage: "Manage backups resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.StringFlag{Name: "backup-id", Usage: "The backup id.", Required: false},
							&cli.IntFlag{Name: "delete-lock-days", Usage: "The delete lock days.", Required: false},
							&cli.IntFlag{Name: "retain-days", Usage: "The retain days.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.CreateBackupRequest{Parent: parent}
							req.BackupId = cmd.String("backup-id")
							req.Backup = &gkebackuppb.Backup{
								DeleteLockDays: int32(cmd.Int("delete-lock-days")),
								RetainDays:     int32(cmd.Int("retain-days")),
								Description:    cmd.String("description"),
							}
							op, err := client.CreateBackup(ctx, req)
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
						Name:  "list",
						Usage: "list backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkebackuppb.ListBackupsRequest{Parent: parent}
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
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"), cmd.String("backup"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetBackupRequest{Name: name}
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
						Name:  "update",
						Usage: "update backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
							&cli.IntFlag{Name: "delete-lock-days", Usage: "The delete lock days.", Required: false},
							&cli.IntFlag{Name: "retain-days", Usage: "The retain days.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"), cmd.String("backup"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.UpdateBackupRequest{}
							req.Backup = &gkebackuppb.Backup{
								Name:           name,
								DeleteLockDays: int32(cmd.Int("delete-lock-days")),
								RetainDays:     int32(cmd.Int("retain-days")),
								Description:    cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("delete-lock-days") {
								paths = append(paths, "delete_lock_days")
							}
							if cmd.IsSet("retain-days") {
								paths = append(paths, "retain_days")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateBackup(ctx, req)
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
						Usage: "delete backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"), cmd.String("backup"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.DeleteBackupRequest{Name: name}
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
						Name:  "describe",
						Usage: "describe backups",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing describe...")
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
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
							client, err := gkebackup.NewBackupForGKEClient(ctx)
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
				Name:  "restore-channels",
				Usage: "Manage restore-channels resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore-channel-id", Usage: "The restore channel id.", Required: false},
							&cli.StringFlag{Name: "destination-project", Usage: "The destination project.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.CreateRestoreChannelRequest{Parent: parent}
							req.RestoreChannelId = cmd.String("restore-channel-id")
							req.RestoreChannel = &gkebackuppb.RestoreChannel{
								DestinationProject: cmd.String("destination-project"),
								Description:        cmd.String("description"),
							}
							op, err := client.CreateRestoreChannel(ctx, req)
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
						Name:  "list",
						Usage: "list restore-channels",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_channel", Usage: "The restore_channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_channel"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetRestoreChannelRequest{Name: name}
							resp, err := client.GetRestoreChannel(ctx, req)
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
						Usage: "update restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_channel", Usage: "The restore_channel.", Required: true},
							&cli.StringFlag{Name: "destination-project", Usage: "The destination project.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_channel"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.UpdateRestoreChannelRequest{}
							req.RestoreChannel = &gkebackuppb.RestoreChannel{
								Name:               name,
								DestinationProject: cmd.String("destination-project"),
								Description:        cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("destination-project") {
								paths = append(paths, "destination_project")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateRestoreChannel(ctx, req)
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
						Usage: "delete restore-channels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_channel", Usage: "The restore_channel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_channel"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.DeleteRestoreChannelRequest{Name: name}
							op, err := client.DeleteRestoreChannel(ctx, req)
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
				Name:  "restore-plan-bindings",
				Usage: "Manage restore-plan-bindings resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list restore-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkebackuppb.ListRestorePlanBindingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListRestorePlanBindings(ctx, req)
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
						Usage: "describe restore-plan-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_channel", Usage: "The restore_channel.", Required: true},
							&cli.StringFlag{Name: "restore_plan_binding", Usage: "The restore_plan_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s/restorePlanBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_channel"), cmd.String("restore_plan_binding"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetRestorePlanBindingRequest{Name: name}
							resp, err := client.GetRestorePlanBinding(ctx, req)
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
				Name:  "restore-plans",
				Usage: "Manage restore-plans resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore-plan-id", Usage: "The restore plan id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "backup-plan", Usage: "The backup plan.", Required: true},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.CreateRestorePlanRequest{Parent: parent}
							req.RestorePlanId = cmd.String("restore-plan-id")
							req.RestorePlan = &gkebackuppb.RestorePlan{
								Description: cmd.String("description"),
								BackupPlan:  cmd.String("backup-plan"),
								Cluster:     cmd.String("cluster"),
							}
							op, err := client.CreateRestorePlan(ctx, req)
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
						Name:  "list",
						Usage: "list restore-plans",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetRestorePlanRequest{Name: name}
							resp, err := client.GetRestorePlan(ctx, req)
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
						Usage: "update restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "backup-plan", Usage: "The backup plan.", Required: false},
							&cli.StringFlag{Name: "cluster", Usage: "The cluster.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.UpdateRestorePlanRequest{}
							req.RestorePlan = &gkebackuppb.RestorePlan{
								Name:        name,
								Description: cmd.String("description"),
								BackupPlan:  cmd.String("backup-plan"),
								Cluster:     cmd.String("cluster"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("backup-plan") {
								paths = append(paths, "backup_plan")
							}
							if cmd.IsSet("cluster") {
								paths = append(paths, "cluster")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateRestorePlan(ctx, req)
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
						Usage: "delete restore-plans",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.DeleteRestorePlanRequest{Name: name}
							op, err := client.DeleteRestorePlan(ctx, req)
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
				Name:  "restores",
				Usage: "Manage restores resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.StringFlag{Name: "restore-id", Usage: "The restore id.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.CreateRestoreRequest{Parent: parent}
							req.RestoreId = cmd.String("restore-id")
							req.Restore = &gkebackuppb.Restore{
								Description: cmd.String("description"),
								Backup:      cmd.String("backup"),
							}
							op, err := client.CreateRestore(ctx, req)
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
						Name:  "list",
						Usage: "list restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkebackuppb.ListRestoresRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListRestores(ctx, req)
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
						Usage: "describe restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The restore.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"), cmd.String("restore"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetRestoreRequest{Name: name}
							resp, err := client.GetRestore(ctx, req)
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
						Usage: "update restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The restore.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"), cmd.String("restore"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.UpdateRestoreRequest{}
							req.Restore = &gkebackuppb.Restore{
								Name:        name,
								Description: cmd.String("description"),
								Backup:      cmd.String("backup"),
							}
							var paths []string
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("backup") {
								paths = append(paths, "backup")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateRestore(ctx, req)
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
						Usage: "delete restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The restore.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"), cmd.String("restore"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.DeleteRestoreRequest{Name: name}
							op, err := client.DeleteRestore(ctx, req)
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
				Name:  "volume-backups",
				Usage: "Manage volume-backups resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list volume-backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkebackuppb.ListVolumeBackupsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListVolumeBackups(ctx, req)
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
						Usage: "describe volume-backups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backup_plan", Usage: "The backup_plan.", Required: true},
							&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
							&cli.StringFlag{Name: "volume_backup", Usage: "The volume_backup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s/volumeBackups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_plan"), cmd.String("backup"), cmd.String("volume_backup"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetVolumeBackupRequest{Name: name}
							resp, err := client.GetVolumeBackup(ctx, req)
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
				Name:  "volume-restores",
				Usage: "Manage volume-restores resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list volume-restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkebackuppb.ListVolumeRestoresRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListVolumeRestores(ctx, req)
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
						Usage: "describe volume-restores",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "restore_plan", Usage: "The restore_plan.", Required: true},
							&cli.StringFlag{Name: "restore", Usage: "The restore.", Required: true},
							&cli.StringFlag{Name: "volume_restore", Usage: "The volume_restore.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s/volumeRestores/%s", cmd.String("project"), cmd.String("location"), cmd.String("restore_plan"), cmd.String("restore"), cmd.String("volume_restore"))
							client, err := gkebackup.NewBackupForGKEClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkebackuppb.GetVolumeRestoreRequest{Name: name}
							resp, err := client.GetVolumeRestore(ctx, req)
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
