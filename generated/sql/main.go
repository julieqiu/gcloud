package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	sql "cloud.google.com/go/sql/apiv1"
	"cloud.google.com/go/sql/apiv1/sqlpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "sqladmin",
				Usage: "manage Cloud SQL Admin API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "acquire-ssrs-lease",
						Usage: "Manage acquire-ssrs-lease resources",
						Commands: []*cli.Command{
							{
								Name:  "acquire-ssrs-lease",
								Usage: "acquire-ssrs-lease acquire-ssrs-lease",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing acquire-ssrs-lease...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-entra-id-certificate",
						Usage: "Manage add-entra-id-certificate resources",
						Commands: []*cli.Command{
							{
								Name:  "add-entra-id-certificate",
								Usage: "add-entra-id-certificate add-entra-id-certificate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-entra-id-certificate...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-server-ca",
						Usage: "Manage add-server-ca resources",
						Commands: []*cli.Command{
							{
								Name:  "add-server-ca",
								Usage: "add-server-ca add-server-ca",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-server-ca...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-server-certificate",
						Usage: "Manage add-server-certificate resources",
						Commands: []*cli.Command{
							{
								Name:  "add-server-certificate",
								Usage: "add-server-certificate add-server-certificate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-server-certificate...")
									return nil
								},
							},
						},
					},
					{
						Name:  "backup-runs",
						Usage: "Manage backup-runs resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete backup-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe backup-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert backup-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list backup-runs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: false},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := sql.NewSqlBackupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &sqlpb.CreateBackupRequest{Parent: parent}
									req.Backup = &sqlpb.Backup{
										Description: cmd.String("description"),
										Instance:    cmd.String("instance"),
										Location:    cmd.String("location"),
									}
									resp, err := client.CreateBackup(ctx, req)
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
								Usage: "describe backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/backups/%s", cmd.String("project"), cmd.String("backup"))
									client, err := sql.NewSqlBackupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &sqlpb.GetBackupRequest{Name: name}
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
								Name:  "list",
								Usage: "list backups",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := sql.NewSqlBackupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &sqlpb.ListBackupsRequest{Parent: parent}
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
								Name:  "update",
								Usage: "update backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "instance", Usage: "The instance.", Required: false},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/backups/%s", cmd.String("project"), cmd.String("backup"))
									client, err := sql.NewSqlBackupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &sqlpb.UpdateBackupRequest{}
									req.Backup = &sqlpb.Backup{
										Name:        name,
										Description: cmd.String("description"),
										Instance:    cmd.String("instance"),
										Location:    cmd.String("location"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("instance") {
										paths = append(paths, "instance")
									}
									if cmd.IsSet("location") {
										paths = append(paths, "location")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateBackup(ctx, req)
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
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/backups/%s", cmd.String("project"), cmd.String("backup"))
									client, err := sql.NewSqlBackupsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &sqlpb.DeleteBackupRequest{Name: name}
									if err := client.DeleteBackup(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "cancel",
						Usage: "Manage cancel resources",
						Commands: []*cli.Command{
							{
								Name:  "cancel",
								Usage: "cancel cancel",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
									return nil
								},
							},
						},
					},
					{
						Name:  "clone",
						Usage: "Manage clone resources",
						Commands: []*cli.Command{
							{
								Name:  "clone",
								Usage: "clone clone",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing clone...")
									return nil
								},
							},
						},
					},
					{
						Name:  "connect-settings",
						Usage: "Manage connect-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe connect-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "create-ephemeral",
						Usage: "Manage create-ephemeral resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create create-ephemeral",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
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
								Name:  "delete",
								Usage: "delete databases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe databases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert databases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list databases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch databases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update databases",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "demote",
						Usage: "Manage demote resources",
						Commands: []*cli.Command{
							{
								Name:  "demote",
								Usage: "demote demote",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing demote...")
									return nil
								},
							},
						},
					},
					{
						Name:  "demote-master",
						Usage: "Manage demote-master resources",
						Commands: []*cli.Command{
							{
								Name:  "demote-master",
								Usage: "demote-master demote-master",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing demote-master...")
									return nil
								},
							},
						},
					},
					{
						Name:  "execute-sql",
						Usage: "Manage execute-sql resources",
						Commands: []*cli.Command{
							{
								Name:  "execute-sql",
								Usage: "execute-sql execute-sql",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing execute-sql...")
									return nil
								},
							},
						},
					},
					{
						Name:  "export",
						Usage: "Manage export resources",
						Commands: []*cli.Command{
							{
								Name:  "export",
								Usage: "export export",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing export...")
									return nil
								},
							},
						},
					},
					{
						Name:  "failover",
						Usage: "Manage failover resources",
						Commands: []*cli.Command{
							{
								Name:  "failover",
								Usage: "failover failover",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing failover...")
									return nil
								},
							},
						},
					},
					{
						Name:  "flags",
						Usage: "Manage flags resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list flags",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-disk-shrink-config",
						Usage: "Manage get-disk-shrink-config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-disk-shrink-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-latest-recovery-time",
						Usage: "Manage get-latest-recovery-time resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-latest-recovery-time",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "import",
						Usage: "Manage import resources",
						Commands: []*cli.Command{
							{
								Name:  "import",
								Usage: "import import",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing import...")
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
								Name:  "generate-ephemeral-cert",
								Usage: "generate-ephemeral-cert instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing generate-ephemeral-cert...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-entra-id-certificates",
						Usage: "Manage list-entra-id-certificates resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-entra-id-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-server-cas",
						Usage: "Manage list-server-cas resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-server-cas",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-server-certificates",
						Usage: "Manage list-server-certificates resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-server-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
								Name:  "describe",
								Usage: "describe operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "perform-disk-shrink",
						Usage: "Manage perform-disk-shrink resources",
						Commands: []*cli.Command{
							{
								Name:  "perform-disk-shrink",
								Usage: "perform-disk-shrink perform-disk-shrink",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing perform-disk-shrink...")
									return nil
								},
							},
						},
					},
					{
						Name:  "pre-check-major-version-upgrade",
						Usage: "Manage pre-check-major-version-upgrade resources",
						Commands: []*cli.Command{
							{
								Name:  "pre-check-major-version-upgrade",
								Usage: "pre-check-major-version-upgrade pre-check-major-version-upgrade",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing pre-check-major-version-upgrade...")
									return nil
								},
							},
						},
					},
					{
						Name:  "projects",
						Usage: "Manage projects resources",
						Commands: []*cli.Command{
							{
								Name:  "point-in-time-restore",
								Usage: "point-in-time-restore projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing point-in-time-restore...")
									return nil
								},
							},
						},
					},
					{
						Name:  "promote-replica",
						Usage: "Manage promote-replica resources",
						Commands: []*cli.Command{
							{
								Name:  "promote-replica",
								Usage: "promote-replica promote-replica",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing promote-replica...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reencrypt",
						Usage: "Manage reencrypt resources",
						Commands: []*cli.Command{
							{
								Name:  "reencrypt",
								Usage: "reencrypt reencrypt",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reencrypt...")
									return nil
								},
							},
						},
					},
					{
						Name:  "release-ssrs-lease",
						Usage: "Manage release-ssrs-lease resources",
						Commands: []*cli.Command{
							{
								Name:  "release-ssrs-lease",
								Usage: "release-ssrs-lease release-ssrs-lease",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing release-ssrs-lease...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reschedule-maintenance",
						Usage: "Manage reschedule-maintenance resources",
						Commands: []*cli.Command{
							{
								Name:  "reschedule-maintenance",
								Usage: "reschedule-maintenance reschedule-maintenance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reschedule-maintenance...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reset-replica-size",
						Usage: "Manage reset-replica-size resources",
						Commands: []*cli.Command{
							{
								Name:  "reset-replica-size",
								Usage: "reset-replica-size reset-replica-size",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reset-replica-size...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reset-ssl-config",
						Usage: "Manage reset-ssl-config resources",
						Commands: []*cli.Command{
							{
								Name:  "reset-ssl-config",
								Usage: "reset-ssl-config reset-ssl-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reset-ssl-config...")
									return nil
								},
							},
						},
					},
					{
						Name:  "restart",
						Usage: "Manage restart resources",
						Commands: []*cli.Command{
							{
								Name:  "restart",
								Usage: "restart restart",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing restart...")
									return nil
								},
							},
						},
					},
					{
						Name:  "restore-backup",
						Usage: "Manage restore-backup resources",
						Commands: []*cli.Command{
							{
								Name:  "restore-backup",
								Usage: "restore-backup restore-backup",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing restore-backup...")
									return nil
								},
							},
						},
					},
					{
						Name:  "rotate-entra-id-certificate",
						Usage: "Manage rotate-entra-id-certificate resources",
						Commands: []*cli.Command{
							{
								Name:  "rotate-entra-id-certificate",
								Usage: "rotate-entra-id-certificate rotate-entra-id-certificate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing rotate-entra-id-certificate...")
									return nil
								},
							},
						},
					},
					{
						Name:  "rotate-server-ca",
						Usage: "Manage rotate-server-ca resources",
						Commands: []*cli.Command{
							{
								Name:  "rotate-server-ca",
								Usage: "rotate-server-ca rotate-server-ca",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing rotate-server-ca...")
									return nil
								},
							},
						},
					},
					{
						Name:  "rotate-server-certificate",
						Usage: "Manage rotate-server-certificate resources",
						Commands: []*cli.Command{
							{
								Name:  "rotate-server-certificate",
								Usage: "rotate-server-certificate rotate-server-certificate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing rotate-server-certificate...")
									return nil
								},
							},
						},
					},
					{
						Name:  "ssl-certs",
						Usage: "Manage ssl-certs resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete ssl-certs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ssl-certs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert ssl-certs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list ssl-certs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "start-external-sync",
						Usage: "Manage start-external-sync resources",
						Commands: []*cli.Command{
							{
								Name:  "start-external-sync",
								Usage: "start-external-sync start-external-sync",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-external-sync...")
									return nil
								},
							},
						},
					},
					{
						Name:  "start-replica",
						Usage: "Manage start-replica resources",
						Commands: []*cli.Command{
							{
								Name:  "start-replica",
								Usage: "start-replica start-replica",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-replica...")
									return nil
								},
							},
						},
					},
					{
						Name:  "stop-replica",
						Usage: "Manage stop-replica resources",
						Commands: []*cli.Command{
							{
								Name:  "stop-replica",
								Usage: "stop-replica stop-replica",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-replica...")
									return nil
								},
							},
						},
					},
					{
						Name:  "switchover",
						Usage: "Manage switchover resources",
						Commands: []*cli.Command{
							{
								Name:  "switchover",
								Usage: "switchover switchover",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing switchover...")
									return nil
								},
							},
						},
					},
					{
						Name:  "tiers",
						Usage: "Manage tiers resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tiers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "truncate-log",
						Usage: "Manage truncate-log resources",
						Commands: []*cli.Command{
							{
								Name:  "truncate-log",
								Usage: "truncate-log truncate-log",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing truncate-log...")
									return nil
								},
							},
						},
					},
					{
						Name:  "users",
						Usage: "Manage users resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete users",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe users",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert users",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list users",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update users",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "verify-external-sync-settings",
						Usage: "Manage verify-external-sync-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "verify-external-sync-settings",
								Usage: "verify-external-sync-settings verify-external-sync-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing verify-external-sync-settings...")
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
